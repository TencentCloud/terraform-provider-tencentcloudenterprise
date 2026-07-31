package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	cwp "terraform-provider-tencentcloudenterprise/sdk/cwp/v20180228"
	tat "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat/v20201028"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type ClusterBasicSetting struct {
	ClusterId               string
	ClusterOs               string
	ClusterOsType           string
	ClusterVersion          string
	ClusterName             string
	ClusterDescription      string
	ClusterLevel            *string
	AutoUpgradeClusterLevel *bool
	VpcId                   string
	SubnetId                string
	CdcId                   string
	ProjectId               int64
	NeedWorkSecurityGroup   bool
	ClusterNodeNum          int64
	ClusterStatus           string
	Tags                    map[string]string
}

type ClusterAdvancedSettings struct {
	Ipvs                    bool
	AsEnabled               bool
	EnableCustomizedPodCIDR bool
	BasePodNumber           int64
	ContainerRuntime        string
	RuntimeVersion          string
	NodeNameType            string
	ExtraArgs               ClusterExtraArgs
	NetworkType             string
	VpcCniType              string
	IsNonStaticIpMode       bool
	DeletionProtection      bool
	KubeProxyMode           string
	AuditEnabled            bool
	AuditLogsetId           string
	AuditLogTopicId         string
	DataPlaneV2             bool
	QGPUShareEnable         bool
	IsDualStack             bool
}

type ClusterExtraArgs struct {
	KubeAPIServer         []string
	KubeControllerManager []string
	KubeScheduler         []string
}

type RunInstancesForNode struct {
	Master []string
	Work   []string
}

type InstanceAdvancedSettings struct {
	MountTarget        string
	DockerGraphPath    string
	UserScript         string
	PreStartUserScript string
	Unschedulable      int64
	DesiredPodNum      int64
	DesiredPodNumSet   bool
	DataDiskPartition  string
	GPUArgs            *tke.GPUArgs
	Taints             []*tke.Taint
	Labels             []*tke.Label
	DataDisks          []*tke.DataDisk
	ExtraArgs          tke.InstanceExtraArgs
}

type ClusterCidrSettings struct {
	ClusterCidr               string
	IgnoreClusterCidrConflict bool
	MaxNodePodNum             int64
	MaxClusterServiceNum      int64
	ServiceCIDR               string
	IgnoreServiceCIDRConflict bool
	EniSubnetIds              []string
	ClaimExpiredSeconds       int64
}

type ClusterInfo struct {
	ClusterBasicSetting
	ClusterCidrSettings
	ClusterAdvancedSettings

	DeployType  string
	CreatedTime string
}

type InstanceInfo struct {
	InstanceId                   string
	InstanceRole                 string
	InstanceState                string
	FailedReason                 string
	NodePoolId                   string
	CreatedTime                  string
	InstanceAdvancedSettings     *tke.InstanceAdvancedSettings
	InstanceDataDiskMountSetting *tke.InstanceDataDiskMountSetting
	LanIp                        string
}

type PrometheusConfigIds struct {
	InstanceId  string
	ClusterType string
	ClusterId   string
}

type TkeService struct {
	client *connectivity.TencentCloudClient
}

func (me *TkeService) DescribeClusterInstances(ctx context.Context, id string) (masters []InstanceInfo, workers []InstanceInfo, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeClusterInstancesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &id
	allRole := "ALL"
	request.InstanceRole = &allRole
	masters = make([]InstanceInfo, 0, 100)
	workers = make([]InstanceInfo, 0, 100)
	var offset int64 = 0
	var limit int64 = 20
	var has = map[string]bool{}
	var total int64 = -1

getMoreData:
	if total >= 0 && offset >= total {
		return
	}
	request.Limit = &limit
	request.Offset = &offset
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterInstances(request)
	if err != nil {
		errRet = err
		return
	}
	if total < 0 {
		total = int64(*response.Response.TotalCount)
	}

	if len(response.Response.InstanceSet) > 0 {
		offset += limit
	} else {
		// get empty set, we're done
		return
	}

	for _, item := range response.Response.InstanceSet {
		if has[*item.InstanceId] {
			errRet = fmt.Errorf("get repeated instance_id[%s] when doing DescribeClusterInstances", *item.InstanceId)
			return
		}
		has[*item.InstanceId] = true
		instanceInfo := InstanceInfo{
			InstanceId:               *item.InstanceId,
			InstanceRole:             *item.InstanceRole,
			InstanceState:            *item.InstanceState,
			FailedReason:             *item.FailedReason,
			InstanceAdvancedSettings: item.InstanceAdvancedSettings,
		}
		if item.CreatedTime != nil {
			instanceInfo.CreatedTime = *item.CreatedTime
		}
		//if item.NodePoolId != nil {
		//	instanceInfo.NodePoolId = *item.NodePoolId
		//}
		//if item.LanIP != nil {
		//	instanceInfo.LanIp = *item.LanIP
		//}
		if instanceInfo.InstanceRole == TKE_ROLE_WORKER {
			workers = append(workers, instanceInfo)
		} else {
			masters = append(masters, instanceInfo)
		}
	}
	goto getMoreData

}

func (me *TkeService) ScaleOutClusterMaster(ctx context.Context, clusterId string, runInstancesForNode []*tke.RunInstancesForNode) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewScaleOutClusterMasterRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &clusterId
	request.RunInstancesForNode = runInstancesForNode

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ScaleOutClusterMaster(request)
	if err != nil {
		return err
	}
	return nil
}

func (me *TkeService) ScaleInClusterMaster(ctx context.Context, clusterId string, scaleInMasters []*tke.ScaleInMaster) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewScaleInClusterMasterRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &clusterId
	request.ScaleInMasters = scaleInMasters

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ScaleInClusterMaster(request)
	if err != nil {
		return err
	}
	return nil
}

func (me *TkeService) DescribeClusters(ctx context.Context, id string, name string) (clusterInfos []ClusterInfo, errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDescribeClustersRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	if id != "" && name != "" {
		errRet = fmt.Errorf("cluster_id, cluster_name only one can be set one")
		return
	}

	if id != "" {
		request.ClusterIds = []*string{&id}
	}

	if name != "" {
		filter := &tke.Filter{
			Name:   helper.String("ClusterName"),
			Values: []*string{&name},
		}
		request.Filters = []*tke.Filter{filter}
	}

	response, err := me.client.UseTkeClient().DescribeClusters(request)

	if err != nil {
		errRet = err
		return
	}

	lenClusters := len(response.Response.Clusters)

	if lenClusters == 0 {
		return
	}
	clusterInfos = make([]ClusterInfo, 0, lenClusters)

	for index := range response.Response.Clusters {
		cluster := response.Response.Clusters[index]
		var clusterInfo ClusterInfo

		clusterInfo.ClusterId = *cluster.ClusterId
		clusterInfo.ClusterOs = *cluster.ClusterOs
		clusterInfo.ClusterVersion = *cluster.ClusterVersion
		clusterInfo.ClusterDescription = *cluster.ClusterDescription
		clusterInfo.ClusterName = *cluster.ClusterName
		clusterInfo.ClusterStatus = *cluster.ClusterStatus
		if cluster.ClusterLevel != nil {
			clusterInfo.ClusterLevel = cluster.ClusterLevel
		}
		if cluster.AutoUpgradeClusterLevel != nil {
			clusterInfo.AutoUpgradeClusterLevel = cluster.AutoUpgradeClusterLevel
		}

		clusterInfo.ProjectId = int64(*cluster.ProjectId)
		clusterInfo.VpcId = *cluster.ClusterNetworkSettings.VpcId
		clusterInfo.ClusterNodeNum = int64(*cluster.ClusterNodeNum)

		clusterInfo.IgnoreClusterCidrConflict = *cluster.ClusterNetworkSettings.IgnoreClusterCIDRConflict
		clusterInfo.ClusterCidr = *cluster.ClusterNetworkSettings.ClusterCIDR
		if cluster.ClusterNetworkSettings.IgnoreServiceCIDRConflict != nil {
			clusterInfo.IgnoreServiceCIDRConflict = *cluster.ClusterNetworkSettings.IgnoreServiceCIDRConflict
		}
		clusterInfo.MaxClusterServiceNum = int64(*cluster.ClusterNetworkSettings.MaxClusterServiceNum)

		clusterInfo.MaxNodePodNum = int64(*cluster.ClusterNetworkSettings.MaxNodePodNum)
		clusterInfo.DeployType = strings.ToUpper(*cluster.ClusterType)
		clusterInfo.Ipvs = *cluster.ClusterNetworkSettings.Ipvs
		clusterInfo.CreatedTime = *cluster.CreatedTime

		if len(cluster.TagSpecification) > 0 {
			clusterInfo.Tags = make(map[string]string)
			for _, tag := range cluster.TagSpecification[0].Tags {
				clusterInfo.Tags[*tag.Key] = *tag.Value
			}
		}

		clusterInfos = append(clusterInfos, clusterInfo)
	}
	return
}

