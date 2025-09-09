package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbListenersDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClbListenersDataSource,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckClbListenerExists("cloud_clb_listener.listener"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_listeners.listeners", "listener_list.0.clb_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_listeners.listeners", "listener_list.0.listener_id"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.listener_name", "mylistener1234"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.port", "1"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.protocol", "TCP"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.session_expire_time", "30"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.scheduler", "WRR"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_type", "HTTP"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_port", "0"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_http_code", "16"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_http_path", "/"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_http_domain", "www.tencent.com"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_http_method", "HEAD"),
					resource.TestCheckResourceAttr("data.cloud_clb_listeners.listeners", "listener_list.0.health_check_http_version", "HTTP/1.1"),
				),
			},
		},
	})
}

const testAccClbListenersDataSource = `
resource "cloud_clb_instance" "clb" {
    network_type = "OPEN"
    clb_name     = "tf-clb-listeners"
}

resource "cloud_clb_listener" "listener" {
    clb_id                     = cloud_clb_instance.clb.id
    port                       = 1
    protocol                   = "TCP"
    listener_name              = "mylistener1234"
    session_expire_time        = 30
    scheduler                  = "WRR"
    health_check_type          = "HTTP"
    health_check_http_domain   = "www.tencent.com"
    health_check_http_code     = 16
    health_check_http_version  = "HTTP/1.1"
    health_check_http_method   = "HEAD"
    health_check_http_path     = "/"
}

data "cloud_clb_listeners" "listeners" {
    clb_id      = cloud_clb_instance.clb.id
    listener_id = cloud_clb_listener.listener.listener_id
}
`
