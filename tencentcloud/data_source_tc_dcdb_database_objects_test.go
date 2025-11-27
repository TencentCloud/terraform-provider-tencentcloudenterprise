package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDCDBDatabaseObjectsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDCDBDatabaseObjectsDataSource, defaultDcdbInstanceId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_dcdb_database_objects.database_objects"),
					resource.TestCheckResourceAttrSet("data.cloud_dcdb_database_objects.database_objects", "tables.#"),
					resource.TestCheckResourceAttrSet("data.cloud_dcdb_database_objects.database_objects", "views.#"),
					resource.TestCheckResourceAttrSet("data.cloud_dcdb_database_objects.database_objects", "procs.#"),
					resource.TestCheckResourceAttrSet("data.cloud_dcdb_database_objects.database_objects", "funcs.#"),
					resource.TestCheckResourceAttr("data.cloud_dcdb_database_objects.database_objects", "db_name", "SysDB"),
				),
			},
		},
	})
}

const testAccDCDBDatabaseObjectsDataSource = `

data "tencentcloudenterprise_dcdb_database_objects" "database_objects" {
	instance_id = "%s"
	db_name = "SysDB"
}

`
