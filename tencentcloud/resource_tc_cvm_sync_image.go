/*
Provides a resource to create a cvm sync_image

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_sync_image" "sync_image" {
	  image_id = "img-xxxxxx"
	  destination_regions =["ap-guangzhou", "ap-shanghai"]
	}

```
*/
package tencentcloud

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_sync_image", CNDescription{
		TerraformTypeCN: "同步镜像",
		DescriptionCN:   "提供同步镜像资源，用于同步镜像到其他地域。",
		AttributesCN: map[string]string{
			"image_id": "镜像ID指定的镜像必须满足以下条件：镜像状态为NORMAL",
			"destination_regions": "同步镜像的目的地区域列表限制：必须是有效的地域对于自定义镜像，目的地区域不能是源地域" +
				"对于共享镜像，目的地区域必须是源地域，表示在同一地域下创建镜像的副本作为自定义镜像",
		},
	})
}

func resourceTencentCloudCvmSyncImage() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a cvm sync_image",
		Create:      resourceTencentCloudCvmSyncImageCreate,
		Read:        resourceTencentCloudCvmSyncImageRead,
		Delete:      resourceTencentCloudCvmSyncImageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"image_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.",
			},

			"destination_regions": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.",
			},
		},
	}
}

func resourceTencentCloudCvmSyncImageCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_sync_image.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := cvm.NewSyncImagesRequest()
	imageId := d.Get("image_id").(string)
	request.ImageIds = []*string{&imageId}

	if v, ok := d.GetOk("destination_regions"); ok {
		destinationRegionsSet := v.(*schema.Set).List()
		for i := range destinationRegionsSet {
			destinationRegions := destinationRegionsSet[i].(string)
			request.DestinationRegions = append(request.DestinationRegions, &destinationRegions)
		}
	}

	/*
		if v, _ := d.GetOk("dry_run"); v != nil {
			request.DryRun = helper.Bool(v.(bool))
		}

		if v, ok := d.GetOk("image_name"); ok {
			request.ImageName = helper.String(v.(string))
		}

		if v, _ := d.GetOk("image_set_required"); v != nil {
			request.ImageSetRequired = helper.Bool(v.(bool))
		}

	*/

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().SyncImages(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate cvm syncImages failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(imageId)

	service := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}

	conf := BuildStateChangeConf([]string{}, []string{"NORMAL"}, 20*readRetryTimeout, time.Second, service.CvmSyncImagesStateRefreshFunc(d.Id(), []string{}))

	if _, e := conf.WaitForState(); e != nil {
		return e
	}

	return resourceTencentCloudCvmSyncImageRead(d, meta)
}

func resourceTencentCloudCvmSyncImageRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_sync_image.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudCvmSyncImageDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_sync_image.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
