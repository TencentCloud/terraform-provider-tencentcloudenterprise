package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("tencentcloudenterprise_route_table", &resource.Sweeper{
		Name: "tencentcloudenterprise_route_table",
		F:    testSweepRouteTable,
	})
}

func testSweepRouteTable(region string) error {
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

	instances, err := vpcService.DescribeRouteTables(ctx, "", "", "", nil, nil, "")
	if err != nil {
		return fmt.Errorf("get instance list error: %s", err.Error())
	}

	for _, v := range instances {

		instanceId := v.routeTableId
		instanceName := v.name
		now := time.Now()
		createTime := stringTotime(v.createTime)
		interval := now.Sub(createTime).Minutes()

		if strings.HasPrefix(instanceName, keepResource) || strings.HasPrefix(instanceName, defaultResource) {
			continue
		}

		if needProtect == 1 && int64(interval) < 30 {
			continue
		}

		if err = vpcService.DeleteRouteTable(ctx, instanceId); err != nil {
			log.Printf("[ERROR] sweep instance %s error: %s", instanceId, err.Error())
		}
	}

	return nil
}

func TestAccDataSourceTencentCloudRouteTable_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudRouteTableConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_route_table.foo"),
					resource.TestCheckResourceAttr("data.cloud_route_table.foo", "name", "tf-ci-test"),
				),
			},
		},
	})
}

const testAccDataSourceTencentCloudRouteTableConfig = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "tf-ci-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_route_table" "route_table" {
  vpc_id = cloud_vpc.foo.id
  name   = "tf-ci-test"
}

data "tencentcloudenterprise_route_table" "foo" {
  route_table_id = cloud_route_table.route_table.id
}
`
