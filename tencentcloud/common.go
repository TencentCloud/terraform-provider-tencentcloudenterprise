package tencentcloud

import (
	cos "github.com/tencentyun/cos-go-sdk-v5"

	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	publicSdkErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"time"
)

var contextNil context.Context = nil

type contextLogId string

const logIdKey = contextLogId("logId")

const (
	PROVIDER_READ_RETRY_TIMEOUT  = "TENCENTCLOUD_READ_RETRY_TIMEOUT"
	PROVIDER_WRITE_RETRY_TIMEOUT = "TENCENTCLOUD_WRITE_RETRY_TIMEOUT"
	PROVIDER_WAIT_READ_TIMEOUT   = "TENCENTCLOUD_WAIT_READ_TIMEOUT"

	SWEEPER_NEED_PROTECT            = "SWEEPER_NEED_PROTECT"
	TENCENTCLOUD_COMMON_TIME_LAYOUT = "2006-01-02 15:04:05"
)

var logFirstTime = ""

var logAtomicId int64 = 0

// readRetryTimeout is read retry timeout
// const readRetryTimeout = 3 * time.Minute
var readRetry = getEnvDefault(PROVIDER_READ_RETRY_TIMEOUT, 3)

var readRetryTimeout = time.Duration(readRetry) * time.Minute

// writeRetryTimeout is write retry timeout
// const writeRetryTimeout = 5 * time.Minute
var writeRetry = getEnvDefault(PROVIDER_WRITE_RETRY_TIMEOUT, 5)

var writeRetryTimeout = time.Duration(writeRetry) * time.Minute

// InternalError common internalError, do not add in retryableErrorCode,
// because when some product return this error, retry won't fix anything.
const InternalError = "InternalError"

// retryableErrorCode is retryable error code
var retryableErrorCode = []string{
	// client
	"ClientError.NetworkError",
	"ClientError.HttpStatusCodeError",
	// common
	"FailedOperation",
	"TradeUnknownError",
	"RequestLimitExceeded",
	"ResourceInUse",
	"ResourceUnavailable",
	// cbs
	"ResourceBusy",
	// teo
	"InvalidParameter.ActionInProgress",
	// posgresql
	"OperationDenied.InstanceStatusLimitError",
}

// retryableCosErrorCode is retryable error code for COS/CI SDK
var retryableCosErrorCode = []string{
	"RequestTimeout",
	"InternalError",
	"KmsInternalException",
	"ServiceUnavailable",
	"SlowDown",
}

func init() {
	logFirstTime = fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
}

func getEnvDefault(key string, defVal int) int {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	int, err := strconv.Atoi(val)
	if err != nil {
		panic("TENCENTCLOUD_READ_RETRY_TIMEOUT or TENCENTCLOUD_WRITE_RETRY_TIMEOUT must be int.")
	}
	return int
}

// getLogId get logId for trace, return a new logId if ctx is nil
func getLogId(ctx context.Context) string {
	if ctx != nil {
		logId, ok := ctx.Value(logIdKey).(string)
		if ok {
			return logId
		}
	}

	return fmt.Sprintf("%s-%d", logFirstTime, atomic.AddInt64(&logAtomicId, 1))
}

// logElapsed log func elapsed time, using in defer
func logElapsed(mark ...string) func() {
	startAt := time.Now()
	return func() {
		log.Printf("[DEBUG] [ELAPSED] %s elapsed %d ms\n", strings.Join(mark, " "), int64(time.Since(startAt)/time.Millisecond))
	}
}

// for Provider produced inconsistent result after apply
func inconsistentCheck(d *schema.ResourceData, meta interface{}) func() {
	oldJson, _ := json.Marshal(d.State())
	return func() {
		newJson, _ := json.Marshal(d.State())
		if !reflect.DeepEqual(oldJson, newJson) {
			log.Printf("[Resource id %s data changes after reading old:%s, new:%s", d.Id(), oldJson, newJson)
		}
	}
}

