package tencentcloud

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pkg/errors"

	"github.com/tencentyun/cos-go-sdk-v5"

	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CosService struct {
	client       *connectivity.TencentCloudClient
	useCspClient bool
}

type CosBucketDomainCertItem struct {
	bucket     string
	domainName string
}

const (
	CERT_ENABLED  = "Enabled"
	CERT_DISABLED = "Disabled"
)

const PUBLIC_GRANTEE = "http://cam.qcloud.com/groups/global/AllUsers"

// GetClient - get csp or cos client
func (c *CosService) GetClient(bucket string) *cos.Client {
	if c.useCspClient {
		return c.client.UseTencentCspClient(bucket)
	}
	return c.client.UseTencentCosClient(bucket)
}

func (me *CosService) HeadObject(ctx context.Context, bucket, key string) (info *s3.HeadObjectOutput, errRet error) {
	logId := getLogId(ctx)

	request := s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	ratelimit.Check("HeadObject")
	response, err := me.client.UseCosS3Client(me.useCspClient).HeadObject(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "head object", request.String(), err.Error())
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "head object", request.String(), response.String())

	return response, nil
}

func (me *CosService) DeleteObject(ctx context.Context, bucket, key string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "delete object", request.String(), errRet.Error())
		}
	}()
	ratelimit.Check("DeleteObject")
	response, err := me.client.UseCosS3Client(me.useCspClient).DeleteObject(&request)
	if err != nil {
		errRet = fmt.Errorf("cos delete object error: %s, bucket: %s, object: %s", err.Error(), bucket, key)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "delete object", request.String(), response.String())

	return nil
}

func (me *CosService) PutObjectAcl(ctx context.Context, bucket, key, acl string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.PutObjectAclInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		ACL:    aws.String(acl),
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "put object acl", request.String(), errRet.Error())
		}
	}()
	ratelimit.Check("PutObjectAcl")
	response, err := me.client.UseCosS3Client(me.useCspClient).PutObjectAcl(&request)
	if err != nil {
		errRet = fmt.Errorf("cos put object acl error: %s, bucket: %s, object: %s", err.Error(), bucket, key)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "put object acl", request.String(), response.String())

	return nil
}

// PutBucket - base on aws s3
func (me *CosService) PutBucket(ctx context.Context, bucket, acl string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String(acl),
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "put bucket", request.String(), errRet.Error())
		}
	}()
	ratelimit.Check("CreateBucket")
	client := me.client.UseCosS3Client(me.useCspClient)
	response, err := client.CreateBucket(&request)
	if err != nil {
		errRet = fmt.Errorf("cos put bucket error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s], endpoint %s\n",
		logId, "put bucket", request.String(), response.String(), client.Endpoint)

	return nil
}

// TencentCosPutBucket - To support MAZ config, We use tencentcloud cos sdk instead of aws s3
func (c *CosService) TencentCosPutBucket(ctx context.Context, bucket string, opt *cos.BucketPutOptions) (errRet error) {
	logId := getLogId(ctx)

	req, _ := json.Marshal(opt)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request [%s], reason[%s]\n",
				logId, "put bucket", req, errRet.Error())
		}
	}()

	ratelimit.Check("TencentcloudCosPutBucket")
	client := c.GetClient(bucket)
	response, err := client.Bucket.Put(ctx, opt)

	if err != nil {
		errRet = fmt.Errorf("cos put bucket error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	resp, _ := json.Marshal(response.Response.Body)

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s], baseUrl %s\n",
		logId, "put bucket", req, resp, client.BaseURL.BucketURL)

	return nil
}

func (c *CosService) TencentCosPutBucketACL(
	ctx context.Context,
	bucket string,
	reqBody string,
	header string,
) (errRet error) {
	logId := getLogId(ctx)

	acl := &cos.ACLXml{}

	opt := &cos.BucketPutACLOptions{}
	if reqBody != "" {
		err := xml.Unmarshal([]byte(reqBody), acl)

		if err != nil {
			errRet = fmt.Errorf("cos [PutBucketACL] XML Unmarshal error: %s, bucket: %s", err.Error(), bucket)
			return
		}
		opt.Body = acl
	} else if header != "" {
		opt.Header = &cos.ACLHeaderOptions{
			XCosACL: header,
		}
	}

	req, _ := json.Marshal(opt)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request [%s], reason[%s]\n",
				logId, "PutBucketACL", req, errRet.Error())
		}
	}()

	ratelimit.Check("TencentcloudCosPutBucketACL")
	response, err := c.GetClient(bucket).Bucket.PutACL(ctx, opt)

	if err != nil {
		errRet = fmt.Errorf("cos [PutBucketACL] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	resp, _ := json.Marshal(response.Response.Body)

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "PutBucketACL", req, resp)

	return nil
}

func (me *CosService) HeadBucket(ctx context.Context, bucket string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("HeadBucket")
	response, err := me.client.UseCosS3Client(me.useCspClient).HeadBucket(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "head bucket", request.String(), err.Error())
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "head bucket", request.String(), response.String())

	return nil
}

func (me *CosService) TencentcloudHeadBucket(ctx context.Context, bucket string) (code int, header http.Header, errRet error) {
	logId := getLogId(ctx)

	// 使用统一的S3 SDK，确保使用服务域名
	request := s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	}
	
	ratelimit.Check("HeadBucket")
	_, err := me.client.UseCosS3Client(me.useCspClient).HeadBucket(&request)

	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
			logId, "HeadBucket", err.Error())
		errRet = err
		// 对于S3 SDK，我们需要从错误中提取状态码
		if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == "NotFound" {
				code = 404
			} else {
				code = 500
			}
		} else {
			code = 500
		}
		return
	}

	// 成功的情况
	code = 200
	header = make(http.Header)
	log.Printf("[DEBUG]%s api[%s] success\n",
		logId, "HeadBucket")

	return
}

func (me *CosService) DeleteBucket(ctx context.Context, bucket string, forced bool, versioned bool) (errRet error) {
	logId := getLogId(ctx)

	if forced {
		log.Printf("[DEBUG]%s api[%s] triggered, bucket [%s], versioned [%v]\n",
			logId, "ForceCleanObject", bucket, versioned)
		err := me.ForceCleanObject(ctx, bucket, versioned)
		if err != nil {
			return err
		}
	}

	request := s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("DeleteBucket")
	response, err := me.client.UseCosS3Client(me.useCspClient).DeleteBucket(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "delete bucket", request.String(), err.Error())
		return fmt.Errorf("cos delete bucket error: %s, bucket: %s", err.Error(), bucket)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "delete bucket", request.String(), response.String())

	return nil
}

