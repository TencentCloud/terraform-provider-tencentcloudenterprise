package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTsfMicroserviceResource_basic -v
func TestAccTencentCloudTsfMicroserviceResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:     func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTsfMicroserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfMicroservice,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTsfMicroserviceExists("tencentcloudenterprise_tsf_microservice.microservice"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_microservice.microservice", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_microservice.microservice", "microservice_name", "test-microservice"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_microservice.microservice", "microservice_desc", "desc-microservice"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_microservice.microservice", "tags.createdBy", "terraform"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tsf_microservice.microservice",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTsfMicroserviceDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tsf_microservice" {
			continue
		}

		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		namespaceId := idSplit[0]
		microserviceId := idSplit[1]

		res, err := service.DescribeTsfMicroserviceById(ctx, namespaceId, microserviceId, "")
		if err != nil {
			return err
		}

		if res != nil {
			return fmt.Errorf("tsf microservice %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckTsfMicroserviceExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}
		idSplit := strings.Split(rs.Primary.ID, FILED_SP)
		if len(idSplit) != 2 {
			return fmt.Errorf("id is broken,%s", rs.Primary.ID)
		}
		namespaceId := idSplit[0]
		microserviceId := idSplit[1]

		service := TsfService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		res, err := service.DescribeTsfMicroserviceById(ctx, namespaceId, microserviceId, "")
		if err != nil {
			return err
		}

		if res == nil {
			return fmt.Errorf("tsf microservice %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTsfMicroserviceVar = `
variable "namespace_id" {
	default = "` + defaultNamespaceId + `"
}
`

const testAccTsfMicroservice = testAccTsfMicroserviceVar + `

resource "tencentcloudenterprise_tsf_microservice" "microservice" {
	namespace_id = "namespace-6ymb2bvg"
	microservice_name = "test-microservice"
	microservice_desc = "desc-microservice"
	tags = {
	  "createdBy" = "terraform"
	}
}

`
