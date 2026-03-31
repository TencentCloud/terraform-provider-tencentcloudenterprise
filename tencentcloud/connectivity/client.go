package connectivity

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	profile2 "terraform-provider-tencentcloudenterprise/sdk/common/profile"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/tencentyun/cos-go-sdk-v5"

	account20181225 "terraform-provider-tencentcloudenterprise/sdk/account/v20181225"
	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	apm "terraform-provider-tencentcloudenterprise/sdk/apm/v20210622"
	as "terraform-provider-tencentcloudenterprise/sdk/as/v20180419"
	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	bms "terraform-provider-tencentcloudenterprise/sdk/bms/v20180813"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	cbs "terraform-provider-tencentcloudenterprise/sdk/cbs/v20170312"
	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	cdc "terraform-provider-tencentcloudenterprise/sdk/cdc/v20201214"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	cfw "terraform-provider-tencentcloudenterprise/sdk/cfw/v20190904"
	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	csp "terraform-provider-tencentcloudenterprise/sdk/csp/v20200107"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	cwp "terraform-provider-tencentcloudenterprise/sdk/cwp/v20180228"
	dc "terraform-provider-tencentcloudenterprise/sdk/dc/v20180410"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	drc "terraform-provider-tencentcloudenterprise/sdk/drc/v20230615"
	location "terraform-provider-tencentcloudenterprise/sdk/location/v20191128"
	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	open "terraform-provider-tencentcloudenterprise/sdk/open/v20201202"
	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
	tbase "terraform-provider-tencentcloudenterprise/sdk/tbase/v20190107"
	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	tke2 "terraform-provider-tencentcloudenterprise/sdk/tke/v20220501"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	intlProfile "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/profile"
	mdl "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/mdl/v20200326"
	antiddos "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos/v20200309"
	api "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/api/v20201106"
	apigateway "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway/v20180808"
	cat "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cat/v20180409"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	cdn "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn/v20180606"
	chdfs "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs/v20201112"
	audit "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit/v20190319"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cynosdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb/v20190107"
	dayu "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dayu/v20180709"
	dbbrain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain/v20210527"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	domain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain/v20180808"
	dts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts/v20211206"
	emr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr/v20190103"
	es "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es/v20180416"
	gaap "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap/v20180529"
	kms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms/v20190118"
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"
	css "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live/v20180801"
	mariadb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb/v20170312"
	mongodb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"
	monitor "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor/v20180724"
	mps "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps/v20190612"
	postgre "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres/v20170312"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	pts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts/v20210728"
	rum "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum/v20210622"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	sqlserver "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver/v20180328"
	sslCertificate "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
	ssm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm/v20190923"
	sts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts/v20180813"
	tat "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat/v20201028"
	tcaplusdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb/v20190823"
	tcm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcm/v20210413"
	tdcpg "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdcpg/v20211118"
	teo "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo/v20220901"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wss/v20180426"
)

const (
	PROVIDER_CVM_REQUEST_TIMEOUT = "TENCENTCLOUD_CVM_REQUEST_TIMEOUT"
	PROVIDER_CBS_REQUEST_TIMEOUT = "TENCENTCLOUD_CBS_REQUEST_TIMEOUT"
)

