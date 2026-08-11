package tencentcloud

import (
	"strings"
	"testing"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
)

func TestUnwrapParseJsonBusinessError_NameAlreadyExist(t *testing.T) {
	raw := `Fail to parse json content: {"Response":{"Error":{"Code":"FailedOperation.NameAlreadyExist","Message":"username already exist; 子账号名称 c2-tf-sa1-nprd 已存在","code":1210,"msg":"username already exist","value":null},"RequestId":"973ba99a-6650-e63e-6445-a5254c8823ff"}}, because: json: cannot unmarshal number into Go struct field .Response.Error.Code of type string`
	err := sdkErrors.NewCloudSDKError("ClientError.ParseJsonError", raw, "973ba99a-6650-e63e-6445-a5254c8823ff")

	unwrapped := unwrapParseJsonBusinessError(err)
	sdkErr, ok := unwrapped.(*sdkErrors.CloudSDKError)
	if !ok {
		t.Fatalf("expected CloudSDKError, got %T: %v", unwrapped, unwrapped)
	}
	if sdkErr.Code != "FailedOperation.NameAlreadyExist" {
		t.Fatalf("unexpected code: %s", sdkErr.Code)
	}
	if !strings.Contains(sdkErr.Message, "c2-tf-sa1-nprd") {
		t.Fatalf("unexpected message: %s", sdkErr.Message)
	}
	if sdkErr.RequestId != "973ba99a-6650-e63e-6445-a5254c8823ff" {
		t.Fatalf("unexpected request id: %s", sdkErr.RequestId)
	}
}

func TestUnwrapParseJsonBusinessError_Passthrough(t *testing.T) {
	err := sdkErrors.NewCloudSDKError("FailedOperation.Other", "keep me", "rid")
	if got := unwrapParseJsonBusinessError(err); got != err {
		t.Fatalf("expected passthrough, got %#v", got)
	}
}