// retryError returns retry error
func retryError(err error, additionRetryableError ...string) *resource.RetryError {
	switch realErr := errors.Cause(err).(type) {
	case *sdkErrors.CloudSDKError:
		if isExpectError(realErr, retryableErrorCode) {
			log.Printf("[CRITAL] Retryable defined error: %v", err)
			return resource.RetryableError(err)
		}

		if len(additionRetryableError) > 0 {
			if isExpectError(realErr, additionRetryableError) {
				log.Printf("[CRITAL] Retryable addition error: %v", err)
				return resource.RetryableError(err)
			}
		}
	case *publicSdkErrors.TencentCloudSDKError:
		if isExpectError(realErr, retryableErrorCode) {
			log.Printf("[CRITAL] Retryable defined error: %v", err)
			return resource.RetryableError(err)
		}

		if len(additionRetryableError) > 0 {
			if isExpectError(realErr, additionRetryableError) {
				log.Printf("[CRITAL] Retryable addition error: %v", err)
				return resource.RetryableError(err)
			}
		}
	case *cos.ErrorResponse:
		if isCosExpectedError(realErr, retryableCosErrorCode) {
			log.Printf("[CRITAL] Retryable defined error: %v", err)
			return resource.RetryableError(err)
		}
		if len(additionRetryableError) > 0 {
			if isCosExpectedError(realErr, additionRetryableError) {
				log.Printf("[CRITAL] Retryable additional error: %v", err)
				return resource.RetryableError(err)
			}
		}
	default:
	}

	log.Printf("[CRITAL] NonRetryable error: %v", err)
	return resource.NonRetryableError(err)
}

// isCosExpectedError returns whether error is expected error when using COS SDK
func isCosExpectedError(err error, expectedError []string) bool {
	e, ok := err.(*cos.ErrorResponse)
	if !ok {
		return false
	}

	errCode := e.Code
	if IsContains(expectedError, errCode) {
		return true
	} else {
		return false
	}
}

// isExpectError returns whether error is expected error
func isExpectError(err error, expectError []string) bool {
	var longCode string
	switch e := errors.Cause(err).(type) {
	case *sdkErrors.CloudSDKError:
		longCode = e.Code
	case *publicSdkErrors.TencentCloudSDKError:
		longCode = e.Code
	default:
		return false
	}

	if IsContains(expectError, longCode) {
		return true
	}

	if strings.Contains(longCode, ".") {
		shortCode := strings.Split(longCode, ".")[0]
		if IsContains(expectError, shortCode) {
			return true
		}
	}

	return false
}

// IsNil Determine whether i is empty
func IsNil(v interface{}) bool {

	valueOf := reflect.ValueOf(v)

	k := valueOf.Kind()

	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return valueOf.IsNil()
	default:
		return v == nil
	}
}

// IsString Determine whether data is a string
func IsString(data interface{}) bool {
	if IsNil(data) {
		return false
	}

	if _, ok := data.(string); ok {
		return true
	}

	return false
}

// writeToFile write data to file
func writeToFile(filePath string, data interface{}) error {
	if strings.HasPrefix(filePath, "~") {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("Get current user fail,reason %s", err.Error())
		}
		if usr.HomeDir != "" {
			filePath = strings.Replace(filePath, "~", usr.HomeDir, 1)
		}
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("stat old file error,reason %s", err.Error())
	}

	if !os.IsNotExist(err) {
		if fileInfo.IsDir() {
			return fmt.Errorf("old filepath is a dir,can not delete")
		}
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("delete old file error,reason %s", err.Error())
		}
	}

	if IsString(data) {
		return ioutil.WriteFile(filePath, []byte(data.(string)), 0422)
	}

	jsonStr, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("json decode error,reason %s", err.Error())
	}

	return ioutil.WriteFile(filePath, jsonStr, 0422)
}

