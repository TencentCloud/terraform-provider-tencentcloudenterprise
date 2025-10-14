---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_cos_shipper"
sidebar_current: "docs-cloud-resources-cls_cos_shipper"
description: |-
  Provides a resource to create a cls cos shipper.
---

# cloud_cls_cos_shipper

Provides a resource to create a cls cos shipper.

## Example Usage

```hcl
resource "cloud_cls_cos_shipper" "shipper" {
  bucket       = "preset-scf-bucket-1308919341"
  interval     = 300
  max_size     = 200
  partition    = "/%Y/%m/%d/%H/"
  prefix       = "ap-guangzhou-fffsasad-1649734752"
  shipper_name = "ap-guangzhou-fffsasad-1649734752"
  topic_id     = "4d07fba0-b93e-4e0b-9a7f-d58542560bbb"

  compress {
    format = "lzop"
  }

  content {
    format = "json"

    json {
      enable_tag = true
      meta_fields = [
        "__FILENAME__",
        "__SOURCE__",
        "__TIMESTAMP__",
      ]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required, String) Destination bucket in the shipping rule to be created.
* `prefix` - (Required, String) Prefix of the shipping directory in the shipping rule to be created.
* `shipper_name` - (Required, String) Shipping rule name.
* `topic_id` - (Required, String) ID of the log topic to which the shipping rule to be created belongs.
* `compress` - (Optional, List) Compression configuration of shipped log.
* `content` - (Optional, List) Format configuration of shipped log content.
* `custom_uin` - (Optional, Int) Cross-account UIN for supporting cross-account bucket shipping.
* `end_time` - (Optional, Int) End time point for the shipping data range, cannot be a future time. If not filled by user, defaults to continuous shipping, i.e., infinite.
* `filename_mode` - (Optional, Int) File naming configuration for shipped logs. 0: random number naming, 1: shipping time naming. Default value: 0 (random number naming).
* `filter_rules` - (Optional, List) Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. If the array is empty, no filtering will be performed, and all logs will be shipped.
* `interval` - (Optional, Int) Shipping time interval in seconds. Default value: 300. Value range: 300~900.
* `max_size` - (Optional, Int) Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.
* `partition` - (Optional, String) Partition rule of shipped log, which can be represented in strftime time format.
* `start_time` - (Optional, Int) Start time point for the shipping data range, cannot exceed the lifecycle start point of the log topic. If not filled by user, defaults to the time when the user creates the shipping task.

The `compress` object supports the following:

* `format` - (Required, String) Compression format. Valid values: gzip, lzop, none (no compression).

The `content` object supports the following:

* `format` - (Required, String) Content format. Valid values: json, csv.
* `csv` - (Optional, List) CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.
* `json` - (Optional, List) JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.
* `parquet` - (Optional, List) 

The `csv` object supports the following:

* `delimiter` - (Required, String) Field delimiter.
* `escape_char` - (Required, String) Field delimiter.
* `keys` - (Required, List) Names of keys.Note: this field may return null, indicating that no valid values can be obtained.
* `non_existing_field` - (Required, String) Content used to populate non-existing fields.
* `print_key` - (Required, Bool) Whether to print key on the first row of the CSV file.

The `filter_rules` object supports the following:

* `key` - (Required, String) Filter rule key.
* `regex` - (Required, String) Filter rule.
* `value` - (Required, String) Filter rule value.

The `json` object supports the following:

* `enable_tag` - (Required, Bool) Enablement flag.
* `meta_fields` - (Required, List) Metadata information list. Note: this field may return null, indicating that no valid values can be obtained..
* `json_type` - (Optional, Int) JSON delivery format. 0: string delivery; 1: structured delivery.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `shipper_id` - Shipping rule ID.


## Import

cls cos shipper can be imported using the id, e.g.

```
$ terraform import cloud_cls_cos_shipper.shipper 5d1b7b2a-c163-4c48-bb01-9ee00584d761
```

