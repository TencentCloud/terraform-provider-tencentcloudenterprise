package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfClusterResource_basic -v
func TestAccTencentCloudTsfClusterResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfCluster,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfClusterExists("tencentcloudenterprise_tsf_cluster.cluster"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_cluster.cluster", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_name", "terraform-test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "vpc_id", "vpc-kphn8u93"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_cidr", "9.165.120.0/24"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_desc", "test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "tsf_region_id", "ap-guangzhou"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_version", "1.18.4"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "max_node_pod_num", "32"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "max_cluster_service_num", "128"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "tags.createdBy", "terraform"),
				),
			},
			// {
			// 	ResourceName:      "tencentcloudenterprise_tsf_cluster.cluster",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

// go test -i; go test -test.run TestAccTencentCloudTsfClusterVResource_basic -v
// 创建虚拟集群
func TestAccTencentCloudTsfClusterVResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckTsfClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfClusterV,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfClusterExists("tencentcloudenterprise_tsf_cluster.cluster"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_cluster.cluster", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_name", "terraform-test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "vpc_id", "vpc-2l7uk2q1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "cluster_desc", "test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "tsf_region_id", "chongqing"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_cluster.cluster", "tags.createdBy", "terraform"),
				),
			},
			// {
			// 	ResourceName:      "tencentcloudenterprise_tsf_cluster.cluster",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckTsfClusterDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tsf_cluster" {
			continue
		}

		res, err := service.DescribeTsfClusterById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf cluster %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfClusterExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfClusterById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf cluster %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfCluster = `

resource "tencentcloudenterprise_tsf_cluster" "cluster" {
	cluster_name = "terraform-test"
	cluster_type = "C"
	vpc_id = "vpc-2l7uk2q1"
	cluster_cidr = "9.165.120.0/24"
	cluster_desc = "test"
	tsf_region_id = "ap-guangzhou"
	cluster_version = "1.18.4"
	max_node_pod_num = 32
	max_cluster_service_num = 128
	tags = {
	  "createdBy" = "terraform"
	}
}
`

const testAccTsfClusterV = `
resource "tencentcloudenterprise_tsf_cluster" "cluster" {
	cluster_name = "terraform-test"
	cluster_type = "V"
	vpc_id = "vpc-2l7uk2q1"
	cluster_desc = "test"
	tsf_region_id = "chongqing"
	tags = {
	  "createdBy" = "terraform"
	}
}
`
