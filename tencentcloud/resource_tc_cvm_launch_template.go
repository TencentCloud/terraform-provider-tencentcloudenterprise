/*
Provides a resource to create a cvm launch_template

# Example Usage

```hcl

	resource "tencentcloudenterprise_cvm_launch_template" "launch_template" {
	  launch_template_name = "test_launch_template"
	  placement {
	    project_id = 0
	    zone       = yfm18
	  }
	  image_id                            = "img-95xgn7er"
	  launch_template_version_description = "test111"
	  instance_type                       = "S5l.SMALL1"
	  system_disk {
	    disk_size = 50
	    disk_type = "CLOUD_BASIC"
	  }
	  virtual_private_cloud {
	    vpc_id             = "vpc-cs6ffr73"
	    subnet_id          = "subnet-38oi34ta"
	    as_vpc_gateway     = false
	    ipv6_address_count = 0
	  }
	  internet_accessible {
	    internet_charge_type       = "TRAFFIC_POSTPAID_BY_HOUR"
	    internet_max_bandwidth_out = 0
	    public_ip_assigned         = false
	  }
	  instance_count     = 1
	  instance_name      = "test_instance_name"
	  security_group_ids = ["sg-9s7k6qgw"]
	  enhanced_service {
	    security_service {
	      enabled = true
	    }
	    monitor_service {
	      enabled = true
	    }
	  }
	  instance_charge_type = "POSTPAID_BY_HOUR"
	}

```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cvm_launch_template", CNDescription{
		TerraformTypeCN: "云服务器启动模板",
		DescriptionCN:   "提供云服务器启动模板资源，用于创建和管理云服务器启动模板。",
		AttributesCN: map[string]string{
			"launch_template_id":                  "启动模板ID",
			"launch_template_name":                "启动模板名称",
			"launch_template_version":             "启动模板版本号",
			"latest_version_number":               "最新版本号",
			"create_time":                         "创建时间",
			"update_time":                         "更新时间",
			"created_by":                          "创建者",
			"default_version_number":              "默认版本号",
			"total_count":                         "启动模板总数",
			"security_group_ids":                  "安全组ID",
			"tags":                                "标签",
			"client_token":                        "客户端令牌",
			"hpc_cluster_id":                      "HPC集群ID",
			"instance_count":                      "实例数量",
			"user_data":                           "用户数据",
			"dry_run":                             "是否预检此请求",
			"instance_name":                       "实例名称",
			"disaster_recover_group_ids":          "容灾组ID",
			"instance_charge_type":                "实例计费类型",
			"instance_type":                       "实例类型",
			"image_id":                            "镜像ID",
			"launch_template_version_description": "启动模板版本描述",
			"host_name":                           "主机名",
			"cam_role_name":                       "CAM角色名称",
			"data_disks":                          "数据盘",
			"enhanced_service":                    "增强服务",
			"instance_charge_prepaid":             "预付费计费模式",
			"placement":                           "位置信息",
			"system_disk":                         "系统盘",
			"virtual_private_cloud":               "虚拟专用网络",
			"internet_accessible":                 "公网带宽",
			"action_timer":                        "定时任务",
			"tag_specification":                   "标签描述",
			"instance_market_options":             "市场选项",
			"login_settings":                      "登录设置",
			"action_time":                         "执行时间",
			"as_vpc_gateway":                      "它是否用作公共网络网关，TRUE或FALSE",
			"automation_service":                  "启用TencentCloud自动化工具（TAT）",
			"bandwidth_package_id":                "带宽包的ID",
			"delete_with_instance":                "数据磁盘是否随实例一起销毁，true或false",
			"disk_id":                             "数据磁盘ID",
			"disk_size":                           "数据磁盘的大小",
			"disk_type":                           "数据磁盘的类型",
			"enabled":                             "是否启用TencentCloud自动化工具（TAT），TRUE或FALSE",
			"encrypt":                             "数据磁盘是否加密，TRUE或FALSE",
			"externals":                           "扩展数据",
			"host_ids":                            "实例的CDH ID列表（输入）",
			"host_ips":                            "指定主机ip",
			"internet_charge_type":                "互联网收费的类型",
			"internet_max_bandwidth_out":          "互联网出站带宽上限，Mbps",
			"ipv6_address_count":                  "弹性网络接口的ipv6地址数",
			"keep_image_login":                    "保持镜像的原始设置",
			"key":                                 "标签的钥匙",
			"key_ids":                             "密钥ID列表",
			"kms_key_id":                          "自定义CMK的id",
			"market_type":                         "市场期权类型，目前仅支持价值现货",
			"max_price":                           "投标",
			"max_size":                            "HDD本地存储的最大容量",
			"min_size":                            "HDD本地存储的最小容量",
			"monitor_service":                     "启用云监控服务",
			"password":                            "实例的登录密码",
			"period":                              "采购实例的周期",
			"private_ip_addresses":                "私有ip地址",
			"project_id":                          "实例的项目ID",
			"public_ip_assigned":                  "是否分配公网IP，TRUE或FALSE",
			"release_address":                     "发布地址",
			"renew_flag":                          "自动续订标志",
			"resource_type":                       "资源的类型",
			"security_service":                    "启用云安全服务",
			"snapshot_id":                         "数据磁盘快照ID",
			"spot_instance_type":                  "招标请求类型，目前仅支持一次性类型",
			"spot_options":                        "与投标相关的选项",
			"storage_block_attr":                  "HDD本地存储属性",
			"subnet_id":                           "子网的id",
			"throughput_performance":              "云磁盘性能，MB/s",
			"timer_action":                        "定时器名称",
			"type":                                "HDD本地存储的类型",
			"unsupport_networks":                  "不支持的网络类型",
			"value":                               "标签的值",
			"vpc_id":                              "VPC的id",
			"zone":                                "实例的可用区域ID",
		},
	})
}
func resourceTencentCloudCvmLaunchTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Create cvm launch_template",
		Create:      resourceTencentCloudCvmLaunchTemplateCreate,
		Read:        resourceTencentCloudCvmLaunchTemplateRead,
		Delete:      resourceTencentCloudCvmLaunchTemplateDelete,
		Schema: map[string]*schema.Schema{
			"launch_template_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The name of launch template.",
			},

			"placement": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The location of instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The available zone ID of the instance.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The project ID of the instance.",
						},
						/*
							"host_ids": {
								Type: schema.TypeSet,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
								Optional:    true,
								Description: "The CDH ID list of the instance(input).",
							},
							"host_ips": {
								Type: schema.TypeSet,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
								Optional:    true,
								Description: "Specify the host machine ip.",
							},
						*/
					},
				},
			},

			"image_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Image ID.",
			},

			"launch_template_version_description": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Instance launch template version description.",
			},

			"instance_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The type of the instance. If this parameter is not specified, the system will dynamically specify the default model according to the resource sales in the current region.",
			},

			"system_disk": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "System disk configuration information of the instance. If this parameter is not specified, it is assigned according to the system default.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of system disk.",
						},
						"disk_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "System disk ID.",
						},
						"disk_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The size of system disk.",
						},
					},
				},
			},

			"data_disks": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "Data disk configuration information of the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_size": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The size of the data disk.",
						},
						"disk_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of data disk.",
						},
						"disk_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Data disk ID.",
						},
						"delete_with_instance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the data disk is destroyed along with the instance, true or false.",
						},
						"snapshot_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Data disk snapshot ID.",
						},
						"encrypt": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether the data disk is encrypted, TRUE or FALSE.",
						},
						"kms_key_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The id of custom CMK.",
						},
						"throughput_performance": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Cloud disk performance, MB/s.",
						},
					},
				},
			},

			"virtual_private_cloud": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The configuration information of VPC. If this parameter is not specified, the basic network is used by default.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of VPC.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of subnet.",
						},
						"as_vpc_gateway": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Is it used as a Public network gateway, TRUE or FALSE.",
						},
						"private_ip_addresses": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "The address of private ip.",
						},
						"ipv6_address_count": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The number of ipv6 addresses for Elastic Network Interface.",
						},
					},
				},
			},

			"internet_accessible": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The information settings of public network bandwidth. If you do not specify this parameter, the default Internet bandwidth is 0 Mbps.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internet_charge_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The type of internet charge.",
						},
						"internet_max_bandwidth_out": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Internet outbound bandwidth upper limit, Mbps.",
						},
						"public_ip_assigned": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to allocate public network IP, TRUE or FALSE.",
						},
						/*
							"bandwidth_package_id": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "The ID of bandwidth package.",
							},
						*/
					},
				},
			},

			"instance_count": {
				Optional:    true,
				ForceNew:    true,
				Default:     1,
				Type:        schema.TypeInt,
				Description: "The number of instances purchased.",
			},

			"instance_name": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The name of instance. If you do not specify an instance display name, 'Unnamed' is displayed by default.",
			},

			"login_settings": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The login settings of instance. By default, passwords are randomly generated and notified to users via internal messages.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The login password of instance.",
						},
						"key_ids": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "List of key ID.",
						},
						"keep_image_login": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Keep the original settings of the mirror.",
						},
					},
				},
			},

			"security_group_ids": {
				Optional: true,
				Computed: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The security group ID of instance. If this parameter is not specified, the default security group is bound.",
			},

			"enhanced_service": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Enhanced service. If this parameter is not specified, cloud monitoring and cloud security services will be enabled by default in public images.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_service": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Enable cloud security service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to enable cloud security service, TRUE or FALSE.",
									},
								},
							},
						},
						"monitor_service": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Enable cloud monitor service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to enable cloud monitor service, TRUE or FALSE.",
									},
								},
							},
						},
						"automation_service": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Enable TencentCloud Automation Tools(TAT).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to enable TencentCloud Automation Tools(TAT), TRUE or FALSE.",
									},
								},
							},
						},
					},
				},
			},

			"client_token": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "A string to used guarantee request idempotency.",
			},

			"host_name": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The host name of CVM.",
			},

			"action_timer": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Timed task.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timer_action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Timer name.",
						},
						"action_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Execution time.",
						},
						"externals": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Extended data.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"release_address": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Release address.",
									},
									"unsupport_networks": {
										Type: schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional:    true,
										Description: "Unsupported network type.",
									},
									"storage_block_attr": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "HDD local storage attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The type of HDD local storage.",
												},
												"min_size": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The minimum capacity of HDD local storage.",
												},
												"max_size": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "The maximum capacity of HDD local storage.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"disaster_recover_group_ids": {
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The ID of disaster recover group.",
			},

			"tag_specification": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Description: "Tag description list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of resource.",
						},
						"tags": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Tag list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The key of tag.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of tag.",
									},
								},
							},
						},
					},
				},
			},

			"instance_market_options": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The marketplace options of instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"spot_options": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Required:    true,
							Description: "Bidding related options.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_price": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Bidding.",
									},
									"spot_instance_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Bidding request type, currently only supported type: one-time.",
									},
								},
							},
						},
						"market_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Market option type, currently only supports value: spot.",
						},
					},
				},
			},

			"user_data": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The data of users.",
			},

			"dry_run": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Description: "Whether to preflight only this request, true or false.",
			},

			"cam_role_name": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The role name of CAM.",
			},

			"hpc_cluster_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The ID of HPC cluster.",
			},

			"instance_charge_type": {
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The charge type of instance. Default value: POSTPAID_BY_HOUR.",
			},

			"instance_charge_prepaid": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "The configuration of charge prepaid.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"period": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The period of purchasing instances.",
						},
						"renew_flag": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Automatic renew flag.",
						},
					},
				},
			},

			// "disable_api_termination": {
			// 	Optional:    true,
			// 	ForceNew:    true,
			// 	Type:        schema.TypeBool,
			// 	Description: "Instance destruction protection flag.",
			// },

			"tags": {
				Type:        schema.TypeMap,
				ForceNew:    true,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudCvmLaunchTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_launch_template.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request          = cvm.NewCreateLaunchTemplateRequest()
		response         = cvm.NewCreateLaunchTemplateResponse()
		launchTemplateId string
	)
	if v, ok := d.GetOk("launch_template_name"); ok {
		request.LaunchTemplateName = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "placement"); ok {
		placement := cvm.Placement{}
		if v, ok := helper.GetStringValue(dMap, "zone"); ok {
			placement.Zone = helper.String(v)
		}
		if v, ok := dMap["project_id"]; ok {
			placement.ProjectId = helper.IntInt64(v.(int))
		}
		if v, ok := helper.GetStringValue(dMap, "host_id"); ok {
			placement.HostId = helper.String(v)
		}
		request.Placement = &placement
	}

	if v, ok := d.GetOk("image_id"); ok {
		request.ImageId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("launch_template_version_description"); ok {
		request.LaunchTemplateVersionDescription = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_type"); ok {
		request.InstanceType = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "system_disk"); ok {
		systemDisk := cvm.SystemDisk{}
		if v, ok := helper.GetStringValue(dMap, "disk_type"); ok {
			systemDisk.DiskType = helper.String(v)
		}
		if v, ok := helper.GetStringValue(dMap, "disk_id"); ok {
			systemDisk.DiskId = helper.String(v)
		}
		if v, ok := dMap["disk_size"]; ok {
			systemDisk.DiskSize = helper.IntInt64(v.(int))
		}
		request.SystemDisk = &systemDisk
	}

	if v, ok := d.GetOk("data_disks"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			dataDisk := cvm.DataDisk{}
			if v, ok := dMap["disk_size"]; ok {
				dataDisk.DiskSize = helper.IntInt64(v.(int))
			}
			if v, ok := helper.GetStringValue(dMap, "disk_type"); ok {
				dataDisk.DiskType = helper.String(v)
			}
			if v, ok := helper.GetStringValue(dMap, "disk_id"); ok {
				dataDisk.DiskId = helper.String(v)
			}
			if v, ok := dMap["delete_with_instance"]; ok {
				dataDisk.DeleteWithInstance = helper.Bool(v.(bool))
			}
			if v, ok := helper.GetStringValue(dMap, "snapshot_id"); ok {
				dataDisk.SnapshotId = helper.String(v)
			}
			request.DataDisks = append(request.DataDisks, &dataDisk)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "virtual_private_cloud"); ok {
		virtualPrivateCloud := cvm.VirtualPrivateCloud{}
		if v, ok := helper.GetStringValue(dMap, "vpc_id"); ok {
			virtualPrivateCloud.VpcId = helper.String(v)
		}
		if v, ok := helper.GetStringValue(dMap, "subnet_id"); ok {
			virtualPrivateCloud.SubnetId = helper.String(v)
		}
		if v, ok := dMap["as_vpc_gateway"]; ok {
			virtualPrivateCloud.AsVpcGateway = helper.Bool(v.(bool))
		}
		if v, ok := dMap["private_ip_addresses"]; ok {
			privateIpAddressesSet := v.(*schema.Set).List()
			for i := range privateIpAddressesSet {
				privateIpAddresses := privateIpAddressesSet[i].(string)
				if privateIpAddresses == "" {
					continue
				}
				virtualPrivateCloud.PrivateIpAddresses = append(virtualPrivateCloud.PrivateIpAddresses, &privateIpAddresses)
			}
		}
		if v, ok := dMap["ipv6_address_count"]; ok {
			virtualPrivateCloud.Ipv6AddressCount = helper.IntInt64(v.(int))
		}
		request.VirtualPrivateCloud = &virtualPrivateCloud
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "internet_accessible"); ok {
		internetAccessible := cvm.InternetAccessible{}
		if v, ok := helper.GetStringValue(dMap, "internet_charge_type"); ok {
			internetAccessible.InternetChargeType = helper.String(v)
		}
		if v, ok := dMap["internet_max_bandwidth_out"]; ok {
			internetAccessible.InternetMaxBandwidthOut = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["public_ip_assigned"]; ok {
			internetAccessible.PublicIpAssigned = helper.Bool(v.(bool))
		}

		request.InternetAccessible = &internetAccessible
	}

	if v, _ := d.GetOk("instance_count"); v != nil {
		request.InstanceCount = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("instance_name"); ok {
		request.InstanceName = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "login_settings"); ok {
		loginSettings := cvm.LoginSettings{}
		if v, ok := helper.GetStringValue(dMap, "password"); ok {
			loginSettings.Password = helper.String(v)
		}
		if v, ok := dMap["key_ids"]; ok {
			keyIdsSet := v.(*schema.Set).List()
			for i := range keyIdsSet {
				keyIds := keyIdsSet[i].(string)
				loginSettings.KeyIds = append(loginSettings.KeyIds, &keyIds)
			}
		}
		if v, ok := helper.GetStringValue(dMap, "keep_image_login"); ok {
			loginSettings.KeepImageLogin = &v
		}
		request.LoginSettings = &loginSettings
	}

	if v, ok := d.GetOk("security_group_ids"); ok {
		securityGroupIdsSet := v.(*schema.Set).List()
		for i := range securityGroupIdsSet {
			securityGroupIds := securityGroupIdsSet[i].(string)
			request.SecurityGroupIds = append(request.SecurityGroupIds, &securityGroupIds)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "enhanced_service"); ok {
		enhancedService := cvm.EnhancedService{}
		if securityServiceMap, ok := helper.InterfaceToMap(dMap, "security_service"); ok {
			runSecurityServiceEnabled := cvm.RunSecurityServiceEnabled{}
			if v, ok := securityServiceMap["enabled"]; ok {
				runSecurityServiceEnabled.Enabled = helper.Bool(v.(bool))
			}
			enhancedService.SecurityService = &runSecurityServiceEnabled
		}
		if monitorServiceMap, ok := helper.InterfaceToMap(dMap, "monitor_service"); ok {
			runMonitorServiceEnabled := cvm.RunMonitorServiceEnabled{}
			if v, ok := monitorServiceMap["enabled"]; ok {
				runMonitorServiceEnabled.Enabled = helper.Bool(v.(bool))
			}
			enhancedService.MonitorService = &runMonitorServiceEnabled
		}

		request.EnhancedService = &enhancedService
	}

	if v, ok := d.GetOk("client_token"); ok {
		request.ClientToken = helper.String(v.(string))
	}

	if v, ok := d.GetOk("host_name"); ok {
		request.HostName = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "action_timer"); ok {
		// actionTimer := cvm.ActionTimer{}
		/*
			if v, ok := dMap["timer_action"]; ok {
				actionTimer.TimerAction = helper.String(v.(string))
			}
			if v, ok := dMap["action_time"]; ok {
				actionTimer.ActionTime = helper.String(v.(string))
			}
		*/
		if externalsMap, ok := helper.InterfaceToMap(dMap, "externals"); ok {
			externals := cvm.Externals{}
			if v, ok := externalsMap["release_address"]; ok {
				externals.ReleaseAddress = helper.Bool(v.(bool))
			}
			if v, ok := externalsMap["unsupport_networks"]; ok {
				unsupportNetworksSet := v.(*schema.Set).List()
				for i := range unsupportNetworksSet {
					unsupportNetworks := unsupportNetworksSet[i].(string)
					externals.UnsupportNetworks = append(externals.UnsupportNetworks, &unsupportNetworks)
				}
			}
			if storageBlockAttrMap, ok := helper.InterfaceToMap(externalsMap, "storage_block_attr"); ok {
				storageBlock := cvm.StorageBlock{}
				if v, ok := helper.GetStringValue(storageBlockAttrMap, "type"); ok {
					storageBlock.Type = helper.String(v)
				}
				if v, ok := storageBlockAttrMap["min_size"]; ok {
					storageBlock.MinSize = helper.IntInt64(v.(int))
				}
				if v, ok := storageBlockAttrMap["max_size"]; ok {
					storageBlock.MaxSize = helper.IntInt64(v.(int))
				}
				externals.StorageBlockAttr = &storageBlock
			}
			//actionTimer.Externals = &externals
		}
		//request.ActionTimer = &actionTimer
	}

	if v, ok := d.GetOk("disaster_recover_group_ids"); ok {
		disasterRecoverGroupIdsSet := v.(*schema.Set).List()
		for i := range disasterRecoverGroupIdsSet {
			disasterRecoverGroupIds := disasterRecoverGroupIdsSet[i].(string)
			request.DisasterRecoverGroupIds = append(request.DisasterRecoverGroupIds, &disasterRecoverGroupIds)
		}
	}

	if v, ok := d.GetOk("tag_specification"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			tagSpecification := cvm.TagSpecification{}
			if v, ok := helper.GetStringValue(dMap, "resource_type"); ok {
				tagSpecification.ResourceType = helper.String(v)
			}
			if v, ok := dMap["tags"]; ok {
				for _, item := range v.([]interface{}) {
					tagsMap := item.(map[string]interface{})
					tag := cvm.Tag{}
					if v, ok := tagsMap["key"]; ok {
						tag.Key = helper.String(v.(string))
					}
					if v, ok := tagsMap["value"]; ok {
						tag.Value = helper.String(v.(string))
					}
					tagSpecification.Tags = append(tagSpecification.Tags, &tag)
				}
			}
			request.TagSpecification = append(request.TagSpecification, &tagSpecification)
		}
	}

	/*
		if dMap, ok := helper.InterfacesHeadMap(d, "instance_market_options"); ok {
			instanceMarketOptionsRequest := cvm.InstanceMarketOptionsRequest{}
			if spotOptionsMap, ok := helper.InterfaceToMap(dMap, "spot_options"); ok {
				spotMarketOptions := cvm.SpotMarketOptions{}
				if v, ok := spotOptionsMap["max_price"]; ok {
					spotMarketOptions.MaxPrice = helper.String(v.(string))
				}
				if v, ok := spotOptionsMap["spot_instance_type"]; ok {
					spotMarketOptions.SpotInstanceType = helper.String(v.(string))
				}
				instanceMarketOptionsRequest.SpotOptions = &spotMarketOptions
			}
			if v, ok := dMap["market_type"]; ok {
				instanceMarketOptionsRequest.MarketType = helper.String(v.(string))
			}
			request.InstanceMarketOptions = &instanceMarketOptionsRequest
		}

	*/

	if v, ok := d.GetOk("user_data"); ok {
		request.UserData = helper.String(v.(string))
	}

	/*
		if v, _ := d.GetOk("dry_run"); v != nil {
			request.DryRun = helper.Bool(v.(bool))
		}
	*/

	if v, ok := d.GetOk("cam_role_name"); ok {
		request.CamRoleName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("hpc_cluster_id"); ok {
		request.HpcClusterId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("instance_charge_type"); ok {
		request.InstanceChargeType = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "instance_charge_prepaid"); ok {
		instanceChargePrepaid := cvm.InstanceChargePrepaid{}
		if v, ok := dMap["period"]; ok {
			instanceChargePrepaid.Period = helper.IntInt64(v.(int))
		}
		if v, ok := helper.GetStringValue(dMap, "renew_flag"); ok {
			instanceChargePrepaid.RenewFlag = helper.String(v)
		}
		request.InstanceChargePrepaid = &instanceChargePrepaid
	}

	//if v, _ := d.GetOk("disable_api_termination"); v != nil {
	//	request.DisableApiTermination = helper.Bool(v.(bool))
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().CreateLaunchTemplate(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cvm launchTemplate failed, reason:%+v", logId, err)
		return err
	}

	launchTemplateId = *response.Response.LaunchTemplateId
	d.SetId(launchTemplateId)

	return resourceTencentCloudCvmLaunchTemplateRead(d, meta)
}

func resourceTencentCloudCvmLaunchTemplateRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_launch_template.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}

	launchTemplateId := d.Id()

	launchTemplate, err := service.DescribeCvmLaunchTemplateById(ctx, launchTemplateId)
	if err != nil {
		return err
	}

	if launchTemplate == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CvmLaunchTemplate` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if launchTemplate.LaunchTemplateName != nil {
		_ = d.Set("launch_template_name", launchTemplate.LaunchTemplateName)
	}

	if launchTemplate.LaunchTemplateVersionData != nil && launchTemplate.LaunchTemplateVersionData.Placement != nil {
		placementMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.Placement.Zone != nil {
			placementMap["zone"] = launchTemplate.LaunchTemplateVersionData.Placement.Zone
		}

		if launchTemplate.LaunchTemplateVersionData.Placement.ProjectId != nil {
			placementMap["project_id"] = launchTemplate.LaunchTemplateVersionData.Placement.ProjectId
		}

		_ = d.Set("placement", []interface{}{placementMap})
	}

	if launchTemplate.LaunchTemplateVersionData != nil && launchTemplate.LaunchTemplateVersionData.ImageId != nil {
		_ = d.Set("image_id", launchTemplate.LaunchTemplateVersionData.ImageId)
	}

	if launchTemplate.LaunchTemplateVersionDescription != nil {
		_ = d.Set("launch_template_version_description", launchTemplate.LaunchTemplateVersionDescription)
	}

	if launchTemplate.LaunchTemplateVersionData.InstanceType != nil {
		_ = d.Set("instance_type", launchTemplate.LaunchTemplateVersionData.InstanceType)
	}

	if launchTemplate.LaunchTemplateVersionData != nil && launchTemplate.LaunchTemplateVersionData.SystemDisk != nil {
		systemDiskMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskType != nil {
			systemDiskMap["disk_type"] = launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskType
		}

		if launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskId != nil {
			systemDiskMap["disk_id"] = launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskId
		}

		if launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskSize != nil {
			systemDiskMap["disk_size"] = launchTemplate.LaunchTemplateVersionData.SystemDisk.DiskSize
		}

		//if launchTemplate.LaunchTemplateVersionData.SystemDisk.CdcId != nil {
		//	systemDiskMap["cdc_id"] = launchTemplate.LaunchTemplateVersionData.SystemDisk.CdcId
		//}

		_ = d.Set("system_disk", []interface{}{systemDiskMap})
	}

	if launchTemplate.LaunchTemplateVersionData.DataDisks != nil {
		dataDisksList := []interface{}{}
		for _, dataDisk := range launchTemplate.LaunchTemplateVersionData.DataDisks {
			dataDisksMap := map[string]interface{}{}

			if dataDisk.DiskSize != nil {
				dataDisksMap["disk_size"] = dataDisk.DiskSize
			}

			if dataDisk.DiskType != nil {
				dataDisksMap["disk_type"] = dataDisk.DiskType
			}

			if dataDisk.DiskId != nil {
				dataDisksMap["disk_id"] = dataDisk.DiskId
			}

			if dataDisk.DeleteWithInstance != nil {
				dataDisksMap["delete_with_instance"] = dataDisk.DeleteWithInstance
			}

			if dataDisk.SnapshotId != nil {
				dataDisksMap["snapshot_id"] = dataDisk.SnapshotId
			}

			/*
				if dataDisk.Encrypt != nil {
					dataDisksMap["encrypt"] = dataDisk.Encrypt
				}

				if dataDisk.KmsKeyId != nil {
					dataDisksMap["kms_key_id"] = dataDisk.KmsKeyId
				}

				if dataDisk.ThroughputPerformance != nil {
					dataDisksMap["throughput_performance"] = dataDisk.ThroughputPerformance
				}

				if dataDisk.CdcId != nil {
					dataDisksMap["cdc_id"] = dataDisk.CdcId
				}

			*/

			dataDisksList = append(dataDisksList, dataDisksMap)
		}

		_ = d.Set("data_disks", dataDisksList)

	}

	if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud != nil {
		virtualPrivateCloudMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.VpcId != nil {
			virtualPrivateCloudMap["vpc_id"] = launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.VpcId
		}

		if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.SubnetId != nil {
			virtualPrivateCloudMap["subnet_id"] = launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.SubnetId
		}

		if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.AsVpcGateway != nil {
			virtualPrivateCloudMap["as_vpc_gateway"] = launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.AsVpcGateway
		}

		if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.PrivateIpAddresses != nil {
			virtualPrivateCloudMap["private_ip_addresses"] = launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.PrivateIpAddresses
		}

		if launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.Ipv6AddressCount != nil {
			virtualPrivateCloudMap["ipv6_address_count"] = launchTemplate.LaunchTemplateVersionData.VirtualPrivateCloud.Ipv6AddressCount
		}

		_ = d.Set("virtual_private_cloud", []interface{}{virtualPrivateCloudMap})
	}

	if launchTemplate.LaunchTemplateVersionData.InternetAccessible != nil {
		internetAccessibleMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.InternetAccessible.InternetChargeType != nil {
			internetAccessibleMap["internet_charge_type"] = launchTemplate.LaunchTemplateVersionData.InternetAccessible.InternetChargeType
		}

		if launchTemplate.LaunchTemplateVersionData.InternetAccessible.InternetMaxBandwidthOut != nil {
			internetAccessibleMap["internet_max_bandwidth_out"] = launchTemplate.LaunchTemplateVersionData.InternetAccessible.InternetMaxBandwidthOut
		}

		if launchTemplate.LaunchTemplateVersionData.InternetAccessible.PublicIpAssigned != nil {
			internetAccessibleMap["public_ip_assigned"] = launchTemplate.LaunchTemplateVersionData.InternetAccessible.PublicIpAssigned
		}

		//if launchTemplate.LaunchTemplateVersionData.InternetAccessible.BandwidthPackageId != nil {
		//	internetAccessibleMap["bandwidth_package_id"] = launchTemplate.LaunchTemplateVersionData.InternetAccessible.BandwidthPackageId
		//}

		_ = d.Set("internet_accessible", []interface{}{internetAccessibleMap})
	}

	if launchTemplate.LaunchTemplateVersionData.InstanceCount != nil {
		_ = d.Set("instance_count", launchTemplate.LaunchTemplateVersionData.InstanceCount)
	}

	if launchTemplate.LaunchTemplateVersionData.InstanceName != nil {
		_ = d.Set("instance_name", launchTemplate.LaunchTemplateVersionData.InstanceName)
	}

	if launchTemplate.LaunchTemplateVersionData.LoginSettings != nil {
		loginSettingsMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.LoginSettings.Password != nil {
			loginSettingsMap["password"] = launchTemplate.LaunchTemplateVersionData.LoginSettings.Password
		}

		if launchTemplate.LaunchTemplateVersionData.LoginSettings.KeyIds != nil {
			loginSettingsMap["key_ids"] = launchTemplate.LaunchTemplateVersionData.LoginSettings.KeyIds
		}

		if launchTemplate.LaunchTemplateVersionData.LoginSettings.KeepImageLogin != nil {
			loginSettingsMap["keep_image_login"] = launchTemplate.LaunchTemplateVersionData.LoginSettings.KeepImageLogin
		}

		_ = d.Set("login_settings", []interface{}{loginSettingsMap})
	}

	if launchTemplate.LaunchTemplateVersionData.SecurityGroupIds != nil {
		_ = d.Set("security_group_ids", launchTemplate.LaunchTemplateVersionData.SecurityGroupIds)
	}

	if launchTemplate.LaunchTemplateVersionData.EnhancedService != nil {
		enhancedServiceMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.EnhancedService.SecurityService != nil {
			securityServiceMap := map[string]interface{}{}

			if launchTemplate.LaunchTemplateVersionData.EnhancedService.SecurityService.Enabled != nil {
				securityServiceMap["enabled"] = launchTemplate.LaunchTemplateVersionData.EnhancedService.SecurityService.Enabled
			}

			enhancedServiceMap["security_service"] = []interface{}{securityServiceMap}
		}

		if launchTemplate.LaunchTemplateVersionData.EnhancedService.MonitorService != nil {
			monitorServiceMap := map[string]interface{}{}

			if launchTemplate.LaunchTemplateVersionData.EnhancedService.MonitorService.Enabled != nil {
				monitorServiceMap["enabled"] = launchTemplate.LaunchTemplateVersionData.EnhancedService.MonitorService.Enabled
			}

			enhancedServiceMap["monitor_service"] = []interface{}{monitorServiceMap}
		}

		_ = d.Set("enhanced_service", []interface{}{enhancedServiceMap})
	}

	if launchTemplate.LaunchTemplateVersionData.ClientToken != nil {
		_ = d.Set("client_token", launchTemplate.LaunchTemplateVersionData.ClientToken)
	}

	if launchTemplate.LaunchTemplateVersionData.HostName != nil {
		_ = d.Set("host_name", launchTemplate.LaunchTemplateVersionData.HostName)
	}

	if launchTemplate.LaunchTemplateVersionData.ActionTimer != nil {
		actionTimerMap := map[string]interface{}{}

		if launchTemplate.LaunchTemplateVersionData.ActionTimer.TimerAction != nil {
			actionTimerMap["timer_action"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.TimerAction
		}

		if launchTemplate.LaunchTemplateVersionData.ActionTimer.ActionTime != nil {
			actionTimerMap["action_time"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.ActionTime
		}

		if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals != nil {
			externalsMap := map[string]interface{}{}

			if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.ReleaseAddress != nil {
				externalsMap["release_address"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.ReleaseAddress
			}

			if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.UnsupportNetworks != nil {
				externalsMap["unsupport_networks"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.UnsupportNetworks
			}

			if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr != nil {
				storageBlockAttrMap := map[string]interface{}{}

				if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.Type != nil {
					storageBlockAttrMap["type"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.Type
				}

				if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.MinSize != nil {
					storageBlockAttrMap["min_size"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.MinSize
				}

				if launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.MaxSize != nil {
					storageBlockAttrMap["max_size"] = launchTemplate.LaunchTemplateVersionData.ActionTimer.Externals.StorageBlockAttr.MaxSize
				}

				externalsMap["storage_block_attr"] = []interface{}{storageBlockAttrMap}
			}

			actionTimerMap["externals"] = []interface{}{externalsMap}
		}

		_ = d.Set("action_timer", []interface{}{actionTimerMap})
	}

	if launchTemplate.LaunchTemplateVersionData.DisasterRecoverGroupIds != nil {
		_ = d.Set("disaster_recover_group_ids", launchTemplate.LaunchTemplateVersionData.DisasterRecoverGroupIds)
	}

	if launchTemplate.LaunchTemplateVersionData.TagSpecification != nil {
		tagSpecificationList := []interface{}{}
		for _, tagSpecification := range launchTemplate.LaunchTemplateVersionData.TagSpecification {
			tagSpecificationMap := map[string]interface{}{}

			if tagSpecification.ResourceType != nil {
				tagSpecificationMap["resource_type"] = tagSpecification.ResourceType
			}

			if tagSpecification.Tags != nil {
				tagsList := []interface{}{}
				for _, tag := range tagSpecification.Tags {
					tagsMap := map[string]interface{}{}

					if tag.Key != nil {
						tagsMap["key"] = tag.Key
					}

					if tag.Value != nil {
						tagsMap["value"] = tag.Value
					}

					tagsList = append(tagsList, tagsMap)
				}

				tagSpecificationMap["tags"] = []interface{}{tagsList}
			}

			tagSpecificationList = append(tagSpecificationList, tagSpecificationMap)
		}

		_ = d.Set("tag_specification", tagSpecificationList)

	}

	if launchTemplate.LaunchTemplateVersionData.UserData != nil {
		_ = d.Set("user_data", launchTemplate.LaunchTemplateVersionData.UserData)
	}

	if launchTemplate.LaunchTemplateVersionData.CamRoleName != nil {
		_ = d.Set("cam_role_name", launchTemplate.LaunchTemplateVersionData.CamRoleName)
	}

	if launchTemplate.LaunchTemplateVersionData.HpcClusterId != nil {
		_ = d.Set("hpc_cluster_id", launchTemplate.LaunchTemplateVersionData.HpcClusterId)
	}

	if launchTemplate.LaunchTemplateVersionData.InstanceChargeType != nil {
		_ = d.Set("instance_charge_type", launchTemplate.LaunchTemplateVersionData.InstanceChargeType)
	}

	return nil
}

func resourceTencentCloudCvmLaunchTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cvm_launch_template.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CvmService{client: meta.(*TencentCloudClient).apiV3Conn}
	launchTemplateId := d.Id()

	if err := service.DeleteCvmLaunchTemplateById(ctx, launchTemplateId); err != nil {
		return err
	}

	return nil
}
