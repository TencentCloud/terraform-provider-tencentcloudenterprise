locals {
  app_id = data.tencentcloudenterprise_user_info.info.app_id
  uin = data.tencentcloudenterprise_user_info.info.uin
  owner_uin = data.tencentcloudenterprise_user_info.info.owner_uin
}

data "tencentcloudenterprise_user_info" "info" {}