---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_ckafka_consumer"
sidebar_current: "docs-cloud-resources-cls_ckafka_consumer"
description: |-
  Provides a resource to create a cls ckafka_consumer
---

# cloud_cls_ckafka_consumer

Provides a resource to create a cls ckafka_consumer

## Example Usage

```hcl
resource "cloud_cls_ckafka_consumer" "ckafka_consumer" {
  compression  = 1
  need_content = true
  topic_id     = "7e34a3a7-635e-4da8-9005-88106c1fde69"

  ckafka {
    instance_id   = "ckafka-qzoeaqx8"
    instance_name = "ckafka-instance"
    topic_id      = "topic-c6tm4kpm"
    topic_name    = "name"
    vip           = "203.0.113.23"
    vport         = "9092"
  }

  content {
    enable_tag = true
    meta_fields = [
      "__FILENAME__",
      "__HOSTNAME__",
      "__PKGID__",
      "__SOURCE__",
      "__TIMESTAMP__",
    ]
    tag_json_not_tiled = true
    timestamp_accuracy = 2
  }
}
```

## Argument Reference

The following arguments are supported:

* `ckafka` - (Required, List) CKafka description.
* `topic_id` - (Required, String, ForceNew) ID of the log topic bound to the delivery task.
* `compression` - (Optional, Int) Compression method for delivery. Values: 0 (NONE), 2 (SNAPPY), 3 (LZ4).
* `content` - (Optional, List) Description of metadata information if metadata needs to be delivered.
* `need_content` - (Optional, Bool) Whether to deliver log metadata information. Default is true. When NeedContent is true, the Content field is valid. When NeedContent is false, the Content field is invalid.

The `ckafka` object supports the following:

* `instance_id` - (Required, String) CKafka instance ID.
* `instance_name` - (Required, String) CKafka instance name.
* `topic_id` - (Required, String) CKafka topic ID.
* `topic_name` - (Required, String) CKafka topic name.
* `vip` - (Required, String) CKafka VIP address.
* `vport` - (Required, String) CKafka VPort.

The `content` object supports the following:

* `enable_tag` - (Required, Bool) Whether to deliver TAG information. When EnableTag is true, it means delivering TAG metadata information.
* `meta_fields` - (Required, List) List of metadata to be delivered. Currently only supports: __SOURCE__, __FILENAME__, __TIMESTAMP__, __HOSTNAME__ and __PKGID__.
* `json_type` - (Optional, Int) Delivery JSON format. JsonType 0: consistent with original log, no escaping. JsonType 1: escaped.
* `tag_json_not_tiled` - (Optional, Bool) When EnableTag is true, this field must be filled. TagJsonNotTiled is used to identify whether tag information is JSON flattened. When TagJsonNotTiled is true, it is not flattened. When TagJsonNotTiled is false, it is flattened.
* `timestamp_accuracy` - (Optional, Int) Delivery timestamp precision. Options: [1: second; 2: millisecond]. Default is 1.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls ckafka_consumer can be imported using the id, e.g.

```
terraform import cloud_cls_ckafka_consumer.ckafka_consumer topic_id
```