func (c *CosService) ForceCleanObject(ctx context.Context, bucket string, versioned bool) error {
	logId := getLogId(ctx)

	// Get the object list of bucket with all versions
	verOpt := cos.BucketGetObjectVersionsOptions{}
	objList, resp, err := c.GetClient(bucket).Bucket.GetObjectVersions(ctx, &verOpt)

	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, resp body [%s], reason[%s]\n",
			logId, "GetObjectVersions", resp.Body, err.Error())
		return fmt.Errorf("cos force clean object error: %s, bucket: %s", err.Error(), bucket)
	}
	if objList.IsTruncated {
		return fmt.Errorf("cos force clean object error: the list of objects is truncated and the bucket[%s] needs to be deleted manually!!!", bucket)
	}

	verCnt := len(objList.Version)
	markerCnt := len(objList.DeleteMarker)
	log.Printf("[DEBUG][ForceCleanObject]%s api[%s] success, get [%v] versions of object, get [%v] deleteMarker, versioned[%v].\n", logId, "GetObjectVersions", verCnt, markerCnt, versioned)

	delCnt := verCnt + markerCnt
	if delCnt == 0 {
		return nil
	}

	delObjs := make([]cos.Object, 0, delCnt)
	if versioned {
		//add the versions
		for _, v := range objList.Version {
			delObjs = append(delObjs, cos.Object{
				Key:       v.Key,
				VersionId: v.VersionId,
			})
		}
		// add the delete-marker
		for _, m := range objList.DeleteMarker {
			delObjs = append(delObjs, cos.Object{
				Key:       m.Key,
				VersionId: m.VersionId,
			})
		}
	} else {
		for _, v := range objList.Version {
			delObjs = append(delObjs, cos.Object{
				Key: v.Key,
			})
		}
	}

	opt := cos.ObjectDeleteMultiOptions{
		Quiet:   true,
		Objects: delObjs,
	}

	// Multi-delete by specified object.
	result, resp, err := c.GetClient(bucket).Object.DeleteMulti(ctx, &opt)

	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, resp body [%s], reason[%s], opt[%v]\n",
			logId, "DeleteMulti", resp.Body, err.Error(), opt)
		return fmt.Errorf("cos force clean object error: %s, bucket: %s", err.Error(), bucket)
	}
	log.Printf("[DEBUG][ForceCleanObject]%s api[%s] completed, removed [%v] versions of object. [%v] failed to remove.\n",
		logId, "DeleteMulti", len(result.DeletedObjects), len(result.Errors))

	// Clean the failed removal version.
	if len(result.Errors) > 0 {
		log.Printf("[CRITAL]%s api[%s] it still [%v] objects have not been removed, need try DeleteMulti again.\n",
			logId, "DeleteMulti", len(result.Errors))

		if err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			unDelObjs := make([]cos.Object, 0, len(result.Errors))
			for _, v := range result.Errors {
				unDelObjs = append(unDelObjs, cos.Object{
					Key:       v.Key,
					VersionId: v.VersionId,
				})
			}
			unDelOpt := cos.ObjectDeleteMultiOptions{
				Quiet:   true,
				Objects: unDelObjs,
			}

			result, resp, err := c.GetClient(bucket).Object.DeleteMulti(ctx, &unDelOpt)
			if err != nil {
				log.Printf("[CRITAL][retry]%s api[%s] fail, resp body [%s], reason[%s]\n",
					logId, "DeleteMulti ", resp.Body, err.Error())
				return retryError(err, InternalError)
			}
			if len(result.Errors) > 0 {
				return resource.RetryableError(fmt.Errorf("[CRITAL][retry]%s api[%s] it still %v objects have not been removed, need try DeleteMulti again.\n",
					logId, "DeleteMulti", len(result.Errors)))
			}
			return nil
		}); err != nil {
			return err
		}
	}

	log.Printf("[DEBUG][ForceCleanObject]%s api[%s] success, [%v] objects have been cleaned.\n",
		logId, "ForceCleanObject", len(result.DeletedObjects))
	return nil
}

func (me *CosService) GetBucketCors(ctx context.Context, bucket string) (corsRules []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketCorsInput{
		Bucket: aws.String(bucket),
	}

	ratelimit.Check("GetBucketCors")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketCors(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if !ok || awsError.Code() != "NoSuchCORSConfiguration" {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "get bucket cors", request.String(), err.Error())
			errRet = fmt.Errorf("cos get bucket cors error: %s, bucket: %s", err.Error(), bucket)
			return
		}
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket cors", request.String(), response.String())

	corsRules = make([]map[string]interface{}, 0, len(response.CORSRules))
	if len(response.CORSRules) > 0 {
		for _, value := range response.CORSRules {
			rule := make(map[string]interface{})
			rule["allowed_origins"] = helper.StringsInterfaces(value.AllowedOrigins)
			rule["allowed_methods"] = helper.StringsInterfaces(value.AllowedMethods)
			rule["allowed_headers"] = helper.StringsInterfaces(value.AllowedHeaders)

			if value.ExposeHeaders != nil {
				rule["expose_headers"] = helper.StringsInterfaces(value.ExposeHeaders)
			}
			if value.MaxAgeSeconds != nil {
				rule["max_age_seconds"] = int(*value.MaxAgeSeconds)
			}

			corsRules = append(corsRules, rule)
		}
	}
	return
}

