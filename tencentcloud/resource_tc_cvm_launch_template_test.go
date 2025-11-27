package tencentcloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccTencentCloudCvmLaunchTemplateResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmLaunchTemplate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cvm_launch_template.launch_template", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cvm_launch_template.launch_template", "image_id", "img-9qrfy1xt"),
				),
			},
		},
	})
}

const testAccCvmLaunchTemplate = `
resource "tencentcloudenterprise_cvm_launch_template" "demo" {
launch_template_name = "test"
placement {
zone = "yfm18"
project_id = 0
}
image_id = "img-95xgn7er"
data_disks {
disk_size = 50
delete_with_instance = false
disk_type = "CLOUD_BASIC"
encrypt = false
snapshot_id = "asp-38z35ryr"
}
host_name = "tichost"
instance_charge_type = "POSTPAID_BY_HOUR"
instance_count = 1
instance_name = "shiliname"
instance_type = "S5l.SMALL1"
launch_template_version_description = "miaoshu"
login_settings {
key_ids = ["skey-5l8kel12"]
}
security_group_ids = ["sg-8vzmnlc4"]
system_disk {
disk_size = 50
disk_type = "CLOUD_BASIC"
}
virtual_private_cloud {
vpc_id = "vpc-03dr982r"
subnet_id = "subnet-eiy0gnni"
as_vpc_gateway = false
ipv6_address_count = 0

}
user_data = "dGVzdA=="
}
`
