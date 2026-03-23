/*
Provide a resource to manage TKE addons via InstallAddon/DescribeAddon/UpdateAddon/DeleteAddon.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_addon" "cbs" {
	  cluster_id = "cls-rkeuubqw"
	  addon_name = "cbs"

	  # raw_values accepts JSON and will be base64-encoded by the provider.
	  raw_values = jsonencode({
	    tolerations = [
	      {
	        key      = "123"
	        operator = "Exists"
	      }
	    ]
	  })

	  values = [
	    "global.image.host=ccr.d12-x86.fsphere.cn",
	    "global.cluster.id=cls-rkeuubqw",
	    "global.cluster.appid=1255000044",
	    "global.cluster.uin=110000000047",
	    "global.cluster.subuin=110000000370",
	    "global.cluster.type=tke",
	    "global.cluster.clustertype=INDEPENDENT_CLUSTER",
	    "global.cluster.kubeversion=1.22.5",
	    "global.cluster.kubeminor=22.0",
	    "cbs.url=cbs.api3.d12-x86.fsphere.cn",
	    "cvm.url=cvm.api3.d12-x86.fsphere.cn",
	    "cfs.url=cfs.api3.d12-x86.fsphere.cn",
	    "metadata.url=http://product-cvm-metadata.ap-qingyuan-region-devtest-ops.d12-x86.fsphere.cn/meta-data",
	  ]
	}

```
*/
package tencentcloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_addon", CNDescription{
		TerraformTypeCN: "集群Addon",
		DescriptionCN:   "提供TKE集群Addon资源，用于安装、更新和删除Addon。",
		AttributesCN: map[string]string{
			"cluster_id":    "集群ID",
			"addon_name":    "Addon名称",
			"addon_version": "Addon版本",
			"raw_values":    "Addon参数（JSON字符串，Provider将进行base64编码）",
			"values":        "Addon参数列表",
			"phase":         "Addon状态",
			"reason":        "失败原因",
		},
	})
}

func resourceTencentCloudTkeKubernetesAddon() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to manage TKE addons via InstallAddon/DescribeAddon/UpdateAddon/DeleteAddon.",
		Create:      resourceTencentCloudTkeKubernetesAddonCreate,
		Read:        resourceTencentCloudTkeKubernetesAddonRead,
		Update:      resourceTencentCloudTkeKubernetesAddonUpdate,
		Delete:      resourceTencentCloudTkeKubernetesAddonDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(3 * time.Minute),
			Update: schema.DefaultTimeout(3 * time.Minute),
			Delete: schema.DefaultTimeout(3 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of cluster.",
			},
			"addon_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of addon.",
			},
			"addon_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Version of addon. If no set, the latest version will be installed by default.",
			},
			"raw_values": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Addon params in JSON format. Provider will base64-encode before sending.",
			},
			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Addon params list.",
			},
			"phase": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of addon.",
			},
			"reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Reason of addon failed.",
			},
		},
	}
}

func resourceTencentCloudTkeKubernetesAddonCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Get("cluster_id").(string)
	addonName := d.Get("addon_name").(string)

	request := tke.NewInstallAddonRequest()
	request.ClusterId = helper.String(clusterId)
	request.AddonName = helper.String(addonName)

	if v, ok := d.GetOk("addon_version"); ok {
		request.AddonVersion = helper.String(v.(string))
	}
	if v, ok := d.GetOk("raw_values"); ok {
		rawValues := base64.StdEncoding.EncodeToString([]byte(v.(string)))
		request.RawValues = helper.String(rawValues)
	}
	if v, ok := d.GetOk("values"); ok {
		request.Values = tkeAddonExpandStringList(v.([]interface{}))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().InstallAddon(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{clusterId, addonName}, FILED_SP))

	if err := tkeAddonWaitReady(ctx, meta, clusterId, addonName, d.Timeout(schema.TimeoutCreate)); err != nil {
		return err
	}

	return resourceTencentCloudTkeKubernetesAddonRead(d, meta)
}

