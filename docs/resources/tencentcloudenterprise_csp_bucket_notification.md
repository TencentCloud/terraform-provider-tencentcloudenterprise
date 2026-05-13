---
subcategory: "Cloud Storage Platform(CSP)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_csp_bucket_notification"
sidebar_current: "docs-tencentcloudenterprise-resource-csp_bucket_notification"
description: |-
  Provides a CSP resource to manage COS bucket event notification configuration.
---

# tencentcloudenterprise_csp_bucket_notification

Provides a CSP resource to manage COS bucket event notification configuration.

Event notifications allow you to receive messages when certain events happen in your bucket,
such as object creation or deletion. Notifications are delivered to CKafka.

## Example Usage

### Without SASL authentication:

```hcl
resource "tencentcloudenterprise_csp_bucket_notification" "example" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "rule-1"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
    topic              = "my-notification-topic"
  }
}
```

### With SASL authentication:

```hcl
resource "tencentcloudenterprise_csp_bucket_notification" "example_sasl" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "rule-sasl"
    events             = ["cos:ObjectCreated:*", "cos:ObjectRemove:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
    topic              = "my-notification-topic"
    sasl_user          = "123123"
    sasl_password      = "Tencent@321"
  }
}
```

### With resource prefix/suffix filter:

```hcl
resource "tencentcloudenterprise_csp_bucket_notification" "example_filter" {
  bucket = "est123-1255000115"

  notification_rule {
    id                 = "rule-filter"
    events             = ["cos:ObjectCreated:*"]
    ckafka_instance_id = "ckafka-7k3pve8e"
    topic              = "my-notification-topic"
    filter_prefix      = "logs/"
    filter_suffix      = ".json"
  }
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String, ForceNew) The name of the bucket. Bucket format should be [custom name]-[appid], for example `mybucket-1258798060`.
* `notification_rule` - (Required, List) A list of notification rules for the bucket.

The `notification_rule` object supports the following:

* `ckafka_instance_id` - (Required, String) The CKafka instance ID to deliver notifications to. The provider will automatically resolve the Kafka endpoint from this instance.
* `events` - (Required, List) List of event types that trigger the notification. Valid values include: `cos:ObjectCreated:*`, `cos:ObjectCreated:Put`, `cos:ObjectCreated:Copy`, `cos:ObjectCreated:Post`, `cos:ObjectCreated:CompleteMultipartUpload`, `cos:ObjectRemove:*`, `cos:ObjectRemove:Delete`, `cos:ObjectRemove:DeleteMarkerCreated`.
* `id` - (Required, String) Unique identifier for the notification rule.
* `topic` - (Required, String) The Kafka topic name to deliver notification messages to.
* `filter_prefix` - (Optional, String) Object key prefix for filtering notifications. Only objects matching this prefix will trigger notifications. For example `adc` means only objects under the `adc/` path.
* `filter_suffix` - (Optional, String) Object key suffix for filtering notifications. Only objects matching this suffix will trigger notifications. For example `.jpg`.
* `sasl_password` - (Optional, String) SASL password for Kafka authentication.
* `sasl_user` - (Optional, String) SASL username for Kafka authentication.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_csp_bucket_notification can be imported using the id, e.g.

```
CSP bucket notification can be imported using the bucket name, e.g.

```
$ terraform import tencentcloudenterprise_csp_bucket_notification.example mybucket-1258798060
```
```

