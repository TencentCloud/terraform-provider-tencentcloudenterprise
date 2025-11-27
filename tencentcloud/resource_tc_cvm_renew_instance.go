/*
Provides a resource to create a cvm renew_instance

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_renew_instance" "renew_instance" {
	  instance_id = "xxx"
	  instance_charge_prepaid {
		period = 1
		renew_flag = "NOTIFY_AND_AUTO_RENEW"
	  }
	  renew_portable_data_disk = true
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_renew_instance", CNDescription{
		TerraformTypeCN: "续费实例",
		DescriptionCN:   "提供续费实例资源，用于续费云服务器实例。",
		AttributesCN: map[string]string{
			"instance_id": "实例ID",
			"instance_charge_prepaid": "预付费模式，即包年包月相关参数设置" +
				"通过该参数可以指定包年包月实例的购买时长以及是否设置自动续费如果指定实例的付费模式为预付费，则该参数必传",
			"renew_portable_data_disk": "是否续费弹性数据盘",
			"period":                   "认购期；单位月；有效值1、2、3、4、5、6、7、8、9、10、11、12、24、36、48、60。注意：此字段可能返回null，表示找不到有效值",
			"renew_flag": "自动续订标志。有效值\n" +
				"- `NOTIFY_AND_AUTO_RENEW`:到期时通知并自动续订；\n" +
				"- `NOTIFY_AND_MANUAL_RENEW `:到期时通知，但不自动续订；\n" +
				"- `DISABLE_NOTIFY_AND_MANUAL_RENEW`:到期时不通知，也不自动续订；\n" +
				"默认值:通知_和_手动_续订。如果此参数被指定为NOTIFY_AND_AUTO_RENEW，则如果帐户余额充足，例程将每月自动续订一次。注意:该字段可能返回null，表示没有找到有效值。",
		},
	})
}
func resourceTencentCloudCvmRenewInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cvm renew_instance",
		Create:      resourceTencentCloudCvmRenewInstanceCreate,
		Read:        resourceTencentCloudCvmRenewInstanceRead,
		Delete:      resourceTencentCloudCvmRenewInstanceDelete,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance ID.",
			},

			"instance_charge_prepaid": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Prepaid mode, that is, yearly and monthly subscription related parameter settings. Through this parameter, you can specify the renewal duration of the Subscription instance, whether to set automatic renewal, and other attributes. For yearly and monthly subscription instances, this parameter is required.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"period": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Subscription period; unit: month; valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60. Note: This field may return null, indicating that no valid value is found.",
						},
						"renew_flag": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Auto renewal flag. Valid values:\n" +
								"- `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically;\n" +
								"- `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically;\n" +
								"- `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically;\n" +
								"Default value: NOTIFY_AND_MANUAL_RENEW. If this parameter is specified as NOTIFY_AND_AUTO_RENEW, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. Note: This field may return null, indicating that no valid value is found.",
						},
					},
				},
			},

			"renew_portable_data_disk": {
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeBool,
				Description: "Whether to renew the elastic data disk. Valid values:\n" +
					"- `TRUE`: Indicates to renew the subscription instance and renew the attached elastic data disk at the same time\n" +
					"- `FALSE`: Indicates that the subscription instance will be renewed and the elastic data disk attached to it will not be renewed\n" +
					"Default value: TRUE.",
			},
		},
	}
}

func resourceTencentCloudCvmRenewInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_renew_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cvm.NewRenewInstancesRequest()
	instanceId := d.Get("instance_id").(string)
	request.InstanceIds = []*string{&instanceId}

	/*
		if dMap, ok := helper.InterfacesHeadMap(d, "instance_charge_prepaid"); ok {
			instanceChargePrepaid := cvm.InstanceChargePrepaid{}
			if v, ok := dMap["period"]; ok {
				instanceChargePrepaid.Period = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["renew_flag"]; ok && v.(string) != "" {
				instanceChargePrepaid.RenewFlag = helper.String(v.(string))
			}
			request.InstanceChargePrepaid = &instanceChargePrepaid
		}

		if v, _ := d.GetOk("renew_portable_data_disk"); v != nil {
			request.RenewPortableDataDisk = helper.Bool(v.(bool))
		}

	*/

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().RenewInstances(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate cvm renewInstance failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(instanceId)

	return resourceTencentCloudCvmRenewInstanceRead(d, meta)
}

func resourceTencentCloudCvmRenewInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_renew_instance.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCvmRenewInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_renew_instance.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
