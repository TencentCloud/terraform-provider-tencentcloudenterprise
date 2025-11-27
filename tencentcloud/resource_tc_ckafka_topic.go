/*
Use this resource to create ckafka topic.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ckafka_topic" "foo" {
		instance_id                     = "ckafka-f9ife4zz"
		topic_name                      = "example"
		note                            = "topic note"
		replica_num                     = 2
		partition_num                   = 1
		enable_white_list               = true
		ip_white_list                   = ["ip1","ip2"]
		clean_up_policy                 = "delete"
		sync_replica_min_num            = 1
		unclean_leader_election_enable  = false
		segment                         = 3600000
		retention                       = 60000
		max_message_bytes               = 0
	}

```

# Import

ckafka topic can be imported using the instance_id#topic_name, e.g.

```
$ terraform import cloud_ckafka_topic.foo ckafka-f9ife4zz#example
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ckafka_topic", CNDescription{
		TerraformTypeCN: "ckafka topic",
		DescriptionCN:   "提供CKafka主题资源，用于创建和管理CKafka主题。",
		AttributesCN: map[string]string{
			"instance_id":                    "实例id",
			"topic_name":                     "主题名称",
			"partition_num":                  "分区数",
			"replica_num":                    "副本个数，不能多于 broker 数，最大为3",
			"enable_white_list":              "是否开启ip鉴权白名单，true表示开启，false表示不开启",
			"ip_white_list":                  "Ip白名单列表，配额限制,enable_white_list=true时必选",
			"note":                           "备注",
			"retention":                      "可选参数。消息保留时间，单位ms，当前最小值为60000ms",
			"sync_replica_min_num":           "最小同步复制数",
			"clean_up_policy":                "清理日志策略，日志清理模式，默认为\"delete\"。\"delete\"：日志按保存时间删除，\"compact\"：日志按 key 压缩，\"compact, delete\"：日志按 key 压缩且会按保存时间删除。",
			"unclean_leader_election_enable": "0表示 false。 1表示 true。",
			"max_message_bytes":              "主题消息最大值，单位为 Byte，最小值1024Byte(即1KB)，最大值为12582912Byte（即12MB）",
			"segment":                        "Segment分片滚动的时长，单位ms，当前最小为3600000ms",
			"create_time":                    "创建时间",
			"forward_interval":               "数据备份到cos的周期频率",
			"forward_cos_bucket":             "数据备份cos bucket: 转存到cos 的bucket地址",
			"forward_status":                 "数据备份cos 状态： 1 不开启数据备份，0 开启数据备份",
			"segment_bytes":                  "Segment 分片滚动的字节数",
		},
	})
}

func resourceTencentCloudCkafkaTopic() *schema.Resource {
	return &schema.Resource{
		Description: "Use this resource to create ckafka topic.",
		Create:      resourceTencentCloudCkafkaTopicCreate,
		Read:        resourceTencentCloudCkafkaTopicRead,
		Update:      resourceTencentCloudCkafkaTopicUpdate,
		Delete:      resourceTencentCLoudCkafkaTopicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Ckafka instance ID.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-).",
			},
			"partition_num": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The number of partition.",
			},
			"replica_num": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Number of replicas. The value cannot exceed the number of brokers. Maximum value: 3.",
			},
			"enable_white_list": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to open the ip whitelist, `true`: open, `false`: close.",
			},
			"ip_white_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "IP allowlist used for specifying a quota. You are required to specify this parameter when EnableWhiteList is set to true.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"note": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The subject note. It must start with a letter, and the remaining part can contain letters, numbers and dashes (-).",
			},
			"retention": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60000,
				ValidateFunc: validateIntegerMin(60000),
				Description:  "Message retention period. Unit: ms. Minimum value: 60000.",
			},
			"sync_replica_min_num": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Min number of sync replicas, Default is `1`.",
			},
			"clean_up_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "delete",
				Description: "Policy or mode to clear logs. Default value: delete. delete\": logs are deleted based on the corresponding retention period. \"compact\": logs are compressed based on the corresponding keys. \"compact, delete\": logs are compressed based on the corresponding keys and deleted based on the corresponding retention period.",
			},
			"unclean_leader_election_enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.",
			},
			"max_message_bytes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of messages in a topic. Unit: Byte. Maximum value: 12582912. The maximum value is equal to 12 MB.",
			},
			"segment": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validateIntegerMin(3600000),
				Description:  "Period after which a segment is rolled. Unit: ms. Minimum value: 3600000.",
			},
			//"message_storage_location": {
			//	Type:        schema.TypeString,
			//	Computed:    true,
			//	Description: "Message storage location.",
			//},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the CKafka topic.",
			},
			"forward_interval": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Periodic frequency of data backup to cos.",
			},
			"forward_cos_bucket": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data backup cos bucket: the bucket address that is dumped to cos.",
			},
			"forward_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Data backup cos status. Valid values: `0`, `1`. `1`: do not open data backup, `0`: open data backup.",
			},
			"segment_bytes": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of bytes rolled by shard.",
			},
		},
	}
}

func resourceTencentCloudCkafkaTopicCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_topic.create")()
	var (
		logId         = getLogId(contextNil)
		ctx           = context.WithValue(context.TODO(), logIdKey, logId)
		ckafkcService = CkafkaService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		request               = ckafka.NewCreateTopicRequest()
		instanceId            = d.Get("instance_id").(string)
		topicName             = d.Get("topic_name").(string)
		note                  string
		ipWhiteLists          = d.Get("ip_white_list").([]interface{})
		ipWhiteList           = make([]*string, 0, len(ipWhiteLists))
		syncReplicaMinNum     = int64(d.Get("sync_replica_min_num").(int))
		uncleanLeaderElection = d.Get("unclean_leader_election_enable").(bool)
		cleanUpPolicy         = d.Get("clean_up_policy").(string)
		whiteListSwitch       = d.Get("enable_white_list").(bool)
		retention             = d.Get("retention").(int)
	)
	for _, value := range ipWhiteLists {
		ipWhiteList = append(ipWhiteList, helper.String(value.(string)))
	}
	request.EnableWhiteList = helper.BoolToInt64Ptr(whiteListSwitch)
	if whiteListSwitch {
		if len(ipWhiteList) == 0 {
			return fmt.Errorf("this Topic %s Create Failed, reason: ip whitelist switch is on, ip whitelist cannot be empty", topicName)
		}
		request.IpWhiteList = ipWhiteList
	}
	if v, ok := d.GetOk("note"); ok {
		note = v.(string)
		request.Note = &note
	}
	if v, ok := d.GetOk("segment"); ok {
		if v.(int) != 0 {
			request.SegmentMs = helper.IntInt64(v.(int))
		}
	}
	if v, ok := d.GetOk("max_message_bytes"); ok {
		request.MaxMessageBytes = helper.IntInt64(v.(int))
	}
	request.InstanceId = &instanceId
	request.TopicName = &topicName
	request.PartitionNum = helper.IntInt64(d.Get("partition_num").(int))
	request.ReplicaNum = helper.IntInt64(d.Get("replica_num").(int))
	request.MinInsyncReplicas = &syncReplicaMinNum
	request.UncleanLeaderElectionEnable = helper.BoolToInt64Ptr(uncleanLeaderElection)
	request.CleanUpPolicy = &cleanUpPolicy
	request.RetentionMs = helper.IntInt64(retention)
	//Before create topic,Check if kafka exists
	_, has, err := ckafkcService.DescribeInstanceById(ctx, instanceId)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("ckafka %s does not exist", instanceId)
	}
	errCreate := ckafkcService.CreateCkafkaTopic(ctx, request)
	if errCreate != nil {
		return errCreate
	}
	_, hasExist, err := ckafkcService.DescribeCkafkaTopicByName(ctx, instanceId, topicName)
	if err != nil {
		return err
	}
	if !hasExist {
		return fmt.Errorf("this Topic %s Create Failed", topicName)
	}
	if len(ipWhiteList) > 0 && whiteListSwitch {
		err = ckafkcService.AddCkafkaTopicIpWhiteList(ctx, instanceId, topicName, ipWhiteList)
		if err != nil {
			return err
		}
	}
	resourceId := instanceId + FILED_SP + topicName
	d.SetId(resourceId)
	return resourceTencentCloudCkafkaTopicRead(d, meta)
}

func resourceTencentCloudCkafkaTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_topic.read")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ckafkcService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	items := strings.Split(d.Id(), FILED_SP)
	if len(items) < 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := items[0]
	topicName := items[1]
	topicListInfo, hasExist, e := ckafkcService.DescribeCkafkaTopicByName(ctx, instanceId, topicName)
	if e != nil {
		return e
	}
	if !hasExist {
		d.SetId("")
		return nil
	}

	var topicinfo *ckafka.TopicAttributesResponse
	errInfo := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		topicDetail, e := ckafkcService.DescribeCkafkaTopicAttributes(ctx, instanceId, topicName)
		if e != nil {
			return retryError(e)
		}
		if topicDetail == nil {
			d.SetId("")
			return nil
		}
		topicinfo = topicDetail
		return nil
	})
	if errInfo != nil {
		return fmt.Errorf("[API]Describe kafka topic fail,reason:%s", errInfo.Error())
	}
	_ = d.Set("instance_id", instanceId)
	_ = d.Set("note", topicinfo.Note)
	_ = d.Set("ip_white_list", topicinfo.IpWhiteList)
	_ = d.Set("enable_white_list", *topicinfo.EnableWhiteList == 1)
	_ = d.Set("replica_num", topicListInfo.ReplicaNum)
	_ = d.Set("create_time", helper.FormatUnixTime(uint64(*topicinfo.CreateTime)))
	_ = d.Set("partition_num", topicinfo.PartitionNum)
	_ = d.Set("topic_name", topicListInfo.TopicName)
	_ = d.Set("forward_interval", topicListInfo.ForwardInterval)
	_ = d.Set("forward_cos_bucket", topicListInfo.ForwardCosBucket)
	_ = d.Set("forward_status", topicListInfo.ForwardStatus)
	_ = d.Set("clean_up_policy", topicinfo.Config.CleanUpPolicy)
	_ = d.Set("max_message_bytes", topicinfo.Config.MaxMessageBytes)
	_ = d.Set("sync_replica_min_num", topicinfo.Config.MinInsyncReplicas)
	_ = d.Set("retention", topicinfo.Config.Retention)
	_ = d.Set("segment_bytes", topicinfo.Config.SegmentBytes)
	_ = d.Set("segment", topicinfo.Config.SegmentMs)
	if topicinfo.Config.UncleanLeaderElectionEnable != nil {
		_ = d.Set("unclean_leader_election_enable", *topicinfo.Config.UncleanLeaderElectionEnable == 1)
	}
	return nil
}

func resourceTencentCloudCkafkaTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_topic.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ckafkcService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	items := strings.Split(d.Id(), FILED_SP)
	if len(items) < 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := items[0]
	topicName := items[1]

	request := ckafka.NewModifyTopicAttributesRequest()
	//replicaNum := d.Get("replica_num").(int)
	whiteListSwitch := d.Get("enable_white_list").(bool)
	cleanUpPolicy := d.Get("clean_up_policy").(string)
	retention := d.Get("retention").(int)
	var note string
	if v, ok := d.GetOk("note"); ok {
		note = v.(string)
		request.Note = &note
	}
	if v, ok := d.GetOk("segment"); ok {
		if v.(int) != 0 {
			request.SegmentMs = helper.IntInt64(v.(int))
		}
	}
	request.InstanceId = &instanceId
	request.TopicName = &topicName
	//request.ReplicaNum = helper.IntInt64(replicaNum)
	request.EnableWhiteList = helper.BoolToInt64Ptr(whiteListSwitch)
	request.MinInsyncReplicas = helper.IntInt64(d.Get("sync_replica_min_num").(int))
	request.UncleanLeaderElectionEnable = helper.BoolToInt64Ptr(d.Get("unclean_leader_election_enable").(bool))
	request.CleanUpPolicy = &cleanUpPolicy
	request.RetentionMs = helper.IntInt64(retention)
	if d.Get("max_message_bytes").(int) != 0 {
		request.MaxMessageBytes = helper.IntInt64(d.Get("max_message_bytes").(int))
	}
	//Update ip white List
	if whiteListSwitch {
		_, newInterface := d.GetChange("ip_white_list")
		newIpWhiteListInterface := newInterface.([]interface{})
		var newIpWhiteList []*string
		for _, value := range newIpWhiteListInterface {
			newIpWhiteList = append(newIpWhiteList, helper.String(value.(string)))
		}
		if len(newIpWhiteList) == 0 {
			return fmt.Errorf("this Topic %s Create Failed, reason: ip whitelist switch is on, ip whitelist cannot be empty", topicName)
		}
		request.IpWhiteList = newIpWhiteList
	} else {
		//IP whiteList Switch not turned on, and the ip whitelist cannot be modified
		if d.HasChange("ip_white_list") {
			return fmt.Errorf("this Topic %s IP whitelist Modification failed, reason: The Ip Whitelist Switch is not turned on", topicName)
		}
	}
	//Update partition num
	oldPartitionNum, newPartitionNum := d.GetChange("partition_num")
	if newPartitionNum.(int) < oldPartitionNum.(int) {
		return fmt.Errorf("this Topic %s partition Modification failed, reason: The partitonNum must more than current partitionNum", topicName)
	} else {
		if newPartitionNum.(int) > oldPartitionNum.(int) {
			err := ckafkcService.AddCkafkaTopicPartition(ctx, instanceId, topicName, int64(newPartitionNum.(int)))
			if err != nil {
				return err
			}
		}
	}
	err := ckafkcService.ModifyCkafkaTopicAttribute(ctx, request)
	if err != nil {
		return err
	}
	_, has, errDes := ckafkcService.DescribeCkafkaTopicByName(ctx, instanceId, topicName)
	if errDes != nil {
		return errDes
	}
	if !has {
		errDes = fmt.Errorf("this Topic %s Update Failed", topicName)
		return errDes
	}

	return resourceTencentCloudCkafkaTopicRead(d, meta)
}

func resourceTencentCLoudCkafkaTopicDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_topic.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ckafkcService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	items := strings.Split(d.Id(), FILED_SP)
	if len(items) < 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := items[0]
	topicName := items[1]

	err := ckafkcService.DeleteCkafkaTopic(ctx, instanceId, topicName)
	if err != nil {
		return err
	}

	return nil
}
