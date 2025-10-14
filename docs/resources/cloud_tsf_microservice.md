---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_microservice"
sidebar_current: "docs-cloud-resources-tsf_microservice"
description: |-
  Provides a resource to create a tsf microservice
---

# cloud_tsf_microservice

Provides a resource to create a tsf microservice

## Example Usage

```hcl
resource "cloud_tsf_microservice" "microservice" {
  namespace_id      = "namespace-vjlkzkgy"
  microservice_name = "test-microservice"
  microservice_desc = "desc-microservice"
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `microservice_name` - (Required, String) Microservice name.
* `namespace_id` - (Required, String) Namespace ID.
* `microservice_desc` - (Optional, String) Microservice description information.
* `tags` - (Optional, Map) Tag description list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tsf microservice can be imported using the namespaceId#microserviceId, e.g.

```
terraform import cloud_tsf_microservice.microservice namespace-vjlkzkgy#ms-vjeb43lw
```

