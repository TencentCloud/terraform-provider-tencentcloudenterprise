/*
Provides a resource to create a dcdb db_parameters

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_db_parameters" "db_parameters" {
	  instance_id = "tdsqlshard-973xatu3"
	  params {
	    param = "audit_txsql_audit_mode"
	    value = "OFF"
	  }
	  params {
	    param = "audit_txsql_log_file_max_size"
	    value = "536870912"
	  }
	}

```

# Import

dcdb db_parameters can be imported using the id, e.g.

```
terraform import cloud_dcdb_db_parameters.db_parameters instanceId#paramName
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_db_parameters", CNDescription{
		TerraformTypeCN: "数据库实例参数",
		DescriptionCN:   "提供DCDB数据库实例参数资源，用于修改DCDB实例的配置参数。",
		AttributesCN: map[string]string{
			"instance_id": "实例的ID",
			"params":      "参数列表中，每个元素都是Param和Value的组合",
			"param":       "参数的名称",
			"参数的值":        "参数的值",
		},
	})
}

func resourceTencentCloudDcdbDbParameters() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb db_parameters",
		Create:      resourceTencentCloudDcdbDbParametersCreate,
		Read:        resourceTencentCloudDcdbDbParametersRead,
		Update:      resourceTencentCloudDcdbDbParametersUpdate,
		Delete:      resourceTencentCloudDcdbDbParametersDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"params": {
				Required: true,
				Type:     schema.TypeList,
				//MaxItems:    1,
				Description: "Parameter list, each element is a combination of Param and Value.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"param": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of parameter.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of parameter.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudDcdbDbParametersCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_db_parameters.create")()
	defer inconsistentCheck(d, meta)()

	return resourceTencentCloudDcdbDbParametersUpdate(d, meta)
}

func resourceTencentCloudDcdbDbParametersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_db_parameters.read")()
	defer inconsistentCheck(d, meta)()
	logId := getLogId(contextNil)

	var (
		instanceId = strings.Split(d.Id(), FILED_SP)[0]
		params     = d.Get("params").([]interface{})
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	dbParametersResponse, err := service.DescribeDcdbDbParametersById(ctx, instanceId)
	if err != nil {
		return err
	}
	dbParameters := dbParametersResponse.Response
	if dbParameters == nil {
		d.SetId("")
		return fmt.Errorf("resource `dbParameters` %s does not exist", d.Id())
	}

	if dbParameters.InstanceId != nil {
		_ = d.Set("instance_id", dbParameters.InstanceId)
	}

	var paramsList []interface{}
	for _, p := range params {
		paramMap := p.(map[string]interface{})
		paramName := paramMap["param"].(string)
		for _, param := range dbParameters.Params {
			if param.Param != nil && *param.Param == paramName {
				paramsMap := map[string]interface{}{}
				paramsMap["param"] = param.Param
				if param.Value != nil {
					paramsMap["value"] = param.Value
				}
				paramsList = append(paramsList, paramsMap)
				break
			}
		}
	}
	_ = d.Set("params", paramsList)

	return nil
}

func resourceTencentCloudDcdbDbParametersUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_db_parameters.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		instanceId string
		params     = d.Get("params").([]interface{})
		request    = dcdb.NewModifyDBParametersRequest()
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
		instanceId = v.(string)
	}

	if d.HasChange("params") {
		for _, p := range params {
			dBParamValue := dcdb.DBParamValue{}
			paramMap := p.(map[string]interface{})
			param := paramMap["param"].(string)
			value := paramMap["value"].(string)
			dBParamValue.Param = helper.String(param)
			dBParamValue.Value = helper.String(value)
			request.Params = append(request.Params, &dBParamValue)
		}

		//if dMap, ok := helper.InterfacesHeadMap(d, "params"); ok {
		//	dBParamValue := dcdb.DBParamValue{}
		//	if v, ok := dMap["param"]; ok {
		//		dBParamValue.Param = helper.String(v.(string))
		//		paramName = v.(string)
		//	}
		//	if v, ok := dMap["value"]; ok {
		//		dBParamValue.Value = helper.String(v.(string))
		//	}
		//	request.Params = append(request.Params, &dBParamValue)
		//}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().ModifyDBParameters(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update dcdb dbParameters failed, reason:%+v", logId, err)
		return err
	}

	err = resource.Retry(3*readRetryTimeout, func() *resource.RetryError {
		dbParametersResponse, err := service.DescribeDcdbDbParametersById(ctx, instanceId)
		if err != nil {
			return retryError(err)
		}
		dbParameters := dbParametersResponse.Response

		if dbParameters == nil || dbParameters.Params == nil {
			return resource.NonRetryableError(fmt.Errorf("DescribeDcdbDbParametersById return result(dcdb db parameter) is nil!"))
		}

		for _, param := range dbParameters.Params {
			if param.Param == nil {
				continue
			}
			for _, p := range params {
				name := p.(map[string]interface{})["param"].(string)
				if *param.Param == name && *param.Value != *param.SetValue {
					return resource.RetryableError(fmt.Errorf("db parameter initializing, retry..."))
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(strings.Join([]string{instanceId}, FILED_SP))

	return resourceTencentCloudDcdbDbParametersRead(d, meta)
}

func resourceTencentCloudDcdbDbParametersDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_db_parameters.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
