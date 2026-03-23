package tencentcloud

import (
	"context"
	"fmt"
	"log"

	dc "terraform-provider-tencentcloudenterprise/sdk/dc/v20180410"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type DcService struct {
	client *connectivity.TencentCloudClient
}

// ///////common
func (me *DcService) fillFilter(ins []*dc.Filter, key, value string) (outs []*dc.Filter) {
	if ins == nil {
		ins = make([]*dc.Filter, 0, 2)
	}

	var filter = dc.Filter{Name: &key, Values: []*string{&value}}
	ins = append(ins, &filter)
	outs = ins
	return
}

func (me *DcService) strPt2str(pt *string) (ret string) {
	if pt == nil {
		return
	} else {
		return *pt
	}
}

/*
func (me *DcService) intPt2int(pt *int) (ret int) {
	if pt == nil {
		return
	} else {
		return *pt
	}
}
*/

func (me *DcService) int64Pt2int64(pt *int64) (ret int64) {
	if pt == nil {
		return
	} else {
		return *pt
	}
}

func (me *DcService) DescribeDirectConnects(ctx context.Context, dcId,
	name string) (infos []dc.DirectConnect, errRet error) {

	logId := getLogId(ctx)
	request := dc.NewDescribeDirectConnectsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var offset int64 = 0
	var limit int64 = 100
	var total int64 = -1
	var has = map[string]bool{}

	var filters []*dc.Filter
	if dcId != "" {
		filters = me.fillFilter(filters, "direct-connect-id", dcId)
	}
	if name != "" {
		filters = me.fillFilter(filters, "direct-connect-name", name)
	}
	if len(filters) > 0 {
		request.Filters = filters
	}
	infos = make([]dc.DirectConnect, 0, 10)

getMoreData:
	if total >= 0 && offset >= total {
		return
	}
	request.Limit = &limit
	request.Offset = &offset
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseDcClient().DescribeDirectConnects(request)
	if err != nil {
		errRet = err
		return
	}
	if total < 0 {
		total = *response.Response.TotalCount
	}

	if len(response.Response.DirectConnectSet) > 0 {
		offset += limit
	} else {
		//get empty set,we're done
		return
	}

	for _, item := range response.Response.DirectConnectSet {
		if has[*item.DirectConnectId] {
			errRet = fmt.Errorf("get repeated dc_id[%s] when doing DescribeDirectConnects", *item.DirectConnectId)
			return
		}
		has[*item.DirectConnectId] = true
		infos = append(infos, *item)
	}
	goto getMoreData
}

func (me *DcService) DescribeDirectConnectTunnel(ctx context.Context, dcxId string) (info dc.DirectConnectTunnel, has int64, errRet error) {

	infos, err := me.DescribeDirectConnectTunnels(ctx, dcxId, "")

	if err != nil {
		errRet = err
		return
	}
	has = int64(len(infos))

	if has > 0 {
		info = infos[0]

	}
	return
}

func (me *DcService) DescribeDirectConnectTunnels(ctx context.Context, dcxId,
	name string) (infos []dc.DirectConnectTunnel, errRet error) {

	logId := getLogId(ctx)
	request := dc.NewDescribeDirectConnectTunnelsRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	var offset int64 = 0
	var limit int64 = 100
	var total int64 = -1
	var has = map[string]bool{}

	var filters []*dc.Filter
	if dcxId != "" {
		filters = me.fillFilter(filters, "direct-connect-tunnel-id", dcxId)
	}
	if name != "" {
		filters = me.fillFilter(filters, "direct-connect-tunnel-name", name)
	}
	if len(filters) > 0 {
		request.Filters = filters
	}
	infos = make([]dc.DirectConnectTunnel, 0, 10)
getMoreData:
	if total >= 0 && offset >= total {
		return
	}
	request.Limit = &limit
	request.Offset = &offset
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseDcClient().DescribeDirectConnectTunnels(request)
	if err != nil {
		errRet = err
		return
	}
	if total < 0 {
		total = *response.Response.TotalCount
	}

	if len(response.Response.DirectConnectTunnelSet) > 0 {
		offset += limit
	} else {
		//get empty set,we're done
		return
	}
	for _, item := range response.Response.DirectConnectTunnelSet {
		if has[*item.DirectConnectTunnelId] {
			errRet = fmt.Errorf("get repeated dcx_id[%s] when doing DescribeDirectConnectTunnels", *item.DirectConnectTunnelId)
			return
		}
		has[*item.DirectConnectTunnelId] = true
		infos = append(infos, *item)
	}
	goto getMoreData
}

func (me *DcService) CreateDirectConnectTunnel(ctx context.Context, dcId, dcxName, networkType,
	networkRegion, vpcName, routeType, bgpAuthKey, cloudAddress, customerAddress, dcgId, loadMode, relatedDirectConnectTunnelId, ipType, idcRoutes string,
	bgpAsn, vlan, bandwidth, bfdInterval, vpcId int64, connectSubnetMask uint64, enableBfd, bgpPeerExist bool, ownerAccount string) (dcxId string, errRet error) {

	logId := getLogId(ctx)
	request := dc.NewCreateDirectConnectTunnelRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	// 设置基本参数
	request.DirectConnectId = &dcId
	request.DirectConnectTunnelName = &dcxName
	request.NetworkRegion = &networkRegion
	request.RouteType = &routeType
	request.DirectConnectGatewayId = &dcgId

	// 设置专线拥有者账户
	if ownerAccount != "" {
		request.DirectConnectOwnerAccount = &ownerAccount
	}

	// 设置VPC相关参数
	request.VpcId = &vpcId
	if vpcName != "" {
		request.VpcName = &vpcName
	}

	// 设置网络参数
	if bandwidth >= 0 {
		request.Bandwidth = &bandwidth
	}
	request.Vlan = &vlan

	// 设置地址参数
	if cloudAddress != "" {
		request.CloudAddress = &cloudAddress
	}
	if customerAddress != "" {
		request.CustomerAddress = &customerAddress
	}

	request.ConnectSubnetMask = &connectSubnetMask

	// 设置负载均衡模式
	if loadMode != "" {
		request.LoadMode = &loadMode
	}

	// 设置关联的冗余通道ID
	if relatedDirectConnectTunnelId != "" {
		request.RelatedDirectConnectTunnelId = &relatedDirectConnectTunnelId
	}

	// 设置IP类型
	if ipType != "" {
		request.IpType = &ipType
	}

	// 设置BFD参数
	request.EnableBfd = &enableBfd
	if bfdInterval > 0 {
		if !enableBfd {
			errRet = fmt.Errorf("bfd_interval can only be set when enable_bfd is true")
			return
		}
		request.BfdInterval = &bfdInterval
	}

	// 设置BGP参数
	if bgpPeerExist {
		var peer dc.BgpPeer
		peer.Asn = &bgpAsn
		peer.AuthKey = &bgpAuthKey
		request.BgpPeer = &peer
	}

	// 设置IDC路由（静态路由）
	if idcRoutes != "" {
		request.IdcRoutes = &idcRoutes
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseDcClient().CreateDirectConnectTunnel(request)
	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.DirectConnectTunnelIdSet) != 1 {
		errRet = fmt.Errorf("CreateDirectConnectTunnel  return %d DirectConnectTunnelIdSet",
			len(response.Response.DirectConnectTunnelIdSet))
		return
	}
	dcxId = *response.Response.DirectConnectTunnelIdSet[0]
	return
}

func (me *DcService) DeleteDirectConnectTunnel(ctx context.Context, dcxId string) (errRet error) {

	logId := getLogId(ctx)
	request := dc.NewDeleteDirectConnectTunnelRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.DirectConnectTunnelId = &dcxId
	ratelimit.Check(request.GetAction())
	_, err := me.client.UseDcClient().DeleteDirectConnectTunnel(request)
	if err != nil {
		errRet = err
	}
	return
}

func (me *DcService) ModifyDirectConnectTunnelAttribute(ctx context.Context, dcxId,
	name, bgpAuthKey, idcRoutes string, bandwidth, bgpAsn, bfdInterval int64,
	enableBfd bool, enableMulticast bool, multicastGroups, vpcId string) (errRet error) {

	logId := getLogId(ctx)
	request := dc.NewModifyDirectConnectTunnelAttributeRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.DirectConnectTunnelId = &dcxId
	if name != "" {
		request.DirectConnectTunnelName = &name
	}

	if bgpAsn >= 0 {
		var peer dc.BgpPeer
		peer.Asn = &bgpAsn
		peer.AuthKey = &bgpAuthKey
		request.BgpPeer = &peer
	}

	if bandwidth > 0 {
		request.Bandwidth = &bandwidth
	}

	// 设置BFD参数
	request.EnableBfd = &enableBfd
	if bfdInterval > 0 {
		if !enableBfd {
			errRet = fmt.Errorf("bfd_interval can only be set when enable_bfd is true")
			return
		}
		request.BfdInterval = &bfdInterval
	}

	// 设置IDC路由参数
	if idcRoutes != "" {
		request.IdcRoutes = &idcRoutes
	}

	
	// 如果要开启专线通道组播，需要先检查VPC组播状态
	if enableMulticast {
		// 检查VPC是否已开启组播
		vpcMulticastEnabled, err := me.CheckVpcMulticastEnabled(ctx, vpcId)
		if err != nil {
			errRet = fmt.Errorf("failed to check VPC multicast status: %v", err)
			return
		}
		
		if !vpcMulticastEnabled {
			errRet = fmt.Errorf("VPC %s multicast is not enabled. Please enable multicast in VPC resource first before enabling tunnel multicast", vpcId)
			return
		}
	}

	// 设置组播参数
	request.EnableMulticast = &enableMulticast
	if multicastGroups != "" {
		request.MulticastGroups = &multicastGroups
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseDcClient().ModifyDirectConnectTunnelAttribute(request)
	if err != nil {
		errRet = err
	}
	return
}

// CheckVpcMulticastEnabled 检查VPC是否开启组播功能
func (me *DcService) CheckVpcMulticastEnabled(ctx context.Context, vpcId string) (enabled bool, errRet error) {
	logId := getLogId(ctx)
	request := vpc.NewDescribeVpcsRequest()
	request.VpcIds = []*string{&vpcId}
	
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseVpcClient().DescribeVpcs(request)
	if err != nil {
		errRet = err
		return
	}

	if len(response.Response.VpcSet) == 0 {
		errRet = fmt.Errorf("VPC %s not found", vpcId)
		return
	}

	vpcInfo := response.Response.VpcSet[0]
	if vpcInfo.EnableMulticast != nil {
		enabled = *vpcInfo.EnableMulticast
	}
	
	return
}

func (me *DcService) DescribeDcShareDcxConfigById(ctx context.Context, directConnectTunnelId string) (ShareDcxConfig *dc.DirectConnectTunnel, errRet error) {
	logId := getLogId(ctx)

	request := dc.NewDescribeDirectConnectTunnelsRequest()
	//request.DirectConnectTunnelIds = []*string{&directConnectTunnelId}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DescribeDirectConnectTunnels(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.DirectConnectTunnelSet) < 1 {
		return
	}

	ShareDcxConfig = response.Response.DirectConnectTunnelSet[0]
	return
}

/*
func (me *DcService) DescribeDcInternetAddressById(ctx context.Context, instanceId string) (internetAddress *dc.InternetAddressDetail, errRet error) {
	logId := getLogId(ctx)

	request := dc.NewDescribeInternetAddressRequest()

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DescribeInternetAddress(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	for _, address := range response.Response.Subnets {
		if *address.InstanceId == instanceId {
			internetAddress = address
			break
		}
	}
	return
}
*/

/*
func (me *DcService) DeleteDcInternetAddressById(ctx context.Context, instanceId string) (errRet error) {
	logId := getLogId(ctx)

	request := dc.NewReleaseInternetAddressRequest()
	request.InstanceId = &instanceId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().ReleaseInternetAddress(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

*/

/*
func (me *DcService) DescribeDcxExtraConfigById(ctx context.Context, directConnectTunnelId string) (dcxExtraConfig *dc.DirectConnectTunnelExtra, errRet error) {
	logId := getLogId(ctx)

	request := dc.NewDescribeDirectConnectTunnelExtraRequest()
	request.DirectConnectTunnelId = &directConnectTunnelId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DescribeDirectConnectTunnelExtra(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	dcxExtraConfig = response.Response.DirectConnectTunnelExtra
	return
}
*/

/*
func (me *DcService) DescribeDcInternetAddressQuota(ctx context.Context) (InternetAddressQuota *dc.DescribeInternetAddressQuotaResponse, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = dc.NewDescribeInternetAddressQuotaRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DescribeInternetAddressQuota(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	InternetAddressQuota = response

	return
}
*/

/*
func (me *DcService) DescribeDcInternetAddressStatistics(ctx context.Context) (statistics []*dc.InternetAddressStatistics, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = dc.NewDescribeInternetAddressStatisticsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DescribeInternetAddressStatistics(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	statistics = response.Response.InternetAddressStatistics
	return
}
*/

/*
func (me *DcService) DescribeDcPublicDirectConnectTunnelRoutesByFilter(ctx context.Context, param map[string]interface{}) (PublicDirectConnectTunnelRoutes []*dc.DirectConnectTunnelRoute, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = dc.NewDescribePublicDirectConnectTunnelRoutesRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	for k, v := range param {
		if k == "DirectConnectTunnelId" {
			request.DirectConnectTunnelId = v.(*string)
		}
		if k == "Filters" {
			request.Filters = v.([]*dc.Filter)
		}
	}

	ratelimit.Check(request.GetAction())

	var (
		offset int64 = 0
		limit  int64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseDcClient().DescribePublicDirectConnectTunnelRoutes(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.Routes) < 1 {
			break
		}
		PublicDirectConnectTunnelRoutes = append(PublicDirectConnectTunnelRoutes, response.Response.Routes...)
		if len(response.Response.Routes) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

*/

func (me *DcService) DeleteDcInstanceById(ctx context.Context, directConnectId string) (errRet error) {
	logId := getLogId(ctx)

	request := dc.NewDeleteDirectConnectRequest()
	request.DirectConnectId = &directConnectId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())

	response, err := me.client.UseDcClient().DeleteDirectConnect(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *DcService) DescribeDcAccessPointsByFilter(ctx context.Context, param map[string]interface{}) (AccessPoints []*dc.AccessPoint, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = dc.NewDescribeAccessPointsRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	//for k, v := range param {
	//	if k == "RegionId" {
	//		request.RegionId = v.(*string)
	//	}
	//}

	ratelimit.Check(request.GetAction())

	var (
		offset int64 = 0
		limit  int64 = 20
	)
	for {
		request.Offset = &offset
		request.Limit = &limit
		response, err := me.client.UseDcClient().DescribeAccessPoints(request)
		if err != nil {
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.AccessPointSet) < 1 {
			break
		}
		AccessPoints = append(AccessPoints, response.Response.AccessPointSet...)
		if len(response.Response.AccessPointSet) < int(limit) {
			break
		}

		offset += limit
	}

	return
}

// UpdateVifAssociated 更新专线通道的负载均衡模式和关联通道ID
func (me *DcService) UpdateVifAssociated(ctx context.Context, dcxId, loadMode, relatedDirectConnectTunnelId string) (errRet error) {
	logId := getLogId(ctx)
	request := dc.NewUpdateVifAssociatedRequest()
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	request.DirectConnectTunnelId = &dcxId
	
	if loadMode != "" {
		request.LoadMode = &loadMode
	}
	
	if relatedDirectConnectTunnelId != "" {
		request.RelatedDirectConnectTunnelId = &relatedDirectConnectTunnelId
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseDcClient().UpdateVifAssociated(request)
	if err != nil {
		errRet = err
	}
	return
}


