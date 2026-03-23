/*
Provides a resource to create a CCN route table.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ccn_route_table" "example" {
	  ccn_id      = "ccn-cwm743gl"
	  name        = "test-ccn-route-table"
	  description = "test ccn route table description"
	}

```

# Import

CCN route table can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table.example ccnrtb-gbaaugtl
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ccn_route_table", CNDescription{
		TerraformTypeCN: "云联网路由表",
		DescriptionCN:   "用于创建和管理云联网路由表。",
		AttributesCN: map[string]string{
			"ccn_id":           "云联网实例 ID",
			"name":             "路由表名称",
			"description":      "路由表描述",
			"is_default_table": "是否为默认路由表",
			"create_time":      "创建时间",
		},
	})
}

func resourceTencentCloudCcnRouteTable() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a CCN route table.",
		Create:      resourceTencentCloudCcnRouteTableCreate,
		Read:        resourceTencentCloudCcnRouteTableRead,
		Update:      resourceTencentCloudCcnRouteTableUpdate,
		Delete:      resourceTencentCloudCcnRouteTableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "CCN instance ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CCN route table name.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the CCN route table.",
			},
			"is_default_table": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this route table is the default route table.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the route table.",
			},
		},
	}
}

func resourceTencentCloudCcnRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := ccn.NewCreateCcnRouteTablesRequest()
	request.RouteTable = []*ccn.CcnBatchRouteTable{
		{
			CcnId:       helper.String(d.Get("ccn_id").(string)),
			Name:        helper.String(d.Get("name").(string)),
			Description: helper.String(d.Get("description").(string)),
		},
	}

	var response *ccn.CreateCcnRouteTablesResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().CreateCcnRouteTables(request)
		if e != nil {
			return retryError(e)
		}
		if result == nil || result.Response == nil || len(result.Response.CcnRouteTableSet) != 1 {
			return resource.NonRetryableError(fmt.Errorf("create ccn route table failed"))
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create ccn route table failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(*response.Response.CcnRouteTableSet[0].CcnRouteTableId)
	return resourceTencentCloudCcnRouteTableRead(d, meta)
}

func resourceTencentCloudCcnRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	ccnRouteTable, err := service.DescribeVpcCcnRouteTablesById(ctx, d.Id())
	if err != nil {
		return err
	}
	if ccnRouteTable == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_ccn_route_table` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if ccnRouteTable.CcnId != nil {
		_ = d.Set("ccn_id", *ccnRouteTable.CcnId)
	}
	if ccnRouteTable.RouteTableName != nil {
		_ = d.Set("name", *ccnRouteTable.RouteTableName)
	}
	if ccnRouteTable.RouteTableDescription != nil {
		_ = d.Set("description", *ccnRouteTable.RouteTableDescription)
	}
	if ccnRouteTable.IsDefaultTable != nil {
		_ = d.Set("is_default_table", *ccnRouteTable.IsDefaultTable)
	}
	if ccnRouteTable.CreateTime != nil {
		_ = d.Set("create_time", *ccnRouteTable.CreateTime)
	}

	return nil
}

func resourceTencentCloudCcnRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table.update")()
	defer inconsistentCheck(d, meta)()

	if !d.HasChange("name") && !d.HasChange("description") {
		return resourceTencentCloudCcnRouteTableRead(d, meta)
	}

	logId := getLogId(contextNil)
	request := ccn.NewModifyCcnRouteTablesRequest()
	request.RouteTableInfo = []*ccn.ModifyRouteTableInfo{
		{
			RouteTableId: helper.String(d.Id()),
			Name:         helper.String(d.Get("name").(string)),
			Description:  helper.String(d.Get("description").(string)),
		},
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().ModifyCcnRouteTables(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s modify ccn route table failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCcnRouteTableRead(d, meta)
}

func resourceTencentCloudCcnRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ccn_route_table.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := ccn.NewDeleteCcnRouteTablesRequest()
	request.RouteTableId = helper.Strings([]string{d.Id()})

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCcnClient().DeleteCcnRouteTables(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s], request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete ccn route table failed, reason:%+v", logId, err)
		return err
	}

	return nil
}
