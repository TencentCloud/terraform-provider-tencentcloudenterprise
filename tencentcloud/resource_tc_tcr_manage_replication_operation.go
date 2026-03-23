/*
Provides a resource to create a tcr manage_replication_operation

# Example Usage

```hcl

	resource "tencentcloudenterprise_tcr_instance" "mytcr_dest" {
		name        = "tf-test-tcr-%s"
		instance_type = "premium"
		delete_bucket = true
	  }

	resource "tencentcloudenterprise_tcr_namespace" "myns_dest" {
		instance_id 	 = tencentcloudenterprise_tcr_instance.mytcr_dest.id
		name			 = "tf_test_ns_dest"
		is_public		 = true
		is_auto_scan	 = true
		is_prevent_vul   = true
		severity		 = "medium"
		cve_whitelist_items	{
			cve_id = "cve-xxxxx"
		}
	}

	resource "tencentcloudenterprise_tcr_manage_replication_operation" "my_replica" {
	  source_registry_id = local.tcr_id
	  destination_registry_id = tencentcloudenterprise_tcr_instance.mytcr_dest.id
	  rule {
			name = "test_sync_%d"
			dest_namespace = tencentcloudenterprise_tcr_namespace.myns_dest.name
			override = true
			filters {
				type = "name"
				value = join("/", [var.tcr_namespace, "**"])
			}
			filters {
				type = "tag"
				value = ""
			}
			filters {
				type = "resource"
				value = ""
			}
	  }
	  description = "this is the tcr sync operation"
	  destination_region_id = 1 // "ap-region"
	  peer_replication_option {
			peer_registry_uin = ""
			peer_registry_token = ""
			enable_peer_replication = false
	  }
	}

```
*/
package tencentcloud

import (
	"log"
	"strings"

	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tcr_manage_replication_operation", CNDescription{
		TerraformTypeCN: "TCR管理同步复制操作",
		DescriptionCN:   "提供TCR管理同步复制操作资源，用于管理容器镜像仓库的同步复制",
		AttributesCN: map[string]string{
			"source_registry_id":      "源实例ID",
			"destination_registry_id": "目标实例ID",
			"rule":                    "同步规则",
			"description":             "规则描述",
			"dest_namespace":          "目标命名空间",
			"override":                "是否覆盖",
			"name":                    "规则名称",
			"dest_region_id":          "目标地域ID",
			"filters":                 "过滤器",
		},
	})
}

func resourceTencentCloudTcrManageReplicationOperation() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudTcrManageReplicationOperationCreate,
		Read:        resourceTencentCloudTcrManageReplicationOperationRead,
		Delete:      resourceTencentCloudTcrManageReplicationOperationDelete,
		Description: "Provides a resource to create TCR manage replication operation",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"source_registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Copy source instance Id.",
			},

			"destination_registry_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Copy destination instance Id.",
			},

			"rule": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Synchronization rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Synchronization rule names.",
						},
						"dest_namespace": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Target namespace.",
						},
						"override": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to cover.",
						},
						"filters": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Sync filters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Type (name, tag, and resource).",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Empty by default.",
									},
								},
							},
						},
					},
				},
			},

			"description": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Rule description.",
			},

			"destination_region_id": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "The region ID of the target instance.",
			},

			"peer_replication_option": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Enable synchronization of configuration items across master account instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"peer_registry_uin": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Uin of the instance to be synchronized.",
						},
						"peer_registry_token": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Access permanent token of the instance to be synchronized.",
						},
						"enable_peer_replication": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether to enable cross-master account instance synchronization.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudTcrManageReplicationOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_manage_replication_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request               = tcr.NewManageReplicationRequest()
		sourceRegistryId      string
		destinationRegistryId string
		ruleName              string
	)
	if v, ok := d.GetOk("source_registry_id"); ok {
		sourceRegistryId = v.(string)
		request.SourceRegistryId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("destination_registry_id"); ok {
		destinationRegistryId = v.(string)
		request.DestinationRegistryId = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "rule"); ok {
		replicationRule := tcr.ReplicationRule{}
		if v, ok := dMap["name"]; ok {
			ruleName = v.(string)
			replicationRule.Name = helper.String(v.(string))
		}
		if v, ok := dMap["dest_namespace"]; ok {
			replicationRule.DestNamespace = helper.String(v.(string))
		}
		if v, ok := dMap["override"]; ok {
			replicationRule.Override = helper.Bool(v.(bool))
		}
		if v, ok := dMap["filters"]; ok {
			for _, item := range v.([]interface{}) {
				filtersMap := item.(map[string]interface{})
				replicationFilter := tcr.ReplicationFilter{}
				if v, ok := filtersMap["type"]; ok {
					replicationFilter.Type = helper.String(v.(string))
				}
				if v, ok := filtersMap["value"]; ok {
					replicationFilter.Value = helper.String(v.(string))
				}
				replicationRule.Filters = append(replicationRule.Filters, &replicationFilter)
			}
		}
		request.Rule = &replicationRule
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, _ := d.GetOk("destination_region_id"); v != nil {
		request.DestinationRegionId = helper.IntUint64(v.(int))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "peer_replication_option"); ok {
		peerReplicationOption := tcr.PeerReplicationOption{}
		if v, ok := dMap["peer_registry_uin"]; ok {
			peerReplicationOption.PeerRegistryUin = helper.String(v.(string))
		}
		if v, ok := dMap["peer_registry_token"]; ok {
			peerReplicationOption.PeerRegistryToken = helper.String(v.(string))
		}
		if v, ok := dMap["enable_peer_replication"]; ok {
			peerReplicationOption.EnablePeerReplication = helper.Bool(v.(bool))
		}
		request.PeerReplicationOption = &peerReplicationOption
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTCRClient().ManageReplication(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s operate tcr ManageReplicationOperation failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(strings.Join([]string{sourceRegistryId, destinationRegistryId, ruleName}, FILED_SP))

	return resourceTencentCloudTcrManageReplicationOperationRead(d, meta)
}

func resourceTencentCloudTcrManageReplicationOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_manage_replication_operation.read")()
	defer inconsistentCheck(d, meta)()

	return nil
}

func resourceTencentCloudTcrManageReplicationOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tcr_manage_replication_operation.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}
