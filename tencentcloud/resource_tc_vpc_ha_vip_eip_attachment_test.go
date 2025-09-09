package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudHaVipEipAttachment_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHaVipEipAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHaVipEipAttachment_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaVipEipAttachmentExists("cloud_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment_basic"),
					resource.TestCheckResourceAttrSet("cloud_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment_basic", "havip_id"),
					resource.TestCheckResourceAttrSet("cloud_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment_basic", "address_ip"),
				),
			},
			{
				ResourceName:      "cloud_vpc_ha_vip_eip_attachment.ha_vip_eip_attachment_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckHaVipEipAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcService := VpcService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_vpc_ha_vip_eip_attachment" {
			continue
		}

		_, _, has, err := vpcService.DescribeHaVipEipById(ctx, rs.Primary.ID)
		if err == nil && has {
			return fmt.Errorf("HA VIP EIP attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckHaVipEipAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("HA VIP EIP attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("HA VIP EIP attachment id is not set")
		}
		vpcService := VpcService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		_, _, has, err := vpcService.DescribeHaVipEipById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("HA VIP EIP attachment does not exist: %s", rs.Primary.ID)
		}
		return nil
	}
}

const testAccHaVipEipAttachment_basic = defaultVpcVariable + `
#Create EIP
resource "cloud_eip" "eip" {
  name = "havip_eip"
}
resource "cloud_vpc_ha_vip" "havip" {
  name      = "terraform_test"
  vpc_id    = var.vpc_id
  subnet_id = var.subnet_id
}
resource "cloud_vpc_ha_vip_eip_attachment" "ha_vip_eip_attachment_basic"{
  havip_id = cloud_vpc_ha_vip.havip.id
  address_ip = cloud_eip.eip.public_ip
}
`
