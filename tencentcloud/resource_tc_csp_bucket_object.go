/*
Provides a csp object resource to put an object(content or file) to the bucket.

# Example Usage

# Uploading a file to a bucket

```hcl

	resource "tencentcloudenterprise_csp_bucket_object" "myobject" {
	  bucket = "mycsp-1258798060"
	  key    = "new_object_key"
	  acl    = "public-read"
	  source = "path/to/file"
	}

```

# Uploading a content to a bucket

```hcl

	resource "tencentcloudenterprise_csp_bucket" "mycsp" {
	  bucket = "mycsp-1258798060"
	  acl    = "public-read"
	}

	resource "tencentcloudenterprise_csp_bucket_object" "myobject" {
	  bucket  = cloud_csp_bucket.mycsp.bucket
	  key     = "new_object_key"
	  content = "the content that you want to upload."
	}

```
*/
package tencentcloud

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
	"io"
	"log"
	"os"
	"time"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_csp_bucket_object", CNDescription{
		TerraformTypeCN: "上传存储对象",
		DescriptionCN:   "提供CSP对象资源，用于将对象（内容或文件）上传到存储桶。",
		AttributesCN: map[string]string{
			"bucket":              "存储桶名称",
			"key":                 "对象名称",
			"source":              "对象源文件路径，请勿填写，可视化模版不支持设置源文件路径",
			"content":             "对象内容，需要填写，可以输入需要上传的纯文本内容",
			"acl":                 "对象访问权限控制, 可取值包括private(私有读)和public-read(公有读)",
			"cache_control":       "缓存控制",
			"content_disposition": "内容描述",
			"content_encoding":    "内容编码",
			"content_type":        "内容类型",
			"tags":                "标签",
			// "storage_class":       "存储类型",
			// "etag":                "对象的ETag",
		},
	})
}

func resourceTencentCloudCspBucketObject() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a csp object resource to put an object(content or file) to the bucket.",
		Create:      resourceTencentCloudCspBucketObjectCreate,
		Read:        resourceTencentCloudCspBucketObjectRead,
		Update:      resourceTencentCloudCspBucketObjectUpdate,
		Delete:      resourceTencentCloudCspBucketObjectDelete,

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of a bucket to use. Bucket format should be [custom name]-[appid], for example `mycsp-1258798060`.",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the object once it is in the bucket.",
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
				Description:   "The path to the source file being uploaded to the bucket.",
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
				Description:   "Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.",
			},
			"acl": {
				Type:     schema.TypeString,
				Optional: true,
				// 同步csp bucket acl
				Default: COS_BUCKET_OBJECT_ACL,
				ValidateFunc: validateAllowedStringValue([]string{
					COS_BUCKET_OBJECT_ACL,
					s3.ObjectCannedACLPrivate,
					s3.ObjectCannedACLPublicRead,
				}),
				Description: "The canned ACL to apply. Available values include `private` and `public-read`.",
			},
			"cache_control": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Specifies caching behavior along the request/reply chain. For further details, RFC2616 can be referred.",
			},
			"content_disposition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies presentational information for the object.",
			},
			"content_encoding": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag of the object.",
			},
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A standard MIME type describing the format of the object data.",
			},
			//"storage_class": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "Object storage type, Available values include `STANDARD_IA`, `MAZ_STANDARD_IA`, `INTELLIGENT_TIERING`, `MAZ_INTELLIGENT_TIERING`, `ARCHIVE`, `DEEP_ARCHIVE`. For more information, please refer to: https://cloud.tencent.com/document/product/436/33417.",
			//},
			//"etag": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "The ETag generated for the object (an MD5 sum of the object content).",
			//},
		},
	}
}

func resourceTencentCloudCspBucketObjectCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_object.create")()

	logId := getLogId(contextNil)

	bucket := d.Get("bucket").(string)
	key := d.Get("key").(string)
	var body io.ReadSeeker
	if v, ok := d.GetOk("source"); ok {
		source := v.(string)
		path, err := homedir.Expand(source)
		if err != nil {
			return fmt.Errorf("csp object source (%s) homedir expand error: %s", source, err.Error())
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("csp object source (%s) open error: %s", source, err.Error())
		}
		body = file
		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("closing csp object source (%s) error: %s", path, err.Error())
			}
		}()
	} else if v, ok := d.GetOk("content"); ok {
		content := v.(string)
		body = bytes.NewReader([]byte(content))
	} else {
		return fmt.Errorf("must specify \"source\" or \"content\" field")
	}

	request := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	}

	if v, ok := d.GetOk("acl"); ok {
		request.ACL = aws.String(v.(string))
	}
	if v, ok := d.GetOk("cache_control"); ok {
		request.CacheControl = aws.String(v.(string))
	}
	if v, ok := d.GetOk("content_disposition"); ok {
		request.ContentDisposition = aws.String(v.(string))
	}
	if v, ok := d.GetOk("content_encoding"); ok {
		request.ContentEncoding = aws.String(v.(string))
	}
	if v, ok := d.GetOk("content_type"); ok {
		request.ContentType = aws.String(v.(string))
	}
	//if v, ok := d.GetOk("storage_class"); ok {
	//	request.StorageClass = aws.String(v.(string))
	//}

	response, err := meta.(*TencentCloudClient).apiV3Conn.UseCosS3Client(true).PutObject(request)
	if err != nil {
		return fmt.Errorf("putting object (%s) in csp bucket (%s) error: %s", key, bucket, err.Error())
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "put object", request.String(), response.String())

	if v, ok := d.GetOk("tags"); ok {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		service := CosService{
			client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true,
		}
		tags := make(map[string]string)

		for key, val := range v.(map[string]interface{}) {
			tags[key] = val.(string)
		}

		if err := service.SetObjectTags(ctx, bucket, key, tags); err != nil {
			log.Printf("[WARN] set object tags error, skip processing")
		}
	}

	d.SetId(bucket + key)
	cosService := CosService{
		client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true,
	}
	if err = WaitCspBucketObjectCreated(contextNil, &cosService, bucket, key); err != nil {
		return err
	}
	return resourceTencentCloudCspBucketObjectRead(d, meta)
}

func resourceTencentCloudCspBucketObjectRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_object.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Get("bucket").(string)
	key := d.Get("key").(string)

	cosService := CosService{
		client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true,
	}
	response, err := cosService.HeadObject(ctx, bucket, key)
	if err != nil {
		if awsError, ok := err.(awserr.RequestFailure); ok && awsError.StatusCode() == 404 {
			log.Printf("[WARN]%s object (%s) in bucket (%s) not found, error code (404)", logId, key, bucket)
			d.SetId("")
			return nil
		}
		return err
	}

	_ = d.Set("cache_control", response.CacheControl)
	_ = d.Set("content_disposition", response.ContentDisposition)
	_ = d.Set("content_encoding", response.ContentEncoding)
	_ = d.Set("content_type", response.ContentType)
	//_ = d.Set("etag", strings.Trim(*response.ETag, `"`))
	//_ = d.Set("storage_class", s3.StorageClassStandard)
	//if response.StorageClass != nil {
	//	_ = d.Set("storage_class", response.StorageClass)
	//}

	var tags map[string]string
	tags, err = cosService.GetObjectTags(ctx, bucket, key)
	if err != nil {
		if awsError, ok := err.(awserr.RequestFailure); ok && awsError.StatusCode() == 404 {
			log.Printf("[WARN]%s tags in object (%s) of bucket (%s) not found, error code (404)", logId, key, bucket)
			d.SetId("")
			return nil
		}
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudCspBucketObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_object.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fields := []string{
		"cache_control",
		"content_disposition",
		"content_encoding",
		"content_type",
		"source",
		"content",
		// "storage_class",
		// "etag",
	}
	for _, key := range fields {
		if d.HasChange(key) {
			return resourceTencentCloudCspBucketObjectCreate(d, meta)
		}
	}

	bucket := d.Get("bucket").(string)
	key := d.Get("key").(string)
	cosService := CosService{
		client:       meta.(*TencentCloudClient).apiV3Conn,
		useCspClient: true,
	}

	if d.HasChange("acl") {
		acl := d.Get("acl").(string)
		err := cosService.PutObjectAcl(ctx, bucket, key, acl)
		if err != nil {
			return err
		}
	}

	if d.HasChange("tags") {
		v := d.Get("tags").(map[string]interface{})
		tags := make(map[string]string)
		for key, val := range v {
			tags[key] = val.(string)
		}
		if err := cosService.SetObjectTags(ctx, bucket, key, tags); err != nil {
			return err
		}
	}

	return nil
}

func resourceTencentCloudCspBucketObjectDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_object.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Get("bucket").(string)
	key := d.Get("key").(string)

	cosService := CosService{
		client:       meta.(*TencentCloudClient).apiV3Conn,
		useCspClient: true,
	}
	err := cosService.DeleteObject(ctx, bucket, key)
	if err != nil {
		return err
	}

	return nil
}

// WaitCspBucketObjectCreated wait for the object to be created in the bucket
func WaitCspBucketObjectCreated(ctx context.Context, service *CosService, bucket, key string) error {
	logId := getLogId(ctx)

	for i := 0; i < 10; i++ {
		_, err := service.HeadObject(ctx, bucket, key)
		if err != nil {
			time.Sleep(5 * time.Second)
			if awsError, ok := err.(awserr.RequestFailure); ok && awsError.StatusCode() == 404 {
				log.Printf("[WARN]%s object (%s) in bucket (%s) not found, error code (404)", logId, key, bucket)
				continue
			}
			log.Printf("[error]%s heade object (%s) in bucket (%s) err: %+v", logId, key, bucket, err)
		} else {
			return nil
		}

	}

	return fmt.Errorf("wait created bucket object timeout: %s ", bucket)
}
