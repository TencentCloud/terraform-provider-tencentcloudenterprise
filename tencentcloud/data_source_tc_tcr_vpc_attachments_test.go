package tencentcloud

// TEMPORARILY DISABLED - SDK conflict

//
// import (
// 	"testing"
//
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
// )
//
// var testDataTCRVPCAttachmentsNameAll = "data.tencentcloudenterprise_tcr_vpc_attachments.id_test"
//
// func TestAccTencentCloudTcrVPCAttachmentsData(t *testing.T) {
// 	t.Parallel()
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_COMMON) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckTCRNamespaceDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccTencentCloudDataTCRVPCAttachmentsBasic,
// 				PreConfig: func() {
// 					testAccStepSetRegion(t, "ap-shanghai")
// 					testAccPreCheckCommon(t, ACCOUNT_TYPE_COMMON)
// 				},
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckTCRVPCAttachmentExists("tencentcloudenterprise_tcr_vpc_attachment.mytcr_vpc_attachment"),
// 					resource.TestCheckResourceAttr(testDataTCRVPCAttachmentsNameAll, "vpc_attachment_list.#", "1"),
// 					resource.TestCheckResourceAttrSet(testDataTCRVPCAttachmentsNameAll, "vpc_attachment_list.0.status"),
// 				),
// 			},
// 		},
// 	})
// }
//
// const defaultTcrVpcSubnets = `
//
// data "tencentcloudenterprise_vpc_subnets" "sh" {
//   availability_zone = "ap-shanghai-1"
// }
//
// locals {
//   vpc_id = data.tencentcloudenterprise_vpc_subnets.sh.instance_list.0.vpc_id
//   subnet_id = data.tencentcloudenterprise_vpc_subnets.sh.instance_list.0.subnet_id
// }`
//
// const testAccTencentCloudDataTCRVPCAttachmentsBasic = defaultTcrVpcSubnets + `
// resource "tencentcloudenterprise_tcr_instance" "mytcr_instance" {
//   name        = "test-tcr-attach"
//   instance_type = "basic"
//   delete_bucket = true
//
//   tags ={
// 	test = "test"
//   }
// }
//
// resource "tencentcloudenterprise_tcr_vpc_attachment" "mytcr_vpc_attachment" {
//   instance_id = tencentcloudenterprise_tcr_instance.mytcr_instance.id
//   vpc_id = local.vpc_id
//   subnet_id = local.subnet_id
// }
//
// data "tencentcloudenterprise_tcr_vpc_attachments" "id_test" {
//   instance_id = tencentcloudenterprise_tcr_vpc_attachment.mytcr_vpc_attachment.instance_id
// }
// `
