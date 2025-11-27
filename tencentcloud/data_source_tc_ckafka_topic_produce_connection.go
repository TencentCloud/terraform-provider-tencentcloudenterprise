// /*
// Use this data source to query detailed information of ckafka topic_produce_connection
//
// # Example Usage
//
// ```hcl
//
//	data "tencentcloudenterprise_ckafka_topic_produce_connection" "topic_produce_connection" {
//	  instance_id = "ckafka-xxxxxx"
//	  topic_name = "topic-xxxxxx"
//	}
//
// ```
// */
package tencentcloud

//
//import (
//	"context"
//
//	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	ckafka "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
//)
//
//func init() {
//	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_topic_produce_connection", CNDescription{
//		TerraformTypeCN: "ckafka 订阅生产者连接信息",
//		AttributesCN: map[string]string{
//			"instance_id":           "实例Id",
//			"topic_name":            "主题名称",
//			"result":                "链接信息返回结果集",
//			"ip_addr":               "ip地址",
//			"time":                  "连接时间",
//			"is_un_support_version": "是支持的版本",
//			"result_output_file":    "用于保存结果",
//		},
//	})
//}
//
//func dataSourceTencentCloudCkafkaTopicProduceConnection() *schema.Resource {
//	return &schema.Resource{
//		Description: "Use this data source to query detailed information of ckafka topic_produce_connection",
//		Read:        dataSourceTencentCloudCkafkaTopicProduceConnectionRead,
//		Schema: map[string]*schema.Schema{
//			"instance_id": {
//				Required:    true,
//				Type:        schema.TypeString,
//				Description: "InstanceId.",
//			},
//
//			"topic_name": {
//				Required:    true,
//				Type:        schema.TypeString,
//				Description: "TopicName.",
//			},
//
//			"result": {
//				Computed:    true,
//				Type:        schema.TypeList,
//				Description: "Link information return result set.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"ip_addr": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Ip address.",
//						},
//						"time": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Connect time.",
//						},
//						"is_un_support_version": {
//							Type:        schema.TypeBool,
//							Computed:    true,
//							Description: "Is the supported version.",
//						},
//					},
//				},
//			},
//
//			"result_output_file": {
//				Type:        schema.TypeString,
//				Optional:    true,
//				Description: "Used to save results.",
//			},
//		},
//	}
//}
//
//func dataSourceTencentCloudCkafkaTopicProduceConnectionRead(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("data_source.cloud_ckafka_topic_produce_connection.read")()
//	defer inconsistentCheck(d, meta)()
//
//	logId := getLogId(contextNil)
//
//	ctx := context.WithValue(context.TODO(), logIdKey, logId)
//	var (
//		instanceId string
//		topicName  string
//	)
//	paramMap := make(map[string]interface{})
//	if v, ok := d.GetOk("instance_id"); ok {
//		instanceId = v.(string)
//		paramMap["instance_id"] = helper.String(instanceId)
//	}
//
//	if v, ok := d.GetOk("topic_name"); ok {
//		topicName = v.(string)
//		paramMap["topic_name"] = helper.String(topicName)
//	}
//
//	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}
//
//	var result []*ckafka.DescribeConnectInfoResultDTO
//
//	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
//		topicProduceConnection, e := service.DescribeCkafkaTopicProduceConnectionByFilter(ctx, paramMap)
//		if e != nil {
//			return retryError(e)
//		}
//		result = topicProduceConnection
//		return nil
//	})
//	if err != nil {
//		return err
//	}
//
//	tmpList := make([]map[string]interface{}, 0)
//
//	if result != nil {
//		for _, describeConnectInfoResultDTO := range result {
//			describeConnectInfoResultDTOMap := map[string]interface{}{}
//
//			if describeConnectInfoResultDTO.IpAddr != nil {
//				describeConnectInfoResultDTOMap["ip_addr"] = describeConnectInfoResultDTO.IpAddr
//			}
//
//			if describeConnectInfoResultDTO.Time != nil {
//				describeConnectInfoResultDTOMap["time"] = describeConnectInfoResultDTO.Time
//			}
//
//			if describeConnectInfoResultDTO.IsUnSupportVersion != nil {
//				describeConnectInfoResultDTOMap["is_un_support_version"] = describeConnectInfoResultDTO.IsUnSupportVersion
//			}
//
//			tmpList = append(tmpList, describeConnectInfoResultDTOMap)
//		}
//
//		_ = d.Set("result", tmpList)
//	}
//
//	d.SetId(instanceId + FILED_SP + topicName)
//
//	output, ok := d.GetOk("result_output_file")
//	if ok && output.(string) != "" {
//		if e := writeToFile(output.(string), tmpList); e != nil {
//			return e
//		}
//	}
//	return nil
//}
