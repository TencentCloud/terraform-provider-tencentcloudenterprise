# Key Management Service (KMS) Examples

# ========== Data Sources ==========

# Query KMS keys
data "tencentcloudenterprise_kms_keys" "keys" {
  key_id = "key-xxxxx"
}

# ========== Resources ==========

# KMS Key
resource "tencentcloudenterprise_kms_key" "key" {
  alias                = "example-key"
  description          = "Example KMS key"
  key_rotation_enabled = true
  is_enabled           = true
  
  tags = {
    env = "test"
  }
}

# KMS External Key (BYOK - Bring Your Own Key)
resource "tencentcloudenterprise_kms_external_key" "external" {
  alias                = "example-external-key"
  description          = "Example external KMS key"
  key_material_base64  = "base64-encoded-key-material"
  valid_to             = 1735660800  # Unix timestamp
  is_enabled           = true
  
  tags = {
    env = "production"
  }
}
