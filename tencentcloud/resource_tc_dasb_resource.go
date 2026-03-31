/*
Provide a resource to create a DASB resource

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_resource" "example" {
	  deploy_region    = "ap-guangzhou"
	  vpc_id           = "vpc-example"
	  subnet_id        = "subnet-example"
	  resource_edition = "standard"
	  resource_node    = 50
	  auto_renew_flag  = 1
	  deploy_zone      = "ap-guangzhou-1"
	  cidr_block       = "10.0.0.0/24"
	  vpc_cidr_block   = "10.0.0.0/16"
	}

```

# Import

DASB resource can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_resource.example resource-example
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_resource", CNDescription{
		TerraformTypeCN: "DASB 堡垒机资源",
		DescriptionCN:   "提供 DASB 堡垒机资源，用于创建和管理堡垒机服务实例。",
		AttributesCN: map[string]string{
			"deploy_region":     "部署地域",
			"vpc_id":            "VPC ID",
			"subnet_id":         "子网ID",
			"resource_edition":  "资源类型",
			"resource_node":     "资源节点数",
			"time_unit":         "计费周期",
			"time_span":         "计费时长",
			"auto_renew_flag":   "自动续费标识",
			"deploy_zone":       "部署可用区",
			"cidr_block":        "子网网段",
			"vpc_cidr_block":    "VPC网段",
			"package_bandwidth": "带宽扩展包",
		},
	})
}

func ResourceTencentCloudDasbResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbResourceCreate,
		Read:   resourceTencentCloudDasbResourceRead,
		Update: resourceTencentCloudDasbResourceUpdate,
		Delete: resourceTencentCloudDasbResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"deploy_region": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Deploy region.",
			},
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Deploy resource vpcId.",
			},
			"subnet_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Deploy resource subnetId.",
			},
			"resource_edition": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(RESOURCE_EDITION),
				Description:  "Resource type.Value:standard/pro.",
			},
			"resource_node": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Number of resource nodes.",
			},
			"time_unit": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Billing cycle, only support m: month. This field is mandatory, fill in m.",
			},
			"time_span": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Billing time. This field is mandatory, with a minimum value of 1.",
			},
			"auto_renew_flag": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "Automatic renewal. 1 is auto renew flag, 0 is not.",
			},
			"deploy_zone": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Deploy zone.",
			},
			"cidr_block": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Subnet segments that require service activation.",
			},
			"vpc_cidr_block": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The network segment corresponding to the VPC that requires service activation.",
			},
			"package_bandwidth": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of bandwidth expansion packets (4M), The set value is an integer multiple of 4.",
			},
			//"package_node": {
			//	Optional:    true,
			//	Computed:    true,
			//	Type:        schema.TypeInt,
			//	Description: "Number of authorized point extension packages (50 points). Cannot exceed 100.",
			//},
		},
	}
}

func resourceTencentCloudDasbResourceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_resource.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId           = getLogId(contextNil)
		request         = bhsaas.NewCreateResourceRequest()
		response        = bhsaas.NewCreateResourceResponse()
		deployRequest   = bhsaas.NewDeployResourceRequest()
		describeRequest = bhsaas.NewDescribeResourcesRequest()
		modifyRequest   = bhsaas.NewModifyResourceRequest()
		resourceId      string
		vpcId           string
		subnetId        string
		deployRegion    string
		deployZone      string
		cidrBlock       string
		vpcCidrBlock    string
	)

	if v, ok := d.GetOk("deploy_region"); ok {
		request.DeployRegion = helper.String(v.(string))
		deployRegion = v.(string)
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		request.VpcId = helper.String(v.(string))
		vpcId = v.(string)
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		request.SubnetId = helper.String(v.(string))
		subnetId = v.(string)
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

	request.PayMode = helper.IntInt64(1)

	if v, ok := d.GetOkExists("auto_renew_flag"); ok {
		request.AutoRenewFlag = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("deploy_zone"); ok {
		request.DeployZone = helper.String(v.(string))
		deployZone = v.(string)
	}

	if v, ok := d.GetOk("cidr_block"); ok {
		cidrBlock = v.(string)
	}

	if v, ok := d.GetOk("vpc_cidr_block"); ok {
		vpcCidrBlock = v.(string)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || *result.Response.ResourceId == "" {
			e = fmt.Errorf("dasb Resource not exists")
			return resource.NonRetryableError(e)
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb Resource failed, reason:%+v", logId, err)
		return err
	}

	resourceId = *response.Response.ResourceId
	d.SetId(resourceId)

	// deploy resource
	deployRequest.ResourceId = helper.String(resourceId)
	deployRequest.ApCode = helper.String(deployRegion)
	deployRequest.Zone = helper.String(deployZone)
	deployRequest.VpcId = helper.String(vpcId)
	deployRequest.SubnetId = helper.String(subnetId)
	deployRequest.CidrBlock = helper.String(cidrBlock)
	deployRequest.VpcCidrBlock = helper.String(vpcCidrBlock)

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DeployResource(deployRequest)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, deployRequest.GetAction(), deployRequest.ToJsonString(), deployRequest.ToJsonString())
		}

		if result == nil {
			e = fmt.Errorf("dasb Resource deploy error")
			return resource.NonRetryableError(e)
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s deploy dasb Resource failed, reason:%+v", logId, err)
		return err
	}

	// wait
	describeRequest.ResourceIds = helper.Strings([]string{resourceId})
	err = resource.Retry(writeRetryTimeout*6, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DescribeResources(describeRequest)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, describeRequest.GetAction(), describeRequest.ToJsonString(), result.ToJsonString())
		}

		if result == nil || len(result.Response.ResourceSet) != 1 {
			e = fmt.Errorf("dasb Resource not exists")
			return resource.NonRetryableError(e)
		}

		if *result.Response.ResourceSet[0].Status == 4 {
			e = fmt.Errorf("dasb Resource deploy error")
			return resource.NonRetryableError(e)
		}

		if *result.Response.ResourceSet[0].Status == 1 {
			return nil
		}

		return resource.RetryableError(fmt.Errorf("dasb Resource is still in running, state %d", *result.Response.ResourceSet[0].Status))
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb Resource failed, reason:%+v", logId, err)
		return err
	}

	// modify
	if v, ok := d.GetOkExists("package_bandwidth"); ok {
		modifyRequest.PackageBandwidth = helper.IntInt64(v.(int))
	}

	//if v, ok := d.GetOkExists("package_node"); ok {
	//	modifyRequest.PackageNode = helper.IntInt64(v.(int))
	//}

	if modifyRequest.PackageBandwidth != nil {
		modifyRequest.ResourceId = &resourceId
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyResource(modifyRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, modifyRequest.GetAction(), modifyRequest.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update dasb Resource failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudDasbResourceRead(d, meta)
}

func resourceTencentCloudDasbResourceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_resource.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceId = d.Id()
	)

	Resource, err := service.DescribeDasbResourceById(ctx, resourceId)
	if err != nil {
		return err
	}

	if Resource == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbResource` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if Resource.ApCode != nil {
		_ = d.Set("deploy_region", Resource.ApCode)
	}

	if Resource.VpcId != nil {
		_ = d.Set("vpc_id", Resource.VpcId)
	}

	if Resource.SubnetId != nil {
		_ = d.Set("subnet_id", Resource.SubnetId)
	}

	if Resource.SvArgs != nil {
		svArgs := strings.Split(*Resource.SvArgs, "_")
		var tmpStr string
		for _, item := range svArgs {
			if item == RESOURCE_EDITION_PRO || item == RESOURCE_EDITION_STANDARD {
				tmpStr = item
				break
			}
		}

		_ = d.Set("resource_edition", tmpStr)
	}

	if Resource.Nodes != nil {
		_ = d.Set("resource_node", Resource.Nodes)
	}

	if Resource.RenewFlag != nil {
		_ = d.Set("auto_renew_flag", Resource.RenewFlag)
	}

	if Resource.Zone != nil {
		_ = d.Set("deploy_zone", Resource.Zone)
	}

	if Resource.CidrBlock != nil {
		_ = d.Set("cidr_block", Resource.CidrBlock)
	}

	if Resource.VpcCidrBlock != nil {
		_ = d.Set("vpc_cidr_block", Resource.VpcCidrBlock)
	}

	if Resource.PackageBandwidth != nil {
		_ = d.Set("package_bandwidth", Resource.PackageBandwidth)
	}

	//if Resource.PackageNode != nil {
	//	_ = d.Set("package_node", Resource.PackageNode)
	//}

	return nil
}

func resourceTencentCloudDasbResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_resource.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = bhsaas.NewModifyResourceRequest()
		resourceId = d.Id()
	)

	immutableArgs := []string{"deploy_region", "vpc_id", "subnet_id", "time_unit", "time_span", "pay_mode", "deploy_zone", "cidr_block", "vpc_cidr_block"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.ResourceId = &resourceId
	if d.HasChange("resource_edition") {
		if v, ok := d.GetOk("resource_edition"); ok {
			request.ResourceEdition = helper.String(v.(string))
		}
	}

	if d.HasChange("resource_node") {
		if v, ok := d.GetOkExists("resource_node"); ok {
			request.ResourceNode = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("auto_renew_flag") {
		if v, ok := d.GetOkExists("auto_renew_flag"); ok {
			request.AutoRenewFlag = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("package_bandwidth") {
		if v, ok := d.GetOkExists("package_bandwidth"); ok {
			request.PackageBandwidth = helper.IntInt64(v.(int))
		}
	}

	//if d.HasChange("package_node") {
	//	if v, ok := d.GetOkExists("package_node"); ok {
	//		request.PackageNode = helper.IntInt64(v.(int))
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyResource(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb Resource failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbResourceRead(d, meta)
}

func resourceTencentCloudDasbResourceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_resource.delete")()
	defer inconsistentCheck(d, meta)()

	return fmt.Errorf("tencentcloud dasb resource not supported delete, please contact the work order for processing")
}
