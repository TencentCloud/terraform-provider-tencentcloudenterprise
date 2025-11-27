/*
Provides a resource to create a CLB attachment.

Example Usage
```hcl

	resource "tencentcloudenterprise_clb_attachment" "foo" {
	  clb_id      = "lb-k2zjp9lv"
	  listener_id = "lbl-hh141sn9"
	  rule_id     = "loc-4xxr2cy7"

	  targets {
	    instance_id = "ins-1flbqyp8"
	    port        = 80
	    weight      = 10
	  }
	}

```

Import
CLB attachment can be imported using the id, e.g.

```
$ terraform import cloud_clb_attachment.foo loc-4xxr2cy7#lbl-hh141sn9#lb-7a0t6zqb
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_attachment", CNDescription{
		TerraformTypeCN: "CLB绑定",
		DescriptionCN:   "提供CLB绑定资源，用于将后端服务器绑定到CLB监听器。",
		AttributesCN: map[string]string{
			"clb_id":          "CLB实例ID",
			"listener_id":     "监听器ID",
			"rule_id":         "监听器规则ID",
			"protocol_type":   "协议类型",
			"targets":         "绑定目标",
			"instance_id":     "后端服务器的CVM实例Id，与'eni_ip', 'instance_bms_id'冲突，但必须指定其中之一",
			"instance_bms_id": "后端服务器的BMS实例Id，与'eni_ip', 'instance_id'冲突，但必须指定其中之一",
			"eni_ip":          "后端服务器的Eni IP地址，与'instance_id', 'instance_bms_id'冲突，但必须指定其中之一",
			"port":            "后端服务器的端口。有效值范围:(0~65535)",
			"weight":          "后端服务的转发权重。有效值范围:(0~100)。默认为“10”",
		},
	})
}
func resourceTencentCloudClbServerAttachment() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a CLB attachment.",
		Create:      resourceTencentCloudClbServerAttachmentCreate,
		Read:        resourceTencentCloudClbServerAttachmentRead,
		Update:      resourceTencentCloudClbServerAttachmentUpdate,
		Delete:      resourceTencentCloudClbServerAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the CLB.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the CLB listener.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "ID of the CLB listener rule. Only supports listeners of `HTTPS` and `HTTP` protocol.",
			},
			"protocol_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of protocol within the listener.",
			},
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				MaxItems:    100,
				Description: "Information of the backends to be attached.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:     schema.TypeString,
							Optional: true,
							//ConflictsWith: []string{"instance_bms_id", "eni_ip"},
							Description: "CVM Instance Id of the backend server, conflict with `eni_ip`  `instance_bms_id` but must specify one of them.",
						},
						"eni_ip": {
							Type:     schema.TypeString,
							Optional: true,
							//ConflictsWith: []string{"instance_bms_id", "instance_id"},
							Description: "Eni IP address of the backend server, conflict with `instance_id` `instance_bms_id` but must specify one of them.",
						},
						"instance_bms_id": {
							Type:     schema.TypeString,
							Optional: true,
							//ConflictsWith: []string{"instance_id", "eni_ip"},
							Description: "Instance_bms_id address of the backend server, conflict with `instance_id`  `eni_ip` but must specify one of them.",
						},
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validateIntegerInRange(0, 65535),
							Description:  "Port of the backend server. Valid value ranges: (0~65535).",
						},
						"weight": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      10,
							ValidateFunc: validateIntegerInRange(0, 100),
							Description:  "Forwarding weight of the backend service. Valid value ranges: (0~100). defaults to `10`.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudClbServerAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.create")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	logId := getLogId(contextNil)
	listenerId := d.Get("listener_id").(string)
	checkErr := ListenerIdCheck(listenerId)
	if checkErr != nil {
		return checkErr
	}
	clbId := d.Get("clb_id").(string)
	locationId := ""
	request := clb.NewRegisterTargetsRequest()
	request.LoadBalancerId = helper.String(clbId)
	request.ListenerId = helper.String(listenerId)
	if v, ok := d.GetOk("rule_id"); ok {
		locationId = v.(string)
		checkErr := RuleIdCheck(locationId)
		if checkErr != nil {
			return checkErr
		}
		if locationId != "" {
			request.LocationId = helper.String(locationId)
		}
	}

	insList := d.Get("targets").(*schema.Set).List()
	// log.Printf("[WARN] inputTarget %#v", insList)
	insLen := len(insList)
	for count := 0; count < insLen; count += 20 {
		//this request only support 20 targets at most once time
		request.Targets = make([]*clb.Target, 0, 20)
		for i := 0; i < 20; i++ {
			index := count + i
			if index >= insLen {
				break
			}
			inst := insList[index].(map[string]interface{})
			request.Targets = append(request.Targets, clbNewTarget(inst["instance_id"], inst["eni_ip"], inst["port"], inst["weight"], inst["instance_bms_id"]))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			requestId := ""
			ratelimit.Check(request.GetAction())
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().RegisterTargets(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				requestId = *result.Response.RequestId
				retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
				if retryErr != nil {
					return resource.NonRetryableError(errors.WithStack(retryErr))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s create CLB attachment failed, reason:%+v", logId, err)
			return err
		}
	}
	id := fmt.Sprintf("%s#%v#%v", locationId, d.Get("listener_id"), d.Get("clb_id"))
	d.SetId(id)

	return resourceTencentCloudClbServerAttachmentRead(d, meta)
}

func resourceTencentCloudClbServerAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.delete")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	clbService := ClbService{client: meta.(*TencentCloudClient).apiV3Conn}

	attachmentId := d.Id()
	items := strings.Split(attachmentId, "#")
	if len(items) < 3 {
		return fmt.Errorf("[CHECK][CLB attachment][Delete] check: id %s of resource.cloud_clb_attachment is not match loc-xxx#lbl-xxx#lb-xxx", attachmentId)
	}

	locationId := items[0]
	listenerId := items[1]
	clbId := items[2]

	request := clb.NewDeregisterTargetsRequest()
	request.ListenerId = &listenerId
	request.LoadBalancerId = helper.String(clbId)
	if locationId != "" {
		request.LocationId = helper.String(locationId)
	}

	//insList := d.Get("targets").(*schema.Set).List()

	// filter target group which cvm not existed
	inputTarget := d.Get("targets").(*schema.Set).List()
	// log.Printf("[WARN] inputTarget %#v", inputTarget)
	insList := getRemoveCandidates(ctx, clbService, clbId, listenerId, locationId, inputTarget)
	insLen := len(insList)
	log.Printf("tencentcloudenterprise_clb_attachment delete insLen:" + strconv.Itoa(insLen))
	for count := 0; count < insLen; count += 20 {
		//this request only support 20 targets at most once time
		request.Targets = make([]*clb.Target, 0, 20)
		for i := 0; i < 20; i++ {
			index := count + i
			if index >= insLen {
				break
			}
			inst := insList[index].(map[string]interface{})
			request.Targets = append(request.Targets, clbNewTarget(inst["instance_id"], inst["eni_ip"], inst["port"], inst["weight"], inst["instance_bms_id"]))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			requestId := ""
			ratelimit.Check(request.GetAction())
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().DeregisterTargets(request)
			if e != nil {
				ee, ok := e.(*sdkErrors.CloudSDKError)
				if ok && ee.GetCode() == "InvalidParameter" {
					return nil
				}
				return retryError(e)

			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				requestId = *result.Response.RequestId
				retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
				if retryErr != nil {
					return resource.NonRetryableError(errors.WithStack(retryErr))
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s create CLB attachment failed, reason:%+v", logId, err)
			return err
		}
	}

	return nil
}

func resourceTencentCloudClbServerAttachmentRemove(d *schema.ResourceData, meta interface{}, remove []interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.remove")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	attachmentId := d.Id()
	items := strings.Split(attachmentId, "#")
	if len(items) < 3 {
		return fmt.Errorf("[CHECK][CLB attachment][Remove] check: id %s of resource.cloud_clb_attachment is not match loc-xxx#lbl-xxx#lb-xxx", attachmentId)
	}
	locationId := items[0]
	listenerId := items[1]
	clbId := items[2]

	request := clb.NewDeregisterTargetsRequest()
	request.ListenerId = helper.String(listenerId)
	request.LoadBalancerId = helper.String(clbId)
	if locationId != "" {
		request.LocationId = helper.String(locationId)
	}

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		removeCandidates := getRemoveCandidates(ctx, clbService, clbId, listenerId, locationId, remove)
		if len(removeCandidates) == 0 {
			return nil
		}
		e := clbService.DeleteAttachmentById(ctx, clbId, listenerId, locationId, removeCandidates)
		if e != nil {
			return retryError(e)
		}

		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s reason[%+v]", logId, err)
		return err
	}

	return nil
}

func resourceTencentCloudClbServerAttachmentAdd(d *schema.ResourceData, meta interface{}, add []interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.add")()
	logId := getLogId(contextNil)

	listenerId := d.Get("listener_id").(string)
	clbId := d.Get("clb_id").(string)
	locationId := ""
	request := clb.NewRegisterTargetsRequest()
	request.LoadBalancerId = helper.String(clbId)
	request.ListenerId = helper.String(listenerId)

	if v, ok := d.GetOk("rule_id"); ok {
		locationId = v.(string)
		if locationId != "" {
			request.LocationId = helper.String(locationId)
		}
	}

	for _, v := range add {
		inst := v.(map[string]interface{})
		request.Targets = append(request.Targets, clbNewTarget(inst["instance_id"], inst["eni_ip"], inst["port"], inst["weight"], inst["instance_bms_id"]))
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		requestId := ""
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseClbClient().RegisterTargets(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			requestId = *response.Response.RequestId
			retryErr := waitForTaskFinish(requestId, meta.(*TencentCloudClient).apiV3Conn.UseClbClient())
			if retryErr != nil {
				return resource.NonRetryableError(errors.WithStack(retryErr))
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s add CLB attachment failed, reason:%+v", logId, err)
		return err
	}
	return nil
}

func resourceTencentCloudClbServerAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.update")()

	clbActionMu.Lock()
	defer clbActionMu.Unlock()

	if d.HasChange("targets") {
		o, n := d.GetChange("targets")
		os := o.(*schema.Set)
		ns := n.(*schema.Set)
		add := ns.Difference(os).List()
		// log.Printf("[WARN] add %#v", add)
		remove := os.Difference(ns).List()
		// log.Printf("[WARN] remove %#v", remove)
		addLen := len(add)
		removeLen := len(remove)
		if removeLen > 0 {
			for count := 0; count < removeLen; count += 20 {
				removeList := make([]interface{}, 0, 20)
				for i := 0; i < 20; i++ {
					index := count + i
					if index >= removeLen {
						break
					}
					removeList = append(removeList, remove[index])
				}
				// log.Printf("[WARN] removeList %#v", removeList)
				err := resourceTencentCloudClbServerAttachmentRemove(d, meta, removeList)
				if err != nil {
					return err
				}
			}
		}
		if addLen > 0 {
			for count := 0; count < addLen; count += 20 {
				addList := make([]interface{}, 0, 20)
				for i := 0; i < 20; i++ {
					index := count + i
					if index >= addLen {
						break
					}
					addList = append(addList, add[index])
				}
				err := resourceTencentCloudClbServerAttachmentAdd(d, meta, addList)
				if err != nil {
					return err
				}
			}
		}
		return resourceTencentCloudClbServerAttachmentRead(d, meta)
	}

	return nil
}

func resourceTencentCloudClbServerAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_clb_attachment.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	items := strings.Split(d.Id(), "#")
	locationId := items[0]
	listenerId := items[1]
	clbId := items[2]

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var instance *clb.ListenerBackend
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := clbService.DescribeAttachmentByPara(ctx, clbId, listenerId, locationId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB attachment failed, reason:%+v", logId, err)
		return err
	}
	//see if read empty

	if instance == nil || (len(instance.Targets) == 0 && locationId == "") || (len(instance.Rules) == 0 && locationId != "") {
		d.SetId("")
		return nil
	}

	_ = d.Set("clb_id", clbId)
	_ = d.Set("listener_id", listenerId)
	_ = d.Set("protocol_type", instance.Protocol)

	var onlineTargets []*clb.Backend
	if *instance.Protocol == CLB_LISTENER_PROTOCOL_HTTP || *instance.Protocol == CLB_LISTENER_PROTOCOL_HTTPS {
		_ = d.Set("rule_id", locationId)
		if len(instance.Rules) > 0 {
			for _, loc := range instance.Rules {
				if locationId == "" || locationId == *loc.LocationId {
					onlineTargets = loc.Targets
				}
			}
		}
		// TCP / UDP / TCP_SSL
	} else if instance.Targets != nil {
		onlineTargets = instance.Targets
	}

	//this may cause problems when there are members in two dimensions array
	//need to read state of the tfstate file to clear the relationships
	//in this situation, import action is not supported
	// TL,DR: just update partial targets which this resource declared.
	stateTargets := d.Get("targets").(*schema.Set)
	log.Printf("[WARN] stateTargets %#v", stateTargets)
	if stateTargets.Len() != 0 {
		//the old state exist
		//create a new attachment with state
		exactTargets := make([]interface{}, 0)
		log.Printf("[WARN] onlineTargets %#v", onlineTargets)
		for i := range onlineTargets {
			v := onlineTargets[i]
			// log.Printf("[WARN] v %#v, v.Type:%#v, v.InstanceId:%#v, v.PrivateIpAddresses:%#v",
			// v, *v.Type, v.InstanceId, v.PrivateIpAddresses)
			if *v.Type == "CVM" && v.InstanceId != nil {
				target := map[string]interface{}{
					"weight":      int(*v.Weight),
					"port":        int(*v.Port),
					"instance_id": *v.InstanceId,
				}
				// log.Printf("[WARN] type=cvm target %#v", target)
				// para in main.tf with no instance_id will not be in stateTargets
				// main.tf para stateTargets:
				// targets {
				// 	port = 10002
				// 	weight = 0
				// 	eni_ip = "192.168.100.6"
				// }
				// DescribeTargets returns onlineTargets[i]:
				//   {
				//     "EniId": "eni-rvugv4iv",
				//     "InstanceId": "ins-2rtl1aoe",
				//     "InstanceName": "QTA_NET勿动_CVM_内网_1.1.2",
				//     "Port": 10002,
				//     "PrivateIpAddresses": [
				//       "192.168.100.6"
				//     ],
				//     "PublicIpAddresses": null,
				//     "RegisteredTime": "2025-05-29 01:29:22",
				//     "Type": "CVM",
				//     "Weight": 0
				//   }
				if stateTargets.Contains(target) {
					// log.Printf("[WARN] type=cvm stateTargets.Contains(target) use instance_id")
					exactTargets = append(exactTargets, map[string]interface{}{
						"weight":      int(*v.Weight),
						"port":        int(*v.Port),
						"instance_id": *v.InstanceId,
					})
				} else {
					// log.Printf("[WARN] type=cvm stateTargets.Contains(target) use eni_ip, %#v", v.PrivateIpAddresses)
					if v.PrivateIpAddresses != nil && len(v.PrivateIpAddresses) > 0 && *v.PrivateIpAddresses[0] != "" {
						target = map[string]interface{}{
							"weight": int(*v.Weight),
							"port":   int(*v.Port),
							"eni_ip": *v.PrivateIpAddresses[0],
						}
						// log.Printf("[WARN] type=cvm stateTargets.Contains(target) use eni_ip, %#v", target)
						if stateTargets.Contains(target) {
							// log.Printf("[WARN] type=cvm stateTargets.Contains(target) use eni_ip")
							exactTargets = append(exactTargets, map[string]interface{}{
								"weight":      int(*v.Weight),
								"port":        int(*v.Port),
								"instance_id": *v.InstanceId,
							})
						}
					}
				}
			} else if len(v.PrivateIpAddresses) > 0 && *v.PrivateIpAddresses[0] != "" {
				target := map[string]interface{}{
					"weight": int(*v.Weight),
					"port":   int(*v.Port),
					"eni_ip": *v.PrivateIpAddresses[0],
				}
				// log.Printf("[WARN] type=%#v target %#v", *v.Type, target)
				if stateTargets.Contains(target) {
					// log.Printf("[WARN] type=%#v stateTargets.Contains(target)", *v.Type)
					exactTargets = append(exactTargets, map[string]interface{}{
						"weight": int(*v.Weight),
						"port":   int(*v.Port),
						"eni_ip": *v.PrivateIpAddresses[0],
					})
				}

				for _, target := range stateTargets.List() {
					if m, ok := target.(map[string]interface{}); ok {
						// log.Printf("[WARN] m %#v", m)
						// log.Printf("[WARN] *v.Port %#v", *v.Port)
						if m["instance_bms_id"] != "" {
							// log.Printf("[WARN] m %#v", m)
							// log.Printf("[WARN] *v.Port %#v", *v.Port)
							if port, ok := m["port"].(int); ok && port == int(*v.Port) {
								// log.Printf("[WARN] m %#v", m)
								// log.Printf("[WARN] *v.Port %#v", *v.Port)
								exactTargets = append(exactTargets, map[string]interface{}{
									"weight":          int(*v.Weight),
									"port":            int(*v.Port),
									"instance_bms_id": m["instance_bms_id"],
								})
								log.Printf("[WARN] exactTargets %#v", exactTargets)
							}
						}
					}
				}
			}
		}
		_ = d.Set("targets", exactTargets)
		log.Printf("[WARN] exactTargets %#v", exactTargets)
	} else {
		_ = d.Set("targets", flattenBackendList(onlineTargets))
	}

	return nil
}

// Destroy CVM instance will dispatch async task to deregister target group from cloudApi backend. Duplicate deregister target groups here will cause Error response.
// If remove diffs created, check existing cvm instance first, filter target groups which already deregister
func getRemoveCandidates(ctx context.Context, clbService ClbService, clbId string, listenerId string, locationId string, remove []interface{}) []interface{} {
	removeCandidates := make([]interface{}, 0)
	existAttachments, err := clbService.DescribeAttachmentByPara(ctx, clbId, listenerId, locationId)
	if err != nil {
		return removeCandidates
	}
	// DescribeTargets  IPv4内网CLB实例4层监听器 返回的Targets非空
	existTargetGroups := existAttachments.Targets
	log.Printf("[WARN] existTargetGroups %v", existTargetGroups)
	// DescribeTargets  IPv4内网CLB实例7层监听器 返回的Rules为空
	if existTargetGroups == nil || len(existTargetGroups) == 0 {
		rules := existAttachments.Rules
		for _, rule := range rules {
			if rule.Targets != nil && len(rule.Targets) > 0 {
				existTargetGroups = append(existTargetGroups, rule.Targets...)
			}
		}
	}
	log.Printf("[WARN] existTargetGroups %v", existTargetGroups)

	for _, item := range remove {
		target := item.(map[string]interface{})
		log.Printf("[WARN] target %#v", target)
		if targetGroupContainsInstance(existTargetGroups, target["instance_id"]) ||
			targetGroupContainsEni(existTargetGroups, target["eni_ip"]) ||
			targetGroupContainsBmsInstance(existTargetGroups, target["instance_bms_id"]) {
			removeCandidates = append(removeCandidates, target)
		}
	}
	return removeCandidates
}

func targetGroupContainsInstance(targets []*clb.Backend, instanceId interface{}) (contains bool) {
	contains = false
	id, ok := instanceId.(string)
	if !ok || id == "" {
		return false
	}
	for _, target := range targets {
		if target.InstanceId == nil {
			continue
		}
		if id == *target.InstanceId {
			log.Printf("[WARN] Instance %s applied.", id)
			return true
		}
	}
	log.Printf("[WARN] Instance %s not exist, skip deregister.", id)

	return
}

func targetGroupContainsBmsInstance(targets []*clb.Backend, instanceId interface{}) (contains bool) {
	contains = false
	id, ok := instanceId.(string)
	if !ok || id == "" {
		return false
	}
	for _, target := range targets {
		if target.InstanceId == nil {
			continue
		}
		if id == *target.InstanceId {
			log.Printf("[WARN] Instance %s applied.", id)
			return true
		}
	}
	log.Printf("[WARN] Instance %s not exist, skip deregister.", id)

	return
}

func targetGroupContainsEni(targets []*clb.Backend, eniIp interface{}) (contains bool) {
	contains = false
	ip, ok := eniIp.(string)
	if !ok || ip == "" {
		return
	}
	for _, target := range targets {
		if len(target.PrivateIpAddresses) == 0 ||
			(len(target.PrivateIpAddresses) > 0 && target.PrivateIpAddresses[0] == nil) {
			continue
		}
		if ip == *target.PrivateIpAddresses[0] {
			log.Printf("[WARN] IP %s applied.", ip)
			return true
		}
	}
	log.Printf("[WARN] IP %s not exist, skip deregister.", ip)

	return
}
