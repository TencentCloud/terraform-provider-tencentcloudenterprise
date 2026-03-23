package tencentcloud

import (
	"context"
	"log"

	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"
)

type TagService struct {
	client *connectivity.TencentCloudClient
}

// DiffTags 比较两个标签集合的差异
func (me *TagService) DiffTags(oldTags, newTags map[string]interface{}) (replaceTags map[string]string, deleteTags []string) {
	replaceTags = make(map[string]string)
	deleteTags = make([]string, 0)
	for k, v := range newTags {
		_, ok := oldTags[k]
		if !ok || oldTags[k].(string) != v.(string) {
			replaceTags[k] = v.(string)
		}
	}
	for k := range oldTags {
		_, ok := newTags[k]
		if !ok {
			deleteTags = append(deleteTags, k)
		}
	}
	return
}

func (me *TagService) ModifyTags(ctx context.Context, resourceName string, replaceTags map[string]string, deleteKeys []string) error {
	request := tag.NewModifyResourceTagsRequest()
	request.Resource = &resourceName
	if len(replaceTags) > 0 {
		request.ReplaceTags = make([]*tag.Tag, 0, len(replaceTags))
		for k, v := range replaceTags {
			key := k
			value := v
			replaceTag := &tag.Tag{
				TagKey:   &key,
				TagValue: &value,
			}
			request.ReplaceTags = append(request.ReplaceTags, replaceTag)
		}
	}
	if len(deleteKeys) > 0 {
		request.DeleteTags = make([]*tag.TagKeyObject, 0, len(deleteKeys))
		for _, v := range deleteKeys {
			key := v
			deleteKey := &tag.TagKeyObject{
				TagKey: &key,
			}
			request.DeleteTags = append(request.DeleteTags, deleteKey)
		}
	}

	return resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())

		if _, err := me.client.UseTagClient().ModifyResourceTags(request); err != nil {
			return retryError(errors.WithStack(err))
		}

		return nil
	})
}

func (me *TagService) DescribeResourceTags(ctx context.Context, serviceType, resourceType, region, resourceId string) (tags map[string]string, err error) {
	request := tag.NewDescribeResourceTagsByResourceIdsRequest()
	request.ServiceType = &serviceType
	request.ResourcePrefix = &resourceType
	request.ResourceRegion = &region
	request.ResourceIds = []*string{&resourceId}
	request.Limit = helper.IntUint64(DESCRIBE_TAGS_LIMIT)

	var offset uint64
	request.Offset = &offset

	// for run loop at least once
	count := DESCRIBE_TAGS_LIMIT
	for count == DESCRIBE_TAGS_LIMIT {
		if err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			ratelimit.Check(request.GetAction())

			response, err := me.client.UseTagClient().DescribeResourceTagsByResourceIds(request)
			if err != nil {
				count = 0

				return retryError(errors.WithStack(err))
			}

			allTags := response.Response.Tags
			count = len(allTags)

			for _, t := range allTags {
				if *t.ResourceId != resourceId {
					continue
				}
				if tags == nil {
					tags = make(map[string]string)
				}

				tags[*t.TagKey] = *t.TagValue
			}

			return nil
		}); err != nil {
			return nil, err
		}

		offset += uint64(count)
	}

	return
}

func diffTags(oldTags, newTags map[string]interface{}) (replaceTags map[string]string, deleteTags []string) {
	replaceTags = make(map[string]string)
	deleteTags = make([]string, 0)
	for k, v := range newTags {
		_, ok := oldTags[k]
		if !ok || oldTags[k].(string) != v.(string) {
			replaceTags[k] = v.(string)
		}
	}
	for k := range oldTags {
		_, ok := newTags[k]
		if !ok {
			deleteTags = append(deleteTags, k)
		}
	}
	return
}

// DescribeTagById describes a tag by key and value using DescribeTags API
func (me *TagService) DescribeTagById(ctx context.Context, tagKey string, tagValue string) (tagRes *tag.TagWithDelete, errRet error) {
	logId := getLogId(ctx)

	request := tag.NewDescribeTagsRequest()
	request.TagKey = &tagKey
	request.TagValue = &tagValue

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTagClient().DescribeTags(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil || len(response.Response.Tags) == 0 {
		return
	}

	for _, v := range response.Response.Tags {
		if v != nil && v.TagKey != nil && v.TagValue != nil {
			if *v.TagKey == tagKey && *v.TagValue == tagValue {
				tagRes = v
				break
			}
		}
	}

	return
}

// DeleteTagById deletes a tag by key and value
func (me *TagService) DeleteTagById(ctx context.Context, tagKey string, tagValue string) (errRet error) {
	logId := getLogId(ctx)

	request := tag.NewDeleteTagRequest()
	request.TagKey = &tagKey
	request.TagValue = &tagValue

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTagClient().DeleteTag(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}

// DescribeTagAttachmentById describes a tag attachment using GetResources API
func (me *TagService) DescribeTagAttachmentById(ctx context.Context, tagKey string, tagValue string, resourceId string) (found bool, errRet error) {
	logId := getLogId(ctx)

	request := tag.NewGetResourcesRequest()
	request.ResourceList = []*string{&resourceId}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTagClient().GetResources(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	if response == nil || response.Response == nil || len(response.Response.ResourceTagMappingList) == 0 {
		return
	}

	for _, resourceTagMap := range response.Response.ResourceTagMappingList {
		if resourceTagMap.Resource != nil && *resourceTagMap.Resource == resourceId {
			for _, t := range resourceTagMap.Tags {
				if t.TagKey != nil && t.TagValue != nil && *t.TagKey == tagKey && *t.TagValue == tagValue {
					found = true
					return
				}
			}
		}
	}

	return
}

// DeleteTagAttachmentById deletes a tag attachment using ModifyResourceTags API
func (me *TagService) DeleteTagAttachmentById(ctx context.Context, tagKey string, resourceId string) (errRet error) {
	logId := getLogId(ctx)

	request := tag.NewModifyResourceTagsRequest()
	request.Resource = &resourceId
	request.DeleteTags = []*tag.TagKeyObject{
		{
			TagKey: &tagKey,
		},
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseTagClient().ModifyResourceTags(request)
	if err != nil {
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

	return
}
