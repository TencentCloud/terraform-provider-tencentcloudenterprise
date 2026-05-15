/*
Provides a resource to create an exclusive CLB Logset.

Example Usage

```hcl
resource "tencentcloudenterprise_clb_log_set" "foo" {
  period = 7
}
```

Import

CLB log set can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_logset.foo 4eb9e3a8-9c42-4b32-9ddf-e215e9c92764
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_log_set", CNDescription{
		TerraformTypeCN: "CLB 日志集",
		DescriptionCN:   "提供 CLB 专有日志集资源，用于创建和管理 CLB 的 CLS 日志集。",
		AttributesCN: map[string]string{
			"period":      "日志集保存时间（天），最大值为 90",
			"name":        "日志集名称，固定为 clb_logset，在所有 CLS 日志集中唯一",
			"create_time": "日志集创建时间",
			"topic_count": "日志集内日志主题数量",
		},
	})
}

func resourceTencentCloudClbLogSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbLogSetCreate,
		Read:   resourceTencentCloudClbLogSetRead,
		Delete: resourceTencentCloudClbLogSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Logset retention period in days. Maximun value is `90`.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Logset name, which unique and fixed `clb_logset` among all CLS logsets.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Logset creation time.",
			},
			"topic_count": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Number of log topics in logset.",
			},
		},
	}
}

func resourceTencentCloudClbLogSetRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_logset.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	info, err := service.DescribeClsLogset(ctx, id)

	if err != nil {
		return err
	}

	if info == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", info.LogsetName)
	_ = d.Set("create_time", info.CreateTime)
	if info.TopicCount != nil {
		_ = d.Set("topic_count", helper.Int64ToStr(*info.TopicCount))
	}

	return nil
}

func resourceTencentCloudClbLogSetCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_logset.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		period = d.Get("period").(int)
	)

	id, err := service.CreateClsLogset(ctx, "clb_logset", period)
	if err != nil {
		return err
	}
	if id == "" {
		return fmt.Errorf("[CRITAL]%s create CLB log set failed, logsetId is empty", logId)
	}
	// 创建保护
	time.Sleep(3 * time.Second)
	d.SetId(id)

	return resourceTencentCloudClbLogSetRead(d, meta)
}

func resourceTencentCloudClbLogSetDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_logset.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsLogsetById(ctx, id); err != nil {
		log.Printf("[CRITAL]%s delete CLB log set failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
