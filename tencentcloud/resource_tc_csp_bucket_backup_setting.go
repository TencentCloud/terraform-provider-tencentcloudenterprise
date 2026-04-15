/*
Provides a resource to manage CSP bucket backup setting.

# Example Usage

```hcl

	resource "tencentcloudenterprise_csp_bucket" "src" {
	  bucket = "src-bucket-1258798060"
	  acl    = "private"
	}

	resource "tencentcloudenterprise_csp_bucket_backup_setting" "example" {
	  bucket             = tencentcloudenterprise_csp_bucket.src.bucket
	  region             = "kazakhstan-1"
	  backup_enable      = true
	  backup_bucket_name = "dst-bucket-1258798060"
	  backup_endpoint    = "cos.kazakhstan-1.csp.example.com"
	  access_key         = "AKIDxxxxxxxxxxxxxxxx"
	  secret_key         = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	  backsource_enabled = false
	  use_https          = false
	}

```

# Import

csp bucket backup setting can be imported using the id (bucket#region), e.g.

```
$ terraform import tencentcloudenterprise_csp_bucket_backup_setting.example src-bucket-1258798060#kazakhstan-1
```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strings"

	csp "terraform-provider-tencentcloudenterprise/sdk/csp/v20200107"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_csp_bucket_backup_setting", CNDescription{
		TerraformTypeCN: "CSP存储桶备份设置",
		DescriptionCN:   "管理TCE CSP存储桶的外部备份配置，支持将桶数据备份到另一个存储桶，并可选开启回源功能。",
		AttributesCN: map[string]string{
			"bucket":             "源存储桶名称，格式为[自定义名称]-[appid]",
			"region":             "存储桶所在地域，如 kazakhstan-1",
			"backup_enable":      "是否启用备份",
			"backsource_enabled": "是否启用回源（备份桶作为回源源站）",
			"backup_bucket_name": "备份目标存储桶名称，格式为[自定义名称]-[appid]",
			"backup_endpoint":    "备份存储桶的访问域名后缀，例如 cos.ap-region.mydomain.com",
			"access_key":         "访问备份存储桶的 SecretId（AccessKey）",
			"secret_key":         "访问备份存储桶的 SecretKey",
			"use_https":          "是否使用 HTTPS 连接备份存储桶",
		},
	})
}

func resourceTencentCloudCspBucketBackupSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCspBucketBackupSettingCreate,
		Read:   resourceTencentCloudCspBucketBackupSettingRead,
		Update: resourceTencentCloudCspBucketBackupSettingUpdate,
		Delete: resourceTencentCloudCspBucketBackupSettingDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateCosBucketName,
				Description:  "Source bucket name. Format: [custom name]-[appid], e.g. `mybucket-1258798060`.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Region of the source bucket, e.g. `kazakhstan-1`.",
			},
			"backup_enable": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to enable bucket backup.",
			},
			"backsource_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable image origin-pull (back-to-source) from the backup bucket.",
			},
			"backup_bucket_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Backup destination bucket name. Required when `backup_enable` is true. Format: [custom name]-[appid].",
			},
			"backup_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "COS domain name suffix for accessing the backup bucket. Required when `backup_enable` is true. Example: `cos.ap-region.mydomain.com`.",
			},
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "SecretId (AccessKey) for accessing the backup bucket. Required when `backup_enable` is true.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "SecretKey for accessing the backup bucket. Required when `backup_enable` is true.",
			},
			"use_https": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to use HTTPS when connecting to the backup bucket.",
			},
		},
	}
}

func resourceTencentCloudCspBucketBackupSettingCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_backup_setting.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	bucket := d.Get("bucket").(string)
	region := d.Get("region").(string)

	if err := cspBucketBackupSettingSet(logId, d, meta); err != nil {
		return err
	}

	d.SetId(bucket + FILED_SP + region)
	log.Printf("[DEBUG]%s CSP bucket backup setting created, id: %s", logId, d.Id())

	// Server has eventual consistency: read back until BackupEnable matches what we set.
	wantEnable := d.Get("backup_enable").(bool)
	cspClient := meta.(*TencentCloudClient).apiV3Conn.UseCspClient()
	waitErr := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		req := csp.NewGetBackupSettingRequest()
		req.Bucket = helper.String(bucket)
		req.CosRegion = helper.String(region)
		resp, e := cspClient.GetBackupSetting(req)
		if e != nil {
			return retryError(e)
		}
		if resp.Response == nil || resp.Response.Response == nil || resp.Response.Response.BackupEnable == nil {
			return resource.RetryableError(fmt.Errorf("backup setting not ready yet"))
		}
		if *resp.Response.Response.BackupEnable != wantEnable {
			return resource.RetryableError(fmt.Errorf("backup_enable not yet %v, retrying", wantEnable))
		}
		return nil
	})
	if waitErr != nil {
		log.Printf("[WARN]%s wait for backup_enable=%v timed out: %v", logId, wantEnable, waitErr)
	}

	return resourceTencentCloudCspBucketBackupSettingRead(d, meta)
}

func resourceTencentCloudCspBucketBackupSettingRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_backup_setting.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	parts := strings.SplitN(d.Id(), FILED_SP, 2)
	if len(parts) != 2 {
		return nil
	}
	bucket, region := parts[0], parts[1]

	cspClient := meta.(*TencentCloudClient).apiV3Conn.UseCspClient()
	request := csp.NewGetBackupSettingRequest()
	request.Bucket = helper.String(bucket)
	request.CosRegion = helper.String(region)

	var resp *csp.GetBackupSettingResponse
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := cspClient.GetBackupSetting(request)
		if e != nil {
			return retryError(e)
		}
		resp = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CSP bucket backup setting failed, reason: %+v", logId, err)
		return err
	}

	if resp == nil || resp.Response == nil || resp.Response.Response == nil {
		d.SetId("")
		log.Printf("[WARN]%s CSP bucket backup setting [%s] not found", logId, d.Id())
		return nil
	}

	data := resp.Response.Response
	_ = d.Set("bucket", bucket)
	_ = d.Set("region", region)

	if data.BackupEnable != nil {
		_ = d.Set("backup_enable", *data.BackupEnable)
	}
	if data.BacksourceEnabled != nil {
		_ = d.Set("backsource_enabled", *data.BacksourceEnabled)
	}
	if data.BackupBucketName != nil {
		_ = d.Set("backup_bucket_name", *data.BackupBucketName)
	}
	if data.BackupEndpoint != nil {
		_ = d.Set("backup_endpoint", *data.BackupEndpoint)
	}
	if data.AccessKey != nil {
		_ = d.Set("access_key", *data.AccessKey)
	}
	if data.SecretKey != nil && *data.SecretKey != "" {
		_ = d.Set("secret_key", *data.SecretKey)
	}
	// Server may return null for UseHttps; treat null as false (default off)
	useHttps := data.UseHttps != nil && *data.UseHttps
	_ = d.Set("use_https", useHttps)

	return nil
}

func resourceTencentCloudCspBucketBackupSettingUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_backup_setting.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	if err := cspBucketBackupSettingSet(logId, d, meta); err != nil {
		return err
	}

	return resourceTencentCloudCspBucketBackupSettingRead(d, meta)
}

func resourceTencentCloudCspBucketBackupSettingDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_csp_bucket_backup_setting.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	parts := strings.SplitN(d.Id(), FILED_SP, 2)
	if len(parts) != 2 {
		return nil
	}
	bucket, region := parts[0], parts[1]

	cspClient := meta.(*TencentCloudClient).apiV3Conn.UseCspClient()
	request := csp.NewSetBackupSettingRequest()
	request.Bucket = helper.String(bucket)
	request.CosRegion = helper.String(region)
	request.BackupEnable = helper.Bool(false)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := cspClient.SetBackupSetting(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[SetBackupSetting] disable success, response: %s", logId, result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete (disable) CSP bucket backup setting failed, reason: %+v", logId, err)
		return err
	}

	return nil
}

func cspBucketBackupSettingSet(logId string, d *schema.ResourceData, meta interface{}) error {
	parts := strings.SplitN(d.Id(), FILED_SP, 2)
	var bucket, region string
	if len(parts) == 2 {
		bucket, region = parts[0], parts[1]
	} else {
		bucket = d.Get("bucket").(string)
		region = d.Get("region").(string)
	}

	cspClient := meta.(*TencentCloudClient).apiV3Conn.UseCspClient()
	request := csp.NewSetBackupSettingRequest()
	request.Bucket = helper.String(bucket)
	request.CosRegion = helper.String(region)
	request.BackupEnable = helper.Bool(d.Get("backup_enable").(bool))
	request.BacksourceEnabled = helper.Bool(d.Get("backsource_enabled").(bool))

	if v, ok := d.GetOk("backup_bucket_name"); ok {
		request.BackupBucketName = helper.String(v.(string))
	}
	if v, ok := d.GetOk("backup_endpoint"); ok {
		request.BackupEndpoint = helper.String(v.(string))
	}
	if v, ok := d.GetOk("access_key"); ok {
		request.AccessKey = helper.String(v.(string))
	}
	if v, ok := d.GetOk("secret_key"); ok {
		request.SecretKey = helper.String(v.(string))
	}
	request.UseHttps = helper.Bool(d.Get("use_https").(bool))

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := cspClient.SetBackupSetting(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[SetBackupSetting] success, request: %s, response: %s",
			logId, request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s SetBackupSetting failed, reason: %+v", logId, err)
		return err
	}
	return nil
}
