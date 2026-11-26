package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	publicerrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/sdk/common"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
)

type testTransport func(*http.Request) (*http.Response, error)

func (f testTransport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func testClient() *connectivity.TencentCloudClient {
	return &connectivity.TencentCloudClient{CredentialTce: common.NewCredential("test", "test"), Region: "example", Domain: "example.invalid", Protocol: "HTTPS"}
}
func requestAction(r *http.Request) string {
	for key, values := range r.Header {
		if strings.EqualFold(key, "X-TC-Action") && len(values) > 0 {
			return values[0]
		}
	}
	return ""
}

func testResponse(body string) *http.Response {
	return &http.Response{StatusCode: http.StatusOK, Header: make(http.Header), Body: io.NopCloser(strings.NewReader(body))}
}

func TestCVMDeletionProtectionRequest(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input map[string]interface{}
		want  bool
	}{
		{"omitted", map[string]interface{}{}, false},
		{"disabled", map[string]interface{}{"disable_api_termination": false}, false},
		{"enabled", map[string]interface{}{"disable_api_termination": true}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			var called bool
			conn.UseCvmClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				if requestAction(r) != "RunInstances" {
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
				called = true
				var payload map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
					return nil, err
				}
				value, exists := payload["DisableApiTermination"]
				if exists != tc.want || (tc.want && value != true) {
					t.Errorf("unexpected deletion protection value %v, present %t", value, exists)
				}
				return testResponse(`{"Response":{"Error":{"Code":"InvalidParameterValue","Message":"test sentinel"}}}`), nil
			}))
			resource := resourceTencentCloudInstance()
			data := schema.TestResourceDataRaw(t, resource.Schema, tc.input)
			err := resource.Create(data, &TencentCloudClient{apiV3Conn: conn})
			if !called || err == nil || !strings.Contains(err.Error(), "test sentinel") {
				t.Fatalf("expected intercepted RunInstances, got %v", err)
			}
		})
	}
}

func TestCAMReadFindsExactNameOnLaterPage(t *testing.T) {
	conn := testClient()
	var calls int
	conn.UseCamClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		calls++
		if requestAction(r) != "ListSubAccounts" {
			return nil, fmt.Errorf("unexpected action")
		}
		var query struct {
			Offset  uint64
			Keyword string
		}
		if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
			return nil, err
		}
		if query.Keyword != "example" || query.Offset != uint64(calls-1) {
			return nil, fmt.Errorf("unexpected filter or offset")
		}
		if calls == 1 {
			return testResponse(`{"Response":{"UserInfo":[{"Name":"example-prefix"}],"TotalNum":2}}`), nil
		}
		if calls == 2 {
			return testResponse(`{"Response":{"UserInfo":[{"Name":"example"}],"TotalNum":2}}`), nil
		}
		return nil, fmt.Errorf("unexpected page")
	}))
	service := CamService{client: conn}
	user, err := service.DescribeUserById(context.Background(), "example")
	if err != nil || user == nil || user.Name == nil || *user.Name != "example" || calls != 2 {
		t.Fatalf("unexpected account or page count, calls=%d err=%v", calls, err)
	}
}

