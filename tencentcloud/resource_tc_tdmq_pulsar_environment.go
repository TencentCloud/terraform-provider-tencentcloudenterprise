package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_environment", CNDescription{
		TerraformTypeCN: "TDMQ Pulsar环境",
		DescriptionCN:   "提供TDMQ Pulsar环境资源，用于创建和管理TDMQ Pulsar环境。",
		AttributesCN: map[string]string{
			"environ_name":     "环境名称",
			"msg_ttl":          "消息过期时间,单位(秒)",
			"remark":           "备注",
			"cluster_id":       "集群ID",
			"retention_policy": "消息保留策略，格式为：`{time_in_minutes:Int，size_in_mb:Int}``time_in_minutes`：保留消息的时间`size_in_mb`：要保留的消息的大小",
		},
	})
}
func resourceTencentCloudTdmqPulsarEnvironment() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a TDMQ Pulsar Environment.",
		Create:      resourceTencentCloudTdmqPulsarEnvironmentCreate,
		Read:        resourceTencentCloudTdmqPulsarEnvironmentRead,
		Update:      resourceTencentCloudTdmqPulsarEnvironmentUpdate,
		Delete:      resourceTencentCloudTdmqPulsarEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"environ_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of namespace to be created.",
			},
			"msg_ttl": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The expiration time of unconsumed message.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Dedicated Cluster Id.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the namespace.",
			},
			"retention_policy": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_in_minutes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "the time of message to retain.",
						},
						"size_in_mb": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "the size of message to retain.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudTdmqPulsarEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_environment.create")()

	var (
		logId           = getLogId(contextNil)
		ctx             = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService     = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		retentionPolicy tdmq.RetentionPolicy
		environ_name    string
		msg_ttl         uint64
		remark          string
		clusterId       string
	)

	if temp, ok := d.GetOk("environ_name"); ok {
		environ_name = temp.(string)
		if len(environ_name) < 1 {
			return fmt.Errorf("environ_name should be not empty string")
		}
	}

	msg_ttl = uint64(d.Get("msg_ttl").(int))

	if temp, ok := d.GetOk("cluster_id"); ok {
		clusterId = temp.(string)
	}

	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}

	if temp, ok := d.GetOk("retention_policy"); ok {
		policy := temp.([]interface{})
		for _, item := range policy {
			value := item.(map[string]interface{})
			timeInMinutes := int64(value["time_in_minutes"].(int))
			sizeInMB := int64(value["size_in_mb"].(int))
			retentionPolicy.TimeInMinutes = &timeInMinutes
			retentionPolicy.SizeInMB = &sizeInMB
		}
	}

	environId, err := tdmqService.CreateTdmqNamespace(ctx, environ_name, msg_ttl, clusterId, remark, retentionPolicy)
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{environId, clusterId}, FILED_SP))
	return resourceTencentCloudTdmqPulsarEnvironmentRead(d, meta)
}

func resourceTencentCloudTdmqPulsarEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_environment.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		tdmqService = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	environId := idSplit[0]
	clusterId := idSplit[1]

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqNamespaceById(ctx, environId, clusterId)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("environ_name", info.EnvironmentId)
		_ = d.Set("cluster_id", clusterId)
		_ = d.Set("msg_ttl", info.MsgTTL)
		_ = d.Set("remark", info.Remark)

		tmpList := make([]map[string]interface{}, 0)
		retentionPolicy := make(map[string]interface{}, 2)
		retentionPolicy["time_in_minutes"] = info.RetentionPolicy.TimeInMinutes
		retentionPolicy["size_in_mb"] = info.RetentionPolicy.SizeInMB
		tmpList = append(tmpList, retentionPolicy)
		_ = d.Set("retention_policy", tmpList)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTdmqPulsarEnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_environment.update")()

	var (
		logId           = getLogId(contextNil)
		ctx             = context.WithValue(context.TODO(), logIdKey, logId)
		service         = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		msgTtl          uint64
		remark          string
		retentionPolicy = new(tdmq.RetentionPolicy)
	)

	immutableArgs := []string{"environ_name", "cluster_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	environId := idSplit[0]
	clusterId := idSplit[1]

	old, now := d.GetChange("msg_ttl")
	if d.HasChange("msg_ttl") {
		msgTtl = uint64(now.(int))
	} else {
		msgTtl = uint64(old.(int))
	}

	old, now = d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	_, now = d.GetChange("retention_policy")
	if d.HasChange("retention_policy") {
		policy := now.([]interface{})

		for _, item := range policy {
			value := item.(map[string]interface{})
			timeInMinutes := int64(value["time_in_minutes"].(int))
			sizeInMB := int64(value["size_in_mb"].(int))
			retentionPolicy.TimeInMinutes = &timeInMinutes
			retentionPolicy.SizeInMB = &sizeInMB
		}
	}

	d.Partial(true)

	if err := service.ModifyTdmqNamespaceAttribute(ctx, environId, msgTtl, remark, clusterId, retentionPolicy); err != nil {
		return err
	}

	d.Partial(false)
	return resourceTencentCloudTdmqPulsarEnvironmentRead(d, meta)
}

func resourceTencentCloudTdmqPulsarEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_environment.delete")()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	environId := idSplit[0]
	clusterId := idSplit[1]

	if err := service.DeleteTdmqNamespace(ctx, environId, clusterId); err != nil {
		return err
	}
	return nil
}
