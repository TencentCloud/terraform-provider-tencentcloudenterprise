package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const BasicClbName = "tf-clb-basic"
const SnatClbName = "tf-clb-snat"
const InternalClbName = "tf-clb-internal"
const InternalClbNameUpdate = "tf-clb-update-internal"
const SingleClbName = "single-open-clb"
const MultiClbName = "multi-open-clb"
const OpenClbName = "tf-clb-open"
const OpenClbNameUpdate = "tf-clb-update-open"

func init() {
	// go test -v ./tencentcloud -sweep=ap-guangzhou -sweep-run=tencentcloudenterprise_clb_instance
	resource.AddTestSweepers("tencentcloudenterprise_clb_instance", &resource.Sweeper{
		Name: "tencentcloudenterprise_clb_instance",
		F:    testSweepClbInstance,
	})
}

func testSweepClbInstance(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cli, err := sharedClientForRegion(region)
	if err != nil {
		return err
	}
	client := cli.(*TencentCloudClient).apiV3Conn
	service := ClbService{client: client}

	res, err := service.DescribeLoadBalancerByFilter(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	if len(res) > 0 {
		for _, v := range res {
			id := *v.LoadBalancerId
			instanceName := *v.LoadBalancerName
			createTime := stringTotime(*v.CreateTime)

			now := time.Now()
			interval := now.Sub(createTime).Minutes()
			// keep not delete
			if strings.HasPrefix(instanceName, keepResource) || strings.HasPrefix(instanceName, defaultResource) {
				continue
			}
			// less than 30 minute, not delete
			if needProtect == 1 && int64(interval) < 30 {
				continue
			}
			if err := service.DeleteLoadBalancerById(ctx, id); err != nil {
				log.Printf("Delete %s error: %s", id, err.Error())
				continue
			}
		}
	}
	return nil
}

func TestAccTencentCloudClbInstance_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "clb_name", BasicClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "tags.test", "tf"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "tags.test1", "tf1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_instance.clb_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbInstance_open(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance_open,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_open"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "clb_name", OpenClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "project_id", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "security_groups.#", "1"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_open", "security_groups.0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "target_region_info_region", "ap-guangzhou"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_open", "target_region_info_vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "tags.test", "tf"),
				),
			},
			{
				Config: testAccClbInstance_update_open,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_open"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "clb_name", OpenClbNameUpdate),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "project_id", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_open", "vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "security_groups.#", "1"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_open", "security_groups.0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "target_region_info_region", "ap-guangzhou"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_open", "target_region_info_vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_open", "tags.test", "test"),
				),
			},
		},
	})
}

func TestAccTencentCloudClbInstance_snat(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance_snat,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "clb_name", SnatClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_basic", "snat_pro", "true"),
				),
			},
		},
	})
}

func TestAccTencentCloudClbInstance_internal(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance_internal,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_internal"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "clb_name", InternalClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "network_type", "INTERNAL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "project_id", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_internal", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_internal", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "tags.test", "tf1"),
				),
			},
			{
				Config: testAccClbInstance_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_internal"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "clb_name", InternalClbNameUpdate),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "network_type", "INTERNAL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "project_id", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_internal", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.clb_internal", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal", "tags.test", "test"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_instance.clb_internal",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudClbInstance_exclusiveSpec(t *testing.T) {
	t.Parallel()

	tgwLabel, stgwLabel := testAccClbExclusiveLabels(t)
	name := fmt.Sprintf("tf-clb-exclusive-%d", time.Now().Unix())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbInstance_exclusiveSpec, name, tgwLabel, stgwLabel),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.exclusive"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.exclusive", "network_type", "INTERNAL"),
					resource.TestCheckTypeSetElemAttr("tencentcloudenterprise_clb_instance.exclusive", "tgw_set_labels.*", tgwLabel),
					resource.TestCheckTypeSetElemAttr("tencentcloudenterprise_clb_instance.exclusive", "stgw_set_labels.*", stgwLabel),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_instance.exclusive",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccClbExclusiveLabels(t *testing.T) (string, string) {
	t.Helper()
	testAccPreCheck(t)

	client, err := sharedClientForRegion(os.Getenv("TENCENTCLOUD_REGION"))
	if err != nil {
		t.Fatalf("create CLB acceptance test client: %s", err)
	}
	response, err := client.(*TencentCloudClient).apiV3Conn.UseClbClient().DescribeAppIdLabel(clb.NewDescribeAppIdLabelRequest())
	if err != nil {
		t.Fatalf("query exclusive CLB labels: %s", err)
	}
	if response == nil || response.Response == nil {
		t.Skip("exclusive CLB label query returned an empty response")
	}

	var tgwLabel, stgwLabel string
	for _, item := range response.Response.OwnerLabelSet {
		if item == nil || item.Label == nil || item.SetType == nil {
			continue
		}
		switch *item.SetType {
		case "L4_LAN_CLB":
			tgwLabel = *item.Label
		case "L7_LAN_CLB":
			stgwLabel = *item.Label
		}
	}
	if tgwLabel == "" || stgwLabel == "" {
		t.Skip("the current account has no complete L4/L7 exclusive CLB specification")
	}
	return tgwLabel, stgwLabel
}