func TestCAMUpdateIncludesUidAndUinAndRefreshes(t *testing.T) {
	conn := testClient()
	var updated, readCalled bool
	conn.UseCamClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "UpdateSubAccount":
			var req struct {
				UserInfo []struct {
					Name, Uid, Uin, Remark string
				}
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				return nil, err
			}
			if len(req.UserInfo) != 1 || req.UserInfo[0].Uid != "17" || req.UserInfo[0].Uin != "23" || req.UserInfo[0].Remark != "updated" {
				return nil, fmt.Errorf("unexpected UpdateSubAccount payload: %+v", req.UserInfo)
			}
			updated = true
			return testResponse(`{"Response":{"RequestId":"req-upd"}}`), nil
		case "ListSubAccounts":
			readCalled = true
			return testResponse(`{"Response":{"UserInfo":[{"Name":"example","Uid":17,"Uin":23,"Remark":"updated","ConsoleLogin":0}],"TotalNum":1}}`), nil
		default:
			return nil, fmt.Errorf("unexpected action %s", requestAction(r))
		}
	}))
	conn.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"Tags":[]}}`), nil
	}))
	res := resourceTencentCloudCamUser()
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"name":   "example",
		"remark": "updated",
		"uid":    17,
		"uin":    23,
	})
	d.SetId("example")
	if err := res.Update(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("unexpected update error: %v", err)
	}
	if !updated || !readCalled {
		t.Fatalf("expected UpdateSubAccount and Read refresh, updated=%t read=%t", updated, readCalled)
	}
}

func TestCAMDeleteWhenKeyLookupFails(t *testing.T) {
	for _, tc := range []struct {
		name, code                string
		useAPI, force, wantDelete bool
	}{
		{"normal api user preserves account", "FailedOperation.UserUinNotSet", true, false, false},
		{"no-api user deletes without force", "FailedOperation.UserUinNotSet", false, false, true},
		{"force delete targets account", "FailedOperation.UserUinNotSet", true, true, true},
		{"force delete propagates permission error", "UnauthorizedOperation", true, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			var deleted bool
			conn.UseCamClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				switch requestAction(r) {
				case "ListSubAccounts":
					return testResponse(`{"Response":{"UserInfo":[{"Name":"example","Uid":17,"Uin":23}],"TotalNum":1}}`), nil
				case "QueryApiKey":
					return testResponse(fmt.Sprintf(`{"Response":{"Error":{"Code":%q,"Message":"key lookup failed"}}}`, tc.code)), nil
				case "DeleteSubAccount":
					var req struct{ UserInfo []struct{ Uid, Uin int } }
					if err := json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.UserInfo) != 1 || req.UserInfo[0].Uid != 17 || req.UserInfo[0].Uin != 23 {
						return nil, fmt.Errorf("unexpected deletion target")
					}
					deleted = true
					return testResponse(`{"Response":{"RequestId":"test"}}`), nil
				default:
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
			}))
			res := resourceTencentCloudCamUser()
			d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
				"name":         "example",
				"use_api":      tc.useAPI,
				"force_delete": tc.force,
			})
			d.SetId("example")
			err := res.Delete(d, &TencentCloudClient{apiV3Conn: conn})
			if deleted != tc.wantDelete || (err == nil) != tc.wantDelete || (err != nil && !strings.Contains(err.Error(), tc.code)) {
				t.Fatalf("deleted=%t, err=%v", deleted, err)
			}
		})
	}
}

func TestRabbitMQBillingModeRequest(t *testing.T) {
	for _, tc := range []struct {
		name   string
		config map[string]interface{}
		want   int
	}{
		{"default", map[string]interface{}{}, 1},
		{"postpaid", map[string]interface{}{"pay_mode": 0}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			calls := 0
			conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				calls++
				if requestAction(r) != "CreateRabbitMQVipInstance" {
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
				var body struct{ PayMode *int }
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					return nil, err
				}
				if body.PayMode == nil || *body.PayMode != tc.want {
					t.Errorf("expected explicit PayMode %d", tc.want)
				}
				return testResponse(`{"Response":{"Error":{"Code":"ResourceUnavailable.CreateFailed","Message":"fixture unavailable"}}}`), nil
			}))
			res := resourceTencentCloudTdmqRabbitmqVipInstance()
			data := schema.TestResourceDataRaw(t, res.Schema, tc.config)
			err := res.Create(data, &TencentCloudClient{apiV3Conn: conn})
			if calls != 1 || err == nil || !strings.Contains(err.Error(), "fixture unavailable") {
				t.Fatalf("unexpected create result: calls=%d, err=%v", calls, err)
			}
		})
	}
}

func TestMariaDBSameZoneOrder(t *testing.T) {
	for _, tc := range []struct {
		nodes int
		want  []string
	}{
		{1, []string{"test-zone"}},
		{2, []string{"test-zone", "test-zone"}},
		{3, []string{"test-zone", "test-zone"}},
	} {
		t.Run(fmt.Sprint(tc.nodes), func(t *testing.T) {
			conn := testClient()
			conn.UseMariadbClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				var order struct {
					NodeCount int
					Zones     []string
				}
				if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
					return nil, err
				}
				if requestAction(r) != "CreateHourDBInstance" || order.NodeCount != tc.nodes || !reflect.DeepEqual(order.Zones, tc.want) {
					t.Errorf("unexpected MariaDB order: %+v", order)
				}
				return testResponse(`{"Response":{"Error":{"Code":"InvalidParameter","Message":"test sentinel"}}}`), nil
			}))
			res := resourceTencentCloudMariadbHourDbInstance()
			data := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
				"zones": []interface{}{"test-zone"}, "node_count": tc.nodes,
			})
			err := res.Create(data, &TencentCloudClient{apiV3Conn: conn})
			if err == nil || !strings.Contains(err.Error(), "test sentinel") {
				t.Fatalf("expected intercepted order, got %v", err)
			}
		})
	}
}

func TestCKafkaReadPreservesCatalogBandwidth(t *testing.T) {
	for _, tc := range []struct {
		name, fields string
		wantErr      bool
	}{
		{"zone list", `"ZoneIds":[1]`, false},
		{"single zone", `"ZoneId":1,"ZoneIds":[]`, false},
		{"different zone", `"ZoneId":2,"ZoneIds":[]`, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			conn.UseCkafkaClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				switch requestAction(r) {
				case "DescribeInstancesDetail":
					return testResponse(fmt.Sprintf(`{"Response":{"Result":{"InstanceList":[{"InstanceId":"test",%s,"Bandwidth":320}]}}}`, tc.fields)), nil
				case "DescribeInstanceAttributes":
					return testResponse(`{"Response":{"Result":{"MsgRetentionTime":10080}}}`), nil
				default:
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
			}))
			conn.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				return testResponse(`{"Response":{"Tags":[]}}`), nil
			}))
			res := resourceTencentCloudCkafkaInstance()
			data := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{"zone_id": 1, "band_width": 320})
			data.SetId("test")
			err := res.Read(data, &TencentCloudClient{apiV3Conn: conn})
			if tc.wantErr {
				if err == nil || !strings.Contains(err.Error(), "zone_id 1 not found") {
					t.Fatalf("unexpected read result: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got := data.Get("band_width"); got != 320 {
				t.Errorf("read returned bandwidth %v for catalog bandwidth 320", got)
			}
		})
	}
}

func TestRabbitMQReadReturnsPerNodeStorage(t *testing.T) {
	for _, tc := range []struct {
		name           string
		configuredSize int
		maxStorage     int
		want           int
	}{
		{"total storage divided by nodes", 200, 600, 200},
		{"per-node non-divisible storage does not fail", 200, 200, 200},
		{"per-node storage matching configured value", 300, 300, 300},
		{"import total storage divided by nodes", 0, 600, 200},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				switch requestAction(r) {
				case "DescribeRabbitMQVipInstance":
					return testResponse(fmt.Sprintf(`{"Response":{"ClusterInfo":{},"ClusterSpecInfo":{"NodeCount":3,"MaxStorage":%d}}}`, tc.maxStorage)), nil
				case "DescribeRabbitMQVipInstances":
					return testResponse(`{"Response":{"TotalCount":1,"Instances":[{"InstanceId":"test"}]}}`), nil
				default:
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
			}))
			res := resourceTencentCloudTdmqRabbitmqVipInstance()
			cfg := map[string]interface{}{"node_num": 3}
			if tc.configuredSize > 0 {
				cfg["storage_size"] = tc.configuredSize
			}
			data := schema.TestResourceDataRaw(t, res.Schema, cfg)
			data.SetId("test")
			if err := res.Read(data, &TencentCloudClient{apiV3Conn: conn}); err != nil {
				t.Fatal(err)
			}
			if got := data.Get("storage_size"); got != tc.want {
				t.Errorf("expected %d GB, got %v", tc.want, got)
			}
		})
	}
}

func TestRabbitMQMissingDetailRequiresConfirmedAbsence(t *testing.T) {
	for _, tc := range []struct {
		name, list string
		absent     bool
	}{
		{"absent", `{"TotalCount":0,"Instances":[]}`, true},
		{"still listed", `{"TotalCount":1,"Instances":[{"InstanceId":"example"}]}`, false},
		{"missing count", `{"Instances":[]}`, false},
		{"missing instances", `{"TotalCount":1}`, false},
		{"list failed", `{"Error":{"Code":"UnauthorizedOperation","Message":"denied"}}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			conn := testClient()
			conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				switch requestAction(r) {
				case "DescribeRabbitMQVipInstance":
					return testResponse(`{"Response":{"Error":{"Code":"FailedOperation","Message":"not query instance"}}}`), nil
				case "DescribeRabbitMQVipInstances":
					var body struct {
						Filters []struct {
							Name   string
							Values []string
						}
					}
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil || len(body.Filters) != 1 || body.Filters[0].Name != "instanceIds" || !reflect.DeepEqual(body.Filters[0].Values, []string{"example"}) {
						t.Fatal("list query must match the exact instance ID")
					}
					return testResponse(`{"Response":` + tc.list + `}`), nil
				default:
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
			}))
			res := resourceTencentCloudTdmqRabbitmqVipInstance()
			d := schema.TestResourceDataRaw(t, res.Schema, nil)
			d.SetId("example")
			err := res.Read(d, &TencentCloudClient{apiV3Conn: conn})
			if (err == nil) != tc.absent || (d.Id() == "") != tc.absent {
				t.Fatalf("id=%q err=%v", d.Id(), err)
			}
		})
	}
}

func TestRetryErrorHandlesPublicSDKError(t *testing.T) {
	rateLimitErr := publicerrors.NewTencentCloudSDKError("RequestLimitExceeded", "too many requests", "req-1")
	if r := retryError(rateLimitErr); !r.Retryable {
		t.Fatalf("expected RequestLimitExceeded to be retryable")
	}
	internalErr := publicerrors.NewTencentCloudSDKError(InternalError, "internal", "req-2")
	if r := retryError(internalErr, InternalError); !r.Retryable {
		t.Fatalf("expected additional InternalError to be retryable")
	}
	invalidErr := publicerrors.NewTencentCloudSDKError("InvalidParameter", "bad param", "req-3")
	if r := retryError(invalidErr); r.Retryable {
		t.Fatalf("expected InvalidParameter to be non-retryable")
	}
}

