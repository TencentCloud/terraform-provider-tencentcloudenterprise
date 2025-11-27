/*
Provides a resource to create a Ckafka Acl.

# Example Usage

# Ckafka Acl

```hcl

	resource "tencentcloudenterprise_ckafka_acl" "foo" {
	  instance_id     = "ckafka-f9ife4zz"
	  resource_type   = "TOPIC"
	  resource_type_name   = "topic-tf-test"
	  operation_type  = "WRITE"
	  permission_type = "ALLOW"
	  host            = "*"
	  principal       = cloud_ckafka_user.foo.account_name
	}

```

# Import

Ckafka acl can be imported using the instance_id#permission_type#principal#host#operation_type#resource_type#resource_type_name, e.g.

```
$ terraform import cloud_ckafka_acl.foo ckafka-f9ife4zz#ALLOW#test#*#WRITE#TOPIC#topic-tf-test
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ckafka_acl", CNDescription{
		TerraformTypeCN: "Ckafka ACL",
		DescriptionCN:   "提供CKafka ACL资源，用于创建和管理CKafka访问控制列表。",
		AttributesCN: map[string]string{
			"instance_id": "Ckafka实例ID",
			//"resource_type":      "Acl资源类型，(0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID)，当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用",
			"resource_type": "Acl资源类型，当前只有TOPIC",
			//"resource_type_name": "资源类型名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称",
			"resource_type_name": "资源类型名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称",
			//"operation_type":     "Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)",
			"operation_type": "Acl操作方式(3:READ，4:WRITE)",
			//"permission_type":    "权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)",
			"permission_type": "权限类型(2:DENY，3:ALLOW)",
			"host":            "默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持",
			"principal":       "用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户",
		},
	})
}
func resourceTencentCloudCkafkaAcl() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a Ckafka Acl.",
		Create:      resourceTencentCloudCkafkaAclCreate,
		Read:        resourceTencentCloudCkafkaAclRead,
		Delete:      resourceTencentCloudCkafkaAclDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the ckafka instance.",
			},
			"resource_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				//Description: "ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. and `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.",
				Description: "ACL resource type.  `TOPIC` by default. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.",
			},
			"resource_type_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				//Description: "ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.",
				Description: "ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; ",
			},
			"operation_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				//Description: "ACL operation mode. Valid values: `UNKNOWN`, `ANY`, `ALL`, `READ`, `WRITE`, `CREATE`, `DELETE`, `ALTER`, `DESCRIBE`, `CLUSTER_ACTION`, `DESCRIBE_CONFIGS` and `ALTER_CONFIGS`.",
				Description: "ACL operation mode. Valid values: `READ`, `WRITE`",
			},
			"permission_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				//Description: "ACL permission type. Valid values: `UNKNOWN`, `ANY`, `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.",
				Description: "ACL permission type. Valid values: `DENY`, `ALLOW`. and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.",
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "*",
				ForceNew: true,
				ValidateFunc: func(i interface{}, s string) (rs []string, error []error) {
					if i.(string) == "*" {
						return
					}
					rs, error = validateIp(i, s)
					return
				},
				Description: "IP address allowed to access. The default value is `*`, which means that any host can access.",
			},
			"principal": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "*",
				ForceNew:    true,
				Description: "User list. The default value is `*`, which means that any user can access. The current user can only be one included in the user list.",
			},
		},
	}
}

func resourceTencentCloudCkafkaAclCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_acl.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	d.GetRawConfig()
	instanceId := d.Get("instance_id").(string)
	resourceType := d.Get("resource_type").(string)
	resourceName := d.Get("resource_type_name").(string)
	operation := d.Get("operation_type").(string)
	permissionType := d.Get("permission_type").(string)
	host := d.Get("host").(string)
	principal := d.Get("principal").(string)

	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	if err := ckafkaService.CreateAcl(ctx, instanceId, resourceType, resourceName, operation, permissionType, host, principal); err != nil {
		return fmt.Errorf("[CRITAL]%s create ckafka user failed, reason:%+v", logId, err)
	}
	d.SetId(instanceId + FILED_SP + permissionType + FILED_SP + principal + FILED_SP + host + FILED_SP + operation + FILED_SP + resourceType + FILED_SP + resourceName)

	return resourceTencentCloudCkafkaAclRead(d, meta)
}

func resourceTencentCloudCkafkaAclRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_acl.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	id := d.Id()
	info, has, err := ckafkaService.DescribeAclByAclId(ctx, id)
	if err != nil {
		return err
	}
	if !has {
		d.SetId("")
		return nil
	}
	items := strings.Split(id, FILED_SP)
	_ = d.Set("instance_id", items[0])
	_ = d.Set("resource_type", CKAFKA_ACL_RESOURCE_TYPE_TO_STRING[*info.ResourceType])
	_ = d.Set("resource_type_name", info.ResourceName)
	_ = d.Set("operation_type", CKAFKA_ACL_OPERATION_TO_STRING[*info.Operation])
	_ = d.Set("permission_type", CKAFKA_PERMISSION_TYPE_TO_STRING[*info.PermissionType])
	_ = d.Set("host", info.Host)
	_ = d.Set("principal", strings.TrimLeft(*info.Principal, CKAFKA_ACL_PRINCIPAL_STR))

	return nil
}

func resourceTencentCloudCkafkaAclDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_ckafka_user.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	if err := ckafkaService.DeleteAcl(ctx, d.Id()); err != nil {
		return err
	}
	return nil
}
