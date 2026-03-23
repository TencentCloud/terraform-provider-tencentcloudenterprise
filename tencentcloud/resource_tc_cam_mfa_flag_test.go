package tencentcloud

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// 测试基础场景 - 设置和更新 MFA 保护状态
func TestAccTencentCloudCamMfaFlag_basic(t *testing.T) {
	t.Parallel()

	testAccPreCheck(t)
	testUin := getTestCamMfaUin(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Step 1: 关闭所有 MFA 保护
			{
				Config: testAccCamMfaFlagDisabled(testUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamMfaFlagExists("tencentcloudenterprise_cam_mfa_flag.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "op_uin", testUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.phone", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.stoken", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.ukey", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.phone", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.stoken", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.ukey", "0"),
				),
			},
			// Step 2: 开启 MFA 登录保护和操作保护
			{
				Config: testAccCamMfaFlagEnabled(testUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamMfaFlagExists("tencentcloudenterprise_cam_mfa_flag.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "op_uin", testUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.phone", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.stoken", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.ukey", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.phone", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.stoken", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.ukey", "0"),
				),
			},
			// Step 3: 再次关闭所有 MFA 保护
			{
				Config: testAccCamMfaFlagDisabled(testUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamMfaFlagExists("tencentcloudenterprise_cam_mfa_flag.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "op_uin", testUin),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.phone", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.stoken", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "login_flag.0.ukey", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.phone", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.stoken", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.token", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_mfa_flag.test", "action_flag.0.ukey", "0"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_mfa_flag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func getTestCamMfaUin(t *testing.T) string {
	t.Helper()

	testUin, err := resolveTestCamMfaUin(os.Getenv("TENCENTCLOUD_TEST_CAM_UIN"), lookupCurrentLoginUinForCamMfaTest)
	if err != nil {
		t.Fatalf("resolve CAM MFA test uin failed: %v", err)
	}
	return testUin
}

func resolveTestCamMfaUin(explicitUin string, lookup func() (string, error)) (string, error) {
	if explicitUin != "" {
		return explicitUin, nil
	}

	return lookup()
}

func lookupCurrentLoginUinForCamMfaTest() (string, error) {
	region := os.Getenv(PROVIDER_REGION)
	if region == "" {
		region = defaultRegion
	}

	sharedClient, err := sharedClientForRegion(region)
	if err != nil {
		return "", err
	}

	tcClient := sharedClient.(*TencentCloudClient)
	response, err := tcClient.apiV3Conn.UseAccountClient().GetMaskedUserInfoByLoginUin(account.NewGetMaskedUserInfoByLoginUinRequest())
	if err != nil {
		return "", err
	}
	if response == nil || response.Response == nil || response.Response.Uin == nil {
		return "", fmt.Errorf("GetMaskedUserInfoByLoginUin returned empty uin")
	}

	return strconv.FormatUint(*response.Response.Uin, 10), nil
}

func testAccCheckCamMfaFlagExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM mfa_flag][Exists] resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM mfa_flag][Exists] ID is not set")
		}

		camService := &CamService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		uin, err := strconv.ParseUint(rs.Primary.ID, 10, 64)
		if err != nil {
			return fmt.Errorf("[CHECK][CAM mfa_flag][Exists] parse uin failed: %v", err)
		}

		loginFlag, actionFlag, err := camService.DescribeCamMfaFlagById(context.Background(), uin)
		if err != nil {
			return err
		}

		// API 调用成功即可，配置可能为空
		_ = loginFlag
		_ = actionFlag

		return nil
	}
}

// 关闭所有 MFA 保护的配置
func testAccCamMfaFlagDisabled(uin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_mfa_flag" "test" {
  op_uin = %s
  
  login_flag {
    phone  = 0
    stoken = 0
    token  = 0
    ukey   = 0
  }
  
  action_flag {
    phone  = 0
    stoken = 0
    token  = 0
    ukey   = 0
  }
}
`, uin)
}

// 开启部分 MFA 保护的配置
func testAccCamMfaFlagEnabled(uin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_mfa_flag" "test" {
  op_uin = %s
  
  login_flag {
    phone  = 1
    stoken = 1
    token  = 0
    ukey   = 0
  }
  
  action_flag {
    phone  = 1
    stoken = 0
    token  = 0
    ukey   = 0
  }
}
`, uin)
}
