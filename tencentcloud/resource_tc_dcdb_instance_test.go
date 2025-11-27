package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("tencentcloudenterprise_dcdb_instance", &resource.Sweeper{
		Name: "tencentcloudenterprise_dcdb_instance",
		F:    testSweepDcdbHourdbInstance,
	})
}

// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=cloud_dcdb_instance
func testSweepDcdbHourdbInstance(r string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cli, _ := sharedClientForRegion(r)
	dcdbService := DcdbService{client: cli.(*TencentCloudClient).apiV3Conn}

	instances, err := dcdbService.DescribeDcdbInstancesByFilter(ctx, nil)
	if err != nil {
		return err
	}
	if instances == nil {
		return fmt.Errorf("dcdb hourdb instance not exists.")
	}

	for _, v := range instances {
		delId := *v.InstanceId
		delName := *v.InstanceName

		if strings.HasPrefix(delName, "test_dcdb_") {
			err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
				err := dcdbService.DeleteDcdbHourdbInstanceById(ctx, delId)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
			if err != nil {
				return fmt.Errorf("[ERROR] delete dcdb hourdb instance %s failed! reason:[%s]", delId, err.Error())
			}
		}
	}
	return nil
}

func TestAccTencentCloudDcdbHourdbInstanceResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDcdbHourdbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDcdbHourdbInstance_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcdbHourdbInstanceExists("tencentcloudenterprise_dcdb_instance.hourdb_instance"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "zones.#", "1"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_instance.hourdb_instance", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_instance.hourdb_instance", "subnet_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_instance.hourdb_instance", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "instance_name", "test_dcdb_hourdb_instance"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "shard_memory", "2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "shard_storage", "10"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "shard_node_count", "2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "shard_count", "2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "db_version_id", "8.0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "project_id", "0"),
					// resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "extranet_access", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "resource_tags.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "resource_tags.0.tag_key", "aaa"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "resource_tags.0.tag_value", "bbb"),
				),
			},
			{
				Config: testAccDcdbHourdbInstance_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcdbHourdbInstanceExists("tencentcloudenterprise_dcdb_instance.hourdb_instance"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_instance.hourdb_instance", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_instance.hourdb_instance", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "project_id", defaultProjectId),
					// resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "extranet_access", "false"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "vip", "203.0.113.110"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_dcdb_instance.hourdb_instance", "instance_name", "test_dcdb_hourdb_instance_CHANGED"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_dcdb_instance.hourdb_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDcdbHourdbInstanceDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	dcdbService := DcdbService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_dcdb_instance" {
			continue
		}

		ret, err := dcdbService.DescribeDcdbHourdbInstance(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if len(ret) > 0 {
			return fmt.Errorf("dcdb hourdb instance still exist, instanceId: %v", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckDcdbHourdbInstanceExists(re string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[re]
		if !ok {
			return fmt.Errorf("dcdb hourdb instance  %s is not found", re)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("dcdb hourdb instance id is not set")
		}

		dcdbService := DcdbService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		ret, err := dcdbService.DescribeDcdbHourdbInstance(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if len(ret) == 0 {
			return fmt.Errorf("dcdb hourdb instance not found, instanceId: %v", rs.Primary.ID)
		}

		return nil
	}
}

const testAccDcdbHourdb_vpc_config = defaultAzVariable + `
data "tencentcloudenterprise_vpc_security_groups" "internal" {
	name = "default"
  }
  
  data "tencentcloudenterprise_vpc_instances" "vpc" {
	name = "Default-VPC"
  }
  
  data "tencentcloudenterprise_vpc_subnets" "subnet" {
	vpc_id = data.cloud_vpc_instances.vpc.instance_list.0.vpc_id
  }
  
  resource "tencentcloudenterprise_vpc" "vpc" {
	cidr_block = "172.18.111.0/24"
	name       = "test-pg-network-vpc"
  }
  
  resource "tencentcloudenterprise_vpc_subnet" "subnet" {
	availability_zone = var.default_az
	cidr_block        = "172.18.111.0/24"
	name              = "test-pg-network-sub1"
	vpc_id            = cloud_vpc.vpc.id
  }
  
  locals {
	vpc_id        = data.cloud_vpc_subnets.subnet.instance_list.0.vpc_id
	subnet_id     = data.cloud_vpc_subnets.subnet.instance_list.0.subnet_id
	sg_id         = data.cloud_vpc_security_groups.internal.security_groups.0.security_group_id
	new_vpc_id    = cloud_vpc_subnet.subnet.vpc_id
	new_subnet_id = cloud_vpc_subnet.subnet.id
  }  
`

const testAccDcdbHourdbInstance_basic = testAccDcdbHourdb_vpc_config + `

resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
  instance_name = "test_dcdb_hourdb_instance"
  zones = [var.default_az]
  shard_memory = "2"
  shard_storage = "10"
  shard_node_count = "2"
  shard_count = "2"
  vpc_id = local.vpc_id
  subnet_id = local.subnet_id
  security_group_id = local.sg_id
  db_version_id = "8.0"
  project_id = 0
//   extranet_access = true
  resource_tags {
	tag_key = "aaa"
	tag_value = "bbb"
  }
}

`

const testAccDcdbHourdbInstance_update = testAccDcdbHourdb_vpc_config + defaultProjectVariable + `

resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
  instance_name = "test_dcdb_hourdb_instance_CHANGED"
  zones = [var.default_az]
  shard_memory = "2"
  shard_storage = "10"
  shard_node_count = "2"
  shard_count = "2"
  vpc_id    = local.new_vpc_id
  subnet_id = local.new_subnet_id
  vip       = "203.0.113.110"
  security_group_id = ""
  db_version_id = "8.0"
  project_id = var.default_project
//   extranet_access = false
  resource_tags {
	tag_key = "aaa"
	tag_value = "bbb"
  }
}

`