// TencentCloudClient is client for all TencentCloud service
type TencentCloudClient struct {
	Credential          *common.Credential
	CredentialTce       *common2.Credential
	Region              string
	Protocol            string
	Domain              string
	CspDomain           string
	CosDomain           string
	cosConn             *s3.S3
	tencentCosConn      *cos.Client
	mysqlConn           *cdb.Client
	redisConn           *redis.Client
	asConn              *as.Client
	bmsConn             *bms.Client
	bhsaasConn          *bhsaas.Client
	vpcConn             *vpc.Client
	vpcdnsConn          *vpcdns.Client
	cbsConn             *cbs.Client
	cvmConn             *cvm.Client
	clbConn             *clb.Client
	cspConn             *csp.Client
	drcConn             *drc.Client
	dayuConn            *dayu.Client
	dcConn              *dc.Client
	tagConn             *tag.Client
	mongodbConn         *mongodb.Client
	tkeConn             *tke.Client
	tke2Conn            *tke2.Client
	tdmqConn            *tdmq.Client
	tcrConn             *tcr.Client
	camConn             *cam.Client
	accountConn         *account.Client
	account20181225Conn *account20181225.Client
	openConn            *open.Client
	stsConn             *sts.Client
	gaapConn            *gaap.Client
	sslConn             *ssl.Client
	cfsConn             *cfs.Client
	turbofsConn         *turbofs.Client
	tcaplusConn         *tcaplusdb.Client
	cdnConn             *cdn.Client
	monitorConn         *monitor.Client
	esConn              *es.Client
	sqlserverConn       *sqlserver.Client
	postgreConn         *postgre.Client
	ckafkaConn          *ckafka.Client
	auditConn           *audit.Client
	cynosConn           *cynosdb.Client
	apiGatewayConn      *apigateway.Client
	sslCertificateConn  *sslCertificate.Client
	kmsConn             *kms.Client
	ssmConn             *ssm.Client
	apiConn             *api.Client
	emrConn             *emr.Client
	clsConn             *cls.Client
	dnsPodConn          *dnspod.Client
	privateDnsConn      *privatedns.Client
	antiddosConn        *antiddos.Client
	domainConn          *domain.Client
	lighthouseConn      *lighthouse.Client
	teoConn             *teo.Client
	tcmConn             *tcm.Client
	cssConn             *css.Client
	sesConn             *ses.Client
	dcdbConn            *dcdb.Client
	smsConn             *sms.Client
	catConn             *cat.Client
	mariadbConn         *mariadb.Client
	rumConn             *rum.Client
	ptsConn             *pts.Client
	tatConn             *tat.Client
	tbaseConn           *tbase.Client
	organizationConn    *organization.Client
	tdcpgConn           *tdcpg.Client
	dbbrainConn         *dbbrain.Client
	dtsConn             *dts.Client
	ciConn              *cos.Client
	tsfConn             *tsf.Client
	mpsConn             *mps.Client
	cwpConn             *cwp.Client
	chdfsConn           *chdfs.Client
	mdlConn             *mdl.Client
	apmConn             *apm.Client
	cfwConn             *cfw.Client
	brcConn             *brc.Client
	cdcConn             *cdc.Client
	cicConn             *cic.Client
	locationConn        *location.Client
	ngwafConn           *ngwaf.Client
	ccnConn             *ccn.Client
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
	cpf.HttpProfile.RootDomain = me.Domain
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

// NewClientProfile returns a new ClientProfile
func (me *TencentCloudClient) NewClientIntlProfile(timeout int) *intlProfile.ClientProfile {
	cpf := intlProfile.NewClientProfile()

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

// UseBmsClient returns bms client for service
func (me *TencentCloudClient) UseBmsClient() *bms.Client {
	if me.bmsConn != nil {
		return me.bmsConn
	}
	cpf := me.NewClientProfileTce(300)
	me.bmsConn, _ = bms.NewClient(me.CredentialTce, me.Region, cpf)
	me.bmsConn.WithHttpTransport(&LogRoundTripper{})
	return me.bmsConn
}

// UseBhsaasClient returns bhsaas client for service
func (me *TencentCloudClient) UseBhsaasClient() *bhsaas.Client {
	if me.bhsaasConn != nil {
		return me.bhsaasConn
	}
	cpf := me.NewClientProfileTce(300)
	me.bhsaasConn, _ = bhsaas.NewClient(me.CredentialTce, me.Region, cpf)
	me.bhsaasConn.WithHttpTransport(&LogRoundTripper{})
	return me.bhsaasConn
}

// UseCcnClient returns ccn client for service
func (me *TencentCloudClient) UseCcnClient() *ccn.Client {
	if me.ccnConn != nil {
		return me.ccnConn
	}
	cpf := me.NewClientProfileTce(300)
	me.ccnConn, _ = ccn.NewClient(me.CredentialTce, me.Region, cpf)
	me.ccnConn.WithHttpTransport(&LogRoundTripper{})
	return me.ccnConn
}

// UseCosClient returns cos client for service
func (me *TencentCloudClient) UseCosClient() *s3.S3 {
	if me.cosConn != nil {
		return me.cosConn
	}
	resolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				URL:           fmt.Sprintf("%s://cos.%s.%s", me.Protocol, me.Region, me.CosDomain),
				SigningRegion: region,
			}, nil
		}
		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	creds := credentials.NewStaticCredentials(me.Credential.SecretId, me.Credential.SecretKey, me.Credential.Token)

	awsCfg := &aws.Config{
		Credentials:      creds,
		Region:           aws.String(me.Region),
		EndpointResolver: endpoints.ResolverFunc(resolver),
		// HTTPClient:       httpClient,
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(awsCfg))

	return s3.New(sess)
}

