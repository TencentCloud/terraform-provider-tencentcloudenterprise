/*
Provides a csp resource to create a Csp bucket policy and set its attributes.

# Example Usage

```hcl

	resource "tencentcloudenterprise_csp_bucket_policy" "csp_policy" {
	  bucket = "mycsp-1258798060"

	  policy = <<EOF

	{
	  "version": "2.0",
	  "Statement": [
	    {
	      "Principal": {
	        "qcs": [
	          "qcs::cam::uin/<your-account-id>:uin/<your-account-id>"
	        ]
	      },
	      "Action": [
	        "name/csp:DeleteBucket",
	        "name/csp:PutBucketACL"
	      ],
	      "Effect": "allow",
	      "Resource": [
	        "qcs::csp:<bucket region>:uid/<your-account-id>:<bucket name>/*"
	      ]
	    }
	  ]
	}

EOF
}
```

# Import

csp bucket policy can be imported, e.g.

```
$ terraform import cloud_csp_bucket_policy.bucket bucket-name
```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_csp_bucket_policy", CNDescription{
		TerraformTypeCN: "添加policy策略",
		DescriptionCN:   "提供CSP存储桶策略资源，用于创建和管理CSP存储桶的访问策略。",
		AttributesCN: map[string]string{
			"bucket": "存储桶名称",
			"policy": "存储桶策略,支持 json 格式",
		},
	})
}

func resourceTencentCloudCspBucketPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a csp resource to create a Csp bucket policy and set its attributes.",
		Create:      resourceTencentCloudCspBucketPolicyCreate,
		Read:        resourceTencentCloudCspBucketPolicyRead,
		Update:      resourceTencentCloudCspBucketPolicyUpdate,
		Delete:      resourceTencentCloudCspBucketPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateCosBucketName,
				Description:  "The name of a bucket to be created. Bucket format should be [custom name]-[appid], for example `mycsp-1258798060`.",
			},
			"policy": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
				DiffSuppressFunc: func(k, olds, news string, d *schema.ResourceData) bool {
					var oldJson interface{}
					err := json.Unmarshal([]byte(olds), &oldJson)
					if err != nil {
						return olds == news
					}
					var newJson interface{}
					err = json.Unmarshal([]byte(news), &newJson)
					if err != nil {
						return olds == news
					}
					flag := reflect.DeepEqual(oldJson, newJson)
					return flag
				},
				Description: "The text of the policy. For more info please refer to [Tencent official doc](https://intl.cloud.tencent.com/document/product/436/18023).",
			},
		},
	}
}

func resourceTencentCloudCspBucketPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_policy.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	bucket := d.Get("bucket").(string)
	policy := d.Get("policy").(string)

	cosService := CosService{
		client:       meta.(*TencentCloudClient).apiV3Conn,
		useCspClient: true,
	}
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	policyErr := camService.PolicyDocumentForceCheck(policy)
	if policyErr != nil {
		return policyErr
	}

	err := cosService.PutBucketPolicy(ctx, bucket, policy)
	if err != nil {
		return err
	}
	d.SetId(bucket)

	return resourceTencentCloudCspBucketPolicyRead(d, meta)
}

func resourceTencentCloudCspBucketPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_policy.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Id()
	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}

	var result string
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		policy, e := cosService.DescribePolicyByBucket(ctx, bucket)
		if e != nil {
			return retryError(e)
		}
		result = policy
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read csp bucket policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	result, err = removeSid(result)
	if err != nil {
		log.Printf("[CRITAL]%s read csp bucket policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	if result == "" {
		d.SetId("")
		return nil
	}
	_ = d.Set("policy", result)
	_ = d.Set("bucket", bucket)

	return nil
}

func resourceTencentCloudCspBucketPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_policy.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cosService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}
	bucket := d.Id()

	if d.HasChange("policy") {
		policy := d.Get("policy").(string)
		camService := CamService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		policyErr := camService.PolicyDocumentForceCheck(policy)
		if policyErr != nil {
			return policyErr
		}
		err := cosService.PutBucketPolicy(ctx, bucket, policy)
		if err != nil {
			return err
		}
	}

	// wait for update cache
	// if not, the data may be outdated.
	time.Sleep(3 * time.Second)

	return resourceTencentCloudCspBucketPolicyRead(d, meta)
}

func resourceTencentCloudCspBucketPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_csp_bucket_policy.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	bucket := d.Id()
	cosService := CosService{
		client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true,
	}
	err := cosService.DeleteBucketPolicy(ctx, bucket)
	if err != nil {
		return err
	}

	// wait for update cache
	// if not, head bucket may be successful
	time.Sleep(3 * time.Second)

	return nil
}
