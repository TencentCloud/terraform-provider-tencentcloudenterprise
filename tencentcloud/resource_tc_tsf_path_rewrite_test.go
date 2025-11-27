package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfPathRewriteResource_basic -v
func TestAccTencentCloudTsfPathRewriteResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfPathRewriteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfPathRewriteTce,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfPathRewritekExists("tencentcloudenterprise_tsf_path_rewrite.path_rewrite"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "id"),
					//resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "gateway_group_id", defaultTsfGroupId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "regex", "/test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "replacement", "/tt"),
					//resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "blocked", "N"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_path_rewrite.path_rewrite", "order", "2"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tsf_path_rewrite.path_rewrite",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTsfPathRewriteDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tsf_path_rewrite" {
			continue
		}

		res, err := service.DescribeTsfPathRewriteById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf PathRewrite %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfPathRewritekExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfPathRewriteById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf PathRewrite %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfPathRewriteVar = `
variable "group_id" {
	default = "` + defaultTsfGroupId + `"
}
`

const testAccTsfPathRewrite = testAccTsfPathRewriteVar + `

resource "tencentcloudenterprise_tsf_path_rewrite" "path_rewrite" {
	gateway_group_id = var.group_id
	regex = "/test"
	replacement = "/tt"
	blocked = "N"
	order = 2
}

`

const testAccTsfPathRewriteTce = `
resource "tencentcloudenterprise_tsf_path_rewrite" "path_rewrite" {
	gateway_group_id = "group-6a7eg6y5"
	regex = "/test"
	replacement = "/tt"
	blocked = "Y"
	order = 2
}
`
