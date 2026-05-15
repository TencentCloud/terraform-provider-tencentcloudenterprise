/*
Provides a resource to create a dc instance

# Example Usage

```hcl

	resource "tencentcloudenterprise_dc_instance" "instance" {
	  access_point_id         = "ap-shenzhen-b-ft"
		  bandwidth               = 10
		  customer_contact_number = "0"
		  direct_connect_name     = "terraform-for-test"
		  line_operator           = "In-houseWiring"
		  tencentcloudenterprise_port_type               = "10GBase-LR"
		  sign_law                = true
		  vlan                    = -1
		}

```

# Import

dc instance can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_dc_instance.instance dc_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	dc "terraform-provider-tencentcloudenterprise/sdk/dc/v20180410"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dc_instance", CNDescription{
		TerraformTypeCN: "专线接入实例",
		AttributesCN: map[string]string{
			"access_point_id":             "接入点ID",
			"bandwidth":                   "专线带宽",
			"customer_contact_number":     "联系人电话",
			"customer_contact_mail":       "联系人邮箱",
			"direct_connect_name":         "专线名称",
			"line_operator":               "运营商类型：ChinaTelecom（中国电信）、ChinaMobile（中国移动）、ChinaUnicom（中国联通）、In-houseWiring（自建线路）、ChinaOther（其他中国运营商）、InternationalOperator（国际运营商）",
			"tencentcloudenterprise_port_type":             "云端端口类型：100Base-T（百兆电口）、1000Base-T（千兆电口）、1000Base-LX（千兆单模光口10公里）、10GBase-T（万兆电口）、10GBase-LR（万兆单模光口10公里）",
			"idc_port_type":               "物理专线接入IDC侧端口类型,取值：100Base-T：百兆电口,1000Base-T（默认值）：千兆电口,1000Base-LX：千兆单模光口（10千米）,10GBase-T：万兆电口10GBase-LR：万兆单模光口（10千米），默认值，千兆单模光口（10千米）。",
			"idc_city":                    "本地数据中心所在城市",
			"location":                    "本地数据中心的地理位置。",
			"redundant_direct_connect_id": "冗余物理专线的ID。",
			"customer_name":               "物理专线申请者姓名。默认从账户体系获取。",
			"is_share":                    "是否共享该专线连接，仅可在修改时设置",
			"state":                       "专线状态",
			"created_time":                "创建时间",
			"enabled_time":                "开通时间",
			"fault_report_contact_person": "报障联系人",
			"fault_report_contact_number": "报障联系电话",
			"apply_id":                    "申请ID",
		},
	})
}
func resourceTencentCloudDcInstance() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dc instance",
		Create:      resourceTencentCloudDcInstanceCreate,
		Read:        resourceTencentCloudDcInstanceRead,
		Update:      resourceTencentCloudDcInstanceUpdate,
		Delete:      resourceTencentCloudDcInstanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"direct_connect_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Connection name.",
			},

			"access_point_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Access point of connection.The selected access point must exist and be available.",
			},

			"line_operator": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "ISP that provides connections.",
			},

			"tencentcloudenterprise_port_type": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Port type of connection. Valid values: 100Base-T (100-Megabit electrical Ethernet interface), 1000Base-T (1-Gigabit electrical Ethernet interface), 1000Base-LX (1-Gigabit single-module optical Ethernet interface; 10 KM), 10GBase-T (10-Gigabit electrical Ethernet interface), 10GBase-LR (10-Gigabit single-module optical Ethernet interface; 10 KM). Default value: 1000Base-LX.",
			},

			"location": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Local IDC location.",
			},

			"bandwidth": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Connection port bandwidth in Mbps. Value range: [2,10240]. Default value: 1000.",
			},

			"redundant_direct_connect_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "ID of redundant connection.",
			},

			"customer_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of connection applicant, which is obtained from the account system by default.",
			},

			"customer_contact_mail": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateEmail,
				Description:  "Email address of connection applicant, which is obtained from the account system by default.",
			},

			"customer_contact_number": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateInternationalPhone,
				Description:  "Contact number of connection applicant. Format：Area code: 1-3 digits, phone number: 5-15 digits (e.g.: 1-5551234567)",
			},

			"idc_port_type": {
				Required: true,
				Type:     schema.TypeString,
				ForceNew: true,
				Description: "IDC-side port type for physical dedicated line access. " +
					"Values: 100Base-T (100M electrical port), 1000Base-T (default, 1000M electrical port), " +
					"1000Base-LX (1000M single-mode optical port, 10km), 10GBase-T (10G electrical port), " +
					"10GBase-LR (10G single-mode optical port, 10km, default value).",
			},

			"idc_city": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "City where the local data center is located",
			},

			"is_share": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Whether the direct connect instance is shared. Can only be modified after creation via updates.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Direct connect state. Possible values: PENDING, REJECTED, ALLOCATED, AVAILABLE, DELETING, DELETED.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the direct connect.",
			},
			"enabled_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enabled time of the direct connect.",
			},
			"fault_report_contact_person": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fault report contact person.",
			},
			"fault_report_contact_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Fault report contact number.",
			},
			"apply_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Application ID of the direct connect.",
			},
		},
	}
}

func resourceTencentCloudDcInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_instance.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = dc.NewCreateDirectConnectRequest()
		response = dc.NewCreateDirectConnectResponse()
	)
	if v, ok := d.GetOk("direct_connect_name"); ok {
		request.DirectConnectName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("access_point_id"); ok {
		request.AccessPointId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("line_operator"); ok {
		request.LineOperator = helper.String(v.(string))
	}

	if v, ok := d.GetOk("tencentcloudenterprise_port_type"); ok {
		request.CloudPortType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("location"); ok {
		request.Location = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("bandwidth"); ok {
		request.Bandwidth = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("redundant_direct_connect_id"); ok {
		log.Printf("[DEBUG] redundant_direct_connect_id value: '%s'", v.(string))
		request.RedundantDirectConnectId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("customer_name"); ok {
		request.CustomerName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("customer_contact_mail"); ok {
		request.CustomerContactMail = helper.String(v.(string))
	}

	if v, ok := d.GetOk("customer_contact_number"); ok {
		request.CustomerContactNumber = helper.String(v.(string))
	}

	if v, ok := d.GetOk("idc_port_type"); ok {
		request.IdcPortType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("idc_city"); ok {
		request.IdcCity = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcClient().CreateDirectConnect(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create dc instance failed, reason:%+v", logId, err)
		return err
	}

	dcSet := response.Response.DirectConnectIdSet
	if len(dcSet) < 1 {
		return fmt.Errorf("create direct connect failed")
	}

	d.SetId(*dcSet[0])

	// CreateDirectConnect API does not support IsShare, so we need to call
	// ModifyDirectConnectAttribute to set it after creation.
	if v, ok := d.GetOkExists("is_share"); ok {
		modifyRequest := dc.NewModifyDirectConnectAttributeRequest()
		modifyRequest.DirectConnectId = helper.String(d.Id())
		modifyRequest.IsShare = helper.Bool(v.(bool))
		// ModifyDirectConnectAttribute requires these fields even for partial updates
		if name, ok := d.GetOk("direct_connect_name"); ok {
			modifyRequest.DirectConnectName = helper.String(name.(string))
		}
		if name, ok := d.GetOk("customer_name"); ok {
			modifyRequest.CustomerName = helper.String(name.(string))
		}
		if mail, ok := d.GetOk("customer_contact_mail"); ok {
			modifyRequest.CustomerContactMail = helper.String(mail.(string))
		}
		if phone, ok := d.GetOk("customer_contact_number"); ok {
			modifyRequest.CustomerContactNumber = helper.String(phone.(string))
		}
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcClient().ModifyDirectConnectAttribute(modifyRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, modifyRequest.GetAction(), modifyRequest.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s modify dc instance is_share after create failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudDcInstanceRead(d, meta)
}

func resourceTencentCloudDcInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_instance.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	directConnectId := d.Id()

	instances, err := service.DescribeDirectConnects(ctx, directConnectId, "")
	if err != nil {
		return err
	}

	if len(instances) < 1 {
		d.SetId("")
		log.Printf("[WARN]%s resource `DcInstance` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	instance := instances[0]

	if instance.DirectConnectName != nil {
		_ = d.Set("direct_connect_name", *instance.DirectConnectName)
	}

	if instance.AccessPointId != nil {
		_ = d.Set("access_point_id", *instance.AccessPointId)
	}

	if instance.LineOperator != nil {
		_ = d.Set("line_operator", *instance.LineOperator)
	}

	if instance.CloudPortType != nil {
		_ = d.Set("tencentcloudenterprise_port_type", *instance.CloudPortType)
	}

	if instance.Location != nil {
		_ = d.Set("location", *instance.Location)
	}

	if instance.Bandwidth != nil {
		_ = d.Set("bandwidth", *instance.Bandwidth)
	}

	if instance.RedundantDirectConnectId != nil {
		// Only set redundant_direct_connect_id in state if the user configured it.
		// The API may auto-populate this field on the primary DC when a redundant DC references it,
		// which would cause an unexpected diff and ForceNew recreation.
		if _, exists := d.GetOk("redundant_direct_connect_id"); exists {
			_ = d.Set("redundant_direct_connect_id", *instance.RedundantDirectConnectId)
		}
	}

	if instance.CustomerName != nil {
		_ = d.Set("customer_name", *instance.CustomerName)
	}

	if instance.CustomerContactMail != nil {
		_ = d.Set("customer_contact_mail", *instance.CustomerContactMail)
	}

	if instance.CustomerContactNumber != nil {
		_ = d.Set("customer_contact_number", *instance.CustomerContactNumber)
	}

	if instance.IdcPortType != nil {
		_ = d.Set("idc_port_type", *instance.IdcPortType)
	}

	if instance.IdcCity != nil {
		_ = d.Set("idc_city", *instance.IdcCity)
	}

	if instance.IsShare != nil {
		_ = d.Set("is_share", *instance.IsShare)
	}

	if instance.State != nil {
		_ = d.Set("state", *instance.State)
	}

	if instance.CreatedTime != nil {
		_ = d.Set("created_time", *instance.CreatedTime)
	}

	if instance.EnabledTime != nil {
		_ = d.Set("enabled_time", *instance.EnabledTime)
	}

	if instance.FaultReportContactPerson != nil {
		_ = d.Set("fault_report_contact_person", *instance.FaultReportContactPerson)
	}

	if instance.FaultReportContactNumber != nil {
		_ = d.Set("fault_report_contact_number", *instance.FaultReportContactNumber)
	}

	if instance.ApplyId != nil {
		_ = d.Set("apply_id", int(*instance.ApplyId))
	}

	return nil
}

func resourceTencentCloudDcInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_instance.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := dc.NewModifyDirectConnectAttributeRequest()

	directConnectId := d.Id()

	request.DirectConnectId = &directConnectId
	needChange := false

	immutableArgs := []string{
		"access_point_id", "line_operator", "tencentcloudenterprise_port_type", "location",
		"redundant_direct_connect_id", "idc_port_type", "idc_city",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	mutableArgs := []string{
		"direct_connect_name", "bandwidth", "customer_name",
		"customer_contact_mail", "customer_contact_number", "is_share",
	}

	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {

		if v, ok := d.GetOk("direct_connect_name"); ok {
			request.DirectConnectName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("bandwidth"); ok {
			request.Bandwidth = helper.IntUint64(v.(int))
		}

		if v, ok := d.GetOk("customer_name"); ok {
			request.CustomerName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("customer_contact_mail"); ok {
			request.CustomerContactMail = helper.String(v.(string))
		}

		if v, ok := d.GetOk("customer_contact_number"); ok {
			request.CustomerContactNumber = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("is_share"); ok {
			request.IsShare = helper.Bool(v.(bool))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcClient().ModifyDirectConnectAttribute(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update dc instance failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudDcInstanceRead(d, meta)
}

func resourceTencentCloudDcInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dc_instance.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}
	directConnectId := d.Id()

	if err := service.DeleteDcInstanceById(ctx, directConnectId); err != nil {
		return err
	}

	return nil
}
