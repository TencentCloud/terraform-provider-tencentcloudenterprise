package tencentcloud

import (
	"context"
	"log"
	"sort"
	"time"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type CvmService struct {
	client *connectivity.TencentCloudClient
}

func (me *CvmService) DescribeInstanceById(ctx context.Context, instanceId string) (instance *cvm.Instance, errRet error) {
	logId := getLogId(ctx)
	request := cvm.NewDescribeInstancesRequest()
	request.InstanceIds = []*string{&instanceId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().DescribeInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.InstanceSet) < 1 {
		return
	}
	instance = response.Response.InstanceSet[0]
	return
}

func (me *CvmService) DescribeInstanceByFilter(ctx context.Context, instancesId []*string, filters map[string]string) (instances []*cvm.Instance, errRet error) {
	logId := getLogId(ctx)
	request := cvm.NewDescribeInstancesRequest()
	if instancesId != nil {
		request.InstanceIds = instancesId
	} else {
		request.Filters = make([]*cvm.Filter, 0, len(filters))
		for k, v := range filters {
			filter := cvm.Filter{
				Name:   helper.String(k),
				Values: []*string{helper.String(v)},
			}
			request.Filters = append(request.Filters, &filter)
		}
	}

	var offset int64 = 0
	var pageSize int64 = 100
	instances = make([]*cvm.Instance, 0)
	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCvmClient().DescribeInstances(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.InstanceSet) < 1 {
			break
		}
		instances = append(instances, response.Response.InstanceSet...)
		if len(response.Response.InstanceSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}
	return
}

func (me *CvmService) ModifyInstanceName(ctx context.Context, instanceId, instanceName string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesAttributeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.InstanceName = &instanceName

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesAttribute(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifySecurityGroups(ctx context.Context, instanceId string, securityGroups []*string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesAttributeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.SecurityGroups = securityGroups

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesAttribute(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyDisableApiTermination(ctx context.Context, instanceId string, disableApiTermination bool) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesAttributeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.DisableApiTermination = &disableApiTermination

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesAttribute(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyProjectId(ctx context.Context, instanceId string, projectId string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesProjectRequest()
	request.InstanceIds = []*string{&instanceId}
	request.ProjectId = &projectId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesProject(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyInstanceType(ctx context.Context, instanceId, instanceType string) error {
	logId := getLogId(ctx)
	request := cvm.NewResetInstancesTypeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.InstanceType = &instanceType
	request.ForceStop = helper.Bool(true)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ResetInstancesType(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyPassword(ctx context.Context, instanceId, password string) error {
	logId := getLogId(ctx)
	request := cvm.NewResetInstancesPasswordRequest()
	request.InstanceIds = []*string{&instanceId}
	request.Password = &password
	forceStop := true
	request.ForceStop = &forceStop

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ResetInstancesPassword(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyInternetMaxBandwidthOut(ctx context.Context, instanceId, internetChargeType string, internetMaxBandWidthOut int64) error {
	logId := getLogId(ctx)
	request := cvm.NewResetInstancesInternetMaxBandwidthRequest()
	request.InstanceIds = []*string{&instanceId}
	request.InternetAccessible = &cvm.InternetAccessible{
		InternetChargeType:      &internetChargeType,
		InternetMaxBandwidthOut: &internetMaxBandWidthOut,
	}

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCvmClient().ResetInstancesInternetMaxBandwidth(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	return nil
}

func (me *CvmService) ModifyVpc(ctx context.Context, instanceId, vpcId, subnetId, privateIp string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesVpcAttributeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
		VpcId:    &vpcId,
		SubnetId: &subnetId,
	}
	if privateIp != "" {
		request.VirtualPrivateCloud.PrivateIpAddresses = []*string{&privateIp}
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesVpcAttribute(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) StopInstance(ctx context.Context, instanceId string, stoppedMode string) error {
	logId := getLogId(ctx)
	request := cvm.NewStopInstancesRequest()
	request.InstanceIds = []*string{&instanceId}
	if stoppedMode != "" {
		request.StoppedMode = &stoppedMode
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().StopInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) StartInstance(ctx context.Context, instanceId string) error {
	logId := getLogId(ctx)
	request := cvm.NewStartInstancesRequest()
	request.InstanceIds = []*string{&instanceId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().StartInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) DeleteInstance(ctx context.Context, instanceId string) error {
	logId := getLogId(ctx)
	request := cvm.NewTerminateInstancesRequest()
	request.InstanceIds = []*string{&instanceId}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().TerminateInstances(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ResetInstance(ctx context.Context, request *cvm.ResetInstanceRequest) (errRet error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseCvmClient().ResetInstance(request)

	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) UnbindKeyPair(ctx context.Context, keyIds []*string, instanceIds []*string) error {
	logId := getLogId(ctx)
	request := cvm.NewDisassociateInstancesKeyPairsRequest()
	request.KeyIds = keyIds
	request.InstanceIds = instanceIds
	request.ForceStop = helper.Bool(true)

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().DisassociateInstancesKeyPairs(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) BindKeyPair(ctx context.Context, keyIds []*string, instanceId string) error {
	logId := getLogId(ctx)
	request := cvm.NewAssociateInstancesKeyPairsRequest()
	request.KeyIds = keyIds
	request.InstanceIds = []*string{&instanceId}
	request.ForceStop = helper.Bool(true)

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCvmClient().AssociateInstancesKeyPairs(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}

	return nil
}

func (me *CvmService) DescribeZones(ctx context.Context) (zones []*cvm.ZoneInfo, errRet error) {
	logId := getLogId(ctx)
	request := cvm.NewDescribeZonesRequest()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().DescribeZones(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	zones = response.Response.ZoneSet
	return
}

func flattenCvmTagsMapping(tags []*cvm.Tag) (mapping map[string]string) {
	mapping = make(map[string]string)
	for _, tag := range tags {
		mapping[*tag.Key] = *tag.Value
	}
	return
}

type cvmImages []*cvm.Image

func (a cvmImages) Len() int {
	return len(a)
}

func (me *CvmService) ModifyRenewParam(ctx context.Context, instanceId string, renewFlag string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesRenewFlagRequest()
	request.InstanceIds = []*string{&instanceId}
	request.RenewFlag = &renewFlag

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesRenewFlag(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ModifyInstanceChargeType(ctx context.Context, instanceId string, chargeType string, period int,
	renewFlag string) error {
	logId := getLogId(ctx)
	request := cvm.NewModifyInstancesChargeTypeRequest()
	request.InstanceIds = []*string{&instanceId}
	request.InstanceChargeType = &chargeType
	if chargeType == CVM_CHARGE_TYPE_PREPAID {
		request.InstanceChargePrepaid = &cvm.InstanceChargePrepaid{}
		if period != -1 {
			request.InstanceChargePrepaid.Period = helper.IntInt64(period)
		}
		if renewFlag != "" {
			request.InstanceChargePrepaid.RenewFlag = &renewFlag
		}
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ModifyInstancesChargeType(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CvmService) ResizeInstanceDisks(ctx context.Context, request *cvm.ResizeInstanceDisksRequest) (errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCvmClient().ResizeInstanceDisks(request)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (a cvmImages) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a cvmImages) Less(i, j int) bool {
	if a[i].CreatedTime == nil || a[j].CreatedTime == nil {
		return false
	}

	itime, _ := time.Parse(time.RFC3339, *a[i].CreatedTime)
	jtime, _ := time.Parse(time.RFC3339, *a[j].CreatedTime)

	return itime.Unix() < jtime.Unix()
}

// Sort images by creation date, in descending order.
func sortImages(images cvmImages) cvmImages {
	sortedImages := images
	sort.Sort(sort.Reverse(sortedImages))
	return sortedImages
}

func (me *CvmService) DescribeImagesByFilter(ctx context.Context, filters map[string][]string, instanceType string) (images []*cvm.Image, errRet error) {
	logId := getLogId(ctx)

	request := cvm.NewDescribeImagesRequest()
	request.Filters = make([]*cvm.Filter, 0, len(filters))
	for k, v := range filters {
		filter := cvm.Filter{
			Name:   helper.String(k),
			Values: []*string{},
		}
		for _, vv := range v {
			filter.Values = append(filter.Values, helper.String(vv))
		}
		request.Filters = append(request.Filters, &filter)
	}
	if instanceType != "" {
		request.InstanceType = helper.String(instanceType)
	}
	var offset uint64 = 0
	var pageSize uint64 = 100
	images = make([]*cvm.Image, 0)
	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCvmClient().DescribeImages(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil || len(response.Response.ImageSet) < 1 {
			break
		}
		images = append(images, response.Response.ImageSet...)
		if len(response.Response.ImageSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}

	return
}
