---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_alarm_notice"
sidebar_current: "docs-tencentcloudenterprise-resource-cls_alarm_notice"
description: |-
  Provides a resource to create a CLS alarm notice.
---

# tencentcloudenterprise_cls_alarm_notice

Provides a resource to create a CLS alarm notice.

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice" {
  name = "terraform-alarm-notice"
  type = "All"

  notice_receivers {
    receiver_type     = "Uin"
    receiver_ids      = [10000001]
    receiver_channels = ["Email"]
    notice_content_id = "Default-zh"
    start_time        = "00:00:00"
    end_time          = "23:59:59"
  }

  web_callbacks {
    callback_type = "Http"
    url           = "https://example.com/callback"
    method        = "POST"
  }

  tags = {
    "createdBy" = "terraform"
  }
}
```

### # Advanced Mode Example

```hcl
resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice_rule" {
  name = "terraform-alarm-notice-rule"

  notice_rules {
    rule = jsonencode({
      Value = "AND"
      Type  = "Operation"
      Children = [
        {
          Type  = "Condition"
          Value = "NotifyType"
          Children = [
            { Value = "In", Type = "Compare" },
            { Value = "[1,2]", Type = "Value" }
          ]
        }
      ]
    })

    notice_receivers {
      receiver_type     = "Uin"
      receiver_ids      = [4364, 3563]
      receiver_channels = ["Sms", "WeChat"]
      start_time        = "00:00:00"
      end_time          = "23:59:59"
      notice_content_id = "Default-zh"
    }

    escalate = false
    interval = 10
    type     = 1
  }

  deliver_status      = 1
  alarm_shield_status = 1
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Alarm notice name.
* `alarm_shield_status` - (Optional, Int) Alarm shield status. 1: Off; 2: On.
* `deliver_config` - (Optional, List) Log delivery configuration.
* `deliver_status` - (Optional, Int) Log delivery switch. 1: Off; 2: On.
* `jump_domain` - (Optional, String) Jump domain for query links.
* `notice_receivers` - (Optional, List) Notice receivers.
* `notice_rules` - (Optional, List) Advanced mode notice rules.
* `tags` - (Optional, Map) Tag description list.
* `type` - (Optional, String) Notice type. Valid values: Trigger, Recovery, All.
* `web_callbacks` - (Optional, List) Callback information.

The `deliver_config` object supports the following:

* `region` - (Optional, String) Delivery region.
* `scope` - (Optional, Int) Delivery scope. 0: All logs; 1: Only alarm trigger and recovery logs.
* `topic_id` - (Optional, String) Delivery topic ID.

The `notice_receivers` object supports the following:

* `receiver_channels` - (Required, Set) Receiver channels. Valid values: Email, Sms, WeChat, Phone.
* `receiver_ids` - (Required, Set) Receiver ID list.
* `receiver_type` - (Required, String) Receiver type. Valid values: Uin, Group.
* `end_time` - (Optional, String) End time allowed to receive messages.
* `index` - (Optional, Int) Index. Input is invalid, output is valid.
* `notice_content_id` - (Optional, String) Notice content template ID.
* `start_time` - (Optional, String) Start time allowed to receive messages.

The `notice_rules` object supports the following:

* `rule` - (Required, String) Rule JSON string.
* `escalate` - (Optional, Bool) Whether to enable escalation.
* `interval` - (Optional, Int) Escalation interval in minutes.
* `notice_receivers` - (Optional, List) Notice receivers.
* `type` - (Optional, Int) Escalation condition type.
* `web_callbacks` - (Optional, List) Callback information.

The `web_callbacks` object supports the following:

* `callback_type` - (Required, String) Callback type. Valid values: Http, WeCom, DingTalk, Lark.
* `url` - (Required, String) Callback URL.
* `body` - (Optional, String) Request body.
* `headers` - (Optional, Set) Request headers.
* `index` - (Optional, Int) Index. Input is invalid, output is valid.
* `method` - (Optional, String) Method. Valid values: POST, PUT.
* `mobiles` - (Optional, Set) Telephone list.
* `notice_content_id` - (Optional, String) Notice content template ID.
* `remind_type` - (Optional, Int) Remind type. 0: Do not remind; 1: Specified person; 2: Everyone.
* `user_ids` - (Optional, Set) User ID list.
* `web_callback_id` - (Optional, String) Integration configuration ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `alarm_notice_id` - Alarm notice ID.
* `deliver_err_msg` - Log delivery error message.
* `deliver_flag` - Log delivery status flag.

## Import

tencentcloudenterprise_cls_alarm_notice can be imported using the id, e.g.

```
cls alarm notice can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_alarm_notice.alarm_notice alarm_notice_id
```
```

