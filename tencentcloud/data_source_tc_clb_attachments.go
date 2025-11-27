/*
Use this data source to query detailed information of CLB attachments

Example Usage
```hcl

	data "tencentcloudenterprise_clb_attachments" "clblab" {
	  listener_id = "lbl-hh141sn9"
	  clb_id      = "lb-k2zjp9lv"
	  rule_id     = "loc-4xxr2cy7"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_clb_attachments", CNDescription{
		TerraformTypeCN: "负载均衡接入",
		DescriptionCN:   "提供负载均衡接入数据源，用于查询CLB接入的详细信息。",
		AttributesCN: map[string]string{
			"clb_id":             "要查询的CLB的ID",
			"listener_id":        "要查询的CLB侦听器的ID",
			"rule_id":            "CLB侦听器规则的ID如果侦听器的协议是“HTTP”/“HTTPS”，则需要此段落",
			"result_output_file": "用于保存结果",
			"attachment_list":    "云负载平衡器接入配置的列表每个元素都包含以下属性：",
			"protocol_type":      "侦听器中的协议类型，可用值包括“TCP”、“UDP”、“HTTP”、“HTTPS”和“TCP_SSL”注意：“TCP_SSL”正在内部测试，如果需要使用，请申请",
			"targets":            "要附加的后端信息",
			"instance_id":        "后端服务器的Id",
			"port":               "后端服务器的端口",
			"weight":             "后端服务的转发权重，范围为[0100]，默认为“10”",
		},
	})
}

func dataSourceTencentCloudClbServerAttachments() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud CLB Server Attachments",
		Read:        dataSourceTencentCloudClbServerAttachmentsRead,

		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the CLB to be queried.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the CLB listener to be queried.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the CLB listener rule. If the protocol of listener is `HTTP`/`HTTPS`, this para is required.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"attachment_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of cloud load balancer attachment configurations. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"clb_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CLB.",
						},
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CLB listener.",
						},
						"protocol_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of protocol within the listener, and available values include `TCP`, `UDP`, `HTTP`, `HTTPS` and `TCP_SSL`. NOTES: `TCP_SSL` is testing internally, please apply if you need to use.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CLB listener rule.",
						},
						"targets": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Information of the backends to be attached.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of the backend server.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port of the backend server.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Forwarding weight of the backend service, the range of [0, 100], defaults to `10`.",
									},
									"eni_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ENI IP of the backend server.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudClbServerAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_attachments.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]string)
	clbId := d.Get("clb_id").(string)
	listenerId := d.Get("listener_id").(string)
	checkErr := ListenerIdCheck(listenerId)
	if checkErr != nil {
		return checkErr
	}
	locationId := ""
	if v, ok := d.GetOk("rule_id"); ok {
		locationId = v.(string)
		checkErr := RuleIdCheck(locationId)
		if checkErr != nil {
			return checkErr
		}
		params["rule_id"] = locationId
	}
	params["clb_id"] = clbId
	params["listener_id"] = listenerId

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var attachments []*clb.ListenerBackend
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := clbService.DescribeAttachmentsByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		attachments = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB attachments failed, reason:%+v", logId, err)
		return err
	}
	attachmentList := make([]map[string]interface{}, 0, len(attachments))
	ids := make([]string, 0, len(attachments))
	for _, attachment := range attachments {
		mapping := map[string]interface{}{
			"clb_id":        clbId,
			"listener_id":   listenerId,
			"protocol_type": attachment.Protocol,
		}
		if locationId != "" {
			mapping["rule_id"] = locationId
		}
		if *attachment.Protocol == CLB_LISTENER_PROTOCOL_HTTP || *attachment.Protocol == CLB_LISTENER_PROTOCOL_HTTPS {
			if len(attachment.Rules) > 0 {
				for _, loc := range attachment.Rules {
					if locationId == "" || locationId == *loc.LocationId {
						mapping["targets"] = flattenBackendList(loc.Targets)
					}
				}
			}
		} else if *attachment.Protocol == CLB_LISTENER_PROTOCOL_TCP || *attachment.Protocol == CLB_LISTENER_PROTOCOL_UDP {
			mapping["targets"] = flattenBackendList(attachment.Targets)
		}
		attachmentList = append(attachmentList, mapping)
		ids = append(ids, locationId+"#"+listenerId+"#"+clbId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("attachment_list", attachmentList); e != nil {
		log.Printf("[CRITAL]%s provider set attachment list fail, reason:%+v", logId, e)
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), attachmentList); e != nil {
			return e
		}
	}

	return nil
}
