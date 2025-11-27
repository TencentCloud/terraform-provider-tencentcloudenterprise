package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudEipAssociationWithInstance(t *testing.T) {
	t.Parallel()
	id := "tencentcloudenterprise_eip_association.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEipAssociationDestroy,
		Steps: []resource.TestStep{
			{
				PreConfig: func() { testAccStepPreConfigSetTempAKSK(t, ACCOUNT_TYPE_COMMON) },
				Config:    testAccTencentCloudEipAssociationWithInstance,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEipAssociationExists(id),
					resource.TestCheckResourceAttrSet(id, "eip_id"),
					resource.TestCheckResourceAttrSet(id, "instance_id"),
					resource.TestCheckNoResourceAttr(id, "network_interface_id"),
					resource.TestCheckNoResourceAttr(id, "private_ip"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_eip.foo", "public_ip"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_eip.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_eip.foo", "status", "UNBIND"),
				),
			},
			{
				ResourceName:      id,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudEipAssociationWithNetworkInterface(t *testing.T) {
	t.Parallel()
	id := "tencentcloudenterprise_eip_association.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEipAssociationDestroy,
		Steps: []resource.TestStep{
			{
				PreConfig: func() { testAccStepPreConfigSetTempAKSK(t, ACCOUNT_TYPE_COMMON) },
				Config:    testAccTencentCloudEipAssociationWithNetworkInterface,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEipAssociationExists(id),
					resource.TestCheckResourceAttrSet(id, "eip_id"),
					resource.TestCheckResourceAttrSet(id, "network_interface_id"),
					resource.TestCheckResourceAttrSet(id, "private_ip"),
					resource.TestCheckNoResourceAttr(id, "instance_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_eip.foo", "public_ip"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_eip.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_eip.foo", "status", "UNBIND"),
				),
			},
		},
	})
}

func testAccCheckEipAssociationDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	vpcService := VpcService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_eip_association" {
			continue
		}

		associationId, err := parseEipAssociationId(rs.Primary.ID)

		if err != nil {
			return err
		}

		eip, err := vpcService.DescribeEipById(ctx, associationId.EipId)
		if err != nil {
			err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
				eip, err = vpcService.DescribeEipById(ctx, associationId.EipId)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		if eip == nil {
			return nil
		}
		if eip.InstanceId != nil && *eip.InstanceId == associationId.InstanceId {
			return fmt.Errorf("eip association still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckEipAssociationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("eip association %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("eip association id is not set")
		}
		vpcService := VpcService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		associationId, err := parseEipAssociationId(rs.Primary.ID)

		if err != nil {
			return err
		}
		eip, err := vpcService.DescribeEipById(ctx, associationId.EipId)
		if err != nil {
			err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
				eip, err = vpcService.DescribeEipById(ctx, associationId.EipId)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		if eip == nil || (eip.InstanceId != nil && *eip.InstanceId != associationId.InstanceId) {
			return fmt.Errorf("eip %s is not found", associationId.EipId)
		}
		return nil
	}
}

const testAccTencentCloudEipAssociationWithInstance = defaultInstanceVariable + `
resource "tencentcloudenterprise_eip" "foo" {
  name = var.instance_name
}

resource "tencentcloudenterprise_cvm_instance" "foo" {
  instance_name      = var.instance_name
  availability_zone  = var.availability_cvm_zone
  image_id           = data.cloud_cvm_images.default.images.0.image_id
  instance_type      = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type   = "CLOUD_PREMIUM"
}

resource "tencentcloudenterprise_eip_association" "foo" {
  eip_id      = cloud_eip.foo.id
  instance_id = cloud_cvm_instance.foo.id
}
`

const testAccTencentCloudEipAssociationWithNetworkInterface = defaultVpcVariable + `
resource "tencentcloudenterprise_eip" "foo" {
  name = var.instance_name
}

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = var.instance_name
  vpc_id      = var.vpc_id
  subnet_id   = var.subnet_id
  description = var.instance_name
  ipv4_count  = 1
}

resource "tencentcloudenterprise_eip_association" "foo" {
  eip_id               = cloud_eip.foo.id
  network_interface_id = cloud_vpc_eni.foo.id
  private_ip           = cloud_vpc_eni.foo.ipv4_info.0.ip
}
`
