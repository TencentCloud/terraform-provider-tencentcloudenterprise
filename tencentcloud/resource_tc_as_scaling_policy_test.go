package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudAsScalingPolicy(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAsScalingPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAsScalingPolicy(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAsScalingPolicyExists("tencentcloudenterprise_as_scaling_policy.scaling_policy"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_as_scaling_policy.scaling_policy", "scaling_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "policy_name", "tf-as-scaling-policy"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "adjustment_type", "EXACT_CAPACITY"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "adjustment_value", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "comparison_operator", "GREATER_THAN"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "metric_name", "CPU_UTILIZATION"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "threshold", "80"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "period", "300"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "continuous_time", "10"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "statistic", "AVERAGE"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "cooldown", "360"),
				),
			},
			// test update case
			{
				Config: testAccAsScalingPolicy_update(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAsScalingPolicyExists("tencentcloudenterprise_as_scaling_policy.scaling_policy"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_as_scaling_policy.scaling_policy", "scaling_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "adjustment_type", "CHANGE_IN_CAPACITY"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "adjustment_value", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "comparison_operator", "GREATER_THAN_OR_EQUAL_TO"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "metric_name", "MEM_UTILIZATION"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "threshold", "85"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "period", "60"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "continuous_time", "9"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "statistic", "MAXIMUM"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_as_scaling_policy.scaling_policy", "cooldown", "300"),
				),
			},
		},
	})
}

func testAccCheckAsScalingPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("auto scaling policy %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("auto scaling policy id is not set")
		}
		asService := AsService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		_, has, err := asService.DescribeScalingPolicyById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if has < 1 {
			return fmt.Errorf("auto scaling policy not exists: %s", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckAsScalingPolicyDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	asService := AsService{
		client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_as_scaling_policy" {
			continue
		}

		_, has, err := asService.DescribeScalingPolicyById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if has > 0 {
			return fmt.Errorf("auto scaling policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccAsScalingPolicy() string {
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
  configuration_name = "tf-as-configuration-policy"
  image_id           = "img-9qabwvbn"
  instance_types     = [data.cloud_cvm_instance_types.default.instance_types.0.instance_type]
}

resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name = "tf-as-scaling-group-policy"
  configuration_id   = cloud_as_scaling_config.launch_configuration.id
  max_size           = 1
  min_size           = 0
  vpc_id             = cloud_vpc.vpc.id
  subnet_ids         = [cloud_vpc_subnet.subnet.id]
}

resource "tencentcloudenterprise_as_scaling_policy" "scaling_policy" {
  scaling_group_id    = cloud_as_scaling_group.scaling_group.id
  policy_name         = "tf-as-scaling-policy"
  adjustment_type     = "EXACT_CAPACITY"
  adjustment_value    = 0
  comparison_operator = "GREATER_THAN"
  metric_name         = "CPU_UTILIZATION"
  threshold           = 80
  period              = 300
  continuous_time     = 10
  statistic           = "AVERAGE"
  cooldown            = 360
}
`
}

func testAccAsScalingPolicy_update() string {
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
  configuration_name = "tf-as-configuration-policy"
  image_id           = "img-9qabwvbn"
  instance_types     = [data.cloud_cvm_instance_types.default.instance_types.0.instance_type]
}

resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name = "tf-as-scaling-group-policy"
  configuration_id   = cloud_as_scaling_config.launch_configuration.id
  max_size           = 1
  min_size           = 0
  vpc_id             = cloud_vpc.vpc.id
  subnet_ids         = [cloud_vpc_subnet.subnet.id]
}

resource "tencentcloudenterprise_as_scaling_policy" "scaling_policy" {
  scaling_group_id    = cloud_as_scaling_group.scaling_group.id
  policy_name         = "tf-as-scaling-policy"
  adjustment_type     = "CHANGE_IN_CAPACITY"
  adjustment_value    = 1
  comparison_operator = "GREATER_THAN_OR_EQUAL_TO"
  metric_name         = "MEM_UTILIZATION"
  threshold           = 85
  period              = 60
  continuous_time     = 9
  statistic           = "MAXIMUM"
  cooldown            = 300
}
`
}
