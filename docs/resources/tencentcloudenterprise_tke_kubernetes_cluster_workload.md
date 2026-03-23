---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_cluster_workload"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_cluster_workload"
description: |-
  Provide a resource to manage Kubernetes workload via TKE forward request.
---

# tencentcloudenterprise_tke_kubernetes_cluster_workload

Provide a resource to manage Kubernetes workload via TKE forward request.

## Example Usage

### # Deployment (YAML)

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

### # Deployment (JSON)

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

### # StatefulSet (YAML)

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

### # StatefulSet (JSON)

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
          volumes = [{ name = "data", emptyDir = map[string] interface {} {} }]
        }
      }
    }
  })
}
```

### # DaemonSet (YAML)

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

### # DaemonSet (JSON)

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

### # Job (YAML)

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

### # Job (JSON)

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

### # CronJob (YAML)

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

### # CronJob (JSON)

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
      name        = "tf-123"
      namespace   = "default"
      labels      = { "qcloud-app" = "tf-123", "k8s-app" = "tf-123" }
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

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the cluster.
* `namespace` - (Required, String, ForceNew) Namespace name.
* `request_body` - (Required, String, ForceNew) Workload manifest in JSON or YAML. Can be base64 encoded when encoded_body=true.
* `workload_kind` - (Required, String, ForceNew) Workload kind: deployment/statefulset/daemonset/job/cronjob.
* `workload_name` - (Required, String, ForceNew) Workload name.
* `encoded_body` - (Optional, Bool, ForceNew) Whether request_body is base64 encoded.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


