package common

import (
	"testing"

	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
)

func TestBaseResponseParseErrorFromHTTPResponse_MixedCaseErrorFields(t *testing.T) {
	body := []byte(`{"Response":{"Error":{"Code":"FailedOperation.MfaBindError","Message":"bind token parameter error","code":11019,"msg":"bind token parameter error","value":null},"RequestId":"ab4ea20b-6627-e20b-a1a9-017add38b9ec"}}`)

	err := (&BaseResponse{}).ParseErrorFromHTTPResponse(body)
	if err == nil {
		t.Fatal("expected sdk error")
	}

	sdkErr, ok := err.(*sdkErrors.CloudSDKError)
	if !ok {
		t.Fatalf("expected *CloudSDKError, got %T", err)
	}
	if sdkErr.GetCode() != "FailedOperation.MfaBindError" {
		t.Fatalf("expected code FailedOperation.MfaBindError, got %q", sdkErr.GetCode())
	}
	if sdkErr.GetMessage() != "bind token parameter error" {
		t.Fatalf("expected message %q, got %q", "bind token parameter error", sdkErr.GetMessage())
	}
	if sdkErr.GetRequestId() != "ab4ea20b-6627-e20b-a1a9-017add38b9ec" {
		t.Fatalf("expected request id %q, got %q", "ab4ea20b-6627-e20b-a1a9-017add38b9ec", sdkErr.GetRequestId())
	}
}
