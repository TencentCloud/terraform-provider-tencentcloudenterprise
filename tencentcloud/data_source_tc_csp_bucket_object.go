/*
Use this data source to query the metadata of an object stored inside a bucket.

# Example Usage

```hcl

	data "tencentcloudenterprise_csp_bucket_object" "mycsp" {
	  bucket             = "mycos-test-1258798060"
	  prefix                = "hello-world."
	  result_output_file = "TFCspresults"
	}

```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_csp_bucket_object", CNDescription{
		TerraformTypeCN: "获取存储桶对象列表（CSP）",
		DescriptionCN:   "提供CSP存储桶对象数据源，用于查询存储桶对象的元数据。",
		AttributesCN: map[string]string{
			"bucket":             "存储桶名称",
			"prefix":             "对象键匹配前缀",
			"delimiter":          "对象键分隔符",
			"marker":             "起始对象键标记",
			"max_keys":           "单次返回最大的条目数量",
			"name":               "存储桶名称",
			"is_truncated":       "响应条目是否被截断",
			"next_marker":        "下一个标记，仅当响应条目有截断（isTruncated为true）才会返回该节点，该节点的值为当前响应条目中的最后一个对象键，当需要继续请求后续条目时，将该节点的值作为下一次请求的marker参数传入",
			"common_prefixes":    "公共前缀，从prefix或从头（如未指定prefix）到首个delimiter之间的部分，定义为CommonPrefix",
			"contents":           "对象列表",
			"key":                "对象键",
			"last_modified":      "对象最后修改时间",
			"etag":               "对象的实体标签",
			"size":               "对象大小，单位为byte",
			"owner_id":           "对象持有者的APPID",
			"display_name":       "对象持有者的名称",
			"storage_class":      "对象存储类型",
			"result_output_file": "用于保存数据源查询结果, 在前端可视化界面使用时，该参数不可用",
			"encoding_type":      "指定返回值的编码方式，可选值为 url，当指定时，返回的对象键（key）等字段会进行 URL 编码处理，例如原始 Key 为 “key 1.jpg”，编码后会变为 “key%201.jpg”",
		},
	})
}

// dataSourceTencentCloudCspBucketObject defines the schema for the data source
// cloud_csp_bucket_object get bucket object
func dataSourceTencentCloudCspBucketObject() *schema.Resource {
	return &schema.Resource{
		Description:        "Use this data source to query the metadata of an object stored inside a bucket.",
		ReadWithoutTimeout: dataSourceTencentCloudCspBucketObjectsRead,

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the bucket that contains the objects to query.",
			},
			"prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "bucket object key prefix.",
			},
			"delimiter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "bucket object key delimiter.",
			},
			"encoding_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Specifies the encoding method of the returned key. Valid values: url. " +
					"If this parameter is not specified, the returned key is not encoded.",
			},
			"marker": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Specifies the object name to start with when listing objects in a bucket. " +
					"The object name in the response must be greater than this parameter in lexicographic order.",
			},
			"max_keys": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Specifies the maximum number of objects returned in the response. " +
					"The default value is 1000. The maximum value is 1000.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the object name.",
			},
			"is_truncated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the returned object list is truncated.",
			},
			"next_marker": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Specifies the object name to start with when listing objects in a bucket. If the " +
					"returned value of IsTruncated is true, the value of NextMarker is the object " +
					"name that starts with" +
					" the value of Marker in the request. If the returned value of IsTruncated is false, " +
					"the value of NextMarker is null.",
			},
			"common_prefixes": {
				Type:        schema.TypeList,
				Description: "Specifies the object name to start with when listing objects in a bucket.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"prefix": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the object name to start with when listing objects in a bucket.",
						},
					},
				},
			},
			"contents": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies the object name to start with when listing objects in a bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key name to start with when listing objects in a bucket.",
						},
						"last_modified": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modify change.",
						},
						"etag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ETag generated for the object, which is may not equal to MD5 value.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "bucket size.",
						},
						"owner_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account owner id.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account owner name.",
						},
						"storage_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object storage type such as STANDARD.",
						},
					},
				},
			},
			//"cache_control": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Specifies caching behavior along the request/reply chain.",
			//},
			//"content_disposition": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Specifies presentational information for the object.",
			//},
			//"content_encoding": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.",
			//},
			//"content_type": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "A standard MIME type describing the format of the object data.",
			//},
			//"etag": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "ETag generated for the object, which is may not equal to MD5 value.",
			//},
			//"last_modified": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Last modified date of the object.",
			//},
			//"storage_class": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Object storage type such as STANDARD.",
			//},
		},
	}
}

// dataSourceTencentCloudCspBucketObjectsRead reads the metadata of an object stored inside a bucket.
// api3: tcencent cos sdk GetBucket: https://cloud.tencent.com/document/product/436/7734
func dataSourceTencentCloudCspBucketObjectsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	defer logElapsed("data_source.cloud_csp_bucket_object.read")()
	var diags diag.Diagnostics
	bucket := d.Get("bucket").(string)
	getOptions := &cos.BucketGetOptions{}
	v, ok := d.GetOk("prefix")
	if ok {
		prefix := v.(string)
		getOptions.Prefix = prefix
	}
	v, ok = d.GetOk("delimiter")
	if ok {
		delimiter := v.(string)
		getOptions.Delimiter = delimiter
	}
	v, ok = d.GetOk("encoding_type")
	if ok {
		encodingType := v.(string)
		getOptions.EncodingType = encodingType
	}
	v, ok = d.GetOk("marker")
	if ok {
		marker := v.(string)
		getOptions.Marker = marker
	}
	v, ok = d.GetOk("max_keys")
	if ok {
		maxKeys := v.(int)
		getOptions.MaxKeys = maxKeys
	}

	cosService := CosService{
		client:       meta.(*TencentCloudClient).apiV3Conn,
		useCspClient: true,
	}
	info, err := cosService.GetBucketObject(ctx, bucket, getOptions)
	log.Printf("[DEBUG] get bucket options: %+v result: %+v", getOptions, info)
	if err != nil {
		return helper.AppendErrorf(diags, "get bucket object (%s): %s", d.Id(), err)
	}

	ids := []string{bucket}
	d.SetId(helper.DataResourceIdsHash(ids))
	if err = d.Set("name", info.Name); err != nil {
		return helper.AppendErrorf(diags, "set name (%s): %s", d.Id(), err)
	}
	if err = d.Set("encoding_type", info.EncodingType); err != nil {
		return helper.AppendErrorf(diags, "set encoding_type (%s): %s", d.Id(), err)
	}
	if err = d.Set("is_truncated", info.IsTruncated); err != nil {
		return helper.AppendErrorf(diags, "set is_truncated (%s): %s", d.Id(), err)
	}
	if err = d.Set("next_marker", info.NextMarker); err != nil {
		return helper.AppendErrorf(diags, "set next_marker (%s): %s", d.Id(), err)
	}
	commonPrefixes := make([]map[string]interface{}, 0, len(info.CommonPrefixes))
	for _, v := range info.CommonPrefixes {
		prefix := make(map[string]interface{})
		prefix["prefix"] = v
		commonPrefixes = append(commonPrefixes, prefix)
	}
	if err = d.Set("common_prefixes", commonPrefixes); err != nil {
		return helper.AppendErrorf(diags, "set common_prefixes (%s): %s", d.Id(), err)
	}
	contents := make([]map[string]interface{}, 0, len(info.Contents))
	for _, v := range info.Contents {
		content := make(map[string]interface{})
		content["key"] = v.Key
		content["last_modified"] = v.LastModified
		content["etag"] = v.ETag
		content["size"] = v.Size
		content["owner_id"] = v.Owner.ID
		content["display_name"] = v.Owner.DisplayName
		content["storage_class"] = v.StorageClass
		contents = append(contents, content)
	}
	_ = d.Set("contents", contents)
	outputMap := make(map[string]interface{})
	outputMap["bucket"] = bucket
	outputMap["name"] = info.Name
	outputMap["encoding_type"] = info.EncodingType
	outputMap["is_truncated"] = info.IsTruncated
	outputMap["next_marker"] = info.NextMarker
	outputMap["common_prefixes"] = commonPrefixes
	outputMap["contents"] = contents

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), outputMap); err != nil {
			return helper.AppendErrorf(diags, "write to file (%s): %s", d.Id(), err)
		}
	}

	return diags
}
