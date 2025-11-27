# Secret Manager (SSM) Examples

# ========== Data Sources ==========

# Query SSM secrets
data "tencentcloudenterprise_ssm_secrets" "secrets" {
  secret_name = "example-secret"
}

# Query SSM secret versions
data "tencentcloudenterprise_ssm_secret_versions" "versions" {
  secret_name = "example-secret"
}

# ========== Resources ==========

# SSM Secret
resource "tencentcloudenterprise_ssm_secret" "secret" {
  secret_name = "example-secret"
  description = "Example secret"
  
  tags = {
    env = "test"
  }
}

# SSM Secret Version
resource "tencentcloudenterprise_ssm_secret_version" "version" {
  secret_name   = cloud_ssm_secret.secret.secret_name
  version_id    = "v1"
  secret_string = jsonencode({
    username = "admin"
    password = "password123"
  })
}
