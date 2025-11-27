/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "tencentcloudenterprise_kubernetes_cluster_pvc" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"pvc\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"pvc\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
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
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_pvc", CNDescription{
		TerraformTypeCN: "集群持久化卷声明",
		DescriptionCN:   "提供集群持久化卷声明资源，用于配置集群的PVC。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"request_body": "请求体",
			"pvc_name":     "持久化卷声明名称",
		},
	})
}

func resourceTencentCloudTkeClusterPvc() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to pvc to cluster",
		Create:      resourceTencentCloudTkeTkeClusterPvcCreate,
		Read:        resourceTencentCloudTkeTkeClusterPvcRead,
		Delete:      resourceTencentCloudTkeTkeClusterPvcDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the cluster.",
			},
			"namespace": {
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
			"pvc_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "pvc_name",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterPvcCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pvc.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		ns          = d.Get("namespace").(string)
		requestBody = d.Get("request_body").(string)

		pvc  = d.Get("pvc_name").(string)
		path = fmt.Sprintf("/api/v1/namespaces/%s/persistentvolumeclaims", ns)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequestEncode(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody, false)
	if err != nil {
		return err
	}
	d.SetId(pvc)
	return resourceTencentCloudTkeTkeClusterPvcRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterPvcRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pvc.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		pvc       = d.Get("pvc_name").(string)
		clusterId = d.Get("cluster_id").(string)
		ns        = d.Get("namespace").(string)
		path      = fmt.Sprintf("/api/v1/namespaces/%s/persistentvolumeclaims?fieldSelector=metadata.name=%s", ns, pvc)
	)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	// var response pvcList

	// err = json.Unmarshal([]byte(body), &response)
	// if err != nil {
	// 	return err
	// }
	// if len(response.Items) == 0 {
	// 	return fmt.Errorf("pvc not found")
	// }
	return nil
}
func resourceTencentCloudTkeTkeClusterPvcDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_pvc.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId   = d.Get("cluster_id").(string)
		pvc         = d.Get("pvc_name").(string)
		ns          = d.Get("namespace").(string)
		path        = fmt.Sprintf("/api/v1/namespaces/%s/persistentvolumeclaims/%s", ns, pvc)
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)

	// delete pvc
	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}

	return nil
}
