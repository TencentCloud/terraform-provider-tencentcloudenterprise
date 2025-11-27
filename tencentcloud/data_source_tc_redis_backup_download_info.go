/*
Use this data source to query detailed information of backup_download_info

# Example Usage

```hcl

	data "tencentcloudenterprise_redis_backup_download_info" "backup_download_info" {
	  instance_id = "crs-iw7d9wdd"
	  backup_id = "641186639-8362913-1516672770"
	  # limit_type = "NoLimit"
	  # vpc_comparison_symbol = "In"
	  # ip_comparison_symbol = "In"
	  # limit_vpc {
		# 	region = "ap-guangzhou"
		# 	vpc_list = [""]
	  # }
	  # limit_ip = [""]
	}

```
*/
package tencentcloud

import (
	"context"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_redis_backup_download_info", CNDescription{
		TerraformTypeCN: "Redis®备份下载信息",
		DescriptionCN:   "提供Redis备份下载信息数据源，用于查询Redis备份下载信息的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":           "实例的ID",
			"backup_id":             "备份ID，可通过[DescriptionInstanceBackups]访问(https://cloud.tencent.com/document/product/239/20011)接口返回要获取的参数RedisBackupSet",
			"limit_type":            "下载备份文件的网络限制类型：-无限制：没有限制，可以从和内外网下载备份文件-LimitOnlyIntranet:只有自动分配的内网地址才能下载备份文件-自定义：指用户定义的专用网络可下载的备份文件",
			"vpc_comparison_symbol": "此参数仅支持输入In，这意味着自定义LimitVpc可以下载备份文件",
			"ip_comparison_symbol":  "标识自定义的LimitIP地址是否可以下载备份文件-在：自定义IP地址可供下载-NotIn:自定义IP不可下载",
			"limit_vpc":             "可下载备份文件的自定义VPC ID如果参数LimitType为**自定义**，则需要配置此参数",
			"region":                "自定义备份文件下载到的专有网络的区域",
			"vpc_list":              "自定义要下载备份文件的VPCs列表",
			"limit_ip":              "可下载备份文件的自定义VPC IP地址如果参数LimitType为**自定义**，则需要配置此参数",
			"backup_infos":          "备份文件信息的列表",
			"file_name":             "备份文件名",
			"file_size":             "备份文件大小以B为单位，如果为0，则无效",
			"download_url":          "Internet上的备份文件下载地址（6小时）",
			"inner_download_url":    "备份文件intranet下载地址（6小时）",
			"result_output_file":    "用于保存结果",
		},
	})
}

func dataSourceTencentCloudRedisBackupDownloadInfo() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of backup_download_info",
		Read:        dataSourceTencentCloudRedisBackupDownloadInfoRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"backup_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The backup ID, which can be accessed via [DescribeInstanceBackups](https://cloud.tencent.com/document/product/239/20011) interface returns the parameter RedisBackupSet to get.",
			},

			// "limit_type": {
			// 	Optional:    true,
			// 	Type:        schema.TypeString,
			// 	Description: "Types of network restrictions for downloading backup files:- NoLimit: There is no limit, and backup files can be downloaded from both Tencent Cloud and internal and external networks.- LimitOnlyIntranet: Only intranet addresses automatically assigned by Tencent Cloud can download backup files.- Customize: refers to a user-defined private network downloadable backup file.",
			// },

			// "vpc_comparison_symbol": {
			// 	Optional:    true,
			// 	Type:        schema.TypeString,
			// 	Description: "This parameter only supports entering In, which means that the custom LimitVpc can download the backup file.",
			// },

			// "ip_comparison_symbol": {
			// 	Optional:    true,
			// 	Type:        schema.TypeString,
			// 	Description: "Identifies whether the customized LimitIP address can download the backup file.- In: Custom IP addresses are available for download.- NotIn: Custom IPs are not available for download.",
			// },

			// "limit_vpc": {
			// 	Optional:    true,
			// 	Type:        schema.TypeList,
			// 	Description: "A custom VPC ID for a downloadable backup file.If the parameter LimitType is **Customize**, you need to configure this parameter.",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"region": {
			// 				Type:        schema.TypeString,
			// 				Required:    true,
			// 				Description: "Customize the region of the VPC to which the backup file is downloaded.",
			// 			},
			// 			"vpc_list": {
			// 				Type: schema.TypeSet,
			// 				Elem: &schema.Schema{
			// 					Type: schema.TypeString,
			// 				},
			// 				Required:    true,
			// 				Description: "Customize the list of VPCs to download backup files.",
			// 			},
			// 		},
			// 	},
			// },

			// "limit_ip": {
			// 	Optional: true,
			// 	Type:     schema.TypeSet,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// 	Description: "A custom VPC IP address for downloadable backup files.If the parameter LimitType is **Customize**, you need to configure this parameter.",
			// },

			"backup_infos": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "A list of backup file information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup file name.",
						},
						"file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The backup file size is in unit B, if it is 0, it is invalid.",
						},
						"download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup file download address on the Internet (6 hours).",
						},
						"inner_download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup file intranet download address (6 hours).",
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

