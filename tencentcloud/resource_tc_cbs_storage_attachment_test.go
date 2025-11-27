package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCbsStorageAttachment(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsbStorageAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCbsStorageAttachmentConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCbsStorageAttachmentExists("tencentcloudenterprise_cbs_storage_attachment.foo"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cbs_storage_attachment.foo", "storage_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cbs_storage_attachment.foo", "instance_id"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cbs_storage_attachment.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCsbStorageAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	cbsService := CbsService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_cbs_storage_attachment" {
			continue
		}

		storage, err := cbsService.DescribeDiskById(ctx, rs.Primary.ID)
		if storage == nil {
			continue
		}
		if err != nil {
			return err
		}
		if *storage.Attached {
			return fmt.Errorf("cbs storage attchment still exists")
		}
	}

	return nil
}

func testAccCheckCbsStorageAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("cbs storage attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("cbs storage attachment id is not set")
		}
		cbsService := CbsService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		storage, err := cbsService.DescribeDiskById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if storage == nil || *storage.Attached == false {
			return fmt.Errorf("cbs storage attchment not exists")
		}
		return nil
	}
}

const testAccCbsStorageAttachmentConfig = defaultInstanceVariable + defaultAzVariable + `
resource "tencentcloudenterprise_cvm_instance" "test_cbs_attach" {
  instance_name     = "test-cbs-attach-cvm"
  availability_zone = var.default_az
  image_id          = data.cloud_cvm_images.default.images.0.image_id
  system_disk_type  = "CLOUD_PREMIUM"
  instance_type     = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
}

resource "tencentcloudenterprise_cbs_storage" "foo" {
  availability_zone = var.default_az
  storage_size      = 100
  storage_type      = "CLOUD_PREMIUM"
  storage_name      = "test-cbs-attachment"
  charge_type       = "POSTPAID_BY_HOUR"
}

resource "tencentcloudenterprise_cbs_storage_attachment" "foo" {
  storage_id  = cloud_cbs_storage.foo.id
  instance_id = cloud_cvm_instance.test_cbs_attach.id
}
`
