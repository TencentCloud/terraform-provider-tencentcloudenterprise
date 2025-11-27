package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_cluster", CNDescription{
		TerraformTypeCN: "TDMQ 实例",
		DescriptionCN:   "提供TDMQ Pulsar集群资源，用于创建和管理TDMQ Pulsar集群。",
		AttributesCN: map[string]string{
			"cluster_name": "TDMQ实例名称",
			"remark":       "备注",
			"tags":         "标签",
			//"project_id":        "项目ID",
			"bind_cluster_id":   "绑定的集群ID",
			"bind_cluster_name": "绑定的集群名称",
		},
	})
}
func resourceTencentCloudTdmqPulsarCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a TDMQ Pulsar cluster.",
		Create:      resourceTencentCloudTdmqPulsarClusterCreate,
		Read:        resourceTencentCloudTdmqPulsarClusterRead,
		Update:      resourceTencentCloudTdmqPulsarClusterUpdate,
		Delete:      resourceTencentCloudTdmqPulsarClusterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of tdmq cluster to be created.",
			},
			"bind_cluster_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Dedicated Cluster Id.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the tdmq cluster.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project ID.",
			},
			"bind_cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the bind cluster.",
			},
			//compute
			"cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The automatically generated ID of the TDMQ Pulsar cluster after creation.",
			},
		},
	}
}

func resourceTencentCloudTdmqPulsarClusterCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_cluster.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		request  = tdmq.NewCreateClusterRequest()
		response *tdmq.CreateClusterResponse
	)
	if v, ok := d.GetOk("cluster_name"); ok {
		request.ClusterName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bind_cluster_id"); ok {
		request.BindClusterId = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bind_cluster_name"); ok {
		request.BindClusterName = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().CreateCluster(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create tdmq instance failed, reason:%+v", logId, err)
		return err
	}

	clusterId := *response.Response.ClusterId

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::tdmq:%s:uin/:cluster/%s", region, clusterId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	d.SetId(clusterId)

	return resourceTencentCloudTdmqPulsarClusterRead(d, meta)
}

func resourceTencentCloudTdmqPulsarClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_cluster.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	tdmqService := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		info, has, e := tdmqService.DescribeTdmqInstanceById(ctx, id)
		if e != nil {
			return retryError(e)
		}
		if !has {
			d.SetId("")
			return nil
		}

		_ = d.Set("cluster_name", info.ClusterName)
		_ = d.Set("remark", info.Remark)
		_ = d.Set("cluster_id", info.ClusterId)
		return nil
	})
	if err != nil {
		return err
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "tdmq", "cluster", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudTdmqPulsarClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_cluster.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterName string
		remark      string
	)
	old, now := d.GetChange("cluster_name")
	if d.HasChange("cluster_name") {
		clusterName = now.(string)
	} else {
		clusterName = old.(string)
	}

	old, now = d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	if err := service.ModifyTdmqInstanceAttribute(ctx, id, clusterName, remark); err != nil {
		return err
	}

	if d.HasChange("tags") {
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("tdmq", "cluster", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}
	return resourceTencentCloudTdmqPulsarClusterRead(d, meta)
}

func resourceTencentCloudTdmqPulsarClusterDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tdmq_pulsar_cluster.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
	clusterId := d.Id()

	if err := service.DeleteTdmqInstance(ctx, clusterId); err != nil {
		return err
	}

	return nil
}
