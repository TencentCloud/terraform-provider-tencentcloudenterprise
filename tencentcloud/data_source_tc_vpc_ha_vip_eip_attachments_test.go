package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudHaVipEipAttachmentsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHaVipEipAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHaVipEipAttachmentsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHaVipEipAttachmentExists("tencentcloudenterprise_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment"),
					resource.TestCheckResourceAttr("data.cloud_vpc_ha_vip_eip_attachments.ha_vip_eip_attachments", "ha_vip_eip_attachment_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_ha_vip_eip_attachments.ha_vip_eip_attachments", "ha_vip_eip_attachment_list.0.havip_id"),
				),
			},
		},
	})
}

const testAccHaVipEipAttachmentsDataSource_basic = defaultVpcVariable + `
#Create EIP
resource "tencentcloudenterprise_eip" "eip" {
  name = "havip_eip"
}
resource "tencentcloudenterprise_vpc_ha_vip" "havip" {
  name       = "terraform_test"
  vpc_id     = var.vpc_id
  subnet_id  = var.subnet_id
}
resource "tencentcloudenterprise_vpc_ha_vip_eip_attachment" "ha_vip_eip_attachment" {
  havip_id   = cloud_vpc_ha_vip.havip.id
  address_ip = cloud_eip.eip.public_ip
}

data "tencentcloudenterprise_vpc_ha_vip_eip_attachments" "ha_vip_eip_attachments" {
  havip_id = cloud_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment.havip_id
}
`
