package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"net/url"
	"os"
	"strconv"
	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

const (
	PROVIDER_SECRET_ID                    = "TENCENTCLOUD_SECRET_ID"
	PROVIDER_SECRET_KEY                   = "TENCENTCLOUD_SECRET_KEY"
	PROVIDER_SECURITY_TOKEN               = "TENCENTCLOUD_SECURITY_TOKEN"
	PROVIDER_REGION                       = "TENCENTCLOUD_REGION"
	PROVIDER_PROTOCOL                     = "TENCENTCLOUD_PROTOCOL"
	PROVIDER_DOMAIN                       = "TENCENTCLOUD_DOMAIN"
	PROVIDER_ASSUME_ROLE_ARN              = "TENCENTCLOUD_ASSUME_ROLE_ARN"
	PROVIDER_ASSUME_ROLE_SESSION_NAME     = "TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME"
	PROVIDER_ASSUME_ROLE_SESSION_DURATION = "TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION"
	PROVIDER_CSP_DOMAIN                   = "TENCENTCLOUD_CSP_DOMAIN"
	PROVIDER_COS_DOMAIN                   = "TENCENTCLOUD_COS_DOMAIN"
)

type TencentCloudClient struct {
	apiV3Conn *connectivity.TencentCloudClient
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "This is the TencentCloud access key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_ID` environment variable.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "This is the TencentCloud secret key. It must be provided, but it can also be sourced from the `TENCENTCLOUD_SECRET_KEY` environment variable.",
				Sensitive:   true,
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECURITY_TOKEN, nil),
				Description: "TencentCloud Security Token of temporary access credentials. It can be sourced from the `TENCENTCLOUD_SECURITY_TOKEN` environment variable. Notice: for supported products, please refer to: [temporary key supported products](https://intl.cloud.tencent.com/document/product/598/10588).",
				Sensitive:   true,
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description:  "This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is ap-guangzhou.",
				InputDefault: "ap-guangzhou",
			},
			"protocol": {
				Type:         schema.TypeString,
				Optional:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_PROTOCOL, "HTTP"),
				ValidateFunc: validateAllowedStringValue([]string{"HTTP", "HTTPS"}),
				Description:  "The protocol of the API request. Valid values: `HTTP` and `HTTPS`. Default is `HTTP`.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_DOMAIN, nil),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"csp_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_CSP_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"cos_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_COS_DOMAIN, ""),
				Description: "The root domain of the API request, Default is `tencentcloudapi.com`.",
			},
			"assume_role": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "The `assume_role` block. If provided, terraform will attempt to assume this role using the supplied credentials.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_ARN, nil),
							Description: "The ARN of the role to assume. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_ARN`.",
						},
						"session_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc(PROVIDER_ASSUME_ROLE_SESSION_NAME, nil),
							Description: "The session name to use when making the AssumeRole call. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_NAME`.",
						},
						"session_duration": {
							Type:     schema.TypeInt,
							Required: true,
							DefaultFunc: func() (interface{}, error) {
								if v := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); v != "" {
									return strconv.Atoi(v)
								}
								return 7200, nil
							},
							ValidateFunc: validateIntegerInRange(0, 43200),
							Description:  "The duration of the session when making the AssumeRole call. Its value ranges from 0 to 43200(seconds), and default is 7200 seconds. It can be sourced from the `TENCENTCLOUD_ASSUME_ROLE_SESSION_DURATION`.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A more restrictive policy when making the AssumeRole call. Its content must not contains `principal` elements. Notice: more syntax references, please refer to: [policies syntax logic](https://intl.cloud.tencent.com/document/product/598/10603).",
						},
					},
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tencentcloudenterprise_cam_users":                  dataSourceTencentCloudCamUsers(),
			"tencentcloudenterprise_tdmq_rabbitmq_vip_instance": dataSourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloudenterprise_cvm_instances":              dataSourceTencentCloudInstances(),
			"tencentcloudenterprise_clb_instances":              dataSourceTencentCloudClbInstances(),
			"tencentcloudenterprise_cbs_storages":               dataSourceTencentCloudCbsStorages(),
			"tencentcloudenterprise_cfs_file_systems":           dataSourceTencentCloudCfsFileSystems(),
			"tencentcloudenterprise_redis_instances":            dataSourceTencentRedisInstances(),
			"tencentcloudenterprise_mariadb_db_instances":       dataSourceTencentCloudMariadbDbInstances(),
			"tencentcloudenterprise_tbase_pg_instances":         dataSourceTencentCloudTbasePGInstances(),
			"tencentcloudenterprise_tdmq_rocketmq_cluster":      dataSourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloudenterprise_ckafka_instances":           dataSourceTencentCloudCkafkaInstances(),
			"tencentcloudenterprise_availability_zones":         dataSourceTencentCloudAvailabilityZones(),
			"tencentcloudenterprise_cvm_image":                  dataSourceTencentCloudImage(),
			"tencentcloudenterprise_cvm_images":                 dataSourceTencentCloudImages(),
			"tencentcloudenterprise_vpc_subnets":                dataSourceTencentCloudVpcSubnets(),
			"tencentcloudenterprise_vpc_instances":              dataSourceTencentCloudVpcInstances(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"tencentcloudenterprise_cvm_instance":               resourceTencentCloudInstance(),
			"tencentcloudenterprise_clb_instance":               resourceTencentCloudClbInstance(),
			"tencentcloudenterprise_cbs_storage":                resourceTencentCloudCbsStorage(),
			"tencentcloudenterprise_cfs_file_system":            resourceTencentCloudCfsFileSystem(),
			"tencentcloudenterprise_redis_instance":             resourceTencentCloudRedisInstance(),
			"tencentcloudenterprise_mariadb_hour_db_instance":   resourceTencentCloudMariadbHourDbInstance(),
			"tencentcloudenterprise_tbase_pg_instance":          resourceTencentCloudTbasePGInstance(),
			"tencentcloudenterprise_cam_user":                   resourceTencentCloudCamUser(),
			"tencentcloudenterprise_tdmq_rocketmq_cluster":      resourceTencentCloudTdmqRocketmqCluster(),
			"tencentcloudenterprise_tdmq_rabbitmq_vip_instance": resourceTencentCloudTdmqRabbitmqVipInstance(),
			"tencentcloudenterprise_ckafka_instance":            resourceTencentCloudCkafkaInstance(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	secretId := d.Get("secret_id").(string)
	secretKey := d.Get("secret_key").(string)
	securityToken := d.Get("security_token").(string)
	region := d.Get("region").(string)
	protocol := d.Get("protocol").(string)
	domain := d.Get("domain").(string)
	cspDomain := d.Get("csp_domain").(string)
	cosDomain := d.Get("cos_domain").(string)
	// standard client
	var tcClient TencentCloudClient
	tcClient.apiV3Conn = &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		CredentialTce: common2.NewTokenCredential(
			secretId,
			secretKey,
			securityToken,
		),
		Region:    region,
		Protocol:  protocol,
		Domain:    domain,
		CspDomain: cspDomain,
		CosDomain: cosDomain,
	}

	envRoleArn := os.Getenv(PROVIDER_ASSUME_ROLE_ARN)
	envSessionName := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_NAME)

	// get assume role from env
	if envRoleArn != "" && envSessionName != "" {
		var assumeRoleSessionDuration int
		if envSessionDuration := os.Getenv(PROVIDER_ASSUME_ROLE_SESSION_DURATION); envSessionDuration != "" {
			var err error
			assumeRoleSessionDuration, err = strconv.Atoi(envSessionDuration)
			if err != nil {
				return nil, diag.FromErr(err)
			}
		}
		if assumeRoleSessionDuration == 0 {
			assumeRoleSessionDuration = 7200
		}

		_ = genClientWithSTS(&tcClient, envRoleArn, envSessionName, assumeRoleSessionDuration, "")
	}

	// get assume role from tf config
	assumeRoleList := d.Get("assume_role").(*schema.Set).List()
	if len(assumeRoleList) == 1 {
		assumeRole := assumeRoleList[0].(map[string]interface{})
		assumeRoleArn := assumeRole["role_arn"].(string)
		assumeRoleSessionName := assumeRole["session_name"].(string)
		assumeRoleSessionDuration := assumeRole["session_duration"].(int)
		assumeRolePolicy := assumeRole["policy"].(string)

		_ = genClientWithSTS(&tcClient, assumeRoleArn, assumeRoleSessionName, assumeRoleSessionDuration, assumeRolePolicy)
	}
	return &tcClient, nil
}

func genClientWithSTS(tcClient *TencentCloudClient, assumeRoleArn, assumeRoleSessionName string,
	assumeRoleSessionDuration int, assumeRolePolicy string) error {
	// applying STS credentials
	request := sts.NewAssumeRoleRequest()
	request.RoleArn = helper.String(assumeRoleArn)
	request.RoleSessionName = helper.String(assumeRoleSessionName)
	request.DurationSeconds = helper.IntUint64(assumeRoleSessionDuration)
	if assumeRolePolicy != "" {
		request.Policy = helper.String(url.QueryEscape(assumeRolePolicy))
	}
	ratelimit.Check(request.GetAction())
	response, err := tcClient.apiV3Conn.UseStsClient().AssumeRole(request)
	if err != nil {
		return err
	}
	// using STS credentials
	tcClient.apiV3Conn.Credential = common.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)
	tcClient.apiV3Conn.CredentialTce = common2.NewTokenCredential(
		*response.Response.Credentials.TmpSecretId,
		*response.Response.Credentials.TmpSecretKey,
		*response.Response.Credentials.Token,
	)
	return nil
}
