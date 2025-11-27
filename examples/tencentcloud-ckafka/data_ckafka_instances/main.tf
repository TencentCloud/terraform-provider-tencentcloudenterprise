data "tencentcloudenterprise_ckafka_instances" "get_instances" {
  search_word = "test"
  tag_key = "ckafka"
  status = [1]
  filters  {
    name = "VpcId"
    values = ["1"]
  }

}