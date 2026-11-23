package tencentcloud

import (
	"context"
	"log"
	cbs "terraform-provider-tencentcloudenterprise/sdk/cbs/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type CbsService struct {
	client *connectivity.TencentCloudClient
}

func (me *CbsService) DescribeDiskById(ctx context.Context, diskId string) (disk *cbs.Disk, errRet error) {
	disks, err := me.DescribeDiskList(ctx, []*string{&diskId})
	if err != nil {
		errRet = err
		return
	}
	if len(disks) > 0 {
		disk = disks[0]
	}
	return
}

func (me *CbsService) DescribeDiskList(ctx context.Context, diskIds []*string) (disk []*cbs.Disk, errRet error) {
	logId := getLogId(ctx)
	request := cbs.NewDescribeDisksRequest()
	request.DiskIds = diskIds
	request.Limit = helper.IntUint64(100)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().DescribeDisks(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if len(response.Response.DiskSet) > 0 {
		disk = response.Response.DiskSet
	}
	return
}

func (me *CbsService) DescribeDisksByFilter(ctx context.Context, params map[string]interface{}) (disks []*cbs.Disk, errRet error) {
	logId := getLogId(ctx)
	request := cbs.NewDescribeDisksRequest()
	request.Filters = make([]*cbs.Filter, 0, len(params))
	for k, v := range params {
		filter := &cbs.Filter{
			Name: helper.String(k),
		}
		switch v := v.(type) {
		case string:
			filter.Values = []*string{helper.String(v)}
		case []*string:
			filter.Values = v
		}
		request.Filters = append(request.Filters, filter)
	}

	offset := 0
	pageSize := 100
	disks = make([]*cbs.Disk, 0)
	for {
		request.Offset = helper.IntUint64(offset)
		request.Limit = helper.IntUint64(pageSize)
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCbsClient().DescribeDisks(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.DiskSet) < 1 {
			break
		}

		disks = append(disks, response.Response.DiskSet...)

		if len(response.Response.DiskSet) < pageSize {
			break
		}
		offset += pageSize
	}
	return
}

func (me *CbsService) ModifyDiskAttributes(ctx context.Context, diskId, diskName string, projectId int) error {
	logId := getLogId(ctx)
	request := cbs.NewModifyDiskAttributesRequest()
	request.DiskIds = []*string{&diskId}
	if diskName != "" {
		request.DiskName = &diskName
	}
	if projectId >= 0 {
		request.ProjectId = helper.IntUint64(projectId)
	}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().ModifyDiskAttributes(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CbsService) DeleteDiskById(ctx context.Context, diskId string) error {
	logId := getLogId(ctx)
	request := cbs.NewTerminateDisksRequest()
	request.DiskIds = []*string{&diskId}
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().TerminateDisks(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CbsService) ResizeDisk(ctx context.Context, diskId string, diskSize int) error {
	logId := getLogId(ctx)
	request := cbs.NewResizeDiskRequest()
	request.DiskId = &diskId
	request.DiskSize = helper.IntUint64(diskSize)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().ResizeDisk(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CbsService) ModifyThroughputPerformance(ctx context.Context, diskId string, throughputPerformance int) error {
	logId := getLogId(ctx)
	request := cbs.NewModifyDiskExtraPerformanceRequest()
	request.DiskId = &diskId
	request.ThroughputPerformance = helper.IntUint64(throughputPerformance)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().ModifyDiskExtraPerformance(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CbsService) ApplySnapshot(ctx context.Context, diskId, snapshotId string) error {
	logId := getLogId(ctx)
	request := cbs.NewApplySnapshotRequest()
	request.DiskId = &diskId
	request.SnapshotId = &snapshotId
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCbsClient().ApplySnapshot(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
	return nil
}

func (me *CbsService) ModifyDisksRenewFlag(ctx context.Context, storageId string, renewFlag string) error {
	logId := getLogId(ctx)
	request := cbs.NewModifyDisksRenewFlagRequest()
	request.DiskIds = []*string{&storageId}
	request.RenewFlag = &renewFlag

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCbsClient().ModifyDisksRenewFlag(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	return nil
}

func (me *CbsService) DescribeSnapshotByIds(ctx context.Context, snapshotIdsParam []*string) (snapshots []*cbs.Snapshot, errRet error) {
	if len(snapshotIdsParam) == 0 {
		return
	}

	var (
		logId            = getLogId(ctx)
		request          = cbs.NewDescribeSnapshotsRequest()
		err              error
		response         *cbs.DescribeSnapshotsResponse
		offset, pageSize uint64 = 0, 100
	)
	request.SnapshotIds = snapshotIdsParam

	for {
		request.Offset = &offset
		request.Limit = &pageSize

		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())
			response, err = me.client.UseCbsClient().DescribeSnapshots(request)
			if err != nil {
				return retryError(err, InternalError)
			}
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}

		if response == nil || response.Response == nil || len(response.Response.SnapshotSet) == 0 {
			break
		}
		snapshots = append(snapshots, response.Response.SnapshotSet...)
		if len(response.Response.SnapshotSet) < int(pageSize) {
			break
		}
		offset += pageSize
	}
	return
}
