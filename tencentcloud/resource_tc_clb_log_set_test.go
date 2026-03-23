package tencentcloud

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudClbLogSet_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbLogSet_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbLogSetExists("tencentcloudenterprise_clb_log_set.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_log_set.foo", "create_time"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_set.foo", "name", "clb_logset"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_set.foo", "period", "7"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_log_set.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckClbLogSetExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB log set][Exists] check: CLB log set %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB log set][Exists] check: CLB log set id is not set")
		}
		service := ClsService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		resourceId := rs.Primary.ID
		instance, err := service.DescribeClsLogset(ctx, resourceId)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CLB log set][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckClbLogSetDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clsService := ClsService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_clb_log_set" {
			continue
		}
		time.Sleep(5 * time.Second)
		resourceId := rs.Primary.ID
		info, err := clsService.DescribeClsLogset(ctx, resourceId)
		if info != nil && err == nil {
			return fmt.Errorf("[CHECK][CLB log set][Destroy] check: CLB log set still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

const testAccClbLogSet_basic = `
resource "tencentcloudenterprise_clb_log_set" "foo" {
  period = 7
}
`
