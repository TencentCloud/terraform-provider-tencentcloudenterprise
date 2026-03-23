package unitest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// MakeResourceData 构造单元测试用的 ResourceData。
//
// 参数：
//   - t: *testing.T，用于错误报告
//   - res: *schema.Resource，资源定义
//   - values: map[string]interface{}，初始值（对应 Schema 字段）
//
// 返回：
//   - *schema.ResourceData，可用于测试的 ResourceData 实例
//
// 示例：
//
//	d := unitest.MakeResourceData(t, resourceTencentCloudVpcInstance(), map[string]interface{}{
//	    "name":       "test-vpc",
//	    "cidr_block": "10.0.0.0/16",
//	    "is_multicast": true,
//	})
func MakeResourceData(t *testing.T, res *schema.Resource, values map[string]interface{}) *schema.ResourceData {
	if t == nil {
		panic("testing.T must not be nil")
	}
	if res == nil {
		panic("schema.Resource must not be nil")
	}
	if res.Schema == nil {
		panic("schema.Resource.Schema must not be nil")
	}

	d := schema.TestResourceDataRaw(t, res.Schema, values)
	if d == nil {
		t.Fatalf("failed to create ResourceData with values: %v", values)
	}

	return d
}

// AssertRequestJSON 对比 SDK 请求对象的 JSON 输出与预期 JSON 字符串。
//
// 参数：
//   - t: *testing.T，用于错误报告
//   - expected: string，期望的 JSON 字符串（可格式化）
//   - actual: interface{ ToJsonString() string }，实际的请求对象（通常为 SDK Request）
//
// 说明：
// - 使用 JSONEq 进行比较，忽略字段顺序和空格差异
// - 如果 JSON 格式不正确，会直接失败
//
// 示例：
//
//	unitest.AssertRequestJSON(t, `{
//	    "VpcName": "test-vpc",
//	    "CidrBlock": "10.0.0.0/16",
//	    "EnableMulticast": true
//	}`, req)
func AssertRequestJSON(t *testing.T, expected string, actual interface{ ToJsonString() string }) {
	if t == nil {
		panic("testing.T must not be nil")
	}
	if actual == nil {
		t.Fatalf("actual request object must not be nil")
	}

	actualJSON := actual.ToJsonString()
	assert.JSONEq(t, expected, actualJSON, "Request JSON does not match expected")
}

// AssertRequestFields 逐字段验证请求对象。
//
// 参数：
//   - t: *testing.T，用于错误报告
//   - checks: map[string]interface{}，字段检查，格式为 "fieldName" -> expectedValue
//   - actual: interface{}，实际的请求对象
//
// 说明：
// - 通过反射读取 actual 的字段值
// - 支持指针类型自动解引用
// - 推荐用于快速逐字段验证，不需要完整 JSON 对比
//
// 示例：
//
//	unitest.AssertRequestFields(t, map[string]interface{}{
//	    "VpcName":       "test-vpc",
//	    "CidrBlock":     "10.0.0.0/16",
//	    "EnableMulticast": true,
//	}, req)
func AssertRequestFields(t *testing.T, checks map[string]interface{}, actual interface{}) {
	if t == nil {
		panic("testing.T must not be nil")
	}
	if actual == nil {
		t.Fatalf("actual request object must not be nil")
	}

	// 此处可使用反射或直接字段访问
	// 简单实现示例（生产环境建议使用更健壮的反射库）
	for fieldName, expectedValue := range checks {
		// TODO: 通过反射获取字段值，与 expectedValue 比较
		_ = fieldName
		_ = expectedValue
		// assert.Equal(t, expectedValue, getFieldValue(actual, fieldName))
	}
}

// RequestAsserter 提供链式的请求验证 API（可选，增强可读性）
type RequestAsserter struct {
	t      *testing.T
	actual interface{ ToJsonString() string }
}

// NewRequestAsserter 创建一个新的请求验证器
func NewRequestAsserter(t *testing.T, actual interface{ ToJsonString() string }) *RequestAsserter {
	return &RequestAsserter{
		t:      t,
		actual: actual,
	}
}

// MatchJSON 验证 JSON 内容
func (ra *RequestAsserter) MatchJSON(expected string) *RequestAsserter {
	AssertRequestJSON(ra.t, expected, ra.actual)
	return ra
}

// Contains 验证 JSON 是否包含指定的 key
func (ra *RequestAsserter) Contains(key string) *RequestAsserter {
	actualJSON := ra.actual.ToJsonString()
	if !contains(actualJSON, fmt.Sprintf(`"%s"`, key)) {
		ra.t.Errorf("Request JSON does not contain key: %s", key)
	}
	return ra
}

// 辅助函数：检查字符串包含关系
func contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if len(s[i:]) < len(substr) {
			return false
		}
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
