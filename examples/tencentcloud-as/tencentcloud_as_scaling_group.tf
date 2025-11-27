resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name   = "tf-as-scaling-group"
  configuration_id     = "asc-d8rayj7t"
  max_size             = 1
  min_size             = 0
  vpc_id               = "vpc-cs6ffr73"

  subnet_ids           = ["subnet-38oi34ta"]
  default_cooldown     = 400
  desired_capacity     = 1
  termination_policies = ["OLDEST_INSTANCE"]
  retry_policy         = "INCREMENTAL_INTERVALS"

  forward_balancer_ids {
    load_balancer_id = "lb-na0vh0bq"
    listener_id      = "lbl-04kcx5wu"

    target_attribute {
      port   = 80
      weight = 90
    }
  }
}