func TestCKafkaImportDeleteAndFailedStatus(t *testing.T) {
	t.Run("import infers zone_id from API", func(t *testing.T) {
		conn := testClient()
		conn.UseCkafkaClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			switch requestAction(r) {
			case "DescribeInstancesDetail":
				return testResponse(`{"Response":{"Result":{"InstanceList":[{"InstanceId":"ckafka-1","InstanceName":"demo","ZoneId":100001,"ZoneIds":[],"Bandwidth":320,"DiskSize":300,"Status":1}]}}}`), nil
			case "DescribeInstanceAttributes":
				return testResponse(`{"Response":{"Result":{"MsgRetentionTime":10080}}}`), nil
			default:
				return nil, fmt.Errorf("unexpected action %s", requestAction(r))
			}
		}))
		conn.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			return testResponse(`{"Response":{"Tags":[]}}`), nil
		}))
		res := resourceTencentCloudCkafkaInstance()
		d := schema.TestResourceDataRaw(t, res.Schema, nil)
		d.SetId("ckafka-1")
		if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("import read failed: %v", err)
		}
		if got := d.Get("zone_id"); got != 100001 {
			t.Fatalf("expected inferred zone_id 100001, got %v", got)
		}
	})

	t.Run("delete treats InstanceNotExist as already gone", func(t *testing.T) {
		conn := testClient()
		conn.UseCkafkaClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			switch requestAction(r) {
			case "DeleteInstance":
				return testResponse(`{"Response":{"Error":{"Code":"InvalidParameterValue.InstanceNotExist","Message":"not exist"}}}`), nil
			case "DescribeInstancesDetail":
				return testResponse(`{"Response":{"Error":{"Code":"InvalidParameterValue.InstanceNotExist","Message":"not exist"}}}`), nil
			default:
				return nil, fmt.Errorf("unexpected action %s", requestAction(r))
			}
		}))
		res := resourceTencentCloudCkafkaInstance()
		d := schema.TestResourceDataRaw(t, res.Schema, nil)
		d.SetId("ckafka-gone")
		if err := res.Delete(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("expected idempotent delete, got %v", err)
		}
	})

	t.Run("create failed status fails immediately", func(t *testing.T) {
		conn := testClient()
		conn.UseCkafkaClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			return testResponse(`{"Response":{"Result":{"InstanceList":[{"InstanceId":"ckafka-bad","Status":-1}]}}}`), nil
		}))
		svc := CkafkaService{client: conn}
		_, ready, err := svc.CheckCkafkaInstanceReady(context.Background(), "ckafka-bad")
		if err == nil || ready {
			t.Fatalf("expected immediate error for status -1, got ready=%t err=%v", ready, err)
		}
	})
}

func TestRabbitMQDeleteHandlesTerminalStatusesAndNoDuplicateOrder(t *testing.T) {
	for _, status := range []int{RabbitMQVipInstanceIsolated, RabbitMQVipInstanceDestroyed} {
		t.Run(fmt.Sprintf("delete_status_%d", status), func(t *testing.T) {
			conn := testClient()
			conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				switch requestAction(r) {
				case "DeleteRabbitMQVipInstance":
					return testResponse(`{"Response":{"RequestId":"req-del"}}`), nil
				case "DescribeRabbitMQVipInstances":
					return testResponse(fmt.Sprintf(`{"Response":{"TotalCount":1,"Instances":[{"InstanceId":"rmq-1","Status":%d}]}}`, status)), nil
				default:
					return nil, fmt.Errorf("unexpected action %s", requestAction(r))
				}
			}))
			res := resourceTencentCloudTdmqRabbitmqVipInstance()
			d := schema.TestResourceDataRaw(t, res.Schema, nil)
			d.SetId("rmq-1")
			if err := res.Delete(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
				t.Fatalf("expected delete to finish for status %d, got %v", status, err)
			}
		})
	}

	t.Run("read clears destroyed status", func(t *testing.T) {
		conn := testClient()
		conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			return testResponse(`{"Response":{"ClusterInfo":{"ClusterStatus":3},"ClusterSpecInfo":{"NodeCount":3,"MaxStorage":600}}}`), nil
		}))
		res := resourceTencentCloudTdmqRabbitmqVipInstance()
		d := schema.TestResourceDataRaw(t, res.Schema, nil)
		d.SetId("rmq-gone")
		if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("unexpected read err: %v", err)
		}
		if d.Id() != "" {
			t.Fatalf("expected cleared id for destroyed cluster, got %q", d.Id())
		}
	})

	t.Run("create does not retry generic FailedOperation", func(t *testing.T) {
		conn := testClient()
		calls := 0
		conn.UseTdmqClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			calls++
			return testResponse(`{"Response":{"Error":{"Code":"FailedOperation","Message":"order state unknown"}}}`), nil
		}))
		res := resourceTencentCloudTdmqRabbitmqVipInstance()
		d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{"pay_mode": 0})
		err := res.Create(d, &TencentCloudClient{apiV3Conn: conn})
		if err == nil || calls != 1 {
			t.Fatalf("expected single create attempt without duplicate retry, calls=%d err=%v", calls, err)
		}
	})
}

func TestTbasePGLifecycleIsolateReleaseAndReadAbsent(t *testing.T) {
	t.Run("delete isolates and permanently releases", func(t *testing.T) {
		conn := testClient()
		stage := "online"
		var isolatedCalled, releasedCalled bool
		conn.UseTbaseClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			switch requestAction(r) {
			case "DescribePGInstances":
				if stage == "deleted" {
					return testResponse(`{"Response":{"Items":[]}}`), nil
				}
				return testResponse(fmt.Sprintf(`{"Response":{"Items":[{"InstanceId":"pg-1","Status":%q}]}}`, stage)), nil
			case "IsolatePGInstanceHour":
				isolatedCalled = true
				stage = "isolate"
				return testResponse(`{"Response":{"RequestId":"req-iso"}}`), nil
			case "ReleaseIsolatedPGInstance":
				releasedCalled = true
				stage = "deleted"
				return testResponse(`{"Response":{"RequestId":"req-rel"}}`), nil
			default:
				return nil, fmt.Errorf("unexpected action %s", requestAction(r))
			}
		}))
		res := resourceTencentCloudTbasePGInstance()
		d := schema.TestResourceDataRaw(t, res.Schema, nil)
		d.SetId("pg-1")
		if diags := res.DeleteContext(context.Background(), d, &TencentCloudClient{apiV3Conn: conn}); diags.HasError() {
			t.Fatalf("unexpected delete diagnostics: %v", diags)
		}
		if !isolatedCalled || !releasedCalled {
			t.Fatalf("expected both isolate and release calls, got isolated=%t released=%t", isolatedCalled, releasedCalled)
		}
	})

	t.Run("read clears state when isolated or missing", func(t *testing.T) {
		for _, resp := range []string{
			`{"Response":{"Items":[]}}`,
			`{"Response":{"Items":[{"InstanceId":"pg-1","Status":"isolate"}]}}`,
		} {
			conn := testClient()
			conn.UseTbaseClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
				return testResponse(resp), nil
			}))
			res := resourceTencentCloudTbasePGInstance()
			d := schema.TestResourceDataRaw(t, res.Schema, nil)
			d.SetId("pg-1")
			if diags := res.ReadContext(context.Background(), d, &TencentCloudClient{apiV3Conn: conn}); diags.HasError() {
				t.Fatalf("expected clean absent read for %s, got %v", resp, diags)
			}
			if d.Id() != "" {
				t.Fatalf("expected cleared id for %s, got %q", resp, d.Id())
			}
		}
	})
}

