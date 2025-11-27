/*
Use this data source to query detailed information of CBS snapshots.

# Example Usage

```hcl

	data "tencentcloudenterprise_cbs_snapshots" "snapshots" {
	  snapshot_id        = "snap-f3io7adt"
	  result_output_file = "mytestpath"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cbs_snapshots", CNDescription{
		TerraformTypeCN: "cbs快照",
		DescriptionCN:   "提供CBS快照数据源，用于查询CBS快照的详细信息。",
		AttributesCN: map[string]string{
			"snapshot_id":        "要查询的快照的ID",
			"snapshot_name":      "要查询的快照的名称",
			"storage_id":         "快照的CBS的ID",
			"storage_usage":      "快照的CBS的类型以及可用值包括SYSTEM_DISK和DATA_DISK",
			"project_id":         "快照中项目的ID",
			"availability_zone":  "CBS实例所在的可用区域",
			"result_output_file": "用于保存结果",
			"snapshot_list":      "快照列表每个元素都包含以下属性：",
			"storage_size":       "快照的存储卷",
			"percent":            "快照创建进度百分比",
			"create_time":        "快照的创建时间",
			"encrypt":            "指示快照是否已加密",
		},
	})
}

func dataSourceTencentCloudCbsSnapshots() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides detailed information of CBS snapshots.",
		Read:        dataSourceTencentCloudCbsSnapshotsRead,

		Schema: map[string]*schema.Schema{
			"snapshot_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the snapshot to be queried.",
			},
			"snapshot_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the snapshot to be queried.",
			},
			"storage_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the the CBS which this snapshot created from.",
			},
			"storage_usage": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateAllowedStringValue(CBS_STORAGE_USAGE),
				Description:  "Types of CBS which this snapshot created from, and available values include SYSTEM_DISK and DATA_DISK.",
			},
			"project_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the project within the snapshot.",
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The available zone that the CBS instance locates at.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"snapshot_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of snapshot. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"snapshot_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the snapshot.",
						},
						"snapshot_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the snapshot.",
						},
						"storage_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the the CBS which this snapshot created from.",
						},
						"storage_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Types of CBS which this snapshot created from.",
						},
						"storage_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Volume of storage which this snapshot created from.",
						},
						"availability_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available zone that the CBS instance locates at.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the project within the snapshot.",
						},
						"percent": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Snapshot creation progress percentage.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of snapshot.",
						},
						"encrypt": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the snapshot is encrypted.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCbsSnapshotsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cbs_snapshots.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]string)
	if v, ok := d.GetOk("snapshot_id"); ok {
		params["snapshot-id"] = v.(string)
	}
	if v, ok := d.GetOk("snapshot_name"); ok {
		params["snapshot-name"] = v.(string)
	}
	if v, ok := d.GetOk("storage_id"); ok {
		params["disk-id"] = v.(string)
	}
	if v, ok := d.GetOk("storage_usage"); ok {
		params["disk-usage"] = v.(string)
	}
	if v, ok := d.GetOkExists("project_id"); ok {
		params["project-id"] = v.(string)
	}
	if v, ok := d.GetOk("availability_zone"); ok {
		params["zone"] = v.(string)
	}

	cbsService := CbsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		snapshots, e := cbsService.DescribeSnapshotsByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		ids := make([]string, 0, len(snapshots))
		snapshotList := make([]map[string]interface{}, 0, len(snapshots))
		for _, snapshot := range snapshots {
			mapping := map[string]interface{}{
				"snapshot_id":       *snapshot.SnapshotId,
				"snapshot_name":     *snapshot.SnapshotName,
				"storage_id":        *snapshot.DiskId,
				"storage_usage":     *snapshot.DiskUsage,
				"storage_size":      *snapshot.DiskSize,
				"availability_zone": *snapshot.Placement.Zone,
				"project_id":        *snapshot.Placement.ProjectId,
				"percent":           *snapshot.Percent,
				"create_time":       *snapshot.CreateTime,
				"encrypt":           *snapshot.Encrypt,
			}
			snapshotList = append(snapshotList, mapping)
			ids = append(ids, *snapshot.SnapshotId)
		}

		d.SetId(helper.DataResourceIdsHash(ids))
		if e = d.Set("snapshot_list", snapshotList); e != nil {
			log.Printf("[CRITAL]%s provider set snapshot list fail, reason:%s\n ", logId, e.Error())
			return resource.NonRetryableError(e)
		}

		output, ok := d.GetOk("result_output_file")
		if ok && output.(string) != "" {
			if e := writeToFile(output.(string), snapshotList); e != nil {
				return resource.NonRetryableError(e)
			}
		}

		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read cbs snapshots failed, reason:%s\n ", logId, err.Error())
		return err
	}

	return nil
}
