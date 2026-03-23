package tencentcloud

// TEMPORARILY DISABLED - SDK conflict

//
// import (
// 	"fmt"
// 	"testing"
// 	"time"
//
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// )
//
// func TestAccTencentCloudTcrManageReplicationOperationResource_basic(t *testing.T) {
// 	t.Parallel()
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_COMMON) },
// 		Providers: testAccProviders,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: fmt.Sprintf(testAccTcrManageReplicationOperation, "sync", time.Now().Nanosecond()),
// 				PreConfig: func() {
// 					testAccStepSetRegion(t, "ap-shanghai")
// 					testAccPreCheckCommon(t, ACCOUNT_TYPE_COMMON)
// 				},
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "id"),
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "source_registry_id"),
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "destination_registry_id"),
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.#"),
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.name"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.override", "true"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.filters.#", "3"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.filters.0.type", "name"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.filters.1.type", "tag"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "rule.0.filters.2.type", "resource"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "description", "this is the tcr sync operation"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "destination_region_id", "4"),
// 					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "peer_replication_option.#"),
// 					resource.TestCheckResourceAttr("tencentcloudenterprise_tcr_manage_replication_operation.my_replica", "peer_replication_option.0.enable_peer_replication", "false"),
// 				),
// 			},
// 		},
// 	})
// }
//
// const testAccTcrManageReplicationOperation = defaultTCRInstanceData + `
//
// resource "tencentcloudenterprise_tcr_instance" "mytcr_dest" {
// 	name        = "tf-test-tcr-%s"
// 	instance_type = "premium"
// 	delete_bucket = true
//   }
//
// resource "tencentcloudenterprise_tcr_namespace" "myns_dest" {
// 	instance_id 	 = tencentcloudenterprise_tcr_instance.mytcr_dest.id
// 	name			 = "tf_test_ns_dest"
// 	is_public		 = true
// 	is_auto_scan	 = true
// 	is_prevent_vul   = true
// 	severity		 = "medium"
// 	cve_whitelist_items	{
// 		cve_id = "cve-xxxxx"
// 	}
// }
//
// resource "tencentcloudenterprise_tcr_manage_replication_operation" "my_replica" {
//   source_registry_id = local.tcr_id
//   destination_registry_id = tencentcloudenterprise_tcr_instance.mytcr_dest.id
//   rule {
// 		name = "test_sync_%d"
// 		dest_namespace = tencentcloudenterprise_tcr_namespace.myns_dest.name
// 		override = true
// 		filters {
// 			type = "name"
// 			value = join("/", [var.tcr_namespace, "**"])
// 		}
// 		filters {
// 			type = "tag"
// 			value = ""
// 		}
// 		filters {
// 			type = "resource"
// 			value = ""
// 		}
//   }
//   description = "this is the tcr sync operation"
//   destination_region_id = 4 // "ap-shanghai"
//   peer_replication_option {
// 		peer_registry_uin = ""
// 		peer_registry_token = ""
// 		enable_peer_replication = false
//   }
// }
//
// `
