/*
Provide a resource to create a TKE native node pool.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_native_node_pool" "example" {
	  cluster_id = "cls-xxx"
	  name       = "native-pool"
	  type       = "Native"

	  native {
	    subnet_ids           = ["subnet-xxx"]
	    instance_charge_type = "POSTPAID_BY_HOUR"
	    system_disk {
	      disk_type = "CLOUD_SSD"
	      disk_size = 50
	    }
	    instance_types     = ["S3.MEDIUM2"]
	    security_group_ids = ["sg-xxx"]
	    scaling {
	      min_replicas = 1
	      max_replicas = 10
	    }
	  }
	}

```
*/
package tencentcloud

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"strings"
//
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//
//	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
//	tke2 "terraform-provider-tencentcloudenterprise/sdk/tke/v20220501"
//	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
//)
//
//func resourceTencentCloudKubernetesNativeNodePool() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceTencentCloudKubernetesNativeNodePoolCreate,
//		Read:   resourceTencentCloudKubernetesNativeNodePoolRead,
//		Update: resourceTencentCloudKubernetesNativeNodePoolUpdate,
//		Delete: resourceTencentCloudKubernetesNativeNodePoolDelete,
//		Importer: &schema.ResourceImporter{
//			State: schema.ImportStatePassthrough,
//		},
//		Schema: map[string]*schema.Schema{
//			"cluster_id": {
//				Type:        schema.TypeString,
//				Required:    true,
//				ForceNew:    true,
//				Description: "ID of the cluster.",
//			},
//			"name": {
//				Type:        schema.TypeString,
//				Required:    true,
//				Description: "Node pool name.",
//			},
//			"type": {
//				Type:        schema.TypeString,
//				Required:    true,
//				ForceNew:    true,
//				Description: "Node pool type. Optional value is `Native`.",
//			},
//			"labels": {
//				Type:        schema.TypeSet,
//				Optional:    true,
//				Description: "Node Labels.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"name": {
//							Type:        schema.TypeString,
//							Required:    true,
//							Description: "Label name.",
//						},
//						"value": {
//							Type:        schema.TypeString,
//							Required:    true,
//							Description: "Label value.",
//						},
//					},
//				},
//			},
//			"taints": {
//				Type:        schema.TypeList,
//				Optional:    true,
//				Description: "Node taint.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"key": {
//							Type:        schema.TypeString,
//							Optional:    true,
//							Description: "Key of the taint.",
//						},
//						"value": {
//							Type:        schema.TypeString,
//							Optional:    true,
//							Description: "Value of the taint.",
//						},
//						"effect": {
//							Type:        schema.TypeString,
//							Optional:    true,
//							Description: "Effect of the taint.",
//						},
//					},
//				},
//			},
//			"deletion_protection": {
//				Type:        schema.TypeBool,
//				Optional:    true,
//				Computed:    true,
//				Description: "Whether to enable deletion protection.",
//			},
//			"unschedulable": {
//				Type:        schema.TypeBool,
//				Optional:    true,
//				Computed:    true,
//				Description: "Whether the node is not schedulable by default.",
//			},
//			"native": {
//				Type:        schema.TypeList,
//				Required:    true,
//				MaxItems:    1,
//				Description: "Native node pool creation parameters.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"scaling": {
//							Type:        schema.TypeList,
//							Optional:    true,
//							Computed:    true,
//							MaxItems:    1,
//							Description: "Node pool scaling configuration.",
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									"min_replicas": {
//										Type:        schema.TypeInt,
//										Optional:    true,
//										Computed:    true,
//										Description: "Minimum number of replicas in node pool.",
//									},
//									"max_replicas": {
//										Type:        schema.TypeInt,
//										Optional:    true,
//										Computed:    true,
//										Description: "Maximum number of replicas in node pool.",
//									},
//									"create_policy": {
//										Type:        schema.TypeString,
//										Optional:    true,
//										Computed:    true,
//										Description: "Node pool expansion strategy.",
//									},
//								},
//							},
//						},
//						"subnet_ids": {
//							Type:        schema.TypeList,
//							Required:    true,
//							Description: "Subnet list.",
//							Elem:        &schema.Schema{Type: schema.TypeString},
//						},
//						"instance_charge_type": {
//							Type:        schema.TypeString,
//							Required:    true,
//							ForceNew:    true,
//							Description: "Node billing type. `PREPAID` or `POSTPAID_BY_HOUR`.",
//						},
//						"system_disk": {
//							Type:        schema.TypeList,
//							Required:    true,
//							ForceNew:    true,
//							MaxItems:    1,
//							Description: "System disk configuration.",
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									"disk_type": {
//										Type:        schema.TypeString,
//										Required:    true,
//										ForceNew:    true,
//										Description: "Cloud disk type.",
//									},
//									"disk_size": {
//										Type:        schema.TypeInt,
//										Required:    true,
//										ForceNew:    true,
//										Description: "Cloud disk size (G).",
//									},
//								},
//							},
//						},
//						"instance_types": {
//							Type:        schema.TypeList,
//							Required:    true,
//							ForceNew:    true,
//							Description: "Model list.",
//							Elem:        &schema.Schema{Type: schema.TypeString},
//						},
//						"security_group_ids": {
//							Type:        schema.TypeList,
//							Required:    true,
//							Description: "Security group list.",
//							Elem:        &schema.Schema{Type: schema.TypeString},
//						},
//						"auto_repair": {
//							Type:        schema.TypeBool,
//							Optional:    true,
//							Description: "Whether to enable self-healing ability.",
//						},
//						"enable_autoscaling": {
//							Type:        schema.TypeBool,
//							Optional:    true,
//							Description: "Whether to enable elastic scaling.",
//						},
//						"replicas": {
//							Type:        schema.TypeInt,
//							Optional:    true,
//							Computed:    true,
//							Description: "Desired number of nodes.",
//						},
//						"data_disks": {
//							Type:        schema.TypeList,
//							Optional:    true,
//							Description: "Data disk list.",
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									"disk_type": {
//										Type:        schema.TypeString,
//										Required:    true,
//										Description: "Cloud disk type.",
//									},
//									"disk_size": {
//										Type:        schema.TypeInt,
//										Required:    true,
//										Description: "Cloud disk size (G).",
//									},
//									"auto_format_and_mount": {
//										Type:        schema.TypeBool,
//										Optional:    true,
//										Description: "Whether to automatically format and mount.",
//									},
//									"mount_target": {
//										Type:        schema.TypeString,
//										Optional:    true,
//										Description: "Mount directory.",
//									},
//									"file_system": {
//										Type:        schema.TypeString,
//										Optional:    true,
//										Description: "File system (ext3/ext4/xfs).",
//									},
//								},
//							},
//						},
//						"key_ids": {
//							Type:        schema.TypeList,
//							Optional:    true,
//							Description: "SSH public key id array.",
//							Elem:        &schema.Schema{Type: schema.TypeString},
//						},
//					},
//				},
//			},
//			"annotations": {
//				Type:        schema.TypeSet,
//				Optional:    true,
//				Computed:    true,
//				Description: "Node Annotation List.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"name": {
//							Type:        schema.TypeString,
//							Required:    true,
//							Description: "Annotation name.",
//						},
//						"value": {
//							Type:        schema.TypeString,
//							Required:    true,
//							Description: "Annotation value.",
//						},
//					},
//				},
//			},
//			// Computed fields
//			"life_state": {
//				Type:        schema.TypeString,
//				Computed:    true,
//				Description: "Node pool status.",
//			},
//			"created_at": {
//				Type:        schema.TypeString,
//				Computed:    true,
//				Description: "Creation time.",
//			},
//		},
//	}
//}
//
//func resourceTencentCloudKubernetesNativeNodePoolCreate(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_native_node_pool.create")()
//
//	logId := getLogId(contextNil)
//	_ = context.WithValue(context.TODO(), logIdKey, logId)
//
//	var (
//		clusterId  string
//		nodePoolId string
//	)
//
//	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()
//	request := tke2.NewCreateNodePoolRequest()
//
//	if v, ok := d.GetOk("cluster_id"); ok {
//		clusterId = v.(string)
//		request.ClusterId = helper.String(clusterId)
//	}
//
//	if v, ok := d.GetOk("name"); ok {
//		request.Name = helper.String(v.(string))
//	}
//
//	if v, ok := d.GetOk("type"); ok {
//		request.Type = helper.String(v.(string))
//	}
//
//	// Labels
//	if v, ok := d.GetOk("labels"); ok {
//		for _, item := range v.(*schema.Set).List() {
//			labelsMap := item.(map[string]interface{})
//			label := tke2.Label{}
//			if v, ok := labelsMap["name"]; ok {
//				label.Name = helper.String(v.(string))
//			}
//			if v, ok := labelsMap["value"]; ok {
//				label.Value = helper.String(v.(string))
//			}
//			request.Labels = append(request.Labels, &label)
//		}
//	}
//
//	// Taints
//	if v, ok := d.GetOk("taints"); ok {
//		for _, item := range v.([]interface{}) {
//			taintsMap := item.(map[string]interface{})
//			taint := tke2.Taint{}
//			if v, ok := taintsMap["key"]; ok {
//				taint.Key = helper.String(v.(string))
//			}
//			if v, ok := taintsMap["value"]; ok {
//				taint.Value = helper.String(v.(string))
//			}
//			if v, ok := taintsMap["effect"]; ok {
//				taint.Effect = helper.String(v.(string))
//			}
//			request.Taints = append(request.Taints, &taint)
//		}
//	}
//
//	if v, ok := d.GetOkExists("deletion_protection"); ok {
//		request.DeletionProtection = helper.Bool(v.(bool))
//	}
//
//	if v, ok := d.GetOkExists("unschedulable"); ok {
//		request.Unschedulable = helper.Bool(v.(bool))
//	}
//
//	// Native configuration
//	if nativeList, ok := d.GetOk("native"); ok {
//		nativeListArr := nativeList.([]interface{})
//		if len(nativeListArr) > 0 {
//			nativeMap := nativeListArr[0].(map[string]interface{})
//			createNativeNodePoolParam := tke2.CreateNativeNodePoolParam{}
//
//			// Scaling
//			if scalingList, ok := nativeMap["scaling"]; ok {
//				scalingListArr := scalingList.([]interface{})
//				if len(scalingListArr) > 0 {
//					scalingMap := scalingListArr[0].(map[string]interface{})
//					machineSetScaling := tke2.MachineSetScaling{}
//					if v, ok := scalingMap["min_replicas"]; ok {
//						machineSetScaling.MinReplicas = helper.Int64(int64(v.(int)))
//					}
//					if v, ok := scalingMap["max_replicas"]; ok {
//						machineSetScaling.MaxReplicas = helper.Int64(int64(v.(int)))
//					}
//					if v, ok := scalingMap["create_policy"]; ok {
//						machineSetScaling.CreatePolicy = helper.String(v.(string))
//					}
//					createNativeNodePoolParam.Scaling = &machineSetScaling
//				}
//			}
//
//			// SubnetIds
//			if v, ok := nativeMap["subnet_ids"]; ok {
//				subnetIdsSet := v.([]interface{})
//				for i := range subnetIdsSet {
//					subnetIds := subnetIdsSet[i].(string)
//					createNativeNodePoolParam.SubnetIds = append(createNativeNodePoolParam.SubnetIds, helper.String(subnetIds))
//				}
//			}
//
//			// InstanceChargeType
//			if v, ok := nativeMap["instance_charge_type"]; ok {
//				createNativeNodePoolParam.InstanceChargeType = helper.String(v.(string))
//			}
//
//			// SystemDisk
//			if systemDiskList, ok := nativeMap["system_disk"]; ok {
//				systemDiskListArr := systemDiskList.([]interface{})
//				if len(systemDiskListArr) > 0 {
//					systemDiskMap := systemDiskListArr[0].(map[string]interface{})
//					disk := tke2.Disk{}
//					if v, ok := systemDiskMap["disk_type"]; ok {
//						disk.DiskType = helper.String(v.(string))
//					}
//					if v, ok := systemDiskMap["disk_size"]; ok {
//						disk.DiskSize = helper.Int64(int64(v.(int)))
//					}
//					createNativeNodePoolParam.SystemDisk = &disk
//				}
//			}
//
//			// InstanceTypes
//			if v, ok := nativeMap["instance_types"]; ok {
//				instanceTypesSet := v.([]interface{})
//				for i := range instanceTypesSet {
//					instanceTypes := instanceTypesSet[i].(string)
//					createNativeNodePoolParam.InstanceTypes = append(createNativeNodePoolParam.InstanceTypes, helper.String(instanceTypes))
//				}
//			}
//
//			// SecurityGroupIds
//			if v, ok := nativeMap["security_group_ids"]; ok {
//				securityGroupIdsSet := v.([]interface{})
//				for i := range securityGroupIdsSet {
//					securityGroupIds := securityGroupIdsSet[i].(string)
//					createNativeNodePoolParam.SecurityGroupIds = append(createNativeNodePoolParam.SecurityGroupIds, helper.String(securityGroupIds))
//				}
//			}
//
//			// AutoRepair
//			if v, ok := nativeMap["auto_repair"]; ok {
//				createNativeNodePoolParam.AutoRepair = helper.Bool(v.(bool))
//			}
//
//			// EnableAutoscaling
//			if v, ok := nativeMap["enable_autoscaling"]; ok {
//				createNativeNodePoolParam.EnableAutoscaling = helper.Bool(v.(bool))
//			}
//
//			// Replicas
//			if v, ok := nativeMap["replicas"]; ok {
//				createNativeNodePoolParam.Replicas = helper.Int64(int64(v.(int)))
//			}
//
//			// DataDisks
//			if v, ok := nativeMap["data_disks"]; ok {
//				dataDisksSet := v.([]interface{})
//				for _, item := range dataDisksSet {
//					dataDiskMap := item.(map[string]interface{})
//					dataDisk := tke2.DataDisk{}
//					if v, ok := dataDiskMap["disk_type"]; ok {
//						dataDisk.DiskType = helper.String(v.(string))
//					}
//					if v, ok := dataDiskMap["disk_size"]; ok {
//						dataDisk.DiskSize = helper.Int64(int64(v.(int)))
//					}
//					if v, ok := dataDiskMap["auto_format_and_mount"]; ok {
//						dataDisk.AutoFormatAndMount = helper.Bool(v.(bool))
//					}
//					if v, ok := dataDiskMap["mount_target"]; ok {
//						dataDisk.MountTarget = helper.String(v.(string))
//					}
//					if v, ok := dataDiskMap["file_system"]; ok {
//						dataDisk.FileSystem = helper.String(v.(string))
//					}
//					createNativeNodePoolParam.DataDisks = append(createNativeNodePoolParam.DataDisks, &dataDisk)
//				}
//			}
//
//			// KeyIds
//			if v, ok := nativeMap["key_ids"]; ok {
//				keyIdsSet := v.([]interface{})
//				for i := range keyIdsSet {
//					keyIds := keyIdsSet[i].(string)
//					createNativeNodePoolParam.KeyIds = append(createNativeNodePoolParam.KeyIds, helper.String(keyIds))
//				}
//			}
//
//			request.Native = &createNativeNodePoolParam
//		}
//	}
//
//	// Annotations
//	if v, ok := d.GetOk("annotations"); ok {
//		for _, item := range v.(*schema.Set).List() {
//			annotationMap := item.(map[string]interface{})
//			annotation := tke2.Annotation{}
//			if v, ok := annotationMap["name"]; ok {
//				annotation.Name = helper.String(v.(string))
//			}
//			if v, ok := annotationMap["value"]; ok {
//				annotation.Value = helper.String(v.(string))
//			}
//			request.Annotations = append(request.Annotations, &annotation)
//		}
//	}
//
//	log.Printf("[DEBUG]%s api[CreateNodePool] request: %s", logId, request.ToJsonString())
//
//	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
//		result, e := client.CreateNodePool(request)
//		if e != nil {
//			return retryError(e)
//		}
//		if result == nil || result.Response == nil || result.Response.NodePoolId == nil {
//			return resource.NonRetryableError(fmt.Errorf("CreateNodePool response is nil"))
//		}
//		nodePoolId = *result.Response.NodePoolId
//		return nil
//	})
//	if err != nil {
//		log.Printf("[CRITAL]%s create native node pool failed, reason:%+v", logId, err)
//		return err
//	}
//
//	d.SetId(clusterId + FILED_SP + nodePoolId)
//
//	return resourceTencentCloudKubernetesNativeNodePoolRead(d, meta)
//}
//
//func resourceTencentCloudKubernetesNativeNodePoolRead(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_native_node_pool.read")()
//
//	logId := getLogId(contextNil)
//	_ = context.WithValue(context.TODO(), logIdKey, logId)
//
//	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()
//
//	idSplit := strings.Split(d.Id(), FILED_SP)
//	if len(idSplit) != 2 {
//		return fmt.Errorf("id is broken, id is %s", d.Id())
//	}
//	clusterId := idSplit[0]
//	nodePoolId := idSplit[1]
//
//	request := tke2.NewDescribeNodePoolsRequest()
//	request.ClusterId = helper.String(clusterId)
//	request.Filters = []*tke2.Filter{
//		{
//			Name:   helper.String("NodePoolsId"),
//			Values: []*string{helper.String(nodePoolId)},
//		},
//	}
//
//	var nodePool *tke2.NodePool
//	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
//		result, e := client.DescribeNodePools(request)
//		if e != nil {
//			return retryError(e)
//		}
//		if result == nil || result.Response == nil {
//			return resource.NonRetryableError(fmt.Errorf("DescribeNodePools response is nil"))
//		}
//		if len(result.Response.NodePools) == 0 {
//			d.SetId("")
//			return nil
//		}
//		nodePool = result.Response.NodePools[0]
//		return nil
//	})
//	if err != nil {
//		log.Printf("[CRITAL]%s read native node pool failed, reason:%+v", logId, err)
//		return err
//	}
//
//	if nodePool == nil {
//		d.SetId("")
//		return nil
//	}
//
//	_ = d.Set("cluster_id", clusterId)
//	_ = d.Set("name", nodePool.Name)
//	_ = d.Set("type", nodePool.Type)
//	_ = d.Set("life_state", nodePool.LifeState)
//	_ = d.Set("created_at", nodePool.CreatedAt)
//
//	if nodePool.DeletionProtection != nil {
//		_ = d.Set("deletion_protection", nodePool.DeletionProtection)
//	}
//	if nodePool.Unschedulable != nil {
//		_ = d.Set("unschedulable", nodePool.Unschedulable)
//	}
//
//	// Labels
//	if nodePool.Labels != nil {
//		labels := make([]map[string]interface{}, 0, len(nodePool.Labels))
//		for _, label := range nodePool.Labels {
//			labels = append(labels, map[string]interface{}{
//				"name":  label.Name,
//				"value": label.Value,
//			})
//		}
//		_ = d.Set("labels", labels)
//	}
//
//	// Taints
//	if nodePool.Taints != nil {
//		taints := make([]map[string]interface{}, 0, len(nodePool.Taints))
//		for _, taint := range nodePool.Taints {
//			taints = append(taints, map[string]interface{}{
//				"key":    taint.Key,
//				"value":  taint.Value,
//				"effect": taint.Effect,
//			})
//		}
//		_ = d.Set("taints", taints)
//	}
//
//	// Native
//	if nodePool.Native != nil {
//		native := nodePool.Native
//		nativeMap := make(map[string]interface{})
//
//		if native.Scaling != nil {
//			scalingMap := map[string]interface{}{
//				"min_replicas":  native.Scaling.MinReplicas,
//				"max_replicas":  native.Scaling.MaxReplicas,
//				"create_policy": native.Scaling.CreatePolicy,
//			}
//			nativeMap["scaling"] = []interface{}{scalingMap}
//		}
//
//		if native.SubnetIds != nil {
//			subnetIds := make([]string, 0, len(native.SubnetIds))
//			for _, v := range native.SubnetIds {
//				subnetIds = append(subnetIds, *v)
//			}
//			nativeMap["subnet_ids"] = subnetIds
//		}
//
//		nativeMap["instance_charge_type"] = native.InstanceChargeType
//		nativeMap["auto_repair"] = native.AutoRepair
//		nativeMap["enable_autoscaling"] = native.EnableAutoscaling
//		nativeMap["replicas"] = native.Replicas
//
//		if native.SystemDisk != nil {
//			systemDiskMap := map[string]interface{}{
//				"disk_type": native.SystemDisk.DiskType,
//				"disk_size": native.SystemDisk.DiskSize,
//			}
//			nativeMap["system_disk"] = []interface{}{systemDiskMap}
//		}
//
//		if native.InstanceTypes != nil {
//			instanceTypes := make([]string, 0, len(native.InstanceTypes))
//			for _, v := range native.InstanceTypes {
//				instanceTypes = append(instanceTypes, *v)
//			}
//			nativeMap["instance_types"] = instanceTypes
//		}
//
//		if native.SecurityGroupIds != nil {
//			securityGroupIds := make([]string, 0, len(native.SecurityGroupIds))
//			for _, v := range native.SecurityGroupIds {
//				securityGroupIds = append(securityGroupIds, *v)
//			}
//			nativeMap["security_group_ids"] = securityGroupIds
//		}
//
//		if native.DataDisks != nil {
//			dataDisks := make([]map[string]interface{}, 0, len(native.DataDisks))
//			for _, disk := range native.DataDisks {
//				diskMap := map[string]interface{}{
//					"disk_type":             disk.DiskType,
//					"disk_size":             disk.DiskSize,
//					"auto_format_and_mount": disk.AutoFormatAndMount,
//					"mount_target":          disk.MountTarget,
//					"file_system":           disk.FileSystem,
//				}
//				dataDisks = append(dataDisks, diskMap)
//			}
//			nativeMap["data_disks"] = dataDisks
//		}
//
//		if native.KeyIds != nil {
//			keyIds := make([]string, 0, len(native.KeyIds))
//			for _, v := range native.KeyIds {
//				keyIds = append(keyIds, *v)
//			}
//			nativeMap["key_ids"] = keyIds
//		}
//
//		_ = d.Set("native", []interface{}{nativeMap})
//	}
//
//	// Annotations
//	if nodePool.Annotations != nil {
//		annotations := make([]map[string]interface{}, 0, len(nodePool.Annotations))
//		for _, annotation := range nodePool.Annotations {
//			annotations = append(annotations, map[string]interface{}{
//				"name":  annotation.Name,
//				"value": annotation.Value,
//			})
//		}
//		_ = d.Set("annotations", annotations)
//	}
//
//	return nil
//}
//
//func resourceTencentCloudKubernetesNativeNodePoolUpdate(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_native_node_pool.update")()
//
//	logId := getLogId(contextNil)
//	_ = context.WithValue(context.TODO(), logIdKey, logId)
//
//	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()
//
//	idSplit := strings.Split(d.Id(), FILED_SP)
//	if len(idSplit) != 2 {
//		return fmt.Errorf("id is broken, id is %s", d.Id())
//	}
//	clusterId := idSplit[0]
//	nodePoolId := idSplit[1]
//
//	request := tke2.NewModifyNodePoolRequest()
//	request.ClusterId = helper.String(clusterId)
//	request.NodePoolId = helper.String(nodePoolId)
//
//	needUpdate := false
//
//	if d.HasChange("name") {
//		request.Name = helper.String(d.Get("name").(string))
//		needUpdate = true
//	}
//
//	if d.HasChange("labels") {
//		if v, ok := d.GetOk("labels"); ok {
//			for _, item := range v.(*schema.Set).List() {
//				labelsMap := item.(map[string]interface{})
//				label := tke2.Label{}
//				if v, ok := labelsMap["name"]; ok {
//					label.Name = helper.String(v.(string))
//				}
//				if v, ok := labelsMap["value"]; ok {
//					label.Value = helper.String(v.(string))
//				}
//				request.Labels = append(request.Labels, &label)
//			}
//		}
//		needUpdate = true
//	}
//
//	if d.HasChange("taints") {
//		if v, ok := d.GetOk("taints"); ok {
//			for _, item := range v.([]interface{}) {
//				taintsMap := item.(map[string]interface{})
//				taint := tke2.Taint{}
//				if v, ok := taintsMap["key"]; ok {
//					taint.Key = helper.String(v.(string))
//				}
//				if v, ok := taintsMap["value"]; ok {
//					taint.Value = helper.String(v.(string))
//				}
//				if v, ok := taintsMap["effect"]; ok {
//					taint.Effect = helper.String(v.(string))
//				}
//				request.Taints = append(request.Taints, &taint)
//			}
//		}
//		needUpdate = true
//	}
//
//	if d.HasChange("deletion_protection") {
//		request.DeletionProtection = helper.Bool(d.Get("deletion_protection").(bool))
//		needUpdate = true
//	}
//
//	if d.HasChange("unschedulable") {
//		request.Unschedulable = helper.Bool(d.Get("unschedulable").(bool))
//		needUpdate = true
//	}
//
//	if d.HasChange("annotations") {
//		if v, ok := d.GetOk("annotations"); ok {
//			for _, item := range v.(*schema.Set).List() {
//				annotationMap := item.(map[string]interface{})
//				annotation := tke2.Annotation{}
//				if v, ok := annotationMap["name"]; ok {
//					annotation.Name = helper.String(v.(string))
//				}
//				if v, ok := annotationMap["value"]; ok {
//					annotation.Value = helper.String(v.(string))
//				}
//				request.Annotations = append(request.Annotations, &annotation)
//			}
//		}
//		needUpdate = true
//	}
//
//	if d.HasChange("native") {
//		updateNativeParam := tke2.UpdateNativeNodePoolParam{}
//		hasNativeChange := false
//
//		if nativeList, ok := d.GetOk("native"); ok {
//			nativeListArr := nativeList.([]interface{})
//			if len(nativeListArr) > 0 {
//				nativeMap := nativeListArr[0].(map[string]interface{})
//
//				// Scaling
//				if scalingList, ok := nativeMap["scaling"]; ok {
//					scalingListArr := scalingList.([]interface{})
//					if len(scalingListArr) > 0 {
//						scalingMap := scalingListArr[0].(map[string]interface{})
//						machineSetScaling := tke2.MachineSetScaling{}
//						if v, ok := scalingMap["min_replicas"]; ok {
//							machineSetScaling.MinReplicas = helper.Int64(int64(v.(int)))
//						}
//						if v, ok := scalingMap["max_replicas"]; ok {
//							machineSetScaling.MaxReplicas = helper.Int64(int64(v.(int)))
//						}
//						if v, ok := scalingMap["create_policy"]; ok {
//							machineSetScaling.CreatePolicy = helper.String(v.(string))
//						}
//						updateNativeParam.Scaling = &machineSetScaling
//						hasNativeChange = true
//					}
//				}
//
//				// SubnetIds
//				if v, ok := nativeMap["subnet_ids"]; ok {
//					subnetIdsSet := v.([]interface{})
//					for i := range subnetIdsSet {
//						subnetIds := subnetIdsSet[i].(string)
//						updateNativeParam.SubnetIds = append(updateNativeParam.SubnetIds, helper.String(subnetIds))
//					}
//					hasNativeChange = true
//				}
//
//				// SecurityGroupIds
//				if v, ok := nativeMap["security_group_ids"]; ok {
//					securityGroupIdsSet := v.([]interface{})
//					for i := range securityGroupIdsSet {
//						securityGroupIds := securityGroupIdsSet[i].(string)
//						updateNativeParam.SecurityGroupIds = append(updateNativeParam.SecurityGroupIds, helper.String(securityGroupIds))
//					}
//					hasNativeChange = true
//				}
//
//				// AutoRepair
//				if v, ok := nativeMap["auto_repair"]; ok {
//					updateNativeParam.AutoRepair = helper.Bool(v.(bool))
//					hasNativeChange = true
//				}
//
//				// EnableAutoscaling
//				if v, ok := nativeMap["enable_autoscaling"]; ok {
//					updateNativeParam.EnableAutoscaling = helper.Bool(v.(bool))
//					hasNativeChange = true
//				}
//
//				// Replicas
//				if v, ok := nativeMap["replicas"]; ok {
//					updateNativeParam.Replicas = helper.Int64(int64(v.(int)))
//					hasNativeChange = true
//				}
//			}
//		}
//
//		if hasNativeChange {
//			request.Native = &updateNativeParam
//			needUpdate = true
//		}
//	}
//
//	if needUpdate {
//		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
//			_, e := client.ModifyNodePool(request)
//			if e != nil {
//				return retryError(e)
//			}
//			return nil
//		})
//		if err != nil {
//			log.Printf("[CRITAL]%s update native node pool failed, reason:%+v", logId, err)
//			return err
//		}
//	}
//
//	return resourceTencentCloudKubernetesNativeNodePoolRead(d, meta)
//}
//
//func resourceTencentCloudKubernetesNativeNodePoolDelete(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_native_node_pool.delete")()
//
//	logId := getLogId(contextNil)
//	_ = context.WithValue(context.TODO(), logIdKey, logId)
//
//	client := meta.(*TencentCloudClient).apiV3Conn.UseTke2Client()
//
//	idSplit := strings.Split(d.Id(), FILED_SP)
//	if len(idSplit) != 2 {
//		return fmt.Errorf("id is broken, id is %s", d.Id())
//	}
//	clusterId := idSplit[0]
//	nodePoolId := idSplit[1]
//
//	request := tke2.NewDeleteNodePoolRequest()
//	request.ClusterId = helper.String(clusterId)
//	request.NodePoolId = helper.String(nodePoolId)
//
//	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
//		_, e := client.DeleteNodePool(request)
//		if e != nil {
//			return retryError(e)
//		}
//		return nil
//	})
//	if err != nil {
//		log.Printf("[CRITAL]%s delete native node pool failed, reason:%+v", logId, err)
//		return err
//	}
//
//	// Wait for deletion
//	err = resource.Retry(5*readRetryTimeout, func() *resource.RetryError {
//		descRequest := tke2.NewDescribeNodePoolsRequest()
//		descRequest.ClusterId = helper.String(clusterId)
//		descRequest.Filters = []*tke2.Filter{
//			{
//				Name:   helper.String("NodePoolsId"),
//				Values: []*string{helper.String(nodePoolId)},
//			},
//		}
//		result, e := client.DescribeNodePools(descRequest)
//		if e != nil {
//			if sdkErr, ok := e.(*sdkErrors.CloudSDKError); ok {
//				if strings.Contains(sdkErr.Message, "not found") || strings.Contains(sdkErr.Message, "NotFound") {
//					return nil
//				}
//			}
//			return retryError(e)
//		}
//		if result == nil || result.Response == nil || len(result.Response.NodePools) == 0 {
//			return nil
//		}
//		return resource.RetryableError(fmt.Errorf("native node pool %s still exists", nodePoolId))
//	})
//	if err != nil {
//		log.Printf("[CRITAL]%s wait for native node pool deletion failed, reason:%+v", logId, err)
//		return err
//	}
//
//	return nil
//}
