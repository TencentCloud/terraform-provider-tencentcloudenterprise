---
subcategory: "Cloud Redis®(Redis)"
layout: "cloud"
page_title: "TencentCloud: cloud_redis_instance"
sidebar_current: "docs-cloud-resource-redis_instance"
description: |-
  Provides a resource to create a instance and set its attributes.
---

# cloud_redis_instance

Provides a resource to create a instance and set its attributes.

~> **NOTE:** The argument vpc_id and subnet_id is now required because Basic Network Instance is no longer supported.

~> **NOTE:** Both adding and removing replications in one change is supported but not recommend.

## Example Usage

```hcl
data "cloud_redis_zone_config" "zone" {
}

resource "cloud_redis_instance" "redis_instance_test_2" {
  availability_zone  = data.cloud_redis_zone_config.zone.list[0].zone
  type_id            = data.cloud_redis_zone_config.zone.list[0].type_id
  password           = "test12345789"
  mem_size           = 8192
  redis_shard_num    = data.cloud_redis_zone_config.zone.list[0].redis_shard_nums[0]
  redis_replicas_num = data.cloud_redis_zone_config.zone.list[0].redis_replicas_nums[0]
  name               = "terrform_test"
  port               = 6379
}
```

### Using multi replica zone set

```hcl
data "cloud_availability_zones" "az" {

}

variable "redis_replicas_num" {
  default = 3
}

resource "cloud_redis_instance" "red1" {
  availability_zone  = data.cloud_availability_zones.az.zones[0].name
  charge_type        = "POSTPAID"
  mem_size           = 1024
  name               = "test-redis"
  port               = 6379
  project_id         = 0
  redis_replicas_num = var.redis_replicas_num
  redis_shard_num    = 1
  security_groups = [
    "sg-d765yoec",
  ]
  subnet_id = "subnet-ie01x91v"
  type_id   = 6
  vpc_id    = "vpc-k4lrsafc"
  password  = "a12121312334"

  replica_zone_ids = [
    for i in range(var.redis_replicas_num)
  : data.cloud_availability_zones.az.zones[i % length(data.cloud_availability_zones.az.zones)].id]
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone ID of an instance to be created, please refer to `cloud_redis_zone_config.list`.
* `mem_size` - (Required, Int) The memory volume of an available instance(in MB), please refer to `cloud_redis_zone_config.list[zone].shard_memories`. When redis is standard type, it represents total memory size of the instance; when Redis is cluster type, it represents memory size of per sharding.
* `security_groups` - (Required, Set: [`String`]) ID of security group. If both vpc_id and subnet_id are not set, this argument should not be set either.
* `subnet_id` - (Required, String) Specifies which subnet the instance should belong to. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
* `type_id` - (Required, Int, ForceNew) Instance type. Available values reference data source `cloud_redis_zone_config` or [document](https://intl.cloud.tencent.com/document/product/239/32069), toggle immediately when modified.
* `vpc_id` - (Required, String) ID of the vpc with which the instance is to be associated. When the `operation_network` is `changeVpc` or `changeBaseToVpc`, this parameter needs to be configured.
* `auto_renew_flag` - (Optional, Int, ForceNew) Auto-renew flag. 0 - default state (manual renewal); 1 - automatic renewal; 2 - explicit no automatic renewal.
* `charge_type` - (Optional, String, ForceNew) The charge type of instance. Valid values: `PREPAID` and `POSTPAID`. Default value is `POSTPAID`. Note: TencentCloud International only supports `POSTPAID`. Caution that update operation on this field will delete old instances and create new with new charge type.
* `force_delete` - (Optional, Bool) Indicate whether to delete Redis instance directly or not. Default is false. If set true, the instance will be deleted instead of staying recycle bin. Note: only works for `PREPAID` instance.
* `ip` - (Optional, String) IP address of an instance. When the `operation_network` is `changeVip`, this parameter needs to be configured.
* `name` - (Optional, String) Instance name.
* `no_auth` - (Optional, Bool) Indicates whether the redis instance support no-auth access. NOTE: Only available in private cloud environment.
* `operation_network` - (Optional, String) Refers to the category of the pre-modified network, including: `changeVip`: refers to switching the private network, including its intranet IPv4 address and port; `changeVpc`: refers to switching the subnet to which the private network belongs; `changeBaseToVpc`: refers to switching the basic network to a private network; `changeVPort`: refers to only modifying the instance network port.
* `password` - (Optional, String) Password for a Redis user, which should be 8 to 16 characters. NOTE: Only `no_auth=true` specified can make password empty.
* `platform_project_id` - (Optional, String, ForceNew) The platform project ID is used to specify which project this resource belongs to.
* `port` - (Optional, Int) The port used to access a redis instance. The default value is 6379. When the `operation_network` is `changeVPort` or `changeVip`, this parameter needs to be configured.
* `prepaid_period` - (Optional, Int) The tenancy (time unit is month) of the prepaid instance, NOTE: it only works when charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
* `redis_replicas_num` - (Optional, Int) The number of instance copies. This is not required for standalone and master slave versions and must equal to count of `replica_zone_ids`.
* `redis_shard_num` - (Optional, Int) The number of instance shard, default is 1. This is not required for standalone and master slave versions.
* `replica_zone_ids` - (Optional, List: [`Int`]) ID of replica nodes available zone. This is not required for standalone and master slave versions. NOTE: Removing some of the same zone of replicas (e.g. removing 100001 of [100001, 100001, 100002]) will pick the first hit to remove.
* `replicas_read_only` - (Optional, Bool) Whether copy read-only is supported, Redis 2.8 Standard Edition and CKV Standard Edition do not support replica read-only, turn on replica read-only, the instance will automatically read and write separate, write requests are routed to the primary node, read requests are routed to the replica node, if you need to open replica read-only, the recommended number of replicas >=2.
* `resource_pool_id` - (Optional, String, ForceNew) The ID of the resource pool to which the instance belongs.
* `tags` - (Optional, Map) Instance tags.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - The time when the instance was created.
* `status` - Current status of an instance, maybe: init, processing, online, isolate and todelete.


## Import

Redis instance can be imported, e.g.

```
$ terraform import cloud_redis_instance.redislab redis-id
```

