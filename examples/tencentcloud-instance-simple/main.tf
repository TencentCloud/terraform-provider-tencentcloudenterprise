resource "tencentcloudenterprise_cvm_instance" "my-server" {
  image_id          = var.image_id
  availability_zone = var.availability_zone
}

output "instance_id" {
  value = cloud_cvm_instance.my-server.id
}