func (me *TkeService) DescribeCluster(ctx context.Context, id string) (
	clusterInfo ClusterInfo,
	has bool,
	errRet error,
) {

	logId := getLogId(ctx)
	request := tke.NewDescribeClustersRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterIds = []*string{&id}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusters(request)

	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.Clusters) == 0 {
		return
	}

	has = true
	cluster := response.Response.Clusters[0]
	clusterInfo.ClusterId = *cluster.ClusterId
	clusterInfo.ClusterOs = *cluster.ClusterOs
	clusterInfo.ClusterVersion = *cluster.ClusterVersion
	clusterInfo.ClusterDescription = *cluster.ClusterDescription
	clusterInfo.ClusterName = *cluster.ClusterName
	clusterInfo.ClusterStatus = *cluster.ClusterStatus
	if cluster.ClusterLevel != nil {
		clusterInfo.ClusterLevel = cluster.ClusterLevel
	}
	if cluster.AutoUpgradeClusterLevel != nil {
		clusterInfo.AutoUpgradeClusterLevel = cluster.AutoUpgradeClusterLevel
	}

	clusterInfo.ProjectId = int64(*cluster.ProjectId)
	clusterInfo.VpcId = *cluster.ClusterNetworkSettings.VpcId
	clusterInfo.ClusterNodeNum = int64(*cluster.ClusterNodeNum)

	clusterInfo.IgnoreClusterCidrConflict = *cluster.ClusterNetworkSettings.IgnoreClusterCIDRConflict
	clusterInfo.ClusterCidr = *cluster.ClusterNetworkSettings.ClusterCIDR
	if cluster.ClusterNetworkSettings.IgnoreServiceCIDRConflict != nil {
		clusterInfo.IgnoreServiceCIDRConflict = *cluster.ClusterNetworkSettings.IgnoreServiceCIDRConflict
	}
	clusterInfo.MaxClusterServiceNum = int64(*cluster.ClusterNetworkSettings.MaxClusterServiceNum)

	clusterInfo.MaxNodePodNum = int64(*cluster.ClusterNetworkSettings.MaxNodePodNum)
	clusterInfo.DeployType = strings.ToUpper(*cluster.ClusterType)
	clusterInfo.Ipvs = *cluster.ClusterNetworkSettings.Ipvs
	if cluster.ClusterNetworkSettings.Cni != nil && *cluster.ClusterNetworkSettings.Cni {
		clusterInfo.NetworkType = TKE_CLUSTER_NETWORK_TYPE_VPC_CNI
	} else {
		clusterInfo.NetworkType = TKE_CLUSTER_NETWORK_TYPE_GR
	}

	if len(cluster.TagSpecification) > 0 {
		clusterInfo.Tags = make(map[string]string)
		for _, tag := range cluster.TagSpecification[0].Tags {
			clusterInfo.Tags[*tag.Key] = *tag.Value
		}
	}

	return
}