func resourceTencentCloudTkeKubernetesAddonRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	addonName := idSplit[1]

	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("addon_name", addonName)

	addon, err := tkeAddonDescribe(ctx, meta, clusterId, addonName)
	if err != nil {
		return err
	}
	if addon == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tke_kubernetes_addon` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if addon.AddonVersion != nil {
		_ = d.Set("addon_version", addon.AddonVersion)
	}
	if addon.RawValues != nil {
		rawValues, err := base64.StdEncoding.DecodeString(*addon.RawValues)
		if err == nil {
			_ = d.Set("raw_values", string(rawValues))
		}
	}
	if addon.Values != nil {
		values := make([]string, 0, len(addon.Values))
		for _, v := range addon.Values {
			if v != nil {
				values = append(values, *v)
			}
		}
		_ = d.Set("values", values)
	}
	if addon.Phase != nil {
		_ = d.Set("phase", addon.Phase)
	}
	if addon.Reason != nil {
		_ = d.Set("reason", addon.Reason)
	}

	return nil
}

func resourceTencentCloudTkeKubernetesAddonUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	addonName := idSplit[1]

	needChange := false
	mutableArgs := []string{"addon_version", "raw_values", "values"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}
	if !needChange {
		return resourceTencentCloudTkeKubernetesAddonRead(d, meta)
	}

	request := tke.NewUpdateAddonRequest()
	request.ClusterId = helper.String(clusterId)
	request.AddonName = helper.String(addonName)

	if v, ok := d.GetOk("addon_version"); ok {
		request.AddonVersion = helper.String(v.(string))
	}
	if v, ok := d.GetOk("raw_values"); ok {
		rawValues := base64.StdEncoding.EncodeToString([]byte(v.(string)))
		request.RawValues = helper.String(rawValues)
	}
	if v, ok := d.GetOk("values"); ok {
		request.Values = tkeAddonExpandStringList(v.([]interface{}))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().UpdateAddon(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	if err := tkeAddonWaitReady(ctx, meta, clusterId, addonName, d.Timeout(schema.TimeoutUpdate)); err != nil {
		return err
	}

	return resourceTencentCloudTkeKubernetesAddonRead(d, meta)
}

func resourceTencentCloudTkeKubernetesAddonDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	clusterId := idSplit[0]
	addonName := idSplit[1]

	request := tke.NewDeleteAddonRequest()
	request.ClusterId = helper.String(clusterId)
	request.AddonName = helper.String(addonName)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DeleteAddon(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	waitRequest := tke.NewDescribeAddonRequest()
	waitRequest.ClusterId = helper.String(clusterId)
	waitRequest.AddonName = helper.String(addonName)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DescribeAddon(waitRequest)
		if e != nil {
			if sdkErr, ok := e.(*sdkErrors.CloudSDKError); ok && sdkErr.Code == "ResourceNotFound" {
				return nil
			}
			return retryError(e)
		}
		if result == nil || result.Response == nil || result.Response.Addons == nil || len(result.Response.Addons) == 0 {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("deleting kubernetes addon. retry..."))
	})
	if err != nil {
		return err
	}

	return nil
}

func tkeAddonExpandStringList(raw []interface{}) []*string {
	if len(raw) == 0 {
		return nil
	}
	values := make([]*string, 0, len(raw))
	for _, v := range raw {
		if v == nil {
			continue
		}
		item := v.(string)
		values = append(values, helper.String(item))
	}
	return values
}

func tkeAddonDescribe(ctx context.Context, meta interface{}, clusterId, addonName string) (*tke.Addon, error) {
	request := tke.NewDescribeAddonRequest()
	request.ClusterId = helper.String(clusterId)
	request.AddonName = helper.String(addonName)

	var resp *tke.DescribeAddonResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DescribeAddon(request)
		if e != nil {
			return retryError(e)
		}
		resp = result
		return nil
	})
	if err != nil {
		return nil, err
	}
	if resp == nil || resp.Response == nil || resp.Response.Addons == nil || len(resp.Response.Addons) == 0 {
		return nil, nil
	}
	return resp.Response.Addons[0], nil
}

func tkeAddonWaitReady(ctx context.Context, meta interface{}, clusterId, addonName string, timeout time.Duration) error {
	request := tke.NewDescribeAddonRequest()
	request.ClusterId = helper.String(clusterId)
	request.AddonName = helper.String(addonName)

	return resource.Retry(timeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().DescribeAddon(request)
		if e != nil {
			return retryError(e)
		}
		if result == nil || result.Response == nil || result.Response.Addons == nil || len(result.Response.Addons) == 0 {
			return resource.NonRetryableError(fmt.Errorf("addons is nil"))
		}
		addon := result.Response.Addons[0]
		if addon.Phase == nil {
			return resource.NonRetryableError(fmt.Errorf("phase is nil"))
		}
		if *addon.Phase == "Succeeded" {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("addon status is %s", *addon.Phase))
	})
}
