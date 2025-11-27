package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var testDataAddressTemplateGroupsNameAll = "data.cloud_vpc_address_template_groups.all_test"

func TestAccTencentCloudDataAddressTemplateGroups(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAddressTemplateGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataAddressTemplateGroupsBasic,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAddressTemplateGroupExists("tencentcloudenterprise_vpc_address_template_group.mygroup"),
					resource.TestCheckResourceAttrSet(testDataAddressTemplateGroupsNameAll, "group_list.0.id"),
					resource.TestCheckResourceAttr(testDataAddressTemplateGroupsNameAll, "group_list.0.name", "mygroup"),
					resource.TestCheckResourceAttr(testDataAddressTemplateGroupsNameAll, "group_list.0.template_ids.#", "1"),
				),
			},
		},
	})
}

const testAccTencentCloudDataAddressTemplateGroupsBasic = `
resource "tencentcloudenterprise_vpc_address_template" "myaddress_template" {
  name        = "testacctcrtemplate"
  addresses = ["1.1.1.1"]
}

resource "tencentcloudenterprise_vpc_address_template_group" "mygroup" {
  name        = "mygroup"
  template_ids = [cloud_vpc_address_template.myaddress_template.id]
}

data "tencentcloudenterprise_vpc_address_template_groups" "all_test" {
  name = cloud_vpc_address_template_group.mygroup.name
  id = cloud_vpc_address_template_group.mygroup.id
}

`
