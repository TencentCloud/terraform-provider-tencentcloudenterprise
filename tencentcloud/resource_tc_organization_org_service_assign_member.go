/*
Provides a resource to manage an organization trusted-service delegated administrator (assign member).

Each resource instance represents a single delegated administrator binding for one
trusted service (`service_id`) on one member account (`member_uin`).

Example Usage

```hcl
resource "tencentcloudenterprise_organization_org_service_assign_member" "demo" {
  service_id        = 24
  member_uin        = 110000003055
  management_scope  = 1
}
```

Import

Organization service assign member can be imported using `service_id#member_uin`, e.g.

```
$ terraform import tencentcloudenterprise_organization_org_service_assign_member.demo 24#110000003055
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudOrganizationOrgServiceAssignMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudOrganizationOrgServiceAssignMemberCreate,
		Read:   resourceTencentCloudOrganizationOrgServiceAssignMemberRead,
		Delete: resourceTencentCloudOrganizationOrgServiceAssignMemberDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"service_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Trusted service ID. e.g. 24 stands for cloudaudit.",
			},
			"member_uin": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The delegated administrator member UIN.",
			},
			"management_scope": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Default:     1,
				Description: "Management scope. `1`: all members; `2`: partial members (must specify `management_scope_uins` and/or `management_scope_node_ids`).",
			},
			"management_scope_uins": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Member UIN list. Effective only when `management_scope` is `2`.",
			},
			"management_scope_node_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Node ID list. Effective only when `management_scope` is `2`.",
			},
			"member_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Member account name.",
			},
			"product_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Product name of the trusted service.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the assignment is created.",
			},
		},
	}
}

func resourceTencentCloudOrganizationOrgServiceAssignMemberCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_service_assign_member.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		serviceId uint64
		memberUin int64
	)

	request := organization.NewCreateOrgServiceAssignRequest()
	if v, ok := d.GetOkExists("service_id"); ok {
		serviceId = uint64(v.(int))
		request.ServiceId = helper.Uint64(serviceId)
	}
	if v, ok := d.GetOkExists("member_uin"); ok {
		memberUin = int64(v.(int))
		request.MemberUins = []*int64{helper.Int64(memberUin)}
	}
	if v, ok := d.GetOkExists("management_scope"); ok {
		request.ManagementScope = helper.Uint64(uint64(v.(int)))
	}
	if v, ok := d.GetOk("management_scope_uins"); ok {
		set := v.(*schema.Set).List()
		for _, u := range set {
			request.ManagementScopeUins = append(request.ManagementScopeUins, helper.Int64(int64(u.(int))))
		}
	}
	if v, ok := d.GetOk("management_scope_node_ids"); ok {
		set := v.(*schema.Set).List()
		for _, n := range set {
			request.ManagementScopeNodeIds = append(request.ManagementScopeNodeIds, helper.Int64(int64(n.(int))))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().CreateOrgServiceAssign(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create organization org_service_assign_member failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(helper.UInt64ToStr(serviceId) + FILED_SP + helper.Int64ToStr(memberUin))
	return resourceTencentCloudOrganizationOrgServiceAssignMemberRead(d, meta)
}

func resourceTencentCloudOrganizationOrgServiceAssignMemberRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_service_assign_member.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	serviceId := helper.StrToUInt64(idSplit[0])
	memberUin := helper.StrToInt64(idSplit[1])

	service := OrganizationService{client: meta.(*TencentCloudClient).apiV3Conn}
	member, err := service.DescribeOrgServiceAssignMember(ctx, serviceId, memberUin)
	if err != nil {
		return err
	}
	if member == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_organization_org_service_assign_member` [%s] not found, removing from state.", logId, d.Id())
		return nil
	}

	_ = d.Set("service_id", serviceId)
	_ = d.Set("member_uin", memberUin)
	if member.ManagementScope != nil {
		_ = d.Set("management_scope", member.ManagementScope)
	}
	if member.ManagementScopeMembers != nil {
		uins := make([]int, 0, len(member.ManagementScopeMembers))
		for _, m := range member.ManagementScopeMembers {
			if m != nil && m.MemberUin != nil {
				uins = append(uins, int(*m.MemberUin))
			}
		}
		_ = d.Set("management_scope_uins", uins)
	}
	if member.ManagementScopeNodes != nil {
		nodes := make([]int, 0, len(member.ManagementScopeNodes))
		for _, n := range member.ManagementScopeNodes {
			if n != nil && n.NodeId != nil {
				nodes = append(nodes, int(*n.NodeId))
			}
		}
		_ = d.Set("management_scope_node_ids", nodes)
	}
	if member.MemberName != nil {
		_ = d.Set("member_name", member.MemberName)
	}
	if member.ProductName != nil {
		_ = d.Set("product_name", member.ProductName)
	}
	if member.CreateTime != nil {
		_ = d.Set("create_time", member.CreateTime)
	}

	return nil
}

func resourceTencentCloudOrganizationOrgServiceAssignMemberDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_organization_org_service_assign_member.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	serviceId := helper.StrToUInt64(idSplit[0])
	memberUin := helper.StrToInt64(idSplit[1])

	request := organization.NewDeleteOrgServiceAssignRequest()
	request.ServiceId = helper.Uint64(serviceId)
	request.MemberUin = helper.Int64(memberUin)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOrganizationClient().DeleteOrgServiceAssign(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete organization org_service_assign_member failed, reason:%+v", logId, err)
		return err
	}
	return nil
}
