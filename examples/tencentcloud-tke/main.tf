# Tencent Kubernetes Engine (TKE) Examples

# ========== Data Sources ==========

# Query TKE clusters
data "cloud_tke_kubernetes_clusters" "clusters" {
  cluster_id = "cls-xxxxx"
}

# Query TKE available cluster versions
data "cloud_tke_kubernetes_available_cluster_versions" "versions" {
  cluster_ids = ["cls-xxxxx"]
}

# Query TKE charts
data "cloud_tke_kubernetes_charts" "charts" {
  kind = "log"
}

# Query TKE cluster common names
data "cloud_tke_kubernetes_cluster_common_names" "names" {
  cluster_id = "cls-xxxxx"
}

# ========== Resources ==========

# TKE Cluster
resource "cloud_tke_kubernetes_cluster" "cluster" {
  cluster_name              = "example-tke"
  cluster_version           = "1.24.4"
  cluster_cidr              = "172.16.0.0/16"
  cluster_max_pod_num       = 64
  cluster_max_service_num   = 256
  vpc_id                    = "vpc-xxxxx"
  cluster_internet          = true
  cluster_internet_security_group = "sg-xxxxx"
  
  worker_config {
    count                      = 3
    availability_zone          = "ap-guangzhou-3"
    instance_type              = "S5.MEDIUM4"
    system_disk_type           = "CLOUD_PREMIUM"
    system_disk_size           = 50
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 10
    public_ip_assigned         = true
    subnet_id                  = "subnet-xxxxx"
    
    data_disk {
      disk_type = "CLOUD_PREMIUM"
      disk_size = 100
    }
    
    enhanced_security_service = false
    enhanced_monitor_service  = false
  }
  
  cluster_deploy_type = "MANAGED_CLUSTER"
  
  tags = {
    env = "test"
  }
}

# TKE Cluster Attachment (Add existing nodes)
resource "cloud_tke_kubernetes_cluster_attachment" "attachment" {
  cluster_id  = cloud_tke_kubernetes_cluster.cluster.id
  instance_id = "ins-xxxxx"
  hostname    = "node-1"
  
  labels = {
    role = "worker"
  }
}

# TKE Scale Worker (Add nodes)
resource "cloud_tke_kubernetes_scale_worker" "scale" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  
  worker_config {
    count                      = 2
    availability_zone          = "ap-guangzhou-3"
    instance_type              = "S5.MEDIUM4"
    system_disk_type           = "CLOUD_PREMIUM"
    system_disk_size           = 50
    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
    internet_max_bandwidth_out = 10
    subnet_id                  = "subnet-xxxxx"
  }
}

# TKE Cluster Endpoint (Public access)
resource "cloud_tke_kubernetes_cluster_endpoint" "endpoint" {
  cluster_id                 = cloud_tke_kubernetes_cluster.cluster.id
  cluster_internet           = true
  cluster_internet_security_group = "sg-xxxxx"
}

# TKE Cluster Namespace
resource "cloud_tke_kubernetes_cluster_namespace" "namespace" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-namespace"
  
  labels = {
    env = "test"
  }
}

# TKE Cluster Secret
resource "cloud_tke_kubernetes_cluster_secret" "secret" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-secret"
  namespace  = "default"
  type       = "Opaque"
  
  data = {
    username = base64encode("admin")
    password = base64encode("password123")
  }
}

# TKE Cluster PV (Persistent Volume)
resource "cloud_tke_kubernetes_cluster_pv" "pv" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-pv"
  
  spec {
    capacity = {
      storage = "10Gi"
    }
    access_modes = ["ReadWriteOnce"]
    storage_class_name = "cbs"
    
    persistent_volume_source {
      cbs {
        volume_id = "disk-xxxxx"
        fs_type   = "ext4"
      }
    }
  }
}

# TKE Cluster PVC (Persistent Volume Claim)
resource "cloud_tke_kubernetes_cluster_pvc" "pvc" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-pvc"
  namespace  = "default"
  
  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "10Gi"
      }
    }
    storage_class_name = "cbs"
  }
}

# TKE Cluster Deploy (Deployment)
resource "cloud_tke_kubernetes_cluster_deploy" "deploy" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-deploy"
  namespace  = "default"
  
  replicas = 3
  
  selector {
    match_labels = {
      app = "nginx"
    }
  }
  
  template {
    metadata {
      labels = {
        app = "nginx"
      }
    }
    
    spec {
      containers {
        name  = "nginx"
        image = "nginx:latest"
        
        ports {
          container_port = 80
        }
      }
    }
  }
}

# TKE Cluster Affinity
resource "cloud_tke_kubernetes_cluster_affinity" "affinity" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  namespace  = "default"
  name       = "example-affinity"
  
  affinity {
    node_affinity {
      required_during_scheduling_ignored_during_execution {
        node_selector_terms {
          match_expressions {
            key      = "role"
            operator = "In"
            values   = ["worker"]
          }
        }
      }
    }
  }
}

# TKE Cluster Ingress
resource "cloud_tke_kubernetes_cluster_ing" "ingress" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  name       = "example-ingress"
  namespace  = "default"
  
  rules {
    host = "example.com"
    http {
      paths {
        path = "/"
        backend {
          service_name = "example-service"
          service_port = 80
        }
      }
    }
  }
}

# TKE Cluster Plugin
resource "cloud_tke_kubernetes_cluster_plugin" "plugin" {
  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
  plugin     = "cos"
  version    = "1.0.0"
}
