/*
Use this data source to query detailed information of cfs mount_targets

# Example Usage

```hcl

	data "tencentcloudenterprise_cfs_mount_targets" "mount_targets" {
	  file_system_id = "cfs-iobiaxtj"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cfs_mount_targets", CNDescription{
		TerraformTypeCN: "CFS挂载点",
		DescriptionCN:   "提供CFS文件系统挂载点数据源，用于查询文件系统挂载点的详细信息。",
		AttributesCN: map[string]string{
			"file_system_id":     "文件系统ID",
			"mount_targets":      "装载目标详细信息",
			"mount_target_id":    "装载目标ID",
			"ip_address":         "装载目标IP",
			"fs_id":              "装载根目录",
			"life_cycle_state":   "装载目标状态",
			"network_interface":  "网络类型",
			"vpc_id":             "VPC ID",
			"vpc_name":           "VPC名称",
			"subnet_id":          "子网ID",
			"subnet_name":        "子网名称",
			"ccn_id":             "CFS Turbo使用的CCN实例ID",
			"cidr_block":         "CFS Turbo使用的CCN IP范围",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudCfsMountTargets() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of cfs mount_targets",
		Read:        dataSourceTencentCloudCfsMountTargetsRead,
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "File system ID.",
			},

			"mount_targets": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Mount target details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_system_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "File system ID.",
						},
						"mount_target_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mount target ID.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mount target IP.",
						},
						"fs_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mount root-directory.",
						},
						"life_cycle_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Mount target status.",
						},
						"network_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network type.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC ID.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC name.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet ID.",
						},
						"subnet_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet name.",
						},
						"ccn_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CCN instance ID used by CFS Turbo.",
						},
						"cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CCN IP range used by CFS Turbo.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCfsMountTargetsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cfs_mount_targets.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var mountTargets []*cfs.MountInfo

	fsId := d.Get("file_system_id").(string)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCfsMountTargetsById(ctx, fsId)
		if e != nil {
			return retryError(e)
		}
		mountTargets = result
		return nil
	})
	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(mountTargets))

	if mountTargets != nil {
		for _, mountInfo := range mountTargets {
			mountInfoMap := map[string]interface{}{}

			if mountInfo.FileSystemId != nil {
				mountInfoMap["file_system_id"] = mountInfo.FileSystemId
			}

			if mountInfo.MountTargetId != nil {
				mountInfoMap["mount_target_id"] = mountInfo.MountTargetId
			}

			if mountInfo.IpAddress != nil {
				mountInfoMap["ip_address"] = mountInfo.IpAddress
			}

			if mountInfo.FSID != nil {
				mountInfoMap["fs_id"] = mountInfo.FSID
			}

			if mountInfo.LifeCycleState != nil {
				mountInfoMap["life_cycle_state"] = mountInfo.LifeCycleState
			}

			if mountInfo.NetworkInterface != nil {
				mountInfoMap["network_interface"] = mountInfo.NetworkInterface
			}

			if mountInfo.VpcId != nil {
				mountInfoMap["vpc_id"] = mountInfo.VpcId
			}

			if mountInfo.VpcName != nil {
				mountInfoMap["vpc_name"] = mountInfo.VpcName
			}

			if mountInfo.SubnetId != nil {
				mountInfoMap["subnet_id"] = mountInfo.SubnetId
			}

			if mountInfo.SubnetName != nil {
				mountInfoMap["subnet_name"] = mountInfo.SubnetName
			}

			/*
				if mountInfo.CcnID != nil {
					mountInfoMap["ccn_id"] = mountInfo.CcnID
				}

				if mountInfo.CidrBlock != nil {
					mountInfoMap["cidr_block"] = mountInfo.CidrBlock
				}

			*/

			tmpList = append(tmpList, mountInfoMap)
		}

		_ = d.Set("mount_targets", tmpList)
	}

	d.SetId(fsId)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
