---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_notice_content"
sidebar_current: "docs-tencentcloudenterprise-resource-cls_notice_content"
description: |-
  Provides a resource to create a CLS notice content template.
---

# tencentcloudenterprise_cls_notice_content

Provides a resource to create a CLS notice content template.

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_notice_content" "notice_content" {
  name = "terraform-notice-content-test"
  type = 0

  notice_contents {
    type = "Email"

    trigger_content {
      title   = "Trigger title"
      content = "Trigger content"
    }

    recovery_content {
      title   = "Recovery title"
      content = "Recovery content"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Notice content template name.
* `notice_contents` - (Optional, List) Notice content template details.
* `type` - (Optional, Int) Notice content language. 0: Chinese; 1: English.

The `notice_contents` object supports the following:

* `type` - (Required, String) Channel type. Valid values: Email, Sms, WeChat, Phone, WeCom, DingTalk, Lark, Http.
* `recovery_content` - (Optional, List) Alarm recovery notice content.
* `trigger_content` - (Optional, List) Alarm trigger notice content.

The `recovery_content` object supports the following:

* `content` - (Optional, String) Notice content body.
* `headers` - (Optional, Set) Request headers, supported only for Http callbacks.
* `title` - (Optional, String) Notice content title.

The `trigger_content` object supports the following:

* `content` - (Optional, String) Notice content body.
* `headers` - (Optional, Set) Request headers, supported only for Http callbacks.
* `title` - (Optional, String) Notice content title.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `notice_content_id` - Notice content template ID.

## Import

tencentcloudenterprise_cls_notice_content can be imported using the id, e.g.

```
cls notice_content can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_notice_content.notice_content notice_content_id
```
```

