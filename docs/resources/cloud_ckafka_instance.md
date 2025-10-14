---
subcategory: "Cloud Kafka(CKafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_instance"
sidebar_current: "docs-cloud-resources-ckafka_instance"
description: |-
  Use this resource to create ckafka instance.
---

# cloud_ckafka_instance

Use this resource to create ckafka instance.

~> **NOTE:** It only support create prepaid ckafka instance.

## Example Usage

### Basic Instance

```hcl
variable "vpc_id" {
  default = "vpc-68vi2d3h"
}

variable "subnet_id" {
  default = "subnet-ob6clqwk"
}

variable "product_info_list_standard" {
  type = list(object({
    name  = string
    value = string
  }))
  default = [
    {
      name  = "Region"
      value = "shenzhen"
    },
    {
      name  = "Cluster"
      value = "ap-shenzhen-region-jcctest-ops-1 cluster"
    },
    {
      name  = "AZ"
      value = "ap-shenzhen-region-jcctest-ops-1 & ap-shenzhen-region-jcctest-ops-2"
    },
    {
      name  = "Instance Name"
      value = "Not named"
    },
    {
      name  = "Specs Type"
      value = "Standard Edition"
    },
    {
      name  = "Product Model"
      value = "Basic"
    },
    {
      name  = "Peak Bandwidth"
      value = "40MB/s"
    },
    {
      name  = "Disk Capacity"
      value = "300GB"
    },
    {
      name  = "Message Retention Period"
      value = "24 hours"
    },
    {
      name  = "Network"
      value = "vpc-hqa3zhut"
    },
    {
      name  = "Subnet"
      value = "subnet-dot6weka",
    }
  ]
}

variable "product_info_list_profession" {
  type = list(object({
    name  = string
    value = string
  }))
  default = [
    {
      name  = "Region"
      value = "shenzhen"
    },
    {
      name  = "Multi-AZ"
      value = "ap-shenzhen-region-jcctest-ops-1 & ap-shenzhen-region-jcctest-ops-2"
    },
    {
      name  = "Instance Name"
      value = "Not named"
    },
    {
      name  = "Specs Type"
      value = "Pro Edition"
    },
    {
      name  = "Kafka Version"
      value = "2.4.1"
    },
    {
      name  = "Peak Bandwidth"
      value = "1600MB/s"
    },
    {
      name  = "Disk Capacity"
      value = "10000GB"
    },
    {
      name  = "Topic"
      value = 2000
    },
    {
      name  = "Partition"
      value = 4000
    },
    {
      name  = "Message Retention Period"
      value = "24 hours"
    },
    {
      name  = "Network"
      value = "vpc-hqa3zhut"
    },
    {
      name  = "Subnet"
      value = "subnet-dot6weka"
    }
  ]
}

data "cloud_availability_zones" "gz" {
  name    = "ap-guangzhou-3"
  product = "ckafka"
}

resource "cloud_ckafka_instance" "kafka_instance" {
  instance_name      = "ckafka-instance-type-tf-test"
  zone_id            = data.cloud_availability_zones.gz.zones.0.id
  region_id          = 80000052
  region_name        = "深圳"
  pid                = 1004667
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  msg_retention_time = 1300
  renew_flag         = 0
  kafka_version      = "2.4.1"
  disk_size          = 1000
  disk_type          = "SSD"
  instance_type      = "Basic"
  topic              = 20
  partition          = 4
  dynamic "product_info" {
    for_each = var.product_info_list_profession
    content {
      name  = product_info.value["name"]
      value = product_info.value["value"]
    }
  }
}
```

### Multi zone Instance

```hcl
variable "vpc_id" {
  default = "vpc-68vi2d3h"
}

variable "subnet_id" {
  default = "subnet-ob6clqwk"
}

data "cloud_availability_zones" "gz3" {
  name    = "ap-guangzhou-3"
  product = "ckafka"
}

data "cloud_availability_zones" "gz6" {
  name    = "ap-guangzhou-6"
  product = "ckafka"
}

resource "cloud_ckafka_instance" "kafka_instance" {
  instance_name   = "ckafka-instance-maz-tf-test"
  zone_id         = data.cloud_availability_zones.gz3.zones.0.id
  multi_zone_flag = true
  zone_ids = [
    data.cloud_availability_zones.gz3.zones.0.id,
    data.cloud_availability_zones.gz6.zones.0.id
  ]
  period             = 1
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  msg_retention_time = 1300
  renew_flag         = 0
  kafka_version      = "1.1.1"
  disk_size          = 500
  disk_type          = "CLOUD_BASIC"
  region_id          = 80000052
  region_name        = "深圳"
  pid                = 1
}
```

## Argument Reference

The following arguments are supported:

* `band_width` - (Required, Int, ForceNew) Instance bandwidth in MBps.
* `disk_size` - (Required, Int, ForceNew) Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
* `goods_num` - (Required, Int) Quantity.
* `instance_name` - (Required, String) Instance name.
* `pid` - (Required, Int, ForceNew) Pricing formula ID. 1-9.
* `product_info` - (Required, List) Product information, when 规格类型=标准版:
 When '规格类型' is '标准版','实例名' is optional; when '规格类型' is '专业版','实例名' and '产品型号'
 are optional, all other fields are required:
  - name: 地域, value: 2R3AZ仲裁区集成测试环境北京
  - name: 集群, value: cqyfm7 cluster
  - name: 可用区, value: 重庆云福M7
  - name: 实例名, value: test111
  - name: 规格类型, value: 标准版
  - name: 产品型号, value: 入门型
  - name: 峰值带宽, value: 40MB/s
  - name: 磁盘容量, value: 300GB
  - name: 信息保留时长, value: 72小时
  - name: 网络, value: vpc-kltzarib
  - name: 子网, value: subnet-7qt1q9h6.
* `region_id` - (Required, Int) Region ID.
* `region_name` - (Required, String) Region Name.
* `subnet_id` - (Required, String) Subnet id.
* `vpc_id` - (Required, String) Vpc id.
* `zone_id` - (Required, Int) Available zone id.
* `cluster_id` - (Optional, Int, ForceNew) Cluster-ID represents  the cluster to which Ckafa belongs, associated with cluster_name only needs to select one of them.
* `disk_type` - (Optional, String) Type of disk. [SSD].
* `instance_id` - (Optional, String) Ckafka instance ID.
* `instance_type` - (Optional, String) Description of instance type. [Basic, High Performance].
* `kafka_version` - (Optional, String) Kafka version (2.4.1/2.8.1).
* `msg_retention_time` - (Optional, Int) The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
* `multi_zone_flag` - (Optional, Bool) Indicates whether the instance is multi zones. NOTE: if set to `true`, `zone_ids` must set together.
* `partition` - (Optional, Int) Partition Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
* `renew_flag` - (Optional, Int) Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
* `tag_set` - (Optional, Map) Tag set of instance.
* `topic` - (Optional, Int) Topic Size.
* `zone_ids` - (Optional, Set: [`Int`]) List of available zone id. NOTE: this argument must set together with `multi_zone_flag`.

The `product_info` object supports the following:

* `name` - (Required, String) Item name.
* `value` - (Required, String) Item value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `vip` - Vip of instance.
* `vport` - Type of instance.


## Import

ckafka instance can be imported using the instance_id, e.g.

```
$ terraform import cloud_ckafka_instance.foo ckafka-f9ife4zz
```

