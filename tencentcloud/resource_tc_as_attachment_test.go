package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"terraform-provider-tencentcloudenterprise/sdk/common/errors"
)

func init() {
	resource.AddTestSweepers("tencentcloudenterprise_as_attachment", &resource.Sweeper{
		Name: "tencentcloudenterprise_as_attachment",
		F:    testSweepAsAttachment,
	})
}

// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=cloud_as_attachment
func testSweepAsAttachment(r string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cli, _ := sharedClientForRegion(r)
	asService := AsService{client: cli.(*TencentCloudClient).apiV3Conn}

	scalingGroups, err := asService.DescribeAutoScalingGroupByFilter(ctx, "", "", "", nil)
	if err != nil {
		return fmt.Errorf("list scaling group error: %s", err.Error())
	}

	for _, v := range scalingGroups {
		scalingGroupId := *v.AutoScalingGroupId
		scalingGroupName := *v.AutoScalingGroupName
		if !strings.HasPrefix(scalingGroupName, "tf-as-") {
			continue
		}

		var instanceIds []string
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, errRet := asService.DescribeAutoScalingAttachment(ctx, scalingGroupId, true)
			if errRet != nil {
				return retryError(errRet)
			}
			instanceIds = result
			return nil
		})
		if err != nil {
			return err
		}
		if len(instanceIds) == 0 {
			continue
		}

		if err = asService.DetachInstances(ctx, scalingGroupId, instanceIds); err != nil {
			log.Printf("[ERROR] delete scaling group %s error: %s", scalingGroupId, err.Error())
		}
	}
	return nil
}

func TestAccTencentCloudAsAttachment(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAsAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAsAttachment(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAsAttachmentExists("tencentcloudenterprise_as_attachment.attachment"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_as_attachment.attachment", "scaling_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_attachment.attachment", "instance_ids.#", "1"),
				),
			},
			// test update case
			{
				Config: testAccAsAttachment_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAsAttachmentExists("tencentcloudenterprise_as_attachment.attachment"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_as_attachment.attachment", "scaling_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_attachment.attachment", "instance_ids.#", "2"),
				),
			},
		},
	})
}

func testAccCheckAsAttachmentExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("auto scaling attachment %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("auto scaling attachment id is not set")
		}
		asService := AsService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		instances, err := asService.DescribeAutoScalingAttachment(ctx, rs.Primary.ID, false)
		if err != nil {
			return err
		}
		if len(instances) < 1 {
			return fmt.Errorf("auto scaling attachement not exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAsAttachmentDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	asService := AsService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_as_attachment" {
			continue
		}

		instances, err := asService.DescribeAutoScalingAttachment(ctx, rs.Primary.ID, false)
		if err != nil {
			if sdkErr, ok := err.(*errors.CloudSDKError); ok {
				if sdkErr.Code == AsScalingGroupNotFound {
					return nil
				}
			}
			return err
		}
		if len(instances) > 0 {
			return fmt.Errorf("auto scaling attachment still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccAsAttachment() string {
	return defaultAsVariable + `
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "tf-as-vpc"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = cloud_vpc.vpc.id
  name              = "tf-as-subnet"
  cidr_block        = "10.2.11.0/24"
  availability_zone = var.availability_zone
}

resource "tencentcloudenterprise_as_scaling_config" "launch_configuration" {
  configuration_name = "tf-as-attachment-config"
  image_id           = "img-2lr9q49h"
  instance_types     = [data.cloud_cvm_instance_types.default.instance_types.0.instance_type]
}

resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name = "tf-as-attachment-group"
  configuration_id   = cloud_as_scaling_config.launch_configuration.id
  max_size           = 5
  min_size           = 0
  vpc_id             = cloud_vpc.vpc.id
  subnet_ids         = [cloud_vpc_subnet.subnet.id]
}

resource "tencentcloudenterprise_cvm_instance" "cvm_instance" {
  instance_name     = "tf_as_instance"
  availability_zone = var.availability_zone
  image_id          = "img-2lr9q49h"
  instance_type     = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type  = "CLOUD_SSD"
  vpc_id            = cloud_vpc.vpc.id
  subnet_id         = cloud_vpc_subnet.subnet.id
}

resource "tencentcloudenterprise_as_attachment" "attachment" {
  scaling_group_id = cloud_as_scaling_group.scaling_group.id
  instance_ids     = [cloud_cvm_instance.cvm_instance.id]
}
`
}

func testAccAsAttachment_update() string {
	return defaultAsVariable + `
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "tf-as-vpc"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = cloud_vpc.vpc.id
  name              = "tf-as-subnet"
  cidr_block        = "10.2.11.0/24"
  availability_zone = var.availability_zone
}

resource "tencentcloudenterprise_as_scaling_config" "launch_configuration" {
  configuration_name = "tf-as-attachment-config"
  image_id           = "img-2lr9q49h"
  instance_types     = [data.cloud_cvm_instance_types.default.instance_types.0.instance_type]
}

resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name = "tf-as-attachment-group"
  configuration_id   = cloud_as_scaling_config.launch_configuration.id
  max_size           = 5
  min_size           = 0
  vpc_id             = cloud_vpc.vpc.id
  subnet_ids         = [cloud_vpc_subnet.subnet.id]
}

resource "tencentcloudenterprise_cvm_instance" "cvm_instance" {
  instance_name     = "tf_as_instance"
  availability_zone = var.availability_zone
  image_id          = "img-2lr9q49h"
  instance_type     = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type  = "CLOUD_SSD"
  vpc_id            = cloud_vpc.vpc.id
  subnet_id         = cloud_vpc_subnet.subnet.id
}

resource "tencentcloudenterprise_cvm_instance" "cvm_instance_1" {
  instance_name     = "tf_as_instance_1"
  availability_zone = var.availability_zone
  image_id          = "img-2lr9q49h"
  instance_type     = data.cloud_cvm_instance_types.default.instance_types.0.instance_type
  system_disk_type  = "CLOUD_SSD"
  vpc_id            = cloud_vpc.vpc.id
  subnet_id         = cloud_vpc_subnet.subnet.id
}

resource "tencentcloudenterprise_as_attachment" "attachment" {
  scaling_group_id = cloud_as_scaling_group.scaling_group.id
  instance_ids     = [cloud_cvm_instance.cvm_instance.id, cloud_cvm_instance.cvm_instance_1.id]
}
`
}