// BuildTagResourceName builds the Tencent Cloud specific name of a resource description.
// The format is `qcs:project_id:service_type:region:account:resource`.
// For more information, go to https://cloud.tencent.com/document/product/598/10606.
func BuildTagResourceName(serviceType, resourceType, region, id string) string {
	switch serviceType {
	case "cos":
		return fmt.Sprintf("qcs::%s:%s:uid/:%s/%s", serviceType, region, resourceType, id)

	default:
		return fmt.Sprintf("qcs::%s:%s:uin/:%s/%s", serviceType, region, resourceType, id)
	}
}

// IsContains returns whether value is within array
func IsContains(array interface{}, value interface{}) bool {
	vv := reflect.ValueOf(array)
	if vv.Kind() == reflect.Ptr || vv.Kind() == reflect.Interface {
		if vv.IsNil() {
			return false
		}
		vv = vv.Elem()
	}

	switch vv.Kind() {
	case reflect.Invalid:
		return false
	case reflect.Slice:
		for i := 0; i < vv.Len(); i++ {
			if reflect.DeepEqual(value, vv.Index(i).Interface()) {
				return true
			}
		}
		return false
	case reflect.Map:
		s := vv.MapKeys()
		for i := 0; i < len(s); i++ {
			if reflect.DeepEqual(value, s[i].Interface()) {
				return true
			}
		}
		return false
	case reflect.String:
		ss := reflect.ValueOf(value)
		switch ss.Kind() {
		case reflect.String:
			return strings.Contains(vv.String(), ss.String())
		}
		return false
	default:
		return reflect.DeepEqual(array, value)
	}
}

func GetListDiffs(o []int, n []int) (adds []int, lacks []int) {
	fillArr := func(arr []int, v int, count int) []int {
		for i := 0; i < count; i++ {
			arr = append(arr, v)
		}
		return arr
	}
	diffs := map[int]int{}
	for _, v := range o {
		diffs[v] -= 1
	}
	for _, v := range n {
		diffs[v] += 1
	}
	log.Printf("DIFFS: %v", diffs)
	for num, count := range diffs {
		if count < 0 {
			lacks = fillArr(lacks, num, -count)
		} else if count > 0 {
			adds = fillArr(adds, num, count)
		}
	}
	return
}

// unwrapParseJsonBusinessError recovers a business API error that was masked by
// ClientError.ParseJsonError. Some TCE error payloads include both string "Code"
// and numeric "code"; Go's encoding/json is case-insensitive and fails the parse
// before the SDK can surface the real FailedOperation.* code.
func unwrapParseJsonBusinessError(err error) error {
	if err == nil {
		return nil
	}
	sdkErr, ok := errors.Cause(err).(*sdkErrors.CloudSDKError)
	if !ok || sdkErr.Code != "ClientError.ParseJsonError" {
		return err
	}

	const prefix = "Fail to parse json content: "
	msg := sdkErr.Message
	start := strings.Index(msg, prefix)
	if start < 0 {
		return err
	}
	rest := msg[start+len(prefix):]
	end := strings.LastIndex(rest, ", because:")
	if end < 0 {
		return err
	}
	raw := strings.TrimSpace(rest[:end])

	var root map[string]json.RawMessage
	if e := json.Unmarshal([]byte(raw), &root); e != nil {
		return err
	}
	respRaw, ok := root["Response"]
	if !ok {
		return err
	}
	var response map[string]json.RawMessage
	if e := json.Unmarshal(respRaw, &response); e != nil {
		return err
	}
	errRaw, ok := response["Error"]
	if !ok {
		return err
	}
	var apiErr map[string]json.RawMessage
	if e := json.Unmarshal(errRaw, &apiErr); e != nil {
		return err
	}

	var code, message, requestId string
	if v, ok := apiErr["Code"]; ok {
		_ = json.Unmarshal(v, &code)
	}
	if v, ok := apiErr["Message"]; ok {
		_ = json.Unmarshal(v, &message)
	}
	if v, ok := response["RequestId"]; ok {
		_ = json.Unmarshal(v, &requestId)
	}
	if code == "" {
		return err
	}

	log.Printf("[WARN] recovered business error from ParseJsonError: code=%s message=%s requestId=%s", code, message, requestId)
	return sdkErrors.NewCloudSDKError(code, message, requestId)
}
