/*
Use this data source to query the detail information of TurboFS file systems.

# Example Usage

```hcl

	data "tencentcloudenterprise_turbofs_file_systems" "file_systems" {
	  file_system_id    = "turbofs-6hgquxmj"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_turbofs_file_systems", CNDescription{
		TerraformTypeCN: "文件系统信息",
		DescriptionCN:   "提供TurboFS信息数据源，用于查询TurboFS的详细信息。",
		AttributesCN: map[string]string{
			"result_output_file":        "用于保存结果信息，可视化界面不可用",
			"file_system_list":          "文件系统信息列表",
			"tag_name":                  "cfs资源池tag名称",
			"size_limit_max":            "容量限额最大值",
			"bandwidth_limit":           "带宽上限",
			"tags":                      "文件系统标签",
			"tiering_detail":            "TieringDetail",
			"tiering_size_in_bytes":     "TieringSizeInBytes",
			"alloced_space":             "已分配空间",
			"storage_resource_pkg":      "文件系统绑定的预付费存储包（暂未支持）",
			"spec_policy":               "SpecPolicy",
			"auto_snapshot_policy_id":   "自动快照策略ID",
			"auto_scale_up_rule":        "AutoScaleUpRule",
			"snap_status":               "快照状态",
			"bandwidth_resource_pkg":    "带宽上限",
			"app_id":                    "appid",
			"tiering_state":             "TieringState",
			"ip_address":                "挂载点IP地址",
			"snap_id":                   "快照id",
			"zone_id":                   "可用区ID",
			"zone":                      "可用区名称，例如:ap-zone-1",
			"creation_token":            "用户自定义文件系统名称，优先级低于fs_name",
			"protocol":                  "文件系统协议类型， 值为TURBO",
			"fs_name":                   "用户自定义文件系统名称,与CreationToken 两者必须填一项",
			"availability_zone":         "可用区",
			"p_group_id":                "权限组ID",
			"p_group":                   "权限组信息",
			"name":                      "权限组名称",
			"cfs_version":               "TurboFS文件系统版本",
			"capacity":                  "TURBO文件系统的容量（TiB）",
			"pool_id":                   "资源池ID",
			"snapshot_id":               "快照ID",
			"storage_type":              "文件系统存储类型，值为 TB/TP ；其中 TB为标准型存储, TP为性能型存储",
			"net_interface":             "网络类型，值为 VPC，BASIC；其中 VPC 为私有网络，BASIC 为基础网络",
			"vpc_id":                    "私有网路（VPC） ID;当网络类型值为 VPC时，与UnVpcId 两者必须填一项",
			"un_vpc_id":                 "系统分配的VPC统一ID",
			"subnet_id":                 "子网， 当网络类型值为 VPC时，与UnSubnetId 两者必须填一项",
			"un_subnet_id":              "系统分配的子网统一 ID",
			"mount_ip":                  "指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP",
			"storage_resource_pkg_id":   "文件系统绑定的存储包，每个文件系统只能绑定一个",
			"bandwidth_resource_pkg_id": "文件系统绑定的带宽包，每个文件系统只能绑定一个",
			"encrypted":                 "文件系统是否加密，若留空则默认为不加密",
			"kms_key_id":                "加密密钥 ID",
			"resource_tags":             "资源标签",
			"project_id":                "项目ID",
			"creation_time":             "文件系统创建时间",
			"cfs_id":                    "文件系统ID",
			"tag_id":                    "TurboFS资源池id",
			"file_system_id":            "文件系统 ID",
			"life_cycle_state":          "文件系统状态",
			"size_byte":                 "文件系统已使用容量大小",
			"tag_key":                   "标签键",
			"tag_value":                 "标签值",
			"size_limit":                "文件系统最大空间限制",
		},
	})
}

func dataSourceTencentCloudTurbofsFileSystems() *schema.Resource {
	return &schema.Resource{
		Description: "Query the detail information of cloud file systems(Turbofs).",
		Read:        dataSourceTencentCloudTurbofsFileSystemsRead,

		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A specified file system ID used to query.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"file_system_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of cloud file system. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation datetime of the file system.",
						},
						"creation_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-defined file system name, with lower priority than fs_name.",
						},
						"file_system_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the file system.",
						},
						"fs_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of a file system.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone that the file system locates at.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The zone id that the file system locates at.",
						},
						"life_cycle_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of the file system",
						},
						"size_byte": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Usage of the file system",
						},
						"size_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max usage of the file system",
						},
						"size_limit_max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max limit of the usage of the file system",
						},
						"p_group": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"p_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of a permission group.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of a permission group.",
									},
								},
							},
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "File service protocol. Valid value, the default is `TURBO`.",
						},
						"cfs_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "TurboFS file system version.",
						},
						"capacity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Capacity of the TURBO file system in TiB. The value should be an integer multiple of the spec policy",
						},
						"pool_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the resource pool. If not specified, the system will automatically select the default resource pool.",
						},
						"snap_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snapshot ID.",
						},
						"snap_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Snapshot status.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "File service StorageType. Valid values are `TP` and `TB`.",
						},
						"bandwidth_limit": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Limit of the file system bandwidth.",
						},
						"app_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "App id of the file system.",
						},
						"tiering_detail": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tiering_size_in_bytes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "tiering size in bytes",
									},
								},
							},
							Description: "Tiering detail",
						},
						"tiering_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tiering state.",
						},
						"auto_snapshot_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auto snapshot policy id.",
						},
						"auto_scale_up_rule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auto scale up rule.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Allocated space of the file system.",
						},
						"alloced_space": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IP address of the mount target.",
						},
						"storage_resource_pkg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The storage resource package bound to the file system. Each file system can only be bound to one storage resource package.",
						},
						"bandwidth_resource_pkg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.",
						},
						"encrypted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the file system is encrypted.",
						},
						"kms_key_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.",
						},
						"tag_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TurboFS pool id.",
						},
						"tag_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "TurboFS pool name.",
						},
						"tags": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Tags.",
						},
						"spec_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Spec policy.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTurbofsFileSystemsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_turbofs_file_systems.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var fileSystemId *string
	//var vpcId string
	//var subnetId string
	//var name string
	//var zone string
	if v, ok := d.GetOk("file_system_id"); ok {
		fileSystemId = helper.String(v.(string))
	}
	//if v, ok := d.GetOk("vpc_id"); ok {
	//	vpcId = v.(string)
	//}
	//if v, ok := d.GetOk("subnet_id"); ok {
	//	subnetId = v.(string)
	//}
	//if v, ok := d.GetOk("name"); ok {
	//	name = v.(string)
	//}
	//if v, ok := d.GetOk("availability_zone"); ok {
	//	zone = v.(string)
	//}

	var fileSystems []*turbofs.FileSystemInfo
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		fileSystems, errRet = turbofsService.DescribeFileSystem(ctx, fileSystemId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	fileSystemList := make([]map[string]interface{}, 0, len(fileSystems))
	ids := make([]string, 0, len(fileSystems))

	for _, fileSystem := range fileSystems {
		fileSystemMap := make(map[string]interface{})

		if fileSystem.FileSystemId != nil {
			fileSystemMap["file_system_id"] = *fileSystem.FileSystemId
			ids = append(ids, *fileSystem.FileSystemId)
		}

		if fileSystem.CreationTime != nil {
			fileSystemMap["creation_time"] = *fileSystem.CreationTime
		}

		if fileSystem.CreationToken != nil {
			fileSystemMap["creation_token"] = *fileSystem.CreationToken
		}

		if fileSystem.FsName != nil {
			fileSystemMap["fs_name"] = *fileSystem.FsName
		}

		if fileSystem.Zone != nil {
			fileSystemMap["zone"] = *fileSystem.Zone
		}

		if fileSystem.ZoneId != nil {
			fileSystemMap["zone_id"] = helper.UInt64ToStr(*fileSystem.ZoneId)
		}

		if fileSystem.LifeCycleState != nil {
			fileSystemMap["life_cycle_state"] = *fileSystem.LifeCycleState
		}

		if fileSystem.SizeByte != nil {
			fileSystemMap["size_byte"] = int(*fileSystem.SizeByte)
		}

		if fileSystem.SizeLimit != nil {
			fileSystemMap["size_limit"] = int(*fileSystem.SizeLimit)
		}

		if fileSystem.SizeLimitMax != nil {
			fileSystemMap["size_limit_max"] = int(*fileSystem.SizeLimitMax)
		}

		if fileSystem.Protocol != nil {
			fileSystemMap["protocol"] = *fileSystem.Protocol
		}

		if fileSystem.CfsVersion != nil {
			fileSystemMap["cfs_version"] = *fileSystem.CfsVersion
		}

		if fileSystem.Capacity != nil {
			fileSystemMap["capacity"] = int(*fileSystem.Capacity)
		}

		if fileSystem.PoolId != nil {
			fileSystemMap["pool_id"] = *fileSystem.PoolId
		}

		if fileSystem.SnapshotId != nil {
			fileSystemMap["snap_id"] = *fileSystem.SnapshotId
		}

		if fileSystem.SnapStatus != nil {
			fileSystemMap["snap_status"] = *fileSystem.SnapStatus
		}

		if fileSystem.StorageType != nil {
			fileSystemMap["storage_type"] = *fileSystem.StorageType
		}

		if fileSystem.BandwidthLimit != nil {
			fileSystemMap["bandwidth_limit"] = *fileSystem.BandwidthLimit
		}

		if fileSystem.AppId != nil {
			fileSystemMap["app_id"] = int(*fileSystem.AppId)
		}

		if fileSystem.IpAddress != nil {
			fileSystemMap["ip_address"] = *fileSystem.IpAddress
		}

		if fileSystem.AllocedSpace != nil {
			fileSystemMap["alloced_space"] = int(*fileSystem.AllocedSpace)
		}

		if fileSystem.StorageResourcePkg != nil {
			fileSystemMap["storage_resource_pkg"] = *fileSystem.StorageResourcePkg
		}

		if fileSystem.BandwidthResourcePkg != nil {
			fileSystemMap["bandwidth_resource_pkg"] = *fileSystem.BandwidthResourcePkg
		}

		if fileSystem.Encrypted != nil {
			fileSystemMap["encrypted"] = *fileSystem.Encrypted
		}

		if fileSystem.KmsKeyId != nil {
			fileSystemMap["kms_key_id"] = *fileSystem.KmsKeyId
		}

		if fileSystem.TagId != nil {
			fileSystemMap["tag_id"] = int(*fileSystem.TagId)
		}

		if fileSystem.TagName != nil {
			fileSystemMap["tag_name"] = *fileSystem.TagName
		}

		// Handle PGroup
		if fileSystem.PGroup != nil {
			pGroupList := make([]map[string]interface{}, 0, 1)
			pGroupMap := make(map[string]interface{})

			if fileSystem.PGroup.PGroupId != nil {
				pGroupMap["p_group_id"] = *fileSystem.PGroup.PGroupId
			}

			if fileSystem.PGroup.Name != nil {
				pGroupMap["name"] = *fileSystem.PGroup.Name
			}

			pGroupList = append(pGroupList, pGroupMap)
			fileSystemMap["p_group"] = pGroupList
		}

		// Handle TieringDetail
		if fileSystem.TieringDetail != nil {
			tieringDetailList := make([]map[string]interface{}, 0, 1)
			tieringDetailMap := make(map[string]interface{})

			if fileSystem.TieringDetail.TieringSizeInBytes != nil {
				tieringDetailMap["tiering_size_in_bytes"] = int(*fileSystem.TieringDetail.TieringSizeInBytes)
			}

			tieringDetailList = append(tieringDetailList, tieringDetailMap)
			fileSystemMap["tiering_detail"] = tieringDetailList
		} else {
			fileSystemMap["tiering_detail"] = []map[string]interface{}{}
		}

		// Handle TieringState
		if fileSystem.TieringState != nil {
			fileSystemMap["tiering_state"] = *fileSystem.TieringState
		} else {
			fileSystemMap["tiering_state"] = ""
		}

		// Handle AutoSnapshotPolicyId
		if fileSystem.AutoSnapshotPolicyId != nil {
			fileSystemMap["auto_snapshot_policy_id"] = *fileSystem.AutoSnapshotPolicyId
		} else {
			fileSystemMap["auto_snapshot_policy_id"] = ""
		}

		// Handle AutoScaleUpRule
		if fileSystem.AutoScaleUpRule != nil {
			fileSystemMap["auto_scale_up_rule"] = *fileSystem.AutoScaleUpRule
		} else {
			fileSystemMap["auto_scale_up_rule"] = ""
		}

		// Handle SpecPolicy
		if fileSystem.SpecPolicy != nil {
			fileSystemMap["spec_policy"] = *fileSystem.SpecPolicy
		} else {
			fileSystemMap["spec_policy"] = ""
		}

		// Handle Tags
		if fileSystem.Tags != nil && len(fileSystem.Tags) > 0 {
			// Convert tags to string representation or handle as needed
			tagsStr := ""
			for i, tag := range fileSystem.Tags {
				if tag.TagKey != nil && tag.TagValue != nil {
					if i > 0 {
						tagsStr += ","
					}
					tagsStr += *tag.TagKey + ":" + *tag.TagValue
				}
			}
			fileSystemMap["tags"] = tagsStr
		} else {
			fileSystemMap["tags"] = ""
		}

		fileSystemList = append(fileSystemList, fileSystemMap)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("file_system_list", fileSystemList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set turbofs file system list fail, reason:%s\n ", logId, err.Error())
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), fileSystemList); err != nil {
			return err
		}
	}
	return nil
}
