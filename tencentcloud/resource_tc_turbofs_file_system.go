/*
Provides a resource to create a parallel file system(TurboFS).

# Example Usage

```hcl

	resource "tencentcloudenterprise_turbofs_file_system" "example" {
	  fs_name      = "tf-test"
	  zone         = "az"
	  p_group_id   = "pgroupbasic"
	  storage_type = "TP"
	  capacity     = 2560 # unit GiB
	  vpc_id       = "vpc-cvukkbpd"
	  subnet_id    = "subnet-kt2ffuim"
	  pool_id      = "pool-wzFg3qtSu"
	}

```

# Import

Cloud file system can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_file_system.foo turbofs-6hgquxmj
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_turbofs_file_system", CNDescription{
		TerraformTypeCN: "文件系统",
		DescriptionCN:   "提供TurboFS文件系统资源，用于创建和管理TurboFS文件系统。",
		AttributesCN: map[string]string{
			//"zone_id":                   "可用区 ID",
			"zone": "可用区名称，例如:ap-zone-1",
			//"creation_token":            "用户自定义文件系统名称，优先级低于fs_name",
			"protocol":                  "文件系统协议类型， 值为TURBO",
			"fs_name":                   "用户自定义文件系统名称,与CreationToken 两者必须填一项",
			"p_group_id":                "权限组ID",
			"cfs_version":               "CFS文件系统版本",
			"capacity":                  "TURBO文件系统的容量（GiB）",
			"pool_id":                   "资源池ID",
			"snapshot_id":               "快照ID",
			"storage_type":              "文件系统存储类型，值为 TB/TP ；其中 TB为标准型存储, TP为性能型存储",
			"net_interface":             "网络类型，值为 VPC，BASIC；其中 VPC 为私有网络，BASIC 为基础网络",
			"vpc_id":                    "私有网路（VPC） ID",
			"un_vpc_id":                 "系统分配的VPC统一ID",
			"subnet_id":                 "VPC子网ID",
			"un_subnet_id":              "系统分配的子网统一 ID",
			"mount_ip":                  "指定IP地址，仅VPC网络支持；若不填写、将在该子网下随机分配 IP",
			"storage_resource_pkg_id":   "文件系统绑定的存储包，每个文件系统只能绑定一个",
			"bandwidth_resource_pkg_id": "文件系统绑定的带宽包，每个文件系统只能绑定一个",
			"encrypted":                 "文件系统是否加密，若留空则默认为不加密",
			"kms_key_id":                "加密密钥 ID",
			"resource_tags":             "资源标签",
			"project_id":                "项目ID",
			"creation_time":             "文件系统创建时间",
			"cfs_id":                    "文件系统ID",
			"tag_id":                    "TurboFS资源池id",
			"file_system_id":            "文件系统 ID",
			"life_cycle_state":          "文件系统状态",
			"size_byte":                 "文件系统已使用容量大小",
			"tag_key":                   "标签键",
			"tag_value":                 "标签值",
		},
	})
}
func resourceTencentCloudTurbofsFileSystem() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a parallel file storage(TurboFS).",
		Create:      resourceTencentCloudTurbofsFileSystemCreate,
		Read:        resourceTencentCloudTurbofsFileSystemRead,
		Update:      resourceTencentCloudTurbofsFileSystemUpdate,
		Delete:      resourceTencentCloudTurbofsFileSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of a file system.",
			},
			//"creation_token": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "User-defined file system name, with lower priority than fs_name.",
			//},
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone that the file system locates at.",
			},
			//"zone_id": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	ForceNew:    true,
			//	Description: "The available zone id that the file system locates at.",
			//},
			"p_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of a permission group.",
			},
			//"protocol": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	ForceNew:    true,
			//	Description: "File service protocol. Valid value, the default is `TURBO`.",
			//},
			//"cfs_version": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "TurboFS file system version.",
			//},
			"capacity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Capacity of the TURBO file system in GiB. The value should be an integer multiple of the spec policy",
			},
			"pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ID of the resource pool. If not specified, the system will automatically select the default resource pool.",
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Snapshot ID.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "File service StorageType. Valid values are `TP` and `TB`.",
			},
			//"net_interface": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	ForceNew:    true,
			//	Description: "Network type. Valid value is `vpc`.",
			//},
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of a VPC network.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of a subnet.",
			},
			//"un_vpc_id": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Description: "Uniform id of a VPC network.",
			//},
			//"un_subnet_id": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Description: "Uniform id of a subnet.",
			//},
			"mount_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IP of mount point.",
			},
			"storage_resource_pkg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the storage resource package bound to the file system. Each file system can only be bound to one storage resource package.",
			},
			"bandwidth_resource_pkg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.",
			},
			"encrypted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the file system is encrypted.",
			},
			"kms_key_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the project.",
			},
			"tag_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TurboFS pool id.",
			},
			"resource_tags": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag key.",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Tag value.",
						},
					},
				},
				Optional:    true,
				Description: "TurboFS resource tag.",
			},
			// computed
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the file system.",
			},
			"file_system_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Id of the file system.",
			},
			"life_cycle_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Life cycle state of the file system.",
			},
			"size_byte": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Usage of the file system.",
			},
			"cfs_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Id of the file system",
			},
		},
	}
}

func resourceTencentCloudTurbofsFileSystemCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.create")()
	logId := getLogId(contextNil)
	request := turbofs.NewCreateCfsFileSystemRequest()
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	if v, ok := d.GetOk("zone"); ok {
		request.Zone = helper.String(v.(string))
	}
	if v, ok := d.GetOk("p_group_id"); ok {
		request.PGroupId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("storage_type"); ok {
		request.StorageType = helper.String(v.(string))
	}
	if v, ok := d.GetOk("fs_name"); ok {
		request.FsName = helper.String(v.(string))
	}
	//if v, ok := d.GetOk("creation_token"); ok {
	//	request.CreationToken = helper.String(v.(string))
	//}
	//if v, ok := d.GetOk("net_interface"); ok {
	//	request.NetInterface = helper.String(v.(string))
	//}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.UnVpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.UnSubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("capacity"); ok {
		capacity := v.(int)
		request.Capacity = helper.IntUint64(capacity)
	}

	//if v, ok := d.GetOk("zone_id"); ok {
	//	request.ZoneId = helper.IntUint64(v.(int))
	//}

	if v, ok := d.GetOk("pool_id"); ok {
		request.PoolId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("snapshot_id"); ok {
		request.SnapshotId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("mount_ip"); ok {
		request.MountIP = helper.String(v.(string))
	}

	if v, ok := d.GetOk("storage_resource_pkg_id"); ok {
		request.StorageResourcePkgId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bandwidth_resource_pkg_id"); ok {
		request.BandwidthResourcePkgId = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("encrypted"); ok {
		request.Encrypted = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("kms_key_id"); ok {
		request.KmsKeyId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("project_id"); ok {
		request.ProjectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_tags"); ok {
		request.ResourceTags = make([]*turbofs.ResourceTags, 0, len(v.(map[string]interface{})))
		for tagKey, tagValue := range v.(map[string]interface{}) {
			tag := turbofs.ResourceTags{
				TagKey:   helper.String(tagKey),
				TagValue: helper.String(tagValue.(string)),
			}
			request.ResourceTags = append(request.ResourceTags, &tag)
		}
	}

	request.Protocol = helper.String(TURBOFS_PROTCOL)
	request.CfsVersion = helper.String(TURBOFS_CFS_VERSION)
	request.NetInterface = helper.String("VPC")

	var fsId *string

	// Create file system
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		response, err := meta.(*TencentCloudClient).apiV3Conn.UseTurbofsClient().CreateCfsFileSystem(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			return retryError(err)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response.Response.FileSystemId == nil {
			err = fmt.Errorf("file system id is nil")
			return resource.NonRetryableError(err)
		}
		fsId = response.Response.FileSystemId
		return nil
	})
	if err != nil {
		return err
	}

	// Wait for file system to be available
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err = turbofsService.WaitForFileSystemAvailable(ctx, *fsId)
	if err != nil {
		return err
	}
	d.SetId(*fsId)

	return resourceTencentCloudTurbofsFileSystemRead(d, meta)
}

func resourceTencentCloudTurbofsFileSystemRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fsId := helper.String(d.Id())
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var fileSystem *turbofs.FileSystemInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		fileSystems, errRet := turbofsService.DescribeFileSystem(ctx, fsId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if len(fileSystems) > 0 {
			fileSystem = fileSystems[0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if fileSystem == nil {
		d.SetId("")
		return nil
	}

	// Set basic file system information with null checks
	if fileSystem.FsName != nil {
		_ = d.Set("fs_name", fileSystem.FsName)
	}
	if fileSystem.Zone != nil {
		_ = d.Set("zone", fileSystem.Zone)
	}
	if fileSystem.PGroup != nil && fileSystem.PGroup.PGroupId != nil {
		_ = d.Set("p_group_id", fileSystem.PGroup.PGroupId)
	}
	if fileSystem.SnapshotId != nil {
		_ = d.Set("snapshot_id", fileSystem.SnapshotId)
	}
	if fileSystem.CreationTime != nil {
		_ = d.Set("creation_time", fileSystem.CreationTime)
	}
	if fileSystem.StorageType != nil {
		_ = d.Set("storage_type", fileSystem.StorageType)
	}
	if fileSystem.Encrypted != nil {
		_ = d.Set("encrypted", fileSystem.Encrypted)
	}
	if fileSystem.KmsKeyId != nil {
		_ = d.Set("kms_key_id", fileSystem.KmsKeyId)
	}
	if fileSystem.FileSystemId != nil {
		_ = d.Set("file_system_id", fileSystem.FileSystemId)
	}
	if fileSystem.LifeCycleState != nil {
		_ = d.Set("life_cycle_state", fileSystem.LifeCycleState)
	}
	if fileSystem.SizeByte != nil {
		_ = d.Set("size_byte", fileSystem.SizeByte)
	}
	if fileSystem.SizeLimit != nil {
		_ = d.Set("capacity", fileSystem.SizeLimit)
	}
	if fileSystem.PoolId != nil {
		_ = d.Set("pool_id", fileSystem.PoolId)
	}

	if fileSystem.CfsId != nil {
		_ = d.Set("cfs_id", fileSystem.CfsId)
	}

	if fileSystem.ProjectId != nil {
		_ = d.Set("project_id", fileSystem.ProjectId)
	}

	if fileSystem.TagId != nil {
		_ = d.Set("tag_id", int(*fileSystem.TagId))
	}

	if fileSystem.Tags != nil && len(fileSystem.Tags) > 0 {
		resourceTags := make([]map[string]interface{}, 0, len(fileSystem.Tags))
		for _, tag := range fileSystem.Tags {
			if tag.TagKey != nil && tag.TagValue != nil {
				resourceTags = append(resourceTags, map[string]interface{}{
					"tag_key":   *tag.TagKey,
					"tag_value": *tag.TagValue,
				})
			}
		}
		if len(resourceTags) > 0 {
			_ = d.Set("resource_tags", resourceTags)
		}
	}

	// Get mount target information to supplement vpc_id, subnet_id, mount_ip
	mountTargets, err := turbofsService.DescribeMountTargets(ctx, *fsId)
	if err != nil {
		log.Printf("[DEBUG]%s describe mount targets fail, reason[%s]\n", logId, err.Error())
	} else if len(mountTargets) > 0 {
		mountTarget := mountTargets[0]

		if mountTarget.IpAddress != nil {
			_ = d.Set("mount_ip", mountTarget.IpAddress)
		}

		if mountTarget.VpcId != nil {
			_ = d.Set("vpc_id", mountTarget.VpcId)
		}

		if mountTarget.SubnetId != nil {
			_ = d.Set("subnet_id", mountTarget.SubnetId)
		}
	}

	return nil
}

func resourceTencentCloudTurbofsFileSystemUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	fsId := d.Id()
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	immutableArgs := []string{"vpc_id", "subnet_id", "pool_id", "storage_type", "zone"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	d.Partial(true)

	if d.HasChange("fs_name") {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			errRet := turbofsService.ModifyFileSystemName(ctx, fsId, d.Get("fs_name").(string))
			if errRet != nil {
				return retryError(errRet)
			}
			return nil
		})
		if err != nil {
			return err
		}

	}

	if d.HasChange("p_group_id") {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			errRet := turbofsService.ModifyFileSystemPGroup(ctx, fsId, d.Get("p_group_id").(string))
			if errRet != nil {
				return retryError(errRet)
			}
			return nil
		})
		if err != nil {
			return err
		}

	}

	if d.HasChange("capacity") {

		oldCapacity, newCapacity := d.GetChange("capacity")
		if newCapacity.(int) <= oldCapacity.(int) {
			return fmt.Errorf("new capacity %d must be greater than current capacity %d",
				newCapacity.(int), oldCapacity.(int))
		}

		targetCapacity := uint64(d.Get("capacity").(int))
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			errRet := turbofsService.ScaleUpFileSystem(ctx, fsId, targetCapacity)
			if errRet != nil {
				return retryError(errRet)
			}
			return nil
		})
		if err != nil {
			return err
		}

		// Wait for file system to be available after scaling
		err = turbofsService.WaitForFileSystemAvailable(ctx, fsId)
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceTencentCloudTurbofsFileSystemRead(d, meta)
}

func resourceTencentCloudTurbofsFileSystemDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_turbofs_file_system.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	fsId := d.Id()
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := turbofsService.DeleteFileSystem(ctx, fsId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
