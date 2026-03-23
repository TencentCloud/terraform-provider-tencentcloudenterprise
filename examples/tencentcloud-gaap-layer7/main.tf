resource tencentcloudenterprise_gaap_proxy "foo" {
  name              = "ci-test-gaap-proxy"
  bandwidth         = 10
  concurrent        = 2
  access_region     = "SouthChina"
  realserver_region = "NorthChina"
}

resource tencentcloudenterprise_gaap_certificate "foo" {
  type    = "SERVER"
  content = var.server_cert
  key     = var.server_key
}

resource tencentcloudenterprise_gaap_certificate "bar" {
  type    = "CLIENT"
  content = var.client_ca
  key     = var.client_ca_key
}

resource tencentcloudenterprise_gaap_certificate "server" {
  type    = "SERVER"
  content = var.server_cert
  key     = var.server_key
}

resource tencentcloudenterprise_gaap_certificate "client" {
  type    = "CLIENT"
  content = var.client_ca
  key     = var.client_ca_key
}

resource tencentcloudenterprise_gaap_certificate "realserver" {
  type    = "REALSERVER"
  content = var.client_ca
  key     = var.client_ca_key
}

resource tencentcloudenterprise_gaap_certificate "basic" {
  type    = "BASIC"
  content = "test:tx2KGdo3zJg/."
}

resource tencentcloudenterprise_gaap_certificate "gaap" {
  type    = "PROXY"
  content = var.server_cert
  key     = var.server_key
}

resource tencentcloudenterprise_gaap_layer7_listener "foo" {
  protocol               = "HTTPS"
  name                   = "ci-test-gaap-l7-listener"
  port                   = 80
  proxy_id               = tencentcloudenterprise_gaap_proxy.foo.id
  certificate_id         = tencentcloudenterprise_gaap_certificate.foo.id
  client_certificate_ids = [tencentcloudenterprise_gaap_certificate.bar.id]
  forward_protocol       = "HTTPS"
  auth_type              = 1
}

resource tencentcloudenterprise_gaap_realserver "foo" {
  domain = "www.qq.com"
  name   = "ci-test-gaap-realserver"
}

resource tencentcloudenterprise_gaap_realserver "bar" {
  domain = "qq.com"
  name   = "ci-test-gaap-realserver"
}

resource tencentcloudenterprise_gaap_http_domain "foo" {
  listener_id            = tencentcloudenterprise_gaap_layer7_listener.foo.id
  domain                 = "www.qq.com"
  certificate_id         = tencentcloudenterprise_gaap_certificate.server.id
  client_certificate_ids = [tencentcloudenterprise_gaap_certificate.client.id]

  realserver_auth               = true
  realserver_certificate_ids    = [tencentcloudenterprise_gaap_certificate.realserver.id]
  realserver_certificate_domain = "qq.com"

  basic_auth    = true
  basic_auth_id = tencentcloudenterprise_gaap_certificate.basic.id

  gaap_auth    = true
  gaap_auth_id = tencentcloudenterprise_gaap_certificate.gaap.id
}

resource tencentcloudenterprise_gaap_http_rule "foo" {
  listener_id     = tencentcloudenterprise_gaap_layer7_listener.foo.id
  domain          = tencentcloudenterprise_gaap_http_domain.foo.domain
  path            = "/"
  realserver_type = "DOMAIN"
  health_check    = false
  forward_host    = "www.qqq.com"

  realservers {
    id   = tencentcloudenterprise_gaap_realserver.foo.id
    ip   = tencentcloudenterprise_gaap_realserver.foo.domain
    port = 80
  }

  realservers {
    id   = tencentcloudenterprise_gaap_realserver.bar.id
    ip   = tencentcloudenterprise_gaap_realserver.bar.domain
    port = 80
  }
}

resource tencentcloudenterprise_gaap_domain_error_page "foo" {
  listener_id    = tencentcloudenterprise_gaap_layer7_listener.foo.id
  domain         = tencentcloudenterprise_gaap_http_domain.foo.domain
  error_codes    = [406, 504]
  new_error_code = 502
  body           = "bad request"
  clear_headers  = ["Content-Length", "X-TEST"]

  set_headers = {
    "X-TEST" = "test"
  }
}

data "tencentcloudenterprise_gaap_http_domains" "foo" {
  listener_id = tencentcloudenterprise_gaap_layer7_listener.foo.id
  domain      = tencentcloudenterprise_gaap_http_domain.foo.domain
}

data tencentcloudenterprise_gaap_http_rules "foo" {
  listener_id  = tencentcloudenterprise_gaap_layer7_listener.foo.id
  path         = tencentcloudenterprise_gaap_http_rule.foo.path
  forward_host = tencentcloudenterprise_gaap_http_rule.foo.forward_host
}

data tencentcloudenterprise_gaap_domain_error_pages "foo" {
  listener_id = tencentcloudenterprise_gaap_domain_error_page.foo.listener_id
  domain      = tencentcloudenterprise_gaap_domain_error_page.foo.domain
  ids         = [tencentcloudenterprise_gaap_domain_error_page.foo.id]
}
