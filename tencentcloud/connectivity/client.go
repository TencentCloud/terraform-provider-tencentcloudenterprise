package connectivity

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	mariadb "terraform-provider-tencentcloudenterprise/sdk/mariadb/v20170312"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	"os"
	"strconv"
	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	cbs "terraform-provider-tencentcloudenterprise/sdk/cbs/v20170312"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	profile2 "terraform-provider-tencentcloudenterprise/sdk/common/profile"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
)

const (
	PROVIDER_CVM_REQUEST_TIMEOUT = "TENCENTCLOUD_CVM_REQUEST_TIMEOUT"
	PROVIDER_CBS_REQUEST_TIMEOUT = "TENCENTCLOUD_CBS_REQUEST_TIMEOUT"
)

// TencentCloudClient is client for all TencentCloud service
type TencentCloudClient struct {
	Credential    *common.Credential
	CredentialTce *common2.Credential
	Region        string
	Protocol      string
	Domain        string
	CspDomain     string
	CosDomain     string
	redisConn     *redis.Client
	vpcConn       *vpc.Client
	cbsConn       *cbs.Client
	cvmConn       *cvm.Client
	clbConn       *clb.Client
	tagConn       *tag.Client
	tdmqConn      *tdmq.Client
	camConn       *cam.Client
	stsConn       *sts.Client
	cfsConn       *cfs.Client
	ckafkaConn    *ckafka.Client
	mariadbConn   *mariadb.Client
	tbaseConn     *tbase.Client
	tkeConn       *tke.Client
}

// NewClientProfile returns a new ClientProfile
func (me *TencentCloudClient) NewClientProfile(timeout int) *profile.ClientProfile {
	cpf := profile.NewClientProfile()

	// all request use method POST
	cpf.HttpProfile.ReqMethod = "POST"
	// request timeout
	cpf.HttpProfile.ReqTimeout = timeout
	// request protocol
	cpf.HttpProfile.Scheme = me.Protocol
	// request domain
	cpf.HttpProfile.Endpoint = me.Domain
	// default language
	cpf.Language = "en-US"

	return cpf
}

// NewClientProfileTce returns a new ClientProfile
func (me *TencentCloudClient) NewClientProfileTce(timeout int) *profile2.ClientProfile {
	cpf := profile2.NewClientProfile()

	// all request use method POST
	cpf.HttpProfile.ReqMethod = "POST"
	// request timeout
	cpf.HttpProfile.ReqTimeout = timeout
	// request protocol
	cpf.HttpProfile.Scheme = me.Protocol
	// request domain
	cpf.HttpProfile.RootDomain = me.Domain
	// default language
	cpf.Language = "en-US"

	return cpf
}

