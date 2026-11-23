package tencentcloud

import (
	"context"
	"log"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

type TdmqRocketmqService struct {
	client *connectivity.TencentCloudClient
}

func (me *TdmqRocketmqService) DescribeTdmqRocketmqCluster(ctx context.Context, clusterId string) (cluster *tdmqRocketmq.RocketMQClusterInfo, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tdmqRocketmq.NewDescribeRocketMQClusterRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()
	request.ClusterId = &clusterId

	response, err := me.client.UseTdmqClient().DescribeRocketMQCluster(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	cluster = response.Response.ClusterInfo
	return
}

func (me *TdmqRocketmqService) DeleteTdmqRocketmqClusterById(ctx context.Context, clusterId string) (errRet error) {
	logId := getLogId(ctx)

	request := tdmqRocketmq.NewDeleteRocketMQClusterRequest()

	request.ClusterId = &clusterId

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTdmqClient().DeleteRocketMQCluster(request)
	if err != nil {
		errRet = err
		return err
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

func (me *TdmqRocketmqService) DescribeRocketmqClusterByFilter(ctx context.Context, param map[string]interface{}) (cluster []*tdmqRocketmq.RocketMQClusterDetail, errRet error) {
	var (
		logId   = getLogId(ctx)
		request = tdmqRocketmq.NewDescribeRocketMQClustersRequest()
	)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "query object", request.ToJsonString(), errRet.Error())
		}
	}()

	if v, ok := param["id_keyword"]; ok {
		request.IdKeyword = helper.String(v.(string))
	}
	if v, ok := param["name_keyword"]; ok {
		request.NameKeyword = helper.String(v.(string))
	}
	if v, ok := param["cluster_id_list"]; ok {
		request.ClusterIdList = make([]*string, 0)
		for _, cluster := range v.([]interface{}) {
			clusterId := cluster.(string)
			request.ClusterIdList = append(request.ClusterIdList, &clusterId)
		}
	}

	ratelimit.Check(request.GetAction())

	var offset uint64 = 0
	var pageSize uint64 = 20

	for {
		request.Offset = &offset
		request.Limit = &pageSize
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseTdmqClient().DescribeRocketMQClusters(request)
		if err != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), err.Error())
			errRet = err
			return
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || len(response.Response.ClusterList) < 1 {
			break
		}
		cluster = append(cluster, response.Response.ClusterList...)
		if len(response.Response.ClusterList) < int(pageSize) {
			break
		}
		offset += pageSize
	}
	return
}
