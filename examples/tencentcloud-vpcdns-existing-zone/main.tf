terraform {
  required_providers {
    tencentcloudenterprise = {
      source = "TencentCloud/tencentcloudenterprise"
    }
  }
}

provider "tencentcloudenterprise" {
  region = var.region
}

variable "existing_zone_id" {
  type        = string
  description = "ID of the existing private DNS zone."
}

variable "region" {
  type        = string
  description = "Region of the new VPC."
}

variable "vpc_name" {
  type        = string
  description = "Name of the new VPC."
}

variable "vpc_cidr" {
  type        = string
  description = "IPv4 CIDR block of the new VPC."
}

resource "tencentcloudenterprise_vpc" "new" {
  name       = var.vpc_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_vpcdns_zone_vpc_attachment" "existing_zone" {
  zone_id = var.existing_zone_id

  vpc_set {
    uniq_vpc_id = tencentcloudenterprise_vpc.new.id
    region      = var.region
  }
}
