# Cloud Storage Platform (CSP) Examples
# CSP is enterprise-specific COS service

# ========== Data Sources ==========

# Query CSP buckets
data "tencentcloudenterprise_csp_buckets" "buckets" {
  bucket_prefix = "example"
}

# Query CSP bucket object
data "tencentcloudenterprise_csp_bucket_object" "object" {
  bucket = "example-bucket-123456"
  key    = "path/to/object.txt"
}

# ========== Resources ==========

# CSP Bucket
resource "tencentcloudenterprise_csp_bucket" "bucket" {
  bucket = "example-csp-bucket-123456"
  acl    = "private"
  
  tags = {
    env = "production"
  }
}

# CSP Bucket Object
resource "tencentcloudenterprise_csp_bucket_object" "object" {
  bucket  = cloud_csp_bucket.bucket.bucket
  key     = "example/data.json"
  content = jsonencode({
    message = "Hello from CSP"
  })
  acl = "private"
}

# CSP Bucket Policy
resource "tencentcloudenterprise_csp_bucket_policy" "policy" {
  bucket = cloud_csp_bucket.bucket.bucket
  policy = jsonencode({
    version = "2.0"
    statement = [
      {
        effect    = "Allow"
        principal = {
          qcs = ["qcs::cam::uin/123456:uin/123456"]
        }
        action = [
          "name/cos:PutObject",
          "name/cos:GetObject"
        ]
        resource = [
          "qcs::cos:ap-guangzhou:uid/123456:${cloud_csp_bucket.bucket.bucket}/*"
        ]
      }
    ]
  })
}