func TestCVMStoragePoolGroupCreateAndRead(t *testing.T) {
	conn := testClient()
	var runReq map[string]interface{}
	var poolGroupDescribeMode string

	cvmTransport := testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "RunInstances":
			body, _ := io.ReadAll(r.Body)
			if err := json.Unmarshal(body, &runReq); err != nil {
				return nil, err
			}
			return testResponse(`{"Response":{"InstanceIdSet":["ins-pool-1"],"RequestId":"req-run"}}`), nil
		case "DescribeInstances":
			sysPool := ""
			dataPool := ""
			switch poolGroupDescribeMode {
			case "populated":
				sysPool = `,"DiskStoragePoolGroup":"pool-sys-1"`
				dataPool = `,"DiskStoragePoolGroup":"pool-data-1"`
			case "empty":
				sysPool = `,"DiskStoragePoolGroup":""`
				dataPool = `,"DiskStoragePoolGroup":""`
			}
			return testResponse(fmt.Sprintf(`{"Response":{"TotalCount":1,"InstanceSet":[{"InstanceId":"ins-pool-1","InstanceName":"demo","InstanceState":"RUNNING","InstanceChargeType":"POSTPAID_BY_HOUR","InstanceType":"S5.MEDIUM2","ImageId":"img-1","Placement":{"Zone":"zone-1","ProjectId":0},"VirtualPrivateCloud":{"VpcId":"vpc-1","SubnetId":"subnet-1"},"InternetAccessible":{"InternetChargeType":"TRAFFIC_POSTPAID_BY_HOUR","InternetMaxBandwidthOut":0},"LoginSettings":{},"SystemDisk":{"DiskType":"CLOUD_PREMIUM","DiskSize":50,"DiskId":"disk-sys-1"%s},"DataDisks":[{"DiskType":"CLOUD_PREMIUM","DiskSize":100,"DiskId":"disk-data-1","DeleteWithInstance":true%s}]}],"RequestId":"req-desc"}}`, sysPool, dataPool)), nil
		case "DescribeImages":
			return testResponse(`{"Response":{"TotalCount":1,"ImageSet":[{"ImageId":"img-1","ImageName":"Demo Image","OsName":"TencentOS Server 3.1"}],"RequestId":"req-img"}}`), nil
		default:
			return nil, fmt.Errorf("unexpected cvm action %s", requestAction(r))
		}
	})
	conn.UseCvmClient().WithHttpTransport(cvmTransport)
	conn.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"Tags":[],"RequestId":"req-tag"}}`), nil
	}))
	conn.UseCbsClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"TotalCount":1,"DiskSet":[{"DiskId":"disk-data-1","DiskSize":100,"DiskState":"ATTACHED"}],"RequestId":"req-cbs"}}`), nil
	}))

	res := resourceTencentCloudInstance()
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"instance_name":                  "demo",
		"availability_zone":              "zone-1",
		"image_id":                       "img-1",
		"instance_type":                  "S5.MEDIUM2",
		"vpc_id":                         "vpc-1",
		"subnet_id":                      "subnet-1",
		"running_flag":                   true,
		"system_disk_type":               "CLOUD_PREMIUM",
		"system_disk_size":               50,
		"system_disk_storage_pool_group": "pool-sys-1",
		"data_disks": []interface{}{
			map[string]interface{}{
				"data_disk_type":               "CLOUD_PREMIUM",
				"data_disk_size":               100,
				"data_disk_storage_pool_group": "pool-data-1",
				"delete_with_instance":         true,
			},
		},
	})

	poolGroupDescribeMode = "populated"
	if err := res.Create(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}

	sysDisk, ok := runReq["SystemDisk"].(map[string]interface{})
	if !ok || sysDisk["DiskStoragePoolGroup"] != "pool-sys-1" {
		t.Fatalf("expected SystemDisk.DiskStoragePoolGroup=pool-sys-1, got %#v", runReq["SystemDisk"])
	}
	dataDisks, ok := runReq["DataDisks"].([]interface{})
	if !ok || len(dataDisks) != 1 {
		t.Fatalf("expected 1 DataDisk in RunInstances, got %#v", runReq["DataDisks"])
	}
	firstDataDisk, _ := dataDisks[0].(map[string]interface{})
	if firstDataDisk["DiskStoragePoolGroup"] != "pool-data-1" {
		t.Fatalf("expected DataDisks[0].DiskStoragePoolGroup=pool-data-1, got %#v", firstDataDisk)
	}
	if got := d.Get("system_disk_storage_pool_group").(string); got != "pool-sys-1" {
		t.Fatalf("expected state system_disk_storage_pool_group=pool-sys-1, got %q", got)
	}
	if got := d.Get("data_disks.0.data_disk_storage_pool_group").(string); got != "pool-data-1" {
		t.Fatalf("expected state data_disks.0.data_disk_storage_pool_group=pool-data-1, got %q", got)
	}

	// Verify Read preserves storage pool groups even if backend omits them or returns empty strings.
	for _, mode := range []string{"omitted", "empty"} {
		poolGroupDescribeMode = mode
		if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("unexpected read error (%s): %v", mode, err)
		}
		if got := d.Get("system_disk_storage_pool_group").(string); got != "pool-sys-1" {
			t.Fatalf("expected preserved system_disk_storage_pool_group=pool-sys-1 (%s), got %q", mode, got)
		}
		if got := d.Get("data_disks.0.data_disk_storage_pool_group").(string); got != "pool-data-1" {
			t.Fatalf("expected preserved data_disks.0.data_disk_storage_pool_group=pool-data-1 (%s), got %q", mode, got)
		}
	}
}

