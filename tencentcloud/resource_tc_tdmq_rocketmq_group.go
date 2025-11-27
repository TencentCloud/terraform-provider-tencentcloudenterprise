/*
Provides a resource to create a tdmqRocketmq group

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
		cluster_name = "test_rocketmq"
		remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  namespace_name = "test_namespace"
	  ttl = 65000
	  retention_time = 65000
	  remark = "test namespace"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_group" "group" {
	  group_name = "test_rocketmq_group"
	  namespace = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	  read_enable = true
	  broadcast_enable = true
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  remark = "test rocketmq group"
	}

```
Import

tdmqRocketmq group can be imported using the id, e.g.
```
$ terraform import cloud_tdmq_rocketmq_group.group group_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_group", CNDescription{
		TerraformTypeCN: "Rocketmq 消费组",
		DescriptionCN:   "提供Rocketmq消费组资源，用于创建和管理Rocketmq消费组。",
		AttributesCN: map[string]string{
			"group_name":          "Group名称，8~64个字符",
			"remark":              "说明信息，最长128个字符",
			"namespace":           "命名空间，目前只支持单个命名空间",
			"read_enable":         "是否开启消费",
			"broadcast_enable":    "是否开启广播消费",
			"cluster_id":          "集群ID",
			"group_type":          "Group类型（TCP/HTTP）",
			"consumer_num":        "在线消费者数量",
			"tps":                 "消费TPS",
			"total_accumulative":  "堆积消息总数",
			"consumption_mode":    "消费模式，0：集群消费，1：广播消费，-1：未知",
			"retry_partition_num": "重试队列分区数",
			"create_time":         "创建时间戳",
			"update_time":         "更新时间戳",
			"client_protocol":     "客户端协议",
			"consumer_type":       "消费者类型，主动或者被动",
		},
	})
}
func resourceTencentCloudTdmqRocketmqGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tdmqRocketmq group",
		Create:      resourceTencentCloudTdmqRocketmqGroupCreate,
		Read:        resourceTencentCloudTdmqRocketmqGroupRead,
		Update:      resourceTencentCloudTdmqRocketmqGroupUpdate,
		Delete:      resourceTencentCloudTdmqRocketmqGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			// Group名称，8~64个字符
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Group name (8-64 characters).",
			},
			// 说明信息，最长128个字符
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remarks (up to 128 characters).",
			},
			// 命名空间，目前只支持单个命名空间
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Namespace. Currently, only one namespace is supported.",
			},
			// 是否开启消费
			"read_enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable consumption.",
			},
			// 是否开启广播消费
			"broadcast_enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable broadcast consumption.",
			},
			// 集群ID
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID.",
			},
			// Group类型（TCP/HTTP）
			"group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group type. Enumerated values: `NORMAL`: TCP; `RETRY`: HTTP.",
			},

			// Computed
			//
			"consumer_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of online consumers.",
			},

			"tps": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Consumption TPS.",
			},

			"total_accumulative": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of heaped messages.",
			},

			"consumption_mode": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "`0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.",
			},

			"retry_partition_num": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of partitions in a retry topic.",
			},

			"create_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation time in milliseconds.",
			},

			"update_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Modification time in milliseconds.",
			},

			"client_protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Client protocol.",
			},

			"consumer_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Consumer type. Enumerated values: ACTIVELY or PASSIVELY.",
			},
		},
	}
}

func resourceTencentCloudTdmqRocketmqGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_group.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request       = tdmqRocketmq.NewCreateRocketMQGroupRequest()
		clusterId     string
		namespaceName string
		groupName     string
	)

	if v, ok := d.GetOk("group_name"); ok {
		groupName = v.(string)
		request.GroupId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("namespace"); ok {
		namespaceName = v.(string)
		request.Namespaces = []*string{&namespaceName}
	}

	if v, _ := d.GetOk("read_enable"); v != nil {
		request.ReadEnable = helper.Bool(v.(bool))
	}

	if v, _ := d.GetOk("broadcast_enable"); v != nil {
		request.BroadcastEnable = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId = v.(string)
		request.ClusterId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("remark"); ok {

		request.Remark = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateRocketMQGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmqRocketmq group failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(clusterId + FILED_SP + namespaceName + FILED_SP + groupName)
	return resourceTencentCloudTdmqRocketmqGroupRead(d, meta)
}

func resourceTencentCloudTdmqRocketmqGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	namespaceName := idSplit[1]
	groupName := idSplit[2]

	groupList, err := service.DescribeTdmqRocketmqGroup(ctx, clusterId, namespaceName, groupName)

	if err != nil {
		return err
	}

	if len(groupList) == 0 {
		d.SetId("")
		return fmt.Errorf("resource `group` %s does not exist", groupName)
	}
	group := groupList[0]
	_ = d.Set("group_name", group.Name)
	_ = d.Set("namespace", namespaceName)
	_ = d.Set("read_enable", group.ReadEnabled)
	_ = d.Set("broadcast_enable", group.BroadcastEnabled)
	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("remark", group.Remark)
	_ = d.Set("consumer_num", group.ConsumerNum)
	_ = d.Set("tps", group.TPS)
	_ = d.Set("total_accumulative", group.TotalAccumulative)
	_ = d.Set("consumption_mode", group.ConsumptionMode)
	_ = d.Set("retry_partition_num", group.RetryPartitionNum)
	_ = d.Set("create_time", group.CreateTime)
	_ = d.Set("update_time", group.UpdateTime)
	_ = d.Set("client_protocol", group.ClientProtocol)
	_ = d.Set("consumer_type", group.ConsumerType)

	return nil
}

func resourceTencentCloudTdmqRocketmqGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_group.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tdmqRocketmq.NewModifyRocketMQGroupRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	namespaceId := idSplit[1]
	groupId := idSplit[2]

	request.ClusterId = &clusterId
	request.NamespaceId = &namespaceId
	request.GroupId = &groupId

	if d.HasChange("group_id") {

		return fmt.Errorf("`group_id` do not support change now.")

	}

	if d.HasChange("namespaces") {

		return fmt.Errorf("`namespaces` do not support change now.")

	}

	if d.HasChange("read_enable") {
		if v, ok := d.GetOk("read_enable"); ok {
			request.ReadEnable = helper.Bool(v.(bool))
		}

	}

	if d.HasChange("broadcast_enable") {
		if v, ok := d.GetOk("broadcast_enable"); ok {
			request.BroadcastEnable = helper.Bool(v.(bool))
		}

	}

	if d.HasChange("cluster_id") {

		return fmt.Errorf("`cluster_id` do not support change now.")

	}

	if d.HasChange("remark") {
		if v, ok := d.GetOk("remark"); ok {
			request.Remark = helper.String(v.(string))
		}

	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().ModifyRocketMQGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmqRocketmq group failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTdmqRocketmqGroupRead(d, meta)
}

func resourceTencentCloudTdmqRocketmqGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmqRocketmq_group.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	namespaceId := idSplit[1]
	groupId := idSplit[2]

	if err := service.DeleteTdmqRocketmqGroupById(ctx, clusterId, namespaceId, groupId); err != nil {
		return err
	}

	return nil
}
