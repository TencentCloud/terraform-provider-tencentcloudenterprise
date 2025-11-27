package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cwp "terraform-provider-tencentcloudenterprise/sdk/cwp/v20180228"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cwp_license_order", CNDescription{
		TerraformTypeCN: "主机安全许可证订单",
		DescriptionCN:   "提供主机安全许可证订单资源，用于创建和管理主机安全许可证订单。",
		AttributesCN: map[string]string{
			"alias":        "资源别名",
			"license_type": "许可证类型",
			"license_num":  "许可证数量",
			"region_id":    "购买订单地域",
			"project_id":   "项目ID",
		},
	})
}

func ResourceTencentCloudCwpLicenseOrder() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudCwpLicenseOrderCreate,
		Read:        resourceTencentCloudCwpLicenseOrderRead,
		Update:      resourceTencentCloudCwpLicenseOrderUpdate,
		Delete:      resourceTencentCloudCwpLicenseOrderDelete,
		Description: "Provides a resource to create and manage CWP license order",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"alias": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Resource alias.",
			},
			"license_type": {
				Optional:     true,
				Type:         schema.TypeInt,
				Default:      LICENSE_TYPE_0,
				ValidateFunc: validateAllowedIntValue(LICENSE_TYPE),
				Description:  "LicenseType, 0 CWP Pro - Pay as you go, 1 CWP Pro - Monthly subscription, 2 CWP Ultimate - Monthly subscription. Default is 0.",
			},
			"license_num": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     1,
				Description: "License quantity, Quantity to be purchased.Default is 1.",
			},
			"region_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Purchase order region.",
			},
			"project_id": {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     0,
				Description: "Project ID. Default is 0.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tags of the license order.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "resource id.",
			},
			"license_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "license id.",
			},
		},
	}
}

func resourceTencentCloudCwpLicenseOrderCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cwp_license_order.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		request     = cwp.NewCreateLicenseOrderRequest()
		response    = cwp.NewCreateLicenseOrderResponse()
		tagService  = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId  string
		regionIdInt int
		licenseNum  *uint64
	)

	if v, ok := d.GetOkExists("license_type"); ok {
		request.LicenseType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("license_num"); ok {
		request.LicenseNum = helper.IntUint64(v.(int))
		licenseNum = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("region_id"); ok {
		request.RegionId = helper.IntUint64(v.(int))
		regionIdInt = v.(int)
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().CreateLicenseOrder(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || len(result.Response.ResourceIds) != 1 {
			e = fmt.Errorf("cwp licenseOrder not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cwp licenseOrder failed, reason:%+v", logId, err)
		return err
	}

	resourceId = *response.Response.ResourceIds[0]
	regionId := strconv.Itoa(regionIdInt)
	d.SetId(strings.Join([]string{resourceId, regionId}, FILED_SP))

	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cwp:%s:uin/:order/%s", region, resourceId)
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	// set alias
	if v, ok := d.GetOk("alias"); ok {
		aliasRequest := cwp.NewModifyLicenseOrderRequest()
		aliasRequest.Alias = helper.String(v.(string))
		aliasRequest.ResourceId = &resourceId
		aliasRequest.InquireNum = licenseNum

		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyLicenseOrder(aliasRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s set cwp licenseOrder alias failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cwp_license_order.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		tagService = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]
	regionId := idSplit[1]

	licenseOrder, err := service.DescribeCwpLicenseOrderById(ctx, resourceId)
	if err != nil {
		return err
	}

	if licenseOrder == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CwpLicenseOrder` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	regionIdInt, _ := strconv.Atoi(regionId)
	_ = d.Set("region_id", regionIdInt)

	if licenseOrder.Alias != nil {
		_ = d.Set("alias", licenseOrder.Alias)
	}

	if licenseOrder.ResourceId != nil {
		_ = d.Set("resource_id", licenseOrder.ResourceId)
	}

	if licenseOrder.LicenseId != nil {
		_ = d.Set("license_id", licenseOrder.LicenseId)
	}

	if licenseOrder.LicenseType != nil {
		_ = d.Set("license_type", licenseOrder.LicenseType)
	}

	if licenseOrder.LicenseCnt != nil {
		_ = d.Set("license_num", licenseOrder.LicenseCnt)
	}

	if licenseOrder.ProjectId != nil {
		_ = d.Set("project_id", licenseOrder.ProjectId)
	}

	region := meta.(*TencentCloudClient).apiV3Conn.Region
	tags, err := tagService.DescribeResourceTags(ctx, "cwp", "order", region, resourceId)
	if err != nil {
		return err
	}

	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudCwpLicenseOrderUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cwp_license_order.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		request    = cwp.NewModifyLicenseOrderRequest()
		tagService = TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	immutableArgs := []string{"license_type", "region_id"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]

	request.ResourceId = &resourceId

	if v, ok := d.GetOk("alias"); ok {
		request.Alias = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("license_num"); ok {
		request.InquireNum = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("project_id"); ok {
		request.ProjectId = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCwpClient().ModifyLicenseOrder(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update cwp licenseOrder failed, reason:%+v", logId, err)
		return err
	}

	if d.HasChange("tags") {
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := tagService.DiffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cwp", "order", "", resourceId)
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudCwpLicenseOrderRead(d, meta)
}

func resourceTencentCloudCwpLicenseOrderDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cwp_license_order.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = CwpService{client: meta.(*TencentCloudClient).apiV3Conn}
		licenseType *uint64
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	resourceId := idSplit[0]

	if v, ok := d.GetOkExists("license_type"); ok {
		licenseType = helper.IntUint64(v.(int))
	}

	if err := service.DeleteCwpLicenseOrderById(ctx, resourceId, licenseType); err != nil {
		return err
	}

	return nil
}