func TestAccTencentCloudClbInstance_internalVip(t *testing.T) {
	t.Parallel()

	name := fmt.Sprintf("tf-clb-internal-vip-%d", time.Now().Unix())
	vip := "10.12.23.12"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccClbInstance_internalVip, name, vip),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.clb_internal_vip"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal_vip", "clb_name", name),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal_vip", "network_type", "INTERNAL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal_vip", "vip", vip),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.clb_internal_vip", "clb_vips.0", vip),
				),
			},
		},
	})
}

func TestAccTencentCloudClbInstance_default_enable(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance_default_enable,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.default_enable"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "clb_name", SingleClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "project_id", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "load_balancer_pass_to_target", "true"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "security_groups.0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "target_region_info_region", "ap-guangzhou"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "target_region_info_vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "tags.test", "open"),
				),
			},
			{
				Config: testAccClbInstance_default_enable_open,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.default_enable"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "clb_name", SingleClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "project_id", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "load_balancer_pass_to_target", "true"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "security_groups.0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "target_region_info_region", "ap-guangzhou"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_instance.default_enable", "target_region_info_vpc_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.default_enable", "tags.test", "hello"),
				),
			},
		},
	})
}

func TestAccTencentCloudClbInstance_multiple_instance(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbInstance__multi_instance,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.multiple_instance"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "clb_name", MultiClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "master_zone_id", "100004"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "slave_zone_id", "100003"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "tags.test", "mytest"),
				),
			},
			{
				Config: testAccClbInstance__multi_instance_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbInstanceExists("tencentcloudenterprise_clb_instance.multiple_instance"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "network_type", "OPEN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "clb_name", MultiClbName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "master_zone_id", "100004"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "slave_zone_id", "100003"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_instance.multiple_instance", "tags.test", "open"),
				),
			},
		},
	})
}

func testAccCheckClbInstanceDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbService := ClbService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_clb_instance" {
			continue
		}

		instance, err := clbService.DescribeLoadBalancerById(ctx, rs.Primary.ID)
		if instance != nil && err == nil {
			return fmt.Errorf("[CHECK][CLB instance][Destroy] check: CLB instance still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckClbInstanceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB instance][Exists] check: CLB instance %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB instance][Exists] check: CLB instance id is not set")
		}
		clbService := ClbService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		instance, err := clbService.DescribeLoadBalancerById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CLB instance][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

const testAccClbInstance_basic = `
resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "` + BasicClbName + `"
  tags = {
    test = "tf"
    test1 = "tf1"
  }
}
`

const testAccClbInstance_snat = `
data "tencentcloudenterprise_vpc_instances" "gz3vpc" {
  name = "Default-"
  is_default = true
}

data "tencentcloudenterprise_vpc_subnets" "gz3" {
  vpc_id = data.tencentcloudenterprise_vpc_instances.gz3vpc.instance_list.0.vpc_id
}

locals {
  keep_clb_subnets = [for subnet in data.tencentcloudenterprise_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") if lookup(subnet, "name") == "keep-clb-sub"]
  subnets = [for subnet in data.tencentcloudenterprise_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") ]
  subnet_for_clb_snat = concat(local.keep_clb_subnets, local.subnets)
}

resource "tencentcloudenterprise_clb_instance" "clb_basic" {
  network_type = "OPEN"
  clb_name     = "` + SnatClbName + `"
  snat_pro     = true
  snat_ips {
	subnet_id = local.subnet_for_clb_snat.0
  }
  snat_ips {
    subnet_id = local.subnet_for_clb_snat.1
  }
}
`

const testAccClbInstance_internal = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "clb-instance-internal-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "guagua-ci-temp-test"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "tencentcloudenterprise_clb_instance" "clb_internal" {
  network_type = "INTERNAL"
  clb_name     = "` + InternalClbName + `"
  vpc_id       = tencentcloudenterprise_vpc.foo.id
  subnet_id    = tencentcloudenterprise_vpc_subnet.subnet.id
  project_id   = 0

  tags = {
    test = "tf1"
  }
}
`

const testAccClbInstance_exclusiveSpec = `
data "tencentcloudenterprise_vpc_instances" "available" {
  is_default = true
}

data "tencentcloudenterprise_vpc_subnets" "available" {
  vpc_id = data.tencentcloudenterprise_vpc_instances.available.instance_list.0.vpc_id
}