func (me *TkeService) DescribeClusterCommonNames(ctx context.Context,
	request *tke.DescribeClusterCommonNamesRequest) (commonNames []*tke.CommonName, errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterCommonNames(request)

	if err != nil {
		errRet = err
		return
	}

	commonNames = response.Response.CommonNames

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

//func (me *TkeService) DescribeClusterLevelAttribute(ctx context.Context, id string) (clusterLevels []*tke.ClusterLevelAttribute, errRet error) {
//	logId := getLogId(ctx)
//	request := tke.NewDescribeClusterLevelAttributeRequest()
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
//				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
//		}
//	}()
//
//	if id != "" {
//		request.ClusterID = &id
//	}
//	ratelimit.Check(request.GetAction())
//	response, err := me.client.UseTkeClient().DescribeClusterLevelAttribute(request)
//
//	if err != nil {
//		errRet = err
//		return
//	}
//
//	clusterLevels = response.Response.Items
//
//	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
//		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
//
//	return
//}

func (me *TkeService) DescribeClusterConfig(ctx context.Context, id string, isPublic bool) (config string, errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDescribeClusterKubeconfigRequest()
	// SDK 字段为 Type（yunapi 元数据规范字段名），取值 INNERLB=内网 / INTERNETLB=外网。
	// 旧代码误用不存在的 IsExtranet 字段被注释掉，导致 kube_config / kube_config_intranet 始终为空。
	clusterType := "INNERLB"
	if isPublic {
		clusterType = "INTERNETLB"
	}
	request.Type = &clusterType

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &id

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterKubeconfig(request)

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil {
		return
	}

	config = *response.Response.Kubeconfig
	return
}

func (me *TkeService) GetUpgradeInstanceResult(ctx context.Context, id string) (
	done bool,
	errRet error,
) {

	logId := getLogId(ctx)
	request := tke.NewGetUpgradeInstanceProgressRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &id

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().GetUpgradeInstanceProgress(request)

	if err != nil {
		errRet = err
		return
	}

	lifeState := *response.Response.LifeState

	// all instances success, lifeState=done
	if lifeState == "done" {
		return true, nil
	} else if lifeState != "process" {
		return false, fmt.Errorf("upgrade instances failed, tke response lifeState is:%s", lifeState)
	}

	// parent lifeState=process, check whether all instances in processing.
	for _, inst := range response.Response.Instances {
		if *inst.LifeState == "done" || *inst.LifeState == "pending" {
			continue
		}
		if *inst.LifeState != "process" {
			return false, fmt.Errorf("upgrade instances failed, "+
				"instanceId:%s, lifeState is:%s", *inst.InstanceID, *inst.LifeState)
		}
		// instance lifeState=process, check whether failed or not.
		for _, detail := range inst.Detail {
			if *detail.LifeState == "failed" {
				return false, fmt.Errorf("upgrade instances failed, "+
					"instanceId:%s, detail.lifeState is:%s", *inst.InstanceID, *detail.LifeState)
			}
		}
	}

	return
}

func (me *TkeService) CreateCluster(ctx context.Context,
	basic ClusterBasicSetting,
	advanced ClusterAdvancedSettings,
	cvms RunInstancesForNode,
	runInstancesForNode []*tke.RunInstancesForNode,
	clusterDeployType string,
	iAdvanced InstanceAdvancedSettings,
	cidrSetting tke.ClusterCIDRSettings,
	tags map[string]string,
	existedInstance []*tke.ExistedInstancesForNode,
	overrideSettings *OverrideSettings,
	iDiskMountSettings []*tke.InstanceDataDiskMountSetting,
	clusterArch string,
	extensionAddons []*tke.ExtensionAddon,
	cdcId string,
	disableAddons []*string,
) (id string, errRet error) {

	logId := getLogId(ctx)
	request := tke.NewCreateClusterRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	if clusterArch != "" {
		request.ClusterArch = &clusterArch
	}

	request.ClusterBasicSettings = &tke.ClusterBasicSettings{}
	request.ClusterBasicSettings.ClusterOs = &basic.ClusterOs
	request.ClusterBasicSettings.ClusterVersion = &basic.ClusterVersion
	request.ClusterBasicSettings.ProjectId = &basic.ProjectId
	request.ClusterBasicSettings.VpcId = &basic.VpcId
	request.ClusterBasicSettings.ClusterDescription = &basic.ClusterDescription
	request.ClusterBasicSettings.ClusterName = &basic.ClusterName
	request.ClusterBasicSettings.OsCustomizeType = &basic.ClusterOsType
	if basic.SubnetId != "" {
		request.ClusterBasicSettings.SubnetId = &basic.SubnetId
	}
	request.ClusterBasicSettings.NeedWorkSecurityGroup = &basic.NeedWorkSecurityGroup
	if basic.ClusterLevel != nil {
		request.ClusterBasicSettings.ClusterLevel = basic.ClusterLevel
	}
	if basic.AutoUpgradeClusterLevel != nil {
		request.ClusterBasicSettings.AutoUpgradeClusterLevel = &tke.AutoUpgradeClusterLevel{
			IsAutoUpgrade: basic.AutoUpgradeClusterLevel,
		}
	}
	for k, v := range tags {
		if len(request.ClusterBasicSettings.TagSpecification) == 0 {
			request.ClusterBasicSettings.TagSpecification = []*tke.TagSpecification{{
				ResourceType: helper.String("cluster"),
			}}
		}

		request.ClusterBasicSettings.TagSpecification[0].Tags = append(request.ClusterBasicSettings.TagSpecification[0].Tags, &tke.Tag{
			Key:   helper.String(k),
			Value: helper.String(v),
		})
	}
	if len(tags) == 0 {
		request.ClusterBasicSettings.TagSpecification = []*tke.TagSpecification{{
			ResourceType: helper.String("cluster"),
			Tags:         []*tke.Tag{},
		}}
	}

	request.ClusterAdvancedSettings = &tke.ClusterAdvancedSettings{}
	request.ClusterAdvancedSettings.IPVS = &advanced.Ipvs
	request.ClusterAdvancedSettings.AsEnabled = &advanced.AsEnabled
	request.ClusterAdvancedSettings.ContainerRuntime = &advanced.ContainerRuntime
	request.ClusterAdvancedSettings.RuntimeVersion = &advanced.RuntimeVersion
	request.ClusterAdvancedSettings.NodeNameType = &advanced.NodeNameType
	request.ClusterAdvancedSettings.EnableCustomizedPodCIDR = &advanced.EnableCustomizedPodCIDR
	request.ClusterAdvancedSettings.BasePodNumber = &advanced.BasePodNumber
	// 确保 ExtraArgs 字段始终被初始化为空数组而不是 nil
	request.ClusterAdvancedSettings.ExtraArgs = &tke.ClusterExtraArgs{
		KubeAPIServer:         make([]*string, 0),
		KubeControllerManager: make([]*string, 0),
		KubeScheduler:         make([]*string, 0),
	}
	if len(advanced.ExtraArgs.KubeAPIServer) > 0 {
		request.ClusterAdvancedSettings.ExtraArgs.KubeAPIServer = common.StringPtrs(advanced.ExtraArgs.KubeAPIServer)
	}
	if len(advanced.ExtraArgs.KubeControllerManager) > 0 {
		request.ClusterAdvancedSettings.ExtraArgs.KubeControllerManager = common.StringPtrs(advanced.ExtraArgs.KubeControllerManager)
	}
	if len(advanced.ExtraArgs.KubeScheduler) > 0 {
		request.ClusterAdvancedSettings.ExtraArgs.KubeScheduler = common.StringPtrs(advanced.ExtraArgs.KubeScheduler)
	}
	request.ClusterAdvancedSettings.NetworkType = &advanced.NetworkType
	request.ClusterAdvancedSettings.IsNonStaticIpMode = &advanced.IsNonStaticIpMode
	if advanced.VpcCniType != "" {
		request.ClusterAdvancedSettings.VpcCniType = &advanced.VpcCniType
	}
	request.ClusterAdvancedSettings.DeletionProtection = &advanced.DeletionProtection
	request.ClusterAdvancedSettings.KubeProxyMode = &advanced.KubeProxyMode
	request.ClusterAdvancedSettings.AuditEnabled = &advanced.AuditEnabled
	request.ClusterAdvancedSettings.AuditLogsetId = &advanced.AuditLogsetId
	request.ClusterAdvancedSettings.AuditLogTopicId = &advanced.AuditLogTopicId
	request.ClusterAdvancedSettings.DataPlaneV2 = &advanced.DataPlaneV2
	request.ClusterAdvancedSettings.QGPUShareEnable = &advanced.QGPUShareEnable
	request.ClusterAdvancedSettings.IsDualStack = &advanced.IsDualStack

	hasInstanceAdvanced := iAdvanced.MountTarget != "" ||
		iAdvanced.PreStartUserScript != "" ||
		iAdvanced.DockerGraphPath != "" ||
		iAdvanced.UserScript != "" ||
		iAdvanced.Unschedulable != 0 ||
		iAdvanced.DesiredPodNumSet ||
		iAdvanced.DataDiskPartition != "" ||
		iAdvanced.GPUArgs != nil ||
		len(iAdvanced.Taints) > 0 ||
		len(iAdvanced.Labels) > 0 ||
		len(iAdvanced.DataDisks) > 0 ||
		len(iAdvanced.ExtraArgs.Kubelet) > 0

	if hasInstanceAdvanced {
		request.InstanceAdvancedSettings = &tke.InstanceAdvancedSettings{}
		if iAdvanced.MountTarget != "" {
			request.InstanceAdvancedSettings.MountTarget = &iAdvanced.MountTarget
		}
		if iAdvanced.PreStartUserScript != "" {
			request.InstanceAdvancedSettings.PreStartUserScript = &iAdvanced.PreStartUserScript
		}
		if iAdvanced.DockerGraphPath != "" {
			request.InstanceAdvancedSettings.DockerGraphPath = &iAdvanced.DockerGraphPath
		}
		if iAdvanced.UserScript != "" {
			request.InstanceAdvancedSettings.UserScript = &iAdvanced.UserScript
		}
		if iAdvanced.Unschedulable != 0 {
			request.InstanceAdvancedSettings.Unschedulable = &iAdvanced.Unschedulable
		}
		if iAdvanced.DesiredPodNumSet {
			request.InstanceAdvancedSettings.DesiredPodNumber = &iAdvanced.DesiredPodNum
		}
		if iAdvanced.DataDiskPartition != "" {
			request.InstanceAdvancedSettings.DataDiskPartition = &iAdvanced.DataDiskPartition
		}
		if iAdvanced.GPUArgs != nil {
			request.InstanceAdvancedSettings.GPUArgs = iAdvanced.GPUArgs
		}
		if len(iAdvanced.Taints) > 0 {
			request.InstanceAdvancedSettings.Taints = iAdvanced.Taints
		}
		if len(iAdvanced.ExtraArgs.Kubelet) > 0 {
			request.InstanceAdvancedSettings.ExtraArgs = &iAdvanced.ExtraArgs
		}
		if len(iAdvanced.Labels) > 0 {
			request.InstanceAdvancedSettings.Labels = iAdvanced.Labels
		}
		if len(iAdvanced.DataDisks) > 0 {
			request.InstanceAdvancedSettings.DataDisks = iAdvanced.DataDisks
		}
	}

	if len(extensionAddons) > 0 {
		request.ExtensionAddons = extensionAddons
	}

	if len(disableAddons) > 0 {
		request.DisableAddons = disableAddons
	}

	if cdcId != "" {
		request.CdcId = &cdcId
	}

	if overrideSettings != nil {
		if len(overrideSettings.Master)+len(overrideSettings.Work) > 0 &&
			len(overrideSettings.Master)+len(overrideSettings.Work) != (len(cvms.Master)+len(cvms.Work)) {
			return "", fmt.Errorf("len(overrideSettings) != (len(cvms.Master)+len(cvms.Work))")
		}
	}

	request.RunInstancesForNode = []*tke.RunInstancesForNode{}
	if len(runInstancesForNode) != 0 {
		request.RunInstancesForNode = runInstancesForNode
		if clusterDeployType != "" {
			request.ClusterType = helper.String(clusterDeployType)
		}
	} else if cdcId != "" {
		request.ClusterType = helper.String(clusterDeployType)
	} else if len(cvms.Master) != 0 {
		var node tke.RunInstancesForNode
		node.NodeRole = helper.String(TKE_ROLE_MASTER_ETCD)
		node.RunInstancesPara = []*string{}
		request.ClusterType = helper.String(TKE_DEPLOY_TYPE_INDEPENDENT)
		for v := range cvms.Master {
			node.RunInstancesPara = append(node.RunInstancesPara, &cvms.Master[v])
			if overrideSettings != nil && len(overrideSettings.Master) != 0 {
				node.InstanceAdvancedSettingsOverrides = append(node.InstanceAdvancedSettingsOverrides, &overrideSettings.Master[v])
			}
		}
		request.RunInstancesForNode = append(request.RunInstancesForNode, &node)
	} else {
		request.ClusterType = helper.String(TKE_DEPLOY_TYPE_MANAGED)
	}

	if len(runInstancesForNode) == 0 && len(cvms.Work) != 0 {
		var node tke.RunInstancesForNode
		node.NodeRole = helper.String(TKE_ROLE_WORKER)
		node.RunInstancesPara = []*string{}
		for v := range cvms.Work {
			node.RunInstancesPara = append(node.RunInstancesPara, &cvms.Work[v])
			if overrideSettings != nil && len(overrideSettings.Work) != 0 {
				node.InstanceAdvancedSettingsOverrides = append(node.InstanceAdvancedSettingsOverrides, &overrideSettings.Work[v])
			}
		}
		request.RunInstancesForNode = append(request.RunInstancesForNode, &node)
	}

	if len(iDiskMountSettings) != 0 {
		request.InstanceDataDiskMountSettings = iDiskMountSettings
	}

	// 直接使用传入的 cidrSetting，它已经是 tke.ClusterCIDRSettings 类型
	request.ClusterCIDRSettings = &cidrSetting

	if len(existedInstance) > 0 {
		request.ExistedInstancesForNode = existedInstance
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().CreateCluster(request)

	if err != nil {
		errRet = err
		return
	}

	id = *response.Response.ClusterId
	return
}

func (me *TkeService) CreateClusterInstances(ctx context.Context,
	id string, runInstancePara string,
	iAdvanced tke.InstanceAdvancedSettings) (instanceIds []string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewCreateClusterInstancesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.RunInstancePara = &runInstancePara

	request.InstanceAdvancedSettings = &iAdvanced

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().CreateClusterInstances(request)

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil {
		errRet = fmt.Errorf("CreateClusterInstances return nil response")
		return
	}

	instanceIds = make([]string, 0, len(response.Response.InstanceIdSet))

	for _, v := range response.Response.InstanceIdSet {

		instanceIds = append(instanceIds, *v)
	}
	return
}

func (me *TkeService) CheckOneOfClusterNodeReady(ctx context.Context, clusterId, nodePoolId string, mustHaveWorkers bool) error {
	return resource.Retry(readRetryTimeout*5, func() *resource.RetryError {

		status, err := me.DescribeClusterStatus(ctx, clusterId)
		if err != nil {
			return retryError(err)
		}

		if *status.ClusterState == "Running" {
			return nil
		}
		// _, workers, err := me.DescribeClusterInstances(ctx, clusterId)
		// if err != nil {
		// 	return retryError(err)
		// }

		// // check serverless node
		// virtualNodes, err := me.DescribeClusterVirtualNode(ctx, clusterId, nodePoolId)
		// if err != nil {
		// 	return retryError(err)
		// }

		// if len(workers) == 0 && len(virtualNodes) == 0 {
		// 	if mustHaveWorkers {
		// 		return resource.RetryableError(fmt.Errorf("waiting for workers created"))
		// 	}
		// 	return nil
		// }

		// for i := range workers {
		// 	worker := workers[i]
		// 	if worker.InstanceState == "running" {
		// 		return nil
		// 	}
		// }
		// for i := range virtualNodes {
		// 	virtualNode := virtualNodes[i]
		// 	if virtualNode.Phase != nil && *virtualNode.Phase == "Running" {
		// 		return nil
		// 	}
		// }

		return resource.RetryableError(fmt.Errorf("cluster %s waiting for status ready, current status is %s", clusterId, *status.ClusterState))
	})
}

/*
if cluster is creating, return error:CloudSDKError] Code=InternalError.ClusterState
*/
func (me *TkeService) DeleteClusterInstances(ctx context.Context, id string, instanceIds []string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDeleteClusterInstancesRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.InstanceIds = make([]*string, 0, len(instanceIds))

	for index := range instanceIds {
		request.InstanceIds = append(request.InstanceIds, &instanceIds[index])
	}

	request.InstanceDeleteMode = helper.String("terminate")
	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().DeleteClusterInstances(request)
	return err
}

func (me *TkeService) DeleteCluster(ctx context.Context, id string, instanceDeleteMode string, resourceDeleteOptions []*tke.ResourceDeleteOption) (errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDeleteClusterRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	if instanceDeleteMode != "" {
		request.InstanceDeleteMode = &instanceDeleteMode
	} else {
		request.InstanceDeleteMode = helper.String("terminate")
	}
	if len(resourceDeleteOptions) > 0 {
		request.ResourceDeleteOptions = resourceDeleteOptions
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().DeleteCluster(request)

	return err
}

func (me *TkeService) DescribeClusterSecurity(ctx context.Context, id string) (ret *tke.DescribeClusterSecurityResponse, errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDescribeClusterSecurityRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &id

	return me.client.UseTkeClient().DescribeClusterSecurity(request)
}

func (me *TkeService) DescribeClusterEndpoints(ctx context.Context, id string) (ret *tke.DescribeClusterEndpointsResponse, errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDescribeClusterEndpointsRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &id

	return me.client.UseTkeClient().DescribeClusterEndpoints(request)
}

func (me *TkeService) CreateClusterAsGroup(ctx context.Context, id, groupPara, configPara string, labels []*tke.Label, iAdvanced InstanceAdvancedSettings) (asGroupId string, errRet error) {
	return "", fmt.Errorf("Cluster AS Group has OFFLINE")
}

func (me *TkeService) DescribeClusterAsGroupsByGroupId(ctx context.Context, id string, groupId string) (clusterAsGroupSet *tke.ClusterAsGroup, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeClusterAsGroupsRequest()

	request.ClusterId = &id
	request.AutoScalingGroupIds = []*string{&groupId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterAsGroups(request)

	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), err.Error())
		errRet = err
		return
	}

	if len(response.Response.ClusterAsGroupSet) > 0 {
		clusterAsGroupSet = response.Response.ClusterAsGroupSet[0]
	}
	return
}

func (me *TkeService) DeleteClusterAsGroups(ctx context.Context, id, asGroupId string) (errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDeleteClusterAsGroupsRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.AutoScalingGroupIds = []*string{&asGroupId}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().DeleteClusterAsGroups(request)
	if err != nil {
		errRet = err
	}
	return
}

/*
open internet access
*/
func (me *TkeService) CreateClusterEndpoint(ctx context.Context, id string, subnetId, securityGroupId string, internet bool, domain string, extensiveParameters string, existedLoadBalancerId string) (errRet error) {
	logId := getLogId(ctx)

	request := tke.NewCreateClusterEndpointRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id

	request.IsExtranet = &internet

	if subnetId != "" {
		request.SubnetId = &subnetId
	}

	if securityGroupId != "" && internet {
		request.SecurityGroup = &securityGroupId
	}

	if domain != "" {
		request.Domain = helper.String(domain)
	}

	if extensiveParameters != "" {
		request.ExtensiveParameters = helper.String(extensiveParameters)
	}

	if existedLoadBalancerId != "" {
		request.ExistedLoadBalancerId = helper.String(existedLoadBalancerId)
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().CreateClusterEndpoint(request)
	if err != nil {
		errRet = err
	}
	return
}

func (me *TkeService) DescribeClusterEndpointStatus(ctx context.Context, id string, isExtranet bool) (status string, message string, errRet error) {
	logId := getLogId(ctx)

	request := tke.NewDescribeClusterEndpointStatusRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.IsExtranet = &isExtranet

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTkeClient().DescribeClusterEndpointStatus(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response == nil || response.Response.Status == nil {
		errRet = fmt.Errorf("sdk DescribeClusterEndpointStatus return empty status")
		return
	}
	status = *response.Response.Status
	message = status
	return
}

func (me *TkeService) DeleteClusterEndpoint(ctx context.Context, id string, isInternet bool) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDeleteClusterEndpointRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.IsExtranet = &isInternet

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().DeleteClusterEndpoint(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) ModifyClusterEndpointSP(ctx context.Context, id string, securityPolicies []string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewModifyClusterEndpointSPRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.SecurityPolicies = helper.Strings(securityPolicies)

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().ModifyClusterEndpointSP(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) ModifyClusterEndpointSG(ctx context.Context, id string, securityGroup string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewModifyClusterEndpointSPRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.SecurityPolicies = []*string{&securityGroup}

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().ModifyClusterEndpointSP(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) ModifyClusterAttribute(ctx context.Context, id string, projectId int64, clusterName, clusterDesc, clusterLevel string, autoUpgradeClusterLevel bool) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewModifyClusterAttributeRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.ProjectId = &projectId
	request.ClusterName = &clusterName
	request.ClusterDesc = &clusterDesc

	if clusterLevel != "" {
		request.ClusterLevel = &clusterLevel
	}

	request.AutoUpgradeClusterLevel = &tke.AutoUpgradeClusterLevel{
		IsAutoUpgrade: &autoUpgradeClusterLevel,
	}

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().ModifyClusterAttribute(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) ModifyClusterVersion(ctx context.Context, id string, clusterVersion string, extraArgs interface{}) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewUpdateClusterVersionRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	if extraArgs != nil && len(extraArgs.([]interface{})) > 0 {
		// the first elem is in use
		extraInterface := extraArgs.([]interface{})
		extraMap := extraInterface[0].(map[string]interface{})

		kas := make([]*string, 0)
		if kaArgs, exist := extraMap["kube_apiserver"]; exist {
			args := kaArgs.([]interface{})
			for index := range args {
				str := args[index].(string)
				kas = append(kas, &str)
			}
		}
		kcms := make([]*string, 0)
		if kcmArgs, exist := extraMap["kube_controller_manager"]; exist {
			args := kcmArgs.([]interface{})
			for index := range args {
				str := args[index].(string)
				kcms = append(kcms, &str)
			}
		}
		kss := make([]*string, 0)
		if ksArgs, exist := extraMap["kube_scheduler"]; exist {
			args := ksArgs.([]interface{})
			for index := range args {
				str := args[index].(string)
				kss = append(kss, &str)
			}
		}

		request.ExtraArgs = &tke.ClusterExtraArgs{
			KubeAPIServer:         kas,
			KubeControllerManager: kcms,
			KubeScheduler:         kss,
		}
	}

	request.ClusterId = &id
	request.DstVersion = &clusterVersion

	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().UpdateClusterVersion(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) DescribeKubernetesAvailableClusterVersionsByFilter(ctx context.Context,
	param map[string]interface{}) (availableClusterVersions []*string, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tke.NewDescribeAvailableClusterVersionRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "cluster_id" {
			request.ClusterId = v.(*string)
		}
		if k == "cluster_ids" {
			request.ClusterIds = v.([]*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTkeClient().DescribeAvailableClusterVersion(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil {
		availableClusterVersions = response.Response.Versions
	}
	return
}

func (me *TkeService) CheckClusterVersion(ctx context.Context, id string, clusterVersion string) (isOk bool, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeAvailableClusterVersionRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	ratelimit.Check(request.GetAction())

	resp, err := me.client.UseTkeClient().DescribeAvailableClusterVersion(request)
	if err != nil {
		errRet = err
		return
	}

	if resp == nil || resp.Response == nil || resp.Response.Versions == nil {
		return
	}
	versions := resp.Response.Versions
	for _, v := range versions {
		if *v == clusterVersion {
			isOk = true
			return
		}
	}

	return
}

func (me *TkeService) CheckInstancesUpgradeAble(ctx context.Context, id string, upgradeType string) (instanceIds []string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewCheckInstancesUpgradeAbleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.UpgradeType = &upgradeType
	ratelimit.Check(request.GetAction())

	resp, err := me.client.UseTkeClient().CheckInstancesUpgradeAble(request)
	if err != nil {
		errRet = err
		return
	}

	if resp == nil || resp.Response == nil || resp.Response.UpgradeAbleInstances == nil {
		return
	}
	for _, inst := range resp.Response.UpgradeAbleInstances {
		instanceIds = append(instanceIds, *inst.InstanceId)
	}

	return
}

func (me *TkeService) UpgradeClusterInstances(ctx context.Context, id string, upgradeType string, instanceIds []string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewUpgradeClusterInstancesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	op := "create"
	request.Operation = &op
	request.ClusterId = &id
	request.UpgradeType = &upgradeType
	request.InstanceIds = helper.Strings(instanceIds)
	ratelimit.Check(request.GetAction())

	_, err := me.client.UseTkeClient().UpgradeClusterInstances(request)
	if err != nil {
		errRet = err
		return
	}

	return
}

func (me *TkeService) DescribeImages(ctx context.Context) (imageIds []string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeImagesRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeImages(request)
	if err != nil {
		errRet = err
		return
	}

	for _, image := range response.Response.ImageInstanceSet {
		imageIds = append(imageIds, *image.ImageId)
	}
	return
}

func GetTkeLabels(d *schema.ResourceData, k string) []*tke.Label {
	labels := make([]*tke.Label, 0)
	if raw, ok := d.GetOk(k); ok {
		for k, v := range raw.(map[string]interface{}) {
			labels = append(labels, &tke.Label{Name: helper.String(k), Value: helper.String(v.(string))})
		}
	}
	return labels
}

func GetTkeTaints(d *schema.ResourceData, k string) []*tke.Taint {
	taints := make([]*tke.Taint, 0)
	if raw, ok := d.GetOk(k); ok {
		for _, v := range raw.([]interface{}) {
			vv := v.(map[string]interface{})
			taints = append(taints, &tke.Taint{Key: helper.String(vv["key"].(string)), Value: helper.String(vv["value"].(string)), Effect: helper.String(vv["effect"].(string))})
		}
	}
	return taints
}

func (me *TkeService) ModifyClusterAsGroupAttribute(ctx context.Context, id, asGroupId string, maxSize, minSize int64) (errRet error) {

	logId := getLogId(ctx)
	request := tke.NewModifyClusterAsGroupAttributeRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.ClusterAsGroupAttribute = &tke.ClusterAsGroupAttribute{
		AutoScalingGroupId: &asGroupId,
		AutoScalingGroupRange: &tke.AutoScalingGroupRange{
			MaxSize: &maxSize,
			MinSize: &minSize,
		},
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ModifyClusterAsGroupAttribute(request)
	if err != nil {
		errRet = err
	}
	return
}

func (me *TkeService) CreateClusterNodePool(ctx context.Context, clusterId, name, groupPara, configPara string,
	enableAutoScale bool, nodeOs string, nodeOsType string, labels []*tke.Label, taints []*tke.Taint,
	iAdvanced *tke.InstanceAdvancedSettings, deletionProtection bool, annotations []*tke.AnnotationValue,
	containerRuntime string, runtimeVersion string, tags []*tke.Tag) (nodePoolId string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewCreateClusterNodePoolRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &clusterId
	request.Name = &name
	request.AutoScalingGroupPara = &groupPara
	request.LaunchConfigurePara = &configPara
	if iAdvanced != nil {
		request.InstanceAdvancedSettings = iAdvanced
	}
	request.EnableAutoscale = &enableAutoScale
	request.DeletionProtection = &deletionProtection
	request.NodePoolOs = &nodeOs
	request.OsCustomizeType = &nodeOsType

	if len(labels) > 0 {
		request.Labels = labels
	}

	if len(taints) > 0 {
		request.Taints = taints
	}

	if len(annotations) > 0 {
		request.Annotations = annotations
	}

	if containerRuntime != "" {
		request.ContainerRuntime = &containerRuntime
	}

	if runtimeVersion != "" {
		request.RuntimeVersion = &runtimeVersion
	}

	if len(tags) > 0 {
		request.Tags = tags
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().CreateClusterNodePool(request)
	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil || response.Response.NodePoolId == nil {
		errRet = fmt.Errorf("CreateClusterNodePool return nil response")
		return
	}

	nodePoolId = *response.Response.NodePoolId
	return
}

func (me *TkeService) ModifyClusterNodePool(ctx context.Context, clusterId, nodePoolId string, name string, enableAutoScale bool, minSize int64, maxSize int64, nodeOs string, nodeOsType string, labels []*tke.Label, taints []*tke.Taint, tags map[string]string, deletionProtection *bool, annotations []*tke.AnnotationValue) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewModifyClusterNodePoolRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &clusterId
	request.NodePoolId = &nodePoolId
	request.Name = &name
	request.EnableAutoscale = &enableAutoScale
	request.MaxNodesNum = &maxSize
	request.MinNodesNum = &minSize
	request.OsName = &nodeOs
	request.OsCustomizeType = &nodeOsType

	if len(labels) > 0 {
		request.Labels = labels
	}

	if len(taints) > 0 {
		request.Taints = taints
	}

	if len(tags) > 0 {
		for k, v := range tags {
			key := k
			val := v
			request.Tags = append(request.Tags, &tke.Tag{
				Key:   &key,
				Value: &val,
			})
		}
	}

	if deletionProtection != nil {
		request.DeletionProtection = deletionProtection
	}

	if len(annotations) > 0 {
		request.Annotations = annotations
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ModifyClusterNodePool(request)
	if err != nil {
		errRet = err
		return
	}
	return
}

func (me *TkeService) ModifyClusterNodePoolPreStartUserScript(ctx context.Context, clusterId, nodePoolId, userData, preStartUserScript string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewModifyClusterNodePoolRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	request.ClusterId = &clusterId
	request.NodePoolId = &nodePoolId
	request.UserScript = &userData
	request.PreStartUserScript = &preStartUserScript

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ModifyClusterNodePool(request)
	if err != nil {
		errRet = err
	}
	return
}

//func (me *TkeService) ModifyClusterNodePoolDesiredCapacity(ctx context.Context, clusterId, nodePoolId string, desiredCapacity int64) (errRet error) {
//	logId := getLogId(ctx)
//	request := tke.NewModifyNodePoolDesiredCapacityAboutAsgRequest()
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
//		}
//	}()
//	request.ClusterId = &clusterId
//	request.NodePoolId = &nodePoolId
//	request.DesiredCapacity = &desiredCapacity
//
//	ratelimit.Check(request.GetAction())
//	_, err := me.client.UseTkeClient().ModifyNodePoolDesiredCapacityAboutAsg(request)
//	if err != nil {
//		errRet = err
//		return
//	}
//	return
//}

//func (me *TkeService) ModifyClusterNodePoolInstanceTypes(ctx context.Context, clusterId, nodePoolId string, instanceTypes []*string) (errRet error) {
//	logId := getLogId(ctx)
//	request := tke.NewModifyNodePoolInstanceTypesRequest()
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
//		}
//	}()
//	request.ClusterId = &clusterId
//	request.NodePoolId = &nodePoolId
//	request.InstanceTypes = instanceTypes
//
//	ratelimit.Check(request.GetAction())
//	_, err := me.client.UseTkeClient().ModifyNodePoolInstanceTypes(request)
//	if err != nil {
//		errRet = err
//		return
//	}
//	return
//}

func (me *TkeService) DeleteClusterNodePool(ctx context.Context, id, nodePoolId string, deleteKeepInstance bool) (errRet error) {

	logId := getLogId(ctx)
	request := tke.NewDeleteClusterNodePoolRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()
	request.ClusterId = &id
	request.NodePoolIds = []*string{&nodePoolId}
	request.KeepInstance = &deleteKeepInstance

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().DeleteClusterNodePool(request)
	if err != nil {
		errRet = err
	}
	return
}

func (me *TkeService) DescribeNodePool(ctx context.Context, clusterId string, nodePoolId string) (
	nodePool *tke.NodePool,
	has bool,
	errRet error,
) {

	logId := getLogId(ctx)
	//the error code of cluster not exist is InternalError
	//check cluster exist first
	_, clusterHas, err := me.DescribeCluster(ctx, clusterId)
	if err != nil {
		errRet = err
		return
	}
	if !clusterHas {
		return
	}

	request := tke.NewDescribeClusterNodePoolDetailRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = helper.String(clusterId)
	request.NodePoolId = helper.String(nodePoolId)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterNodePoolDetail(request)

	if err != nil {
		errRet = err
		return
	}

	if response.Response.NodePool == nil {
		return
	}

	has = true
	nodePool = response.Response.NodePool

	return
}

// node pool global config
func (me *TkeService) ModifyClusterNodePoolGlobalConfig(ctx context.Context, request *tke.ModifyClusterAsGroupOptionAttributeRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseTkeClient().ModifyClusterAsGroupOptionAttribute(request)
	if err != nil {
		errRet = err
	}
	return
}

func (me *TkeService) DescribeClusterNodePoolGlobalConfig(ctx context.Context, clusterId string) (
	npGlobalConfig *tke.ClusterAsGroupOption,
	errRet error,
) {

	logId := getLogId(ctx)
	//the error code of cluster not exist is InternalError
	//check cluster exist first
	_, clusterHas, err := me.DescribeCluster(ctx, clusterId)
	if err != nil {
		errRet = err
		return
	}
	if !clusterHas {
		return
	}

	request := tke.NewDescribeClusterAsGroupOptionRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = helper.String(clusterId)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterAsGroupOption(request)

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil || response.Response.ClusterAsGroupOption == nil {
		return
	}

	npGlobalConfig = response.Response.ClusterAsGroupOption

	return
}

func (me *TkeService) WaitForAuthenticationOptionsUpdateSuccess(ctx context.Context, id string) (info *tke.ServiceAccountAuthenticationOptions, errRet error) {
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		options, state, _, err := me.DescribeClusterAuthenticationOptions(ctx, id)
		info = options

		if err != nil {
			return resource.NonRetryableError(err)
		}

		if state == "Success" {
			return nil
		}

		if state == "Updating" {
			return resource.RetryableError(fmt.Errorf("state is %s, retry", state))
		}

		return resource.NonRetryableError(fmt.Errorf("update failed: %s", state))
	})

	if err != nil {
		errRet = err
		return
	}
	return
}

// DescribeClusterAuthenticationOptions
// Field `ServiceAccounts.AutoCreateDiscoveryAnonymousAuth` will always return null by design
func (me *TkeService) DescribeClusterAuthenticationOptions(ctx context.Context, id string) (options *tke.ServiceAccountAuthenticationOptions, state string, oidcConfig *tke.OIDCConfigAuthenticationOptions, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeClusterAuthenticationOptionsRequest()
	request.ClusterId = helper.String(id)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	res, err := me.client.UseTkeClient().DescribeClusterAuthenticationOptions(request)
	if err != nil {
		errRet = err
	}

	if res.Response != nil {
		state = *res.Response.LatestOperationState
		options = res.Response.ServiceAccounts
		oidcConfig = res.Response.OIDCConfig
	}

	return
}

func (me *TkeService) ModifyClusterAuthenticationOptions(ctx context.Context, request *tke.ModifyClusterAuthenticationOptionsRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n", logId, request.GetAction(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().ModifyClusterAuthenticationOptions(request)
	if err != nil {
		errRet = err
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

//func (me *TkeService) ModifyDeletionProtection(ctx context.Context, id string, enable bool) (errRet error) {
//	var (
//		logId  = getLogId(ctx)
//		action string
//	)
//
//	if enable {
//		request := tke.NewEnableClusterDeletionProtectionRequest()
//		request.ClusterId = &id
//		action = request.GetAction()
//		ratelimit.Check(action)
//		response, err := me.client.UseTkeClient().EnableClusterDeletionProtection(request)
//		if err != nil {
//			errRet = err
//			return
//		}
//
//		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
//			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
//	} else {
//		request := tke.NewDisableClusterDeletionProtectionRequest()
//		request.ClusterId = &id
//		action = request.GetAction()
//		ratelimit.Check(action)
//		response, err := me.client.UseTkeClient().DisableClusterDeletionProtection(request)
//		if err != nil {
//			errRet = err
//			return
//		}
//
//		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
//			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
//	}
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
//				logId, action, errRet.Error())
//		}
//	}()
//
//	return
//}

func (me *TkeService) AcquireClusterAdminRole(ctx context.Context, clusterId string) (errRet error) {
	logId := getLogId(ctx)
	request := tke.NewAcquireClusterAdminRoleRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterId = &clusterId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().AcquireClusterAdminRole(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) SwitchLogAgent(ctx context.Context, clusterId, rootDir string, enable bool) error {
	if enable {
		request := tke.NewInstallLogAgentRequest()
		request.ClusterId = &clusterId
		if rootDir != "" {
			request.KubeletRootDir = &rootDir
		}
		return me.InstallLogAgent(ctx, request)
	}
	request := tke.NewUninstallLogAgentRequest()
	request.ClusterId = &clusterId
	return me.UninstallLogAgent(ctx, request)
}

func (me *TkeService) SwitchEventPersistence(ctx context.Context, clusterId, logSetId, topicId string,
	enable, deleteEventLog bool) error {
	if enable {
		request := tke.NewEnableEventPersistenceRequest()
		request.ClusterId = &clusterId
		if logSetId != "" {
			request.LogsetId = &logSetId
		}
		if topicId != "" {
			request.TopicId = &topicId
		}
		return me.EnableEventPersistence(ctx, request)
	}

	request := tke.NewDisableEventPersistenceRequest()
	request.ClusterId = &clusterId
	request.DeleteLogSetAndTopic = &deleteEventLog
	return me.DisableEventPersistence(ctx, request)
}

func (me *TkeService) SwitchClusterAudit(ctx context.Context, clusterId, logSetId, topicId string,
	enable, deleteAuditLog bool) error {
	if enable {
		request := tke.NewEnableClusterAuditRequest()
		request.ClusterId = &clusterId
		if logSetId != "" {
			request.LogsetId = &logSetId
		}
		if topicId != "" {
			request.TopicId = &topicId
		}
		return me.EnableClusterAudit(ctx, request)
	}
	request := tke.NewDisableClusterAuditRequest()
	request.ClusterId = &clusterId
	request.DeleteLogSetAndTopic = &deleteAuditLog
	return me.DisableClusterAudit(ctx, request)
}

func (me *TkeService) InstallLogAgent(ctx context.Context, request *tke.InstallLogAgentRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().InstallLogAgent(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) UninstallLogAgent(ctx context.Context, request *tke.UninstallLogAgentRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().UninstallLogAgent(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) EnableEventPersistence(ctx context.Context, request *tke.EnableEventPersistenceRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().EnableEventPersistence(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) DisableEventPersistence(ctx context.Context, request *tke.DisableEventPersistenceRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DisableEventPersistence(request)

	if err != nil {
		if sdkErr, ok := err.(*sdkErrors.CloudSDKError); ok {
			if sdkErr.Code == "InternalError.KubernetesDeleteOperationError" {
				return
			}
		}
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) EnableClusterAudit(ctx context.Context, request *tke.EnableClusterAuditRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().EnableClusterAudit(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) DisableClusterAudit(ctx context.Context, request *tke.DisableClusterAuditRequest) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DisableClusterAudit(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) DescribeKubernetesLogConfigById(ctx context.Context, clusterId, logConfigName, clusterType string) (ret *tke.DescribeLogConfigsResponse, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeLogConfigsRequest()
	request.ClusterId = helper.String(clusterId)
	request.ClusterType = helper.String(clusterType)
	request.LogConfigNames = helper.String(logConfigName)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeLogConfigs(request)
	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	ret = response
	return
}

func (me *TkeService) DescribeServerlessNodePoolByClusterIdAndNodePoolId(ctx context.Context, clusterId,
	nodePoolId string) (instance *tke.VirtualNodePool, has bool, errRet error) {
	logId := getLogId(ctx)

	request := tke.NewDescribeClusterVirtualNodePoolsRequest()
	request.ClusterId = common.StringPtr(clusterId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterVirtualNodePools(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response != nil && len(response.Response.NodePoolSet) > 0 {
		for _, nodePool := range response.Response.NodePoolSet {
			if nodePool != nil && nodePool.NodePoolId != nil && *nodePool.NodePoolId == nodePoolId {
				has = true
				instance = nodePool
			}
		}
	}
	return
}

func (me *TkeService) CreateClusterVirtualNodePool(ctx context.Context, request *tke.CreateClusterVirtualNodePoolRequest) (id string, errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().CreateClusterVirtualNodePool(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response != nil && response.Response != nil && response.Response.NodePoolId != nil {
		id = *response.Response.NodePoolId
	}

	return
}

func (me *TkeService) DeleteClusterVirtualNodePool(ctx context.Context, request *tke.DeleteClusterVirtualNodePoolRequest) (errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DeleteClusterVirtualNodePool(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) ModifyClusterVirtualNodePool(ctx context.Context, request *tke.ModifyClusterVirtualNodePoolRequest) (errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().ModifyClusterVirtualNodePool(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TkeService) DescribeClusterVirtualNode(ctx context.Context, clusterId,
	nodePoolId string) (virtualNodes []tke.VirtualNode,
	errRet error) {
	logId := getLogId(ctx)

	request := tke.NewDescribeClusterVirtualNodeRequest()
	request.ClusterId = common.StringPtr(clusterId)
	request.NodePoolId = common.StringPtr(nodePoolId)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeClusterVirtualNode(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response.Response != nil && len(response.Response.Nodes) > 0 {
		for _, node := range response.Response.Nodes {
			if node != nil {
				virtualNodes = append(virtualNodes, *node)
			}
		}
	}
	return
}

func ModifySecurityServiceOfCvmInNodePool(ctx context.Context, d *schema.ResourceData, tkeSvc *TkeService, cvmSvc *CvmService, client *connectivity.TencentCloudClient, clusterId string, nodePoolId string) error {
	logId := getLogId(ctx)

	if d.HasChange("auto_scaling_config.0.enhanced_security_service") {
		workersInsIdOfNodePool := make([]string, 0)
		_, workers, err := tkeSvc.DescribeClusterInstances(ctx, clusterId)
		if err != nil {
			return err
		}
		for _, worker := range workers {
			if worker.NodePoolId != "" && worker.NodePoolId == nodePoolId {
				workersInsIdOfNodePool = append(workersInsIdOfNodePool, worker.InstanceId)
			}
		}

		const BatchProcessedInsLimit = 100 // limit 100 items to change each request
		var (
			launchConfigRaw []interface{}
			dMap            map[string]interface{}
		)
		if raw, ok := d.GetOk("auto_scaling_config"); ok {
			launchConfigRaw = raw.([]interface{})
			dMap = launchConfigRaw[0].(map[string]interface{})
		}

		if v, ok := dMap["enhanced_security_service"]; ok && !v.(bool) {
			// uninstall, cwp/DeleteMachine, need uuid
			// https://cloud.tencent.com/document/product/296/19844
			for i := 0; i < len(workersInsIdOfNodePool); i += BatchProcessedInsLimit {
				var reqInstanceIds []string
				if i+BatchProcessedInsLimit <= len(workersInsIdOfNodePool) {
					reqInstanceIds = workersInsIdOfNodePool[i : i+BatchProcessedInsLimit]
				} else {
					reqInstanceIds = workersInsIdOfNodePool[i:]
				}
				// get uuid
				instanceSet, err := cvmSvc.DescribeInstanceSetByIds(ctx, helper.StrListValToStr(reqInstanceIds))
				if err != nil {
					return err
				}
				// call cwp/DeleteMachine
				for _, ins := range instanceSet {
					requestDeleteMachine := cwp.NewDeleteMachineRequest()
					requestDeleteMachine.Uuid = ins.Uuid
					if _, err := client.UseCwpClient().DeleteMachine(requestDeleteMachine); err != nil {
						log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
							logId, requestDeleteMachine.GetAction(), requestDeleteMachine.ToJsonString(), err.Error())
						return err
					}
				}
			}
		} else {
			// default is true, install security agent
			// tat/InvokeCommand, CommandId=cmd-d8jj2skv, instanceId is enough
			// https://cloud.tencent.com/document/product/1340/52678
			for i := 0; i < len(workersInsIdOfNodePool); i += BatchProcessedInsLimit {
				var reqInstanceIds []string
				if i+BatchProcessedInsLimit <= len(workersInsIdOfNodePool) {
					reqInstanceIds = workersInsIdOfNodePool[i : i+BatchProcessedInsLimit]
				} else {
					reqInstanceIds = workersInsIdOfNodePool[i:]
				}
				requestInvokeCommand := tat.NewInvokeCommandRequest()
				requestInvokeCommand.InstanceIds = helper.StringsStringsPoint(reqInstanceIds)
				requestInvokeCommand.CommandId = helper.String(InstallSecurityAgentCommandId)
				requestInvokeCommand.Parameters = helper.String("{}")
				requestInvokeCommand.Timeout = helper.Uint64(60)
				_, err := client.UseTatClient().InvokeCommand(requestInvokeCommand)
				if err != nil {
					log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
						logId, requestInvokeCommand.GetAction(), requestInvokeCommand.ToJsonString(), err.Error())
					return err
				}
			}
		}
	}
	return nil
}

func ModifyClusterInternetOrIntranetAccess(ctx context.Context, d *schema.ResourceData, tkeSvc *TkeService,
	isInternet bool, enable bool, sg string, subnetId string, domain string) error {

	id := d.Id()
	var accessType string
	if isInternet {
		accessType = "cluster internet"
	} else {
		accessType = "cluster intranet"
	}
	// open access
	if enable {
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			inErr := tkeSvc.CreateClusterEndpoint(ctx, id, subnetId, sg, isInternet, domain, "", "")
			if inErr != nil {
				return retryError(inErr)
			}
			return nil
		})
		if err != nil {
			return err
		}
		err = resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
			status, message, inErr := tkeSvc.DescribeClusterEndpointStatus(ctx, id, isInternet)
			if inErr != nil {
				return retryError(inErr)
			}
			if status == TkeInternetStatusCreating {
				return resource.RetryableError(
					fmt.Errorf("%s create %s endpoint status still is %s", id, accessType, status))
			}
			if status == TkeInternetStatusNotfound || status == TkeInternetStatusCreated {
				return nil
			}
			return resource.NonRetryableError(
				fmt.Errorf("%s create %s endpoint error ,status is %s,message is %s", id, accessType, status, message))
		})
		if err != nil {
			return err
		}
	} else { // close access
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			inErr := tkeSvc.DeleteClusterEndpoint(ctx, id, isInternet)
			if inErr != nil {
				return retryError(inErr)
			}
			return nil
		})
		if err != nil {
			return err
		}
		err = resource.Retry(2*readRetryTimeout, func() *resource.RetryError {
			status, message, inErr := tkeSvc.DescribeClusterEndpointStatus(ctx, id, isInternet)
			if inErr != nil {
				return retryError(inErr)
			}
			if status == TkeInternetStatusDeleting {
				return resource.RetryableError(
					fmt.Errorf("%s close %s endpoint status still is %s", id, accessType, status))
			}
			if status == TkeInternetStatusNotfound || status == TkeInternetStatusDeleted || status == TkeInternetStatusCreated {
				return nil
			}
			return resource.NonRetryableError(
				fmt.Errorf("%s close %s endpoint error ,status is %s,message is %s", id, accessType, status, message))
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//func (me *TkeService) createBackupStorageLocation(ctx context.Context, request *tke.CreateBackupStorageLocationRequest) (errRet error) {
//	logId := getLogId(ctx)
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
//				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
//		}
//	}()
//
//	ratelimit.Check(request.GetAction())
//	response, err := me.client.UseTkeClient().CreateBackupStorageLocation(request)
//
//	if err != nil {
//		errRet = err
//		return
//	}
//
//	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
//		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
//
//	return
//}

//func (me *TkeService) describeBackupStorageLocations(ctx context.Context, names []string) (locations []*tke.BackupStorageLocation, errRet error) {
//	logId := getLogId(ctx)
//
//	request := tke.NewDescribeBackupStorageLocationsRequest()
//	request.Names = common.StringPtrs(names)
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
//				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
//		}
//	}()
//
//	ratelimit.Check(request.GetAction())
//	response, err := me.client.UseTkeClient().DescribeBackupStorageLocations(request)
//
//	if err != nil {
//		errRet = err
//		return
//	}
//
//	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
//		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
//
//	if response.Response != nil && len(response.Response.BackupStorageLocationSet) > 0 {
//		locations = append(locations, response.Response.BackupStorageLocationSet...)
//	}
//	return
//
//}

//func (me *TkeService) deleteBackupStorageLocation(ctx context.Context, name string) (errRet error) {
//	logId := getLogId(ctx)
//	request := tke.NewDeleteBackupStorageLocationRequest()
//
//	defer func() {
//		if errRet != nil {
//			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
//				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
//		}
//	}()
//	request.Name = common.StringPtr(name)
//
//	ratelimit.Check(request.GetAction())
//	_, err := me.client.UseTkeClient().DeleteBackupStorageLocation(request)
//	return err
//}

// DescribeHelmChart ...
func (me *TkeService) DescribeHelmChart(ctx context.Context, request *tke.DescribeHelmChartRequest) (
	response *tke.DescribeHelmChartResponse, errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTkeClient().DescribeHelmChart(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return

}

func (me *TkeService) DescribeClusterStatus(ctx context.Context, id string) (status *tke.ClusterStatus, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewDescribeClusterStatusRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.ClusterIds = []*string{&id}
	response, err := me.client.UseTkeClient().DescribeClusterStatus(request)
	if err != nil {
		errRet = err
		return
	}
	if *response.Response.TotalCount < 1 {
		errRet = fmt.Errorf("cluster not found")
		return
	}
	if response.Response.ClusterStatusSet == nil {
		errRet = fmt.Errorf("ClusterStatusSet is Nil")
		return
	}
	return response.Response.ClusterStatusSet[0], nil
}

func (me *TkeService) ForwardRequest(ctx context.Context, method, path, clusterName,
	requestBody string) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardRequestRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	response, err := me.client.UseTkeClient().ForwardRequest(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ClusterStatusSet is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) ForwardPlatformRequestV3(ctx context.Context, method, path, clusterName,
	requestBody string) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardPlatformRequestV3Request()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	response, err := me.client.UseTkeClient().ForwardPlatformRequestV3(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ResponseBody is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) ForwardApplicationRequestV3(ctx context.Context, method, path, clusterName,
	requestBody string) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardApplicationRequestV3Request()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	response, err := me.client.UseTkeClient().ForwardApplicationRequestV3(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ResponseBody is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) ForwardPlatformRequestV3WithOptions(ctx context.Context, method, path, clusterName,
	requestBody, accept, contentType string, encodedBody *bool) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardPlatformRequestV3Request()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if accept != "" {
		request.Accept = common.StringPtr(accept)
	}
	if contentType != "" {
		request.ContentType = common.StringPtr(contentType)
	}
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	if encodedBody != nil {
		request.EncodedBody = common.BoolPtr(*encodedBody)
	}
	response, err := me.client.UseTkeClient().ForwardPlatformRequestV3(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ResponseBody is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) ForwardRequestEncode(ctx context.Context, method, path, clusterName,
	requestBody string, encodeBody bool) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardRequestRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	request.EncodedBody = common.BoolPtr(encodeBody)
	response, err := me.client.UseTkeClient().ForwardRequest(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ClusterStatusSet is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) ForwardRequestContentType(ctx context.Context, method, path, clusterName,
	requestBody, contentType string) (responseBody string, errRet error) {
	logId := getLogId(ctx)
	request := tke.NewForwardRequestRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.Method = common.StringPtr(method)
	request.Path = common.StringPtr(path)
	if clusterName != "" {
		request.ClusterName = common.StringPtr(clusterName)
	}
	if requestBody != "" {
		request.RequestBody = common.StringPtr(requestBody)
	}
	if contentType != "" {
		request.ContentType = common.StringPtr(contentType)
	}
	response, err := me.client.UseTkeClient().ForwardRequest(request)
	if err != nil {
		errRet = err
		return
	}
	if response.Response.ResponseBody == nil {
		errRet = fmt.Errorf("ClusterStatusSet is Nil")
		return
	}
	return *response.Response.ResponseBody, nil
}

func (me *TkeService) DescribeKubernetesChartsByFilter(ctx context.Context, param map[string]interface{}) (ret []*tke.AppChart, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tke.NewGetTkeAppChartListRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "Kind" {
			request.Kind = v.(*string)
		}
		if k == "Arch" {
			request.Arch = v.(*string)
		}
		if k == "ClusterType" {
			request.ClusterType = v.(*string)
		}
	}

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseTkeClient().GetTkeAppChartList(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.AppCharts) < 1 {
		return
	}

	ret = response.Response.AppCharts
	return
}
