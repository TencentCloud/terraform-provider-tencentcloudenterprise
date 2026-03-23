/*
Provides a resource to create a CAM policy.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_policy" "foo" {
	  name        = "cam-policy-test"
	  document    = <<EOF

	{
	  "version": "2.0",
	  "statement": [
	    {
	      "action": [
	        "name/sts:AssumeRole"
	      ],
	      "effect": "allow",
	      "resource": [
	        "*"
	      ]
	    }
	  ]
	}

EOF

	  description = "test"
	}

```

# Import

CAM policy can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_policy.foo 26655801
```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"log"
	"reflect"
	"strconv"
	"strings"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_policy", CNDescription{
		TerraformTypeCN: "CAM策略",
		DescriptionCN:   "提供 CAM 策略资源，用于创建和管理自定义策略。",
		AttributesCN: map[string]string{
			"name":        "策略名称",
			"document":    "策略文档",
			"description": "策略描述",
			"type":        "策略类型",
			"create_time": "创建时间",
			"update_time": "更新时间",
		},
	})
}

func resourceTencentCloudCamPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamPolicyCreate,
		Read:   resourceTencentCloudCamPolicyRead,
		Update: resourceTencentCloudCamPolicyUpdate,
		Delete: resourceTencentCloudCamPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of CAM policy.",
			},
			"document": {
				Type:     schema.TypeString,
				Required: true,
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
				Description: "Document of the CAM policy. The syntax refers to [CAM POLICY](https://intl.cloud.tencent.com/document/product/598/10604). There are some notes when using this para in terraform: 1. The elements in JSON claimed supporting two types as `string` and `array` only support type `array`; 2. Terraform does not support the `root` syntax, when it appears, it must be replaced with the uin it stands for.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the CAM policy.",
			},
			"type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Type of the policy strategy. Valid values: `1`, `2`.  `1` means customer strategy and `2` means preset strategy.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the CAM policy.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of the CAM policy.",
			},
		},
	}
}

func resourceTencentCloudCamPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_policy.create")()

	logId := getLogId(contextNil)

	name := d.Get("name").(string)
	document := d.Get("document").(string)

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	documentErr := camService.PolicyDocumentForceCheck(document)
	if documentErr != nil {
		return documentErr
	}
	request := cam.NewCreatePolicyRequest()
	request.PolicyName = &name
	request.PolicyDocument = &document
	// Ensure custom policy mode so description is honored
	request.CreateMode = helper.Uint64(3)
	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	var response *cam.CreatePolicyResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().CreatePolicy(request)
		if e != nil {
			if ee, ok := e.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				//check if read empty
				if strings.Contains(errCode, "PolicyNameInUse") {
					return resource.NonRetryableError(e)
				}
			}
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	if response.Response.PolicyId == nil {
		return fmt.Errorf("CAM policy id is nil")
	}
	d.SetId(strconv.Itoa(int(*response.Response.PolicyId)))

	//get really instance then read
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	policyId := d.Id()

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribePolicyById(ctx, policyId)
		if e != nil {
			return retryError(e)
		}
		if instance == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	//time.Sleep(10 * time.Second)
	return resourceTencentCloudCamPolicyRead(d, meta)
}

func resourceTencentCloudCamPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_policy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	policyId := d.Id()
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var instance *cam.GetPolicyResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribePolicyById(ctx, policyId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM policy failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil || instance.Response == nil || instance.Response.PolicyName == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", *instance.Response.PolicyName)
	//document with special change rule, the `\\/` must be replaced with `/`
	_ = d.Set("document", strings.Replace(*instance.Response.PolicyDocument, "\\/", "/", -1))
	_ = d.Set("create_time", *instance.Response.AddTime)
	_ = d.Set("update_time", *instance.Response.UpdateTime)
	_ = d.Set("type", int(*instance.Response.Type))
	if instance.Response.Description != nil {
		_ = d.Set("description", *instance.Response.Description)
	}
	return nil
}

func resourceTencentCloudCamPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_policy.update")()

	logId := getLogId(contextNil)

	policyId := d.Id()
	policyIdInt, e := strconv.Atoi(policyId)
	if e != nil {
		return e
	}
	policyIdInt64 := uint64(policyIdInt)
	request := cam.NewUpdatePolicyRequest()
	request.PolicyId = &policyIdInt64
	changeFlag := false

	// Always send name; some backends require it even if not changed
	request.PolicyName = helper.String(d.Get("name").(string))

	if d.HasChange("description") {
		request.Description = helper.String(d.Get("description").(string))
		changeFlag = true

	}
	if d.HasChange("name") {
		changeFlag = true
	}

	if d.HasChange("document") {
		document := d.Get("document").(string)
		camService := CamService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		documentErr := camService.PolicyDocumentForceCheck(document)
		if documentErr != nil {
			return documentErr
		}
		request.PolicyDocument = &document
		changeFlag = true

	}
	if changeFlag {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().UpdatePolicy(request)

			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update CAM policy description failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	return resourceTencentCloudCamPolicyRead(d, meta)
}

func resourceTencentCloudCamPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_policy.delete")()

	logId := getLogId(contextNil)

	policyId := d.Id()
	policyIdInt, e := strconv.Atoi(policyId)
	if e != nil {
		return e
	}
	policyIdInt64 := uint64(policyIdInt)
	request := cam.NewDeletePolicyRequest()
	request.PolicyId = []*uint64{&policyIdInt64}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().DeletePolicy(request)
		if e != nil {
			msg := e.Error()
			// tolerate already-missing policies; backend sometimes wraps in ParseJsonError with numeric code 11026
			if strings.Contains(msg, "PolicyIdNotFound") || strings.Contains(msg, "policy id is not exist") || strings.Contains(msg, "11026") {
				return nil
			}
			if ee, ok := e.(*sdkErrors.CloudSDKError); ok {
				code := ee.GetCode()
				if strings.Contains(code, "PolicyIdNotFound") || strings.Contains(code, "11026") {
					return nil
				}
			}
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM policy failed, reason:%s\n", logId, err.Error())
		return err
	}
	return nil
}
