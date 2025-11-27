/*
Use this data source to query detailed information of dcdb parameters

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_parameters" "parameters" {
	  instance_id = "your_instance_id"
	  }

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_parameters", CNDescription{
		TerraformTypeCN: "数据库参数",
		DescriptionCN:   "提供DCDB参数数据源，用于查询DCDB参数的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例id",
			"list":               "参数列表",
			"param":              "参数名称",
			"参数值":                "参数值",
			"default":            "默认值",
			"constraint":         "params约束",
			"type":               "类型",
			"enum":               "num类型的可选值列表",
			"range":              "范围约束",
			"min":                "最小值",
			"max":                "最大值",
			"string":             "约束类型为字符串",
			"have_set_具有设定值":     "具有设定值",
			"need_restart":       "需要重新启动",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbParameters() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of dcdb parameters",
		Read:        dataSourceTencentCloudDcdbParametersRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Parameter list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"param": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter name.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter value.",
						},
						"default": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Default value.",
						},
						"constraint": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Params constraint.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type.",
									},
									"enum": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A list of optional values of type num.",
									},
									"range": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Range constraint.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"min": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Min value.",
												},
												"max": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Max value.",
												},
											},
										},
									},
									"string": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Constraint type is string.",
									},
								},
							},
						},
						"have_set_value": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Have set value.",
						},
						"need_restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Need restart.",
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

func dataSourceTencentCloudDcdbParametersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_parameters.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	dcdbService := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var params []*dcdb.ParamDesc
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := dcdbService.DescribeDcdbParametersByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		params = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb params failed, reason:%+v", logId, err)
		return err
	}

	paramList := []interface{}{}
	ids := make([]string, 0, len(params))
	if params != nil {
		for _, param := range params {
			paramMap := map[string]interface{}{}
			if param.Param != nil {
				paramMap["param"] = param.Param
			}
			if param.Value != nil {
				paramMap["value"] = param.Value
			}
			if param.Default != nil {
				paramMap["default"] = param.Default
			}
			if param.Constraint != nil {
				constraintMap := map[string]interface{}{}
				if param.Constraint.Type != nil {
					constraintMap["type"] = param.Constraint.Type
				}
				if param.Constraint.Enum != nil {
					constraintMap["enum"] = param.Constraint.Enum
				}
				if param.Constraint.Range != nil {
					rangeMap := map[string]interface{}{}
					if param.Constraint.Range.Min != nil {
						rangeMap["min"] = param.Constraint.Range.Min
					}
					if param.Constraint.Range.Max != nil {
						rangeMap["max"] = param.Constraint.Range.Max
					}

					constraintMap["range"] = []interface{}{rangeMap}
				}
				if param.Constraint.String != nil {
					constraintMap["string"] = param.Constraint.String
				}

				paramMap["constraint"] = []interface{}{constraintMap}
			}
			if param.HaveSetValue != nil {
				paramMap["have_set_value"] = param.HaveSetValue
			}
			//if param.NeedRestart != nil {
			//	paramMap["need_restart"] = param.NeedRestart
			//}
			ids = append(ids, *param.Param+FILED_SP+*param.Value)
			paramList = append(paramList, paramMap)
		}
		d.SetId(helper.DataResourceIdsHash(ids))
		_ = d.Set("list", paramList)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), paramList); e != nil {
			return e
		}
	}

	return nil
}
