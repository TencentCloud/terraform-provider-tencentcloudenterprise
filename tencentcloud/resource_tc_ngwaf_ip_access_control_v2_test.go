package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudNgwafIpAccessControlV2_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNgwafIpAccessControlV2Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNgwafIpAccessControlV2Basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNgwafIpAccessControlV2Exists("tencentcloudenterprise_ngwaf_ip_access_control_v2.example"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "instance_id", defaultNgwafIpAccessControlV2InstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "domain", defaultNgwafIpAccessControlV2Domain),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "action_type", "40"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "note", "note."),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "ip_list.#", "3"),
				),
			},
			{
				Config: testAccNgwafIpAccessControlV2Update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNgwafIpAccessControlV2Exists("tencentcloudenterprise_ngwaf_ip_access_control_v2.example"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "instance_id", defaultNgwafIpAccessControlV2InstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "domain", defaultNgwafIpAccessControlV2Domain),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "action_type", "40"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "note", "note update."),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_ip_access_control_v2.example", "ip_list.#", "4"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_ngwaf_ip_access_control_v2.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckNgwafIpAccessControlV2Exists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("resource %s id is not set", resourceName)
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 3 {
			return fmt.Errorf("resource %s id is broken: %s", resourceName, rs.Primary.ID)
		}

		service := &NgwafService{client: testAccProvider.Meta().(*TencentCloudClient)}
		ctx := context.WithValue(context.Background(), logIdKey, getLogId(contextNil))

		respData, err := service.DescribeWafIpAccessControlV2ById(ctx, idSplit[1], idSplit[2])
		if err != nil {
			return err
		}

		if findNgwafIpAccessControlV2Rule(respData, idSplit[2]) == nil {
			return fmt.Errorf("ngwaf ip access control v2 %s not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckNgwafIpAccessControlV2Destroy(s *terraform.State) error {
	service := &NgwafService{client: testAccProvider.Meta().(*TencentCloudClient)}
	ctx := context.WithValue(context.Background(), logIdKey, getLogId(contextNil))

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_ngwaf_ip_access_control_v2" {
			continue
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 3 {
			return fmt.Errorf("resource %s id is broken: %s", rs.Type, rs.Primary.ID)
		}

		respData, err := service.DescribeWafIpAccessControlV2ById(ctx, idSplit[1], idSplit[2])
		if err != nil {
			return err
		}

		if findNgwafIpAccessControlV2Rule(respData, idSplit[2]) != nil {
			return fmt.Errorf("ngwaf ip access control v2 %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccNgwafIpAccessControlV2Basic = defaultNgwafVariable + `
resource "tencentcloudenterprise_ngwaf_ip_access_control_v2" "example" {
  instance_id = var.ngwaf_ip_access_control_v2_instance_id
  domain      = var.ngwaf_ip_access_control_v2_domain
  action_type = 40
  note        = "note."

  ip_list = [
    "10.0.0.10",
    "172.0.0.16",
    "192.168.0.30",
  ]

  job_type = "TimedJob"

  job_date_time {
    time_t_zone = "UTC+8"

    timed {
      end_date_time   = 0
      start_date_time = 0
    }
  }
}
`

const testAccNgwafIpAccessControlV2Update = defaultNgwafVariable + `
resource "tencentcloudenterprise_ngwaf_ip_access_control_v2" "example" {
  instance_id = var.ngwaf_ip_access_control_v2_instance_id
  domain      = var.ngwaf_ip_access_control_v2_domain
  action_type = 40
  note        = "note update."

  ip_list = [
    "10.0.0.10",
    "172.0.0.16",
    "192.168.0.30",
    "168.10.10.10",
  ]

  job_type = "TimedJob"

  job_date_time {
    time_t_zone = "UTC+8"

    timed {
      end_date_time   = 0
      start_date_time = 0
    }
  }
}
`
