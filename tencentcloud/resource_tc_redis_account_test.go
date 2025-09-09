package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudRedisAccountResource_basic -v
func TestAccTencentCloudRedisAccountResource_basic(t *testing.T) {
	// t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccTencentCloudRedisAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRedisAccount,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccTencentCloudRedisAccountExists("cloud_redis_account.account"),
					resource.TestCheckResourceAttrSet("cloud_redis_account.account", "id"),
					resource.TestCheckResourceAttr("cloud_redis_account.account", "instance_id", defaultCrsInstanceId),
					resource.TestCheckResourceAttr("cloud_redis_account.account", "account_name", "account_test"),
					resource.TestCheckResourceAttr("cloud_redis_account.account", "remark", "redis desc"),
					resource.TestCheckResourceAttr("cloud_redis_account.account", "readonly_policy.#", "1"),
					resource.TestCheckResourceAttr("cloud_redis_account.account", "privilege", "r"),
				),
			},
			{
				ResourceName:            "cloud_redis_account.account",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"account_password"},
			},
		},
	})
}

func testAccTencentCloudRedisAccountExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		items := strings.Split(rs.Primary.ID, FILED_SP)
		if len(items) != 2 {
			return fmt.Errorf("invalid ID %s", rs.Primary.ID)
		}
		instanceId := items[0]
		accountName := items[1]

		service := RedisService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		account, err := service.DescribeRedisAccountById(ctx, instanceId, accountName)
		if err != nil {
			return err
		}
		if account == nil {
			return fmt.Errorf("redis account %s is not found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccTencentCloudRedisAccountDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := RedisService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_redis_account" {
			continue
		}
		time.Sleep(5 * time.Second)

		items := strings.Split(rs.Primary.ID, FILED_SP)
		if len(items) != 2 {
			return fmt.Errorf("invalid ID %s", rs.Primary.ID)
		}
		instanceId := items[0]
		accountName := items[1]

		account, err := service.DescribeRedisAccountById(ctx, instanceId, accountName)
		if err != nil {
			return err
		}
		if account != nil {
			return fmt.Errorf("redis account %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

const testAccRedisAccountVar = `
variable "instance_id" {
	default = "` + defaultCrsInstanceId + `"
}
`

const testAccRedisAccount = testAccRedisAccountVar + `

resource "cloud_redis_account" "account" {
	instance_id 	 = var.instance_id
	account_name 	 = "account_test"
	account_password = "test1234"
	remark 			 = "redis desc"
	readonly_policy  = ["master"]
	privilege 		 = "r"
  }

`
