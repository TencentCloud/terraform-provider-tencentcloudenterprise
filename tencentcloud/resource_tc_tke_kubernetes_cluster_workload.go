/*
Provide a resource to manage Kubernetes workload via TKE forward request.

# Example Usage

# Deployment (YAML)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "deployment" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "deployment"
	  workload_name = "tf-deploy"

	  request_body = <<-YAML
	  apiVersion: apps/v1
	  kind: Deployment
	  metadata:
	    name: tf-deploy
	    namespace: default
	    labels:
	      qcloud-app: tf-deploy
	      k8s-app: tf-deploy
	    annotations:
	      description: demo deployment
	  spec:
	    replicas: 2
	    selector:
	      matchLabels:
	        k8s-app: tf-deploy
	    template:
	      metadata:
	        labels:
	          qcloud-app: tf-deploy
	          k8s-app: tf-deploy
	      spec:
	        containers:
	          - name: nginx
	            image: nginx:stable
	            ports:
	              - name: http
	                containerPort: 80
	            env:
	              - name: ENV
	                value: prod
	            resources:
	              requests:
	                cpu: "100m"
	                memory: "128Mi"
	              limits:
	                cpu: "500m"
	                memory: "512Mi"
	  YAML
	}

```

# Deployment (JSON)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "deployment_json" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "deployment"
	  workload_name = "tf-deploy"

	  request_body = jsonencode({
	    apiVersion = "apps/v1"
	    kind       = "Deployment"
	    metadata = {
	      name      = "tf-deploy"
	      namespace = "default"
	      labels = {
	        "qcloud-app" = "tf-deploy"
	        "k8s-app"    = "tf-deploy"
	      }
	      annotations = { description = "demo deployment" }
	    }
	    spec = {
	      replicas = 2
	      selector = { matchLabels = { "k8s-app" = "tf-deploy" } }
	      template = {
	        metadata = { labels = { "qcloud-app" = "tf-deploy", "k8s-app" = "tf-deploy" } }
	        spec = {
	          containers = [{
	            name  = "nginx"
	            image = "nginx:stable"
	            ports = [{ name = "http", containerPort = 80 }]
	            env   = [{ name = "ENV", value = "prod" }]
	            resources = {
	              requests = { cpu = "100m", memory = "128Mi" }
	              limits   = { cpu = "500m", memory = "512Mi" }
	            }
	          }]
	        }
	      }
	    }
	  })
	}

```

# StatefulSet (YAML)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "statefulset" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "statefulset"
	  workload_name = "tf-ss"

	  request_body = <<-YAML
	  apiVersion: apps/v1
	  kind: StatefulSet
	  metadata:
	    name: tf-ss
	    namespace: default
	    labels:
	      qcloud-app: tf-ss
	      k8s-app: tf-ss
	  spec:
	    serviceName: tf-ss
	    replicas: 1
	    selector:
	      matchLabels:
	        k8s-app: tf-ss
	    template:
	      metadata:
	        labels:
	          qcloud-app: tf-ss
	          k8s-app: tf-ss
	      spec:
	        containers:
	          - name: app
	            image: busybox
	            command: ["sh", "-c", "sleep 3600"]
	            volumeMounts:
	              - name: data
	                mountPath: /data
	        volumes:
	          - name: data
	            emptyDir: {}
	  YAML
	}

```

# StatefulSet (JSON)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "statefulset_json" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "statefulset"
	  workload_name = "tf-ss"

	  request_body = jsonencode({
	    apiVersion = "apps/v1"
	    kind       = "StatefulSet"
	    metadata   = { name = "tf-ss", namespace = "default", labels = { "qcloud-app" = "tf-ss", "k8s-app" = "tf-ss" } }
	    spec = {
	      serviceName = "tf-ss"
	      replicas    = 1
	      selector    = { matchLabels = { "k8s-app" = "tf-ss" } }
	      template = {
	        metadata = { labels = { "qcloud-app" = "tf-ss", "k8s-app" = "tf-ss" } }
	        spec = {
	          containers = [{
	            name         = "app"
	            image        = "busybox"
	            command      = ["sh", "-c", "sleep 3600"]
	            volumeMounts = [{ name = "data", mountPath = "/data" }]
	          }]
	          volumes = [{ name = "data", emptyDir = map[string]interface{}{} }]
	        }
	      }
	    }
	  })
	}

```

# DaemonSet (YAML)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "daemonset" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "daemonset"
	  workload_name = "tf-ds"

	  request_body = <<-YAML
	  apiVersion: apps/v1
	  kind: DaemonSet
	  metadata:
	    name: tf-ds
	    namespace: default
	    labels:
	      qcloud-app: tf-ds
	      k8s-app: tf-ds
	  spec:
	    selector:
	      matchLabels:
	        k8s-app: tf-ds
	    template:
	      metadata:
	        labels:
	          qcloud-app: tf-ds
	          k8s-app: tf-ds
	      spec:
	        containers:
	          - name: agent
	            image: busybox
	            args: ["sh", "-c", "echo running; sleep 3600"]
	            resources:
	              requests:
	                cpu: "50m"
	                memory: "64Mi"
	  YAML
	}

```

# DaemonSet (JSON)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "daemonset_json" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "daemonset"
	  workload_name = "tf-ds"

	  request_body = jsonencode({
	    apiVersion = "apps/v1"
	    kind       = "DaemonSet"
	    metadata   = { name = "tf-ds", namespace = "default", labels = { "qcloud-app" = "tf-ds", "k8s-app" = "tf-ds" } }
	    spec = {
	      selector = { matchLabels = { "k8s-app" = "tf-ds" } }
	      template = {
	        metadata = { labels = { "qcloud-app" = "tf-ds", "k8s-app" = "tf-ds" } }
	        spec = {
	          containers = [{
	            name  = "agent"
	            image = "busybox"
	            args  = ["sh", "-c", "echo running; sleep 3600"]
	            resources = {
	              requests = { cpu = "50m", memory = "64Mi" }
	            }
	          }]
	        }
	      }
	    }
	  })
	}

```

# Job (YAML)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "job" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "job"
	  workload_name = "tf-job"

	  request_body = <<-YAML
	  apiVersion: batch/v1
	  kind: Job
	  metadata:
	    name: tf-job
	    namespace: default
	    labels:
	      qcloud-app: tf-job
	      k8s-app: tf-job
	  spec:
	    completions: 1
	    parallelism: 1
	    template:
	      spec:
	        restartPolicy: OnFailure
	        containers:
	          - name: job
	            image: busybox
	            command: ["sh", "-c", "echo hello; sleep 5"]
	            env:
	              - name: TASK
	                value: demo
	  YAML
	}

```

# Job (JSON)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "job_json" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "job"
	  workload_name = "tf-job"

	  request_body = jsonencode({
	    apiVersion = "batch/v1"
	    kind       = "Job"
	    metadata = {
	      name      = "tf-job"
	      namespace = "default"
	      labels    = { "qcloud-app" = "tf-job", "k8s-app" = "tf-job" }
	    }
	    spec = {
	      completions = 1
	      parallelism = 1
	      template = {
	        spec = {
	          restartPolicy = "OnFailure"
	          containers = [{
	            name    = "job"
	            image   = "busybox"
	            command = ["sh", "-c", "echo hello; sleep 5"]
	            env     = [{ name = "TASK", value = "demo" }]
	          }]
	        }
	      }
	    }
	  })
	}

```

# CronJob (YAML)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "cronjob" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "cronjob"
	  workload_name = "tf-123"

	  # If you have base64-encoded YAML/JSON, set encoded_body = true
	  request_body = <<-YAML
	  apiVersion: batch/v1
	  kind: CronJob
	  metadata:
	    name: tf-123
	    namespace: default
	    labels:
	      qcloud-app: tf-123
	      k8s-app: tf-123
	    annotations:
	      description: demo cronjob
	  spec:
	    schedule: "0 0/12 * * *"
	    successfulJobsHistoryLimit: 3
	    failedJobsHistoryLimit: 1
	    jobTemplate:
	      spec:
	        template:
	          metadata:
	            labels:
	              qcloud-app: tf-123
	              k8s-app: tf-123
	          spec:
	            restartPolicy: OnFailure
	            containers:
	              - name: cc11
	                image: ccr.d12-x86.fsphere.cn/test1/test-private
	                resources:
	                  requests:
	                    cpu: "100m"
	                    memory: "128Mi"
	                  limits:
	                    cpu: "500m"
	                    memory: "512Mi"
	  YAML
	}

```

# CronJob (JSON)

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_workload" "cronjob_json" {
	  cluster_id    = "cls-rkeuubqw"
	  namespace     = "default"
	  workload_kind = "cronjob"
	  workload_name = "tf-123"

	  request_body = jsonencode({
	    apiVersion = "batch/v1"
	    kind       = "CronJob"
	    metadata = {
	      name      = "tf-123"
	      namespace = "default"
	      labels    = { "qcloud-app" = "tf-123", "k8s-app" = "tf-123" }
	      annotations = { description = "demo cronjob" }
	    }
	    spec = {
	      schedule                   = "0 0/12 * * *"
	      successfulJobsHistoryLimit = 3
	      failedJobsHistoryLimit     = 1
	      jobTemplate = {
	        spec = {
	          template = {
	            metadata = { labels = { "qcloud-app" = "tf-123", "k8s-app" = "tf-123" } }
	            spec = {
	              restartPolicy = "OnFailure"
	              containers = [{
	                name  = "cc11"
	                image = "ccr.d12-x86.fsphere.cn/test1/test-private"
	                resources = {
	                  requests = { cpu = "100m", memory = "128Mi" }
	                  limits   = { cpu = "500m", memory = "512Mi" }
	                }
	              }]
	            }
	          }
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
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	tkeWorkloadKindDeployment  = "deployment"
	tkeWorkloadKindStatefulSet = "statefulset"
	tkeWorkloadKindDaemonSet   = "daemonset"
	tkeWorkloadKindJob         = "job"
	tkeWorkloadKindCronJob     = "cronjob"

	tkeWorkloadModeApply  = "apply"
	tkeWorkloadModeDirect = "direct"
	tkeWorkloadModeProxy  = "proxy"
)

type tkeWorkloadSpec struct {
	apiVersion string
	resource   string
}

var tkeWorkloadKindSpecs = map[string]tkeWorkloadSpec{
	tkeWorkloadKindDeployment:  {apiVersion: "apps/v1", resource: "deployments"},
	tkeWorkloadKindStatefulSet: {apiVersion: "apps/v1", resource: "statefulsets"},
	tkeWorkloadKindDaemonSet:   {apiVersion: "apps/v1", resource: "daemonsets"},
	tkeWorkloadKindJob:         {apiVersion: "batch/v1", resource: "jobs"},
	tkeWorkloadKindCronJob:     {apiVersion: "batch/v1", resource: "cronjobs"},
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_workload", CNDescription{
		TerraformTypeCN: "集群工作负载",
		DescriptionCN:   "提供集群工作负载资源，用于创建/读取/删除 Deployment、StatefulSet、DaemonSet、Job 与 CronJob。",
		AttributesCN: map[string]string{
			"cluster_id":    "集群ID",
			"namespace":     "命名空间名称",
			"request_body":  "工作负载清单，支持JSON或YAML（可为base64编码）",
			"encoded_body":  "request_body是否为base64编码",
			"workload_kind": "工作负载类型：deployment/statefulset/daemonset/job/cronjob",
			"workload_name": "工作负载名称",
		},
	})
}

func resourceTencentCloudTkeClusterWorkload() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to manage Kubernetes workloads via TKE forward request.",
		Create:      resourceTencentCloudTkeTkeClusterWorkloadCreate,
		Read:        resourceTencentCloudTkeTkeClusterWorkloadRead,
		Delete:      resourceTencentCloudTkeTkeClusterWorkloadDelete,
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
				Description: "Namespace name.",
			},
			"request_body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Workload manifest in JSON or YAML. Can be base64 encoded when encoded_body=true.",
			},
			"encoded_body": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether request_body is base64 encoded.",
			},
			"workload_kind": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Workload kind: deployment/statefulset/daemonset/job/cronjob.",
			},
			"workload_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Workload name.",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterWorkloadCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_workload.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		requestBody = d.Get("request_body").(string)
		encodedBody = d.Get("encoded_body").(bool)
		name        = d.Get("workload_name").(string)
		kind        = d.Get("workload_kind").(string)
		namespace   = d.Get("namespace").(string)
	)

	kind = strings.ToLower(strings.TrimSpace(kind))
	name = strings.TrimSpace(name)
	namespace = strings.TrimSpace(namespace)

	if kind == "" || name == "" || namespace == "" {
		return errors.New("workload_kind/workload_name/namespace are required")
	}

	resolvedVersion, resource, err := tkeResolveWorkloadKind(kind, "")
	if err != nil {
		return err
	}

	mode := tkeResolveWorkloadMode(kind, "")
	path := tkeBuildWorkloadCreatePath(clusterId, namespace, resolvedVersion, resource, mode)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	contentType := tkeDetectContentType(requestBody, encodedBody)
	_, err = service.ForwardPlatformRequestV3WithOptions(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody, "", contentType, &encodedBody)
	if err != nil {
		return err
	}
	d.SetId(name)
	_ = d.Set("workload_kind", kind)
	_ = d.Set("workload_name", name)
	_ = d.Set("namespace", namespace)
	return resourceTencentCloudTkeTkeClusterWorkloadRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterWorkloadRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_workload.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		ns        = d.Get("namespace").(string)
		name      = d.Get("workload_name").(string)
		kind      = d.Get("workload_kind").(string)
		clusterId = d.Get("cluster_id").(string)
	)

	if kind == "" || name == "" || ns == "" {
		return errors.New("workload_kind/workload_name/namespace are required for read")
	}

	resolvedVersion, resource, err := tkeResolveWorkloadKind(kind, "")
	if err != nil {
		return err
	}

	mode := tkeResolveWorkloadMode(kind, "")
	path := tkeBuildWorkloadReadPath(clusterId, ns, resolvedVersion, resource, name, mode)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err = service.ForwardPlatformRequestV3(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudTkeTkeClusterWorkloadDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_kubernetes_cluster_workload.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId = d.Get("cluster_id").(string)
		kind      = d.Get("workload_kind").(string)
		name      = d.Get("workload_name").(string)
		ns        = d.Get("namespace").(string)
	)

	if kind == "" || name == "" || ns == "" {
		return errors.New("workload_kind/workload_name/namespace are required for delete")
	}

	resolvedVersion, resource, err := tkeResolveWorkloadKind(kind, "")
	if err != nil {
		return err
	}

	mode := tkeResolveWorkloadMode(kind, "")
	path := tkeBuildWorkloadDeletePath(clusterId, ns, resolvedVersion, resource, name, mode)
	requestBody := "{\"propagationPolicy\":\"Background\"}"

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err = service.ForwardPlatformRequestV3(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	return nil
}

func tkeResolveWorkloadKind(kind, apiVersion string) (string, string, error) {
	normalized := strings.ToLower(strings.TrimSpace(kind))
	spec, exists := tkeWorkloadKindSpecs[normalized]
	if !exists {
		return "", "", fmt.Errorf("unsupported workload_kind %q", kind)
	}
	if apiVersion != "" {
		spec.apiVersion = apiVersion
	}
	return spec.apiVersion, spec.resource, nil
}

func tkeResolveWorkloadMode(kind, mode string) string {
	switch mode {
	case tkeWorkloadModeApply, tkeWorkloadModeDirect, tkeWorkloadModeProxy:
		return mode
	case "":
		return tkeDefaultModeForKind(kind)
	default:
		return tkeDefaultModeForKind(kind)
	}
}

func tkeDefaultModeForKind(kind string) string {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case tkeWorkloadKindCronJob:
		return tkeWorkloadModeProxy
	case tkeWorkloadKindJob:
		return tkeWorkloadModeDirect
	default:
		return tkeWorkloadModeApply
	}
}

func tkeBuildWorkloadCreatePath(clusterId, namespace, apiVersion, resource, mode string) string {
	switch mode {
	case tkeWorkloadModeDirect:
		return fmt.Sprintf("/apis/%s/namespaces/%s/%s", apiVersion, namespace, resource)
	case tkeWorkloadModeProxy:
		inner := fmt.Sprintf("/apis/%s/namespaces/%s/%s", apiVersion, namespace, resource)
		return tkeBuildProxyPath(clusterId, inner)
	default:
		return fmt.Sprintf("/apis/platform.tke/v1/clusters/%s/apply?notUpdate=true", clusterId)
	}
}

func tkeBuildWorkloadReadPath(clusterId, namespace, apiVersion, resource, name, mode string) string {
	inner := fmt.Sprintf("/apis/%s/namespaces/%s/%s?fieldSelector=metadata.name=%s", apiVersion, namespace, resource, name)
	if mode == tkeWorkloadModeProxy {
		return tkeBuildProxyPath(clusterId, inner)
	}
	return inner
}

func tkeBuildWorkloadDeletePath(clusterId, namespace, apiVersion, resource, name, mode string) string {
	inner := fmt.Sprintf("/apis/%s/namespaces/%s/%s/%s", apiVersion, namespace, resource, name)
	if mode == tkeWorkloadModeProxy {
		return tkeBuildProxyPath(clusterId, inner)
	}
	return inner
}

func tkeBuildProxyPath(clusterId, innerPath string) string {
	return fmt.Sprintf("/apis/platform.tke/v1/clusters/%s/proxy?path=%s", clusterId, url.QueryEscape(innerPath))
}

func tkeDetectContentType(requestBody string, encoded bool) string {
	body := strings.TrimSpace(requestBody)
	if body == "" {
		return ""
	}
	if encoded {
		decoded, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			return ""
		}
		body = strings.TrimSpace(string(decoded))
		if body == "" {
			return ""
		}
	}
	if strings.HasPrefix(body, "{") || strings.HasPrefix(body, "[") {
		return "application/json"
	}
	return "application/yaml"
}
