/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "tencentcloudenterprise_kubernetes_cluster_namespace" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"Namespace\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"Secret\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
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
		DescriptionCN:   "提供集群命名空间资源，用于向集群增加命名空间。",
		AttributesCN: map[string]string{
			"cluster_id":            "集群ID",
			"worker_config":         "worker实例配置信息",
			"labels":                "worker实例标签",
			"extra_args":            "worker实例额外参数",
			"gpu_args":              "worker实例GPU参数",
			"unschedulable":         "worker实例是否参与调度",
			"desired_pod_num":       "worker实例期望pod数量",
			"docker_graph_path":     "worker实例docker graph路径",
			"mount_target":          "worker实例挂载目标",
			"data_disk":             "worker实例数据盘配置",
			"worker_instances_list": "worker实例列表",
			"path":                  "命名空间路径",
			"request_body":          "命名空间相关请求主体",
			"namespace":             "集群命名空间",
			"status":                "集群命名空间状态",
			"message":               "接口响应信息",
			"apiversion":            "API版本信息",
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
		Description: "Provide a resource to increase ns to cluster",
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
				Description: "ns name",
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
				Description: "namespace",
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
	defer logElapsed("resource.cloud_kubernetes_cluster_namespace.create")()
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

	defer logElapsed("resource.cloud_kubernetes_cluster_namespace.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId = d.Get("cluster_id").(string)
		namespace = d.Get("namespace").(string)
		path      = fmt.Sprintf("/api/v1/namespaces?fieldSelector=metadata.name=%s", namespace)
	)

	body, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
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
	defer logElapsed("resource.cloud_kubernetes_cluster_namespace.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId = d.Get("cluster_id").(string)
		namespace = d.Get("namespace").(string)
		path      = fmt.Sprintf("/api/v1/namespaces/%s", namespace)
	)

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, "")
	if err != nil {
		return err
	}
	return nil
}