func (me *CosService) GetBucketLifecycle(ctx context.Context, bucket string) (lifecycleRules []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketLifecycleConfiguration")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketLifecycleConfiguration(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if !ok || awsError.Code() != "NoSuchLifecycleConfiguration" {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "get bucket lifecycle", request.String(), err.Error())
			errRet = fmt.Errorf("cos get bucket cors error: %s, bucket: %s", err.Error(), bucket)
			return
		}
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket lifecycle", request.String(), response.String())

	lifecycleRules = make([]map[string]interface{}, 0, len(response.Rules))
	if len(response.Rules) > 0 {
		for _, value := range response.Rules {
			rule := make(map[string]interface{})

			if value.ID != nil {
				rule["id"] = *value.ID
			}

			// filter_prefix and filter_tags
			if value.Filter != nil {
				if value.Filter.And != nil {
					// Handle And operator with prefix and/or tags
					if value.Filter.And.Prefix != nil && *value.Filter.And.Prefix != "" {
						rule["filter_prefix"] = *value.Filter.And.Prefix
					}
					if len(value.Filter.And.Tags) > 0 {
						tags := make(map[string]interface{})
						for _, tag := range value.Filter.And.Tags {
							if tag.Key != nil && tag.Value != nil {
								tags[*tag.Key] = *tag.Value
							}
						}
						rule["filter_tags"] = tags
					}
				} else if value.Filter.Prefix != nil && *value.Filter.Prefix != "" {
					// Handle simple prefix filter
					rule["filter_prefix"] = *value.Filter.Prefix
				} else if value.Filter.Tag != nil {
					// Handle single tag filter
					if value.Filter.Tag.Key != nil && value.Filter.Tag.Value != nil {
						tags := make(map[string]interface{})
						tags[*value.Filter.Tag.Key] = *value.Filter.Tag.Value
						rule["filter_tags"] = tags
					}
				}
			}
			// transition
			if len(value.Transitions) > 0 {
				transitions := make([]interface{}, 0, len(value.Transitions))
				for _, v := range value.Transitions {
					t := make(map[string]interface{})
					if v.Date != nil {
						t["date"] = (*v.Date).Format("2006-01-02")
					}
					if v.Days != nil {
						t["days"] = int(*v.Days)
					}
					if v.StorageClass != nil {
						t["storage_class"] = *v.StorageClass
					}
					transitions = append(transitions, t)
				}
				rule["transition"] = schema.NewSet(transitionHash, transitions)
			}
			// expiration
			if value.Expiration != nil {
				e := make(map[string]interface{})
				if value.Expiration.Date != nil {
					e["date"] = (*value.Expiration.Date).Format("2006-01-02")
				}
				if value.Expiration.Days != nil {
					e["days"] = int(*value.Expiration.Days)
				}
				if value.Expiration.ExpiredObjectDeleteMarker != nil {
					e["delete_marker"] = *value.Expiration.ExpiredObjectDeleteMarker
				}
				rule["expiration"] = schema.NewSet(expirationHash, []interface{}{e})
			}

			// transition
			if len(value.NoncurrentVersionTransitions) > 0 {
				transitions := make([]interface{}, 0, len(value.NoncurrentVersionTransitions))
				for _, v := range value.NoncurrentVersionTransitions {
					t := make(map[string]interface{})
					if v.NoncurrentDays != nil {
						t["non_current_days"] = int(*v.NoncurrentDays)
					}
					if v.StorageClass != nil {
						t["storage_class"] = *v.StorageClass
					}
					transitions = append(transitions, t)
				}
				rule["non_current_transition"] = schema.NewSet(transitionHash, transitions)
			}
			// non current expiration
			if value.NoncurrentVersionExpiration != nil {
				e := make(map[string]interface{})
				if value.NoncurrentVersionExpiration.NoncurrentDays != nil {
					e["non_current_days"] = int(*value.NoncurrentVersionExpiration.NoncurrentDays)
				}
				rule["non_current_expiration"] = schema.NewSet(nonCurrentExpirationHash, []interface{}{e})
			}

			lifecycleRules = append(lifecycleRules, rule)
		}
	}
	return
}

func (me *CosService) GetDataSourceBucketLifecycle(ctx context.Context, bucket string) (lifecycleRules []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(bucket),
	}

	ratelimit.Check("GetBucketLifecycleConfiguration")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketLifecycleConfiguration(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if !ok || awsError.Code() != "NoSuchLifecycleConfiguration" {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "get bucket lifecycle", request.String(), err.Error())
			errRet = fmt.Errorf("cos get bucket cors error: %s, bucket: %s", err.Error(), bucket)
			return
		}
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket lifecycle", request.String(), response.String())

	lifecycleRules = make([]map[string]interface{}, 0, len(response.Rules))
	if len(response.Rules) > 0 {
		for _, value := range response.Rules {
			rule := make(map[string]interface{})

			// filter_prefix and filter_tags
			if value.Filter != nil {
				if value.Filter.And != nil {
					// Handle And operator with prefix and/or tags
					if value.Filter.And.Prefix != nil && *value.Filter.And.Prefix != "" {
						rule["filter_prefix"] = *value.Filter.And.Prefix
					}
					if len(value.Filter.And.Tags) > 0 {
						tags := make(map[string]interface{})
						for _, tag := range value.Filter.And.Tags {
							if tag.Key != nil && tag.Value != nil {
								tags[*tag.Key] = *tag.Value
							}
						}
						rule["filter_tags"] = tags
					}
				} else if value.Filter.Prefix != nil && *value.Filter.Prefix != "" {
					// Handle simple prefix filter
					rule["filter_prefix"] = *value.Filter.Prefix
				} else if value.Filter.Tag != nil {
					// Handle single tag filter
					if value.Filter.Tag.Key != nil && value.Filter.Tag.Value != nil {
						tags := make(map[string]interface{})
						tags[*value.Filter.Tag.Key] = *value.Filter.Tag.Value
						rule["filter_tags"] = tags
					}
				}
			}
			// transition
			if len(value.Transitions) > 0 {
				transitions := make([]interface{}, 0, len(value.Transitions))
				for _, v := range value.Transitions {
					t := make(map[string]interface{})
					if v.Date != nil {
						t["date"] = (*v.Date).Format("2006-01-02")
					}
					if v.Days != nil {
						t["days"] = int(*v.Days)
					}
					if v.StorageClass != nil {
						t["storage_class"] = *v.StorageClass
					}
					transitions = append(transitions, t)
				}
				rule["transition"] = transitions
			}
			// expiration
			if value.Expiration != nil {
				e := make(map[string]interface{})
				if value.Expiration.Date != nil {
					e["date"] = (*value.Expiration.Date).Format("2006-01-02")
				}
				if value.Expiration.Days != nil {
					e["days"] = int(*value.Expiration.Days)
				}
				rule["expiration"] = []interface{}{e}
			}
			// non current transition
			if len(value.NoncurrentVersionTransitions) > 0 {
				transitions := make([]interface{}, 0, len(value.NoncurrentVersionTransitions))
				for _, v := range value.NoncurrentVersionTransitions {
					t := make(map[string]interface{})
					if v.NoncurrentDays != nil {
						t["non_current_days"] = int(*v.NoncurrentDays)
					}
					if v.StorageClass != nil {
						t["storage_class"] = *v.StorageClass
					}
					transitions = append(transitions, t)
				}
				rule["non_current_transition"] = transitions
			}
			// non current expiration
			if value.NoncurrentVersionExpiration != nil {
				e := make(map[string]interface{})
				if value.NoncurrentVersionExpiration.NoncurrentDays != nil {
					e["non_current_days"] = int(*value.NoncurrentVersionExpiration.NoncurrentDays)
				}
				rule["non_current_expiration"] = []interface{}{e}
			}

			lifecycleRules = append(lifecycleRules, rule)
		}
	}
	return
}

func (me *CosService) GetBucketWebsite(ctx context.Context, bucket string) (websites []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketWebsiteInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketWebsite")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketWebsite(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if ok && awsError.Code() == "NoSuchWebsiteConfiguration" {
			return
		}
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket website", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket website error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket website", request.String(), response.String())

	websites = make([]map[string]interface{}, 0, 1)
	website := make(map[string]interface{})
	if response.IndexDocument != nil {
		website["index_document"] = *response.IndexDocument.Suffix
	}
	if response.ErrorDocument != nil {
		website["error_document"] = *response.ErrorDocument.Key
	}
	if len(website) > 0 {
		websites = append(websites, website)
	}

	return
}

