/*
Use this data source to query detailed information of cvm disaster_recover_group_quota

# Example Usage

```hcl
data "tencentcloudenterprise_cvm_disaster_recover_group_quota" "disaster_recover_group_quota" {
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
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	// 注册资源中文描述提供者映射
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_disaster_recover_group_quota", CNDescription{
		TerraformTypeCN: "CVM 置放群组配额",
		DescriptionCN:   "提供CVM 置放群组配额数据源，用于查询CVM 置放群组配额信息。",
		AttributesCN: map[string]string{
			"group_quota":             "可创建的置放群组数量上限",
			"current_num":             "当前用户已创建的置放群组数量",
			"cvm_in_host_group_quota": "物理机类型容灾组内实例的配额数",
			"cvm_in_sw_group_quota":   "交换机类型容灾组内实例的配额数",
			"cvm_in_rack_group_quota": "机架类型容灾组内实例的配额数",
			"result_output_file":      "用于保存结果的文件",
		},
	})
}

func dataSourceTencentCloudCvmDisasterRecoverGroupQuota() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of cvm disaster_recover_group_quota",
		Read:        dataSourceTencentCloudCvmDisasterRecoverGroupQuotaRead,
		Schema: map[string]*schema.Schema{
			"group_quota": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "The maximum number of placement groups that can be created.",
			},

			"current_num": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "The number of placement groups that have been created by the current user.",
			},

			"cvm_in_host_group_quota": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Quota on instances in a physical-machine-type disaster recovery group.",
			},

			"cvm_in_sw_group_quota": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Quota on instances in a switch-type disaster recovery group.",
			},

			"cvm_in_rack_group_quota": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Quota on instances in a rack-type disaster recovery group.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCvmDisasterRecoverGroupQuotaRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_disaster_recover_group_quota.read")()
	defer inconsistentCheck(d, meta)()

	var response *cvm.DescribeDisasterRecoverGroupQuotaResponse

	request := cvm.NewDescribeDisasterRecoverGroupQuotaRequest()
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DescribeDisasterRecoverGroupQuota(request)
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
	_ = d.Set("group_quota", response.Response.GroupQuota)
	_ = d.Set("current_num", response.Response.CurrentNum)
	_ = d.Set("cvm_in_host_group_quota", response.Response.CvmInHostGroupQuota)
	_ = d.Set("cvm_in_sw_group_quota", response.Response.CvmInSwGroupQuota)
	_ = d.Set("cvm_in_rack_group_quota", response.Response.CvmInRackGroupQuota)

	d.SetId(helper.BuildToken())
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), map[string]interface{}{
			"group_quota":             response.Response.GroupQuota,
			"current_num":             response.Response.CurrentNum,
			"cvm_in_host_group_quota": response.Response.CvmInHostGroupQuota,
			"cvm_in_sw_group_quota":   response.Response.CvmInSwGroupQuota,
			"cvm_in_rack_group_quota": response.Response.CvmInRackGroupQuota,
		}); e != nil {
			return e
		}
	}
	return nil
}
