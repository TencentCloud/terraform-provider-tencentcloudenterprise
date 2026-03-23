---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_clb_domain"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_clb_domain"
description: |-
  Provides a resource to create a NGWAF CLB domain.
---

# tencentcloudenterprise_ngwaf_clb_domain

Provides a resource to create a NGWAF CLB domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_clb_domain" "example_with_lb" {
  instance_id                 = "waf-xxxxxxxx"
  domain                      = "example.com"
  status                      = 1
  engine                      = 20
  is_cdn                      = 0
  region                      = "ap-beijing"
  flow_mode                   = 0
  bot_status                  = 0
  alb_type                    = "clb"
  ip_headers                  = []
  tencentcloudenterprise_type = "public"
  note                        = "example domain with load balancer"

  load_balancer_set {
    load_balancer_id   = "lb-xxxxxxxx"
    load_balancer_name = "example-lb-1"
    listener_id        = "lbl-xxxxxxxx-1"
    listener_name      = "example-listener-1"
    vip                = "1.2.3.4"
    vport              = 80
    region             = "ap-beijing"
    protocol           = "http"
    zone               = "ap-beijing-1"
    numerical_vpc_id   = -1
    load_balancer_type = "OPEN"
  }
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) Domain name.
* `instance_id` - (Required, String) Instance unique ID.
* `region` - (Required, String) Regions of LB bound by domain.
* `alb_type` - (Optional, String) Load balancer type: clb, apisix or tsegw, default clb.
* `bot_status` - (Optional, Int) Whether to enable bot, 1 enable, 0 disable.
* `engine` - (Optional, Int) Protection Status: 10: Rule Observation&&AI Off Mode, 11: Rule Observation&&AI Observation Mode, 12: Rule Observation&&AI Interception Mode, 20: Rule Interception&&AI Off Mode, 21: Rule Interception&&AI Observation Mode, 22: Rule Interception&&AI Interception Mode, Default 20.
* `flow_mode` - (Optional, Int) WAF traffic mode, 1 cleaning mode, 0 mirroring mode.
* `ip_headers` - (Optional, List: [`String`]) When is_cdn=3, this parameter needs to be filled in to indicate a custom header.
* `is_cdn` - (Optional, Int) Whether a proxy has been enabled before WAF, 0 no deployment, 1 deployment and use first IP in X-Forwarded-For as client IP, 2 deployment and use remote_addr as client IP, 3 deployment and use values of custom headers as client IP.
* `load_balancer_set` - (Optional, List) List of bound LB.
* `note` - (Optional, String) Domain note information.
* `status` - (Optional, Int) Binding status between waf and LB, 0:not bind, 1:binding.
* `tencentcloudenterprise_type` - (Optional, String) Cloud type: public: public cloud; private: private cloud; hybrid: hybrid cloud.

The `load_balancer_set` object supports the following:

* `listener_id` - (Required, String) Unique ID of listener in LB.
* `listener_name` - (Required, String) Listener name.
* `load_balancer_id` - (Required, String) LoadBalancer unique ID.
* `load_balancer_name` - (Required, String) LoadBalancer name.
* `protocol` - (Required, String) Protocol of listener, http or https.
* `region` - (Required, String) LoadBalancer region.
* `vip` - (Required, String) LoadBalancer IP.
* `vport` - (Required, Int) LoadBalancer port.
* `load_balancer_type` - (Optional, String) Network type for load balancer.
* `numerical_vpc_id` - (Optional, Int) VPCID for load balancer, public network is -1, and internal network is filled in according to actual conditions.
* `zone` - (Optional, String) LoadBalancer zone.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `domain_id` - Domain id.

## Import

tencentcloudenterprise_ngwaf_clb_domain can be imported using the id, e.g.

```
NGWAF CLB domain can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_clb_domain.example_with_lb waf-xxxxxxxx#example.com
```
```