func (me *CosService) GetBucketEncryption(ctx context.Context, bucket string) (encryption string, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketEncryptionInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketEncryption")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketEncryption(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if ok && awsError.Code() == "NoSuchEncryptionConfiguration" {
			return
		}
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket encryption", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket encryption error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket encryption", request.String(), response.String())

	if response.ServerSideEncryptionConfiguration != nil && len(response.ServerSideEncryptionConfiguration.Rules) > 0 {
		rule := response.ServerSideEncryptionConfiguration.Rules[0]
		if rule.ApplyServerSideEncryptionByDefault != nil && rule.ApplyServerSideEncryptionByDefault.SSEAlgorithm != nil {
			encryption = *rule.ApplyServerSideEncryptionByDefault.SSEAlgorithm
		}
	}
	return
}

func (me *CosService) GetBucketVersioning(ctx context.Context, bucket string) (versioningEnable bool, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketVersioningInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketVersioning")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketVersioning(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if ok && awsError.Code() == "NoSuchVersioningConfiguration" {
			return
		}
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket versioning", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket versioning error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket versioning", request.String(), response.String())

	if response.Status == nil || *response.Status == "Suspended" {
		versioningEnable = false
	} else if *response.Status == "Enabled" {
		versioningEnable = true
	}

	return
}

func (me *CosService) GetBucketAccleration(ctx context.Context, bucket string) (accelerationEnable bool, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketAccelerateConfigurationInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketAccelerateConfiguration")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketAccelerateConfiguration(&request)
	if err != nil {
		awsError, ok := err.(awserr.Error)
		if ok && awsError.Code() == "NoSuchAccelerateConfiguration" {
			return
		}
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket acceleration", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket acceleration error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket acceleration", request.String(), response.String())

	if response.Status == nil || *response.Status == "Suspended" {
		accelerationEnable = false
	} else if *response.Status == "Enabled" {
		accelerationEnable = true
	}

	return
}

func (me *CosService) GetBucketLogStatus(ctx context.Context, bucket string) (logEnable bool, logTargetBucket string, logPrefix string, errRet error) {
	logId := getLogId(ctx)

	request := s3.GetBucketLoggingInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("GetBucketVersioning")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketLogging(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket log status", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket log status error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket log status", request.String(), response.String())

	if response.LoggingEnabled == nil || response.LoggingEnabled.TargetBucket == nil || *response.LoggingEnabled.TargetBucket == "" || response.LoggingEnabled.TargetPrefix == nil || *response.LoggingEnabled.TargetPrefix == "" {
		logEnable = false
	} else {
		logEnable = true
		logTargetBucket = *response.LoggingEnabled.TargetBucket
		logPrefix = *response.LoggingEnabled.TargetPrefix
	}

	return
}

func (me *CosService) ListBuckets(ctx context.Context) (buckets []*s3.Bucket, errRet error) {
	logId := getLogId(ctx)

	request := s3.ListBucketsInput{}
	ratelimit.Check("ListBuckets")
	response, err := me.client.UseCosS3Client(me.useCspClient).ListBuckets(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket list", request.String(), err.Error())
		errRet = fmt.Errorf("cos get bucket list error: %s", err.Error())
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket list", request.String(), response.String())

	buckets = response.Buckets
	return
}

func (me *CosService) ListObjects(ctx context.Context, bucket string) (objects []*s3.Object, errRet error) {
	logId := getLogId(ctx)

	request := s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("ListObjects")
	response, err := me.client.UseCosS3Client(me.useCspClient).ListObjects(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get object list", request.String(), err.Error())
		errRet = fmt.Errorf("cos get object list error: %s", err.Error())
		return
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get object list", request.String(), response.String())

	objects = response.Contents
	return
}

// SetBucketTags if len(tags) == 0, only delete tags
func (me *CosService) SetBucketTags(ctx context.Context, bucket string, tags map[string]string) error {
	logId := getLogId(ctx)

	deleteReq := &s3.DeleteBucketTaggingInput{Bucket: aws.String(bucket)}

	ratelimit.Check("DeleteBucketTagging")

	deleteResp, err := me.client.UseCosS3Client(me.useCspClient).DeleteBucketTagging(deleteReq)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]",
			logId, "delete olg tags", deleteReq.String(), err)
		return err
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "delete olg tags", deleteReq.String(), deleteResp.String())

	if len(tags) == 0 {
		return nil
	}

	putReq := &s3.PutBucketTaggingInput{
		Bucket:  aws.String(bucket),
		Tagging: new(s3.Tagging),
	}

	for k, v := range tags {
		putReq.Tagging.TagSet = append(putReq.Tagging.TagSet, &s3.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

	ratelimit.Check("PutBucketTagging")

	resp, err := me.client.UseCosS3Client(me.useCspClient).PutBucketTagging(putReq)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%v]",
			logId, "put new tags", deleteReq.String(), err)
		return err
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "put new tags", putReq.String(), resp.String())

	return nil
}

func (me *CosService) GetBucketTags(ctx context.Context, bucket string) (map[string]string, error) {
	logId := getLogId(ctx)

	req := &s3.GetBucketTaggingInput{Bucket: aws.String(bucket)}

	ratelimit.Check("GetBucketTagging")

	resp, err := me.client.UseCosS3Client(me.useCspClient).GetBucketTagging(req)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); !ok || awsErr.Code() != "404" {
			return nil, nil
		}

		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%v]",
			logId, "get tags", req.String(), err)
		return nil, err
	}

	tags := make(map[string]string, len(resp.TagSet))
	for _, t := range resp.TagSet {
		tags[*t.Key] = *t.Value
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "get tags", req.String(), resp.String())

	return tags, nil
}

func (me *CosService) GetObjectTags(ctx context.Context, bucket string, key string) (map[string]string, error) {
	logId := getLogId(ctx)

	req := &s3.GetObjectTaggingInput{
		Bucket: &bucket,
		Key:    &key,
	}

	ratelimit.Check("GetObjectTagging")
	resp, err := me.client.UseCosS3Client(me.useCspClient).GetObjectTagging(req)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); !ok || awsErr.Code() != "404" {
			return nil, nil
		}

		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%v]",
			logId, "get object tags", req.String(), err)
		return nil, err
	}

	tags := make(map[string]string, len(resp.TagSet))

	for _, tag := range resp.TagSet {
		tags[*tag.Key] = *tag.Value
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "get tags", req.String(), resp.String())

	return tags, nil
}

