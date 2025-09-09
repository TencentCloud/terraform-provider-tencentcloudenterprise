package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfInstancesAttachmentResource_basic -v
func TestAccTencentCloudTsfInstancesAttachmentResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfInstancesAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfInstancesAttachmentTce,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfInstancesAttachmentExists("cloud_tsf_instances_attachment.instances_attachment"),
					resource.TestCheckResourceAttrSet("cloud_tsf_instances_attachment.instances_attachment", "id"),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "cluster_id", defaultTsfClustId),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "image_id", defaultTsfImageId),
					resource.TestCheckResourceAttrSet("cloud_tsf_instances_attachment.instances_attachment", "instance_id"),
					resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "instance_import_mode", "R"),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "os_customize_type", "my_customize"),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "instance_advanced_settings.#", "1"),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "instance_advanced_settings.0.docker_graph_path", "/var/lib/docker"),
					//resource.TestCheckResourceAttr("cloud_tsf_instances_attachment.instances_attachment", "instance_advanced_settings.0.mount_target", "/mnt/data"),
				),
			},
		},
	})
}

func testAccCheckTsfInstancesAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_tsf_instances_attachment" {
			continue
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		clusterId := idSplit[0]
		instanceId := idSplit[1]

		res, err := service.DescribeTsfInstancesAttachmentById(ctx, clusterId, instanceId)
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf instancesAttachment %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfInstancesAttachmentExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}
		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		clusterId := idSplit[0]
		instanceId := idSplit[1]

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfInstancesAttachmentById(ctx, clusterId, instanceId)
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf instancesAttachment %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfInstancesAttachmentVar = `
variable "cluster_id" {
	default = "` + defaultTsfClustId + `"
}

variable "image_id" {
	default = "` + defaultTsfImageId + `"
}

`

const testAccTsfInstancesAttachment = testAccTsfInstancesAttachmentVar + `

data "cloud_cvm_instance_types" "default" {
	filter {
	  name   = "instance-family"
	  values = ["S1"]
	}
  
	cpu_core_count = 2
	memory_size    = 2
  }
  
  resource "cloud_cvm_instance" "foo" {
	instance_name     = "tf-tsf-test"
	availability_zone = "ap-guangzhou-3"
	image_id          = var.image_id
	instance_type     = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
	system_disk_type  = "CLOUD_PREMIUM"
	instance_charge_type       = "PREPAID"
	instance_charge_type_prepaid_period = 1
	instance_charge_type_prepaid_renew_flag = "NOTIFY_AND_MANUAL_RENEW"
	force_delete = true
  }

resource "cloud_tsf_instances_attachment" "instances_attachment" {
	cluster_id = var.cluster_id
	instance_id = cloud_cvm_instance.foo.id
	# os_name = "CentOS Stream 8"
	image_id = var.image_id
	password = "MyP@ssw0rd"
	# key_id = "key-123456"
	# sg_id = "sg-123456"
	instance_import_mode = "R"
	os_customize_type = "my_customize"
	# feature_id_list =
	instance_advanced_settings {
		  mount_target = "/mnt/data"
		  docker_graph_path = "/var/lib/docker"
	}
	# security_group_ids = [""]
  }

`

const testAccTsfInstancesAttachmentTce = `
resource "cloud_tsf_instances_attachment" "instances_attachment" {
	cluster_id = "cluster-zvw7jwy8"
	instance_id = "ins-j7za7rwo"
	image_id = "img-3y126h0t"
	password = "MyP@ssw0rd"
	instance_import_mode = "R"
}
`
