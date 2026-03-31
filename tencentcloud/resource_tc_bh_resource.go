/*
Provide a resource to create a BH resource

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_resource" "example" {
	  deploy_region    = "ap-guangzhou"
	  vpc_id           = "vpc-xxxxxxxx"
	  subnet_id        = "subnet-xxxxxxxx"
	  resource_edition = "standard"
	  resource_node    = 50
	  time_unit        = "m"
	  time_span        = 1
	  pay_mode         = 1
	  auto_renew_flag  = 1
	  deploy_zone      = "ap-guangzhou-3"
	  cidr_block       = "10.0.0.0/24"
	  vpc_cidr_block   = "10.0.0.0/16"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_resource", CNDescription{
		TerraformTypeCN: "BH 资源",
		DescriptionCN:   "提供 BH 资源，用于创建和管理堡垒机实例资源。",
		AttributesCN: map[string]string{
			"deploy_region":    "部署地域",
			"vpc_id":           "部署堡垒机的VPC ID",
			"subnet_id":        "部署堡垒机的子网ID",
			"resource_edition": "资源类型",
			"resource_node":    "资源节点数",
			"time_unit":        "计费周期",
			"time_span":        "计费时长",
			"pay_mode":         "计费模式",
			"auto_renew_flag":  "自动续费",
			"deploy_zone":      "部署可用区",
			"trial":            "是否试用版",
			"share_clb":        "是否共享CLB",
			"cidr_block":       "堡垒机CIDR段",
			"vpc_cidr_block":   "VPC对应网段",
			"web_access":       "Web访问开关",
			"client_access":    "客户端访问开关",
			"intranet_access":  "内网访问开关",
			"external_access":  "外网访问开关",
			"resource_id":      "资源实例ID",
		},
	})
}

func ResourceTencentCloudBhResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhResourceCreate,
		Read:   resourceTencentCloudBhResourceRead,
		Update: resourceTencentCloudBhResourceUpdate,
		Delete: resourceTencentCloudBhResourceDelete,
		Schema: map[string]*schema.Schema{
			"deploy_region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Deployment region.",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC ID for deploying the bastion host.",
			},

			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Subnet ID for deploying the bastion host.",
			},

			"resource_edition": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource type. Values: standard/pro.",
			},

			"resource_node": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Number of resource nodes.",
			},

			"time_unit": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Billing cycle.",
			},

			"time_span": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Billing duration.",
			},

			"pay_mode": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Billing mode, 1 for prepaid.",
			},

			"auto_renew_flag": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Auto-renewal.",
			},

			"deploy_zone": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Deployment zone.",
			},

			"trial": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "0 for non-trial version, 1 for trial version.",
			},

			"share_clb": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Whether to share CLB, 0: not shared, 1: shared.",
			},

			// deploy params
			"cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CIDR block of the bastion host.",
			},

			"vpc_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The network segment corresponding to the VPC that needs to activate the service.",
			},

			"web_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "0 - Disable web access bastion host; 1 - Enable web access bastion host.",
			},

			"client_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "0 - Disable client access to the bastion host; 1 - Enable client access to the bastion host.",
			},

			"intranet_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "0 - Disable internal network access bastion host; 1 - Enable internal network access bastion host.",
			},

			"external_access": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "0 - Disable public network access to the bastion host; 1 - Enable public network access to the bastion host.",
			},

			// computed
			"resource_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource instance ID.",
			},
		},
	}
}

func resourceTencentCloudBhResourceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_resource.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = bhsaas.NewCreateResourceRequest()
		response   = bhsaas.NewCreateResourceResponse()
		resourceId string
	)

	if v, ok := d.GetOk("deploy_region"); ok {
		request.DeployRegion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("resource_edition"); ok {
		request.ResourceEdition = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("resource_node"); ok {
		request.ResourceNode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("time_unit"); ok {
		request.TimeUnit = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("time_span"); ok {
		request.TimeSpan = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("pay_mode"); ok {
		request.PayMode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("auto_renew_flag"); ok {
		request.AutoRenewFlag = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("deploy_zone"); ok {
		request.DeployZone = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("trial"); ok {
		request.Trial = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("share_clb"); ok {
		request.ShareClb = helper.IntUint64(v.(int))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create bh resource failed, Response is nil."))
		}

		response = result
		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create bh resource failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if response.Response.ResourceId == nil {
		return fmt.Errorf("ResourceId is nil.")
	}

	resourceId = *response.Response.ResourceId
	d.SetId(resourceId)

	// deploy
	deployReq := bhsaas.NewDeployResourceRequest()
	if v, ok := d.GetOk("deploy_region"); ok {
		deployReq.ApCode = helper.String(v.(string))
	}

	if v, ok := d.GetOk("deploy_zone"); ok {
		deployReq.Zone = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		deployReq.VpcId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		deployReq.SubnetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cidr_block"); ok {
		deployReq.CidrBlock = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpc_cidr_block"); ok {
		deployReq.VpcCidrBlock = helper.String(v.(string))
	}

	deployReq.ResourceId = &resourceId
	reqErr = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DeployResource(deployReq)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, deployReq.GetAction(), deployReq.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s deploy bh resource failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	// wait
	waitReq := bhsaas.NewDescribeResourcesRequest()
	waitReq.ResourceIds = []*string{&resourceId}
	reqErr = resource.Retry(readRetryTimeout*7, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeResources(waitReq)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.ResourceSet == nil || len(result.Response.ResourceSet) == 0 {
			return resource.NonRetryableError(fmt.Errorf("Describe bh resource failed, Response is nil."))
		}

		resourceStatus := result.Response.ResourceSet[0].Status
		if resourceStatus == nil {
			return resource.NonRetryableError(fmt.Errorf("Status is nil."))
		}

		if *resourceStatus == 1 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("Resource is still deploying...Status is %d", *resourceStatus))
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s deploy bh resource failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudBhResourceRead(d, meta)
}

func resourceTencentCloudBhResourceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_resource.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = BhService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId = d.Id()
	)

	respData, err := service.DescribeBhResourceById(ctx, resourceId)
	if err != nil {
		return err
	}

	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_bh_resource` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if respData.ApCode != nil {
		_ = d.Set("deploy_region", respData.ApCode)
	}

	if respData.VpcId != nil {
		_ = d.Set("vpc_id", respData.VpcId)
	}

	if respData.SubnetId != nil {
		_ = d.Set("subnet_id", respData.SubnetId)
	}

	if respData.Nodes != nil {
		_ = d.Set("resource_node", respData.Nodes)
	}

	if respData.RenewFlag != nil {
		_ = d.Set("auto_renew_flag", respData.RenewFlag)
	}

	if respData.Zone != nil {
		_ = d.Set("deploy_zone", respData.Zone)
	}

	if respData.Trial != nil {
		_ = d.Set("trial", respData.Trial)
	}

	if respData.ShareClb != nil {
		if *respData.ShareClb {
			_ = d.Set("share_clb", 1)
		} else {
			_ = d.Set("share_clb", 0)
		}
	}

	if respData.CidrBlock != nil {
		_ = d.Set("cidr_block", respData.CidrBlock)
	}

	if respData.VpcCidrBlock != nil {
		_ = d.Set("vpc_cidr_block", respData.VpcCidrBlock)
	}

	if respData.WebAccess != nil {
		_ = d.Set("web_access", respData.WebAccess)
	}

	if respData.WebAccess != nil {
		_ = d.Set("web_access", respData.WebAccess)
	}

	if respData.ClientAccess != nil {
		_ = d.Set("client_access", respData.ClientAccess)
	}

	if respData.IntranetAccess != nil {
		_ = d.Set("intranet_access", respData.IntranetAccess)
	}

	if respData.ExternalAccess != nil {
		_ = d.Set("external_access", respData.ExternalAccess)
	}

	return nil
}

func resourceTencentCloudBhResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_resource.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		resourceId = d.Id()
	)

	needChange := false
	mutableArgs := []string{"resource_edition", "resource_node", "auto_renew_flag"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := bhsaas.NewModifyResourceRequest()
		if v, ok := d.GetOk("resource_edition"); ok {
			request.ResourceEdition = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("resource_node"); ok {
			request.ResourceNode = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOkExists("auto_renew_flag"); ok {
			request.AutoRenewFlag = helper.IntInt64(v.(int))
		}

		request.ResourceId = helper.String(resourceId)
		reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyResource(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s update bh resource failed, reason:%+v", logId, reqErr)
			return reqErr
		}

		// wait
		waitReq := bhsaas.NewDescribeResourcesRequest()
		waitReq.ResourceIds = []*string{&resourceId}
		reqErr = resource.Retry(readRetryTimeout*7, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeResources(waitReq)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
			}

			if result == nil || result.Response == nil || result.Response.ResourceSet == nil || len(result.Response.ResourceSet) == 0 {
				return resource.NonRetryableError(fmt.Errorf("Describe bh resource failed, Response is nil."))
			}

			resourceStatus := result.Response.ResourceSet[0].Status
			if resourceStatus == nil {
				return resource.NonRetryableError(fmt.Errorf("Status is nil."))
			}

			if *resourceStatus == 1 {
				return nil
			}

			return resource.RetryableError(fmt.Errorf("Resource is still deploying...Status is %d", *resourceStatus))
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s deploy bh resource failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	if d.HasChange("intranet_access") {
		if v, ok := d.GetOkExists("intranet_access"); ok {
			if v.(int) == 1 {
				request := bhsaas.NewEnableIntranetAccessRequest()
				if v, ok := d.GetOk("vpc_id"); ok {
					request.VpcId = helper.String(v.(string))
				}

				if v, ok := d.GetOk("subnet_id"); ok {
					request.SubnetId = helper.String(v.(string))
				}

				if v, ok := d.GetOk("vpc_cidr_block"); ok {
					request.VpcCidrBlock = helper.String(v.(string))
				}

				request.ResourceId = helper.String(resourceId)
				reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
					result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().EnableIntranetAccess(request)
					if e != nil {
						return retryError(e)
					} else {
						log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
					}

					return nil
				})

				if reqErr != nil {
					log.Printf("[CRITAL]%s enable intranet access bh resource failed, reason:%+v", logId, reqErr)
					return reqErr
				}

				// wait
				waitReq := bhsaas.NewDescribeResourcesRequest()
				waitReq.ResourceIds = []*string{&resourceId}
				reqErr = resource.Retry(readRetryTimeout*7, func() *resource.RetryError {
					result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeResources(waitReq)
					if e != nil {
						return retryError(e)
					} else {
						log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
					}

					if result == nil || result.Response == nil || result.Response.ResourceSet == nil || len(result.Response.ResourceSet) == 0 {
						return resource.NonRetryableError(fmt.Errorf("Describe bh resource failed, Response is nil."))
					}

					intranetAccess := result.Response.ResourceSet[0].IntranetAccess
					if intranetAccess == nil {
						return resource.NonRetryableError(fmt.Errorf("IntranetAccess is nil."))
					}

					if *intranetAccess == 1 {
						return nil
					}

					return resource.RetryableError(fmt.Errorf("Enable intranet access is still running...Intranet access is %d", *intranetAccess))
				})

				if reqErr != nil {
					log.Printf("[CRITAL]%s enable intranet access bh resource failed, reason:%+v", logId, reqErr)
					return reqErr
				}
			} else {
				request := bhsaas.NewDisableIntranetAccessRequest()
				request.ResourceId = helper.String(resourceId)
				reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
					result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DisableIntranetAccess(request)
					if e != nil {
						return retryError(e)
					} else {
						log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
					}

					return nil
				})

				if reqErr != nil {
					log.Printf("[CRITAL]%s disable intranet access bh resource failed, reason:%+v", logId, reqErr)
					return reqErr
				}

				// wait
				waitReq := bhsaas.NewDescribeResourcesRequest()
				waitReq.ResourceIds = []*string{&resourceId}
				reqErr = resource.Retry(readRetryTimeout*7, func() *resource.RetryError {
					result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeResources(waitReq)
					if e != nil {
						return retryError(e)
					} else {
						log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, waitReq.GetAction(), waitReq.ToJsonString(), result.ToJsonString())
					}

					if result == nil || result.Response == nil || result.Response.ResourceSet == nil || len(result.Response.ResourceSet) == 0 {
						return resource.NonRetryableError(fmt.Errorf("Describe bh resource failed, Response is nil."))
					}

					intranetAccess := result.Response.ResourceSet[0].IntranetAccess
					if intranetAccess == nil {
						return resource.NonRetryableError(fmt.Errorf("IntranetAccess is nil."))
					}

					if *intranetAccess == 0 {
						return nil
					}

					return resource.RetryableError(fmt.Errorf("Disable intranet access is still running...Intranet access is %d", *intranetAccess))
				})

				if reqErr != nil {
					log.Printf("[CRITAL]%s disable intranet access bh resource failed, reason:%+v", logId, reqErr)
					return reqErr
				}
			}
		}

	}

	return resourceTencentCloudBhResourceRead(d, meta)
}

func resourceTencentCloudBhResourceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_resource.delete")()
	defer inconsistentCheck(d, meta)()

	return fmt.Errorf("cloud bh resource not supported delete, please contact the work order for processing")
}
