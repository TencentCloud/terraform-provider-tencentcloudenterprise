/*
Provides a resource to create a cls machine group.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_machine_group" "group" {
	  group_name        = "group"
	  service_logging   = true
	  tags              = {
	    "test" = "test1"
	  }
	  update_end_time   = "19:05:40"
	  update_start_time = "17:05:40"

	  machine_group_type {
	    type   = "ip"
	    values = [
	      "203.0.113.101",
	      "203.0.113.102",
	    ]
	  }
	}

```

# Import

cls machine group can be imported using the id, e.g.

```
$ terraform import cloud_cls_machine_group.group caf168e7-32cd-4ac6-bf89-1950a760e09c
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_machine_group", CNDescription{
		TerraformTypeCN: "CLS机器组",
		DescriptionCN:   "提供CLS机器组资源，用于创建和管理日志服务机器组。",
		AttributesCN: map[string]string{
			"group_name":         "机器组名字，不能重复",
			"flag":               "TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE。",
			"delay_cleanup_time": "机器组中机器离线清理时间。单位：天",
			"meta_tags":          "机器组元数据信息列表",
			"os_type":            "系统类型，取值如下：  0：Linux （默认值） 1：Windows",
			"service_logging":    "是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费。默认false",
			"auto_update":        "是否开启机器组自动更新。true表示开启机器组自动更新，false表示关闭机器组自动更新。是否开启机器组自动更新。默认false",
			"update_start_time":  "升级开始时间，建议业务低峰期升级LogListener",
			"update_end_time":    "升级结束时间，建议业务低峰期升级LogListener",
			"machine_group_type": "创建机器组类型。取值如下：  Type：ip，Values中为ip字符串列表创建机器组 Type：label，Values中为标签字符串列表创建机器组",
			"tags":               "标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。",
			"group_id":           "机器组ID",
			"key":                "元数据key",
			"value":              "元数据value",
			"type":               "机器组类型。支持 ip 和 label。  ip：表示该机器组Values中存的是采集机器的ip地址 label：表示该机器组Values中存储的是机器的标签",
			"values":             "机器描述列表。",
		},
	})
}

func resourceTencentCloudClsMachineGroup() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsMachineGroupCreate,
		Read:        resourceTencentCloudClsMachineGroupRead,
		Delete:      resourceTencentCloudClsMachineGroupDelete,
		Update:      resourceTencentCloudClsMachineGroupUpdate,
		Description: "Provides a resource to create and manage CLS machine group",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Machine group name, which must be unique.",
			},
			"machine_group_type": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Type of the machine group to be created. `ip`: machine group created with a list of IP strings in `Values`. `label`: machine group created with a list of label strings in `Values`.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Machine group type. Supported values: `ip` and `label`. `ip` means the `Values` of this machine group stores the IP addresses of the machines to be collected. `label` means the `Values` of this machine group stores the machine labels.",
						},
						"values": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "Machine description list.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list. By specifying this parameter, tags can be bound to the corresponding machine group at the same time. A maximum of 10 tag key-value pairs are supported, and the same resource can only be bound to the same tag key.",
			},
			"auto_update": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable automatic update for the machine group. `true` means to enable automatic update, `false` means to disable. Default is `false`.",
			},
			"update_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Update start time. It is recommended to upgrade LogListener during off-peak hours.",
			},
			"update_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Update end time. It is recommended to upgrade LogListener during off-peak hours.",
			},
			"service_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable service logging, which is used to record logs generated by the LogListener service itself. After it is enabled, the internal log set `cls_service_logging` and log topics `loglistener_status`, `loglistener_alarm`, and `loglistener_business` will be created, which will not incur fees. Default is `false`.",
			},
			"meta_tags": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of metadata tags for the machine group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Metadata key.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Metadata value.",
						},
					},
				},
			},
			"flag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TKE flag. An empty string indicates that the logs are not from TKE, and `label_k8s` indicates that the logs are from TKE.",
			},
			"delay_cleanup_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Offline cleaning time for machines in the machine group. Unit: days.",
			},
			"os_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Operating system type. 0: Linux (default), 1: Windows.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Machine group id.",
			},
		},
	}
}

func resourceTencentCloudClsMachineGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_machine_group.create")()

	logId := getLogId(contextNil)

	var (
		request  = cls.NewCreateMachineGroupRequest()
		response *cls.CreateMachineGroupResponse
	)

	if v, ok := d.GetOk("group_name"); ok {
		request.GroupName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("machine_group_type"); ok {
		machineGroupTypes := make([]*cls.MachineGroupTypeInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			machineGroupType := cls.MachineGroupTypeInfo{}
			if v, ok := dMap["type"]; ok {
				machineGroupType.Type = helper.String(v.(string))
			}
			if v, ok := dMap["values"]; ok {
				for _, u := range v.([]interface{}) {
					machineGroupType.Values = append(machineGroupType.Values, helper.String(u.(string)))
				}
			}
			machineGroupTypes = append(machineGroupTypes, &machineGroupType)
		}
		request.MachineGroupType = machineGroupTypes[0]
	}

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		for k, v := range tags {
			key := k
			value := v
			request.Tags = append(request.Tags, &cls.Tag{
				Key:   &key,
				Value: &value,
			})
		}
	}

	if v, ok := d.GetOk("auto_update"); ok {
		request.AutoUpdate = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("update_start_time"); ok {
		request.UpdateStartTime = helper.String(v.(string))
	}

	if v, ok := d.GetOk("update_end_time"); ok {
		request.UpdateEndTime = helper.String(v.(string))
	}

	if v, ok := d.GetOk("service_logging"); ok {
		request.ServiceLogging = helper.Bool(v.(bool))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateMachineGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cls machine group failed, reason:%+v", logId, err)
		return err
	}

	id := *response.Response.GroupId
	d.SetId(id)
	return resourceTencentCloudClsMachineGroupRead(d, meta)
}

func resourceTencentCloudClsMachineGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_machine_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	machineGroup, err := service.DescribeClsMachineGroupById(ctx, id)

	if err != nil {
		return err
	}

	if machineGroup == nil {
		d.SetId("")
		return fmt.Errorf("resource `MachineGroup` %s does not exist", id)
	}

	_ = d.Set("group_name", machineGroup.GroupName)

	machineGroupType := make([]map[string]interface{}, 0)

	machineGroupType = append(machineGroupType, map[string]interface{}{
		"type":   *machineGroup.MachineGroupType.Type,
		"values": machineGroup.MachineGroupType.Values,
	})

	_ = d.Set("machine_group_type", machineGroupType)

	tags := make(map[string]string, len(machineGroup.Tags))
	for _, tag := range machineGroup.Tags {
		tags[*tag.Key] = *tag.Value
	}
	_ = d.Set("tags", tags)

	_ = d.Set("auto_update", helper.StrToBool(*machineGroup.AutoUpdate))
	_ = d.Set("update_start_time", machineGroup.UpdateStartTime)
	_ = d.Set("update_end_time", machineGroup.UpdateEndTime)
	_ = d.Set("service_logging", machineGroup.ServiceLogging)

	return nil
}

func resourceTencentCloudClsMachineGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_machine_group.update")()
	logId := getLogId(contextNil)
	request := cls.NewModifyMachineGroupRequest()

	request.GroupId = helper.String(d.Id())
	request.GroupName = helper.String(d.Get("group_name").(string))

	if d.HasChange("machine_group_type") {
		if v, ok := d.GetOk("machine_group_type"); ok {
			machineGroupTypes := make([]*cls.MachineGroupTypeInfo, 0, 10)
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				machineGroupType := cls.MachineGroupTypeInfo{}
				if v, ok := dMap["type"]; ok {
					machineGroupType.Type = helper.String(v.(string))
				}
				if v, ok := dMap["values"]; ok {
					for _, u := range v.([]interface{}) {
						machineGroupType.Values = append(machineGroupType.Values, helper.String(u.(string)))
					}
					machineGroupTypes = append(machineGroupTypes, &machineGroupType)
				}
			}
			request.MachineGroupType = machineGroupTypes[0]
		}
	}

	if d.HasChange("tags") {
		tags := d.Get("tags").(map[string]interface{})
		request.Tags = make([]*cls.Tag, 0, len(tags))
		for k, v := range tags {
			key := k
			value := v
			request.Tags = append(request.Tags, &cls.Tag{
				Key:   &key,
				Value: helper.String(value.(string)),
			})
		}
	}

	if d.HasChange("auto_update") || d.HasChange("update_start_time") || d.HasChange("update_end_time") {
		request.AutoUpdate = helper.Bool(d.Get("auto_update").(bool))
		request.UpdateStartTime = helper.String(d.Get("update_start_time").(string))
		request.UpdateEndTime = helper.String(d.Get("update_end_time").(string))
	}

	if d.HasChange("service_logging") {
		request.ServiceLogging = helper.Bool(d.Get("service_logging").(bool))
	}

	if d.HasChange("delay_cleanup_time") {
		request.DelayCleanupTime = helper.IntInt64(d.Get("delay_cleanup_time").(int))
	}

	if d.HasChange("meta_tags") {
		if v, ok := d.GetOk("meta_tags"); ok {
			metaTags := make([]*cls.MetaTagInfo, 0, 10)
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				metaTag := cls.MetaTagInfo{}
				if v, ok := dMap["key"]; ok {
					metaTag.Key = helper.String(v.(string))
				}
				if v, ok := dMap["value"]; ok {
					metaTag.Value = helper.String(v.(string))
				}
				metaTags = append(metaTags, &metaTag)
			}
			request.MetaTags = metaTags
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyMachineGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudClsMachineGroupRead(d, meta)
}

func resourceTencentCloudClsMachineGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_machine_group.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsMachineGroup(ctx, id); err != nil {
		return err
	}

	return nil
}