func TestCVMImageDataSourcesExposeOsKey(t *testing.T) {
	conn := testClient()
	conn.UseCvmClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		if requestAction(r) != "DescribeImages" {
			return nil, fmt.Errorf("unexpected action %s", requestAction(r))
		}
		return testResponse(`{"Response":{"TotalCount":2,"ImageSet":[{"ImageId":"img-new","ImageName":"TencentOS Server 3.1","OsName":"TencentOS Server 3.1 (TK4)","OsKey":"tlinux3","ImageType":"PUBLIC_IMAGE","CreatedTime":"2026-01-02T00:00:00Z","ImageSize":20,"Architecture":"x86_64","ImageState":"NORMAL","Platform":"TencentOS"},{"ImageId":"img-old","ImageName":"Legacy Image","OsName":"CentOS 7.9","ImageType":"PUBLIC_IMAGE","CreatedTime":"2025-01-01T00:00:00Z","ImageSize":20,"Architecture":"x86_64","ImageState":"NORMAL","Platform":"CentOS"}],"RequestId":"req-img"}}`), nil
	}))

	t.Run("cvm_image", func(t *testing.T) {
		ds := dataSourceTencentCloudImage()
		d := schema.TestResourceDataRaw(t, ds.Schema, map[string]interface{}{
			"os_name": "TencentOS",
		})
		if err := ds.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
		if got := d.Get("image_id").(string); got != "img-new" {
			t.Fatalf("expected image_id=img-new, got %q", got)
		}
		if got := d.Get("os_key").(string); got != "tlinux3" {
			t.Fatalf("expected os_key=tlinux3, got %q", got)
		}
	})

	t.Run("cvm_images", func(t *testing.T) {
		ds := dataSourceTencentCloudImages()
		d := schema.TestResourceDataRaw(t, ds.Schema, map[string]interface{}{
			"image_type": []interface{}{"PUBLIC_IMAGE"},
		})
		if err := ds.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("unexpected read error: %v", err)
		}
		if got := d.Get("images.#").(int); got != 2 {
			t.Fatalf("expected 2 images, got %d", got)
		}
		if got := d.Get("images.0.os_key").(string); got != "tlinux3" {
			t.Fatalf("expected images.0.os_key=tlinux3, got %q", got)
		}
		if got := d.Get("images.1.os_key").(string); got != "" {
			t.Fatalf("expected images.1.os_key empty when nil, got %q", got)
		}
	})
}

func TestCAMListSubAccountsFlexibleUnmarshal(t *testing.T) {
	cases := []struct {
		name         string
		payload      string
		wantTotal    *uint64
		wantOwners   int
		wantOwnerUin uint64
	}{
		{
			name:         "numeric TotalNum and array OwnerInfo",
			payload:      `{"Response":{"TotalNum":2,"UserInfo":[{"Uid":1,"Uin":1001,"Name":"alice"}],"OwnerInfo":[{"Uin":9001,"Name":"root"}],"RequestId":"r1"}}`,
			wantTotal:    func() *uint64 { v := uint64(2); return &v }(),
			wantOwners:   1,
			wantOwnerUin: 9001,
		},
		{
			name:         "string TotalNum and single object OwnerInfo",
			payload:      `{"Response":{"TotalNum":"1","UserInfo":[{"Uid":2,"Uin":1002,"Name":"bob"}],"OwnerInfo":{"Uin":9002,"Name":"root2"},"RequestId":"r2"}}`,
			wantTotal:    func() *uint64 { v := uint64(1); return &v }(),
			wantOwners:   1,
			wantOwnerUin: 9002,
		},
		{
			name:       "empty string TotalNum and empty object OwnerInfo",
			payload:    `{"Response":{"TotalNum":"","UserInfo":[],"OwnerInfo":{},"RequestId":"r3"}}`,
			wantTotal:  nil,
			wantOwners: 0,
		},
		{
			name:       "null TotalNum and null OwnerInfo",
			payload:    `{"Response":{"TotalNum":null,"UserInfo":[],"OwnerInfo":null,"RequestId":"r4"}}`,
			wantTotal:  nil,
			wantOwners: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp := cam.NewListSubAccountsResponse()
			if err := resp.FromJsonString(tc.payload); err != nil {
				t.Fatalf("FromJsonString failed: %v", err)
			}
			if resp.Response == nil {
				t.Fatalf("expected non-nil Response")
			}
			if tc.wantTotal == nil {
				if resp.Response.TotalNum != nil {
					t.Fatalf("expected nil TotalNum, got %d", *resp.Response.TotalNum)
				}
			} else if resp.Response.TotalNum == nil || *resp.Response.TotalNum != *tc.wantTotal {
				t.Fatalf("expected TotalNum=%d, got %v", *tc.wantTotal, resp.Response.TotalNum)
			}
			if len(resp.Response.OwnerInfo) != tc.wantOwners {
				t.Fatalf("expected %d OwnerInfo items, got %d", tc.wantOwners, len(resp.Response.OwnerInfo))
			}
			if tc.wantOwners > 0 {
				if resp.Response.OwnerInfo[0] == nil || resp.Response.OwnerInfo[0].Uin == nil || *resp.Response.OwnerInfo[0].Uin != tc.wantOwnerUin {
					t.Fatalf("unexpected OwnerInfo[0]: %+v", resp.Response.OwnerInfo[0])
				}
			}
		})
	}

	// Verify end-to-end data source read when the server returns string TotalNum and object OwnerInfo.
	conn := testClient()
	conn.UseCamClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "ListSubAccounts":
			return testResponse(`{"Response":{"TotalNum":"1","UserInfo":[{"Uid":7,"Uin":10007,"Name":"demo-user","Remark":"r","ConsoleLogin":1}],"OwnerInfo":{"Uin":10000,"Name":"owner"},"RequestId":"r5"}}`), nil
		default:
			return testResponse(`{"Response":{"RequestId":"r-other"}}`), nil
		}
	}))
	ds := dataSourceTencentCloudCamUsers()
	d := schema.TestResourceDataRaw(t, ds.Schema, map[string]interface{}{})
	if err := ds.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("cam_users data source Read failed: %v", err)
	}
	if got := d.Get("user_list.#").(int); got != 1 {
		t.Fatalf("expected 1 user, got %d", got)
	}
	if got := d.Get("user_list.0.name").(string); got != "demo-user" {
		t.Fatalf("expected user_list.0.name=demo-user, got %q", got)
	}
}

