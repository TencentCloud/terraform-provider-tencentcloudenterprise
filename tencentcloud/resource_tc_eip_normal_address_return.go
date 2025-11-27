/*
Provides a resource to create a vpc normal_address_return

# Example Usage

```hcl

	resource "tencentcloudenterprise_eip_normal_address_return" "normal_address_return" {
	  address_ids = ["eip-8zei45vm"]
	}

```
*/
package tencentcloud

import (
	"log"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_eip_normal_address_return", CNDescription{
		TerraformTypeCN: "退还普通公网 IP",
		DescriptionCN:   "提供退还普通公网IP资源，用于退还普通公网IP。",
		AttributesCN: map[string]string{
			"address_ids": "eip实例id",
		},
	})
}

func resourceTencentCloudEipNormalAddressReturn() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a vpc normal_address_return",
		Create:      resourceTencentCloudEipNormalAddressReturnCreate,
		Read:        resourceTencentCloudEipNormalAddressReturnRead,
		Delete:      resourceTencentCloudEipNormalAddressReturnDelete,
		Schema: map[string]*schema.Schema{
			"address_ids": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The Id of the EIP, example: eip-8zei45vm.",
			},
		},
	}
}

func resourceTencentCloudEipNormalAddressReturnCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_normal_address_return.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = cvm.NewReturnAddressesRequest()
		addressIds string
	)
	if v, ok := d.GetOk("address_ids"); ok {
		addressIpsSet := v.(*schema.Set).List()
		for i := range addressIpsSet {
			addressId := addressIpsSet[i].(string)
			request.AddressIds = append(request.AddressIds, &addressId)
			addressIds = addressId + FILED_SP
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().ReturnAddresses(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate vpc normalAddressReturn failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(addressIds)

	return resourceTencentCloudEipNormalAddressReturnRead(d, meta)
}

func resourceTencentCloudEipNormalAddressReturnRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_normal_address_return.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudEipNormalAddressReturnDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_eip_normal_address_return.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
