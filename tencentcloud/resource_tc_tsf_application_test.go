package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfApplicationResource_basic -v
func TestAccTencentCloudTsfApplicationResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfApplication,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfApplicationExists("cloud_tsf_application.application"),
					resource.TestCheckResourceAttrSet("cloud_tsf_application.application", "id"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "application_desc", "This is my application"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "application_name", "terraform-test"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "application_resource_type", "DEF"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "application_runtime_type", "Java"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "application_type", "V"),
					//resource.TestCheckResourceAttr("cloud_tsf_application.application", "ignore_create_image_repository", "true"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "microservice_type", "N"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.#", "1"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.name", "my-service"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.health_check.#", "1"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.health_check.0.path", "/health"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.ports.#", "1"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.ports.0.protocol", "HTTP"),
					resource.TestCheckResourceAttr("cloud_tsf_application.application", "service_config_list.0.ports.0.target_port", "8080"),
				),
			},
			// {
			// 	ResourceName:      "cloud_tsf_application.application",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckTsfApplicationDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_tsf_application_release_config" {
			continue
		}

		res, err := service.DescribeTsfApplicationById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf Application %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfApplicationExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfApplicationById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf Application %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfApplication = `

resource "cloud_tsf_application" "application" {
	application_name = "terraform-test"
	application_type = "V"
	microservice_type = "N"
	application_desc = "This is my application"
	application_runtime_type = "Java"
	service_config_list {
	  name = "my-service"
	  ports {
		target_port = 8080
		protocol = "HTTP"
	  }
	  health_check {
		path = "/health"
	  }
	}
	// ignore_create_image_repository = true
}

`
