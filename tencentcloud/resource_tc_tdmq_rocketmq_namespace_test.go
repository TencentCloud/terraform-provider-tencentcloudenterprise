package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudTdmqRocketmqNamespaceResource_basic(t *testing.T) {
	t.Parallel()
	terraformId := "tencentcloudenterprise_tdmq_rocketmq_namespace.namespace"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTdmqRocketmqNamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqRocketmqNamespace,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRocketmqNamespaceExists(terraformId),
					resource.TestCheckResourceAttr(terraformId, "ttl", "65000"),
					resource.TestCheckResourceAttr(terraformId, "retention_time", "65000"),
					resource.TestCheckResourceAttr(terraformId, "remark", "test namespace"),
				),
			},
			{
				Config: testAccTdmqRocketmqNamespaceUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRocketmqNamespaceExists(terraformId),
					resource.TestCheckResourceAttr(terraformId, "ttl", "66000"),
					resource.TestCheckResourceAttr(terraformId, "retention_time", "66000"),
					resource.TestCheckResourceAttr(terraformId, "remark", "test namespace update"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tdmq_rocketmq_namespace.namespace",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTdmqRocketmqNamespaceDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqRocketmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tdmq_rocketmq_namespace" {
			continue
		}
		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		clusterId := idSplit[0]
		namespaceName := idSplit[1]
		namespaces, err := service.DescribeTdmqRocketmqNamespace(ctx, namespaceName, clusterId)
		if err != nil {
			if sdkerr, ok := err.(*sdkErrors.CloudSDKError); ok {
				if sdkerr.Code == "ResourceUnavailable" {
					return nil
				}
			}
			return err
		}

		if len(namespaces) != 0 {
			return fmt.Errorf("Rocketmq namespace still exist, id: %v", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckTdmqRocketmqNamespaceExists(re string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[re]
		if !ok {
			return fmt.Errorf("Rocketmq namespace  %s is not found", re)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Rocketmq namespace id is not set")
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		clusterId := idSplit[0]
		namespaceName := idSplit[1]
		service := TdmqRocketmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		namespaces, err := service.DescribeTdmqRocketmqNamespace(ctx, namespaceName, clusterId)

		if err != nil {
			return err
		}

		if len(namespaces) == 0 {
			return fmt.Errorf("Rocketmq namespace not found, id: %v", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTdmqRocketmqNamespace = `
resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	cluster_name = "test_rocketmq"
	remark = "test recket mq"
}
  
resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	namespace_name = "test_namespace"
	ttl = 65000
	retention_time = 65000
	remark = "test namespace"
}
`

const testAccTdmqRocketmqNamespaceUpdate = `
resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	cluster_name = "test_rocketmq"
	remark = "test recket mq"
}
  
resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	namespace_name = "test_namespace"
	ttl = 66000
	retention_time = 66000
	remark = "test namespace update"
}
`
