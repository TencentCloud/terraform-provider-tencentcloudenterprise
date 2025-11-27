/*
Use this data source to query detailed information of CBS storages in parallel.

# Example Usage

```hcl

	data "tencentcloudenterprise_cbs_storages_set" "storages" {
	  availability_zone = "ap-guangzhou-3"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cbs_storages_set", CNDescription{
		TerraformTypeCN: "云硬盘列表",
		DescriptionCN:   "提供云硬盘列表数据源，支持并行查询CBS云硬盘的详细信息。",
		AttributesCN: map[string]string{
			"storage_id":             "要查询的CBS的ID",
			"storage_name":           "要查询的CBS的名称",
			"availability_zone":      "CBS实例所在的可用区域",
			"project_id":             "与CBS关联的项目的ID",
			"storage_type":           "按云磁盘介质类型筛选（`cloud_BASIC`：HDD云磁盘|`cloud_PREMIUM`：高级云存储|`cloud_SSD`：SSD云磁盘）",
			"storage_usage":          "按云磁盘类型（`SYSTEM_disk`：系统磁盘|`DATA_disk'：数据磁盘）筛选",
			"charge_type":            "按磁盘费用类型（`POSTPAID_by_HOUR`|`PREPAID`）列出筛选器",
			"portable":               "根据磁盘是否可移植进行筛选（布尔值“true”或“false”）",
			"storage_state":          "按磁盘状态列出筛选器（`UNATTACHED`|`ATTACHING` |`ATTACHED`|`DETACHIG`|`EXPANDING`|`ROLLBACKING`|`TORECYCLE`）",
			"instance_ips":           "按附加的实例公用或专用IP列出筛选器",
			"instance_name":          "按附加的实例名称列出筛选器",
			"tag_keys":               "按标记键列出筛选器",
			"tag_按标记值列出筛选器s":         "按标记值列出筛选器",
			"result_output_file":     "用于保存结果",
			"storage_list":           "存储列表每个元素都包含以下属性：",
			"storage_size":           "CBS的音量",
			"attached":               "指示CBS是否安装在CVM上",
			"instance_id":            "此CBS装载的CVM实例的ID",
			"encrypt":                "指示CBS是否加密",
			"create_time":            "CBS的创建时间",
			"status":                 "CBS的状态",
			"tags":                   "此CBS中的可用标签",
			"prepaid_renew_flag":     "CBS实例在达到预付租赁结束时自动续订或不续订的方式",
			"throughput_performance": "为数据磁盘添加额外的性能仅当磁盘类型为“CLOUD_TSSD”或“CLOUD_HSSD”时才起作用",
		},
	})
}

func dataSourceTencentCloudCbsStoragesSet() *schema.Resource {
	return &schema.Resource{
		Description: "Query detailed information of CBS storages in parallel.",
		Read:        dataSourceTencentCloudCbsStoragesSetRead,

		Schema: map[string]*schema.Schema{
			"storage_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the CBS to be queried.",
			},
			"storage_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the CBS to be queried.",
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone that the CBS instance locates at.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the project with which the CBS is associated.",
			},
			"storage_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CBS_STORAGE_TYPE),
				Description:  "Filter by cloud disk media type (`CLOUD_BASIC`: HDD cloud disk | `CLOUD_PREMIUM`: Premium Cloud Storage | `CLOUD_SSD`: SSD cloud disk).",
			},
			"storage_usage": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by cloud disk type (`SYSTEM_DISK`: system disk | `DATA_DISK`: data disk).",
			},
			"charge_type": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by disk charge type (`POSTPAID_BY_HOUR` | `PREPAID`).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"portable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter by whether the disk is portable (Boolean `true` or `false`).",
			},
			"storage_state": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by disk state (`UNATTACHED` | `ATTACHING` | `ATTACHED` | `DETACHING` | `EXPANDING` | `ROLLBACKING` | `TORECYCLE`).",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"instance_ips": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by attached instance public or private IPs.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"instance_name": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by attached instance name.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tag_keys": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by tag keys.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"tag_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List filter by tag values.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"storage_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of storage. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"storage_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of CBS.",
						},
						"storage_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of CBS.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Types of storage medium.",
						},
						"storage_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Types of CBS.",
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone of CBS.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the project.",
						},
						"storage_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Volume of CBS.",
						},
						"attached": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the CBS is mounted the CVM.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CVM instance that be mounted by this CBS.",
						},
						"encrypt": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether CBS is encrypted.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of CBS.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of CBS.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The available tags within this CBS.",
						},
						"prepaid_renew_flag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The way that CBS instance will be renew automatically or not when it reach the end of the prepaid tenancy.",
						},
						"charge_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pay type of the CBS instance.",
						},
						"throughput_performance": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Add extra performance to the data disk. Only works when disk type is `CLOUD_TSSD` or `CLOUD_HSSD`.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCbsStoragesSetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cbs_storages.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	if v, ok := d.GetOk("storage_id"); ok {
		params["disk-id"] = v.(string)
	}
	if v, ok := d.GetOk("storage_name"); ok {
		params["disk-name"] = v.(string)
	}
	if v, ok := d.GetOk("availability_zone"); ok {
		params["zone"] = v.(string)
	}
	if v, ok := d.GetOkExists("project_id"); ok {
		params["project-id"] = fmt.Sprintf("%d", v.(int))
	}
	if v, ok := d.GetOk("storage_type"); ok {
		params["disk-type"] = v.(string)
	}
	if v, ok := d.GetOk("storage_usage"); ok {
		params["disk-usage"] = v.(string)
	}

	if v, ok := d.GetOk("charge_type"); ok {
		params["disk-charge-type"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}

	if v, ok := d.GetOk("portable"); ok {
		if v.(bool) {
			params["portable"] = "TRUE"
		} else {
			params["portable"] = "FALSE"
		}
	}

	if v, ok := d.GetOk("storage_state"); ok {
		params["disk-state"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	if v, ok := d.GetOk("instance_ips"); ok {
		params["instance-ip-address"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	if v, ok := d.GetOk("instance_name"); ok {
		params["instance-name"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	if v, ok := d.GetOk("tag_keys"); ok {
		params["tag-key"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}
	if v, ok := d.GetOk("tag_values"); ok {
		params["tag-value"] = helper.InterfacesStringsPoint(v.([]interface{}))
	}

	cbsService := CbsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	storages, e := cbsService.DescribeDisksInParallelByFilter(ctx, params)
	if e != nil {
		return e
	}
	ids := make([]string, 0, len(storages))
	storageList := make([]map[string]interface{}, 0, len(storages))
	for _, storage := range storages {
		mapping := map[string]interface{}{
			"storage_id":             storage.DiskId,
			"storage_name":           storage.DiskName,
			"storage_usage":          storage.DiskUsage,
			"storage_type":           storage.DiskType,
			"availability_zone":      storage.Placement.Zone,
			"project_id":             storage.Placement.ProjectId,
			"storage_size":           storage.DiskSize,
			"attached":               storage.Attached,
			"instance_id":            storage.InstanceId,
			"encrypt":                storage.Encrypt,
			"create_time":            storage.CreateTime,
			"status":                 storage.DiskState,
			"prepaid_renew_flag":     storage.RenewFlag,
			"charge_type":            storage.DiskChargeType,
			"throughput_performance": storage.ThroughputPerformance,
		}
		if storage.Tags != nil {
			tags := make(map[string]interface{}, len(storage.Tags))
			for _, t := range storage.Tags {
				tags[*t.Key] = *t.Value
			}
			mapping["tags"] = tags
		}
		storageList = append(storageList, mapping)
		ids = append(ids, *storage.DiskId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e = d.Set("storage_list", storageList); e != nil {
		log.Printf("[CRITAL]%s provider set storage list fail, reason:%s\n ", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), storageList); e != nil {
			return e
		}
	}

	return nil

}
