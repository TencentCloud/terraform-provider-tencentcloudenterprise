# Cloud Load Balancer (CLB) Examples

# ========== Data Sources ==========

# Query CLB instances
data "cloud_clb_instances" "instances" {
  clb_id = "lb-xxxxx"
}

# Query CLB listeners
data "cloud_clb_listeners" "listeners" {
  clb_id      = "lb-xxxxx"
  listener_id = "lbl-xxxxx"
}

# Query CLB listener rules
data "cloud_clb_listener_rules" "rules" {
  clb_id      = "lb-xxxxx"
  listener_id = "lbl-xxxxx"
  rule_id     = "rule-xxxxx"
}

# Query CLB attachments
data "cloud_clb_attachments" "attachments" {
  clb_id      = "lb-xxxxx"
  listener_id = "lbl-xxxxx"
}

# Query CLB redirections
data "cloud_clb_redirections" "redirections" {
  clb_id             = "lb-xxxxx"
  source_listener_id = "lbl-xxxxx"
  source_rule_id     = "rule-xxxxx"
}

# Query CLB certificates
data "cloud_clb_certificates" "certs" {
  certificate_id = "cert-xxxxx"
}

# Query CLB instance by cert ID
data "cloud_clb_instance_by_cert_id" "by_cert" {
  certificate_ids = ["cert-xxxxx"]
}

# Query CLB instance detail
data "cloud_clb_instance_detail" "detail" {
  target_type = "CVM"
}

# Query CLB resources
data "cloud_clb_resources" "resources" {}

# Query CLB target health
data "cloud_clb_target_health" "health" {
  load_balancer_ids = ["lb-xxxxx"]
}

# ========== Resources ==========

# CLB Instance
resource "cloud_clb_instance" "instance" {
  clb_name              = "example-clb"
  network_type          = "OPEN"
  clb_type              = "OPEN"
  vpc_id                = "vpc-xxxxx"
  project_id            = 0
  security_groups       = ["sg-xxxxx"]
  
  tags = {
    env = "test"
  }
}

# CLB Listener (Layer 4 - TCP)
resource "cloud_clb_listener" "tcp_listener" {
  clb_id               = cloud_clb_instance.instance.id
  listener_name        = "tcp-listener"
  port                 = 80
  protocol             = "TCP"
  health_check_switch  = true
  health_check_type    = "TCP"
  health_check_interval = 5
  health_check_timeout  = 2
  scheduler            = "WRR"
}

# CLB Listener (Layer 7 - HTTP)
resource "cloud_clb_listener" "http_listener" {
  clb_id        = cloud_clb_instance.instance.id
  listener_name = "http-listener"
  port          = 80
  protocol      = "HTTP"
  scheduler     = "WRR"
}

# CLB Listener (Layer 7 - HTTPS)
resource "cloud_clb_listener" "https_listener" {
  clb_id           = cloud_clb_instance.instance.id
  listener_name    = "https-listener"
  port             = 443
  protocol         = "HTTPS"
  certificate_id   = "cert-xxxxx"
  scheduler        = "WRR"
}

# CLB Listener Rule (for HTTP/HTTPS)
resource "cloud_clb_listener_rule" "rule" {
  clb_id              = cloud_clb_instance.instance.id
  listener_id         = cloud_clb_listener.http_listener.listener_id
  domain              = "example.com"
  url                 = "/api"
  scheduler           = "WRR"
  health_check_switch = true
  
  health_check_http_code = 200
  health_check_path      = "/health"
  health_check_method    = "GET"
}

# CLB Attachment (Backend servers)
resource "cloud_clb_attachment" "attachment" {
  clb_id      = cloud_clb_instance.instance.id
  listener_id = cloud_clb_listener.tcp_listener.listener_id
  
  targets {
    instance_id = "ins-xxxxx"
    port        = 80
    weight      = 10
  }
}

# CLB Redirection
resource "cloud_clb_redirection" "redirect" {
  clb_id             = cloud_clb_instance.instance.id
  source_listener_id = cloud_clb_listener.http_listener.listener_id
  target_listener_id = cloud_clb_listener.https_listener.listener_id
  source_rule_id     = cloud_clb_listener_rule.rule.rule_id
  is_auto_rewrite    = true
}

# CLB Customized Config
resource "cloud_clb_customized_config" "config" {
  config_name    = "example-config"
  config_content = "client_max_body_size 20m;"
}

# CLB Security Group Attachment
resource "cloud_clb_security_group_attachment" "sg_attachment" {
  clb_id             = cloud_clb_instance.instance.id
  security_group_ids = ["sg-xxxxx"]
}

# CLB Certificates
resource "cloud_clb_certificates" "cert" {
  certificate_name    = "example-cert"
  certificate_type    = "SERVER"
  certificate_content = file("./server.crt")
  certificate_key     = file("./server.key")
}

# CLB Replace Certificate
resource "cloud_clb_replace_cert" "replace" {
  old_certificate_id = "cert-xxxxx"
  certificate {
    cert_id = cloud_clb_certificates.cert.id
  }
}
