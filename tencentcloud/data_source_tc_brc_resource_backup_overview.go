/*
	Use this data source to query detailed information of brc backup overviews

# Example Usage

```hcl

	 data "tencentcloudenterprise_brc_resource_backup_overview" "overview" {
	   result_output_file = "backup_overview.json"
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
	registerDataDescriptionProvider("tencentcloudenterprise_brc_resource_backup_overview", CNDescription{
		TerraformTypeCN: "BRC资源备份概览",
		DescriptionCN:   "提供BRC资源备份概览数据源，用于查询资源备份的概览信息。",
		AttributesCN: map[string]string{
			"resource_type":       "备份资源类型，有效值：INSTANCE(CVM实例)、DISK(CBS云硬盘)、CFS、COS、CSP、MySQL_MariaDB和TDSQL_MySQL",
			"result_output_file":  "用于保存结果",
			"backup_overview":     "备份概览列表",
			"overview_detail_set": "概览详情集合",
		},
	})
}

func dataSourceTencentCloudBrcResourceBackupOverview() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudBrcResourceBackupOverviewRead,
		Description: "Use this data source to query detailed information of brc backup overviews",
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Resource type filter. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.",
				//ValidateFunc: validateAllowedStringValue(BackupResouceTypes),
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"backup_overview": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Backup overview list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type.",
						},
						"backup_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup count.",
						},
						"backup_resource_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup resource count.",
						},
						"backup_size_mb": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup size in MB.",
						},
					},
				},
			},
			"overview_detail_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Overview detail set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type.",
						},
						"is_open": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the service is open.",
						},
						"is_support": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the service is supported.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region.",
						},
						"resource_overview": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Resource overview.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_resource_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Backup resource count.",
									},
									"backup_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Backup count.",
									},
									"backup_size_mb": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Backup size in MB.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudBrcResourceBackupOverviewRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_brc_resource_backup_overview.read")()

	var (
		logId      = getLogId(contextNil)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	request := brc.NewDescribeResourceBackupOverviewRequest()

	if v, ok := d.GetOk("resource_type"); ok {
		request.ResourceType = helper.String(v.(string))
	}

	ratelimit.Check(request.GetAction())
	response, err := brcService.client.UseBrcClient().DescribeResourceBackupOverview(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}

	// 处理backup_overview
	backupOverviewList := make([]map[string]interface{}, 0)
	if response.Response != nil && response.Response.BackupOverview != nil {
		for _, overview := range response.Response.BackupOverview {
			overviewMap := map[string]interface{}{
				"resource_type":         overview.ResourceType,
				"backup_count":          overview.BackupCount,
				"backup_resource_count": overview.BackupResourceCount,
				"backup_size_mb":        overview.BackupSizeMb,
			}
			backupOverviewList = append(backupOverviewList, overviewMap)
		}
	}

	// 处理overview_detail_set
	overviewDetailList := make([]map[string]interface{}, 0)
	if response.Response != nil && response.Response.OverviewDetailSet != nil {
		for _, detail := range response.Response.OverviewDetailSet {
			detailMap := map[string]interface{}{
				"resource_type": detail.ResourceType,
				"is_open":       detail.IsOpen,
				"is_support":    detail.IsSupport,
				"region":        detail.Region,
			}

			// 设置resource_overview
			if detail.ResourceOverview != nil {
				resourceOverview := map[string]interface{}{
					"backup_resource_count": detail.ResourceOverview.BackupResourceCount,
					"backup_count":          detail.ResourceOverview.BackupCount,
					"backup_size_mb":        detail.ResourceOverview.BackupSizeMb,
				}
				detailMap["resource_overview"] = []map[string]interface{}{resourceOverview}
			}

			overviewDetailList = append(overviewDetailList, detailMap)
		}
	}

	// 创建ID用于hash
	resourceTypes := make([]string, 0)
	if v, ok := d.GetOk("resource_type"); ok {
		resourceTypes = append(resourceTypes, v.(string))
	} else {
		for _, overview := range backupOverviewList {
			if resourceType, ok := overview["resource_type"].(*string); ok && resourceType != nil {
				resourceTypes = append(resourceTypes, *resourceType)
			}
		}
	}

	d.SetId(fmt.Sprintf("brc_resource_backup_overview_%s", helper.DataResourceIdsHash(resourceTypes)))
	_ = d.Set("backup_overview", backupOverviewList)
	_ = d.Set("overview_detail_set", overviewDetailList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		result := map[string]interface{}{
			"backup_overview":     backupOverviewList,
			"overview_detail_set": overviewDetailList,
		}
		if err := writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