func TestMariaDBCpuArchCreateAndRead(t *testing.T) {
	conn := testClient()
	var seenCpuArch string
	conn.UseMariadbClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "CreateHourDBInstance":
			var body struct {
				CpuArch string
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				return nil, err
			}
			seenCpuArch = body.CpuArch
			return testResponse(`{"Response":{"Error":{"Code":"InvalidParameter","Message":"test sentinel"}}}`), nil
		case "DescribeDBInstanceDetail":
			return testResponse(`{"Response":{"InstanceId":"tdsql-arm1","InstanceName":"db-arm","Zone":"zone-a","SlaveZones":["zone-a"],"NodeCount":2,"Memory":4,"Storage":50,"ProjectId":0,"Vip":"10.0.0.9","VpcId":"vpc-1","SubnetId":"subnet-1","DbVersion":"10.1.9","CpuArch":"ARM","RequestId":"r1"}}`), nil
		default:
			return testResponse(`{"Response":{"RequestId":"r-other"}}`), nil
		}
	}))
	conn.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"Rows":[],"RequestId":"r-tag"}}`), nil
	}))

	res := resourceTencentCloudMariadbHourDbInstance()
	createData := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"zones":      []interface{}{"zone-a"},
		"node_count": 2,
		"memory":     4,
		"storage":    50,
		"cpu_arch":   "ARM",
	})
	err := res.Create(createData, &TencentCloudClient{apiV3Conn: conn})
	if err == nil || !strings.Contains(err.Error(), "test sentinel") {
		t.Fatalf("expected intercepted create, got %v", err)
	}
	if seenCpuArch != "ARM" {
		t.Fatalf("expected CreateHourDBInstance CpuArch=ARM, got %q", seenCpuArch)
	}

	readData := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{})
	readData.SetId("tdsql-arm1")
	if err := res.Read(readData, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if got := readData.Get("cpu_arch").(string); got != "ARM" {
		t.Fatalf("expected cpu_arch=ARM, got %q", got)
	}
	if got := readData.Get("db_version_id").(string); got != "10.1.9" {
		t.Fatalf("expected db_version_id=10.1.9, got %q", got)
	}

	// Verify configured minor-less db_version_id (e.g. "10.4") does not drift when DescribeDBInstanceDetail
	// only returns full DbVersion ("10.4.12"), and lowercase cpu_arch normalizes to uppercase.
	conn2 := testClient()
	conn2.UseMariadbClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"InstanceId":"tdsql-104","InstanceName":"db-104","Zone":"zone-a","SlaveZones":["zone-a"],"NodeCount":2,"Memory":2,"Storage":10,"ProjectId":0,"Vip":"10.0.0.10","VpcId":"vpc-1","SubnetId":"subnet-1","DbVersion":"10.4.12","CpuArch":"arm","RequestId":"r2"}}`), nil
	}))
	conn2.UseTagClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		return testResponse(`{"Response":{"Rows":[],"RequestId":"r-tag"}}`), nil
	}))
	preserveData := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"db_version_id": "10.4",
		"cpu_arch":      "ARM",
	})
	preserveData.SetId("tdsql-104")
	if err := res.Read(preserveData, &TencentCloudClient{apiV3Conn: conn2}); err != nil {
		t.Fatalf("Read with configured db_version_id failed: %v", err)
	}
	if got := preserveData.Get("db_version_id").(string); got != "10.4" {
		t.Fatalf("expected db_version_id=10.4 preserved without drift, got %q", got)
	}
	if got := preserveData.Get("cpu_arch").(string); got != "ARM" {
		t.Fatalf("expected normalized cpu_arch=ARM, got %q", got)
	}
}

func TestProviderInternalValidateWithTKE(t *testing.T) {
	p := Provider()
	if err := p.InternalValidate(); err != nil {
		t.Fatalf("Provider().InternalValidate() failed: %v", err)
	}
	for _, resName := range []string{
		"tencentcloudenterprise_tke_kubernetes_cluster",
		"tencentcloudenterprise_tke_kubernetes_scale_worker",
		"tencentcloudenterprise_tke_kubernetes_cluster_attachment",
		"tencentcloudenterprise_tke_kubernetes_cluster_endpoint",
	} {
		if _, ok := p.ResourcesMap[resName]; !ok {
			t.Fatalf("expected resource %q to be registered", resName)
		}
	}
	if _, ok := p.DataSourcesMap["tencentcloudenterprise_tke_kubernetes_clusters"]; !ok {
		t.Fatalf("expected data source tencentcloudenterprise_tke_kubernetes_clusters to be registered")
	}
}

func TestTKEClusterNetworkTypeAndKubeconfig(t *testing.T) {
	t.Run("parseTkeClusterNetworkType", func(t *testing.T) {
		for _, tc := range []struct {
			name    string
			prop    *string
			want    string
			wantErr bool
		}{
			{"nil property", nil, "", false},
			{"empty property", new(string), "", false},
			{"GR", func() *string { s := `{"NodeNameType":"lan-ip","NetworkType":"GR"}`; return &s }(), "GR", false},
			{"VPC-CNI", func() *string { s := `{"NetworkType":"VPC-CNI"}`; return &s }(), "VPC-CNI", false},
			{"invalid json", func() *string { s := `{bad`; return &s }(), "", true},
		} {
			t.Run(tc.name, func(t *testing.T) {
				got, err := parseTkeClusterNetworkType(tc.prop)
				if (err != nil) != tc.wantErr {
					t.Fatalf("err=%v, wantErr=%t", err, tc.wantErr)
				}
				if got != tc.want {
					t.Fatalf("got=%q, want=%q", got, tc.want)
				}
			})
		}
	})

	t.Run("cluster read and data source populate network_type and kubeconfig", func(t *testing.T) {
		conn := testClient()
		var seenKubeTypes []string
		conn.UseTkeClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
			switch requestAction(r) {
			case "DescribeClusters":
				return testResponse(`{"Response":{"TotalCount":1,"Clusters":[{"ClusterId":"cls-test1","ClusterName":"demo","ClusterDescription":"desc","ClusterVersion":"1.24.4","ClusterOs":"centos7.6.0_x64","ClusterType":"MANAGED_CLUSTER","ClusterNetworkSettings":{"ClusterCIDR":"172.24.0.0/16","IgnoreClusterCIDRConflict":false,"MaxNodePodNum":32,"MaxClusterServiceNum":32,"Ipvs":true,"VpcId":"vpc-1","Cni":true},"ClusterNodeNum":1,"ProjectId":0,"ClusterStatus":"Running","Property":"{\"NodeNameType\":\"lan-ip\",\"NetworkType\":\"GR\"}","ClusterArch":"x86","CreatedTime":"2026-01-01T00:00:00Z"}],"RequestId":"r1"}}`), nil
			case "DescribeClusterKubeconfig":
				var body struct{ Type string }
				_ = json.NewDecoder(r.Body).Decode(&body)
				seenKubeTypes = append(seenKubeTypes, body.Type)
				return testResponse(fmt.Sprintf(`{"Response":{"Kubeconfig":"apiVersion: v1\nclusters:\n- name: %s\n","RequestId":"r2"}}`, body.Type)), nil
			case "DescribeClusterInstances":
				return testResponse(`{"Response":{"TotalCount":1,"InstanceSet":[{"InstanceId":"ins-1","InstanceRole":"WORKER","InstanceState":"running","FailedReason":""}],"RequestId":"r3"}}`), nil
			case "DescribeClusterSecurity":
				return testResponse(`{"Response":{"UserName":"admin","Password":"secret","CertificationAuthority":"ca-cert","ClusterExternalEndpoint":"1.2.3.4","Domain":"cls-test1.ccs.tencent-cloud.com","PgwEndpoint":"10.0.0.10","SecurityPolicy":["0.0.0.0/0"],"RequestId":"r4"}}`), nil
			case "DescribeClusterAsGroupOption":
				return testResponse(`{"Response":{"ClusterAsGroupOption":{"IsScaleDownEnabled":false,"Expander":"random"},"RequestId":"r5"}}`), nil
			default:
				return nil, fmt.Errorf("unexpected tke action %s", requestAction(r))
			}
		}))

		// Resource Read
		res := resourceTencentCloudTkeCluster()
		d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
			"cluster_os": "centos7.6.0_x64",
		})
		d.SetId("cls-test1")
		if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("cluster Read failed: %v", err)
		}
		if got := d.Get("network_type").(string); got != "GR" {
			t.Fatalf("expected network_type=GR (not VPC-CNI despite Cni=true), got %q", got)
		}
		if got := d.Get("cluster_arch").(string); got != "x86" {
			t.Fatalf("expected cluster_arch=x86, got %q", got)
		}
		if got := d.Get("kube_config").(string); !strings.Contains(got, "INTERNETLB") {
			t.Fatalf("expected kube_config from INTERNETLB, got %q", got)
		}
		if got := d.Get("kube_config_intranet").(string); !strings.Contains(got, "INNERLB") {
			t.Fatalf("expected kube_config_intranet from INNERLB, got %q", got)
		}

		// Data Source Read
		ds := dataSourceTencentCloudKubernetesClusters()
		dsData := schema.TestResourceDataRaw(t, ds.Schema, map[string]interface{}{
			"cluster_id": "cls-test1",
		})
		if err := ds.Read(dsData, &TencentCloudClient{apiV3Conn: conn}); err != nil {
			t.Fatalf("clusters data source Read failed: %v", err)
		}
		if got := dsData.Get("list.0.network_type").(string); got != "GR" {
			t.Fatalf("expected ds network_type=GR, got %q", got)
		}
		if got := dsData.Get("list.0.kube_config").(string); !strings.Contains(got, "INTERNETLB") {
			t.Fatalf("expected ds kube_config from INTERNETLB, got %q", got)
		}
		if got := dsData.Get("list.0.kube_config_intranet").(string); !strings.Contains(got, "INNERLB") {
			t.Fatalf("expected ds kube_config_intranet from INNERLB, got %q", got)
		}
		if got := dsData.Get("list.0.user_name").(string); got != "admin" {
			t.Fatalf("expected ds user_name=admin, got %q", got)
		}
		if len(seenKubeTypes) < 4 {
			t.Fatalf("expected INTERNETLB and INNERLB calls for resource + data source, got %v", seenKubeTypes)
		}
	})
}

