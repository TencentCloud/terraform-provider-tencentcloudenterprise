/*
Use this data source to query detailed information of cvm instance_vnc_url

# Example Usage

```hcl

	data "tencentcloudenterprise_cvm_instance_vnc_url" "instance_vnc_url" {
	  instance_id = "ins-xxxxxxxx"
	}

```
*/
package tencentcloud

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	// 注册资源中文描述提供者映射
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_instance_vnc_url", CNDescription{
		TerraformTypeCN: "CVM 实例 VNC 连接 URL",
		DescriptionCN:   "提供CVM实例VNC连接URL数据源，用于查询CVM实例VNC连接URL的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例 ID要获取实例 ID，您可以调用 `DescribeInstances` 接口查询响应中的 `InstanceId` 字段",
			"instance_vnc_url":   "实例 VNC 连接 URL",
			"result_output_file": "用于保存结果的文件",
		},
	})
}

func dataSourceTencentCloudCvmInstanceVncUrl() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of cvm instance_vnc_url",
		Read:        dataSourceTencentCloudCvmInstanceVncUrlRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance ID. To obtain the instance IDs, you can call `DescribeInstances` and look for `InstanceId` in the response.",
			},

			"instance_vnc_url": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Instance VNC URL.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCvmInstanceVncUrlRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_instance_vnc_url.read")()
	defer inconsistentCheck(d, meta)()

	var response *cvm.DescribeInstanceVncUrlResponse
	request := cvm.NewDescribeInstanceVncUrlRequest()
	instanceId := d.Get("instance_id").(string)
	request.InstanceId = helper.String(instanceId)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {

		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DescribeInstanceVncUrl(request)
		if e != nil {
			return retryError(e)
		}
		response = result
		return nil
	})
	if err != nil {
		return err
	}

	if response == nil || response.Response == nil {
		d.SetId("")
		return fmt.Errorf("Response is nil")

	}
	d.SetId(instanceId)
	_ = d.Set("instance_vnc_url", *response.Response.InstanceVncUrl)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), map[string]interface{}{
			"instance_vnc_url": *response.Response.InstanceVncUrl,
		}); e != nil {
			return e
		}
	}
	return nil
}
