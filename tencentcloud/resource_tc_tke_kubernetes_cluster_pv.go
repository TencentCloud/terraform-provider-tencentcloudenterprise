/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "tencentcloudenterprise_kubernetes_cluster_pv" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"pv\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"pv\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	// "encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_pv", CNDescription{
		TerraformTypeCN: "集群持久化卷",
		DescriptionCN:   "提供集群持久化卷资源，用于配置集群的PV。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"request_body": "请求体",
			"pv_name":      "持久化卷名称",
			"path":         "命名空间路径",
		},
	})
}

func resourceTencentCloudTkeClusterPv() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to pv to cluster",
		Create:      resourceTencentCloudTkeTkeClusterPvCreate,
		Read:        resourceTencentCloudTkeTkeClusterPvRead,
		Delete:      resourceTencentCloudTkeTkeClusterPvDelete,
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
			"pv_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "pv_name",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterPvCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pv.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		requestBody = d.Get("request_body").(string)

		pv = d.Get("pv_name").(string)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequestEncode(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody, false)
	if err != nil {
		return err
	}
	d.SetId(pv)
	return resourceTencentCloudTkeTkeClusterPvRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterPvRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pv.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		pv        = d.Get("pv_name").(string)
		clusterId = d.Get("cluster_id").(string)
		path      = fmt.Sprintf("/api/v1/persistentvolumes?fieldSelector=metadata.name=%s", pv)
	)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	// var response pvList

	// err = json.Unmarshal([]byte(body), &response)
	// if err != nil {
	// 	return err
	// }
	// if len(response.Items) == 0 {
	// 	return fmt.Errorf("pv not found")
	// }
	return nil
}
func resourceTencentCloudTkeTkeClusterPvDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pv.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId   = d.Get("cluster_id").(string)
		pv          = d.Get("pv_name").(string)
		path        = fmt.Sprintf("/api/v1/persistentvolumes/%s", pv)
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)

	// delete pv
	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}

	return nil
}
