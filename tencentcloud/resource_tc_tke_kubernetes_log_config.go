/*
Provide a resource to manage TKE CLS log configs via CreateCLSLogConfig/DescribeLogConfigs/DeleteLogConfigs.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_log_config" "log_config" {
	  cluster_id      = "cls-xxxxxx"
	  log_config_name = "test-tf"
	  logset_id       = "9066db41-1758-4b73-a3a5-310882aba884"

	  log_config = jsonencode({
	    apiVersion = "cls.cloud.tencent.com/v1"
	    kind       = "LogConfig"
	    metadata = {
	      name = "test-tf"
	    }
	    spec = {
	      clsDetail = {
	        logType    = "json_log"
	        logFormat  = "default"
	        region     = "ap-qingyuan-region-devtest-ops"
	        period     = 30
	        indexs     = [{ indexName = "namespace" }]
	        extractRule = {
	          unMatchUpload = "true"
	          unMatchedKey  = "LogParseFailure"
	          backtracking  = "-1"
	          jsonStandard  = "true"
	          isGBK         = "false"
	        }
	      }
	      inputDetail = {
	        type = "container_stdout"
	        containerStdout = {
	          allContainers    = true
	          namespace        = "default"
	          metadataContainer = ["namespace", "container_id", "container_name", "image_name", "cluster_id"]
	          metadataLabels    = ["__NULL__"]
	        }
	      }
	    }
	  })
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_log_config", CNDescription{
		TerraformTypeCN: "集群日志采集配置",
		DescriptionCN:   "提供TKE集群日志采集配置资源，用于创建和删除CLS采集规则。",
		AttributesCN: map[string]string{
			"cluster_id":      "集群ID",
			"log_config_name": "采集规则名称",
			"logset_id":       "CLS日志集ID",
			"log_config":      "日志采集配置JSON",
			"cluster_type":    "集群类型",
		},
	})
}

func resourceTencentCloudTkeKubernetesLogConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to manage TKE CLS log configs via CreateCLSLogConfig/DescribeLogConfigs/DeleteLogConfigs.",
		Create:      resourceTencentCloudTkeKubernetesLogConfigCreate,
		Read:        resourceTencentCloudTkeKubernetesLogConfigRead,
		Delete:      resourceTencentCloudTkeKubernetesLogConfigDelete,
		Schema: map[string]*schema.Schema{
			"log_config": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "JSON expression of log collection configuration.",
			},
			"log_config_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Log config name.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID.",
			},
			"logset_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "CLS log set ID.",
			},
			"cluster_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "tke",
				Description: "Cluster type, supported values: tke, eks. Default is tke.",
			},
		},
	}
}

func resourceTencentCloudTkeKubernetesLogConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_log_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Get("cluster_id").(string)
	logConfigName := d.Get("log_config_name").(string)
	clusterType := d.Get("cluster_type").(string)

	request := tke.NewCreateCLSLogConfigRequest()
	request.ClusterId = helper.String(clusterId)
	request.LogConfig = helper.String(d.Get("log_config").(string))
	request.ClusterType = helper.String(clusterType)
	if v, ok := d.GetOk("logset_id"); ok {
		request.LogsetId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().CreateCLSLogConfig(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{clusterId, logConfigName, clusterType}, FILED_SP))

	waitReq := tke.NewDescribeLogConfigsRequest()
	waitReq.ClusterId = helper.String(clusterId)
	waitReq.ClusterType = helper.String(clusterType)
	waitReq.LogConfigNames = helper.String(logConfigName)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DescribeLogConfigs(waitReq)
		if e != nil {
			if sdkErr, ok := e.(*sdkErrors.CloudSDKError); ok {
				if sdkErr.GetCode() == "FailedOperation.KubernetesGetOperationError" {
					return resource.RetryableError(fmt.Errorf("waiting for kubernetes log config to be created ready"))
				}
			}
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())

		if err := resourceTencentCloudTkeKubernetesLogConfigReadResponse(ctx, result); err != nil {
			return err
		}
		if result == nil || result.Response == nil || result.Response.LogConfigs == nil {
			return resource.RetryableError(fmt.Errorf("waiting for kubernetes log config to be created ready"))
		}
		return nil
	})
	if err != nil {
		return err
	}

	return resourceTencentCloudTkeKubernetesLogConfigRead(d, meta)
}

func resourceTencentCloudTkeKubernetesLogConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_log_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	clusterId := idSplit[0]
	logConfigName := idSplit[1]
	clusterType := idSplit[2]

	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("log_config_name", logConfigName)
	_ = d.Set("cluster_type", clusterType)

	var resp *tke.DescribeLogConfigsResponse
	reqErr := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesLogConfigById(ctx, clusterId, logConfigName, clusterType)
		if e != nil {
			return retryError(e)
		}
		if err := resourceTencentCloudTkeKubernetesLogConfigReadResponse(ctx, result); err != nil {
			return err
		}
		resp = result
		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s read kubernetes log config failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if resp == nil || resp.Response == nil || resp.Response.LogConfigs == nil {
		log.Printf("[WARN]%s resource `tke_kubernetes_log_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	return nil
}

func resourceTencentCloudTkeKubernetesLogConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_log_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	clusterId := idSplit[0]
	logConfigName := idSplit[1]
	clusterType := idSplit[2]

	request := tke.NewDeleteLogConfigsRequest()
	request.ClusterId = helper.String(clusterId)
	request.LogConfigNames = helper.String(logConfigName)
	request.ClusterType = helper.String(clusterType)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DeleteLogConfigs(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		if err := resourceTencentCloudTkeKubernetesLogConfigDeleteResponse(ctx, result); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	waitReq := tke.NewDescribeLogConfigsRequest()
	waitReq.ClusterId = helper.String(clusterId)
	waitReq.ClusterType = helper.String(clusterType)
	waitReq.LogConfigNames = helper.String(logConfigName)
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DescribeLogConfigs(waitReq)
		if e != nil {
			if sdkErr, ok := e.(*sdkErrors.CloudSDKError); ok {
				if sdkErr.GetCode() == "FailedOperation.KubernetesGetOperationError" {
					return nil
				}
			}
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
		return resource.RetryableError(fmt.Errorf("waiting for kubernetes log config to be deleted ready"))
	})
	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudTkeKubernetesLogConfigReadResponse(ctx context.Context, resp *tke.DescribeLogConfigsResponse) *resource.RetryError {
	if resp != nil && resp.Response != nil {
		message := resp.Response.Message
		if message != nil && *message != "" {
			return resource.NonRetryableError(fmt.Errorf(*message))
		}
	}
	return nil
}

func resourceTencentCloudTkeKubernetesLogConfigDeleteResponse(ctx context.Context, resp *tke.DeleteLogConfigsResponse) *resource.RetryError {
	if resp != nil && resp.Response != nil {
		message := resp.Response.Message
		if message != nil && *message != "" {
			return resource.NonRetryableError(fmt.Errorf(*message))
		}
	}
	return nil
}
