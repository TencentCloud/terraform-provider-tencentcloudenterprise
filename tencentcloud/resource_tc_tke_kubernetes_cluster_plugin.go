/*
Provide a resource to manage TKE application in cluster via ForwardApplicationRequestV3.

# Example Usage

Repository chart:

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_plugin" "alertmanager" {
	  cluster_id = "cls-rkeuubqw"
	  namespace  = "tke-addon"

	  request_body = jsonencode({
	    kind       = "App"
	    apiVersion = "application.tkestack.io/v1"
	    metadata = {
	      name      = "alertmanager"
	      namespace = "tke-addon"
	      labels = {
	        "application.tkestack.io/type" = "internal-app"
	      }
	    }
	    spec = {
	      chart = {
	        chartGroupName = "local"
	        chartName      = "alertmanager"
	        chartVersion   = "1.7.0"
	        tenantID       = "local"
	        importedRepo   = true
	      }
	      name          = "alertmanager"
	      targetCluster = "cls-rkeuubqw"
	      type          = "HelmV3"
	      values = {
	        rawValues     = ""
	        rawValuesType = "yaml"
	        values        = ["rootDir="]
	      }
	      dryRun = false
	    }
	  })
	}

```

Local chart package (tgz):

```hcl

	locals {
	  chart_b64  = filebase64("${path.module}/charts/alertmanager-1.7.0.tgz")
	  values_yml = file("${path.module}/values.yaml")
	}

	resource "tencentcloudenterprise_tke_kubernetes_cluster_plugin" "alertmanager" {
	  cluster_id = "cls-rkeuubqw"
	  namespace  = "tke-cluster-inspection"

	  request_body = jsonencode({
	    kind       = "App"
	    apiVersion = "application.tkestack.io/v1"
	    metadata = {
	      name      = "alertmanager"
	      namespace = "tke-cluster-inspection"
	      annotations = {
	        "application.tkestack.io/chart" = local.chart_b64
	      }
	      labels = {
	        "application.tkestack.io/type" = "internal-app"
	      }
	    }
	    spec = {
	      chart = {
	        chartGroupName = "local"
	        chartName      = "alertmanager"
	        chartVersion   = "1.7.0"
	        importedRepo   = true
	      }
	      name          = "alertmanager"
	      targetCluster = "cls-rkeuubqw"
	      type          = "HelmV3"
	      values = {
	        rawValuesType = "yaml"
	        rawValues     = local.values_yml
	      }
	      dryRun = false
	    }
	  })
	}

```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_plugin", CNDescription{
		TerraformTypeCN: "集群插件配置",
		DescriptionCN:   "提供集群应用配置资源，用于创建/删除集群应用。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"request_body": "请求体",
			"plugin_name":  "插件名称",
			"path":         "应用创建路径，例如 /apis/application.tkestack.io/v1/namespaces/<namespace>/apps",
			"app_id":       "AppId信息",
			"apiversion":   "API版本信息",
		},
	})
}

type TkeForwardRequestPluginCreateResponse struct {
	MetaData   map[string]interface{} `json:"metadata"`
	Code       int                    `json:"code"`
	Kind       string                 `json:"kind"`
	ApiVersion string                 `json:"apiVersion"`
	Spec       map[string]interface{} `json:"spec"`
}

func resourceTencentCloudTkeClusterPlugin() *schema.Resource {
	return &schema.Resource{
		Description: "Manage a TKE application via ForwardApplicationRequestV3.",
		Create:      resourceTencentCloudTkeTkeClusterPluginCreate,
		Read:        resourceTencentCloudTkeTkeClusterPluginRead,
		Delete:      resourceTencentCloudTkeTkeClusterPluginDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the cluster.",
			},
			"path": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "Application endpoint path. If empty, defaults to /apis/application.tkestack.io/v1/namespaces/<namespace>/apps.",
			},
			"request_body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "request_body",
			},
			"namespace": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Namespace for the application.",
			},
			// Computed
			"app_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "app_id",
			},
			"apiversion": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "apiversion",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterPluginCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_plugin.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		requestBody = d.Get("request_body").(string)
		namespace   = d.Get("namespace").(string)
	)
	if path == "" {
		if namespace == "" {
			return fmt.Errorf("namespace is required when path is empty")
		}
		path = fmt.Sprintf("/apis/application.tkestack.io/v1/namespaces/%s/apps", namespace)
	}
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	body, err := service.ForwardApplicationRequestV3(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	var response TkeForwardRequestPluginCreateResponse

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return err
	}
	appid := ""
	if response.MetaData != nil {
		if name, exists := response.MetaData["name"].(string); exists && name != "" {
			appid = name
		}
	}
	if appid == "" && response.Spec != nil {
		if name, exists := response.Spec["name"].(string); exists && name != "" {
			appid = name
		}
	}
	if appid == "" {
		return fmt.Errorf("app name not found in response metadata.name or spec.name")
	}
	d.SetId(appid)
	_ = d.Set("app_id", appid)
	_ = d.Set("apiversion", response.ApiVersion)
	return resourceTencentCloudTkeTkeClusterPluginRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterPluginRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_plugin.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}
func resourceTencentCloudTkeTkeClusterPluginDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_plugin.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		namespace   = d.Get("namespace").(string)
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)
	if path == "" {
		if namespace == "" {
			return fmt.Errorf("namespace is required when path is empty")
		}
		path = fmt.Sprintf("/apis/application.tkestack.io/v1/namespaces/%s/apps", namespace)
	}
	trimmedPath := strings.TrimRight(path, "/")
	if !strings.HasSuffix(trimmedPath, "/"+d.Id()) {
		trimmedPath = fmt.Sprintf("%s/%s", trimmedPath, d.Id())
	}

	_, err := service.ForwardApplicationRequestV3(ctx, TKE_FORWARD_METHOD_DELETE, trimmedPath, clusterId, requestBody)
	if err != nil {
		return err
	}
	return nil
}
