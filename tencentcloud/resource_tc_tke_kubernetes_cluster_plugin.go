/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "tencentcloudenterprise_kubernetes_cluster_plugin" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"Plugin\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"Secret\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
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
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_plugin", CNDescription{
		TerraformTypeCN: "集群插件配置",
		DescriptionCN:   "提供集群插件配置资源，用于配置集群的插件。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"request_body": "请求体",
			"plugin_name":  "插件名称",
			"path":         "命名空间路径",
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
		Description: "Provide a resource to increase ns to cluster",
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
				Required:    true,
				Description: "ns name",
			},
			"request_body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "request_body",
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
	defer logElapsed("resource.cloud_kubernetes_cluster_plugin.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		requestBody = d.Get("request_body").(string)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	body, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	var response TkeForwardRequestPluginCreateResponse

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return err
	}
	// get app_id from rsp.metadata
	appid, exists := response.MetaData["name"].(string)
	if !exists {
		fmt.Println("Name field not found in metadata")
	}

	d.SetId(appid)
	_ = d.Set("app_id", appid)
	_ = d.Set("apiversion", response.ApiVersion)
	return resourceTencentCloudTkeTkeClusterPluginRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterPluginRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_plugin.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}
func resourceTencentCloudTkeTkeClusterPluginDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_plugin.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = fmt.Sprintf("/apis/application.tkestack.io/v1/namespaces/default/apps/%s", d.Id())
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	return nil
}
