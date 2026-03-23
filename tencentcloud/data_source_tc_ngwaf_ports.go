package tencentcloud

import (
	"context"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudNgwafPorts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafPortsRead,
		Schema: map[string]*schema.Schema{
			"edition": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Instance type, sparta-waf represents SAAS WAF, clb-waf represents CLB WAF.",
			},
			"instance_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Instance unique ID.",
			},
			"http_ports": {
				Computed:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Http port list for instance.",
			},
			"https_ports": {
				Computed:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Https port list for instance.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudNgwafPortsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_ports.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		rsp     *ngwaf.DescribePortsResponse
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("edition"); ok {
		paramMap["Edition"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["InstanceID"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafPortsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		rsp = result
		return nil
	})

	if err != nil {
		return err
	}

	if rsp != nil && rsp.Response != nil {
		if rsp.Response.HttpPorts != nil {
			_ = d.Set("http_ports", rsp.Response.HttpPorts)
		}

		if rsp.Response.HttpsPorts != nil {
			_ = d.Set("https_ports", rsp.Response.HttpsPorts)
		}

		if rsp.Response.RequestId != nil {
			d.SetId(*rsp.Response.RequestId)
		}
	}
	// 构造要写入文件的数据映射
	portsMap := map[string]interface{}{}
	if rsp != nil && rsp.Response != nil {
		if rsp.Response.HttpPorts != nil {
			portsMap["http_ports"] = rsp.Response.HttpPorts
		}
		if rsp.Response.HttpsPorts != nil {
			portsMap["https_ports"] = rsp.Response.HttpsPorts
		}
	}
	// 添加查询参数
	if v, ok := d.GetOk("edition"); ok {
		portsMap["edition"] = v.(string)
	}
	if v, ok := d.GetOk("instance_id"); ok {
		portsMap["instance_id"] = v.(string)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), portsMap); e != nil {
			return e
		}
	}

	return nil
}
