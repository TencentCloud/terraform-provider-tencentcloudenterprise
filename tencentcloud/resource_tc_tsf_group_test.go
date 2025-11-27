package tencentcloud

import (
	"context"
	"fmt"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfGroupResource_basic -v
func TestAccTencentCloudTsfGroupResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfGroup2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfGroupExists("tencentcloudenterprise_tsf_group.group"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_group.group", "id"),
					//resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "application_id", defaultTsfApplicationId),
					//resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "namespace_id", defaultNamespaceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "group_name", "terraform-test"),
					//resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "cluster_id", defaultTsfClustId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "group_desc", "terraform desc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "alias", "terraform test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_group.group", "tags.createdBy", "terraform"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tsf_group.group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTsfGroupDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tsf_group" {
			continue
		}

		res, err := service.DescribeTsfGroupById(ctx, rs.Primary.ID)
		if err != nil {
			code := err.(*sdkErrors.CloudSDKError).Code
			if code == "ResourceNotFound.GroupNotExist" {
				return nil
			}
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf group %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfGroupExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfGroupById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf group %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfGroupVar = `
variable "application_id" {
	default = "` + defaultTsfApplicationId + `"
}
variable "namespace_id" {
	default = "` + defaultNamespaceId + `"
}
variable "cluster_id" {
	default = "` + defaultTsfClustId + `"
}
`

const testAccTsfGroup = testAccTsfGroupVar + `

resource "tencentcloudenterprise_tsf_group" "group" {
	application_id = var.application_id
	namespace_id = var.namespace_id
	group_name = "terraform-test"
	cluster_id = var.cluster_id
	group_desc = "terraform desc"
	alias = "terraform test"
	tags = {
	  "createdBy" = "terraform"
	}
  }

`

const testAccTsfGroup2 = `
resource "tencentcloudenterprise_tsf_group" "group" {
	application_id = "application-dapbqkyo"
	namespace_id = "namespace-nqab5eal"
	group_name = "terraform-test"
	cluster_id = "cluster-m2vzm9yp"
	group_desc = "terraform desc"
	// alias = "terraform test"
	tags = {
	  "createdBy" = "terraform"
	}
}
`
