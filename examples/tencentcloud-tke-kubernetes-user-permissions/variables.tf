variable "cluster_id" {
  description = "ID of the existing TKE cluster used for validation."
  type        = string

  validation {
    condition     = startswith(var.cluster_id, "cls-")
    error_message = "cluster_id must start with \"cls-\"."
  }
}

variable "target_uin" {
  description = "UIN of the dedicated test user whose TKE RBAC permissions will be managed."
  type        = string

  validation {
    condition     = length(trimspace(var.target_uin)) > 0
    error_message = "target_uin must not be empty."
  }
}

variable "permissions" {
  description = "Complete desired permission set for the test user."
  type = list(object({
    role_name = string
    role_type = string
    namespace = optional(string)
  }))

  default = [
    {
      role_name = "tke:ro"
      role_type = "cluster"
    }
  ]

  validation {
    condition = alltrue([
      for permission in var.permissions :
      contains(["cluster", "namespace"], permission.role_type)
    ])
    error_message = "Each role_type must be either \"cluster\" or \"namespace\"."
  }

  validation {
    condition = alltrue([
      for permission in var.permissions :
      permission.role_type != "namespace" || try(length(trimspace(permission.namespace)) > 0, false)
    ])
    error_message = "namespace must be set when role_type is \"namespace\"."
  }
}