resource "tencentcloudenterprise_clb_instance" "exclusive" {
  network_type    = "INTERNAL"
  clb_name        = "%s"
  vpc_id          = data.tencentcloudenterprise_vpc_instances.available.instance_list.0.vpc_id
  subnet_id       = data.tencentcloudenterprise_vpc_subnets.available.instance_list.0.subnet_id
  tgw_set_labels  = ["%s"]
  stgw_set_labels = ["%s"]
}
`

const testAccClbInstance_open = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "keep-ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "clb-instance-open-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_clb_instance" "clb_open" {
  network_type              = "OPEN"
  clb_name                  = "` + OpenClbName + `"
  project_id                = 0
  vpc_id                    = tencentcloudenterprise_vpc.foo.id
  target_region_info_region = "ap-guangzhou"
  target_region_info_vpc_id = tencentcloudenterprise_vpc.foo.id
  security_groups           = [tencentcloudenterprise_vpc_security_group.foo.id]

  tags = {
    test = "tf"
  }
}
`

const testAccClbInstance_internalVip = `
resource "tencentcloudenterprise_clb_instance" "clb_internal_vip" {
  network_type = "INTERNAL"
  clb_name     = "%s"
  project_id   = 0
  vpc_id       = "vpc-4tapjns3"
  subnet_id    = "subnet-92tohwk0"
  vip          = "%s"

  tags = {
    test = "tf-vip"
  }
}
`

const testAccClbInstance_update = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "clb-instance-internal-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "guagua-ci-temp-test"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "tencentcloudenterprise_clb_instance" "clb_internal" {
  network_type = "INTERNAL"
  clb_name     = "` + InternalClbNameUpdate + `"
  vpc_id       = tencentcloudenterprise_vpc.foo.id
  subnet_id    = tencentcloudenterprise_vpc_subnet.subnet.id
  project_id   = 0

  tags = {
    test = "test"
  }
}
`

const testAccClbInstance_update_open = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "clb-instance-sg"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "clb-instance-open-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_clb_instance" "clb_open" {
  network_type              = "OPEN"
  clb_name                  = "` + OpenClbNameUpdate + `"
  vpc_id                    = tencentcloudenterprise_vpc.foo.id
  project_id                = 0
  target_region_info_region = "ap-guangzhou"
  target_region_info_vpc_id = tencentcloudenterprise_vpc.foo.id
  security_groups           = [tencentcloudenterprise_vpc_security_group.foo.id]

  tags = {
    test = "test"
  }
}
`

const testAccClbInstance_default_enable = `
variable "availability_zone" {
  default = "ap-guangzhou-1"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "keep-sdk-feature-test"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "tencentcloudenterprise_vpc_security_group" "sglab" {
  name        = "clb-instance-enable-sg"
  description = "favourite sg"
  project_id  = 0
}

resource "tencentcloudenterprise_vpc" "foo" {
  name         = "clb-instance-default-vpc"
  cidr_block   = "10.0.0.0/16"

  tags = {
    "test" = "mytest"
  }
}

resource "tencentcloudenterprise_clb_instance" "default_enable" {
  network_type                 = "OPEN"
  clb_name                     = "` + SingleClbName + `"
  project_id                   = 0
  vpc_id                       = tencentcloudenterprise_vpc.foo.id
  load_balancer_pass_to_target = true

  security_groups              = [tencentcloudenterprise_vpc_security_group.sglab.id]
  target_region_info_region    = "ap-guangzhou"
  target_region_info_vpc_id    = tencentcloudenterprise_vpc.foo.id

  tags = {
    test = "open"
  }
}
`

const testAccClbInstance_default_enable_open = `
variable "availability_zone" {
  default = "ap-guangzhou-1"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "keep-sdk-feature-test"
  vpc_id            = tencentcloudenterprise_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "tencentcloudenterprise_vpc_security_group" "sglab" {
  name        = "clb-instance-enable-sg"
  description = "favourite sg"
  project_id  = 0
}

resource "tencentcloudenterprise_vpc" "foo" {
  name         = "clb-instance-default-vpc"
  cidr_block   = "10.0.0.0/16"

  tags = {
    "test" = "mytest"
  }
}

resource "tencentcloudenterprise_clb_instance" "default_enable" {
  network_type                 = "OPEN"
  clb_name                     = "` + SingleClbName + `"
  project_id                   = 0
  vpc_id                       = tencentcloudenterprise_vpc.foo.id
  load_balancer_pass_to_target = true

  security_groups              = [tencentcloudenterprise_vpc_security_group.sglab.id]
  target_region_info_region    = "ap-guangzhou"
  target_region_info_vpc_id    = tencentcloudenterprise_vpc.foo.id

  tags = {
    test = "hello"
  }
}
`

const testAccClbInstance__multi_instance = `
resource "tencentcloudenterprise_clb_instance" "multiple_instance" {
  network_type              = "OPEN"
  clb_name                  = "` + MultiClbName + `"
  master_zone_id = "100004"
  slave_zone_id = "100003"

  tags = {
    test = "mytest"
  }
}
`

const testAccClbInstance__multi_instance_update = `
resource "tencentcloudenterprise_clb_instance" "multiple_instance" {
  network_type              = "OPEN"
  clb_name                  = "` + MultiClbName + `"
  master_zone_id = "100004"
  slave_zone_id = "100003"

  tags = {
    test = "open"
  }
}
`
