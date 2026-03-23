package tencentcloud

import (
	"context"
	"strconv"
	"time"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudNgwafPeakPoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudNgwafPeakPointsRead,
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
				Description: "The domain name to be queried. If all domain name data is queried, this parameter is not filled in.",
			},
			"edition": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Only support sparta-waf and clb-waf. If not passed, there will be no filtering.",
			},
			"instance_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "WAF instance ID, if not passed, there will be no filtering.",
			},
			"metric_name": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(PEAK_POINTS_METRIC_NAMES),
				Description:  "Three values are available: `access`-Peak qps trend chart; `cc`-Trend chart of total number of CC attacks; `attack`-Trend chart of total number of web attacks.",
			},
			"points": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "point list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Second level timestamp.",
						},
						"access": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "qps.",
						},
						"bot_access": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Bot qps.",
						},
						"attack": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of web attacks.",
						},
						"cc": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of cc attacks.",
						},
						"down": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak downlink bandwidth, unit B.",
						},
						"up": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak uplink bandwidth, unit B.",
						},
						"status_server_error": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned by WAF to the server.",
						},
						"status_client_error": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned by WAF to the client.",
						},
						"status_redirect": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned by WAF to the client.",
						},
						"status_ok": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned by WAF to the client.",
						},
						"upstream_server_error": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned to WAF by the origin site.",
						},
						"upstream_client_error": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned to WAF by the origin site.",
						},
						"upstream_redirect": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trend chart of the number of status codes returned to WAF by the origin site.",
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

func dataSourceTencentCloudNgwafPeakPointsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ngwaf_peak_points.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		points  []*ngwaf.PeakPointsItem
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("from_time"); ok {
		paramMap["FromTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("to_time"); ok {
		paramMap["ToTime"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("domain"); ok {
		paramMap["Domain"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		paramMap["Edition"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["InstanceID"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("metric_name"); ok {
		paramMap["MetricName"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeWafPeakPointsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		points = result
		return nil
	})

	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(points))

	if points != nil {
		for _, point := range points {
			dMap := map[string]interface{}{}

			if point.Time != nil {
				dMap["time"] = point.Time
			}

			if point.Access != nil {
				dMap["access"] = point.Access
			}

			if point.BotAccess != nil {
				dMap["bot_access"] = point.BotAccess
			}

			if point.Attack != nil {
				dMap["attack"] = point.Attack
			}

			if point.Cc != nil {
				dMap["cc"] = point.Cc
			}

			if point.Down != nil {
				dMap["down"] = point.Down
			}

			if point.Up != nil {
				dMap["up"] = point.Up
			}

			if point.StatusServerError != nil {
				dMap["status_server_error"] = point.StatusServerError
			}

			if point.StatusClientError != nil {
				dMap["status_client_error"] = point.StatusClientError
			}

			if point.StatusRedirect != nil {
				dMap["status_redirect"] = point.StatusRedirect
			}

			if point.StatusOk != nil {
				dMap["status_ok"] = point.StatusOk
			}

			if point.UpstreamServerError != nil {
				dMap["upstream_server_error"] = point.UpstreamServerError
			}

			if point.UpstreamClientError != nil {
				dMap["upstream_client_error"] = point.UpstreamClientError
			}

			if point.UpstreamRedirect != nil {
				dMap["upstream_redirect"] = point.UpstreamRedirect
			}

			tmpList = append(tmpList, dMap)
		}

		_ = d.Set("points", tmpList)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}