/*
Use this data source to query detailed information of tsf api_detail

# Example Usage

```hcl

	data "tencentcloudenterprise_tsf_api_detail" "api_detail" {
	  microservice_id = "ms-yq3jo6jd"
	  path = "/printRequest"
	  method = "GET"
	  pkg_version = "20210625192923"
	  application_id = "application-a24x29xv"
	}

```
*/
package tencentcloud

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tsf_api_detail", CNDescription{
		TerraformTypeCN: "TSF API详情",
		DescriptionCN:   "提供TSF API详情数据源，用于查询TSF API详情的详细信息。",
		AttributesCN: map[string]string{
			"microservice_id":      "微服务ID",
			"path":                 "API路径",
			"method":               "请求方法",
			"pkg_version":          "包版本",
			"application_id":       "应用ID",
			"result":               "API详情",
			"request":              "API请求描述",
			"response":             "API响应",
			"definitions":          "API数据结构",
			"request_content_type": "API内容类型",
			"can_run":              "是否可调试",
			"status":               "API状态",
			"description":          "API描述",
			"result_output_file":   "结果输出文件",
		},
	})
}

func dataSourceTencentCloudTsfApiDetail() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tsf api_detail",
		Read:        dataSourceTencentCloudTsfApiDetailRead,
		Schema: map[string]*schema.Schema{
			"microservice_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Microservice id.",
			},

			"path": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Api path.",
			},

			"method": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Request method.",
			},

			"pkg_version": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Pkg version.",
			},

			"application_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Application id.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Api detail.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"request": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Api request description.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param name.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type.",
									},
									"in": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param position.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param description.",
									},
									"required": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Require or not.",
									},
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default value.",
									},
								},
							},
						},
						"response": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Api response.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param description.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param type.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Param description.",
									},
								},
							},
						},
						"definitions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Api data struct.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name.",
									},
									"properties": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Object property list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Property name.",
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Property type.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Property description.",
												},
											},
										},
									},
								},
							},
						},
						"request_content_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Api content type.",
						},
						"can_run": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can debug or not.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "API status 0: offline 1: online, default 0. Note: This section may return null, indicating that no valid value can be obtained.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "API description. Note: This field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTsfApiDetailRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tsf_api_detail.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		microserviceId string
		path           string
		method         string
		pkgVersion     string
		applicationId  string
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("microservice_id"); ok {
		microserviceId = v.(string)
		paramMap["MicroserviceId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("path"); ok {
		path = v.(string)
		paramMap["Path"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("method"); ok {
		method = v.(string)
		paramMap["Method"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("pkg_version"); ok {
		pkgVersion = v.(string)
		paramMap["PkgVersion"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("application_id"); ok {
		applicationId = v.(string)
		paramMap["ApplicationId"] = helper.String(v.(string))
	}

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	var apiDetail *tsf.ApiDetailResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTsfApiDetailByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		apiDetail = result
		return nil
	})
	if err != nil {
		return err
	}

	apiDetailResponseMap := map[string]interface{}{}
	if apiDetail != nil {
		if apiDetail.Request != nil {
			requestList := []interface{}{}
			for _, request := range apiDetail.Request {
				requestMap := map[string]interface{}{}

				if request.Name != nil {
					requestMap["name"] = request.Name
				}

				if request.Type != nil {
					requestMap["type"] = request.Type
				}

				if request.In != nil {
					requestMap["in"] = request.In
				}

				if request.Description != nil {
					requestMap["description"] = request.Description
				}

				if request.Required != nil {
					requestMap["required"] = request.Required
				}

				if request.DefaultValue != nil {
					requestMap["default_value"] = request.DefaultValue
				}

				requestList = append(requestList, requestMap)
			}

			apiDetailResponseMap["request"] = requestList
		}

		if apiDetail.Response != nil {
			responseList := []interface{}{}
			for _, response := range apiDetail.Response {
				responseMap := map[string]interface{}{}

				if response.Name != nil {
					responseMap["name"] = response.Name
				}

				if response.Type != nil {
					responseMap["type"] = response.Type
				}

				if response.Description != nil {
					responseMap["description"] = response.Description
				}

				responseList = append(responseList, responseMap)
			}

			apiDetailResponseMap["response"] = responseList
		}

		if apiDetail.Definitions != nil {
			definitionsList := []interface{}{}
			for _, definitions := range apiDetail.Definitions {
				definitionsMap := map[string]interface{}{}

				if definitions.Name != nil {
					definitionsMap["name"] = definitions.Name
				}

				if definitions.Properties != nil {
					propertiesList := []interface{}{}
					for _, properties := range definitions.Properties {
						propertiesMap := map[string]interface{}{}

						if properties.Name != nil {
							propertiesMap["name"] = properties.Name
						}

						if properties.Type != nil {
							propertiesMap["type"] = properties.Type
						}

						if properties.Description != nil {
							propertiesMap["description"] = properties.Description
						}

						propertiesList = append(propertiesList, propertiesMap)
					}

					definitionsMap["properties"] = propertiesList
				}

				definitionsList = append(definitionsList, definitionsMap)
			}

			apiDetailResponseMap["definitions"] = definitionsList
		}

		if apiDetail.RequestContentType != nil {
			apiDetailResponseMap["request_content_type"] = apiDetail.RequestContentType
		}

		if apiDetail.CanRun != nil {
			apiDetailResponseMap["can_run"] = apiDetail.CanRun
		}

		if apiDetail.Status != nil {
			apiDetailResponseMap["status"] = apiDetail.Status
		}

		if apiDetail.Description != nil {
			apiDetailResponseMap["description"] = apiDetail.Description
		}

		_ = d.Set("result", []interface{}{apiDetailResponseMap})
	}

	d.SetId(strings.Join([]string{microserviceId, path, method, pkgVersion, applicationId}, FILED_SP))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), apiDetailResponseMap); e != nil {
			return e
		}
	}
	return nil
}
