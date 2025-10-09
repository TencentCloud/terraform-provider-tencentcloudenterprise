---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application_release_config"
sidebar_current: "docs-cloud-resource-tsf_application_release_config"
description: |-
  Provides a resource to create a tsf application_release_config
---

# cloud_tsf_application_release_config

Provides a resource to create a tsf application_release_config

## Example Usage

```hcl
resource "cloud_tsf_application_release_config" "application_release_config" {
  config_id    = "dcfg-nalqbqwv"
  group_id     = "group-yxmz72gv"
  release_desc = "terraform-test"
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required, String, ForceNew) Configuration ID.
* `group_id` - (Required, String, ForceNew) Deployment group ID.
* `release_desc` - (Optional, String, ForceNew) Release description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `application_id` - Application ID.
* `cluster_id` - Cluster ID.
* `cluster_name` - Cluster name.
* `config_name` - Configuration item name.
* `config_release_id` - Configuration item release ID.
* `config_version` - Configuration item version.
* `group_name` - Deployment group name.
* `namespace_id` - Namespace ID.
* `namespace_name` - Namespace name.
* `release_time` - Release time.


## Import

tsf application_release_config can be imported using the configId#groupId#configReleaseId, e.g.

```
terraform import cloud_tsf_application_release_config.application_release_config dcfg-nalqbqwv#group-yxmz72gv#dcfgr-maeeq2ea
```