func TestTKENodeStoragePoolGroupAndPreStartScript(t *testing.T) {
	cvmJson, count, err := tkeGetCvmRunInstancesPara(map[string]interface{}{
		"count":                          2,
		"availability_zone":              "zone-1",
		"instance_type":                  "S3.LARGE8",
		"subnet_id":                      "subnet-1",
		"system_disk_type":               "CLOUD_HSSD",
		"system_disk_size":               60,
		"system_disk_storage_pool_group": "pool-sys-tke",
		"data_disk": []interface{}{
			map[string]interface{}{
				"disk_type":               "CLOUD_SSD",
				"disk_size":               100,
				"disk_storage_pool_group": "pool-data-tke",
			},
		},
		"public_ip_assigned": false,
		"password":           "TkeTest@2026",
	}, nil, "vpc-1", 0)
	if err != nil {
		t.Fatalf("tkeGetCvmRunInstancesPara failed: %v", err)
	}
	if count != 2 {
		t.Fatalf("expected count=2, got %d", count)
	}
	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(cvmJson), &payload); err != nil {
		t.Fatalf("invalid RunInstances JSON: %v", err)
	}
	sysDisk, _ := payload["SystemDisk"].(map[string]interface{})
	if sysDisk["DiskStoragePoolGroup"] != "pool-sys-tke" || sysDisk["DiskType"] != "CLOUD_HSSD" {
		t.Fatalf("unexpected SystemDisk payload: %#v", sysDisk)
	}
	dataDisks, _ := payload["DataDisks"].([]interface{})
	if len(dataDisks) != 1 || dataDisks[0].(map[string]interface{})["DiskStoragePoolGroup"] != "pool-data-tke" {
		t.Fatalf("unexpected DataDisks payload: %#v", dataDisks)
	}

	adv := tkeGetInstanceAdvancedPara(map[string]interface{}{
		"pre_start_user_script": "ZWNobyBoZWxsbw==",
		"user_data":             "ZWNobyB3b3JsZA==",
	}, nil)
	if adv.PreStartUserScript == nil || *adv.PreStartUserScript != "ZWNobyBoZWxsbw==" {
		t.Fatalf("expected PreStartUserScript to be populated, got %#v", adv.PreStartUserScript)
	}
}

func TestTKEClusterEndpointAntiDriftAndUpdate(t *testing.T) {
	res := resourceTencentCloudTkeClusterEndpoint()
	for _, field := range []string{
		"cluster_internet_security_group",
		"cluster_internet_domain",
		"cluster_intranet_domain",
		"cluster_intranet_subnet_id",
	} {
		s := res.Schema[field]
		if s == nil || !s.Optional || !s.Computed {
			t.Fatalf("expected %s to be Optional+Computed to prevent repeated apply drift", field)
		}
	}
	if np := res.Schema["node_pool_id"]; np == nil || np.Required {
		t.Fatalf("expected node_pool_id to be Optional in 3.10.0")
	}

	// Ensure ModifyClusterAttribute update does not panic when DescribeCluster returns nil ClusterLevel.
	conn := testClient()
	var modifiedName string
	conn.UseTkeClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "DescribeClusters":
			return testResponse(`{"Response":{"TotalCount":1,"Clusters":[{"ClusterId":"cls-upd","ClusterName":"old","ClusterDescription":"old","ClusterVersion":"1.24.4","ClusterOs":"centos7.6.0_x64","ClusterType":"MANAGED_CLUSTER","ClusterNetworkSettings":{"ClusterCIDR":"172.24.0.0/16","IgnoreClusterCIDRConflict":false,"MaxNodePodNum":32,"MaxClusterServiceNum":32,"Ipvs":true,"VpcId":"vpc-1"},"ClusterNodeNum":0,"ProjectId":0,"ClusterStatus":"Running","Property":"{\"NetworkType\":\"GR\"}"}],"RequestId":"r1"}}`), nil
		case "ModifyClusterAttribute":
			var body struct{ ClusterName string }
			_ = json.NewDecoder(r.Body).Decode(&body)
			modifiedName = body.ClusterName
			return testResponse(`{"Response":{"RequestId":"r2"}}`), nil
		case "DescribeClusterKubeconfig":
			return testResponse(`{"Response":{"Kubeconfig":"apiVersion: v1\nclusters: []\n","RequestId":"r3"}}`), nil
		case "DescribeClusterInstances":
			return testResponse(`{"Response":{"TotalCount":0,"InstanceSet":[],"RequestId":"r4"}}`), nil
		case "DescribeClusterSecurity":
			return testResponse(`{"Response":{"UserName":"admin","RequestId":"r5"}}`), nil
		case "DescribeClusterAsGroupOption":
			return testResponse(`{"Response":{"ClusterAsGroupOption":{"IsScaleDownEnabled":false},"RequestId":"r6"}}`), nil
		case "DescribeAvailableClusterVersion":
			return testResponse(`{"Response":{"Versions":["1.10.5","1.24.4"],"RequestId":"r7"}}`), nil
		case "UpdateClusterVersion":
			return testResponse(`{"Response":{"RequestId":"r8"}}`), nil
		default:
			return nil, fmt.Errorf("unexpected action %s", requestAction(r))
		}
	}))

	clusterRes := resourceTencentCloudTkeCluster()
	updData := schema.TestResourceDataRaw(t, clusterRes.Schema, map[string]interface{}{
		"cluster_name": "renamed",
		"cluster_desc": "new desc",
	})
	updData.SetId("cls-upd")
	if err := clusterRes.Update(updData, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("cluster Update should not panic or fail when ClusterLevel is nil: %v", err)
	}
	if modifiedName != "renamed" {
		t.Fatalf("expected ModifyClusterAttribute ClusterName=renamed, got %q", modifiedName)
	}
}