// SetObjectTags same as delete Bucket Tags
func (me *CosService) SetObjectTags(ctx context.Context, bucket string, key string, tags map[string]string) error {
	logId := getLogId(ctx)

	deleteReq := &s3.DeleteObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	ratelimit.Check("DeleteObjectTagging")

	deleteResp, err := me.client.UseCosS3Client(me.useCspClient).DeleteObjectTagging(deleteReq)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]",
			logId, "delete olg object tags", deleteReq.String(), err)
		return err
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "delete olg object tags", deleteReq.String(), deleteResp.String())

	if len(tags) == 0 {
		return nil
	}

	putReq := &s3.PutObjectTaggingInput{
		Key:     aws.String(key),
		Bucket:  aws.String(bucket),
		Tagging: new(s3.Tagging),
	}

	for k, v := range tags {
		putReq.Tagging.TagSet = append(putReq.Tagging.TagSet, &s3.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		})
	}

	ratelimit.Check("PutObjectTagging")

	resp, err := me.client.UseCosS3Client(me.useCspClient).PutObjectTagging(putReq)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%v]",
			logId, "put new object tags", deleteReq.String(), err)
		return err
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]",
		logId, "put new object tags", putReq.String(), resp.String())

	return nil
}

func (me *CosService) PutBucketPolicy(ctx context.Context, bucket, policy string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.PutBucketPolicyInput{
		Bucket: &bucket,
		Policy: &policy,
	}
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, "put bucket", request.String(), errRet.Error())
		}
	}()
	ratelimit.Check("PutBucketPolicy")
	response, err := me.client.UseCosS3Client(me.useCspClient).PutBucketPolicy(&request)
	if err != nil {
		errRet = fmt.Errorf("cos put bucket policy error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "put bucket policy", request.String(), response.String())

	return nil
}

func (me *CosService) DescribePolicyByBucket(ctx context.Context, bucket string) (bucketPolicy string, errRet error) {
	logId := getLogId(ctx)
	request := s3.GetBucketPolicyInput{Bucket: aws.String(bucket)}

	ratelimit.Check("GetBucketPolicy")
	response, err := me.client.UseCosS3Client(me.useCspClient).GetBucketPolicy(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "get bucket policy", request.String(), err.Error())
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "get bucket policy", request.String(), response.String())
	bucketPolicy = *response.Policy
	return
}

func (me *CosService) DeleteBucketPolicy(ctx context.Context, bucket string) (errRet error) {
	logId := getLogId(ctx)

	request := s3.DeleteBucketPolicyInput{
		Bucket: aws.String(bucket),
	}
	ratelimit.Check("DeleteBucketPolicy")
	response, err := me.client.UseCosS3Client(me.useCspClient).DeleteBucketPolicy(&request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, "delete bucket policy", request.String(), err.Error())
		return fmt.Errorf("cos delete bucket policy error: %s, bucket: %s", err.Error(), bucket)
	}
	log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
		logId, "delete bucket policy", request.String(), response.String())

	return nil
}

func (c *CosService) GetBucketACL(ctx context.Context, bucket string) (result *cos.BucketGetACLResult, errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "GetBucketACL", errRet.Error())
		}
	}()

	ratelimit.Check("TencentcloudCosPutBucketACL")
	// PathStyle：直接构造请求到 cos.<region>.<domain>/{bucket}?acl
	client := c.client.UseCosS3Client(c.useCspClient)
	request := &s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	}
	response, err := client.GetBucketAcl(request)
	if err != nil {
		errRet = fmt.Errorf("cos [GetBucketACL] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	// 转换 S3 ACL 响应为 cos.BucketGetACLResult 格式
	if response.Owner == nil {
		errRet = fmt.Errorf("cos [GetBucketACL] error: owner is nil, bucket: %s", bucket)
	}
	result = &cos.BucketGetACLResult{
		Owner: &cos.Owner{
			ID:          *response.Owner.ID,
			DisplayName: *response.Owner.DisplayName,
		},
		AccessControlList: make([]cos.ACLGrant, len(response.Grants)),
	}

	for i, grant := range response.Grants {
		if grant.Grantee == nil {
			continue
		}
		result.AccessControlList[i] = cos.ACLGrant{
			Grantee: &cos.ACLGrantee{
				Type: *grant.Grantee.Type,
				URI:  aws.StringValue(grant.Grantee.URI),
				ID:   aws.StringValue(grant.Grantee.ID),
			},
			Permission: *grant.Permission,
		}
	}

	log.Printf("[DEBUG]%s api[%s] success\n", logId, "GetBucketACL")
	return
}

func GetBucketPublicACL(acl *cos.BucketGetACLResult) string {
	var publicRead, publicWrite bool

	for i := range acl.AccessControlList {
		item := acl.AccessControlList[i]

		if item.Grantee.URI == PUBLIC_GRANTEE && item.Permission == "READ" {
			publicRead = true
		}

		if item.Grantee.URI == PUBLIC_GRANTEE && item.Permission == "WRITE" {
			publicWrite = true
		}
	}

	if publicRead && !publicWrite {
		return s3.ObjectCannedACLPublicRead
	}

	if publicRead && publicWrite {
		return s3.ObjectCannedACLPublicReadWrite
	}

	return s3.ObjectCannedACLPrivate
}

func (c *CosService) GetBucketPullOrigin(ctx context.Context, bucket string) (result []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "GetBucketPullOrigin", errRet.Error())
		}
	}()

	ratelimit.Check("TencentcloudCosGetBucketPullOrigin")
	originConfig, response, err := c.GetClient(bucket).Bucket.GetOrigin(ctx)

	if response.StatusCode == 404 {
		return make([]map[string]interface{}, 0), nil
	}

	if err != nil {
		errRet = fmt.Errorf("cos [GetBucketPullOrigin] error: %s, bucket: %s", err.Error(), bucket)
		return nil, errRet
	}

	resp, _ := json.Marshal(originConfig)

	log.Printf("[DEBUG]%s api[%s] success, request body response body [%s]\n",
		logId, "GetBucketPullOrigin", resp)

	rules := make([]map[string]interface{}, 0)

	for _, rule := range originConfig.Rule {
		item := make(map[string]interface{})
		item["priority"] = helper.Int(rule.RulePriority)
		// item["host"] = helper.String(rule.OriginInfo.HostInfo)

		if rule.OriginCondition != nil {
			item["prefix"] = helper.String(rule.OriginCondition.Prefix)
		}

		if rule.OriginType == "Mirror" {
			item["sync_back_to_source"] = helper.Bool(true)
		} else if rule.OriginType == "Proxy" {
			item["sync_back_to_source"] = helper.Bool(false)
		}

		if rule.OriginParameter != nil {
			if rule.OriginParameter.HttpHeader != nil {
				if len(rule.OriginParameter.HttpHeader.NewHttpHeaders) != 0 {
					headers := make(map[string]interface{})
					for _, header := range rule.OriginParameter.HttpHeader.NewHttpHeaders {
						headers[header.Key] = helper.String(header.Value)
					}
					item["custom_http_headers"] = headers
				}

				if len(rule.OriginParameter.HttpHeader.FollowHttpHeaders) != 0 {
					headers := schema.NewSet(func(i interface{}) int {
						return helper.HashString(i.(string))
					}, nil)
					for _, header := range rule.OriginParameter.HttpHeader.FollowHttpHeaders {
						headers.Add(header.Key)
					}
					item["follow_http_headers"] = headers
				}

			}
			item["protocol"] = helper.String(rule.OriginParameter.Protocol)
			item["follow_redirection"] = helper.Bool(rule.OriginParameter.FollowRedirection)
			item["follow_query_string"] = helper.Bool(rule.OriginParameter.FollowQueryString)
		}

		if rule.OriginInfo.FileInfo != nil {
			// item["host"] = helper.String(rule.OriginInfo.HostInfo)
			//item["redirect_prefix"] = helper.String(rule.OriginInfo.FileInfo.Prefix)
			//item["redirect_suffix"] = helper.String(rule.OriginInfo.FileInfo.Suffix)
		}

		rules = append(rules, item)
	}

	return rules, nil
}

