provider "tencentcloud" {
  region = "ap-guangzhou"
}

resource "cloud_ssm_secret" "foo" {
  secret_name = "test"
  description = "test secret"
  recovery_window_in_days = 0
  is_enabled = true

  tags = {
    test-tag = "test"
  }
}

resource "cloud_ssm_secret_version" "v1" {
  secret_name = cloud_ssm_secret.foo.secret_name
  version_id = "v1"
  secret_binary = "MTIzMTIzMTIzMTIzMTIzQQ=="
}

data "cloud_ssm_secrets" "secret_list" {
  secret_name = cloud_ssm_secret.foo.secret_name
  order_type = 1
  state = 1
}

data "cloud_ssm_secret_versions" "secret_version_list" {
  secret_name = cloud_ssm_secret_version.v1.secret_name
  version_id = cloud_ssm_secret_version.v1.version_id
}