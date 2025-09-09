package tencentcloud

import (
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbSnatIp(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbSnatIpBasic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_clb_snat_ip.snat_ips", "id"),
					resource.TestCheckResourceAttr("cloud_clb_snat_ip.snat_ips", "ips.#", "3"),
				),
			},
			{
				PreConfig: func() {
					time.Sleep(time.Second * 10)
				},
				Config: testAccClbSnatIpBasicUpdate,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_clb_snat_ip.snat_ips", "id"),
					resource.TestCheckResourceAttr("cloud_clb_snat_ip.snat_ips", "ips.#", "3"),
				),
			},
			{
				PreConfig: func() {
					time.Sleep(time.Second * 10)
				},
				Config: testAccClbSnatIpBasicUpdate2,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_clb_snat_ip.snat_ips", "id"),
					resource.TestCheckResourceAttr("cloud_clb_snat_ip.snat_ips", "ips.#", "1"),
				),
			},
		},
	})
}

const testAccClbSnatIpBasic = `
data "cloud_vpc_instances" "gz3vpc" {
  name = "Default-"
  is_default = true
}

data "cloud_vpc_subnets" "gz3" {
  vpc_id = data.cloud_vpc_instances.gz3vpc.instance_list.0.vpc_id
}

locals {
  keep_clb_subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") if lookup(subnet, "name") == "keep-clb-sub"]
  subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") ]
  subnet_for_clb_snat = concat(local.keep_clb_subnets, local.subnets)
}

resource "cloud_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-snat-resource-test"
}

resource "cloud_clb_snat_ip" "snat_ips" {
  clb_id = cloud_clb_instance.foo.id
  ips {
    ip = "203.0.113.17"
	subnet_id = local.subnet_for_clb_snat.0
  }
  ips {
	ip = "203.0.113.15"
	subnet_id = local.subnet_for_clb_snat.0
  }
  ips {
	ip = "203.0.113.138"
	subnet_id = local.subnet_for_clb_snat.1
  }
}
`

const testAccClbSnatIpBasicUpdate = `
data "cloud_vpc_instances" "gz3vpc" {
  name = "Default-"
  is_default = true
}

data "cloud_vpc_subnets" "gz3" {
  vpc_id = data.cloud_vpc_instances.gz3vpc.instance_list.0.vpc_id
}

locals {
  keep_clb_subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") if lookup(subnet, "name") == "keep-clb-sub"]
  subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") ]
  subnet_for_clb_snat = concat(local.keep_clb_subnets, local.subnets)
}

resource "cloud_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-snat-resource-test"
}

resource "cloud_clb_snat_ip" "snat_ips" {
  clb_id = cloud_clb_instance.foo.id
  ips {
    ip = "203.0.113.17"
	subnet_id = local.subnet_for_clb_snat.0
  }
  ips {
	ip = "203.0.113.138"
	subnet_id = local.subnet_for_clb_snat.1
  }
  ips {
	ip = "203.0.113.139"
    subnet_id = local.subnet_for_clb_snat.1
  }
}
`

const testAccClbSnatIpBasicUpdate2 = `
data "cloud_vpc_instances" "gz3vpc" {
  name = "Default-"
  is_default = true
}

data "cloud_vpc_subnets" "gz3" {
  vpc_id = data.cloud_vpc_instances.gz3vpc.instance_list.0.vpc_id
}

locals {
  keep_clb_subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") if lookup(subnet, "name") == "keep-clb-sub"]
  subnets = [for subnet in data.cloud_vpc_subnets.gz3.instance_list: lookup(subnet, "subnet_id") ]
  subnet_for_clb_snat = concat(local.keep_clb_subnets, local.subnets)
}

resource "cloud_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = "tf-clb-snat-resource-test"
}

resource "cloud_clb_snat_ip" "snat_ips" {
  clb_id = cloud_clb_instance.foo.id
  ips {
    ip = "203.0.113.16"
	subnet_id = local.subnet_for_clb_snat.0
  }
}
`
