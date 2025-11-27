/*
Use this data source to query detailed information of dcdb database_objects

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_database_objects" "database_objects" {
	  instance_id = "tdsqlshard-973xatu3"
	  db_name = "SysDB"
	  result_output_file = "database_objects.json"
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
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_database_objects", CNDescription{
		TerraformTypeCN: "数据库实例",
		DescriptionCN:   "提供DCDB数据库对象数据源，用于查询DCDB数据库对象的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例的ID",
			"db_name":            "数据库名称，通过DescribeDatabaseapi获取",
			"tables":             "表列表",
			"table":              "表的名称",
			"views":              "查看列表",
			"view":               "视图的名称",
			"procs":              "程序列表",
			"proc":               "过程的名称",
			"funcs":              "功能列表",
			"func":               "函数的名称",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbDatabaseObjects() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the database_objects of the Cloud database instance.",
		Read:        dataSourceTencentCloudDcdbDatabaseObjectsRead,
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

			"tables": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Table list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"table": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of table.",
						},
					},
				},
			},

			"views": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "View list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"view": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of view.",
						},
					},
				},
			},

			"procs": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Procedure list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of procedure.",
						},
					},
				},
			},

			"funcs": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Function list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"func": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of function.",
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
	}
}

func dataSourceTencentCloudDcdbDatabaseObjectsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_database_objects.read")()
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

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	var result *dcdb.DescribeDatabaseObjectsResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		var e error
		result, e = service.DescribeDcdbDBObjectsByFilter(ctx, paramMap)
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
		tables := result.Response.Tables
		tabList := make([]map[string]interface{}, 0, len(tables))
		if tables != nil {
			for _, databaseTable := range tables {
				databaseTableMap := map[string]interface{}{}

				if databaseTable.Table != nil {
					databaseTableMap["table"] = databaseTable.Table
				}
				tabList = append(tabList, databaseTableMap)
			}
			_ = d.Set("tables", tabList)
			data["tables"] = tabList
		}

		views := result.Response.Views
		viewList := make([]map[string]interface{}, 0, len(views))
		if views != nil {
			for _, databaseView := range views {
				databaseViewMap := map[string]interface{}{}

				if databaseView.View != nil {
					databaseViewMap["view"] = databaseView.View
				}
				viewList = append(viewList, databaseViewMap)
			}
			_ = d.Set("views", viewList)
			data["views"] = viewList
		}

		procs := result.Response.Procs
		procList := make([]map[string]interface{}, 0, len(procs))
		if procs != nil {
			for _, databaseProcedure := range procs {
				databaseProcedureMap := map[string]interface{}{}

				if databaseProcedure.Proc != nil {
					databaseProcedureMap["proc"] = databaseProcedure.Proc
				}
				procList = append(procList, databaseProcedureMap)
			}
			_ = d.Set("procs", procList)
			data["procs"] = procList
		}

		funcs := result.Response.Funcs
		funcList := make([]map[string]interface{}, 0, len(funcs))
		if funcs != nil {
			for _, databaseFunction := range funcs {
				databaseFunctionMap := map[string]interface{}{}

				if databaseFunction.Func != nil {
					databaseFunctionMap["func"] = databaseFunction.Func
				}
				funcList = append(funcList, databaseFunctionMap)
			}
			_ = d.Set("funcs", funcList)
			data["funcs"] = funcList
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