func (c *CosService) PutBucketPullOrigin(ctx context.Context, bucket string, rules []cos.BucketOriginRule) (errRet error) {
	logId := getLogId(ctx)

	opt := &cos.BucketPutOriginOptions{
		Rule: rules,
	}
	ratelimit.Check("PutBucketPullOrigin")
	response, err := c.GetClient(bucket).Bucket.PutOrigin(ctx, opt)
	req, _ := json.Marshal(opt)
	resp, _ := json.Marshal(response.Response.Body)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request:[%s], reason[%s]\n",
				logId, "PutBucketPullOrigin", req, errRet.Error())
		}
	}()

	if err != nil {
		errRet = fmt.Errorf("[PutBucketPullOrigin] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[PutBucketPullOrigin] success, request body [%s], response body [%s]\n",
		logId, req, resp)

	return nil
}

func (c *CosService) DeleteBucketPullOrigin(ctx context.Context, bucket string) (errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "DeleteBucketPullOrigin", errRet.Error())
		}
	}()

	ratelimit.Check("DeleteBucketPullOrigin")
	response, err := c.GetClient(bucket).Bucket.DeleteOrigin(ctx)
	resp, _ := json.Marshal(response.Response.Body)

	if err != nil {
		errRet = fmt.Errorf("[DeleteBucketPullOrigin] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[DeleteBucketPullOrigin] success, response body [%s]\n",
		logId, resp)

	return nil
}

func (c *CosService) GetBucketOriginDomain(ctx context.Context, bucket string) (result []map[string]interface{}, errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "GetBucketOriginDomain", errRet.Error())
		}
	}()

	ratelimit.Check("TencentcloudCosGetBucketOriginDomain")
	domain, response, err := c.GetClient(bucket).Bucket.GetDomain(ctx)

	if response.StatusCode == 404 {
		log.Printf("[WARN] [GetBucketOriginDomain] returns %d, %s", 404, err)
		return make([]map[string]interface{}, 0), nil
	}

	if err != nil {
		errRet = fmt.Errorf("cos [GetBucketOriginDomain] error: %s, bucket: %s", err.Error(), bucket)
		return nil, errRet
	}

	rules := make([]map[string]interface{}, 0)

	for _, rule := range domain.Rules {
		item := make(map[string]interface{})
		item["domain"] = helper.String(rule.Name)
		item["status"] = helper.String(rule.Status)
		item["type"] = helper.String(rule.Type)
		rules = append(rules, item)
	}

	resp, _ := json.Marshal(response.Response.Body)

	log.Printf("[DEBUG]%s api[%s] success, request body response body [%s]\n",
		logId, "GetBucketOriginDomain", resp)

	return rules, nil
}

func (c *CosService) PutBucketOriginDomain(ctx context.Context, bucket string, rules []cos.BucketDomainRule) (errRet error) {
	logId := getLogId(ctx)

	opt := &cos.BucketPutDomainOptions{
		Rules: rules,
	}
	ratelimit.Check("PutBucketOriginDomain")
	response, err := c.GetClient(bucket).Bucket.PutDomain(ctx, opt)
	req, _ := json.Marshal(opt)
	resp, _ := json.Marshal(response.Response.Body)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request:[%s], reason[%s]\n",
				logId, "PutBucketOriginDomain", req, errRet.Error())
		}
	}()

	if err != nil {
		errRet = fmt.Errorf("[PutBucketOriginDomain] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[PutBucketOriginDomain] success, request body [%s], response body [%s]\n",
		logId, req, resp)

	return nil
}

func (me *CosService) DeleteBucketOriginDomain(ctx context.Context, bucket string) (errRet error) {
	logId := getLogId(ctx)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "DeleteBucketOriginDomain", errRet.Error())
		}
	}()

	ratelimit.Check("DeleteBucketOriginDomain")
	response, err := me.client.UseTencentCosClient(bucket).Bucket.DeleteDomain(ctx)
	resp, _ := json.Marshal(response.Response.Body)

	if err != nil {
		errRet = fmt.Errorf("[DeleteBucketOriginDomain] error: %s, bucket: %s", err.Error(), bucket)
		return
	}

	log.Printf("[DEBUG]%s api[DeleteBucketOriginDomain] success, response body [%s]\n",
		logId, resp)

	return nil
}

func (c *CosService) GetBucketReplication(ctx context.Context, bucket string) (result *cos.GetBucketReplicationResult, errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "GetBucketReplication", errRet.Error())
		}
	}()

	ratelimit.Check("GetBucketReplication")
	result, response, err := c.GetClient(bucket).Bucket.GetBucketReplication(ctx)

	if response.StatusCode == 404 {
		log.Printf("[WARN]%s, api[%s] returns %d", logId, "GetBucketReplication", response.StatusCode)
		return
	}

	resp, _ := json.Marshal(response.Response.Body)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] response body [%s]\n",
		logId, "GetBucketReplication", resp)

	return
}

func (me *CosService) PutBucketReplication(ctx context.Context, bucket string, role string, rules []cos.BucketReplicationRule) (errRet error) {
	logId := getLogId(ctx)

	option := &cos.PutBucketReplicationOptions{
		Role: role,
		Rule: rules,
	}

	request, _ := xml.Marshal(option)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request: %s reason[%s]\n",
				logId, "PutBucketReplication", request, errRet.Error())
		}
	}()

	ratelimit.Check("PutBucketReplication")
	response, err := me.client.UseTencentCosClient(bucket).Bucket.PutBucketReplication(ctx, option)

	resp, _ := json.Marshal(response.Response.Body)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] response body [%s]\n",
		logId, "PutBucketReplication", resp)

	return
}

