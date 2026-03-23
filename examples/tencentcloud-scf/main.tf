resource "tencentcloudenterprise_cos_bucket" "foo" {
  bucket = "scf-cos1-1258798060"
  acl    = "public-read"
}

resource "tencentcloudenterprise_cos_bucket" "bar" {
  bucket = "scf-cos2-1258798060"
  acl    = "public-read"
}

resource "tencentcloudenterprise_cos_bucket_object" "myobject" {
  bucket = tencentcloudenterprise_cos_bucket.foo.bucket
  key    = "/new_object_key.zip"
  source = "code.zip"
  acl    = "public-read"
}

resource "tencentcloudenterprise_cam_role" "foo" {
  name          = "ci-scf-role"
  document      = var.role_document
  description   = "ci-scf-role"
  console_login = true
}

resource "tencentcloudenterprise_scf_namespace" "foo" {
  namespace   = "ci-test-scf"
  description = "test1"
}

resource "tencentcloudenterprise_vpc" "test_vpc" {
  name       = "Used for testing subnets"
  cidr_block = "10.1.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "test_subnet" {
  vpc_id            = tencentcloudenterprise_vpc.test_vpc.id
  name              = "terraform test subnet"
  cidr_block        = "10.1.1.0/24"
  availability_zone = var.availability_zone
}

resource "tencentcloudenterprise_scf_function" "foo" {
  name        = "ci-test-function"
  description = "test"
  handler     = "main.do_it"
  runtime     = "Python3.6"
  namespace   = tencentcloudenterprise_scf_namespace.foo.id
  role        = tencentcloudenterprise_cam_role.foo.id
  vpc_id      = tencentcloudenterprise_vpc.test_vpc.id
  subnet_id   = tencentcloudenterprise_vpc_subnet.test_subnet.id

  cos_bucket_name   = tencentcloudenterprise_cos_bucket.foo.id
  cos_object_name   = tencentcloudenterprise_cos_bucket_object.myobject.key
  cos_bucket_region = "ap-guangzhou"

  triggers {
    name         = "ci-test-fn-api-gw"
    type         = "timer"
    trigger_desc = "*/5 * * * * * *"
  }

  triggers {
    name         = tencentcloudenterprise_cos_bucket.bar.id
    type         = "cos"
    trigger_desc = var.trigger_desc
  }

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_scf_functions" "foo" {
  name        = tencentcloudenterprise_scf_function.foo.name
  description = tencentcloudenterprise_scf_function.foo.description
  namespace   = tencentcloudenterprise_scf_function.foo.namespace
  tags        = tencentcloudenterprise_scf_function.foo.tags
}

data "tencentcloudenterprise_scf_namespaces" "foo" {
  namespace   = tencentcloudenterprise_scf_namespace.foo.id
  description = tencentcloudenterprise_scf_namespace.foo.description
}

data "tencentcloudenterprise_scf_logs" "foo" {
  function_name = tencentcloudenterprise_scf_function.foo.name
  namespace     = tencentcloudenterprise_scf_function.foo.namespace
  offset        = 0
  limit         = 100
  order         = "desc"
  order_by      = "duration"
  ret_code      = "UserCodeException"
  start_time    = "2017-05-16 20:00:00"
  end_time      = "2017-05-17 20:00:00"
}
