terraform {
  required_version = ">= 1.5.0"

  required_providers {
    tencentcloudenterprise = {
      source = "TencentCloud/tencentcloudenterprise"
    }
  }
}

provider "tencentcloudenterprise" {}

resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "validation" {
  cluster_id = var.cluster_id
  target_uin = var.target_uin

  dynamic "permissions" {
    for_each = var.permissions

    content {
      role_name = permissions.value.role_name
      role_type = permissions.value.role_type
      namespace = try(permissions.value.namespace, null)
    }
  }
}

output "resource_id" {
  value = tencentcloudenterprise_tke_kubernetes_user_permissions.validation.id
}

output "permissions" {
  value = tencentcloudenterprise_tke_kubernetes_user_permissions.validation.permissions
}