func TestTKEClusterCreateFastFailOnProgressError(t *testing.T) {
	conn := testClient()
	conn.UseTkeClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "CreateCluster":
			return testResponse(`{"Response":{"ClusterId":"cls-fail","RequestId":"r1"}}`), nil
		case "DescribeClusters":
			return testResponse(`{"Response":{"TotalCount":1,"Clusters":[{"ClusterId":"cls-fail","ClusterName":"fail-test","ClusterVersion":"1.24.4","ClusterOs":"centos7.6.0_x64","ClusterType":"MANAGED_CLUSTER","ClusterNetworkSettings":{"ClusterCIDR":"172.24.0.0/16","IgnoreClusterCIDRConflict":false,"MaxNodePodNum":32,"MaxClusterServiceNum":32,"Ipvs":true,"VpcId":"vpc-1"},"ClusterNodeNum":0,"ProjectId":0,"ClusterStatus":"Initializing","Property":"{\"NetworkType\":\"GR\"}"}],"RequestId":"r2"}}`), nil
		case "DescribeClusterCreateProgress":
			return testResponse(`{"Response":{"Progress":[{"Type":"EnsureCreateTceCvm","Status":"False","Message":"Failed to create cvm. This cluster has failed to be created and can be cleaned up. Code=ResourceInsufficient.SpecifiedInstanceType"}],"RequestId":"r3"}}`), nil
		default:
			return nil, fmt.Errorf("unexpected action %s", requestAction(r))
		}
	}))

	clusterRes := resourceTencentCloudTkeCluster()
	d := schema.TestResourceDataRaw(t, clusterRes.Schema, map[string]interface{}{
		"vpc_id":                  "vpc-1",
		"cluster_cidr":            "172.24.0.0/16",
		"cluster_max_pod_num":     32,
		"cluster_max_service_num": 32,
		"cluster_name":            "fail-test",
		"cluster_deploy_type":     "MANAGED_CLUSTER",
		"cluster_version":         "1.24.4",
		"cluster_os":              "centos7.6.0_x64",
	})
	err := clusterRes.Create(d, &TencentCloudClient{apiV3Conn: conn})
	if err == nil {
		t.Fatalf("expected Create to fail fast when DescribeClusterCreateProgress reports fatal failure")
	}
	if !strings.Contains(err.Error(), "EnsureCreateTceCvm") || !strings.Contains(err.Error(), "ResourceInsufficient.SpecifiedInstanceType") {
		t.Fatalf("expected error to include step and cause, got: %v", err)
	}
}

func TestTKEScaleWorkerReadPreservesLabelsWhenAPIReturnsEmpty(t *testing.T) {
	conn := testClient()
	conn.UseTkeClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "DescribeClusters":
			return testResponse(`{"Response":{"TotalCount":1,"Clusters":[{"ClusterId":"cls-sw","ClusterName":"sw","ClusterVersion":"1.26.1","ClusterOs":"tlinux3.3x86_64","ClusterType":"INDEPENDENT_CLUSTER","ClusterNetworkSettings":{"ClusterCIDR":"172.26.0.0/16","IgnoreClusterCIDRConflict":false,"MaxNodePodNum":64,"MaxClusterServiceNum":256,"Ipvs":true,"VpcId":"vpc-1"},"ClusterNodeNum":1,"ProjectId":0,"ClusterStatus":"Running","Property":"{\"NetworkType\":\"GR\"}"}],"RequestId":"r1"}}`), nil
		case "DescribeClusterInstances":
			return testResponse(`{"Response":{"TotalCount":1,"InstanceSet":[{"InstanceId":"ins-w1","InstanceRole":"WORKER","InstanceState":"running","FailedReason":"","InstanceAdvancedSettings":{"Labels":null}}],"RequestId":"r2"}}`), nil
		default:
			return nil, fmt.Errorf("unexpected action %s", requestAction(r))
		}
	}))

	res := resourceTencentCloudTkeScaleWorker()
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"cluster_id": "cls-sw",
		"labels": map[string]interface{}{
			"env": "e2e",
		},
		"worker_instances_list": []interface{}{
			map[string]interface{}{
				"instance_id":    "ins-w1",
				"instance_role":  "WORKER",
				"instance_state": "initializing",
			},
		},
	})
	d.SetId("TkeScaleWorker.test")

	if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("scale worker Read failed: %v", err)
	}
	labels := d.Get("labels").(map[string]interface{})
	if labels["env"] != "e2e" {
		t.Fatalf("expected labels[env]=e2e to be preserved when API returns null Labels, got %#v", labels)
	}
	workers := d.Get("worker_instances_list").([]interface{})
	if len(workers) != 1 || workers[0].(map[string]interface{})["instance_state"] != "running" {
		t.Fatalf("expected worker_instances_list to refresh to running, got %#v", workers)
	}
}

func TestTKEClusterAttachmentReadAllowsNonRunningStateForDestroy(t *testing.T) {
	conn := testClient()
	conn.UseTkeClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		switch requestAction(r) {
		case "DescribeClusters":
			return testResponse(`{"Response":{"TotalCount":1,"Clusters":[{"ClusterId":"cls-att","ClusterName":"att","ClusterVersion":"1.26.1","ClusterOs":"tlinux3.3x86_64","ClusterType":"INDEPENDENT_CLUSTER","ClusterNetworkSettings":{"ClusterCIDR":"172.26.0.0/16","IgnoreClusterCIDRConflict":false,"MaxNodePodNum":64,"MaxClusterServiceNum":256,"Ipvs":true,"VpcId":"vpc-1"},"ClusterNodeNum":1,"ProjectId":0,"ClusterStatus":"Running","Property":"{\"NetworkType\":\"GR\"}"}],"RequestId":"r1"}}`), nil
		case "DescribeClusterInstances":
			return testResponse(`{"Response":{"TotalCount":1,"InstanceSet":[{"InstanceId":"ins-att1","InstanceRole":"WORKER","InstanceState":"failed","FailedReason":"init failed"}],"RequestId":"r2"}}`), nil
		default:
			return nil, fmt.Errorf("unexpected tke action %s", requestAction(r))
		}
	}))
	conn.UseCvmClient().WithHttpTransport(testTransport(func(r *http.Request) (*http.Response, error) {
		if requestAction(r) == "DescribeInstances" {
			return testResponse(`{"Response":{"TotalCount":1,"InstanceSet":[{"InstanceId":"ins-att1","InstanceState":"RUNNING","SecurityGroupIds":["sg-1"]}],"RequestId":"r3"}}`), nil
		}
		return nil, fmt.Errorf("unexpected cvm action %s", requestAction(r))
	}))

	res := resourceTencentCloudTkeClusterAttachment()
	d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
		"cluster_id":  "cls-att",
		"instance_id": "ins-att1",
	})
	d.SetId("ins-att1_cls-att")

	if err := res.Read(d, &TencentCloudClient{apiV3Conn: conn}); err != nil {
		t.Fatalf("attachment Read should not fail on non-running worker state so destroy can proceed: %v", err)
	}
	if got := d.Get("state").(string); got != "failed" {
		t.Fatalf("expected state=failed, got %q", got)
	}
}

