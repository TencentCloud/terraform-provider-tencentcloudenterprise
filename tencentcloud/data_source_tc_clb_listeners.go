/*
Use this data source to query detailed information of CLB listener

Example Usage
```hcl

	data "tencentcloudenterprise_clb_listeners" "foo" {
	  clb_id      = "lb-k2zjp9lv"
	  listener_id = "lbl-mwr6vbtv"
	  protocol    = "TCP"
	  port        = 80
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
	registerDataDescriptionProvider("tencentcloudenterprise_clb_listeners", CNDescription{
		TerraformTypeCN: "负载均衡监听器",
		DescriptionCN:   "提供负载均衡监听器数据源，用于查询CLB监听器的详细信息。",
		AttributesCN: map[string]string{
			"clb_id":                     "要查询的CLB的Id",
			"listener_id":                "要查询的侦听器的Id",
			"protocol":                   "侦听器中的协议类型，可用值为“TCP”、“UDP”、“HTTP”、“HTTPS”",
			"port":                       "CLB侦听器的端口",
			"result_output_file":         "用于保存结果",
			"listener_list":              "云负载平衡器的侦听器列表每个元素都包含以下属性：",
			"listener_name":              "CLB侦听器的名称",
			"health_check_switch":        "指示是否启用运行状况检查",
			"health_check_time_out":      "运行状况检查的响应超时值范围为2-60秒，默认值为“2”秒响应超时需要小于检查间隔注意：TCP/UDP侦听器允许直接配置",
			"health_check_interval_time": "健康检查的间隔时间取值范围为2-300秒，默认为“5”秒注意：TCP/UDP侦听器允许直接配置，HTTP/HTTPS侦听器需要在cloud_clb_listener_rule中配置",
			"health_check_health_num":    "健康检查的健康阈值，默认为“3”如果连续三次返回健康检查的成功结果，则CVM被识别为健康数值范围为2-10注意：TCP/UDP侦听器允许直接配置，HTTP/HTTPS侦听器需要在cloud_clb_listener_rule中配置",
			"health_check_unhealth_num":  "健康检查的不健康阈值，默认值为“3”如果连续三次返回健康检查的成功结果，则CVM被识别为不健康数值范围为2-10注意：TCP/UDP侦听器允许直接配置，HTTP/HTTPS侦听器需要在cloud_clb_listener_rule中配置",
			"health_check_type":          "用于健康检查的协议",
			"health_check_port":          "运行状况检查端口是后端服务的端口",
			"health_check_http_version":  "后端服务的HTTP版本",
			"health_check_http_code":     "TCP侦听器的HTTP健康检查代码",
			"health_check_http_path":     "TCP侦听器的HTTP运行状况检查路径",
			"health_check_http_domain":   "TCP侦听器的HTTP运行状况检查域",
			"health_check_http_method":   "TCP侦听器的HTTP健康检查方法",
			"health_check_context_type":  "健康检查协议",
			"health_check_send_context":  "它表示运行状况检查发送的请求的内容",
			"health_check_recv_context":  "它表示运行状况检查返回的结果",
			"certificate_ssl_mode":       "证书类型，可用值包括“UNIDIRECTIONAL”、“MUTUAL”注意：仅支持“HTTPS”协议的侦听器，并且必须在可用时进行设置",
			"certificate_id":             "服务器证书的Id当协议为“HTTPS”时，必须设置它注意：仅由“HTTPS”协议的侦听器支持，并且必须在可用时进行设置",
			"certificate_ca_id":          "客户端证书的Id当SSLMode为“相互”时，必须设置它注意：仅由“HTTPS”协议的侦听器支持",
			"session_expire_time":        "CLB侦听器中的会话持续时间注意：TCP/UDP侦听器允许直接配置，HTTP/HTTPS侦听器需要在cloud_clb_listener_rule中配置",
			"scheduler":                  "CLB侦听器的调度方法，可用值为“WRR”和“LEAST_CONN”默认值为“WRR”注意：“HTTP”和“HTTPS”协议的侦听器还支持“IP HASH”方法注意：TCP/UDP侦听器允许直接配置，HTTP/HTTPS侦听器需要在cloud_clb_listener_rule中配置",
			"sni_switch":                 "指示是否启用了SNI注意：仅受“HTTPS”协议支持",
		},
	})
}

func dataSourceTencentCloudClbListeners() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the CLB listeners of the specified CLB instance.",
		Read:        dataSourceTencentCloudClbListenersRead,

		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the CLB to be queried.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the listener to be queried.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CLB_LISTENER_PROTOCOL),
				Description:  "Type of protocol within the listener, and available values are `TCP`, `UDP`, `HTTP`, `HTTPS`.",
			},
			"port": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerInRange(1, 65535),
				Description:  "Port of the CLB listener.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"listener_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of listeners of cloud load balancers. Each element contains the following attributes:",
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
							Description: "ID of the listener.",
						},
						"listener_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the CLB listener.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol of the listener. Available values are `HTTP`, `HTTPS`, `TCP`, `UDP`.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port of the CLB listener.",
						},
						"health_check_switch": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether health check is enabled.",
						},
						"health_check_time_out": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Response timeout of health check. The value range is 2-60 sec, and the default is `2` sec. Response timeout needs to be less than check interval. NOTES: TCP/UDP listener allows direct configuration.",
						},
						"health_check_interval_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interval time of health check. The value range is 2-300 sec, and the default is `5` sec. NOTES: TCP/UDP listener allows direct configuration, HTTP/HTTPS listener needs to be configured in cloud_clb_listener_rule.",
						},
						"health_check_health_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Health threshold of health check, and the default is `3`. If a success result is returned for the health check three consecutive times, the CVM is identified as healthy. The value range is 2-10. NOTES: TCP/UDP listener allows direct configuration, HTTP/HTTPS listener needs to be configured in cloud_clb_listener_rule.",
						},
						"health_check_unhealth_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Unhealthy threshold of health check, and the default is `3`. If a success result is returned for the health check three consecutive times, the CVM is identified as unhealthy. The value range is 2-10. NOTES: TCP/UDP listener allows direct configuration, HTTP/HTTPS listener needs to be configured in cloud_clb_listener_rule.",
						},
						"health_check_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol used for health check.",
						},
						"health_check_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The health check port is the port of the backend service.",
						},
						"health_check_http_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP version of the backend service.",
						},
						"health_check_http_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "HTTP health check code of TCP listener.",
						},
						"health_check_http_path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP health check path of TCP listener.",
						},
						"health_check_http_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP health check domain of TCP listener.",
						},
						"health_check_http_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP health check method of TCP listener.",
						},
						"health_check_context_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Health check protocol.",
						},
						"health_check_send_context": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "It represents the content of the request sent by the health check.",
						},
						"health_check_recv_context": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "It represents the result returned by the health check.",
						},
						"certificate_ssl_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of certificate, and available values inclue `UNIDIRECTIONAL`, `MUTUAL`. NOTES: Only supports listeners of `HTTPS` protocol and must be set when it is available.",
						},
						"certificate_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of the server certificate. It must be set when protocol is `HTTPS`. NOTES: only supported by listeners of `HTTPS` protocol and must be set when it is available.",
						},
						"certificate_ca_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of the client certificate. It must be set when SSLMode is `mutual`. NOTES: only supported by listeners of `HTTPS` protocol.",
						},
						"session_expire_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Time of session persistence within the CLB listener. NOTES: TCP/UDP listener allows direct configuration, HTTP/HTTPS listener needs to be configured in cloud_clb_listener_rule.",
						},
						"scheduler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Scheduling method of the CLB listener, and available values are `WRR` and `LEAST_CONN`. The default is `WRR`. NOTES: The listener of 'HTTP' and `HTTPS` protocol additionally supports the `IP HASH` method. NOTES: TCP/UDP listener allows direct configuration, HTTP/HTTPS listener needs to be configured in cloud_clb_listener_rule.",
						},
						"sni_switch": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether SNI is enabled. NOTES: Only supported by `HTTPS` protocol.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudClbListenersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_clb_listeners.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clbId := d.Get("clb_id").(string)

	params := make(map[string]interface{})
	params["clb_id"] = clbId
	if v, ok := d.GetOk("listener_id"); ok {
		listenerId := v.(string)
		params["listener_id"] = listenerId
		checkErr := ListenerIdCheck(listenerId)
		if checkErr != nil {
			return checkErr
		}
	}
	if v, ok := d.GetOk("port"); ok {
		params["port"] = v.(int)
	}
	if v, ok := d.GetOk("protocol"); ok {
		params["protocol"] = v.(string)
	}

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var listeners []*clb.Listener
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := clbService.DescribeListenersByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		listeners = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB listeners failed, reason:%+v", logId, err)
		return err
	}
	listenerList := make([]map[string]interface{}, 0, len(listeners))
	ids := make([]string, 0, len(listeners))
	for _, listener := range listeners {
		mapping := map[string]interface{}{
			"clb_id":        clbId,
			"listener_id":   listener.ListenerId,
			"listener_name": listener.ListenerName,
			"protocol":      listener.Protocol,
			"port":          listener.Port,
		}
		if listener.SessionExpireTime != nil {
			mapping["session_expire_time"] = listener.SessionExpireTime
		}
		if listener.SniSwitch != nil {
			sniSwitch := false
			if *listener.SniSwitch == int64(1) {
				sniSwitch = true
			}
			mapping["sni_switch"] = sniSwitch
		}
		mapping["scheduler"] = listener.Scheduler
		if listener.HealthCheck != nil {
			health_check_switch := false
			if *listener.HealthCheck.HealthSwitch == int64(1) {
				health_check_switch = true
			}
			mapping["health_check_switch"] = health_check_switch
			mapping["health_check_time_out"] = listener.HealthCheck.TimeOut
			mapping["health_check_interval_time"] = listener.HealthCheck.IntervalTime
			mapping["health_check_health_num"] = listener.HealthCheck.HealthNum
			mapping["health_check_unhealth_num"] = listener.HealthCheck.UnHealthNum
			mapping["health_check_http_code"] = listener.HealthCheck.HttpCode
			mapping["health_check_http_path"] = listener.HealthCheck.HttpCheckPath
			mapping["health_check_http_domain"] = listener.HealthCheck.HttpCheckDomain
			mapping["health_check_http_method"] = listener.HealthCheck.HttpCheckMethod
			mapping["health_check_http_version"] = listener.HealthCheck.HttpVersion
			mapping["health_check_context_type"] = listener.HealthCheck.ContextType
			mapping["health_check_send_context"] = listener.HealthCheck.SendContext
			mapping["health_check_recv_context"] = listener.HealthCheck.RecvContext
			mapping["health_check_type"] = listener.HealthCheck.CheckType
			mapping["health_check_port"] = listener.HealthCheck.CheckPort
		}
		if listener.Certificate != nil {
			mapping["certificate_ssl_mode"] = listener.Certificate.SSLMode
			mapping["certificate_id"] = listener.Certificate.CertId
			if listener.Certificate.CertCaId != nil {
				mapping["certificate_ca_id"] = listener.Certificate.CertCaId
			}
		}
		listenerList = append(listenerList, mapping)
		ids = append(ids, *listener.ListenerId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("listener_list", listenerList); e != nil {
		log.Printf("[CRITAL]%s provider set CLB listener list fail, reason:%+v", logId, e)
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), listenerList); e != nil {
			return e
		}
	}

	return nil
}
