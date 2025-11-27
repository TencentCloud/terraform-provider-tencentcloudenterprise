/*
Provides a resource to create a cfs snapshot

# Example Usage

```hcl

	resource "tencentcloudenterprise_cfs_snapshot" "snapshot" {
	  file_system_id = "cfs-iobiaxtj"
	  snapshot_name = "test"
	  tags = {
	    "createdBy" = "terraform"
	  }
	}

```

# Import

cfs snapshot can be imported using the id, e.g.

```
terraform import cloud_cfs_snapshot.snapshot snapshot_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cfs_snapshot", CNDescription{
		TerraformTypeCN: "快照",
		DescriptionCN:   "提供CFS快照资源，用于创建和管理CFS快照。",
		AttributesCN: map[string]string{
			"file_system_id": "文件系统ID",
			"snapshot_name":  "快照名称",
			// "tags":           "标签",
			"job_id": "创建快照任务ID",
		},
	})
}
func resourceTencentCloudCfsSnapshot() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cfs snapshot.",
		Create:      resourceTencentCloudCfsSnapshotCreate,
		Read:        resourceTencentCloudCfsSnapshotRead,
		Update:      resourceTencentCloudCfsSnapshotUpdate,
		Delete:      resourceTencentCloudCfsSnapshotDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Id of file system.",
			},

			"snapshot_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of snapshot.",
			},

			"job_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Job id of create snapshot.",
			},
		},
	}
}

func resourceTencentCloudCfsSnapshotCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_snapshot.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = cfs.NewCreateCfsSnapRequest()
		response = cfs.NewCreateCfsSnapResponse()
		jobId    string
	)
	if v, ok := d.GetOk("file_system_id"); ok {
		request.FileSystemId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("snapshot_name"); ok {
		request.SnapName = helper.String(v.(string))
	}

	/*
		if v := helper.GetTags(d, "tags"); len(v) > 0 {
			for tagKey, tagValue := range v {
				tag := cfs.TagInfo{
					TagKey:   helper.String(tagKey),
					TagValue: helper.String(tagValue),
				}
				request.ResourceTags = append(request.ResourceTags, &tag)
			}
		}

	*/

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfsClient().CreateCfsSnap(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cfs snapshot failed, reason:%+v", logId, err)
		return err
	}
	if response == nil || response.Response == nil || response.Response.JobId == nil {
		return fmt.Errorf("response is nil")
	}

	jobId = *response.Response.JobId
	d.Set("job_id", jobId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"available"}, 2*readRetryTimeout, time.Second,
		service.CfsSnapshotStateRefreshFunc(jobId, []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	// ctx := context.WithValue(context.TODO(), logIdKey, logId)
	// if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
	// 	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	// 	region := meta.(*TencentCloudClient).apiV3Conn.Region
	// 	resourceName := fmt.Sprintf("qcs::cfs:%s:uin/:snap/%s", region, d.Id())
	// 	if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
	// 		return err
	// 	}
	// }

	return resourceTencentCloudCfsSnapshotRead(d, meta)
}

func resourceTencentCloudCfsSnapshotRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_snapshot.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}

	jobId := d.Get("job_id").(string)

	snapshot, err := service.DescribeCfsSnapshotById(ctx, jobId)
	if err != nil {
		return err
	}

	if snapshot == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CfsSnapshot` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if snapshot.FileSystemId != nil {
		_ = d.Set("file_system_id", snapshot.FileSystemId)
	}

	if snapshot.SnapName != nil {
		_ = d.Set("snapshot_name", snapshot.SnapName)
	}

	// tcClient := meta.(*TencentCloudClient).apiV3Conn
	// tagService := &TagService{client: tcClient}
	// tags, err := tagService.DescribeResourceTags(ctx, "cfs", "snap", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	// _ = d.Set("tags", tags)
	d.SetId(*snapshot.SnapId)

	return nil
}

func resourceTencentCloudCfsSnapshotUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_snapshot.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	needChange := false

	request := cfs.NewUpdateCfsSnapNameRequest()

	request.JobId = helper.String(d.Get("job_id").(string))

	if d.HasChange("snapshot_name") {
		needChange = true
		if v, ok := d.GetOk("snapshot_name"); ok {
			request.NewSnapName = helper.String(v.(string))
		}
	}

	if needChange {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCfsClient().UpdateCfsSnapName(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update cfs snapshot failed, reason:%+v", logId, err)
			return err
		}
		d.Set("job_id", *request.JobId)
	}

	// if d.HasChange("tags") {
	// 	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	// 	tcClient := meta.(*TencentCloudClient).apiV3Conn
	// 	tagService := &TagService{client: tcClient}
	// 	oldTags, newTags := d.GetChange("tags")
	// 	replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
	// 	resourceName := BuildTagResourceName("cfs", "snap", tcClient.Region, d.Id())
	// 	if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
	// 		return err
	// 	}
	// }

	return resourceTencentCloudCfsSnapshotRead(d, meta)
}

func resourceTencentCloudCfsSnapshotDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cfs_snapshot.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}
	snapshotId := d.Id()

	if err := service.DeleteCfsSnapshotById(ctx, snapshotId); err != nil {
		return err
	}

	err := resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
		instance, errRet := service.DescribeCfsSnapshotById(ctx, snapshotId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance == nil {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("cfs snapshot status is %v, retry...", *instance.Status))
	})
	if err != nil {
		return err
	}

	return nil
}
