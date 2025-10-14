---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_application"
sidebar_current: "docs-cloud-resources-tsf_application"
description: |-
  Provides a resource to create a tsf application
---

# cloud_tsf_application

Provides a resource to create a tsf application

## Example Usage

```hcl
resource "cloud_tsf_application" "application" {
  application_name         = "terraform-test"
  application_type         = "V"
  microservice_type        = "N"
  application_desc         = "This is my application"
  application_runtime_type = "Java"
  service_config_list {
    name = "my-service"
    ports {
      target_port = 8080
      protocol    = "HTTP"
    }
    health_check {
      path = "/health"
    }
  }
  // ignore_create_image_repository = true
}
```

## Argument Reference

The following arguments are supported:

* `application_name` - (Required, String) Application name.
* `application_type` - (Required, String) Application type: V for virtual machine, C for container.
* `microservice_type` - (Required, String) Application microservice type: M for service mesh, N for normal application, G for gateway application.
* `application_desc` - (Optional, String) Application description.
* `application_log_config` - (Optional, String) Application log configuration, deprecated parameter.
* `application_resource_type` - (Optional, String) Application resource type, deprecated parameter.
* `application_runtime_type` - (Optional, String) Application runtime type.
* `program_id` - (Optional, String) ID of the dataset to be bound.
* `service_config_list` - (Optional, List) List of service configuration information.

The `health_check` object supports the following:

* `path` - (Optional, String) Health check path.

The `ports` object supports the following:

* `protocol` - (Required, String) Port protocol.
* `target_port` - (Required, Int) Service port.

The `service_config_list` object supports the following:

* `name` - (Required, String) Service name.
* `ports` - (Required, List) List of port information.
* `health_check` - (Optional, List) Health check configuration.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



