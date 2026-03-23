/*
Provide a resource to create/read/delete a namespace in cluster

# Example Usage

# JSON

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_namespace" "ns" {
	  cluster_id = var.cluster_id
	  namespace  = "tf-test"

	  path = format("/apis/platform.tkestack.io/v1/clusters/%s/apply", var.cluster_id)

	  request_body = jsonencode({
	    apiVersion = "v1"
	    kind       = "Namespace"
	    metadata = {
	      name = "tf-test"
	      annotations = {
	        description = "ddddd"
	      }
	    }
	  })
	}

```

# YAML

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_namespace" "ns" {
	  cluster_id = var.cluster_id
	  namespace  = "tf-test"

	  path = format("/apis/platform.tkestack.io/v1/clusters/%s/apply", var.cluster_id)

	  request_body = <<-YAML

apiVersion: v1
kind: Namespace
metadata:

	name: tf-test
	labels:
	  env: dev
	annotations:
	  description: created-by-terraform

YAML
}
```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_namespace", CNDescription{
		TerraformTypeCN: "集群命名空间",
		DescriptionCN:   "提供集群命名空间资源。创建时需提供集群ID、apply路径及Namespace清单（apiVersion=v1, kind=Namespace）。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"path":         "Namespace创建的apply路径，例如 /apis/platform.tkestack.io/v1/clusters/<cluster_id>/apply",
			"request_body": "Namespace清单，支持JSON或YAML，需包含 apiVersion=v1 与 kind=Namespace",
			"namespace":    "命名空间名称（用于查询/删除）",
			"status":       "命名空间状态",
			"message":      "接口响应信息",
			"apiversion":   "API版本信息",
		},
	})
}

type TkeForwardRequestNamespaceResponse struct {
	MetaData map[string]interface{} `json:"metadata"`
	Status   string                 `json:"status"`
	Message  string                 `json:"message"`
	Code     int                    `json:"code"`
}

type TkeForwardRequestNamespaceReadResponse struct {
	Kind       string                   `json:"kind"`
	ApiVersion string                   `json:"apiVersion"`
	Metadata   map[string]interface{}   `json:"metadata"`
	Items      []map[string]interface{} `json:"items"`
}

func resourceTencentCloudTkeClusterNamespace() *schema.Resource {
	return &schema.Resource{
		Description: "Create/read/delete a namespace via TKE forward request. `path` should be the cluster apply endpoint, and `request_body` must be a Namespace manifest (apiVersion=v1, kind=Namespace).",
		Create:      resourceTencentCloudTkeTkeClusterNamespaceCreate,
		Read:        resourceTencentCloudTkeTkeClusterNamespaceRead,
		Delete:      resourceTencentCloudTkeTkeClusterNamespaceDelete,
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
				Required:    true,
				Description: "Apply endpoint for namespace creation, e.g. /apis/platform.tkestack.io/v1/clusters/<cluster_id>/apply.",
			},
			"request_body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Namespace manifest in JSON or YAML. Must include apiVersion=v1 and kind=Namespace.",
			},
			"namespace": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Namespace name used for read/delete.",
			},
			// Computed
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "status",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "message",
			},
			"apiversion": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "apiversion",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_namespace.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		requestBody = d.Get("request_body").(string)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	body, err := service.ForwardPlatformRequestV3(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	var rsp TkeForwardRequestNamespaceResponse
	err = json.Unmarshal([]byte(body), &rsp)
	if err != nil {
		return err
	}
	if rsp.Code != 200 {
		return fmt.Errorf(rsp.Message)
	}
	d.SetId(clusterId)
	_ = d.Set("status", rsp.Status)
	_ = d.Set("message", rsp.Message)

	return resourceTencentCloudTkeTkeClusterNamespaceRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterNamespaceRead(d *schema.ResourceData, meta interface{}) error {

	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_namespace.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId = d.Get("cluster_id").(string)
		namespace = d.Get("namespace").(string)
		path      = fmt.Sprintf("/api/v1/namespaces?fieldSelector=metadata.name=%s", namespace)
	)

	body, err := service.ForwardPlatformRequestV3(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	var rsp TkeForwardRequestNamespaceReadResponse
	err = json.Unmarshal([]byte(body), &rsp)
	if err != nil {
		return err
	}
	d.SetId(clusterId)
	_ = d.Set("apiversion", rsp.ApiVersion)
	return nil
}
func resourceTencentCloudTkeTkeClusterNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_namespace.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId = d.Get("cluster_id").(string)
		namespace = d.Get("namespace").(string)
		path      = fmt.Sprintf("/api/v1/namespaces/%s", namespace)
	)

	_, err := service.ForwardPlatformRequestV3(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, "")
	if err != nil {
		return err
	}
	return nil
}
