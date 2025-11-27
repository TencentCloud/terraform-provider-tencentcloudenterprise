/*
Use this data source to query detailed information of brc backups

# Example Usage

```hcl

	data "tencentcloudenterprise_brc_backups" "all_backups" {
	  result_output_file = "all_backups.json"
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_brc_backups", CNDescription{
		TerraformTypeCN: "BRC备份列表",
		DescriptionCN:   "提供BRC备份列表数据源，用于查询备份的详细信息。",
		AttributesCN: map[string]string{
			"backup_id":             "按照备份的 ID 过滤。快照 ID 形如：backup-11112222",
			"backup_name":           "按照备份名称过滤",
			"backup_state":          "按照备份状态过滤,NORMAL（正常）、CREATING（创建中）、ROLLBACKING（回滚中）",
			"disk_usage":            "按创建快照的云盘类型过滤,SYSTEM_DISK（代表系统盘）、DATA_DISK（代表数据盘）",
			"platform_project_id":   "按备份所属项目 ID 过滤",
			"disk_id":               "按照创建备份的云硬盘 ID 过滤",
			"backup_group_id":       "按照备份绑定的备份组 ID 过滤",
			"zone":                  "按照可用区过滤",
			"result_output_file":    "用于保存结果",
			"backup_set":            "备份列表",
			"copy_from_remote":      "是否为远程复制的备份。",
			"is_permanent":          "是否永久保留。",
			"deadline_time":         "备份的到期时间。",
			"percent":               "备份创建的进度。",
			"share_reference":       "备份被共享的次数。",
			"disk_size":             "创建备份的云硬盘大小，单位GB。",
			"copying_to_regions":    "备份当前正在远程复制的目标地域列表。",
			"encrypt":               "备份是否为加密备份。",
			"placement":             "备份所在的位置。",
			"tags":                  "备份绑定的标签列表。",
			"appid":                 "用户AppId。",
			"disk_name":             "云盘名称。",
			"disk_details":          "创建备份时刻，云硬盘各属性的详情。",
			"backup_class":          "全量、增量备份信息；FULL表示全量备份，INC表示增量备份。",
			"account_uin":           "主账号uin。",
			"sub_account_uin":       "创建备份的子账号uin。",
			"auto_backup_policy_id": "创建当前备份的定期备份策略ID，为null则为手动创建的备份。",
			"archive_status":        "归档状态。",
			"backup_size":           "备份大小",
			"backup_type":           "备份类型",
			"create_time":           "创建时间",
			"expire_time":           "过期时间",
			"account_name":          "账户名称",
			"create_speed":          "创建速度",
			"need_archive":          "是否需要归档",
		},
	})
}

func dataSourceTencentCloudBrcBackups() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudBrcBackupsRead,
		Description: "Use this data source to query detailed information of brc backups",
		Schema: map[string]*schema.Schema{
			"backup_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Backup ID filter.",
			},
			"backup_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Backup name filter.",
			},
			"backup_state": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Backup state filter.",
			},
			"disk_usage": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Disk usage type filter, SYSTEM_DISK (system disk) or DATA_DISK (data disk).",
			},
			"platform_project_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Platform project ID filter.",
			},
			"disk_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Disk ID filter for creating backup.",
			},
			"zone": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Availability zone filter.",
			},
			"backup_group_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Backup group ID filter.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"backup_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Backup list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup ID.",
						},
						"backup_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup name.",
						},
						"backup_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup state.",
						},
						"backup_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup size in bytes.",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup type.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expire time.",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account name.",
						},
						"zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Zone.",
						},
						"copy_from_remote": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the backup is copied from remote.",
						},
						"is_permanent": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the backup is permanently retained.",
						},
						"deadline_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup expiration time.",
						},
						"percent": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup creation progress percentage.",
						},
						"share_reference": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of times the backup has been shared.",
						},
						"disk_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Size of the disk used to create the backup, in GB.",
						},
						"disk_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the disk used to create the backup.",
						},
						"platform_project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project ID to which the backup belongs.",
						},
						"copying_to_regions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of target regions where the backup is currently being remotely copied.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"encrypt": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the backup is encrypted.",
						},
						"disk_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk usage type for creating the backup, SYSTEM_DISK (system disk) or DATA_DISK (data disk).",
						},
						"placement": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Location information of the backup.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cage_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cage ID.",
									},
									"zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Availability zone.",
									},
									"project_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Project ID.",
									},
									"cdc_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CDC name.",
									},
									"cdc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "CDC ID.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Project name.",
									},
									"dedicated_cluster_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Dedicated cluster ID.",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of tags bound to the backup.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag key.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tag value.",
									},
								},
							},
						},
						"appid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "User AppId.",
						},
						"disk_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk name.",
						},
						"disk_details": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Details of disk attributes at the time of backup creation.",
						},
						"backup_class": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full or incremental backup information; FULL indicates full backup, INC indicates incremental backup.",
						},
						"backup_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup group ID associated with the backup.",
						},
						"account_uin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Main account UIN.",
						},
						"sub_account_uin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub-account UIN that created the backup.",
						},
						"auto_backup_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the scheduled backup policy that created the current backup, null if manually created.",
						},
						"archive_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Archive status.",
						},
						"create_speed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup creation speed.",
						},
						"need_archive": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether archiving is needed.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudBrcBackupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_brc_backups.read")()

	var (
		logId      = getLogId(contextNil)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	request := brc.NewDescribeBackupsRequest()
	request.Limit = helper.Uint64(100)
	request.Offset = helper.Uint64(0)

	// 添加过滤条件
	filters := make([]*brc.Filter, 0)

	if v, ok := d.GetOk("backup_id"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("backup-id"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("backup_name"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("backup-name"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("disk_id"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("disk-id"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("disk_usage"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("disk-usage"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("platform_project_id"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("platform-project-id"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("backup_group_id"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("backup-group-id"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("zone"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("zone"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("backup_state"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("backup-state"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if len(filters) > 0 {
		request.Filters = filters
	}

	ratelimit.Check(request.GetAction())
	response, err := brcService.client.UseBrcClient().DescribeBackups(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}

	backupSet := make([]map[string]interface{}, 0)
	if response.Response != nil && response.Response.BackupSet != nil {
		for _, backup := range response.Response.BackupSet {
			backupMap := map[string]interface{}{
				"backup_id":           backup.BackupId,
				"backup_name":         backup.BackupName,
				"backup_state":        backup.BackupState,
				"backup_size":         backup.DiskSize,
				"backup_type":         backup.BackupType,
				"create_time":         backup.CreateTime,
				"expire_time":         backup.DeadlineTime,
				"account_name":        backup.AccountName,
				"copy_from_remote":    backup.CopyFromRemote,
				"is_permanent":        backup.IsPermanent,
				"deadline_time":       backup.DeadlineTime,
				"percent":             backup.Percent,
				"share_reference":     backup.ShareReference,
				"disk_size":           backup.DiskSize,
				"disk_id":             backup.DiskId,
				"platform_project_id": backup.PlatformProjectId,
				"copying_to_regions":  backup.CopyingToRegions,
				"encrypt":             backup.Encrypt,
				"disk_usage":          backup.DiskUsage,

				"appid":                 backup.AppId,
				"disk_name":             backup.DiskName,
				"backup_class":          backup.BackupClass,
				"backup_group_id":       backup.BackupGroupId,
				"account_uin":           backup.AccountUin,
				"sub_account_uin":       backup.SubAccountUin,
				"auto_backup_policy_id": backup.AutoBackupPolicyId,
				"archive_status":        backup.ArchiveStatus,
				"create_speed":          backup.CreateSpeed,
				"need_archive":          backup.NeedArchive,
			}

			if backup.Placement != nil {
				placementList := make([]map[string]interface{}, 0, 1)
				placementMap := map[string]interface{}{
					"cage_id":              backup.Placement.CageId,
					"zone":                 backup.Placement.Zone,
					"project_id":           backup.Placement.ProjectId,
					"cdc_name":             backup.Placement.CdcName,
					"cdc_id":               backup.Placement.CdcId,
					"project_name":         backup.Placement.ProjectName,
					"dedicated_cluster_id": backup.Placement.DedicatedClusterId,
				}
				placementList = append(placementList, placementMap)
				backupMap["placement"] = placementList
				backupMap["zone"] = backup.Placement.Zone
			}

			if backup.Tags != nil {
				tagsList := make([]map[string]interface{}, 0, len(backup.Tags))
				for _, tag := range backup.Tags {
					tagMap := map[string]interface{}{
						"key":   tag.Key,
						"value": tag.Value,
					}
					tagsList = append(tagsList, tagMap)
				}
				backupMap["tags"] = tagsList
			}

			backupSet = append(backupSet, backupMap)
		}
	}

	// 创建ID列表用于hash
	backupIds := make([]string, 0, len(backupSet))
	for _, backup := range backupSet {
		if id, ok := backup["backup_id"].(*string); ok && id != nil {
			backupIds = append(backupIds, *id)
		}
	}

	d.SetId(fmt.Sprintf("brc_backups_%s", helper.DataResourceIdsHash(backupIds)))
	_ = d.Set("backup_set", backupSet)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), backupSet); err != nil {
			return err
		}
	}

	return nil
}
