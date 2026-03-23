/*
Provides an eip resource associated with other resource like CVM, CLB, ENI, BMS and HAVIP.

~> **NOTE:** Please DO NOT define `allocate_public_ip` in `tencentcloudenterprise_cvm_instance` resource when using `tencentcloudenterprise_eip_association`.

Example Usage

Bind EIP to CVM instance

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_cvm" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "ins-xxxxxx"
}
```

Bind EIP to CLB instance

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_clb" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "lb-xxxxxx"
}
```

Bind EIP to ENI with specific private IP

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_vpc_eni" "example" {
  name      = "eni-example"
  vpc_id    = "vpc-xxxxxx"
  subnet_id = "subnet-xxxxxx"
}

resource "tencentcloudenterprise_eip_association" "eip_eni" {
  eip_id               = tencentcloudenterprise_eip.example.id
  network_interface_id = tencentcloudenterprise_vpc_eni.example.id
  private_ip           = "10.0.1.22"
}
```

Bind EIP to BMS (Blackstone Metal Server)

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_eip_association" "eip_bms" {
  eip_id      = tencentcloudenterprise_eip.example.id
  instance_id = "bms-xxxxxx"
  private_ip  = "10.24.0.6"
}
```

Bind EIP to HAVIP (High Availability Virtual IP)

```hcl
resource "tencentcloudenterprise_eip" "example" {
  name = "eip-example"
}

resource "tencentcloudenterprise_havip" "example" {
  name      = "havip-example"
  vpc_id    = "vpc-xxxxxx"
  subnet_id = "subnet-xxxxxx"
}

resource "tencentcloudenterprise_eip_association" "eip_havip" {
  eip_id   = tencentcloudenterprise_eip.example.id
  havip_id = tencentcloudenterprise_havip.example.id
}
```

Import

Eip association can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_eip_association.eip_cvm eip-41s6jwy4::ins-34jwj3
$ terraform import tencentcloudenterprise_eip_association.eip_eni eip-41s6jwy4::eni-34jwj3::10.0.1.22
$ terraform import tencentcloudenterprise_eip_association.eip_havip eip-41s6jwy4::havip-34jwj3
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_eip_association", CNDescription{
		TerraformTypeCN: "弹性公网IP绑定到实例",
		DescriptionCN:   "提供弹性公网IP绑定资源，用于将弹性公网IP绑定到CVM实例、CLB实例、ENI、BMS、HAVIP等资源。",
		AttributesCN: map[string]string{
			"eip_id":               "标识 EIP 的唯一 ID，EIP 唯一 ID 形如：`eip-11112222`",
			"instance_id":          "要绑定的实例 ID，实例 ID 形如：`ins-11112222`（CVM）、`lb-11112222`（CLB）、`bms-11112222`（BMS），与 `NetworkInterfaceId` 和 `HaVipId` 不可同时指定",
			"network_interface_id": "要绑定的弹性网卡 ID，弹性网卡 ID 形如：`eni-11112222`，与 `InstanceId` 和 `HaVipId` 不可同时指定",
			"private_ip":           "要绑定的内网 IP，当绑定到 ENI 时必须指定，当绑定到 BMS 时可指定具体的内网 IP，与 `HaVipId` 不可同时指定",
			"havip_id":             "要绑定的 HAVIP ID，HAVIP ID 形如：`havip-11112222`，与 `InstanceId` 和 `NetworkInterfaceId` 不可同时指定",
		},
	})
}

func resourceTencentCloudEipAssociation() *schema.Resource {
	return &schema.Resource{
		Description: "Provides an eip resource associated with other resource like CVM, CLB, ENI, BMS and HAVIP.",
		Create:      resourceTencentCloudEipAssociationCreate,
		Read:        resourceTencentCloudEipAssociationRead,
		Delete:      resourceTencentCloudEipAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"eip_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateStringLengthInRange(1, 25),
				Description:  "The ID of EIP.",
			},
			"instance_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
				ConflictsWith: []string{
					"network_interface_id",
					"havip_id",
				},
				ValidateFunc: validateStringLengthInRange(1, 25),
				Description:  "The CVM, CLB or BMS instance id going to bind with the EIP. This field is conflict with `network_interface_id` and `havip_id`.",
			},
			"network_interface_id": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringLengthInRange(1, 25),
				ConflictsWith: []string{
					"instance_id",
				},
				Description: "Indicates the network interface id like `eni-xxxxxx`. This field is conflict with `instance_id`.",
			},
			"private_ip": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringLengthInRange(7, 25),
				ConflictsWith: []string{
					"havip_id",
				},
				Description: "Indicates a private IP belongs to the `network_interface_id`, or a specific private IP to bind when using with BMS instance. This field is conflict with `havip_id`.",
			},
			"havip_id": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateStringLengthInRange(1, 25),
				ConflictsWith: []string{
					"instance_id",
					"network_interface_id",
				},
				Description: "The HAVip id going to bind with the EIP. This field is conflict with `instance_id` and `network_interface_id`.",
			},
		},
	}
}

func resourceTencentCloudEipAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_eip_association.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	vpcService := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	eipId := d.Get("eip_id").(string)
	var eip *vpc.Address
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		eip, errRet = vpcService.DescribeEipById(ctx, eipId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		if eip == nil {
			return resource.NonRetryableError(fmt.Errorf("eip is not found"))
		}
		return nil
	})
	if err != nil {
		return err
	}
	if *eip.AddressStatus != EIP_STATUS_UNBIND {
		return fmt.Errorf("eip status is illegal %s", *eip.AddressStatus)
	}

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId := v.(string)
		
		// 如果同时传了 instance_id 和 private_ip，则使用 AssociateAddress API（用于 BMS 等场景）
		if privateIp, hasPrivateIp := d.GetOk("private_ip"); hasPrivateIp {
			request := vpc.NewAssociateAddressRequest()
			request.AddressId = &eipId
			request.InstanceId = &instanceId
			privateIpStr := privateIp.(string)
			request.PrivateIpAddress = &privateIpStr
			
			err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				ratelimit.Check(request.GetAction())
				response, err := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AssociateAddress(request)
				if err != nil {
					log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
						logId, request.GetAction(), request.ToJsonString(), err.Error())
					return retryError(err)
				}
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
				return nil
			})
			if err != nil {
				return err
			}
		} else {
			// 只传了 instance_id，使用简化的封装方法（用于 CVM/CLB）
			err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				err := vpcService.AttachEip(ctx, eipId, instanceId)
				if err != nil {
					return retryError(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		
		associationId := fmt.Sprintf("%v::%v", eipId, instanceId)
		d.SetId(associationId)

		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			eip, errRet = vpcService.DescribeEipById(ctx, eipId)
			if errRet != nil {
				return retryError(errRet)
			}
			if eip == nil {
				return resource.NonRetryableError(fmt.Errorf("eip is not found"))
			}
			if *eip.AddressStatus == EIP_STATUS_BIND {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("wait for binding success: %s", *eip.AddressStatus))
		})
		if err != nil {
			return err
		}
		return resourceTencentCloudEipAssociationRead(d, meta)
	}

	needRequest := false
	request := vpc.NewAssociateAddressRequest()
	request.AddressId = &eipId
	var networkId string
	var privateIp string
	if v, ok := d.GetOk("network_interface_id"); ok {
		needRequest = true
		networkId = v.(string)
		request.NetworkInterfaceId = &networkId
	}
	if v, ok := d.GetOk("private_ip"); ok {
		needRequest = true
		privateIp = v.(string)
		request.PrivateIpAddress = &privateIp
	}
	if needRequest {
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			response, err := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AssociateAddress(request)
			if err != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), err.Error())
				return retryError(err)
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			return nil
		})
		if err != nil {
			return err
		}
		id := fmt.Sprintf("%v::%v::%v", eipId, networkId, privateIp)
		d.SetId(id)

		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			eip, errRet = vpcService.DescribeEipById(ctx, eipId)
			if errRet != nil {
				return retryError(errRet)
			}
			if eip == nil {
				return resource.NonRetryableError(fmt.Errorf("eip is not found"))
			}
			if *eip.AddressStatus == EIP_STATUS_BIND || *eip.AddressStatus == EIP_STATUS_BIND_ENI {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("wait for binding success: %s", *eip.AddressStatus))
		})
		if err != nil {
			return err
		}

		return resourceTencentCloudEipAssociationRead(d, meta)
	}

	// bind with havip
	if v, ok := d.GetOk("havip_id"); ok {
		havipId := v.(string)
		havipRequest := vpc.NewHaVipAssociateAddressIpRequest()
		havipRequest.HaVipId = &havipId
		havipRequest.AddressIp = eip.AddressIp

		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(havipRequest.GetAction())
			response, err := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().HaVipAssociateAddressIp(havipRequest)
			if err != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, havipRequest.GetAction(), havipRequest.ToJsonString(), err.Error())
				return retryError(err)
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, havipRequest.GetAction(), havipRequest.ToJsonString(), response.ToJsonString())
			return nil
		})
		if err != nil {
			return err
		}
		id := fmt.Sprintf("%v::%v", eipId, havipId)
		d.SetId(id)

		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			eip, errRet = vpcService.DescribeEipById(ctx, eipId)
			if errRet != nil {
				return retryError(errRet)
			}
			if eip == nil {
				return resource.NonRetryableError(fmt.Errorf("eip is not found"))
			}
			if *eip.AddressStatus == EIP_STATUS_BIND || *eip.AddressStatus == EIP_STATUS_BIND_ENI {
				return nil
			}
			return resource.RetryableError(fmt.Errorf("wait for binding success: %s", *eip.AddressStatus))
		})
		if err != nil {
			return err
		}

		return resourceTencentCloudEipAssociationRead(d, meta)
	}

	return nil
}

func resourceTencentCloudEipAssociationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_eip_association.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	vpcService := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	id := d.Id()
	association, err := parseEipAssociationId(id)
	if err != nil {
		return err
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		eip, errRet := vpcService.DescribeEipById(ctx, association.EipId)
		if errRet != nil {
			return retryError(errRet)
		}
		if eip == nil {
			d.SetId("")
			return nil
		}

		// Determine the binding type based on BindResourceType or InstanceId prefix
		if eip.BindResourceType != nil && *eip.BindResourceType == "BIND_HAVIP" {
			// HAVIP binding case
			if eip.InstanceId != nil {
				_ = d.Set("havip_id", *eip.InstanceId)
			}
		} else if eip.InstanceId != nil && strings.HasPrefix(*eip.InstanceId, "havip-") {
			// Fallback: Check InstanceId prefix for HAVIP
			_ = d.Set("havip_id", *eip.InstanceId)
		} else if eip.InstanceId != nil {
			// CVM/CLB instance binding case
			_ = d.Set("instance_id", *eip.InstanceId)
		} else if eip.NetworkInterfaceId != nil {
			// ENI binding case
			_ = d.Set("network_interface_id", *eip.NetworkInterfaceId)
			if eip.PrivateAddressIp != nil {
				_ = d.Set("private_ip", *eip.PrivateAddressIp)
			}
		}
		
		return nil
	})
	if err != nil {
		return err
	}

	_ = d.Set("eip_id", association.EipId)
	return nil
}

func resourceTencentCloudEipAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_eip_association.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	vpcService := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	id := d.Id()
	association, err := parseEipAssociationId(id)
	if err != nil {
		return err
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := vpcService.UnattachEip(ctx, association.EipId)
		if err != nil {
			return retryError(err, "DesOperation.MutexTaskRunning")
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

type EipAssociationId struct {
	EipId              string
	InstanceId         string
	NetworkInterfaceId string
	PrivateIp          string
	HaVipId            string
}

func parseEipAssociationId(associationId string) (association EipAssociationId, errRet error) {
	ids := strings.Split(associationId, "::")
	if len(ids) < 2 || len(ids) > 3 {
		errRet = fmt.Errorf("Invalid eip association ID: %v", associationId)
		return
	}
	association.EipId = ids[0]

	// associate with instance or havip (2 parts)
	if len(ids) == 2 {
		// Check if it's a havip id (starts with "havip-")
		if strings.HasPrefix(ids[1], "havip-") {
			association.HaVipId = ids[1]
		} else {
			association.InstanceId = ids[1]
		}
		return
	}

	// associate with network interface (3 parts)
	association.NetworkInterfaceId = ids[1]
	association.PrivateIp = ids[2]
	return
}
