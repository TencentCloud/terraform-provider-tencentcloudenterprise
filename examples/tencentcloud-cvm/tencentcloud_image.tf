resource "tencentcloudenterprise_cvm_image" "image_snap" {
  image_name   		= "image-snapshot-keep"
  instance_id = "ins-j3zyyzda"
  force_poweroff 		= true
  image_description 	= "create image with snapshot"
}