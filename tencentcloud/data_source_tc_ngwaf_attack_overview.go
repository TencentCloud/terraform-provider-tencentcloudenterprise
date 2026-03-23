/*
Use this data source to query NGWAF attack overview statistics for a specified time period.

# Example Usage

```hcl

	data "tencentcloudenterprise_ngwaf_attack_overview" "example" {
	  from_time   = "2023-01-01 00:00:00"
	  to_time     = "2023-01-31 23:59:59"
	  domain      = "example.com"
	  instance_id = "waf-xxxxxxxx"
	}

```
*/
package tencentcloud

import (
	"context"
	"strconv"
	"time"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudNgwafAttackOverview() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafAttackOverviewRead,
		Schema: map[string]*schema.Schema{
			"from_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Begin time.",
			},
			"to_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "End time.",
			},
			"domain": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"instance_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Waf instanceId, otherwise not filter.",
			},
			"access_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access count.",
			},
			"attack_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Attack count.",
			},
			"acl_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access control count.",
			},
			"cc_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "CC attack count.",
			},
			"bot_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Bot attack count.",
			},
			"api_assets_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Api asset count.",
			},
			"api_risk_event_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of API risk events.",
			},
			"applet_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Applet attack count.",
			},
			"ip_black_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "IP blacklist count.",
			},
			"leak_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Information leak count.",
			},
			"tamper_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Tamper proof count.",
			},
			"acl_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access control circle count.",
			},
			"access_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Access circle count.",
			},
			"api_assets_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "API assets circle count.",
			},
			"api_risk_event_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "API risk event circle count.",
			},
			"applet_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Applet circle count.",
			},
			"attack_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Attack circle count.",
			},
			"bot_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Bot circle count.",
			},
			"cc_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "CC circle count.",
			},
			"ip_black_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "IP blacklist circle count.",
			},
			"leak_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Leak circle count.",
			},
			"tamper_circle_count": {
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Tamper circle count.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudNgwafAttackOverviewRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_attack_overview.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = NgwafService{client: meta.(*TencentCloudClient)}
		param   = make(map[string]interface{})
	)

	if v, ok := d.GetOk("from_time"); ok {
		param["FromTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("to_time"); ok {
		param["ToTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("domain"); ok {
		param["Domain"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_id"); ok {
		param["InstanceID"] = helper.String(v.(string))
	}

	attackOverview, err := service.DescribeNgwafAttackOverviewByFilter(ctx, param)
	if err != nil {
		return err
	}

	if attackOverview == nil || attackOverview.Response == nil {
		return nil
	}

	if attackOverview.Response.AccessCount != nil {
		_ = d.Set("access_count", attackOverview.Response.AccessCount)
	}

	if attackOverview.Response.AttackCount != nil {
		_ = d.Set("attack_count", attackOverview.Response.AttackCount)
	}

	if attackOverview.Response.ACLCount != nil {
		_ = d.Set("acl_count", attackOverview.Response.ACLCount)
	}

	if attackOverview.Response.CCCount != nil {
		_ = d.Set("cc_count", attackOverview.Response.CCCount)
	}

	if attackOverview.Response.BotCount != nil {
		_ = d.Set("bot_count", attackOverview.Response.BotCount)
	}

	if attackOverview.Response.ApiAssetsCount != nil {
		_ = d.Set("api_assets_count", attackOverview.Response.ApiAssetsCount)
	}

	if attackOverview.Response.ApiRiskEventCount != nil {
		_ = d.Set("api_risk_event_count", attackOverview.Response.ApiRiskEventCount)
	}

	if attackOverview.Response.AppletCount != nil {
		_ = d.Set("applet_count", attackOverview.Response.AppletCount)
	}

	if attackOverview.Response.IPBlackCount != nil {
		_ = d.Set("ip_black_count", attackOverview.Response.IPBlackCount)
	}

	if attackOverview.Response.LeakCount != nil {
		_ = d.Set("leak_count", attackOverview.Response.LeakCount)
	}

	if attackOverview.Response.TamperCount != nil {
		_ = d.Set("tamper_count", attackOverview.Response.TamperCount)
	}

	if attackOverview.Response.ACLCircleCount != nil {
		_ = d.Set("acl_circle_count", attackOverview.Response.ACLCircleCount)
	}

	if attackOverview.Response.AccessCircleCount != nil {
		_ = d.Set("access_circle_count", attackOverview.Response.AccessCircleCount)
	}

	if attackOverview.Response.ApiAssetsCircleCount != nil {
		_ = d.Set("api_assets_circle_count", attackOverview.Response.ApiAssetsCircleCount)
	}

	if attackOverview.Response.ApiRiskEventCircleCount != nil {
		_ = d.Set("api_risk_event_circle_count", attackOverview.Response.ApiRiskEventCircleCount)
	}

	if attackOverview.Response.AppletCircleCount != nil {
		_ = d.Set("applet_circle_count", attackOverview.Response.AppletCircleCount)
	}

	if attackOverview.Response.AttackCircleCount != nil {
		_ = d.Set("attack_circle_count", attackOverview.Response.AttackCircleCount)
	}

	if attackOverview.Response.BotCircleCount != nil {
		_ = d.Set("bot_circle_count", attackOverview.Response.BotCircleCount)
	}

	if attackOverview.Response.CCCircleCount != nil {
		_ = d.Set("cc_circle_count", attackOverview.Response.CCCircleCount)
	}

	if attackOverview.Response.IPBlackCircleCount != nil {
		_ = d.Set("ip_black_circle_count", attackOverview.Response.IPBlackCircleCount)
	}

	if attackOverview.Response.LeakCircleCount != nil {
		_ = d.Set("leak_circle_count", attackOverview.Response.LeakCircleCount)
	}

	if attackOverview.Response.TamperCircleCount != nil {
		_ = d.Set("tamper_circle_count", attackOverview.Response.TamperCircleCount)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		attackOverviewMap := map[string]interface{}{
			"access_count":                d.Get("access_count"),
			"attack_count":                d.Get("attack_count"),
			"acl_count":                   d.Get("acl_count"),
			"cc_count":                    d.Get("cc_count"),
			"bot_count":                   d.Get("bot_count"),
			"api_assets_count":            d.Get("api_assets_count"),
			"api_risk_event_count":        d.Get("api_risk_event_count"),
			"applet_count":                d.Get("applet_count"),
			"ip_black_count":              d.Get("ip_black_count"),
			"leak_count":                  d.Get("leak_count"),
			"tamper_count":                d.Get("tamper_count"),
			"acl_circle_count":            d.Get("acl_circle_count"),
			"access_circle_count":         d.Get("access_circle_count"),
			"api_assets_circle_count":     d.Get("api_assets_circle_count"),
			"api_risk_event_circle_count": d.Get("api_risk_event_circle_count"),
			"applet_circle_count":         d.Get("applet_circle_count"),
			"attack_circle_count":         d.Get("attack_circle_count"),
			"bot_circle_count":            d.Get("bot_circle_count"),
			"cc_circle_count":             d.Get("cc_circle_count"),
			"ip_black_circle_count":       d.Get("ip_black_circle_count"),
			"leak_circle_count":           d.Get("leak_circle_count"),
			"tamper_circle_count":         d.Get("tamper_circle_count"),
			"from_time":                   d.Get("from_time"),
			"to_time":                     d.Get("to_time"),
			"domain":                      d.Get("domain"),
			"instance_id":                 d.Get("instance_id"),
		}
		if e := writeToFile(output.(string), attackOverviewMap); e != nil {
			return e
		}
	}

	return nil
}
