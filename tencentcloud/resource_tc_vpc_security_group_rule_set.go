/*
Provides a resource to create security group rule. This resource is similar with tencentcloudenterprise_vpc_security_group_lite_rule, rules can be ordered and configure descriptions.

~> **NOTE:** This resource must exclusive in one security group, do not declare additional rule resources of this security group elsewhere.

~> **NOTE:** `ingress` and `egress` are ordered lists, the first rule has the highest priority. Inserting a rule in the middle makes Terraform show all subsequent rules as changed because list elements are compared by index; this is a display effect only. On apply, simple changes (append, insert, remove, in-place modify) are applied incrementally and only touch the affected rules; complex changes such as reordering combined with other edits fall back to resetting the whole rule set of the changed direction.

Example Usage

```hcl
resource "tencentcloudenterprise_vpc_security_group" "sglab_1" {
  name        = "mysg_1"
  description = "favourite sg_1"
}

resource "tencentcloudenterprise_vpc_security_group_rule_set" "sglab_1" {
  security_group_id = tencentcloudenterprise_vpc_security_group.sglab_1.id
  ingress {
    cidr_block  = "10.0.0.0/16" # Accept IP or CIDR
    protocol    = "TCP" # Default is ALL
    port        = "80" # Accept port e.g. 80 or PortRange e.g. 8080-8089
    action      = "ACCEPT"
    description = "favourite sg rule_1"
  }
  ingress {
    protocol           = "TCP"
    port               = "80"
    action             = "ACCEPT"
    source_security_id = tencentcloudenterprise_vpc_security_group.sglab_3.id
    description        = "favourite sg rule_2"
  }

  egress {
    action              = "ACCEPT"
    address_template_id = "ipm-xxxxxxxx" # Support address template (group)
    description         = "Allow address template"
  }
  egress {
    action                 = "ACCEPT"
    service_template_group = "ppmg-xxxxxxxx" # Support protocol template (group)
    description            = "Allow protocol template"
  }
  egress {
    cidr_block  = "10.0.0.0/16"
    protocol    = "TCP"
    port        = "80"
    action      = "DROP"
    description = "favourite sg egress rule"
  }
}
```

Import

Resource tencentcloudenterprise_vpc_security_group_rule_set can be imported by passing security grou id:

```
terraform import tencentcloudenterprise_vpc_security_group_rule_set.sglab_1 sg-xxxxxxxx
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_security_group_rule_set", CNDescription{
		TerraformTypeCN: "批量创建安全组规则",
		DescriptionCN:   "提供批量创建安全组规则资源，用于创建和管理安全组规则。",
		AttributesCN: map[string]string{
			"security_group_id": "安全组ID",
			"ingress":           "入站规则",
			"egress":            "出站规则",
			"version":           "安全组版本",
			"action":            "安全组的规则策略。有效值“ACCEPT”和“DROP”",
			"address_template_group": "指定地址模板的组ID，如“ipmg-xxxxxxxxx”，与“source_security_ID”和“cidr_block”冲突",
			"address_template_id": "指定地址模板ID，如“ipm-xxxxxxxx”，与“source_security_ID”和“cidr_block”冲突",
			"cidr_block":        "IP地址网络或CIDR段。注意：“cidr_block”、“ipv6_cidr_block”，“source_security_id”和“address_template_*”是互斥的，不能同时设置",
			"description":       "安全组规则的描述",
			"ipv6_cidr_block":   "IPV6地址网络或CIDR段，与“source_security_id”和“address_template_*”冲突",
			"port":              "端口的范围。可用值可以是一个、多个或一个段。例如，“80”、“80、90”和“80-90”。默认为所有端口，与`service_template_*`冲突",
			"protocol":          "IP协议类型。有效值“TCP”、“UDP”和“ICMP”。默认为所有类型的协议，与`service_template_*`冲突",
			"service_template_group": "指定协议模板ID的组ID，如“ppmg-xxxxxxxxx”，与“cidr_block”和“port”冲突",
			"service_template_id": "指定协议模板ID，如“ppm-xxxxxxxx”，与“cidr_block”和“port”冲突",
			"source_security_id": "嵌套安全组的ID，与“cidr_block”和“address_template_*”冲突",
		},
	})
}

func resourceTencentCloudSecurityGroupRuleSet() *schema.Resource {
	ruleElem := map[string]*schema.Schema{
		"action": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validateAllowedStringValueIgnoreCase([]string{"ACCEPT", "DROP"}),
			Description:  "Rule policy of security group. Valid values: `ACCEPT` and `DROP`.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the security group rule.",
		},
		"cidr_block": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An IP address network or CIDR segment. NOTE: `cidr_block`, `ipv6_cidr_block`, `source_security_id` and `address_template_*` are exclusive and cannot be set in the same time.",
		},
		"ipv6_cidr_block": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An IPV6 address network or CIDR segment, and conflict with `source_security_id` and `address_template_*`.",
		},
		"source_security_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ID of the nested security group, and conflicts with `cidr_block` and `address_template_*`.",
		},
		"address_template_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify Address template ID like `ipm-xxxxxxxx`, conflict with `source_security_id` and `cidr_block`.",
		},
		"address_template_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify Group ID of Address template like `ipmg-xxxxxxxx`, conflict with `source_security_id` and `cidr_block`.",
		},
		"service_template_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify Protocol template ID like `ppm-xxxxxxxx`, conflict with `cidr_block` and `port`.",
		},
		"service_template_group": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Specify Group ID of Protocol template ID like `ppmg-xxxxxxxx`, conflict with `cidr_block` and `port`.",
		},
		"protocol": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressSecurityGroupRuleDefaultValue,
			Description:      "Type of IP protocol. Valid values: `TCP`, `UDP` and `ICMP`. Default to all types protocol, and conflicts with `service_template_*`.",
		},
		"port": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: suppressSecurityGroupRuleDefaultValue,
			Description:      "Range of the port. The available value can be one, multiple or one segment. E.g. `80`, `80,90` and `80-90`. Default to all ports, and conflicts with `service_template_*`.",
		},
	}
	return &schema.Resource{
		Description: "Provides a resource to create security group rule. This resource is similar with tencentcloudenterprise_vpc_security_group_lite_rule, rules can be ordered and configure descriptions.",
		Create:      resourceTencentCloudSecurityGroupRuleSetCreate,
		Read:        resourceTencentCloudSecurityGroupRuleSetRead,
		Update:      resourceTencentCloudSecurityGroupRuleSetUpdate,
		Delete:      resourceTencentCloudSecurityGroupRuleSetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the security group to be queried.",
			},
			"ingress": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of ingress rule. NOTE: this block is ordered, the first rule has the highest priority.",
				Elem: &schema.Resource{
					Schema: ruleElem,
				},
			},
			"egress": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of egress rule. NOTE: this block is ordered, the first rule has the highest priority.",
				Elem: &schema.Resource{
					Schema: ruleElem,
				},
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Security policies version, auto increment for every update.",
			},
		},
	}
}

func resourceTencentCloudSecurityGroupRuleSetCreate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_security_group_rule_set.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: m.(*TencentCloudClient).apiV3Conn}

	var err error
	id := d.Get("security_group_id").(string)
	request := vpc.NewModifySecurityGroupPoliciesRequest()
	request.SecurityGroupId = helper.String(id)
	request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{}

	if v, ok := d.GetOk("ingress"); ok {
		rules := v.([]interface{})
		request.SecurityGroupPolicySet.Ingress, err = unmarshalSecurityPolicy(rules)
		if err != nil {
			return err
		}
	}
	if v, ok := d.GetOk("egress"); ok {
		rules := v.([]interface{})
		request.SecurityGroupPolicySet.Egress, err = unmarshalSecurityPolicy(rules)
		if err != nil {
			return err
		}
	}

	err = service.ModifySecurityGroupPolicies(ctx, request)
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceTencentCloudSecurityGroupRuleSetRead(d, m)
}

func resourceTencentCloudSecurityGroupRuleSetRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_security_group_rule_set.read")()
	defer inconsistentCheck(d, m)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: m.(*TencentCloudClient).apiV3Conn}

	securityGroupId := d.Id()
	request := vpc.NewDescribeSecurityGroupPoliciesRequest()
	request.SecurityGroupId = &securityGroupId

	result, err := service.DescribeSecurityGroupPolicies(ctx, securityGroupId)
	if err != nil {
		return err
	}

	_ = d.Set("security_group_id", securityGroupId)
	d.SetId(securityGroupId)
	_ = d.Set("version", result.Version)
	if len(result.Ingress) > 0 {
		_ = d.Set("ingress", marshalSecurityPolicy(result.Ingress))
	}
	if len(result.Egress) > 0 {
		_ = d.Set("egress", marshalSecurityPolicy(result.Egress))
	}
	return nil
}

func resourceTencentCloudSecurityGroupRuleSetUpdate(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("tencentcloudenterprise_vpc_security_group_rule_set.update")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := m.(*TencentCloudClient).apiV3Conn
	service := VpcService{client}
	sgId := d.Id()

	needFullReplace := make(map[string]bool)
	incrementalDone := false

	for _, direction := range []string{"ingress", "egress"} {
		if !d.HasChange(direction) {
			continue
		}
		oldRaw, newRaw := d.GetChange(direction)
		oldList := oldRaw.([]interface{})
		newList := newRaw.([]interface{})
		edit := diffSecurityGroupRuleList(oldList, newList)
		if edit.typ == securityGroupRuleEditNone {
			continue
		}
		if edit.typ == securityGroupRuleEditFull {
			needFullReplace[direction] = true
			continue
		}
		// Incremental operations are index based: verify the remote list still
		// matches the last known state before touching single rules, otherwise
		// fall back to a full replace which also heals out-of-band drift.
		remote, err := service.DescribeSecurityGroupPolicies(ctx, sgId)
		if err != nil {
			return err
		}
		if !securityGroupRuleListMatchesRemote(remote, direction, oldList) {
			log.Printf("[WARN]%s security group %s %s rules drifted from state, falling back to full replace", logId, sgId, direction)
			needFullReplace[direction] = true
			continue
		}
		if err := applySecurityGroupRuleEdit(ctx, &service, sgId, direction, edit); err != nil {
			log.Printf("[WARN]%s incremental update of security group %s %s failed: %v, falling back to full replace", logId, sgId, direction, err)
			needFullReplace[direction] = true
			continue
		}
		incrementalDone = true
	}

	if len(needFullReplace) > 0 {
		if err := modifySecurityGroupPoliciesFullReplace(ctx, &service, d, needFullReplace, incrementalDone); err != nil {
			return err
		}
	}

	return resourceTencentCloudSecurityGroupRuleSetRead(d, m)
}

// modifySecurityGroupPoliciesFullReplace resets the given directions of the
// security group rule set via ModifySecurityGroupPolicies, which removes all
// existing rules of the given directions and recreates them from the
// configuration. When incrementalOpsDone is true the remote version is
// refreshed first, because preceding incremental calls may have bumped it.
func modifySecurityGroupPoliciesFullReplace(ctx context.Context, service *VpcService, d *schema.ResourceData, directions map[string]bool, incrementalOpsDone bool) error {
	logId := getLogId(ctx)
	sgId := d.Id()

	version := d.Get("version").(string)
	if incrementalOpsDone {
		remote, err := service.DescribeSecurityGroupPolicies(ctx, sgId)
		if err != nil {
			return err
		}
		if remote.Version != nil {
			version = *remote.Version
		}
	}
	ver, vErr := strconv.ParseInt(version, 10, 64)

	request := vpc.NewModifySecurityGroupPoliciesRequest()
	request.SecurityGroupId = helper.String(sgId)
	request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{}
	request.SortPolicys = helper.Bool(true)
	if vErr == nil {
		nextVer := fmt.Sprintf("%d", ver+1)
		request.SecurityGroupPolicySet.Version = &nextVer
	}

	var err error
	if directions["ingress"] {
		request.SecurityGroupPolicySet.Ingress, err = unmarshalSecurityPolicy(d.Get("ingress").([]interface{}))
		if err != nil {
			return err
		}
	}
	if directions["egress"] {
		request.SecurityGroupPolicySet.Egress, err = unmarshalSecurityPolicy(d.Get("egress").([]interface{}))
		if err != nil {
			return err
		}
	}
	if request.SecurityGroupPolicySet.Ingress == nil && request.SecurityGroupPolicySet.Egress == nil {
		log.Printf("[WARN]%s security group %s full replace skipped, no direction to update", logId, sgId)
		return nil
	}
	return service.ModifySecurityGroupPolicies(ctx, request)
}

type securityGroupRuleEditType int

const (
	securityGroupRuleEditNone securityGroupRuleEditType = iota
	securityGroupRuleEditAppend
	securityGroupRuleEditInsert
	securityGroupRuleEditRemove
	securityGroupRuleEditReplace
	securityGroupRuleEditFull
)

// securityGroupRuleEdit describes a minimal incremental change of one rule
// direction (ingress or egress) of a security group.
type securityGroupRuleEdit struct {
	typ     securityGroupRuleEditType
	index   int64         // insert: position the block is inserted at
	indices []int         // remove/replace: indices in the old list
	rules   []interface{} // append/insert: rules to create
	newList []interface{} // replace: the full new list, index aligned with the old list
}

// diffSecurityGroupRuleList computes a minimal edit between the old (state)
// and new (config) rule lists. It returns securityGroupRuleEditFull when the
// change cannot be expressed as a simple append, insert, remove or in-place
// replace, in which case the caller falls back to a full replace.
func diffSecurityGroupRuleList(oldList, newList []interface{}) *securityGroupRuleEdit {
	ruleAt := func(list []interface{}, i int) string {
		m, _ := list[i].(map[string]interface{})
		return canonicalSecurityGroupRuleFromMap(m)
	}

	switch {
	case len(newList) == len(oldList):
		var changed []int
		for i := range oldList {
			if ruleAt(oldList, i) != ruleAt(newList, i) {
				changed = append(changed, i)
			}
		}
		if len(changed) == 0 {
			return &securityGroupRuleEdit{typ: securityGroupRuleEditNone}
		}
		return &securityGroupRuleEdit{typ: securityGroupRuleEditReplace, indices: changed, newList: newList}

	case len(newList) > len(oldList):
		prefix := 0
		for prefix < len(oldList) && ruleAt(oldList, prefix) == ruleAt(newList, prefix) {
			prefix++
		}
		suffix := 0
		for suffix < len(oldList)-prefix && ruleAt(oldList, len(oldList)-1-suffix) == ruleAt(newList, len(newList)-1-suffix) {
			suffix++
		}
		if prefix+suffix == len(oldList) {
			block := newList[prefix : len(newList)-suffix]
			if prefix == len(oldList) {
				return &securityGroupRuleEdit{typ: securityGroupRuleEditAppend, rules: block}
			}
			return &securityGroupRuleEdit{typ: securityGroupRuleEditInsert, index: int64(prefix), rules: block}
		}
		return &securityGroupRuleEdit{typ: securityGroupRuleEditFull}

	default: // len(newList) < len(oldList)
		var removed []int
		j := 0
		for i := range oldList {
			if j < len(newList) && ruleAt(oldList, i) == ruleAt(newList, j) {
				j++
			} else {
				removed = append(removed, i)
			}
		}
		if j == len(newList) {
			return &securityGroupRuleEdit{typ: securityGroupRuleEditRemove, indices: removed}
		}
		return &securityGroupRuleEdit{typ: securityGroupRuleEditFull}
	}
}

// applySecurityGroupRuleEdit executes one incremental edit on the given
// direction of the security group. Every API call is atomic; on failure the
// caller falls back to a full replace which converges from any intermediate
// state.
func applySecurityGroupRuleEdit(ctx context.Context, service *VpcService, sgId, direction string, edit *securityGroupRuleEdit) error {
	switch edit.typ {
	case securityGroupRuleEditAppend, securityGroupRuleEditInsert:
		policies, err := unmarshalSecurityPolicy(edit.rules)
		if err != nil {
			return err
		}
		if edit.typ == securityGroupRuleEditInsert {
			// The API requires all policies of one request to carry the same index.
			for _, policy := range policies {
				policy.PolicyIndex = helper.Int64(edit.index)
			}
		}
		request := vpc.NewCreateSecurityGroupPoliciesRequest()
		request.SecurityGroupId = helper.String(sgId)
		request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{}
		setSecurityGroupPolicySetDirection(request.SecurityGroupPolicySet, direction, policies)
		return service.CreateSecurityGroupPolicies(ctx, request)

	case securityGroupRuleEditRemove:
		policies := make([]*vpc.SecurityGroupPolicy, 0, len(edit.indices))
		// Delete from the highest index down so lower indices stay valid even
		// if the backend applies the policies sequentially.
		for i := len(edit.indices) - 1; i >= 0; i-- {
			policies = append(policies, &vpc.SecurityGroupPolicy{PolicyIndex: helper.Int64(int64(edit.indices[i]))})
		}
		request := vpc.NewDeleteSecurityGroupPoliciesRequest()
		request.SecurityGroupId = helper.String(sgId)
		request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{}
		setSecurityGroupPolicySetDirection(request.SecurityGroupPolicySet, direction, policies)
		return service.DeleteSecurityGroupPolicies(ctx, request)

	case securityGroupRuleEditReplace:
		// Replace does not shift indices, ascending order is fine. The API
		// replaces exactly one rule per request.
		for _, idx := range edit.indices {
			policies, err := unmarshalSecurityPolicy([]interface{}{edit.newList[idx]})
			if err != nil {
				return err
			}
			policies[0].PolicyIndex = helper.Int64(int64(idx))
			request := vpc.NewReplaceSecurityGroupPolicyRequest()
			request.SecurityGroupId = helper.String(sgId)
			request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{}
			setSecurityGroupPolicySetDirection(request.SecurityGroupPolicySet, direction, policies)
			if err := service.ReplaceSecurityGroupPolicy(ctx, request); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func setSecurityGroupPolicySetDirection(set *vpc.SecurityGroupPolicySet, direction string, policies []*vpc.SecurityGroupPolicy) {
	if direction == "ingress" {
		set.Ingress = policies
	} else {
		set.Egress = policies
	}
}

// securityGroupRuleListMatchesRemote checks whether the remote rule list of
// the given direction is identical to the last known state, so that index
// based incremental operations target the intended rules.
func securityGroupRuleListMatchesRemote(remote *vpc.SecurityGroupPolicySet, direction string, oldList []interface{}) bool {
	var sdkRules []*vpc.SecurityGroupPolicy
	if direction == "ingress" {
		sdkRules = remote.Ingress
	} else {
		sdkRules = remote.Egress
	}
	if len(sdkRules) != len(oldList) {
		return false
	}
	for i := range sdkRules {
		m, ok := oldList[i].(map[string]interface{})
		if !ok {
			return false
		}
		if canonicalSecurityGroupRuleFromSdk(sdkRules[i]) != canonicalSecurityGroupRuleFromMap(m) {
			return false
		}
	}
	return true
}

func securityGroupRuleMapField(rule map[string]interface{}, key string) string {
	if v, ok := rule[key]; ok && v != nil {
		switch t := v.(type) {
		case string:
			return t
		case *string:
			if t != nil {
				return *t
			}
		}
	}
	return ""
}

func canonicalSecurityGroupRuleFromMap(rule map[string]interface{}) string {
	return canonicalSecurityGroupRuleString(
		securityGroupRuleMapField(rule, "action"),
		securityGroupRuleMapField(rule, "cidr_block"),
		securityGroupRuleMapField(rule, "ipv6_cidr_block"),
		securityGroupRuleMapField(rule, "source_security_id"),
		securityGroupRuleMapField(rule, "address_template_id"),
		securityGroupRuleMapField(rule, "address_template_group"),
		securityGroupRuleMapField(rule, "service_template_id"),
		securityGroupRuleMapField(rule, "service_template_group"),
		securityGroupRuleMapField(rule, "protocol"),
		securityGroupRuleMapField(rule, "port"),
		securityGroupRuleMapField(rule, "description"),
	)
}

func canonicalSecurityGroupRuleFromSdk(policy *vpc.SecurityGroupPolicy) string {
	str := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}
	var addressId, addressGroupId, serviceId, serviceGroupId string
	if policy.AddressTemplate != nil {
		addressId = str(policy.AddressTemplate.AddressId)
		addressGroupId = str(policy.AddressTemplate.AddressGroupId)
	}
	if policy.ServiceTemplate != nil {
		serviceId = str(policy.ServiceTemplate.ServiceId)
		serviceGroupId = str(policy.ServiceTemplate.ServiceGroupId)
	}
	return canonicalSecurityGroupRuleString(
		str(policy.Action), str(policy.CidrBlock), str(policy.Ipv6CidrBlock), str(policy.SecurityGroupId),
		addressId, addressGroupId, serviceId, serviceGroupId,
		str(policy.Protocol), str(policy.Port), str(policy.PolicyDescription),
	)
}

func canonicalSecurityGroupRuleString(action, cidrBlock, ipv6CidrBlock, securityGroupId, addressTemplateId, addressTemplateGroup, serviceTemplateId, serviceTemplateGroup, protocol, port, description string) string {
	// The backend treats an omitted protocol or port as "ALL" but may return
	// the field empty or as the literal "ALL"; both forms are equivalent.
	withDefault := func(s string) string {
		if s == "" {
			return "ALL"
		}
		return strings.ToUpper(s)
	}
	return strings.Join([]string{
		strings.ToUpper(action), cidrBlock, ipv6CidrBlock, securityGroupId,
		addressTemplateId, addressTemplateGroup, serviceTemplateId, serviceTemplateGroup,
		withDefault(protocol), withDefault(port), description,
	}, "\x00")
}

// suppressSecurityGroupRuleDefaultValue suppresses diffs between an unset
// optional value and the backend default "ALL", so that omitting protocol or
// port does not produce a perpetual diff after the backend fills in the
// default.
func suppressSecurityGroupRuleDefaultValue(k, old, new string, d *schema.ResourceData) bool {
	if strings.EqualFold(old, new) {
		return true
	}
	isAll := func(s string) bool { return strings.EqualFold(s, "ALL") }
	return (old == "" && isAll(new)) || (new == "" && isAll(old))
}

func resourceTencentCloudSecurityGroupRuleSetDelete(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_security_group_rule_set.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: m.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	request := vpc.NewModifySecurityGroupPoliciesRequest()
	request.SecurityGroupId = &id
	request.SecurityGroupPolicySet = &vpc.SecurityGroupPolicySet{
		Version: helper.String("0"),
		Ingress: []*vpc.SecurityGroupPolicy{},
		Egress:  []*vpc.SecurityGroupPolicy{},
	}
	//if v, ok := d.GetOk("ingress"); ok {
	//	rules := v.([]interface{})
	//	request.SecurityGroupPolicySet.Ingress, _ = unmarshalSecurityPolicy(rules)
	//}
	//if v, ok := d.GetOk("egress"); ok {
	//	rules := v.([]interface{})
	//	request.SecurityGroupPolicySet.Egress, _ = unmarshalSecurityPolicy(rules)
	//}
	err := service.ModifySecurityGroupPolicies(ctx, request)
	if err != nil {
		log.Printf("[CRITAL]%s security group rule delete failed: %s\n ", logId, err.Error())
		return err
	}

	return nil
}

func unmarshalSecurityPolicy(policies []interface{}) (output []*vpc.SecurityGroupPolicy, err error) {
	for i := range policies {
		policy := policies[i].(map[string]interface{})
		result := &vpc.SecurityGroupPolicy{
			Action: helper.String(policy["action"].(string)),
		}
		// CidrBlock, Ipv6CidrBlock, SecurityGroupId, AddressTemplate are exclusive, and Protocol + Port, ServiceTemplate are also exclusive
		var (
			cidrBlock            = policy["cidr_block"].(string)
			ipv6CidrBlock        = policy["ipv6_cidr_block"].(string)
			sgId                 = policy["source_security_id"].(string)
			addressTemplateId    = policy["address_template_id"].(string)
			addressTemplateGroup = policy["address_template_group"].(string)
			protocol             = policy["protocol"].(string)
			port                 = policy["port"].(string)
			serviceTemplate      = policy["service_template_id"].(string)
			serviceTemplateGroup = policy["service_template_group"].(string)
			desc                 = policy["description"].(string)
		)

		// check if exclusive arguments both set
		checkExcludeValues := func(item map[string]string) (result []string) {
			for k, v := range item {
				if v != "" {
					result = append(result, k)
				}
			}
			return result
		}

		if excludes := checkExcludeValues(map[string]string{
			"cidr_block":             cidrBlock,
			"ipv6_cidr_block":        ipv6CidrBlock,
			"source_security_id":     sgId,
			"address_template_id":    addressTemplateId,
			"address_template_group": addressTemplateGroup,
		}); len(excludes) > 1 {
			err = fmt.Errorf("conflict at rule.%d, cannot set %s in time", i, strings.Join(excludes, ","))
			return
		}

		if excludes := checkExcludeValues(map[string]string{
			"protocol + port":        protocol + port,
			"service_template_id":    serviceTemplate,
			"service_template_group": serviceTemplateGroup,
		}); len(excludes) > 1 {
			err = fmt.Errorf("conflict at rule.%d, cannot set %s in time", i, strings.Join(excludes, ","))
			return
		}

		if cidrBlock != "" {
			result.CidrBlock = &cidrBlock
		}
		if ipv6CidrBlock != "" {
			result.Ipv6CidrBlock = &ipv6CidrBlock
		}
		if sgId != "" {
			result.SecurityGroupId = &sgId
		}
		if addressTemplateId != "" || addressTemplateGroup != "" {
			result.AddressTemplate = &vpc.AddressTemplateSpecification{}
		}
		if addressTemplateId != "" {
			result.AddressTemplate.AddressId = &addressTemplateId
		}
		if addressTemplateGroup != "" {
			result.AddressTemplate.AddressGroupId = &addressTemplateGroup
		}
		if protocol != "" {
			result.Protocol = &protocol
		}
		if port != "" {
			result.Port = &port
		}
		if serviceTemplate != "" || serviceTemplateGroup != "" {
			result.ServiceTemplate = &vpc.ServiceTemplateSpecification{}
		}
		if serviceTemplate != "" {
			result.ServiceTemplate.ServiceId = &serviceTemplate
		}
		if serviceTemplateGroup != "" {
			result.ServiceTemplate.ServiceGroupId = &serviceTemplateGroup
		}
		if desc != "" {
			result.PolicyDescription = &desc
		}
		//result.PolicyIndex = helper.IntInt64(i)

		output = append(output, result)
	}
	return
}

func marshalSecurityPolicy(policies []*vpc.SecurityGroupPolicy) []interface{} {
	result := make([]interface{}, 0, len(policies))
	for i := range policies {
		policy := policies[i]
		dMap := map[string]interface{}{
			"action": policy.Action,
		}
		if policy.CidrBlock != nil {
			dMap["cidr_block"] = policy.CidrBlock
		}
		if policy.Ipv6CidrBlock != nil {
			dMap["ipv6_cidr_block"] = policy.Ipv6CidrBlock
		}
		if policy.SecurityGroupId != nil {
			dMap["source_security_id"] = policy.SecurityGroupId
		}
		if policy.AddressTemplate != nil && policy.AddressTemplate.AddressId != nil {
			dMap["address_template_id"] = policy.AddressTemplate.AddressId
		}
		if policy.AddressTemplate != nil && policy.AddressTemplate.AddressGroupId != nil {
			dMap["address_template_group"] = policy.AddressTemplate.AddressGroupId
		}
		if policy.Protocol != nil /*!checkPolicyPortIgnore(policy.Protocol, policy)*/ {
			dMap["protocol"] = strings.ToUpper(*policy.Protocol)
		}
		if policy.Port != nil /*!checkPolicyPortIgnore(policy.Port, policy)*/ {
			dMap["port"] = policy.Port
		}
		if policy.ServiceTemplate != nil && policy.ServiceTemplate.ServiceId != nil {
			dMap["service_template_id"] = policy.ServiceTemplate.ServiceId
		}
		if policy.ServiceTemplate != nil && policy.ServiceTemplate.ServiceGroupId != nil {
			dMap["service_template_group"] = policy.ServiceTemplate.ServiceGroupId
		}
		if policy.PolicyDescription != nil {
			dMap["description"] = policy.PolicyDescription
		}
		result = append(result, dMap)
	}
	return result
}
