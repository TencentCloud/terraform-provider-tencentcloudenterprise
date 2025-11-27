/*
Provide a resource to create a bms instance

Example Usage
```hcl

	resource "tencentcloudenterprise_bms_instance" "tencentcloudenterprise_bms_instance-wuua" {
	  operating_system      = "your operating system name"
	  operating_system_type = "Linux"
	  password              = "your password"
	  subnet_id             = "your subnet id"
	  vpc_id                = "your vpc id"
	  flavor_id             = "your flavorid"
	  instance_name         = "your instance"
	  ipv6_address          = true
	  raid_type             = "NORAID"
	  availability_zone     = "your zone name"
	  security_service      = true
	  whistle_service       = true
	}

```

Import
Placement group can be imported using the id, e.g.

```
$ terraform import cloud_bms_placement_group.foo ps-ilan8vjf
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

	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bms_instance", CNDescription{
		TerraformTypeCN: "裸金属服务器",
		DescriptionCN:   "提供裸金属服务器实例资源，用于创建和管理裸金属服务器实例。",
		AttributesCN: map[string]string{
			"instance_id":                             "实例ID",
			"instance_name":                           "实例名称",
			"instance_type":                           "实例机型",
			"hostname":                                "主机名",
			"availability_zone":                       "可用区",
			"project_id":                              "项目ID",
			"instance_charge_type":                    "计费模式",
			"instance_charge_type_prepaid_period":     "预付费时长",
			"instance_charge_type_prepaid_renew_flag": "自动续费标识",
			"spot_instance_type":                      "竞价实例类型",
			"spot_max_price":                          "竞价实例最高价格",
			"internet_charge_type":                    "网络计费类型",
			"bandwidth_package_id":                    "共享带宽包ID",
			"internet_max_bandwidth_out":              "公网出带宽",
			"internet_service_provider":               "网络提供商",
			"ipv6_address":                            "是否使用IPV6",
			"allocate_public_ip":                      "是否分配公网IP",
			"vpc_id":                                  "VPC ID",
			"subnet_id":                               "子网ID",
			"private_ip":                              "私有IP",
			"security_groups":                         "安全组ID",
			"orderly_security_groups":                 "安全组ID",
			"system_disk_type":                        "系统盘类型",
			"system_disk_size":                        "系统盘大小",
			//"system_disk_id":                          "系统盘快照ID",
			"data_disks":       "数据盘",
			"security_service": "是否开启安全服务,默认False关闭",
			//"monitor_service":       "是否开启云监控服务，默认False关闭",
			"whistle_service":  "是否开启云哨/硬件监控服务，默认False关闭",
			"key_name":         "密钥对名称",
			"key_ids":          "密钥对ID",
			"password":         "密码",
			"keep_image_login": "保留镜像原始设置",
			"tags":             "标签",
			//"force_delete":          "是否强制删除",
			"instance_cound":        "实例数量",
			"raid_type":             "RAID类型",
			"flavor_id":             "实例规格",
			"operating_system":      "操作系统",
			"placement_group_id":    "置放群组ID",
			"operating_system_type": "操作系统类型",
			"instance_status":       "实例状态",
			"create_time":           "创建时间",
			// "instance_count":        "实例数量",
		},
	})
}
func resourceTencentCloudBmsInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a bms instance.",
		Create:      resourceTencentCloudBmsInstanceCreate,
		Read:        resourceTencentCloudBmsInstanceRead,
		Update:      resourceTencentCloudBmsInstanceUpdate,
		Delete:      resourceTencentCloudBmsInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			//"image_id": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	Description: "The image to use for the instance. Changing `image_id` will cause the instance reset.",
			//},
			"availability_zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The available zone for the bms instance.",
			},
			"flavor_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The flavor of the instance.",
			},

			//"instance_count": {
			//	Type:     schema.TypeInt,
			//	Optional: true,
			//	//Deprecated:   "It has been deprecated from version 1.59.18. Use built-in `count` instead.",
			//	// ValidateFunc: validateIntegerInRange(1, 100),
			//	ValidateFunc: validateIntegerInRange(1, 10),
			//	// Description:  "The number of instances to be purchased. Value range:[1,100]; default value: 1.",
			//	Description: "The number of instances to be purchased. Value range:[1,10]; default value: 1.",
			//},
			"instance_name": {
				Type:     schema.TypeString,
				Required: true,
				//Default:      "Terraform-bms-Instance",
				ValidateFunc: validateStringLengthInRange(2, 128),
				// Description:  "The name of the instance. The max length of instance_name is 60, and default value is `Terraform-bms-Instance`.",
				Description: "The name of the instance. The max length of instance_name is 60, and default value is `Unnamed`.",
			},
			"operating_system_type": {
				Type:     schema.TypeString,
				Required: true,
				//ValidateFunc: validateInstanceType,
				Description: "The operating system type of the instance.",
			},
			"operating_system": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The operate system of bms instance.",
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
				// Description: "The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.",
				Description: "The hostname of the instance. The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.",
			},
			// "project_id": {
			// 	Type:        schema.TypeInt,
			// 	Optional:    true,
			// 	Default:     0,
			// 	Description: "The project the instance belongs to, default to 0.",
			// },

			"placement_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of a placement group.",
			},
			"raid_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of raid.",
			},
			"internet_max_bandwidth_out": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false. The max values is 1000",
			},
			"internet_service_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The Internet Service Provider (ISP) associated with the instance.",
			},
			"allocate_public_ip": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.",
			},
			// vpc
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of a VPC network.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of a VPC subnet.",
			},
			//"private_ip": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "The private IP to be assigned to this instance, must be in the provided subnet and available.",
			//},
			"private_ip": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The private IP to be assigned to this instance, must be in the provided subnet and available.",
			},

			// storage
			//"system_disk_type": {
			//	Type:         schema.TypeString,
			//	Optional:     true,
			//	Default:      bms_DISK_TYPE_CLOUD_PREMIUM,
			//	ValidateFunc: validateAllowedStringValue(bms_DISK_TYPE),
			//	Description:  "System disk type. For more information on limits of system disk types, see [Storage Overview](https://intl.cloud.tencent.com/document/product/213/4952). Valid values: `LOCAL_BASIC`: local disk, `LOCAL_SSD`: local SSD disk, `CLOUD_SSD`: SSD, `CLOUD_PREMIUM`: Premium Cloud Storage, `CLOUD_BSSD`: Basic SSD. NOTE: If modified, the instance may force stop.",
			//},
			// "system_disk_size": {
			// 	Type:        schema.TypeInt,
			// 	Optional:    true,
			// 	Default:     50,
			// 	Description: "Size of the system disk. unit is GB, Default is 50GB. If modified, the instance may force stop.",
			// },
			//"system_disk_id": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "System disk snapshot ID used to initialize the system disk. When system disk type is `LOCAL_BASIC` and `LOCAL_SSD`, disk id is not supported.",
			//},
			// enhance services
			"security_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enhance service for security, it is disable by default.",
			},
			//"monitor_service": {
			//	Type:        schema.TypeBool,
			//	Optional:    true,
			//	Default:     false,
			//	Description: "Enhance service for monitor, it is disable by default.",
			//},
			"whistle_service": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enhance service for whistle, it is disable by default.",
			},

			// login
			/*
				"key_name": {
					Type:          schema.TypeString,
					Optional:      true,
					Computed:      true,
					Deprecated:    "Please use `key_ids` instead.",
					ConflictsWith: []string{"key_ids"},
					Description:   "The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.",
				},
			*/
			/*
				"key_ids": {
					Type:          schema.TypeSet,
					Optional:      true,
					Computed:      true,
					ConflictsWith: []string{"key_name"},
					Description:   "The key pair to use for the instance, it looks like `skey-16jig7tx`. Modifying will cause the instance reset.",
					Set:           schema.HashString,
					Elem:          &schema.Schema{Type: schema.TypeString},
				},
			*/
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				// Description: "Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.",
				Description: "Password for the instance. The password can only be set when creating a new instance. It is not modifiable after the instance is created.",
			},
			/*
				"keep_image_login": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
					DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						if new == "false" && old == "" || old == "false" && new == "" {
							return true
						} else {
							return old == new
						}
					},
					ConflictsWith: []string{"key_name", "key_ids"},
					Description:   "Whether to keep image login or not, default is `false`. When the image type is private or shared or imported, this parameter can be set `true`. Modifying will cause the instance reset.",
				},
			*/
			//"user_data": {
			//	Type:          schema.TypeString,
			//	Optional:      true,
			//	ForceNew:      true,
			//	ConflictsWith: []string{"user_data_raw"},
			//	Description:   "The user data to be injected into this instance. Must be base64 encoded and up to 16 KB.",
			//},
			//"user_data_raw": {
			//	Type:          schema.TypeString,
			//	Optional:      true,
			//	ForceNew:      true,
			//	ConflictsWith: []string{"user_data"},
			//	Description:   "The user data to be injected into this instance, in plain text. Conflicts with `user_data`. Up to 16 KB after base64 encoded.",
			//},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				// Description: "A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).",
				Description: "A mapping of tags to assign to the resource.",
			},
			// "force_delete": {
			// 	Type:        schema.TypeBool,
			// 	Optional:    true,
			// 	Default:     false,
			// 	Description: "Indicate whether to force delete the instance. Default is `false`. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for `PREPAID` instance.",
			// },

			// Computed values.
			"instance_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current status of the instance.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time of the instance.",
			},
			"ipv6_address": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to allocate an IPv6 address. Defaults to not allocating if not specified.",
			},
		},
	}
}

func resourceTencentCloudBmsInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_instance.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	request := bms.NewRunInstancesRequest()
	//request.ImageId = helper.String(d.Get("image_id").(string))
	request.Placement = &bms.Placement{
		Zone: helper.String(d.Get("availability_zone").(string)),
	}
	if v, ok := d.GetOk("project_id"); ok {
		projectId := int64(v.(int))
		request.Placement.ProjectId = &projectId
	}
	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}
	// if v, ok := d.GetOk("instance_count"); ok {
	// 	request.InstanceCount = helper.Int64(int64(v.(int)))
	// }
	//if v, ok := d.GetOk("instance_type"); ok {
	//	request.InstanceType = helper.String(v.(string))
	//}
	request.RaidType = helper.String(d.Get("raid_type").(string))
	request.FlavorId = helper.String(d.Get("flavor_id").(string))
	request.OperatingSystem = helper.String(d.Get("operating_system").(string))
	request.OperatingSystemType = helper.String(d.Get("operating_system_type").(string))
	if v, ok := d.GetOk("hostname"); ok {
		request.HostName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("placement_group_id"); ok {
		request.GroupId = helper.String(v.(string))
	}

	// network
	request.InternetAccessible = &bms.InternetAccessible{}
	if v, ok := d.GetOk("internet_max_bandwidth_out"); ok {
		maxBandwidthOut := int64(v.(int))
		request.InternetAccessible.InternetMaxBandwidthOut = &maxBandwidthOut
	}
	if v, ok := d.GetOk("internet_service_provider"); ok {
		request.InternetAccessible.InternetServiceProvider = helper.String(v.(string))
	}
	//if v, ok := d.GetOk("bandwidth_package_id"); ok {
	//	request.InternetAccessible.BandwidthPackageId = helper.String(v.(string))
	//}
	if v, ok := d.GetOkExists("allocate_public_ip"); ok {
		allocatePublicIp := v.(bool)
		request.InternetAccessible.PublicIpAssigned = &allocatePublicIp
	}

	// vpc
	if v, ok := d.GetOk("vpc_id"); ok {
		request.VirtualPrivateCloud = &bms.VirtualPrivateCloud{}
		request.VirtualPrivateCloud.VpcId = helper.String(v.(string))

		if v, ok = d.GetOk("ipv6_address"); ok {
			request.VirtualPrivateCloud.Ipv6Address = helper.Bool(v.(bool))
		}

		if v, ok = d.GetOk("subnet_id"); ok {
			request.VirtualPrivateCloud.SubnetId = helper.String(v.(string))
		}

		if v, ok = d.GetOk("private_ip"); ok {
			ips := v.(*schema.Set).List()
			for i := range ips {
				ip := ips[i].(string)
				request.VirtualPrivateCloud.PrivateIpAddresses = append(request.VirtualPrivateCloud.PrivateIpAddresses, helper.String(ip))
			}
		}
	}

	// enhanced service
	request.EnhancedService = &bms.EnhancedService{}
	if v, ok := d.GetOkExists("security_service"); ok {
		securityService := v.(bool)
		if securityService {
			request.EnhancedService.SecurityService = &securityService
		}
	}
	if v, ok := d.GetOkExists("monitor_service"); ok {
		monitorService := v.(bool)
		if monitorService {
			request.EnhancedService.MonitorService = &monitorService
		}
	}
	if v, ok := d.GetOkExists("whistle_service"); ok {
		whistleService := v.(bool)
		request.EnhancedService.WhistleService = &whistleService
	}
	/*
		if v, ok := d.GetOkExists("disable_whistle_service"); ok {
			monitorService := !(v.(bool))
			request.EnhancedService.MonitorService = &monitorService
		}
	*/

	// login
	request.LoginSettings = &bms.LoginSettings{}
	/*
		keyIds := d.Get("key_ids").(*schema.Set).List()
		if len(keyIds) > 0 {
			request.LoginSettings.KeyIds = helper.InterfacesStringsPoint(keyIds)
		} else if v, ok := d.GetOk("key_name"); ok {
			request.LoginSettings.KeyIds = []*string{helper.String(v.(string))}
		}
	*/
	if v, ok := d.GetOk("password"); ok {
		request.LoginSettings.Password = helper.String(v.(string))
	}
	/*
		v := d.Get("keep_image_login").(bool)
		if v {
			request.LoginSettings.KeepImageLogin = helper.String(BMS_IMAGE_LOGIN)
		} else {
			request.LoginSettings.KeepImageLogin = helper.String(BMS_IMAGE_LOGIN_NOT)
		}
	*/

	/*
		if v, ok := d.GetOkExists("disable_api_termination"); ok {
			request.DisableApiTermination = helper.Bool(v.(bool))
		}
	*/

	if v := helper.GetTags(d, "tags"); len(v) > 0 {
		tags := make([]*bms.Tag, 0)
		for tagKey, tagValue := range v {
			tag := bms.Tag{
				TagKey:   helper.String(tagKey),
				TagValue: helper.String(tagValue),
			}
			tags = append(tags, &tag)
		}
		request.Tags = tags
	}

	taskId := ""

	ratelimit.Check("create")
	response, err := meta.(*TencentCloudClient).apiV3Conn.UseBmsClient().RunInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		e, ok := err.(*sdkErrors.CloudSDKError)
		if ok && IsContains(BMS_RETRYABLE_ERROR, e.Code) {
			return fmt.Errorf("bms create error: %s, retrying", e.Error())
		}
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	if len(response.Response.TaskId) < 1 {
		err = fmt.Errorf("task id is nil")
		return err
	}

	if len(response.Response.BmsId) < 1 {
		err = fmt.Errorf("bms id is nil")
		return err
	}
	taskId = *response.Response.TaskId[0]
	bmsId := *response.Response.BmsId[0]

	time.Sleep(5 * time.Second)

	//get system disk ID and data disk ID
	var systemDiskId string
	var dataDiskIds []string
	err = resource.Retry(20*readRetryTimeout, func() *resource.RetryError {
		instance, errRet := bmsService.DescribeInstanceById(ctx, bmsId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance != nil {
			switch *instance.Status {
			case "CREATE_FAILED":
				return resource.NonRetryableError(fmt.Errorf("bms instance create failed,"+
					" status %s", *instance.Status))
			case "PENDING":
				log.Printf("[DEBUG] bms instance status is %s, retry...", *instance.Status)
			case "RUNNING":
				d.SetId(*instance.InstanceId)
				return nil
			}
		} else {
			return resource.NonRetryableError(fmt.Errorf("bms instance is nil"))
		}
		return resource.RetryableError(fmt.Errorf("bms instance status is %s, please retry", *instance.Status))
	})

	if err != nil {
		return err
	}

	// Wait for the tags attached to the vm since tags attachment it's async while vm creation.
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		resourceName := BuildTagResourceName("bms", "instance", tcClient.Region, taskId)
		if err = tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			// If tags attachment failed, the user will be notified, then plan/apply/update with terraform.
			return err
		}

		//except instance ,system disk and data disk will be tagged
		//keep logical consistence with the console
		//tag system disk
		if systemDiskId != "" {
			resourceName = BuildTagResourceName("bms", "volume", tcClient.Region, systemDiskId)
			if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
				// If tags attachment failed, the user will be notified, then plan/apply/update with terraform.
				return err
			}
		}
		//tag disk ids
		for _, diskId := range dataDiskIds {
			if diskId != "" {
				resourceName = BuildTagResourceName("bms", "volume", tcClient.Region, diskId)
				if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
					// If tags attachment failed, the user will be notified, then plan/apply/update with terraform.
					return err
				}
			}
		}
	}

	return resourceTencentCloudBmsInstanceRead(d, meta)
}

func resourceTencentCloudBmsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	// forceDelete := false
	// if v, ok := d.GetOkExists("force_delete"); ok {
	// 	forceDelete = v.(bool)
	// 	_ = d.Set("force_delete", forceDelete)
	// }

	client := meta.(*TencentCloudClient).apiV3Conn
	bmsService := BmsService{
		client: client,
	}
	var instance *bms.Instance
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, errRet = bmsService.DescribeInstanceById(ctx, d.Id())
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance != nil && instance.Status != nil && *instance.Status == BMS_STATUS_CREATE_FAILED {
			return resource.RetryableError(fmt.Errorf("waiting for instance %s operation", *instance.InstanceId))
		}
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(*instance.InstanceId)
	_ = d.Set("availability_zone", instance.Placement.Zone)
	_ = d.Set("instance_name", instance.InstanceName)
	// _ = d.Set("project_id", instance.Placement.ProjectId)
	_ = d.Set("vpc_id", instance.VirtualPrivateCloud.VpcId)
	_ = d.Set("subnet_id", instance.VirtualPrivateCloud.SubnetId)
	_ = d.Set("create_time", instance.CreatedTime)
	_ = d.Set("instance_status", instance.Status)
	//_ = d.Set("disable_api_termination", instance.DisableApiTermination)

	tagService := TagService{client}

	tags, err := tagService.DescribeResourceTags(ctx, "bms", "instance", client.Region, d.Id())
	if err != nil {
		return err
	}
	// as attachment add tencentcloud:autoscaling:auto-scaling-group-id tag automatically
	// we should remove this tag, otherwise it will cause terraform state change
	_ = d.Set("tags", tags)

	if len(instance.PrivateIpAddresses) > 0 {
		_ = d.Set("private_ip", instance.PrivateIpAddresses)
	}

	return nil
}

func resourceTencentCloudBmsInstanceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	defer logElapsed("resource.cloud_cvm_instance.update")()
	return nil
}

func resourceTencentCloudBmsInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_instance.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	instanceId := d.Id()
	//check is force delete or not
	//forceDelete := d.Get("force_delete").(bool)

	bmsService := BmsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := bmsService.DeleteInstance(ctx, instanceId)
		if errRet != nil {
			return retryError(errRet)
		}
		return nil
	})
	if err != nil {
		return err
	}

	//check recycling
	notExist := false

	//check exist
	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
		instance, errRet := bmsService.DescribeInstanceById(ctx, instanceId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance == nil {
			notExist = true
			return nil
		}
		return resource.RetryableError(fmt.Errorf("cvm instance status is %s, retry...", *instance.Status))
	})
	if err != nil {
		return err
	}

	if notExist {
		return nil
	}

	// exist in recycle, delete again
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		errRet := bmsService.DeleteInstance(ctx, instanceId)
		//when state is terminating, do not delete but check exist
		if errRet != nil {
			//check InvalidInstanceState.Terminating
			ee, ok := errRet.(*sdkErrors.CloudSDKError)
			if !ok {
				return retryError(errRet)
			}
			if ee.Code == "InvalidInstanceState.Terminating" {
				return nil
			}
			return retryError(errRet, "OperationDenied.InstanceOperationInProgress")
		}
		return nil
	})
	if err != nil {
		return err
	}

	//describe and check not exist
	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
		instance, errRet := bmsService.DescribeInstanceById(ctx, instanceId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if instance == nil {
			return nil
		}
		return resource.RetryableError(fmt.Errorf("cvm instance status is %s, retry...", *instance.Status))
	})
	if err != nil {
		return err
	}
	return nil
}
