# Cloud Object Storage (COS) Examples

# ========== Data Sources ==========

# Query COS buckets
data "cloud_cos_buckets" "buckets" {
  bucket_prefix = "example"
}

# Query COS bucket object
data "cloud_cos_bucket_object" "object" {
  bucket = "example-bucket-123456"
  key    = "path/to/object.txt"
}

# ========== Resources ==========

# COS Bucket
resource "cloud_cos_bucket" "bucket" {
  bucket = "example-bucket-123456"
  acl    = "private"
  
  versioning {
    enabled = true
  }
  
  lifecycle_rules {
    filter_prefix = "logs/"
    
    transition {
      days          = 30
      storage_class = "STANDARD_IA"
    }
    
    expiration {
      days = 90
    }
  }
  
  cors_rules {
    allowed_origins = ["*"]
    allowed_methods = ["GET", "POST"]
    allowed_headers = ["*"]
    max_age_seconds = 3600
  }
  
  tags = {
    env = "test"
  }
}

# COS Bucket Object
resource "cloud_cos_bucket_object" "object" {
  bucket  = cloud_cos_bucket.bucket.bucket
  key     = "example/object.txt"
  content = "Hello, COS!"
  acl     = "private"
}

# COS Bucket Policy
resource "cloud_cos_bucket_policy" "policy" {
  bucket = cloud_cos_bucket.bucket.bucket
  policy = jsonencode({
    version = "2.0"
    statement = [
      {
        effect    = "Allow"
        principal = {
          qcs = ["qcs::cam::uin/123456:uin/123456"]
        }
        action = [
          "name/cos:GetObject"
        ]
        resource = [
          "qcs::cos:ap-guangzhou:uid/123456:${cloud_cos_bucket.bucket.bucket}/*"
        ]
      }
    ]
  })
}
