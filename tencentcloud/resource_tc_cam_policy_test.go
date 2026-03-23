package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCamPolicy_basic(t *testing.T) {
	t.Parallel()

	rName := fmt.Sprintf("tf-policy-%s", acctest.RandString(6))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamPolicyBasic(rName, "desc-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamPolicyExists("tencentcloudenterprise_cam_policy.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy.test", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy.test", "description", "desc-1"),
				),
			},
			{
				Config: testAccCamPolicyUpdate(rName, "desc-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamPolicyExists("tencentcloudenterprise_cam_policy.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy.test", "name", rName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_policy.test", "description", "desc-2"),
				),
			},
			{
				ResourceName:            "tencentcloudenterprise_cam_policy.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"document"},
			},
		},
	})
}

func testAccCheckCamPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM policy][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM policy][Exists] ID is not set")
		}

		camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		instance, err := camService.DescribePolicyById(context.Background(), rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil || instance.Response == nil || instance.Response.PolicyId == nil {
			return fmt.Errorf("[CHECK][CAM policy][Exists] policy not found")
		}
		return nil
	}
}

func testAccCheckCamPolicyDestroy(s *terraform.State) error {
	camService := CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_policy" {
			continue
		}
		instance, err := camService.DescribePolicyById(context.Background(), rs.Primary.ID)
		if err != nil {
			// tolerate not found in destroy
			if strings.Contains(err.Error(), "PolicyIdNotFound") || strings.Contains(err.Error(), "not exist") {
				continue
			}
			return err
		}
		if instance != nil && instance.Response != nil && instance.Response.PolicyId != nil {
			return fmt.Errorf("[CHECK][CAM policy][Destroy] policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCamPolicyBasic(name, desc string) string {
	// Keep document identical between create and update to avoid serviceperm issues.
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_policy" "test" {
  name        = "%s"
  description = "%s"
  document    = <<EOF
{
  "version": "2.0",
  "statement": [
    {
      "effect": "allow",
      "action": [
        "account:TokenBind",
        "account:AddSubAccount",
        "account:UploadCert",
        "account:ModifyNickname",
        "account:ModifySensitiveAction",
        "account:SendVerifyCode",
        "account:ChangeMailPassword",
        "account:ChangeSubAccountPassword",
        "account:SetSafeAuthFlag",
        "account:SetMfaDevice",
        "account:SetAttributeValues",
        "account:TokenUnBind",
        "account:DescribeSubAccounts",
        "account:DescribeUserCategory",
        "account:GetUserByAttributeValue"
      ],
      "resource": ["*"],
      "condition": {
        "ip_equal": {
          "qcs:ip": ["10.34.55.0/28"]
        }
      }
    }
  ]
}
EOF
}
`, name, desc)
}

func testAccCamPolicyUpdate(name, desc string) string {
	// Use same document as create to avoid serviceperm errors; only description changes.
	return testAccCamPolicyBasic(name, desc)
}
