package tencentcloud

import (
	"context"
	"log"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type CfsService struct {
	client *connectivity.TencentCloudClient
}

func (me *CfsService) DescribeFileSystem(ctx context.Context, fsId, vpcId, subnetId string) (fs []*cfs.FileSystemInfo, errRet error) {
	logId := getLogId(ctx)
	request := cfs.NewDescribeCfsFileSystemsRequest()
	if fsId != "" {
		request.FileSystemId = &fsId
	}
	if vpcId != "" {
		request.VpcId = &vpcId
	}
	if subnetId != "" {
		request.SubnetId = &subnetId
	}

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCfsClient().DescribeCfsFileSystems(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	fs = response.Response.FileSystems
	return
}

func (me *CfsService) DescribeMountTargets(ctx context.Context, fsId string) (targets []*cfs.MountInfo, errRet error) {
	logId := getLogId(ctx)
	request := cfs.NewDescribeMountTargetsRequest()
	request.FileSystemId = &fsId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCfsClient().DescribeMountTargets(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	targets = response.Response.MountTargets
	return
}

func (me *CfsService) ModifyFileSystemName(ctx context.Context, fsId, fsName string) error {
	logId := getLogId(ctx)
	request := cfs.NewUpdateCfsFileSystemNameRequest()
	request.FileSystemId = &fsId
	request.FsName = &fsName

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCfsClient().UpdateCfsFileSystemName(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CfsService) ModifyFileSystemAccessGroup(ctx context.Context, fsId, accessGroupId string) error {
	logId := getLogId(ctx)
	request := cfs.NewUpdateCfsFileSystemPGroupRequest()
	request.FileSystemId = &fsId
	request.PGroupId = &accessGroupId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCfsClient().UpdateCfsFileSystemPGroup(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}

func (me *CfsService) DeleteFileSystem(ctx context.Context, fsId string) error {
	logId := getLogId(ctx)
	request := cfs.NewDeleteCfsFileSystemRequest()
	request.FileSystemId = &fsId

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCfsClient().DeleteCfsFileSystem(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return nil
}
