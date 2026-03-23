resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "tf-eks-vpc"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "sub" {
  vpc_id            = tencentcloudenterprise_vpc.vpc.id
  name              = "tf-as-subnet"
  cidr_block        = "10.2.11.0/24"
  availability_zone = "ap-guangzhou-3"
}
resource "tencentcloudenterprise_vpc_subnet" "sub2" {
  vpc_id            = tencentcloudenterprise_vpc.vpc.id
  name              = "tf-as-subnet"
  cidr_block        = "10.2.10.0/24"
  availability_zone = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_eks_cluster" "foo" {
  cluster_name = "tf-test-eks"
  k8s_version = "1.18.4"
  vpc_id = tencentcloudenterprise_vpc.vpc.id
  subnet_ids = [
    tencentcloudenterprise_vpc_subnet.sub.id,
    tencentcloudenterprise_vpc_subnet.sub2.id,
  ]
  cluster_desc = "test eks cluster created by terraform"
  service_subnet_id =     tencentcloudenterprise_vpc_subnet.sub.id
  dns_servers {
    domain = "www.example1.com"
    servers = ["1.1.1.1:8080", "1.1.1.1:8081", "1.1.1.1:8082"]
  }
  enable_vpc_core_dns = true
  need_delete_cbs = true
  tags = {
    hello = "world"
  }
}