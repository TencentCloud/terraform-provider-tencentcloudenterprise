/*
Use this data source to query detailed information of redis backup

# Example Usage

```hcl

	data "tencentcloudenterprise_redis_backup" "backup" {
	  instance_id = "crs-c1nl9rpv"
	  begin_time = "2023-04-07 03:57:30"
	  end_time = "2023-04-07 03:57:56"
	  status = [2]
	  instance_name = "Keep-terraform"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_redis_backup", CNDescription{
		TerraformTypeCN: "Redis®备份",
		DescriptionCN:   "提供Redis备份数据源，用于查询Redis备份的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例的ID",
			"begin_time":         "开始时间，如2017-02-08 19:09:26查询实例在[beginTime，endTime]时间段内开始备份的备份列表",
			"end_time":           "结束时间，如2017-02-08 19:09:26查询实例在[beginTime，endTime]时间段内开始备份的备份列表",
			"status":             "备份任务的状态：1:备份正在进行中2:备份正常3:备份到RDB文件处理4:RDB转换完成-1:备份已过期-2:备份已删除",
			"instance_name":      "实例名称，支持基于实例名称的模糊搜索",
			"backup_set":         "实例的备份阵列",
			"start_time":         "备份开始时间",
			"backup_id":          "备份ID",
			"backup_type":        "备份类型1:用户启动的手动备份0:系统在凌晨启动的备份",
			"remark":             "注意备份信息",
			"locked":             "备份是否已锁定0:未锁定1:已锁定",
			"backup_size":        "内部字段，用户可以忽略这些字段",
			"full_backup":        "内部字段，用户可以忽略这些字段",
			"instance_type":      "内部字段，用户可以忽略这些字段",
			"region":             "备份所在的区域",
			"file_type":          "备份文件类型",
			"expire_time":        "备份文件过期时间",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudRedisBackup() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud Redis® Backup",
		Read:        dataSourceTencentCloudRedisBackupRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"begin_time": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Start time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.",
			},

			"end_time": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "End time, such as 2017-02-08 19:09:26.Query the list of backups that the instance started backing up during the [beginTime, endTime] time period.",
			},

			"status": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "Status of the backup task:1: Backup is in the process.2: The backup is normal.3: Backup to RDB file processing.4: RDB conversion completed.-1: The backup has expired.-2: Backup deleted.",
			},

			// "instance_name": {
			// 	Optional:    true,
			// 	Type:        schema.TypeString,
			// 	Description: "Instance name, which supports fuzzy search based on instance name.",
			// },

			"backup_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "An array of backups for the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup start time.",
						},
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup ID.",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup type.1: User-initiated manual backup.0: System-initiated backup in the early morning.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup status.1: The backup is locked by another process.2: The backup is normal and not locked by any process.-1: The backup has expired.3: The backup is being exported.4: The backup export is successful.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Notes information for the backup.",
						},
						"locked": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the backup is locked.0: Not locked.1: Has been locked.",
						},
						"backup_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Internal fields, which can be ignored by the user.",
						},
						"full_backup": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Internal fields, which can be ignored by the user.",
						},
						"instance_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Internal fields, which can be ignored by the user.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region where the backup is located.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup end time.",
						},
						"file_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Back up file types.",
						},
						"expire_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup file expiration time.",
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

func dataSourceTencentCloudRedisBackupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_redis_backup.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("begin_time"); ok {
		paramMap["begin_time"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("end_time"); ok {
		paramMap["end_time"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("status"); ok {
		statusSet := v.(*schema.Set).List()
		statusList := []*int64{}
		for i := range statusSet {
			status := statusSet[i].(int)
			statusList = append(statusList, helper.IntInt64(status))
		}
		paramMap["status"] = statusList
	}

	// if v, ok := d.GetOk("instance_name"); ok {
	// 	paramMap["InstanceName"] = helper.String(v.(string))
	// }

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}

	var backupSet []*redis.RedisBackupSet

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeRedisBackupByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		backupSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(backupSet))
	tmpList := make([]map[string]interface{}, 0, len(backupSet))

	if backupSet != nil {
		for _, redisBackupSet := range backupSet {
			redisBackupSetMap := map[string]interface{}{}

			if redisBackupSet.StartTime != nil {
				redisBackupSetMap["start_time"] = redisBackupSet.StartTime
			}

			if redisBackupSet.BackupId != nil {
				redisBackupSetMap["backup_id"] = redisBackupSet.BackupId
			}

			if redisBackupSet.BackupType != nil {
				redisBackupSetMap["backup_type"] = redisBackupSet.BackupType
			}

			if redisBackupSet.Status != nil {
				redisBackupSetMap["status"] = redisBackupSet.Status
			}

			if redisBackupSet.Remark != nil {
				redisBackupSetMap["remark"] = redisBackupSet.Remark
			}

			if redisBackupSet.Locked != nil {
				redisBackupSetMap["locked"] = redisBackupSet.Locked
			}

			if redisBackupSet.BackupSize != nil {
				redisBackupSetMap["backup_size"] = redisBackupSet.BackupSize
			}

			if redisBackupSet.FullBackup != nil {
				redisBackupSetMap["full_backup"] = redisBackupSet.FullBackup
			}

			if redisBackupSet.InstanceType != nil {
				redisBackupSetMap["instance_type"] = redisBackupSet.InstanceType
			}

			//if redisBackupSet.InstanceId != nil {
			//	redisBackupSetMap["instance_id"] = redisBackupSet.InstanceId
			//}
			//
			//if redisBackupSet.InstanceName != nil {
			//	redisBackupSetMap["instance_name"] = redisBackupSet.InstanceName
			//}
			//
			//if redisBackupSet.Region != nil {
			//	redisBackupSetMap["region"] = redisBackupSet.Region
			//}

			if redisBackupSet.EndTime != nil {
				redisBackupSetMap["end_time"] = redisBackupSet.EndTime
			}

			if redisBackupSet.FileType != nil {
				redisBackupSetMap["file_type"] = redisBackupSet.FileType
			}

			if redisBackupSet.ExpireTime != nil {
				redisBackupSetMap["expire_time"] = redisBackupSet.ExpireTime
			}

			//ids = append(ids, *redisBackupSet.InstanceId)
			tmpList = append(tmpList, redisBackupSetMap)
		}

		_ = d.Set("backup_set", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}
