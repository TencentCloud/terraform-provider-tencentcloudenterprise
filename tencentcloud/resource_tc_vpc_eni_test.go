package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("tencentcloudenterprise_vpc_eni", &resource.Sweeper{
		Name: "tencentcloudenterprise_vpc_eni",
		F:    testSweepEniInstance,
	})
}

func testSweepEniInstance(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	sharedClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("getting tencentcloud client error: %s", err.Error())
	}
	client := sharedClient.(*TencentCloudClient)

	vpcService := VpcService{
		client: client.apiV3Conn,
	}

	instances, err := vpcService.DescribeEniByFilters(ctx, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("get instance list error: %s", err.Error())
	}

	for _, v := range instances {
		instanceId := *v.NetworkInterfaceId
		instanceName := v.NetworkInterfaceName

		now := time.Now()

		createTime := stringTotime(*v.CreatedTime)
		interval := now.Sub(createTime).Minutes()
		if instanceName != nil {
			if strings.HasPrefix(*instanceName, keepResource) || strings.HasPrefix(*instanceName, defaultResource) {
				continue
			}
		}

		// less than 30 minute, not delete
		if needProtect == 1 && int64(interval) < 30 {
			continue
		}

		if err = vpcService.DeleteEni(ctx, instanceId); err != nil {
			log.Printf("[ERROR] sweep instance %s error: %s", instanceId, err.Error())
		}
	}
	return nil
}

func TestAccTencentCloudEni_basic(t *testing.T) {
	t.Parallel()
	var eniId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEniDestroy(&eniId),
		Steps: []resource.TestStep{
			{
				Config: testAccEniBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "name", "ci-test-eni"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "description", "eni desc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "security_groups.#", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count", "1"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_eni.foo", "tags"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "mac"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "state", ENI_STATE_AVAILABLE),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "primary", "false"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "create_time"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_eni.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudEni_updateAttr(t *testing.T) {
	t.Parallel()
	var eniId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEniDestroy(&eniId),
		Steps: []resource.TestStep{
			{
				Config: testAccEniBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "name", "ci-test-eni"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "description", "eni desc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "security_groups.#", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count", "1"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_eni.foo", "tags"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "mac"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "state", ENI_STATE_AVAILABLE),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "primary", "false"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "create_time"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "1"),
				),
			},
			{
				Config: testAccEniUpdateAttr,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "name", "ci-test-eni-new"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "description", "eni desc new"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "security_groups.#", "2"),
				),
			},
			{
				Config: testAccEniUpdateTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "tags.test", "test"),
				),
			},
		},
	})
}

func TestAccTencentCloudEni_updateCount(t *testing.T) {
	t.Parallel()
	var eniId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEniDestroy(&eniId),
		Steps: []resource.TestStep{
			{
				Config: testAccEniBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "name", "ci-test-eni"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "description", "eni desc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "security_groups.#", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count", "1"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_eni.foo", "tags"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "mac"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "state", ENI_STATE_AVAILABLE),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "primary", "false"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "create_time"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "1"),
				),
			},
			{
				Config: testAccEniUpdateCountAdd,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count", "30"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "30"),
				),
			},
			{
				Config: testAccEniUpdateCountSub,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count", "20"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "20"),
				),
			},
		},
	})
}

func TestAccTencentCloudEni_updateManually(t *testing.T) {
	t.Parallel()
	var eniId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEniDestroy(&eniId),
		Steps: []resource.TestStep{
			{
				Config: testAccEniManually,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "name", "ci-test-eni"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "description", "eni desc"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "security_groups.#", "0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "subnet_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4s.#", "1"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_count"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_eni.foo", "tags"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "mac"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "state", ENI_STATE_AVAILABLE),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "primary", "false"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_eni.foo", "create_time"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "1"),
				),
			},
			{
				Config: testAccEniManuallyUpdatePrimaryDesc,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "1"),
				),
			},
			{
				Config: testAccEniManuallyUpdateAdd,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4s.#", "30"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "30"),
				),
			},
			{
				Config: testAccEniManuallyUpdateSub,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEniExists("tencentcloudenterprise_vpc_eni.foo", &eniId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4s.#", "15"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_eni.foo", "ipv4_info.#", "15"),
				),
			},
		},
	})
}

func testAccCheckEniExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no eni id is set")
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		enis, err := service.DescribeEniById(context.TODO(), []string{rs.Primary.ID})
		if err != nil {
			return err
		}

		for _, e := range enis {
			if e.NetworkInterfaceId == nil {
				return errors.New("eni id is nil")
			}

			if *e.NetworkInterfaceId == rs.Primary.ID {
				*id = rs.Primary.ID
				return nil
			}
		}

		return fmt.Errorf("eni not found: %s", rs.Primary.ID)
	}
}

func testAccCheckEniDestroy(id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		service := VpcService{client: client}

		enis, err := service.DescribeEniById(context.TODO(), []string{*id})
		if err != nil {
			return err
		}

		if len(enis) > 0 {
			return errors.New("eni still exists")
		}

		return nil
	}
}

const testAccEniVpc = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "ci-test-eni-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "foo" {
  availability_zone = var.availability_zone
  name              = "ci-test-eni-subnet"
  vpc_id            = cloud_vpc.foo.id
  cidr_block        = "10.0.0.0/16"
  is_multicast      = false
}
`

const testAccEniBasic = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  ipv4_count  = 1
}
`

const testAccEniUpdateAttr = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "test-ci-eni-sg1"
}

resource "tencentcloudenterprise_vpc_security_group" "bar" {
  name = "test-ci-eni-sg2"
}

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name            = "ci-test-eni-new"
  vpc_id          = cloud_vpc.foo.id
  subnet_id       = cloud_vpc_subnet.foo.id
  description     = "eni desc new"
  security_groups = [cloud_vpc_security_group.foo.id, cloud_vpc_security_group.bar.id]
  ipv4_count      = 1
}
`

const testAccEniUpdateTags = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "test-ci-eni-sg1"
}

resource "tencentcloudenterprise_vpc_security_group" "bar" {
  name = "test-ci-eni-sg2"
}

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name            = "ci-test-eni-new"
  vpc_id          = cloud_vpc.foo.id
  subnet_id       = cloud_vpc_subnet.foo.id
  description     = "eni desc new"
  security_groups = [cloud_vpc_security_group.foo.id, cloud_vpc_security_group.bar.id]
  ipv4_count      = 1

  tags = {
    "test" = "test"
  }
}
`

const testAccEniUpdateCountAdd = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  ipv4_count  = 30
}
`

const testAccEniUpdateCountSub = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  ipv4_count  = 20
}
`

const testAccEniManually = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  
  ipv4s {
    ip      = "203.0.113.10"
    primary = true
    description = "desc"
  }
}
`

const testAccEniManuallyUpdatePrimaryDesc = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  
  ipv4s {
    ip          = "203.0.113.10"
    primary     = true
    description = ""
  }
}
`

const testAccEniManuallyUpdateAdd = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  
  ipv4s {
    ip          = "203.0.113.10"
    primary     = true
    description = ""
  }

  ipv4s {
    ip      = "203.0.113.11"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.12"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.13"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.14"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.15"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.16"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.17"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.18"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.19"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.21"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.22"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.23"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.24"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.25"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.26"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.27"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.28"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.29"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.30"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.31"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.32"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.33"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.34"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.35"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.36"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.37"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.38"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.39"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.40"
    primary = false
  }
}
`

const testAccEniManuallyUpdateSub = testAccEniVpc + `

resource "tencentcloudenterprise_vpc_eni" "foo" {
  name        = "ci-test-eni"
  vpc_id      = cloud_vpc.foo.id
  subnet_id   = cloud_vpc_subnet.foo.id
  description = "eni desc"
  
  ipv4s {
    ip          = "203.0.113.10"
    primary     = true
    description = "" // set empty desc to test if SDK can set private IP desc empty or not
  }

  ipv4s {
    ip      = "203.0.113.11"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.12"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.13"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.14"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.15"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.16"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.17"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.18"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.19"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.21"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.22"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.23"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.24"
    primary = false
  }

  ipv4s {
    ip      = "203.0.113.25"
    primary = false
  }
}
`
