resource "tencentcloudenterprise_tcr_instance" "example" {
  name        = "testacctcrinstance"
  instance_type = "standard"
  delete_bucket = true

  tags ={
	test = "test"
  }
}

resource "tencentcloudenterprise_tcr_token" "example" {
  instance_id = tencentcloudenterprise_tcr_instance.example.id
  description       = "test"
  enable   = false
}

resource "tencentcloudenterprise_tcr_namespace" "example" {
  instance_id = tencentcloudenterprise_tcr_instance.example.id
  name        = "test"
  is_public   = false
}

resource "tencentcloudenterprise_tcr_repository" "example" {
  instance_id = tencentcloudenterprise_tcr_instance.example.id
  namespace_name        = tencentcloudenterprise_tcr_namespace.example.name
  name = "test"
  brief_desc = "example"
  description = "long example"
}

data "tencentcloudenterprise_tcr_instances" "example" {
  name = tencentcloudenterprise_tcr_instance.example.name
}

data "tencentcloudenterprise_tcr_tokens" "example" {
  instance_id = tencentcloudenterprise_tcr_token.example.instance_id
}

data "tencentcloudenterprise_tcr_namespaces" "example" {
  instance_id = tencentcloudenterprise_tcr_namespace.example.instance_id
}

data "tencentcloudenterprise_tcr_repositories" "example" {
  instance_id = tencentcloudenterprise_tcr_repository.example.instance_id
  namespace_name = tencentcloudenterprise_tcr_namespace.example.name
}

# get internal vpc access
resource "tencentcloudenterprise_vpc" "example" {
  name       = "example"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "example" {
  availability_zone = var.availability_zone
  name              = "example"
  vpc_id            = tencentcloudenterprise_vpc.example.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

resource "tencentcloudenterprise_tcr_vpc_attachment" "example" {
  instance_id = tencentcloudenterprise_tcr_instance.example.id
  vpc_id = tencentcloudenterprise_vpc.example.id
  subnet_id = tencentcloudenterprise_vpc_subnet.example.id
}

data "tencentcloudenterprise_tcr_vpc_attachments" "example" {
  instance_id = tencentcloudenterprise_tcr_vpc_attachment.example.instance_id
}