package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcPeerConnectAcceptExOperationResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{{
			// NOTE: this test requires a real cross-region peering_connection_id in PENDING state.
			// Without valid credentials or a real pending peering connection, this will fail.
			// The test validates the resource schema and SDK call path only.
			Config: testAccVpcPeerConnectAcceptExOperation,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_peer_connect_accept_ex_operation.example", "id"),
			),
		}},
	})
}

// NOTE: replace peering_connection_id with a real pending cross-region peering connection ID for e2e test
const testAccVpcPeerConnectAcceptExOperation = `
resource "tencentcloudenterprise_vpc_peer_connect_accept_ex_operation" "example" {
	peering_connection_id = "pcx-placeholder-pending-cross-region"
}
`