func (me *CosService) DeleteBucketReplication(ctx context.Context, bucket string) (errRet error) {
	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, reason[%s]\n",
				logId, "DeleteBucketReplication", errRet.Error())
		}
	}()

	ratelimit.Check("DeleteBucketReplication")
	response, err := me.client.UseTencentCosClient(bucket).Bucket.DeleteBucketReplication(ctx)

	resp, _ := json.Marshal(response.Response.Body)

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] response body [%s]\n",
		logId, "DeleteBucketReplication", resp)

	return
}

func (me *CosService) DescribeCosBucketDomainCertificate(ctx context.Context, certId string) (result *cos.BucketGetDomainCertificateResult, bucket string, errRet error) {
	logId := getLogId(ctx)

	ids, err := me.parseCertId(certId)
	if err != nil {
		errRet = err
		return
	}

	bucket = ids.bucket
	domainName := ids.domainName
	option := &cos.BucketGetDomainCertificateOptions{
		DomainName: domainName,
	}
	request, _ := xml.Marshal(option)

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request[%s], reason[%s]\n",
				logId, "GetDomainCertificate", request, errRet.Error())
		}
	}()

	result, response, err := me.client.UseTencentCosClient(bucket).Bucket.GetDomainCertificate(ctx, option)
	resp, _ := json.Marshal(response.Response.Body)
	if response.StatusCode == 404 {
		log.Printf("[WARN]%s, api[%s] returns %d", logId, "GetDomainCertificate", response.StatusCode)
		return
	}

	if err != nil {
		errRet = err
		return
	}

	log.Printf("[DEBUG]%s api[%s] success, request [%s], response body [%s], result [%s]\n",
		logId, "GetDomainCertificate", request, resp, result)

	return
}

func (c *CosService) DeleteCosBucketDomainCertificate(ctx context.Context, certId string) (errRet error) {
	logId := getLogId(ctx)

	ids, err := c.parseCertId(certId)
	if err != nil {
		errRet = err
		return
	}

	bucket := ids.bucket
	domainName := ids.domainName
	option := &cos.BucketDeleteDomainCertificateOptions{
		DomainName: domainName,
	}

	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, option [%s], reason[%s]\n",
				logId, "DeleteDomainCertificate", option, errRet.Error())
		}
	}()

	ratelimit.Check("DeleteDomainCertificate")
	response, err := c.GetClient(bucket).Bucket.DeleteDomainCertificate(ctx, option)

	if err != nil {
		errRet = err
		return err
	}

	resp, _ := json.Marshal(response.Response.Body)

	log.Printf("[DEBUG]%s api[%s] success, option [%s], response body [%s]\n",
		logId, "DeleteDomainCertificate", option, resp)

	return
}

func (me *CosService) parseCertId(configId string) (ret *CosBucketDomainCertItem, err error) {
	idSplit := strings.Split(configId, FILED_SP)
	if len(idSplit) != 2 {
		return nil, fmt.Errorf("id is broken,%s", configId)
	}

	bucket := idSplit[0]
	domain := idSplit[1]
	if bucket == "" || domain == "" {
		return nil, fmt.Errorf("id is broken,%s", configId)
	}

	ret = &CosBucketDomainCertItem{bucket, domain}
	return
}

// ListTencentCloudBuckets list all buckets
func (c *CosService) ListTencentCloudBuckets(ctx context.Context) (*cos.ServiceGetResult, error) {
	getServiceResult, _, err := c.GetClient("").Service.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "cos list buckets error")
	}
	return getServiceResult, nil
}

// GetBucketObject get bucket object
func (c *CosService) GetBucketObject(ctx context.Context, bucket string,
	opt *cos.BucketGetOptions) (*cos.BucketGetResult, error) {
	getBucketResult, _, err := c.GetClient(bucket).Bucket.Get(ctx, opt)
	if err != nil {
		return nil, errors.Wrap(err, "cos get bucket object error")
	}
	return getBucketResult, nil
}

// CspNotificationConfiguration represents the full notification config XML with CSP extensions.
type CspNotificationConfiguration struct {
	XMLName             xml.Name                `xml:"NotificationConfiguration"`
	Xmlns               string                  `xml:"xmlns,attr,omitempty"`
	TopicConfigurations []CspTopicConfiguration `xml:"TopicConfiguration,omitempty"`
}

// CspTopicConfiguration represents a single notification rule with CSP-specific Kafka fields.
type CspTopicConfiguration struct {
	Id       string   `xml:"Id"`
	Events   []string `xml:"Event"`
	Topic    string   `xml:"Topic"`              // fixed "csp"
	KafkaID  string   `xml:"KafkaID"`            // CSP extension: ckafka instance id
	Endpoint string   `xml:"Endpoint"`           // CSP extension: kafka://host:port
	User     string   `xml:"User,omitempty"`     // SASL username: format "instanceId#appId"
	Password string   `xml:"Password,omitempty"` // SASL password
}

// PutBucketNotification writes a notification configuration to a CSP bucket.
// Uses virtual-hosted style URL (bucket.cos.region.domain) with cos.AuthorizationTransport signing.
func (me *CosService) PutBucketNotification(ctx context.Context, bucket string, config *CspNotificationConfiguration) error {
	logId := getLogId(ctx)

	xmlBytes, err := xml.Marshal(config)
	if err != nil {
		return fmt.Errorf("marshal notification xml error: %s", err.Error())
	}
	xmlBody := xml.Header + string(xmlBytes)

	log.Printf("[DEBUG]%s PutBucketNotification bucket[%s] body[%s]", logId, bucket, xmlBody)

	notifURL := fmt.Sprintf("%s://%s.cos.%s.%s/?notification",
		me.client.Protocol, bucket, me.client.Region, me.client.CspDomain)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, notifURL, strings.NewReader(xmlBody))
	if err != nil {
		return fmt.Errorf("create put notification request error: %s", err.Error())
	}
	httpReq.Header.Set("Content-Type", "application/xml")

	httpClient := &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:     me.client.Credential.SecretId,
			SecretKey:    me.client.Credential.SecretKey,
			SessionToken: me.client.Credential.Token,
		},
	}

	ratelimit.Check("PutBucketNotificationConfiguration")
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("put bucket notification error: %s, bucket: %s", err.Error(), bucket)
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("put bucket notification failed, status: %d, body: %s", resp.StatusCode, string(body))
	}

	log.Printf("[DEBUG]%s PutBucketNotification success, bucket[%s]", logId, bucket)
	return nil
}

