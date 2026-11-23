package main

// docTPL is the template for generating documentation
const docTPL = `---
subcategory: "{{.product}}"
layout: "{{.cloud_mark}}"
page_title: "{{.cloud_title}}: {{.name}}"
sidebar_current: "docs-{{.cloud_mark}}-{{.dtype_singular}}-{{.resource}}"
description: |-
  {{.description_short}}
---

# {{.name}}

{{.description}}

## Example Usage

{{.example}}

## Argument Reference

The following arguments are supported:

{{.arguments}}

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

{{.attributes}}
{{if .import}}
## Import

{{.cloud_mark}}_{{.resource}} can be imported using the id, e.g.

` + "```" + `
{{.import}}
` + "```" + `
{{end}}
`

// idxTPL is the template for generating index
const idxTPL = `---
layout: "{{.cloud_mark}}"
page_title: "Provider: {{.cloud_mark}}"
sidebar_current: "docs-{{.cloud_mark}}-index"
description: The {{.cloud_title}} provider is used to interact with many resources supported by {{.cloud_title}}. The provider needs to be configured with the proper credentials before it can be used.
---

# {{.cloud_title}} Provider

The {{.cloud_title}} (TCE) provider is used to interact with many resources supported by [{{.cloud_title}}](https://www.tencentcloud.com/solutions/tce).
This provider is implemented as ` + "`terraform-provider-{{.cloud_mark}}`" + ` and requires proper credential configuration before use.

Use the navigation on the left to read about the available resources.
> ℹ️**Note**: This provider is still under development.

## Example Usage

` + "``` terraform" + `
terraform {
  required_providers {
    {{.cloud_mark}} = {
      source = "TencentCloud/{{.cloud_mark}}"
    }
  }
}

# Configure the TCE Provider
provider "{{.cloud_mark}}" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
  region     = "ap-guangzhou"
  domain     = "api3.mydomain.com"
}

# Get availability images
data "{{.cloudPrefix}}cvm_images" "default" {
  image_type = ["PUBLIC_IMAGE"]
  os_name    = "centos"
}

# Create a web server
resource "{{.cloudPrefix}}cvm_instance" "web" {
  instance_name              = "web server"
  availability_zone          = "ap-guangzhou-1"
  image_id                   = data.{{.cloudPrefix}}cvm_images.default.images.0.image_id
  instance_type              = "S5.MEDIUM2"
  system_disk_type           = "CLOUD_PREMIUM"
  system_disk_size           = 50
  vpc_id                     = "vpc-xxxxxxxx"
  subnet_id                  = "subnet-xxxxxxxx"
}
` + "```" + `

## Authentication

The TCE provider offers a flexible means of providing credentials for authentication.
The following methods are supported and explained below:


### Static credentials

> ⚠️**Warning:** Hard-coding credentials into any Terraform configuration is not
recommended, and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding ` + "`secret_id`" + `, ` + "`secret_key`" + `, ` + "`region`" + `, and ` + "`domain`" + ` in-line in the {{.cloud_mark}} provider block:

Usage:

` + "```hcl" + `
provider "{{.cloud_mark}}" {
  secret_id  = "my-secret-id"
  secret_key = "my-secret-key"
  region     = "ap-guangzhou"
  domain     = "api3.mydomain.com"
  csp_domain = "csp.mydomain.com"    # Optional: Custom CSP domain
  cos_domain = "cospub.mydomain.com" # Optional: Custom COS domain
}
` + "```" + `

### Environment variables

You can provide your credentials via ` + "`TENCENTCLOUD_SECRET_ID`" + ` and ` + "`TENCENTCLOUD_SECRET_KEY`" + ` environment variables,
representing your TencentCloud Secret Id and Secret Key respectively. ` + "`TENCENTCLOUD_REGION`" + ` and ` + "`TENCENTCLOUD_DOMAIN`" + ` are also used, if applicable:

` + "```hcl" + `
provider "{{.cloud_mark}}" {

}
` + "```" + `

Usage:

` + "```shell" + `
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
$ export TENCENTCLOUD_REGION="ap-guangzhou"
$ export TENCENTCLOUD_DOMAIN="api3.mydomain.com"
$ export TENCENTCLOUD_PROTOCOL="HTTPS"               # Optional: HTTP or HTTPS
$ export TENCENTCLOUD_CSP_DOMAIN="csp.mydomain.com"  # Optional: Custom CSP domain
$ export TENCENTCLOUD_COS_DOMAIN="cospub.mydomain.com" # Optional: Custom COS domain
$ terraform plan
` + "```" + `

## Argument Reference

In addition to generic provider arguments (e.g. alias and version), the following arguments are supported in the TencentCloud provider block:

* ` + "`secret_id`" + ` - (Optional) This is the TencentCloud secret id. It must be provided, but it can also be sourced from the ` + "`TENCENTCLOUD_SECRET_ID`" + ` environment variable.
* ` + "`secret_key`" + ` - (Optional) This is the TencentCloud secret key. It must be provided, but it can also be sourced from the ` + "`TENCENTCLOUD_SECRET_KEY`" + ` environment variable.
* ` + "`region`" + ` - (Optional) This is the TencentCloud region. It must be provided, but it can also be sourced from the ` + "`TENCENTCLOUD_REGION`" + ` environment variable.
* ` + "`protocol`" + ` - (Optional) The protocol of the API request. Valid values: ` + "`HTTP`" + ` and ` + "`HTTPS`" + `. Default is ` + "`HTTP`" + `.
* ` + "`domain`" + ` - (Optional) The root domain of the API request. Can also be sourced from the ` + "`TENCENTCLOUD_DOMAIN`" + ` environment variable.
* ` + "`csp_domain`" + ` - (Optional) The CSP domain for the API request. Can also be sourced from the ` + "`TENCENTCLOUD_CSP_DOMAIN`" + ` environment variable.
* ` + "`cos_domain`" + ` - (Optional) The COS domain for the API request. Can also be sourced from the ` + "`TENCENTCLOUD_COS_DOMAIN`" + ` environment variable.`
