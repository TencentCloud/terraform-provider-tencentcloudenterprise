/*
Provide a resource to create a KubernetesClusterEndpoint. This resource allows you to create an empty cluster first without any workers. Only all attached node depends create complete, cluster endpoint will finally be enabled.

~> **NOTE:** Recommend using `depends_on` to make sure endpoint create after node pools or workers does.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_cluster_endpoint" "foo" {
	  cluster_id = "cls-xxxxxxxx"
	  cluster_internet = true
	  cluster_intranet = true
	  cluster_internet_security_group = "sg-xxxxxxxx"
	  cluster_intranet_subnet_id = "subnet-xxxxxxxx"
	}

# Use existed CLB

	resource "tencentcloudenterprise_tke_kubernetes_cluster_endpoint" "with_clb" {
	  cluster_id = "cls-xxxxxxxx"
	  cluster_internet = true
	  cluster_internet_security_group = "sg-xxxxxxxx"
	  existed_load_balancer_id = "lb-xxxxxxxx"
	}

```

# Import

KubernetesClusterEndpoint instance can be imported by passing cluster id, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_cluster_endpoint.test cluster-id
```
*/
package tencentcloud

import (
	"context"
	"fmt"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"

	"github.com/hashicorp/go-multierror"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_endpoint", CNDescription{
		TerraformTypeCN: "集群访问端点",
		DescriptionCN:   "提供Kubernetes集群访问端点资源，用于创建和管理集群的访问端点。",
		AttributesCN: map[string]string{
			"cluster_id":                      "集群 ID",
			"cluster_internet":                "是否开启公网访问",
			"cluster_intranet":                "是否开启内网访问",
			"cluster_internet_security_group": "公网访问安全组",
			"cluster_internet_domain":         "公网访问域名",
			"extensive_parameters":            "公网访问扩展参数",
			"cluster_intranet_domain":         "内网访问域名",
			"cluster_intranet_subnet_id":      "内网访问子网 ID",
			"existed_load_balancer_id":        "使用已有的负载均衡开启访问",
			"cluster_external_endpoint":       "公网访问地址",
			"cluster_intranet_endpoint":       "内网访问地址",
			"cluster_domain":                  "集群域名",
			"certification_authority":         "CA证书",
			"cluster_external_acl":            "公网访问ACL列表",
		},
	})
}
func resourceTencentCloudTkeClusterEndpoint() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a KubernetesClusterEndpoint. This resource allows you to create an empty cluster first without any workers. Only all attached node depends create complete, cluster endpoint will finally be enabled.",
		Create:      resourceTencentCloudTkeClusterEndpointCreate,
		Read:        resourceTencentCloudTkeClusterEndpointRead,
		Update:      resourceTencentCloudTkeClusterEndpointUpdate,
		Delete:      resourceTencentCloudTkeClusterEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specify cluster ID.",
			},
			"cluster_internet": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Open internet access or not.",
			},
			"cluster_intranet": {
				Type:        schema.TypeBool,
				Default:     false,
				Optional:    true,
				Description: "Open intranet access or not.",
			},
			"cluster_internet_security_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify security group, NOTE: This argument must not be empty if cluster internet enabled.",
			},
			"cluster_internet_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Domain name for cluster Kube-apiserver internet access. " +
					" Be careful if you modify value of this parameter, the cluster_external_endpoint value may be changed automatically too.",
			},
			"extensive_parameters": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "The LB parameter. Only used for public network access.",
			},
			"cluster_intranet_domain": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Domain name for cluster Kube-apiserver intranet access." +
					" Be careful if you modify value of this parameter, the cluster_intranet_endpoint value may be changed automatically too.",
			},
			"cluster_intranet_subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Subnet id who can access this independent cluster, this field must and can only set  when `cluster_intranet` is true." +
					" `cluster_intranet_subnet_id` can not modify once be set.",
			},
			"existed_load_balancer_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Use existed CLB to enable internet/intranet access.",
			},
			// Computed
			"cluster_external_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External network address to access.",
			},
			"cluster_intranet_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Intranet network address to access.",
			},
			"cluster_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster domain name.",
			},
			"certification_authority": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CA certificate used for access.",
			},
			"cluster_external_acl": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "External network access ACL list.",
			},
		},
	}
}

func resourceTencentCloudTkeClusterEndpointRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster_endpoint.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client}

	id := d.Id()
	_, has, err := service.DescribeCluster(ctx, id)
	if err != nil {
		d.SetId("")
		return err
	}
	if !has {
		d.SetId("")
		return fmt.Errorf("cluster %s not found", id)
	}

	d.SetId(id)
	_ = d.Set("cluster_id", id)

	// Use DescribeClusterEndpoints to get endpoint info
	response, err := service.DescribeClusterEndpoints(ctx, id)
	if err != nil {
		return err
	}

	if response != nil && response.Response != nil {
		resp := response.Response

		// Set computed fields
		if resp.ClusterExternalEndpoint != nil {
			_ = d.Set("cluster_external_endpoint", resp.ClusterExternalEndpoint)
		}
		if resp.ClusterIntranetEndpoint != nil {
			_ = d.Set("cluster_intranet_endpoint", resp.ClusterIntranetEndpoint)
		}
		if resp.ClusterDomain != nil {
			_ = d.Set("cluster_domain", resp.ClusterDomain)
		}
		if resp.CertificationAuthority != nil {
			_ = d.Set("certification_authority", resp.CertificationAuthority)
		}
		if resp.ClusterExternalACL != nil {
			acls := make([]string, 0, len(resp.ClusterExternalACL))
			for _, acl := range resp.ClusterExternalACL {
				if acl != nil {
					acls = append(acls, *acl)
				}
			}
			_ = d.Set("cluster_external_acl", acls)
		}

		// Optionally set input fields from response if not set
		if resp.SecurityGroup != nil && d.Get("cluster_internet_security_group").(string) == "" {
			_ = d.Set("cluster_internet_security_group", resp.SecurityGroup)
		}
		if resp.ClusterIntranetSubnetId != nil && d.Get("cluster_intranet_subnet_id").(string) == "" {
			_ = d.Set("cluster_intranet_subnet_id", resp.ClusterIntranetSubnetId)
		}
		if resp.ClusterExternalDomain != nil && d.Get("cluster_internet_domain").(string) == "" {
			_ = d.Set("cluster_internet_domain", resp.ClusterExternalDomain)
		}
		if resp.ClusterIntranetDomain != nil && d.Get("cluster_intranet_domain").(string) == "" {
			_ = d.Set("cluster_intranet_domain", resp.ClusterIntranetDomain)
		}
	}

	return nil
}

func resourceTencentCloudTkeClusterEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster_endpoint.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client}

	id := d.Get("cluster_id").(string)
	var (
		clusterInternet              = d.Get("cluster_internet").(bool)
		clusterIntranet              = d.Get("cluster_intranet").(bool)
		intranetSubnetId             = d.Get("cluster_intranet_subnet_id").(string)
		clusterInternetSecurityGroup = d.Get("cluster_internet_security_group").(string)
		clusterInternetDomain        = d.Get("cluster_internet_domain").(string)
		clusterIntranetDomain        = d.Get("cluster_intranet_domain").(string)
		extensiveParameters          = d.Get("extensive_parameters").(string)
		existedLoadBalancerId        = d.Get("existed_load_balancer_id").(string)
	)

	if clusterIntranet && intranetSubnetId == "" {
		return fmt.Errorf("`cluster_intranet_subnet_id` must set when `cluster_intranet` is true")
	}
	if !clusterIntranet && intranetSubnetId != "" {
		return fmt.Errorf("`cluster_intranet_subnet_id` can only set when `cluster_intranet` is true")
	}

	// Create Intranet(Private) Network
	if clusterIntranet {
		err := tencentCloudClusterIntranetSwitch(ctx, &service, id, intranetSubnetId, true, clusterIntranetDomain, existedLoadBalancerId)
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, true, false)
		if err != nil {
			return err
		}
	}

	// Open the internet
	if clusterInternet {
		err := tencentCloudClusterInternetSwitch(ctx, &service, id, true, clusterInternetSecurityGroup, clusterInternetDomain, extensiveParameters, existedLoadBalancerId)
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, true, true)
		if err != nil {
			return err
		}
	}

	d.SetId(id)

	return resourceTencentCloudTkeClusterEndpointRead(d, meta)
}

func resourceTencentCloudTkeClusterEndpointUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster_endpoint.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client}
	id := d.Id()

	var (
		clusterInternet              = d.Get("cluster_internet").(bool)
		clusterIntranet              = d.Get("cluster_intranet").(bool)
		clusterInternetSecurityGroup = d.Get("cluster_internet_security_group").(string)
		clusterInternetDomain        = d.Get("cluster_internet_domain").(string)
		clusterIntranetDomain        = d.Get("cluster_intranet_domain").(string)
		subnetId                     = d.Get("cluster_intranet_subnet_id").(string)
		extensiveParameters          = d.Get("extensive_parameters").(string)
	)

	var err error

	if d.HasChange("cluster_internet_security_group") && !d.HasChange("cluster_internet") {
		if clusterInternet {
			err := service.ModifyClusterEndpointSG(ctx, id, clusterInternetSecurityGroup)
			if err != nil {
				return err
			}
		}
	}

	if d.HasChange("cluster_internet") {
		err = tencentCloudClusterInternetSwitch(ctx, &service, id, clusterInternet, clusterInternetSecurityGroup, clusterInternetDomain, extensiveParameters, "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, clusterInternet, true)
		if err != nil {
			return err
		}
	} else if clusterInternet && d.HasChange("cluster_internet_domain") {
		// only domain changed, need to close and reopen
		// close
		err = tencentCloudClusterInternetSwitch(ctx, &service, id, false, clusterInternetSecurityGroup, clusterInternetDomain, "", "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, false, true)
		if err != nil {
			return err
		}
		// reopen
		err = tencentCloudClusterInternetSwitch(ctx, &service, id, true, clusterInternetSecurityGroup, clusterInternetDomain, extensiveParameters, "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, true, true)
		if err != nil {
			return err
		}
	}

	if d.HasChange("cluster_intranet") {
		err = tencentCloudClusterIntranetSwitch(ctx, &service, id, subnetId, clusterIntranet, clusterIntranetDomain, "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, clusterIntranet, false)
		if err != nil {
			return err
		}
	} else if clusterIntranet && d.HasChange("cluster_intranet_domain") {
		// only domain changed, need to close and reopen
		// close
		err = tencentCloudClusterIntranetSwitch(ctx, &service, id, subnetId, false, clusterIntranetDomain, "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, false, false)
		if err != nil {
			return err
		}
		// reopen
		err = tencentCloudClusterIntranetSwitch(ctx, &service, id, subnetId, true, clusterIntranetDomain, "")
		if err != nil {
			return err
		}
		err = waitForClusterEndpointFinish(ctx, &service, id, true, false)
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudTkeClusterEndpointRead(d, meta)
}

func resourceTencentCloudTkeClusterEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_cluster_endpoint.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	client := meta.(*TencentCloudClient).apiV3Conn
	service := TkeService{client}
	var (
		id   = d.Id()
		errs multierror.Error
	)

	// Use DescribeClusterEndpoints to check current status
	response, err := service.DescribeClusterEndpoints(ctx, id)
	if err != nil {
		return err
	}

	var (
		clusterInternet = false
		clusterIntranet = false
	)

	if response != nil && response.Response != nil {
		clusterInternet = response.Response.ClusterExternalEndpoint != nil && *response.Response.ClusterExternalEndpoint != ""
		clusterIntranet = response.Response.ClusterIntranetEndpoint != nil && *response.Response.ClusterIntranetEndpoint != ""
	}

	if clusterInternet {
		err = tencentCloudClusterInternetSwitch(ctx, &service, id, false, "", "", "", "")
		if err != nil {
			errs = *multierror.Append(err)
		} else {
			taskErr := waitForClusterEndpointFinish(ctx, &service, id, false, true)
			if taskErr != nil {
				errs = *multierror.Append(taskErr)
			}
		}
	}

	if clusterIntranet {
		err = tencentCloudClusterIntranetSwitch(ctx, &service, id, "", false, "", "")
		if err != nil {
			errs = *multierror.Append(err)
		} else {
			taskErr := waitForClusterEndpointFinish(ctx, &service, id, false, false)
			if taskErr != nil {
				errs = *multierror.Append(taskErr)
			}
		}
	}

	return errs.ErrorOrNil()
}

func waitForClusterEndpointFinish(ctx context.Context, service *TkeService, id string, enabled bool, isInternet bool) (err error) {
	return resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
		var (
			status       string
			message      string
			inErr        error
			finishStates []string
			failedStates []string
		)

		if enabled {
			finishStates = []string{TkeInternetStatusCreated}
			failedStates = []string{TkeInternetStatusCreateFailed}
		} else {
			finishStates = []string{TkeInternetStatusNotfound, TkeInternetStatusDeleted}
			failedStates = []string{TkeInternetStatusDeletedFailed}
		}

		status, message, inErr = service.DescribeClusterEndpointStatus(ctx, id, isInternet)

		if inErr != nil {
			return retryError(inErr)
		}

		// Check if finished successfully
		if IsContains(finishStates, status) {
			return nil
		}

		// Check if failed
		if IsContains(failedStates, status) {
			return resource.NonRetryableError(
				fmt.Errorf("cluster %s endpoint operation failed, status: %s, message: %s", id, status, message))
		}

		// Still in progress (Creating/Deleting/NotFound for create)
		return resource.RetryableError(
			fmt.Errorf("cluster %s endpoint status is %s, waiting...", id, status))
	})
}

func tencentCloudClusterInternetSwitch(ctx context.Context, service *TkeService, id string, enable bool, sg string, domain string, extensiveParameters string, existedLoadBalancerId string) (err error) {
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if enable {
			err = service.CreateClusterEndpoint(ctx, id, "", sg, true, domain, extensiveParameters, existedLoadBalancerId)
			if err != nil {
				return retryError(err, tke.RESOURCEUNAVAILABLE_CLUSTERSTATE)
			}
		} else {
			err = service.DeleteClusterEndpoint(ctx, id, true)
			if err != nil {
				return retryError(err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func tencentCloudClusterIntranetSwitch(ctx context.Context, service *TkeService, id, subnetId string, enable bool, domain string, existedLoadBalancerId string) (err error) {
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if enable {
			err = service.CreateClusterEndpoint(ctx, id, subnetId, "", false, domain, "", existedLoadBalancerId)
			if err != nil {
				return retryError(err, tke.RESOURCEUNAVAILABLE_CLUSTERSTATE)
			}
		} else {
			err = service.DeleteClusterEndpoint(ctx, id, false)
			if err != nil {
				return retryError(err)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