// UseCosS3Client returns cos client for service
// return csp or cos s3 client
func (me *TencentCloudClient) UseCosS3Client(useCsp bool) *s3.S3 {
	if me.cosConn != nil {
		return me.cosConn
	}
	var s3Url string
	if useCsp {
		s3Url = fmt.Sprintf("%s://cos.%s.%s", me.Protocol, me.Region, me.CspDomain)
	} else {
		s3Url = fmt.Sprintf("%s://cos.%s.%s", me.Protocol, me.Region, me.CosDomain)
	}
	resolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == endpoints.S3ServiceID {
			return endpoints.ResolvedEndpoint{
				//URL:           fmt.Sprintf("https://cos.%s.myqcloud.com", region),
				URL:           s3Url,
				SigningRegion: region,
			}, nil
		}
		return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
	}

	creds := credentials.NewStaticCredentials(me.Credential.SecretId, me.Credential.SecretKey, me.Credential.Token)

	awsCfg := &aws.Config{
		Credentials:      creds,
		Region:           aws.String(me.Region),
		EndpointResolver: endpoints.ResolverFunc(resolver),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.Must(session.NewSession(awsCfg))

	return s3.New(sess)
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

// UseTencentCosClient tencent cloud own client for service instead of aws
func (me *TencentCloudClient) UseTencentCosClient(bucket string) *cos.Client {
	svc, _ := url.Parse(fmt.Sprintf("%s://cos.%s.%s", me.Protocol, me.Region, me.CosDomain))
	bkt, _ := url.Parse(fmt.Sprintf("%s://%s.cos.%s.%s", me.Protocol, bucket, me.Region, me.CosDomain))
	return me.GetTencentCosClient(bkt, svc, bucket)
}

// UseTencentCspClient tencent cloud own client for service instead of aws
func (me *TencentCloudClient) UseTencentCspClient(bucket string) *cos.Client {
	svc, _ := url.Parse(fmt.Sprintf("%s://cos.%s.%s", me.Protocol, me.Region, me.CspDomain))
	bkt, _ := url.Parse(fmt.Sprintf("%s://%s.cos.%s.%s", me.Protocol, bucket, me.Region, me.CspDomain))
	return me.GetTencentCosClient(bkt, svc, bucket)
}

// GetTencentCosClient tencent cloud own client for service instead of aws
func (me *TencentCloudClient) GetTencentCosClient(bucketUrl, serviceUrl *url.URL, bucket string) *cos.Client {
	if me.tencentCosConn != nil && me.tencentCosConn.BaseURL.BucketURL == bucketUrl {
		return me.tencentCosConn
	}

	baseUrl := &cos.BaseURL{
		BucketURL:  bucketUrl,
		ServiceURL: serviceUrl,
	}

	me.tencentCosConn = cos.NewClient(baseUrl, &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:     me.Credential.SecretId,
			SecretKey:    me.Credential.SecretKey,
			SessionToken: me.Credential.Token,
		},
	})

	return me.tencentCosConn
}

