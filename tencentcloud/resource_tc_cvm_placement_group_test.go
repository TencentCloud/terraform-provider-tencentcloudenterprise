package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudPlacementGroup(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPlacementGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPlacementGroupExists("tencentcloudenterprise_bms_placement_group.placement"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_bms_placement_group.placement", "name", "tf-test-placement"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_bms_placement_group.placement", "type", "HOST"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_bms_placement_group.placement", "cvm_quota_total"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_bms_placement_group.placement", "current_num"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_bms_placement_group.placement", "create_time"),
				),
			},
		},
	})
}

func testAccCheckPlacementGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("placement group %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("placement group id is not set")
		}

		cvmService := CvmService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		placement, err := cvmService.DescribePlacementGroupById(ctx, rs.Primary.ID)
		if err != nil {
			err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
				placement, err = cvmService.DescribePlacementGroupById(ctx, rs.Primary.ID)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
		}
		if err != nil {
			return err
		}
		if placement == nil {
			return fmt.Errorf("placement group id is not found")
		}
		return nil
	}
}

func testAccCheckPlacementGroupDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cvmService := CvmService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_bms_placement_group" {
			continue
		}

		placement, err := cvmService.DescribePlacementGroupById(ctx, rs.Primary.ID)
		if err != nil {
			err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
				placement, err = cvmService.DescribePlacementGroupById(ctx, rs.Primary.ID)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
		}
		if err != nil {
			return err
		}
		if placement != nil {
			return fmt.Errorf("placement group still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

const testAccPlacementGroup = `
resource "tencentcloudenterprise_bms_placement_group" "placement" {
	name = "tf-test-placement"
	type = "HOST"
}
`
