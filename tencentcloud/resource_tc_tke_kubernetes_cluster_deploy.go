/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "cloud_kubernetes_cluster_deploy" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"deploy\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"deploy\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
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
	registerResourceDescriptionProvider("cloud_tke_kubernetes_cluster_deploy", CNDescription{
		TerraformTypeCN: "集群部署配置",
		DescriptionCN:   "提供集群部署配置资源，用于向集群部署应用。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"request_body": "请求体",
			"deploy_name":  "部署名称",
		},
	})
}

func resourceTencentCloudTkeClusterDeploy() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to deploy to cluster",
		Create:      resourceTencentCloudTkeTkeClusterDeployCreate,
		Read:        resourceTencentCloudTkeTkeClusterDeployRead,
		Delete:      resourceTencentCloudTkeTkeClusterDeployDelete,
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
			"deploy_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "deploy_name",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterDeployCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_deploy.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		requestBody = d.Get("request_body").(string)

		deploy = d.Get("deploy_name").(string)
		path   = fmt.Sprintf("/apis/platform.tkestack.io/v1/clusters/%s/apply", clusterId)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	d.SetId(deploy)
	return resourceTencentCloudTkeTkeClusterDeployRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterDeployRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_deploy.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		deploy    = d.Get("deploy_name").(string)
		clusterId = d.Get("cluster_id").(string)
		ns        = d.Get("namespace").(string)
		path      = fmt.Sprintf("/apis/apps/v1/namespaces/%s/deployments?fieldSelector=metadata.name=%s", ns, deploy)
	)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	// var response deployList

	// err = json.Unmarshal([]byte(body), &response)
	// if err != nil {
	// 	return err
	// }
	// if len(response.Items) == 0 {
	// 	return fmt.Errorf("deploy not found")
	// }
	return nil
}
func resourceTencentCloudTkeTkeClusterDeployDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_deploy.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		clusterId   = d.Get("cluster_id").(string)
		deploy      = d.Get("deploy_name").(string)
		ns          = d.Get("namespace").(string)
		path        = fmt.Sprintf("/apis/apps/v1/namespaces/%s/deployments/%s", ns, deploy)
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)

	// delete deploy
	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}

	return nil
}
