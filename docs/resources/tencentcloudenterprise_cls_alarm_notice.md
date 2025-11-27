---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_alarm_notice"
sidebar_current: "docs-tencentcloudenterprise-resources-cls_alarm_notice"
description: |-
  Provides a resource to create a cls alarm_notice
---

# tencentcloudenterprise_cls_alarm_notice

Provides a resource to create a cls alarm_notice

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice" {
  name = "terraform-alarm-notice-test"
  tags = {
    "createdBy" = "terraform"
  }
  type = "All"

  notice_receivers {
    index = 0
    receiver_channels = [
      "Sms",
    ]
    receiver_ids = [
      13478043,
      15972111,
    ]
    receiver_type = "Uin"
    start_time    = "00:00:00"
    end_time      = "23:59:59"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Alarm notice name.
* `type` - (Required, String) Notice type.
* `notice_receivers` - (Optional, List) Notice receivers.
* `tags` - (Optional, Map) Tag description list.
* `web_callbacks` - (Optional, List) Callback info.

The `notice_receivers` object supports the following:

* `receiver_channels` - (Required, Set) Receiver channels, Email,Sms,WeChat or Phone.
* `receiver_ids` - (Required, Set) Receiver id.
* `receiver_type` - (Required, String) Receiver type, Uin or Group.
* `end_time` - (Optional, String) End time allowed to receive messages.
* `index` - (Optional, Int) Index.
* `start_time` - (Optional, String) Start time allowed to receive messages.

The `web_callbacks` object supports the following:

* `callback_type` - (Required, String) Callback type, WeCom or Http.
* `url` - (Required, String) Callback url.
* `body` - (Optional, String) Abandoned.
* `headers` - (Optional, Set) Abandoned.
* `index` - (Optional, Int) Index.
* `method` - (Optional, String) Method, POST or PUT.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls alarm_notice can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_alarm_notice.alarm_notice alarm_notice_id
```

