/*
Provide a resource to configure addon that kubernetes comes with.
Unlike tencentcloudenterprise_tke_kubernetes_addon which manages the full lifecycle (install/update/delete),
this resource only manages addon configuration (update). It will not install or delete the addon.

# Example Usage

```hcl

resource "tencentcloudenterprise_tke_kubernetes_addon_config" "example" {
  cluster_id = "cls-rkeuubqw"
  addon_name = "cluster-autoscaler"
  raw_values = jsonencode({
    extraArgs = {
      scale-down-enabled               = true
      max-empty-bulk-delete             = 11
      scale-down-delay-after-add        = "10mm"
      scale-down-unneeded-time          = "10mm"
      scale-down-utilization-threshold  = 0.005
      ignore-daemonsets-utilization     = false
      skip-nodes-with-local-storage     = true
      skip-nodes-with-system-pods       = true
    }
  })
}

```
*/
package tencentcloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_addon_config", CNDescription{
		TerraformTypeCN: "集群Addon配置",
		DescriptionCN:   "提供TKE集群Addon配置资源，用于管理已有Addon的配置参数，不会安装或卸载Addon。",
		AttributesCN: map[string]string{
			"cluster_id":    "集群ID",
			"addon_name":    "Addon名称",
			"addon_version": "Addon版本",
			"raw_values":    "Addon参数（JSON字符串，Provider将进行base64编码）",
			"phase":         "Addon状态",
			"reason":        "失败原因",
		},
	})
}

func resourceTencentCloudTkeKubernetesAddonConfig() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to configure addon that kubernetes comes with.",
		Create:      resourceTencentCloudTkeKubernetesAddonConfigCreate,
		Read:        resourceTencentCloudTkeKubernetesAddonConfigRead,
		Update:      resourceTencentCloudTkeKubernetesAddonConfigUpdate,
		Delete:      resourceTencentCloudTkeKubernetesAddonConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(3 * time.Minute),
			Update: schema.DefaultTimeout(3 * time.Minute),
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
				Description: "Version of addon.",
			},
			"raw_values": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      "Params of addon, base64 encoded json format.",
				DiffSuppressFunc: suppressJSONOrderDiff,
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

func resourceTencentCloudTkeKubernetesAddonConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon_config.create")()
	defer inconsistentCheck(d, meta)()

	clusterId := d.Get("cluster_id").(string)
	addonName := d.Get("addon_name").(string)

	d.SetId(strings.Join([]string{clusterId, addonName}, FILED_SP))
	return resourceTencentCloudTkeKubernetesAddonConfigUpdate(d, meta)
}

func resourceTencentCloudTkeKubernetesAddonConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon_config.read")()
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
		log.Printf("[WARN]%s resource `tke_kubernetes_addon_config` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
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
	if addon.Phase != nil {
		_ = d.Set("phase", addon.Phase)
	}
	if addon.Reason != nil {
		_ = d.Set("reason", addon.Reason)
	}

	return nil
}

func resourceTencentCloudTkeKubernetesAddonConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon_config.update")()
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
	mutableArgs := []string{"addon_version", "raw_values"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
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

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().UpdateAddon(request)
			if e != nil {
				return retryError(e)
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update kubernetes addon config failed, reason:%+v", logId, err)
			return err
		}

		if err := tkeAddonWaitReady(ctx, meta, clusterId, addonName, d.Timeout(schema.TimeoutUpdate)); err != nil {
			return err
		}
	}

	return resourceTencentCloudTkeKubernetesAddonConfigRead(d, meta)
}

func resourceTencentCloudTkeKubernetesAddonConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_addon_config.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

// suppressJSONOrderDiff compares two JSON strings and ignores ordering differences
func suppressJSONOrderDiff(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	if old == "" || new == "" {
		return false
	}

	var oldJSON, newJSON interface{}
	if err := json.Unmarshal([]byte(old), &oldJSON); err != nil {
		log.Printf("[WARN] Failed to unmarshal old value as JSON: %v", err)
		return old == new
	}
	if err := json.Unmarshal([]byte(new), &newJSON); err != nil {
		log.Printf("[WARN] Failed to unmarshal new value as JSON: %v", err)
		return old == new
	}

	return reflect.DeepEqual(oldJSON, newJSON)
}