func dataSourceTencentCloudRedisBackupDownloadInfoRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_redis_backup_download_info.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("backup_id"); ok {
		paramMap["backup_id"] = helper.String(v.(string))
	}

	//if v, ok := d.GetOk("limit_type"); ok {
	//	paramMap["limit_type"] = helper.String(v.(string))
	//}
	//
	//if v, ok := d.GetOk("vpc_comparison_symbol"); ok {
	//	paramMap["vpc_comparison_symbol"] = helper.String(v.(string))
	//}
	//
	//if v, ok := d.GetOk("ip_comparison_symbol"); ok {
	//	paramMap["ip_comparison_symbol"] = helper.String(v.(string))
	//}
	//
	//if v, ok := d.GetOk("limit_vpc"); ok {
	//	limitVpcSet := v.([]interface{})
	//	tmpSet := make([]*redis.BackupLimitVpcItem, 0, len(limitVpcSet))
	//
	//	for _, item := range limitVpcSet {
	//		backupLimitVpcItem := redis.BackupLimitVpcItem{}
	//		backupLimitVpcItemMap := item.(map[string]interface{})
	//
	//		if v, ok := backupLimitVpcItemMap["region"]; ok {
	//			backupLimitVpcItem.Region = helper.String(v.(string))
	//		}
	//		if v, ok := backupLimitVpcItemMap["vpc_list"]; ok {
	//			vpcListSet := v.(*schema.Set).List()
	//			backupLimitVpcItem.VpcList = helper.InterfacesStringsPoint(vpcListSet)
	//		}
	//		tmpSet = append(tmpSet, &backupLimitVpcItem)
	//	}
	//	paramMap["limit_vpc"] = tmpSet
	//}

	//if v, ok := d.GetOk("limit_ip"); ok {
	//	limitIpSet := v.(*schema.Set).List()
	//	paramMap["limit_ip"] = helper.InterfacesStringsPoint(limitIpSet)
	//}

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}

	var backupInfos []*redis.BackupDownloadInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeRedisBackupDownloadInfoByFilter(ctx, paramMap)
		if e != nil {
			if ee, ok := e.(*sdkErrors.CloudSDKError); ok {
				if ee.Code == "FailedOperation.SystemError" {
					return resource.NonRetryableError(e)
				}
			}
			return retryError(e)
		}
		backupInfos = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(backupInfos))
	tmpList := make([]map[string]interface{}, 0, len(backupInfos))

	if backupInfos != nil {
		for _, backupDownloadInfo := range backupInfos {
			backupDownloadInfoMap := map[string]interface{}{}

			if backupDownloadInfo.FileName != nil {
				backupDownloadInfoMap["file_name"] = backupDownloadInfo.FileName
			}

			if backupDownloadInfo.FileSize != nil {
				backupDownloadInfoMap["file_size"] = backupDownloadInfo.FileSize
			}

			if backupDownloadInfo.DownloadUrl != nil {
				backupDownloadInfoMap["download_url"] = backupDownloadInfo.DownloadUrl
			}

			if backupDownloadInfo.InnerDownloadUrl != nil {
				backupDownloadInfoMap["inner_download_url"] = backupDownloadInfo.InnerDownloadUrl
			}

			ids = append(ids, *backupDownloadInfo.FileName)
			tmpList = append(tmpList, backupDownloadInfoMap)
		}

		_ = d.Set("backup_infos", tmpList)
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