// UseMysqlClient returns mysql(cdb) client for service
func (me *TencentCloudClient) UseMysqlClient() *cdb.Client {
	if me.mysqlConn != nil {
		return me.mysqlConn
	}

	cpf := me.NewClientProfile(300)
	me.mysqlConn, _ = cdb.NewClient(me.Credential, me.Region, cpf)
	me.mysqlConn.WithHttpTransport(&LogRoundTripper{})

	return me.mysqlConn
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

// UseAsClient returns as client for service
func (me *TencentCloudClient) UseAsClient() *as.Client {
	if me.asConn != nil {
		return me.asConn
	}

	cpf := me.NewClientProfileTce(300)
	me.asConn, _ = as.NewClient(me.CredentialTce, me.Region, cpf)
	me.asConn.WithHttpTransport(&LogRoundTripper{})

	return me.asConn
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

// UseVpcDnsClient returns vpcdns client for service
func (me *TencentCloudClient) UseVpcDnsClient() *vpcdns.Client {
	if me.vpcdnsConn != nil {
		return me.vpcdnsConn
	}

	cpf := me.NewClientProfileTce(300)
	me.vpcdnsConn, _ = vpcdns.NewClient(me.CredentialTce, me.Region, cpf)
	me.vpcdnsConn.WithHttpTransport(&LogRoundTripper{})

	return me.vpcdnsConn
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

// UseDcClient returns dc client for service
func (me *TencentCloudClient) UseDcClient() *dc.Client {
	if me.dcConn != nil {
		return me.dcConn
	}

	cpf := me.NewClientProfileTce(300)
	me.dcConn, _ = dc.NewClient(me.CredentialTce, me.Region, cpf)
	me.dcConn.WithHttpTransport(&LogRoundTripper{})

	return me.dcConn
}

// UseMongodbClient returns mongodb client for service
func (me *TencentCloudClient) UseMongodbClient() *mongodb.Client {
	if me.mongodbConn != nil {
		return me.mongodbConn
	}

	cpf := me.NewClientProfile(300)
	me.mongodbConn, _ = mongodb.NewClient(me.Credential, me.Region, cpf)
	me.mongodbConn.WithHttpTransport(&LogRoundTripper{})

	return me.mongodbConn
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

// UseTke2Client returns tke v20220501 client for service (native node pool)
func (me *TencentCloudClient) UseTke2Client() *tke2.Client {
	if me.tke2Conn != nil {
		return me.tke2Conn
	}

	cpf := me.NewClientProfileTce(300)
	me.tke2Conn, _ = tke2.NewClient(me.CredentialTce, me.Region, cpf)
	me.tke2Conn.WithHttpTransport(&LogRoundTripper{})

	return me.tke2Conn
}

// UseTdmqClient returns Tdmq client for service
func (me *TencentCloudClient) UseTdmqClient(iacExtInfo ...IacExtInfo) *tdmq.Client {
	var logRoundTripper LogRoundTripper
	if len(iacExtInfo) != 0 {
		logRoundTripper.InstanceId = iacExtInfo[0].InstanceId
	}

	if me.tdmqConn != nil {
		me.tdmqConn.WithHttpTransport(&logRoundTripper)
		return me.tdmqConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tdmqConn, _ = tdmq.NewClient(me.CredentialTce, me.Region, cpf)
	me.tdmqConn.WithHttpTransport(&logRoundTripper)

	return me.tdmqConn
}

// UseGaapClient returns gaap client for service
func (me *TencentCloudClient) UseGaapClient() *gaap.Client {
	if me.gaapConn != nil {
		return me.gaapConn
	}

	cpf := me.NewClientProfile(300)
	me.gaapConn, _ = gaap.NewClient(me.Credential, me.Region, cpf)
	me.gaapConn.WithHttpTransport(&LogRoundTripper{})

	return me.gaapConn
}

// UseSslClient returns ssl client for service
func (me *TencentCloudClient) UseSslClient() *ssl.Client {
	if me.sslConn != nil {
		return me.sslConn
	}

	cpf := me.NewClientProfile(300)
	me.sslConn, _ = ssl.NewClient(me.Credential, me.Region, cpf)
	me.sslConn.WithHttpTransport(&LogRoundTripper{})

	return me.sslConn
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

// UseTurbofsClient returns turbofs client for service
func (me *TencentCloudClient) UseTurbofsClient() *turbofs.Client {
	if me.turbofsConn != nil {
		return me.turbofsConn
	}

	cpf := me.NewClientProfileTce(300)
	me.turbofsConn, _ = turbofs.NewClient(me.CredentialTce, me.Region, cpf)
	me.turbofsConn.WithHttpTransport(&LogRoundTripper{})

	return me.turbofsConn
}

// UseTcaplusClient returns tcaplush client for service
func (me *TencentCloudClient) UseTcaplusClient() *tcaplusdb.Client {
	if me.tcaplusConn != nil {
		return me.tcaplusConn
	}

	cpf := me.NewClientProfile(300)
	me.tcaplusConn, _ = tcaplusdb.NewClient(me.Credential, me.Region, cpf)
	me.tcaplusConn.WithHttpTransport(&LogRoundTripper{})

	return me.tcaplusConn
}

// UseDayuClient returns dayu client for service
func (me *TencentCloudClient) UseDayuClient() *dayu.Client {
	if me.dayuConn != nil {
		return me.dayuConn
	}

	cpf := me.NewClientProfile(300)
	me.dayuConn, _ = dayu.NewClient(me.Credential, me.Region, cpf)
	me.dayuConn.WithHttpTransport(&LogRoundTripper{})

	return me.dayuConn
}

// UseCdnClient returns cdn client for service
func (me *TencentCloudClient) UseCdnClient() *cdn.Client {
	if me.cdnConn != nil {
		return me.cdnConn
	}

	cpf := me.NewClientProfile(300)
	me.cdnConn, _ = cdn.NewClient(me.Credential, me.Region, cpf)
	me.cdnConn.WithHttpTransport(&LogRoundTripper{})

	return me.cdnConn
}

// UseMonitorClient returns monitor client for service
func (me *TencentCloudClient) UseMonitorClient() *monitor.Client {
	if me.monitorConn != nil {
		return me.monitorConn
	}

	cpf := me.NewClientProfile(300)
	me.monitorConn, _ = monitor.NewClient(me.Credential, me.Region, cpf)
	me.monitorConn.WithHttpTransport(&LogRoundTripper{})

	return me.monitorConn
}

// UseEsClient returns es client for service
func (me *TencentCloudClient) UseEsClient() *es.Client {
	if me.esConn != nil {
		return me.esConn
	}

	cpf := me.NewClientProfile(300)
	cpf.Language = "zh-CN"
	me.esConn, _ = es.NewClient(me.Credential, me.Region, cpf)
	me.esConn.WithHttpTransport(&LogRoundTripper{})

	return me.esConn
}

// UsePostgresqlClient returns postgresql client for service
func (me *TencentCloudClient) UsePostgresqlClient() *postgre.Client {
	if me.postgreConn != nil {
		return me.postgreConn
	}

	cpf := me.NewClientProfile(300)
	me.postgreConn, _ = postgre.NewClient(me.Credential, me.Region, cpf)
	me.postgreConn.WithHttpTransport(&LogRoundTripper{})

	return me.postgreConn
}

// UseSqlserverClient returns sqlserver client for service
func (me *TencentCloudClient) UseSqlserverClient() *sqlserver.Client {
	if me.sqlserverConn != nil {
		return me.sqlserverConn
	}

	cpf := me.NewClientProfile(300)
	me.sqlserverConn, _ = sqlserver.NewClient(me.Credential, me.Region, cpf)
	me.sqlserverConn.WithHttpTransport(&LogRoundTripper{})

	return me.sqlserverConn
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

// UseAuditClient returns audit client for service
func (me *TencentCloudClient) UseAuditClient() *audit.Client {
	if me.auditConn != nil {
		return me.auditConn
	}

	cpf := me.NewClientProfile(300)
	me.auditConn, _ = audit.NewClient(me.Credential, me.Region, cpf)
	me.auditConn.WithHttpTransport(&LogRoundTripper{})

	return me.auditConn
}

// UseCynosdbClient returns cynosdb client for service
func (me *TencentCloudClient) UseCynosdbClient() *cynosdb.Client {
	if me.cynosConn != nil {
		return me.cynosConn
	}

	cpf := me.NewClientProfile(300)
	me.cynosConn, _ = cynosdb.NewClient(me.Credential, me.Region, cpf)
	me.cynosConn.WithHttpTransport(&LogRoundTripper{})

	return me.cynosConn
}

// UseAPIGatewayClient returns apigateway client for service
func (me *TencentCloudClient) UseAPIGatewayClient() *apigateway.Client {
	if me.apiGatewayConn != nil {
		return me.apiGatewayConn
	}

	cpf := me.NewClientProfile(300)
	me.apiGatewayConn, _ = apigateway.NewClient(me.Credential, me.Region, cpf)
	me.apiGatewayConn.WithHttpTransport(&LogRoundTripper{})

	return me.apiGatewayConn
}

// UseTCRClient returns apigateway client for service
func (me *TencentCloudClient) UseTCRClient() *tcr.Client {
	if me.tcrConn != nil {
		return me.tcrConn
	}

	cpf := me.NewClientProfileTce(300)
	me.tcrConn, _ = tcr.NewClient(me.CredentialTce, me.Region, cpf)
	me.tcrConn.WithHttpTransport(&LogRoundTripper{})

	return me.tcrConn
}

// UseSSLCertificateClient returns SSL Certificate client for service
func (me *TencentCloudClient) UseSSLCertificateClient() *sslCertificate.Client {
	if me.sslCertificateConn != nil {
		return me.sslCertificateConn
	}

	cpf := me.NewClientProfile(300)
	me.sslCertificateConn, _ = sslCertificate.NewClient(me.Credential, me.Region, cpf)
	me.sslCertificateConn.WithHttpTransport(&LogRoundTripper{})

	return me.sslCertificateConn
}

// UseKmsClient returns KMS client for service
func (me *TencentCloudClient) UseKmsClient() *kms.Client {
	if me.kmsConn != nil {
		return me.kmsConn
	}

	cpf := me.NewClientProfile(300)
	me.kmsConn, _ = kms.NewClient(me.Credential, me.Region, cpf)
	me.kmsConn.WithHttpTransport(&LogRoundTripper{})

	return me.kmsConn
}

// UseSsmClient returns SSM client for service
func (me *TencentCloudClient) UseSsmClient() *ssm.Client {
	if me.ssmConn != nil {
		return me.ssmConn
	}

	cpf := me.NewClientProfile(300)
	me.ssmConn, _ = ssm.NewClient(me.Credential, me.Region, cpf)
	me.ssmConn.WithHttpTransport(&LogRoundTripper{})

	return me.ssmConn
}

// UseApiClient return API client for service
func (me *TencentCloudClient) UseApiClient() *api.Client {
	if me.apiConn != nil {
		return me.apiConn
	}
	cpf := me.NewClientProfile(300)
	me.apiConn, _ = api.NewClient(me.Credential, me.Region, cpf)
	me.apiConn.WithHttpTransport(&LogRoundTripper{})

	return me.apiConn
}

// UseEmrClient return EMR client for service
func (me *TencentCloudClient) UseEmrClient() *emr.Client {
	if me.emrConn != nil {
		return me.emrConn
	}
	cpf := me.NewClientProfile(300)
	me.emrConn, _ = emr.NewClient(me.Credential, me.Region, cpf)
	me.emrConn.WithHttpTransport(&LogRoundTripper{})

	return me.emrConn
}

// UseClsClient return CLS client for service
func (me *TencentCloudClient) UseClsClient() *cls.Client {
	if me.clsConn != nil {
		return me.clsConn
	}
	cpf := me.NewClientProfileTce(300)
	me.clsConn, _ = cls.NewClient(me.CredentialTce, me.Region, cpf)
	me.clsConn.WithHttpTransport(&LogRoundTripper{})

	return me.clsConn
}

// UseLighthouseClient return Lighthouse client for service
func (me *TencentCloudClient) UseLighthouseClient() *lighthouse.Client {
	if me.lighthouseConn != nil {
		return me.lighthouseConn
	}
	cpf := me.NewClientProfile(300)
	me.lighthouseConn, _ = lighthouse.NewClient(me.Credential, me.Region, cpf)
	me.lighthouseConn.WithHttpTransport(&LogRoundTripper{})

	return me.lighthouseConn
}

// UseDnsPodClient return DnsPod client for service
func (me *TencentCloudClient) UseDnsPodClient() *dnspod.Client {
	if me.dnsPodConn != nil {
		return me.dnsPodConn
	}
	cpf := me.NewClientProfile(300)
	me.dnsPodConn, _ = dnspod.NewClient(me.Credential, me.Region, cpf)
	me.dnsPodConn.WithHttpTransport(&LogRoundTripper{})

	return me.dnsPodConn
}

// UsePrivateDnsClient return PrivateDns client for service
func (me *TencentCloudClient) UsePrivateDnsClient() *privatedns.Client {
	if me.dnsPodConn != nil {
		return me.privateDnsConn
	}
	cpf := me.NewClientProfile(300)
	me.privateDnsConn, _ = privatedns.NewClient(me.Credential, me.Region, cpf)
	me.privateDnsConn.WithHttpTransport(&LogRoundTripper{})

	return me.privateDnsConn
}

// UseDomainClient return Domain client for service
func (me *TencentCloudClient) UseDomainClient() *domain.Client {
	if me.dnsPodConn != nil {
		return me.domainConn
	}
	cpf := me.NewClientProfile(300)
	me.domainConn, _ = domain.NewClient(me.Credential, me.Region, cpf)
	me.domainConn.WithHttpTransport(&LogRoundTripper{})

	return me.domainConn
}

// UseAntiddosClient returns antiddos client for service
func (me *TencentCloudClient) UseAntiddosClient() *antiddos.Client {
	if me.antiddosConn != nil {
		return me.antiddosConn
	}

	cpf := me.NewClientProfile(300)
	me.antiddosConn, _ = antiddos.NewClient(me.Credential, me.Region, cpf)
	me.antiddosConn.WithHttpTransport(&LogRoundTripper{})

	return me.antiddosConn
}

// UseTeoClient returns teo client for service
func (me *TencentCloudClient) UseTeoClient() *teo.Client {
	if me.teoConn != nil {
		return me.teoConn
	}

	cpf := me.NewClientProfile(300)
	me.teoConn, _ = teo.NewClient(me.Credential, me.Region, cpf)
	me.teoConn.WithHttpTransport(&LogRoundTripper{})

	return me.teoConn
}

// UseTcmClient returns Tcm client for service
func (me *TencentCloudClient) UseTcmClient() *tcm.Client {
	if me.tcmConn != nil {
		return me.tcmConn
	}

	cpf := me.NewClientProfile(300)
	me.tcmConn, _ = tcm.NewClient(me.Credential, me.Region, cpf)
	me.tcmConn.WithHttpTransport(&LogRoundTripper{})

	return me.tcmConn
}

// UseCssClient returns css client for service
func (me *TencentCloudClient) UseCssClient() *css.Client {
	if me.cssConn != nil {
		return me.cssConn
	}

	cpf := me.NewClientProfile(300)
	me.cssConn, _ = css.NewClient(me.Credential, me.Region, cpf)
	me.cssConn.WithHttpTransport(&LogRoundTripper{})

	return me.cssConn
}

// UseSesClient returns Ses client for service
func (me *TencentCloudClient) UseSesClient() *ses.Client {
	if me.sesConn != nil {
		return me.sesConn
	}

	cpf := me.NewClientProfile(300)
	me.sesConn, _ = ses.NewClient(me.Credential, me.Region, cpf)
	me.sesConn.WithHttpTransport(&LogRoundTripper{})

	return me.sesConn
}

// UseDcdbClient returns dcdb client for service
func (me *TencentCloudClient) UseDcdbClient() *dcdb.Client {
	if me.dcdbConn != nil {
		return me.dcdbConn
	}

	cpf := me.NewClientProfileTce(300)
	me.dcdbConn, _ = dcdb.NewClient(me.CredentialTce, me.Region, cpf)
	me.dcdbConn.WithHttpTransport(&LogRoundTripper{})

	return me.dcdbConn
}

// UseSmsClient returns Sms client for service
func (me *TencentCloudClient) UseSmsClient() *sms.Client {
	if me.smsConn != nil {
		return me.smsConn
	}

	cpf := me.NewClientProfile(300)
	me.smsConn, _ = sms.NewClient(me.Credential, me.Region, cpf)
	me.smsConn.WithHttpTransport(&LogRoundTripper{})

	return me.smsConn
}

// UseCatClient returns Cat client for service
func (me *TencentCloudClient) UseCatClient() *cat.Client {
	if me.catConn != nil {
		return me.catConn
	}

	cpf := me.NewClientProfile(300)
	me.catConn, _ = cat.NewClient(me.Credential, me.Region, cpf)
	me.catConn.WithHttpTransport(&LogRoundTripper{})

	return me.catConn
}

// UseMariadbClient returns mariadb client for service
func (me *TencentCloudClient) UseMariadbClient() *mariadb.Client {
	if me.mariadbConn != nil {
		return me.mariadbConn
	}

	cpf := me.NewClientProfile(300)
	me.mariadbConn, _ = mariadb.NewClient(me.Credential, me.Region, cpf)
	me.mariadbConn.WithHttpTransport(&LogRoundTripper{})

	return me.mariadbConn
}

// UsePtsClient returns pts client for service
func (me *TencentCloudClient) UsePtsClient() *pts.Client {
	if me.ptsConn != nil {
		return me.ptsConn
	}

	cpf := me.NewClientProfile(300)
	me.ptsConn, _ = pts.NewClient(me.Credential, me.Region, cpf)
	me.ptsConn.WithHttpTransport(&LogRoundTripper{})

	return me.ptsConn
}

// UseTatClient returns tat client for service
func (me *TencentCloudClient) UseTatClient() *tat.Client {
	if me.tatConn != nil {
		return me.tatConn
	}

	cpf := me.NewClientProfile(300)
	me.tatConn, _ = tat.NewClient(me.Credential, me.Region, cpf)
	me.tatConn.WithHttpTransport(&LogRoundTripper{})

	return me.tatConn
}

// UseOrganizationClient returns organization client for service
func (me *TencentCloudClient) UseOrganizationClient() *organization.Client {
	if me.organizationConn != nil {
		return me.organizationConn
	}

	cpf := me.NewClientProfileTce(300)
	me.organizationConn, _ = organization.NewClient(me.CredentialTce, me.Region, cpf)
	me.organizationConn.WithHttpTransport(&LogRoundTripper{})

	return me.organizationConn
}

// UseTdcpgClient returns tdcpg client for service
func (me *TencentCloudClient) UseTdcpgClient() *tdcpg.Client {
	if me.tdcpgConn != nil {
		return me.tdcpgConn
	}

	cpf := me.NewClientProfile(300)
	me.tdcpgConn, _ = tdcpg.NewClient(me.Credential, me.Region, cpf)
	me.tdcpgConn.WithHttpTransport(&LogRoundTripper{})

	return me.tdcpgConn
}

// UseDbbrainClient returns dbbrain client for service
func (me *TencentCloudClient) UseDbbrainClient() *dbbrain.Client {
	if me.dbbrainConn != nil {
		return me.dbbrainConn
	}

	cpf := me.NewClientProfile(300)
	cpf.Language = "zh-CN"
	me.dbbrainConn, _ = dbbrain.NewClient(me.Credential, me.Region, cpf)
	me.dbbrainConn.WithHttpTransport(&LogRoundTripper{})

	return me.dbbrainConn
}

// UseRumClient returns rum client for service
func (me *TencentCloudClient) UseRumClient() *rum.Client {
	if me.rumConn != nil {
		return me.rumConn
	}

	cpf := me.NewClientProfile(300)
	me.rumConn, _ = rum.NewClient(me.Credential, me.Region, cpf)
	me.rumConn.WithHttpTransport(&LogRoundTripper{})

	return me.rumConn
}

// UseDtsClient returns dts client for service
func (me *TencentCloudClient) UseDtsClient() *dts.Client {
	if me.dtsConn != nil {
		return me.dtsConn
	}

	cpf := me.NewClientProfile(300)
	me.dtsConn, _ = dts.NewClient(me.Credential, me.Region, cpf)
	me.dtsConn.WithHttpTransport(&LogRoundTripper{})

	return me.dtsConn
}

// UseTsfClient returns tsf client for service
func (me *TencentCloudClient) UseTsfClient() *tsf.Client {
	if me.tsfConn != nil {
		return me.tsfConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.tsfConn, _ = tsf.NewClient(me.CredentialTce, me.Region, cpf)
	me.tsfConn.WithHttpTransport(&LogRoundTripper{})

	return me.tsfConn
}

// UseMpsClient returns mps client for service
func (me *TencentCloudClient) UseMpsClient() *mps.Client {
	if me.mpsConn != nil {
		return me.mpsConn
	}

	cpf := me.NewClientProfile(300)
	cpf.Language = "zh-CN"
	me.mpsConn, _ = mps.NewClient(me.Credential, me.Region, cpf)
	me.mpsConn.WithHttpTransport(&LogRoundTripper{})

	return me.mpsConn
}

// UseTkeClient returns tke client for service
func (me *TencentCloudClient) UseCwpClient() *cwp.Client {
	if me.cwpConn != nil {
		return me.cwpConn
	}

	var reqTimeout = getEnvDefault(PROVIDER_CVM_REQUEST_TIMEOUT, 300)
	cpf := me.NewClientProfileTce(reqTimeout)
	me.cwpConn, _ = cwp.NewClient(me.CredentialTce, me.Region, cpf)
	me.cwpConn.WithHttpTransport(&LogRoundTripper{})

	return me.cwpConn
}

// UseChdfsClient returns chdfs client for service
func (me *TencentCloudClient) UseChdfsClient() *chdfs.Client {
	if me.chdfsConn != nil {
		return me.chdfsConn
	}

	cpf := me.NewClientProfile(300)
	cpf.Language = "zh-CN"
	me.chdfsConn, _ = chdfs.NewClient(me.Credential, me.Region, cpf)
	me.chdfsConn.WithHttpTransport(&LogRoundTripper{})

	return me.chdfsConn
}

// UseMdlClient returns mdl client for service
func (me *TencentCloudClient) UseMdlClient() *mdl.Client {
	if me.mdlConn != nil {
		return me.mdlConn
	}

	cpf := me.NewClientIntlProfile(300)
	cpf.Language = "zh-CN"
	me.mdlConn, _ = mdl.NewClient(me.Credential, me.Region, cpf)
	me.mdlConn.WithHttpTransport(&LogRoundTripper{})

	return me.mdlConn
}

// UseApmClient returns apm client for service
func (me *TencentCloudClient) UseApmClient() *apm.Client {
	if me.apmConn != nil {
		return me.apmConn
	}

	cpf := me.NewClientProfileTce(300)
	me.apmConn, _ = apm.NewClient(me.CredentialTce, me.Region, cpf)
	me.apmConn.WithHttpTransport(&LogRoundTripper{})

	return me.apmConn
}

// UseCspClient returns csp client for service
func (me *TencentCloudClient) UseCspClient() *csp.Client {
	if me.cspConn != nil {
		return me.cspConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.cspConn, _ = csp.NewClient(me.CredentialTce, me.Region, cpf)
	me.cspConn.WithHttpTransport(&LogRoundTripper{})

	return me.cspConn
}

// UseCicClient returns cic client for service
func (me *TencentCloudClient) UseCicClient() *cic.Client {
	if me.cicConn != nil {
		return me.cicConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.cicConn, _ = cic.NewClient(me.CredentialTce, me.Region, cpf)
	me.cicConn.WithHttpTransport(&LogRoundTripper{})

	return me.cicConn
}

func (me *TencentCloudClient) UseDrcClient() *drc.Client {
	if me.drcConn != nil {
		return me.drcConn
	}
	cpf := me.NewClientProfileTce(300)
	me.drcConn, _ = drc.NewClient(me.CredentialTce, me.Region, cpf)
	me.drcConn.WithHttpTransport(&LogRoundTripper{})
	return me.drcConn
}

// UseCfwClient returns cfw client for service
func (me *TencentCloudClient) UseCfwClient() *cfw.Client {
	if me.cfwConn != nil {
		return me.cfwConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.cfwConn, _ = cfw.NewClient(me.CredentialTce, me.Region, cpf)
	me.cfwConn.WithHttpTransport(&LogRoundTripper{})

	return me.cfwConn
}

// UseBrcClient returns brc client for service
func (me *TencentCloudClient) UseBrcClient() *brc.Client {
	if me.brcConn != nil {
		return me.brcConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.brcConn, _ = brc.NewClient(me.CredentialTce, me.Region, cpf)
	me.brcConn.WithHttpTransport(&LogRoundTripper{})

	return me.brcConn
}

// UseNgwafClient returns ngwaf client for service
func (me *TencentCloudClient) UseNgwafClient() *ngwaf.Client {
	if me.ngwafConn != nil {
		return me.ngwafConn
	}

	cpf := me.NewClientProfileTce(300)
	cpf.Language = "zh-CN"
	me.ngwafConn, _ = ngwaf.NewClient(me.CredentialTce, me.Region, cpf)
	me.ngwafConn.WithHttpTransport(&LogRoundTripper{})

	return me.ngwafConn
}

// UseLocationClient returns location client for service
func (me *TencentCloudClient) UseLocationClient() *location.Client {
	if me.locationConn != nil {
		return me.locationConn
	}

	cpf := me.NewClientProfileTce(300)
	me.locationConn, _ = location.NewClient(me.CredentialTce, me.Region, cpf)
	me.locationConn.WithHttpTransport(&LogRoundTripper{})

	return me.locationConn
}

// UseCdcClient returns cdc client for service
func (me *TencentCloudClient) UseCdcClient() *cdc.Client {
	if me.cdcConn != nil {
		return me.cdcConn
	}

	cpf := me.NewClientProfileTce(300)
	me.cdcConn, _ = cdc.NewClient(me.CredentialTce, me.Region, cpf)
	me.cdcConn.WithHttpTransport(&LogRoundTripper{})

	return me.cdcConn
}

// UseCdcV20201214Client returns cdc client for service (alias for compatibility)
func (me *TencentCloudClient) UseCdcV20201214Client() *cdc.Client {
	return me.UseCdcClient()
}

// UseAccountClient returns account client for service
func (me *TencentCloudClient) UseAccountClient() *account.Client {
	if me.accountConn != nil {
		return me.accountConn
	}

	cpf := me.NewClientProfileTce(300)
	me.accountConn, _ = account.NewClient(me.CredentialTce, me.Region, cpf)
	me.accountConn.WithHttpTransport(&LogRoundTripper{})

	return me.accountConn
}

// UseAccount20181225Client returns account client for v20181225
func (me *TencentCloudClient) UseAccount20181225Client() *account20181225.Client {
	if me.account20181225Conn != nil {
		return me.account20181225Conn
	}

	cpf := me.NewClientProfileTce(300)
	me.account20181225Conn, _ = account20181225.NewClient(me.CredentialTce, me.Region, cpf)
	me.account20181225Conn.WithHttpTransport(&LogRoundTripper{})

	return me.account20181225Conn
}

func (me *TencentCloudClient) UseOpenClient() *open.Client {
	if me.openConn != nil {
		return me.openConn
	}

	cpf := me.NewClientProfileTce(300)
	me.openConn, _ = open.NewClient(me.CredentialTce, me.Region, cpf)
	me.openConn.WithHttpTransport(&LogRoundTripper{})

	return me.openConn
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
