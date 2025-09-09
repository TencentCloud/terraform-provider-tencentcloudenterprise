package tencentcloud

/*
import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const (
	targetGroupById     = "data.cloud_clb_target_groups.target_group_info_id"
	targetGroupByName   = "data.cloud_clb_target_groups.target_group_info_name"
	targetGroupResource = "cloud_clb_target_group.test"
)

func TestAccTencentCloudClbTargetGroupDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbTargetGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceClbTargetGroup,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbTargetGroupExists(targetGroupResource),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.#"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.target_group_id"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.vpc_id"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.target_group_name"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.port"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.create_time"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.update_time"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.associated_rule_list.#"),
					resource.TestCheckResourceAttrSet(targetGroupById, "list.0.target_group_instance_list.#"),
				),
			},
			{
				Config: testAccTencentCloudDataSourceClbTargetGroupName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbTargetGroupExists(targetGroupResource),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.#"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.target_group_id"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.vpc_id"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.target_group_name"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.port"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.create_time"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.update_time"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.associated_rule_list.#"),
					resource.TestCheckResourceAttrSet(targetGroupByName, "list.0.target_group_instance_list.#"),
				),
			},
		},
	})
}

const tareGetGroupBase = defaultVpcSubnets + `
resource "cloud_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "tf-clb-rule-data"
  vpc_id = local.vpc_id
}

resource "cloud_clb_listener" "listener_basic" {
  clb_id        = cloud_clb_instance.clb_basic.id
  port          = 1
  protocol      = "HTTP"
  listener_name = "listener_basic"
}

resource "cloud_clb_listener_rule" "rule_basic" {
  clb_id              = cloud_clb_instance.clb_basic.id
  listener_id         = cloud_clb_listener.listener_basic.listener_id
  domain              = "abc.com"
  url                 = "/"
  session_expire_time = 30
  scheduler           = "WRR"
  target_type         = "TARGETGROUP"
}

resource "cloud_clb_target_group" "test"{
    target_group_name = "test-target-keep-1"
    vpc_id = local.vpc_id
}

resource "cloud_clb_target_group_attachment" "group" {
    clb_id          = cloud_clb_instance.clb_basic.id
    listener_id     = cloud_clb_listener.listener_basic.listener_id
    rule_id         = cloud_clb_listener_rule.rule_basic.rule_id
    target_group_id = cloud_clb_target_group.test.id
}
`

const testAccTencentCloudDataSourceClbTargetGroup = tareGetGroupBase + `
data "cloud_clb_target_groups" "target_group_info_id" {
  target_group_id = cloud_clb_target_group.test.id
}
`

const testAccTencentCloudDataSourceClbTargetGroupName = tareGetGroupBase + `
data "cloud_clb_target_groups" "target_group_info_name" {
  target_group_name = cloud_clb_target_group.test.target_group_name
}
`
*/
