/*
Use this data source to query detailed acl information of Ckafka

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_acls" "foo" {
	  instance_id   = "ckafka-f9ife4zz"
	  resource_type = "TOPIC"
	  resource_name = "topic-tf-test"
	  host          = "2"
	}

```
*/
package tencentcloud

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_acls", CNDescription{
		TerraformTypeCN: "ckafka acl列表",
		DescriptionCN:   "提供Ckafka ACL数据源，用于查询Ckafka ACL的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "ckafka实例的Id",
			"resource_type":      "Acl资源类型，(0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID)，当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用",
			"resource_name":      "资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称",
			"host":               "默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持",
			"result_output_file": "结果输出文件",
			"acl_list":           "ACL列表",
			"operation_type":     "Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)",
			"permission_type":    "权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)",
			"principal":          "用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户",
		},
	})
}

func dataSourceTencentCloudCkafkaAcls() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed acl information of Ckafka",
		Read:        dataSourceTencentCloudCkafkaAclsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the ckafka instance.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ACL resource type. Valid values are `UNKNOWN`, `ANY`, `TOPIC`, `GROUP`, `CLUSTER`, `TRANSACTIONAL_ID`. Currently, only `TOPIC` is available, and other fields will be used for future ACLs compatible with open-source Kafka.",
			},
			"resource_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ACL resource name, which is related to `resource_type`. For example, if `resource_type` is `TOPIC`, this field indicates the topic name; if `resource_type` is `GROUP`, this field indicates the group name.",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host substr used for querying.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"acl_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of ckafka acls. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ACL resource type.",
						},
						"resource_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ACL resource name, which is related to `resource_type`.",
						},
						"operation_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ACL operation mode.",
						},
						"permission_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ACL permission type, valid values are `UNKNOWN`, `ANY`, `DENY`, `ALLOW`, and `ALLOW` by default. Currently, CKafka supports `ALLOW` (equivalent to allow list), and other fields will be used for future ACLs compatible with open-source Kafka.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address allowed to access.",
						},
						"principal": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User which can access. `*` means that any user can access.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCkafkaAclsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_acls.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	params["instance_id"] = d.Get("instance_id").(string)
	params["resource_type"] = d.Get("resource_type").(string)
	params["resource_name"] = d.Get("resource_name").(string)
	if v, ok := d.GetOk("host"); ok {
		params["host"] = v.(string)
	}

	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	aclInfos, err := ckafkaService.DescribeAclByFilter(ctx, params)
	if err != nil {
		return err
	}
	aclList := make([]map[string]interface{}, 0, len(aclInfos))
	ids := make([]string, 0, len(aclInfos))
	for _, acl := range aclInfos {
		aclList = append(aclList, map[string]interface{}{
			"resource_type":   CKAFKA_ACL_RESOURCE_TYPE_TO_STRING[*acl.ResourceType],
			"resource_name":   *acl.ResourceName,
			"operation_type":  CKAFKA_ACL_OPERATION_TO_STRING[*acl.Operation],
			"permission_type": CKAFKA_PERMISSION_TYPE_TO_STRING[*acl.PermissionType],
			"host":            *acl.Host,
			"principal":       strings.TrimLeft(*acl.Principal, CKAFKA_ACL_PRINCIPAL_STR),
		})

		ids = append(ids, params["instance_id"].(string)+FILED_SP+CKAFKA_PERMISSION_TYPE_TO_STRING[*acl.PermissionType]+
			FILED_SP+strings.TrimLeft(*acl.Principal, CKAFKA_ACL_PRINCIPAL_STR)+FILED_SP+*acl.Host+FILED_SP+
			CKAFKA_ACL_OPERATION_TO_STRING[*acl.Operation]+FILED_SP+CKAFKA_ACL_RESOURCE_TYPE_TO_STRING[*acl.ResourceType]+
			FILED_SP+*acl.ResourceName)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("acl_list", aclList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), aclList); e != nil {
			return e
		}
	}

	return nil
}