// GetBucketNotification reads the notification configuration from a CSP bucket.
// Uses virtual-hosted style URL with cos.AuthorizationTransport signing.
func (me *CosService) GetBucketNotification(ctx context.Context, bucket string) (*CspNotificationConfiguration, error) {
	logId := getLogId(ctx)

	notifURL := fmt.Sprintf("%s://%s.cos.%s.%s/?notification",
		me.client.Protocol, bucket, me.client.Region, me.client.CspDomain)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, notifURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create get notification request error: %s", err.Error())
	}

	httpClient := &http.Client{
		Timeout: 100 * time.Second,
		Transport: &cos.AuthorizationTransport{
			SecretID:     me.client.Credential.SecretId,
			SecretKey:    me.client.Credential.SecretKey,
			SessionToken: me.client.Credential.Token,
		},
	}

	ratelimit.Check("GetBucketNotificationConfiguration")
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("get bucket notification error: %s, bucket: %s", err.Error(), bucket)
	}
	defer resp.Body.Close()

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read bucket notification response error: %s", err.Error())
	}

	log.Printf("[DEBUG]%s GetBucketNotification bucket[%s] response[%s]", logId, bucket, string(rawBody))

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("get bucket notification failed, status: %d, body: %s", resp.StatusCode, string(rawBody))
	}

	var config CspNotificationConfiguration
	if err := xml.Unmarshal(rawBody, &config); err != nil {
		return nil, fmt.Errorf("unmarshal notification xml error: %s, body: %s", err.Error(), string(rawBody))
	}

	return &config, nil
}

// DeleteBucketNotification removes all notification rules by putting an empty configuration.
func (me *CosService) DeleteBucketNotification(ctx context.Context, bucket string) error {
	emptyConfig := &CspNotificationConfiguration{}
	return me.PutBucketNotification(ctx, bucket, emptyConfig)
}

// ResolveCkafkaEndpoint finds or creates a PLAINTEXT route for the given CKafka instance
// and returns the kafka endpoint in the format "kafka://host:port".
func (me *CosService) ResolveCkafkaEndpoint(ctx context.Context, instanceId string) (string, error) {
	logId := getLogId(ctx)

	ckafkaService := CkafkaService{client: me.client}

	route, err := ckafkaService.DescribeCkafkaRouteByKey(ctx, instanceId, 4, "", "", 0)
	if err != nil {
		return "", fmt.Errorf("describe ckafka route error: %s", err.Error())
	}

	if route != nil && route.Processing != nil && *route.Processing == 0 &&
		route.VipList != nil && len(route.VipList) > 0 &&
		route.VipList[0].Vip != nil && *route.VipList[0].Vip != "" {
		endpoint := fmt.Sprintf("kafka://%s:%s", *route.VipList[0].Vip, *route.VipList[0].Vport)
		log.Printf("[DEBUG]%s found existing ckafka route, endpoint: %s", logId, endpoint)
		return endpoint, nil
	}

	log.Printf("[DEBUG]%s no ready PLAINTEXT route for ckafka %s, creating...", logId, instanceId)

	createReq := ckafka.NewCreateRouteRequest()
	createReq.InstanceId = &instanceId
	vipType := int64(4)
	createReq.VipType = &vipType
	accessType := int64(0)
	createReq.AccessType = &accessType

	ratelimit.Check("CreateRoute")
	_, err = me.client.UseCkafkaClient().CreateRoute(createReq)
	if err != nil {
		return "", fmt.Errorf("create ckafka route error: %s", err.Error())
	}

	var endpoint string
	err = resource.Retry(3*time.Minute, func() *resource.RetryError {
		r, e := ckafkaService.DescribeCkafkaRouteByKey(ctx, instanceId, 4, "", "", 0)
		if e != nil {
			return retryError(e)
		}
		if r == nil {
			return resource.RetryableError(fmt.Errorf("ckafka route not found yet"))
		}
		if r.Processing != nil && *r.Processing != 0 {
			return resource.RetryableError(fmt.Errorf("ckafka route still processing"))
		}
		if r.VipList == nil || len(r.VipList) == 0 || r.VipList[0].Vip == nil || *r.VipList[0].Vip == "" {
			return resource.RetryableError(fmt.Errorf("ckafka route vip not ready"))
		}
		endpoint = fmt.Sprintf("kafka://%s:%s", *r.VipList[0].Vip, *r.VipList[0].Vport)
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("waiting for ckafka route ready error: %s", err.Error())
	}

	log.Printf("[DEBUG]%s created ckafka route, endpoint: %s", logId, endpoint)
	return endpoint, nil
}

// BuildNotificationConfig constructs a CspNotificationConfiguration from terraform resource data.
func BuildNotificationConfig(rules []interface{}, endpoints map[string]string) *CspNotificationConfiguration {
	config := &CspNotificationConfiguration{}
	for _, raw := range rules {
		rule := raw.(map[string]interface{})
		tc := CspTopicConfiguration{
			Id:      rule["id"].(string),
			Topic:   "csp",
			KafkaID: rule["ckafka_instance_id"].(string),
		}
		for _, e := range rule["events"].([]interface{}) {
			tc.Events = append(tc.Events, e.(string))
		}
		if ep, ok := endpoints[tc.KafkaID]; ok {
			tc.Endpoint = ep
		}
		if v, ok := rule["sasl_user"].(string); ok && v != "" {
			// Format: {instanceId}#{appId} as required by CSP
			tc.User = tc.KafkaID + "#" + v
		}
		if v, ok := rule["sasl_password"].(string); ok && v != "" {
			tc.Password = v
		}
		config.TopicConfigurations = append(config.TopicConfigurations, tc)
	}
	return config
}

// FlattenNotificationRules converts CspNotificationConfiguration back to terraform state format.
// sasl_user and sasl_password are write-only fields not returned by GET, so they are preserved
// from existing state via the existingRules parameter.
func FlattenNotificationRules(config *CspNotificationConfiguration, existingRules []interface{}) []map[string]interface{} {
	if config == nil {
		return nil
	}

	// Build a lookup of existing SASL credentials by rule ID so they survive Read.
	existingSasl := make(map[string][2]string) // id -> [user, password]
	for _, raw := range existingRules {
		rule := raw.(map[string]interface{})
		id, _ := rule["id"].(string)
		user, _ := rule["sasl_user"].(string)
		pass, _ := rule["sasl_password"].(string)
		if id != "" {
			existingSasl[id] = [2]string{user, pass}
		}
	}

	var rules []map[string]interface{}
	for _, tc := range config.TopicConfigurations {
		saslUser := ""
		saslPassword := ""
		if creds, ok := existingSasl[tc.Id]; ok {
			saslUser = creds[0]
			saslPassword = creds[1]
		}
		rule := map[string]interface{}{
			"id":                 tc.Id,
			"events":             tc.Events,
			"ckafka_instance_id": tc.KafkaID,
			"endpoint":           tc.Endpoint,
			"sasl_user":          saslUser,
			"sasl_password":      saslPassword,
		}
		rules = append(rules, rule)
	}
	return rules
}
