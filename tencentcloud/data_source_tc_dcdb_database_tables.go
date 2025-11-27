/*
Use this data source to query detailed information of dcdb database_tables

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_database_tables" "database_tables" {
	  instance_id = "tdsqlshard-973xatu3"
	  db_name = "SysDB"
	  table = "StatusTable"
	  result_output_file = "database_tables.json"
	}

```
*/
package tencentcloud

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_database_tables", CNDescription{
		TerraformTypeCN: "实例表信息",
		DescriptionCN:   "提供DCDB数据库表数据源，用于查询DCDB数据库表的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例的ID",
			"db_name":            "数据库名称，通过DescribeDatabaseapi获取",
			"table":              "表名，通过DescribeDatabaseObjects api获得",
			"cols":               "列信息",
			"col":                "列的名称",
			"type":               "列类型",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbDatabaseTables() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDcdbDatabaseTablesRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"db_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Database name, obtained through the DescribeDatabases api.",
			},

			"table": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Table name, obtained through the DescribeDatabaseObjects api.",
			},

			"cols": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Column information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"col": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of column.",
						},
						"type": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Column type.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
		Description: "Query detailed information of dcdb database_tables.",
	}
}

func dataSourceTencentCloudDcdbDatabaseTablesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_database_tables.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("db_name"); ok {
		paramMap["db_name"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("table"); ok {
		paramMap["table"] = helper.String(v.(string))
	}

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	var result *dcdb.DescribeDatabaseTableResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var e error
		result, e = service.DescribeDcdbDBTablesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0)
	data := make(map[string]interface{})

	if result != nil {
		colums := result.Response.Cols
		colList := make([]map[string]interface{}, 0, len(colums))
		if colums != nil {
			for _, databaseCol := range colums {
				databaseColMap := map[string]interface{}{}

				if databaseCol.Col != nil {
					databaseColMap["col"] = databaseCol.Col
				}
				if databaseCol.Type != nil {
					databaseColMap["type"] = databaseCol.Type
				}
				colList = append(colList, databaseColMap)
			}
			_ = d.Set("cols", colList)
			data["cols"] = colList
		}
	}

	ids = append(ids, strings.Join([]string{*result.Response.InstanceId, *result.Response.DbName}, FILED_SP))

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), data); e != nil {
			return e
		}
	}
	return nil
}