// UseTbaseClient returns tbase client for service
func (me *TencentCloudClient) UseTbaseClient() *tbase.Client {
	if me.tbaseConn != nil {
		return me.tbaseConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tbaseConn, _ = tbase.NewClient(me.CredentialTce, me.Region, cpf)
	me.tbaseConn.WithHttpTransport(&LogRoundTripper{})
	return me.tbaseConn
}

// UseRedisClient returns redis client for service
func (me *TencentCloudClient) UseRedisClient() *redis.Client {
	if me.redisConn != nil {
		return me.redisConn
	}

	cpf := me.NewClientProfileTce(300)
	me.redisConn, _ = redis.NewClient(me.CredentialTce, me.Region, cpf)
	me.redisConn.WithHttpTransport(&LogRoundTripper{})

	return me.redisConn
}

// UseVpcClient returns vpc client for service
func (me *TencentCloudClient) UseVpcClient() *vpc.Client {
	if me.vpcConn != nil {
		return me.vpcConn
	}

	cpf := me.NewClientProfileTce(300)
	me.vpcConn, _ = vpc.NewClient(me.CredentialTce, me.Region, cpf)
	me.vpcConn.WithHttpTransport(&LogRoundTripper{})

	return me.vpcConn
}

// UseCbsClient returns cbs client for service
func (me *TencentCloudClient) UseCbsClient() *cbs.Client {
	if me.cbsConn != nil {
		return me.cbsConn
	}

	var reqTimeout = getEnvDefault(PROVIDER_CBS_REQUEST_TIMEOUT, 300)
	cpf := me.NewClientProfileTce(reqTimeout)
	me.cbsConn, _ = cbs.NewClient(me.CredentialTce, me.Region, cpf)
	me.cbsConn.WithHttpTransport(&LogRoundTripper{})

	return me.cbsConn
}

// UseClbClient returns clb client for service
func (me *TencentCloudClient) UseClbClient() *clb.Client {
	if me.clbConn != nil {
		return me.clbConn
	}

	cpf := me.NewClientProfileTce(300)
	me.clbConn, _ = clb.NewClient(me.CredentialTce, me.Region, cpf)
	me.clbConn.WithHttpTransport(&LogRoundTripper{})

	return me.clbConn
}

// UseCvmClient returns cvm client for service
func (me *TencentCloudClient) UseCvmClient() *cvm.Client {
	if me.cvmConn != nil {
		return me.cvmConn
	}

	var reqTimeout = getEnvDefault(PROVIDER_CVM_REQUEST_TIMEOUT, 300)
	cpf := me.NewClientProfileTce(reqTimeout)
	me.cvmConn, _ = cvm.NewClient(me.CredentialTce, me.Region, cpf)
	me.cvmConn.WithHttpTransport(&LogRoundTripper{})

	return me.cvmConn
}

// UseTagClient returns tag client for service
func (me *TencentCloudClient) UseTagClient() *tag.Client {
	if me.tagConn != nil {
		return me.tagConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tagConn, _ = tag.NewClient(me.CredentialTce, me.Region, cpf)
	me.tagConn.WithHttpTransport(&LogRoundTripper{})

	return me.tagConn
}

// UseTdmqClient returns Tdmq client for service
func (me *TencentCloudClient) UseTdmqClient() *tdmq.Client {
	if me.tdmqConn != nil {
		return me.tdmqConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tdmqConn, _ = tdmq.NewClient(me.CredentialTce, me.Region, cpf)
	me.tdmqConn.WithHttpTransport(&LogRoundTripper{})

	return me.tdmqConn
}

// UseCamClient returns cam client for service
func (me *TencentCloudClient) UseCamClient() *cam.Client {
	if me.camConn != nil {
		return me.camConn
	}

	cpf := me.NewClientProfileTce(300)
	me.camConn, _ = cam.NewClient(me.CredentialTce, me.Region, cpf)
	me.camConn.WithHttpTransport(&LogRoundTripper{})

	return me.camConn
}

// UseStsClient returns sts client for service
func (me *TencentCloudClient) UseStsClient() *sts.Client {
	/*
		me.Credential will changed, don't cache it
		if me.stsConn != nil {
			return me.stsConn
		}
	*/

	cpf := me.NewClientProfile(300)
	if me.Domain != "" {
		cpf.HttpProfile.Endpoint = "sts." + me.Domain
	}
	me.stsConn, _ = sts.NewClient(me.Credential, me.Region, cpf)
	me.stsConn.WithHttpTransport(&LogRoundTripper{})

	return me.stsConn
}

// UseCfsClient returns cfs client for service
func (me *TencentCloudClient) UseCfsClient() *cfs.Client {
	if me.cfsConn != nil {
		return me.cfsConn
	}

	cpf := me.NewClientProfileTce(300)
	me.cfsConn, _ = cfs.NewClient(me.CredentialTce, me.Region, cpf)
	me.cfsConn.WithHttpTransport(&LogRoundTripper{})

	return me.cfsConn
}

// UseCkafkaClient returns ckafka client for service
func (me *TencentCloudClient) UseCkafkaClient() *ckafka.Client {
	if me.ckafkaConn != nil {
		return me.ckafkaConn
	}

	cpf := me.NewClientProfileTce(300)
	me.ckafkaConn, _ = ckafka.NewClient(me.CredentialTce, me.Region, cpf)
	me.ckafkaConn.WithHttpTransport(&LogRoundTripper{})

	return me.ckafkaConn
}

// UseMariadbClient returns mariadb client for service
func (me *TencentCloudClient) UseMariadbClient() *mariadb.Client {
	if me.mariadbConn != nil {
		return me.mariadbConn
	}

	cpf := me.NewClientProfileTce(300)
	if me.Domain != "" {
		cpf.HttpProfile.Endpoint = "mariadb." + me.Domain
	}
	me.mariadbConn, _ = mariadb.NewClient(me.CredentialTce, me.Region, cpf)
	me.mariadbConn.WithHttpTransport(&LogRoundTripper{})

	return me.mariadbConn
}

// UseTkeClient returns tke client for service
func (me *TencentCloudClient) UseTkeClient() *tke.Client {
	if me.tkeConn != nil {
		return me.tkeConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tkeConn, _ = tke.NewClient(me.CredentialTce, me.Region, cpf)
	me.tkeConn.WithHttpTransport(&LogRoundTripper{})

	return me.tkeConn
}

func getEnvDefault(key string, defVal int) int {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	timeOut, err := strconv.Atoi(val)
	if err != nil {
		panic("TENCENTCLOUD_XXX_REQUEST_TIMEOUT must be int.")
	}
	return timeOut
}
