package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"testing"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloudenterprise_cam_group
	resource.AddTestSweepers("tencentcloudenterprise_cam_group", &resource.Sweeper{
		Name: "tencentcloudenterprise_cam_group",
		F: func(r string) error {
			logId := getLogId(contextNil)
			ctx := context.WithValue(context.TODO(), logIdKey, logId)
			cli, _ := sharedClientForRegion(r)
			client := cli.(*TencentCloudClient).apiV3Conn

			service := CamService{client: client}

			groups, err := service.DescribeGroupsByFilter(ctx, nil)
			if err != nil {
				return err
			}
			for _, v := range groups {
				name := *v.GroupName

				if persistResource.MatchString(name) {
					continue
				}

				request := cam.NewDeleteGroupRequest()
				// Convert uint64 to int64
				groupIdInt := int64(*v.GroupId)
				request.GroupId = &groupIdInt
				if _, err := client.UseCamClient().DeleteGroup(request); err != nil {
					log.Printf("[%s] error, request: %s \nreason: %s ", request.GetAction(), request.ToJsonString(), err.Error())
					continue
				}
			}

			return nil
		},
	})
}

func TestAccTencentCloudCamGroup_basic(t *testing.T) {
	t.Parallel()
	randomSuffix := acctest.RandString(6)
	name1 := fmt.Sprintf("cam-group-test-%s-1", randomSuffix)
	name2 := fmt.Sprintf("cam-group-test-%s-2", randomSuffix)
	name3 := fmt.Sprintf("cam-group-test-%s-3", randomSuffix)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCamGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCamGroup_basic(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupExists("tencentcloudenterprise_cam_group.group_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "name", name1),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "remark", "test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "channel", "3"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "create_time"),
				),
			}, {
				Config: testAccCamGroup_update_name(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupExists("tencentcloudenterprise_cam_group.group_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "name", name2),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "remark", "test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "channel", "3"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "create_time"),
				),
			},
			{
				Config: testAccCamGroup_update_remark(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupExists("tencentcloudenterprise_cam_group.group_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "name", name2),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "remark", "test2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "channel", "3"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "create_time"),
				),
			},
			{
				Config: testAccCamGroup_update_all(name3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCamGroupExists("tencentcloudenterprise_cam_group.group_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "name", name3),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "remark", "test3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_group.group_basic", "channel", "3"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_group.group_basic", "create_time"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cam_group.group_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCamGroupDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cam_group" {
			continue
		}

		instance, err := camService.DescribeGroupById(ctx, rs.Primary.ID)
		if err == nil && instance != nil {
			return fmt.Errorf("[CHECK][CAM group][Destroy] check: CAM group still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckCamGroupExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CAM group][Exists] check: CAM group %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CAM group][Exists] check: CAM group id is not set")
		}
		camService := CamService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		instance, err := camService.DescribeGroupById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CAM group][Exists] check: CAM group is not exist")
		}
		return nil
	}
}

func testAccCamGroup_basic(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group_basic" {
  name   = "%s"
  remark = "test"
}
`, name)
}

func testAccCamGroup_update_name(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group_basic" {
  name   = "%s"
  remark = "test"
}
`, name)
}

func testAccCamGroup_update_remark(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group_basic" {
  name   = "%s"
  remark = "test2"
}
`, name)
}

func testAccCamGroup_update_all(name string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_cam_group" "group_basic" {
  name   = "%s"
  remark = "test3"
}
`, name)
}
