---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_saas_domain"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_saas_domain"
description: |-
  Provides a resource to create a NGWAF SaaS domain.
---

# tencentcloudenterprise_ngwaf_saas_domain

Provides a resource to create a NGWAF SaaS domain.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_saas_domain" "example" {
  instance_id = "waf-xxxxxxxx"
  domain      = "example.com"

  ports {
    port              = "80"
    protocol          = "http"
    upstream_port     = "80"
    upstream_protocol = "http"
  }

  upstream_type   = 0
  src_list        = ["1.2.3.4"]
  upstream_scheme = "http"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) Domain names that require defense.
* `instance_id` - (Required, String) Unique ID of Instance.
* `ports` - (Required, Set) This field needs to be set for multiple ports in the upstream server.
* `active_check` - (Optional, Int) Whether to enable active health detection, 0 represents disable and 1 represents enable.
* `api_safe_status` - (Optional, Int) Whether to enable api safe, 1 enable, 0 disable.
* `bot_status` - (Optional, Int) Whether to enable bot, 1 enable, 0 disable.
* `cert_type` - (Optional, Int) Certificate type, 0 represents no certificate, CertType=1 represents self owned certificate, and 2 represents managed certificate.
* `cert` - (Optional, String) Certificate content, When CertType=1, this parameter needs to be filled.
* `cipher_template` - (Optional, Int) Encryption Suite Template, 0:default  1:Universal template 2:Security template 3:Custom template.
* `ciphers` - (Optional, List: [`Int`]) Encryption Suite Information.
* `cls_status` - (Optional, Int) Whether to enable access logs, 1 enable, 0 disable.
* `https_rewrite` - (Optional, Int) Whether redirect to https, 1 will redirect and 0 will not.
* `https_upstream_port` - (Optional, String) Upstream port for https, When listen ports has https port and UpstreamScheme is HTTP, the current field needs to be filled.
* `ip_headers` - (Optional, List: [`String`]) When is_cdn=3, this parameter needs to be filled in to indicate a custom header.
* `is_cdn` - (Optional, Int) Whether a proxy has been enabled before WAF, 0 no deployment, 1 deployment and use first IP in X-Forwarded-For as client IP, 2 deployment and use remote_addr as client IP, 3 deployment and use values of custom headers as client IP.
* `is_http2` - (Optional, Int) Whether enable HTTP2, Enabling HTTP2 requires HTTPS support, 1 means enabled, 0 does not.
* `is_keep_alive` - (Optional, String) Whether to enable keep-alive, 0 disable, 1 enable.
* `is_websocket` - (Optional, Int) Is WebSocket support enabled. 1 means enabled, 0 does not.
* `load_balance` - (Optional, String) Load balancing strategy, where 0 represents polling and 1 represents IP hash and 2 weighted round robin.
* `private_key` - (Optional, String) Certificate key, When CertType=1, this parameter needs to be filled.
* `proxy_read_timeout` - (Optional, Int) 300s.
* `proxy_send_timeout` - (Optional, Int) 300s.
* `sni_host` - (Optional, String) When SniType=3, this parameter needs to be filled in to represent a custom host.
* `sni_type` - (Optional, Int) Sni type fo upstream, 0:disable SNI; 1:enable SNI and SNI equal original request host; 2:and SNI equal upstream host 3:enable SNI and equal customize host.
* `src_list` - (Optional, List: [`String`]) Upstream IP List, When UpstreamType=0, this parameter needs to be filled.
* `ssl_id` - (Optional, String) Certificate ID, When CertType=2, this parameter needs to be filled.
* `status` - (Optional, Int) WAF switch status, 1: turn on WAF switch; 0: turn off WAF switch.
* `tls_version` - (Optional, Int) Version of TLS Protocol.
* `upstream_domain` - (Optional, String) Upstream domain, When UpstreamType=1, this parameter needs to be filled.
* `upstream_scheme` - (Optional, String) Upstream scheme for https, http or https.
* `upstream_type` - (Optional, Int) Upstream type, 0 represents IP, 1 represents domain name.
* `weights` - (Optional, List: [`Int`]) Weight of each upstream.
* `xff_reset` - (Optional, Int) 0:disable xff reset; 1:enable xff reset.

The `ports` object supports the following:

* `port` - (Required, String) Listening port.
* `protocol` - (Required, String) The listening protocol of listening port.
* `upstream_port` - (Required, String) The upstream port for listening port.
* `upstream_protocol` - (Required, String) The upstream protocol for listening port.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `domain_id` - Domain id.

