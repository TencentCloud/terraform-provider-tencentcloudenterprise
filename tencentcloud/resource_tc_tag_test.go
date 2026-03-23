package tencentcloud

import (
	"encoding/json"
	"testing"

	tag "terraform-provider-tencentcloudenterprise/sdk/tag/v20180813"
)

// ==================== tencentcloudenterprise_tag 资源测试 ====================

func TestCreateTagRequestConstruction(t *testing.T) {
	// 测试 CreateTagRequest 构造
	request := tag.NewCreateTagRequest()

	tagKey := "env"
	tagValue := "production"

	request.TagKey = &tagKey
	request.TagValue = &tagValue

	// 验证字段是否正确设置
	if request.TagKey == nil || *request.TagKey != tagKey {
		t.Errorf("TagKey not set correctly, expected %s, got %v", tagKey, request.TagKey)
	}

	if request.TagValue == nil || *request.TagValue != tagValue {
		t.Errorf("TagValue not set correctly, expected %s, got %v", tagValue, request.TagValue)
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("CreateTagRequest JSON: %s", jsonStr)

	// 验证 JSON 包含正确的字段
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonMap["TagKey"] != tagKey {
		t.Errorf("JSON TagKey mismatch, expected %s, got %v", tagKey, jsonMap["TagKey"])
	}

	if jsonMap["TagValue"] != tagValue {
		t.Errorf("JSON TagValue mismatch, expected %s, got %v", tagValue, jsonMap["TagValue"])
	}
}

func TestDeleteTagRequestConstruction(t *testing.T) {
	// 测试 DeleteTagRequest 构造
	request := tag.NewDeleteTagRequest()

	tagKey := "env"
	tagValue := "production"

	request.TagKey = &tagKey
	request.TagValue = &tagValue

	// 验证字段
	if request.TagKey == nil || *request.TagKey != tagKey {
		t.Errorf("TagKey not set correctly")
	}

	if request.TagValue == nil || *request.TagValue != tagValue {
		t.Errorf("TagValue not set correctly")
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("DeleteTagRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonMap["TagKey"] != tagKey {
		t.Errorf("JSON TagKey mismatch")
	}

	if jsonMap["TagValue"] != tagValue {
		t.Errorf("JSON TagValue mismatch")
	}
}

func TestDescribeTagsRequestConstruction(t *testing.T) {
	// 测试 DescribeTagsRequest 构造
	request := tag.NewDescribeTagsRequest()

	tagKey := "env"
	tagValue := "production"
	offset := uint64(0)
	limit := uint64(20)

	request.TagKey = &tagKey
	request.TagValue = &tagValue
	request.Offset = &offset
	request.Limit = &limit

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("DescribeTagsRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonMap["TagKey"] != tagKey {
		t.Errorf("JSON TagKey mismatch")
	}

	if jsonMap["TagValue"] != tagValue {
		t.Errorf("JSON TagValue mismatch")
	}

	if uint64(jsonMap["Offset"].(float64)) != offset {
		t.Errorf("JSON Offset mismatch")
	}

	if uint64(jsonMap["Limit"].(float64)) != limit {
		t.Errorf("JSON Limit mismatch")
	}
}

// ==================== tencentcloudenterprise_tag_attachment 资源测试 ====================

func TestAddResourceTagRequestConstruction(t *testing.T) {
	// 测试 AddResourceTagRequest 构造
	request := tag.NewAddResourceTagRequest()

	tagKey := "env"
	tagValue := "production"
	resource := "qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-12345678"

	request.TagKey = &tagKey
	request.TagValue = &tagValue
	request.Resource = &resource

	// 验证字段是否正确设置
	if request.TagKey == nil || *request.TagKey != tagKey {
		t.Errorf("TagKey not set correctly")
	}

	if request.TagValue == nil || *request.TagValue != tagValue {
		t.Errorf("TagValue not set correctly")
	}

	if request.Resource == nil || *request.Resource != resource {
		t.Errorf("Resource not set correctly")
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("AddResourceTagRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonMap["TagKey"] != tagKey {
		t.Errorf("JSON TagKey mismatch")
	}

	if jsonMap["TagValue"] != tagValue {
		t.Errorf("JSON TagValue mismatch")
	}

	if jsonMap["Resource"] != resource {
		t.Errorf("JSON Resource mismatch")
	}
}

func TestDeleteResourceTagRequestConstruction(t *testing.T) {
	// 测试 DeleteResourceTagRequest 构造
	request := tag.NewDeleteResourceTagRequest()

	tagKey := "env"
	resource := "qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-12345678"

	request.TagKey = &tagKey
	request.Resource = &resource

	// 验证字段
	if request.TagKey == nil || *request.TagKey != tagKey {
		t.Errorf("TagKey not set correctly")
	}

	if request.Resource == nil || *request.Resource != resource {
		t.Errorf("Resource not set correctly")
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("DeleteResourceTagRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if jsonMap["TagKey"] != tagKey {
		t.Errorf("JSON TagKey mismatch")
	}

	if jsonMap["Resource"] != resource {
		t.Errorf("JSON Resource mismatch")
	}
}

func TestGetResourcesRequestConstruction(t *testing.T) {
	// 测试 GetResourcesRequest 构造 (用于 tag_attachment 的 Read)
	request := tag.NewGetResourcesRequest()

	resourceList := []*string{
		strPtr("qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-12345678"),
		strPtr("qcs::cvm:ap-guangzhou:uin/123456789:instance/ins-87654321"),
	}
	request.ResourceList = resourceList

	// 验证字段
	if len(request.ResourceList) != 2 {
		t.Errorf("ResourceList length mismatch, expected 2, got %d", len(request.ResourceList))
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("GetResourcesRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	resourceArr, ok := jsonMap["ResourceList"].([]interface{})
	if !ok || len(resourceArr) != 2 {
		t.Errorf("JSON ResourceList mismatch")
	}
}

// ==================== tencentcloudenterprise_tag_keys 数据源测试 ====================

func TestDescribeTagKeysRequestConstruction(t *testing.T) {
	// 测试 DescribeTagKeysRequest 构造
	request := tag.NewDescribeTagKeysRequest()

	limit := int64(100)
	offset := uint64(0)

	request.Limit = &limit
	request.Offset = &offset

	// 验证字段
	if request.Limit == nil || *request.Limit != limit {
		t.Errorf("Limit not set correctly")
	}

	if request.Offset == nil || *request.Offset != offset {
		t.Errorf("Offset not set correctly")
	}

	// 验证 JSON 序列化
	jsonStr := request.ToJsonString()
	t.Logf("DescribeTagKeysRequest JSON: %s", jsonStr)

	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if int64(jsonMap["Limit"].(float64)) != limit {
		t.Errorf("JSON Limit mismatch")
	}

	if uint64(jsonMap["Offset"].(float64)) != offset {
		t.Errorf("JSON Offset mismatch")
	}
}

// ==================== 综合测试 ====================

func TestTagResourceWorkflow(t *testing.T) {
	// 模拟完整的 tag 资源工作流程

	// 1. 创建 tag
	createReq := tag.NewCreateTagRequest()
	tagKey := "project"
	tagValue := "terraform-test"
	createReq.TagKey = &tagKey
	createReq.TagValue = &tagValue

	createJSON := createReq.ToJsonString()
	t.Logf("Step 1 - Create Tag: %s", createJSON)

	// 验证
	var createMap map[string]interface{}
	json.Unmarshal([]byte(createJSON), &createMap)
	if createMap["TagKey"] != tagKey || createMap["TagValue"] != tagValue {
		t.Errorf("Create tag request construction failed")
	}

	// 2. 查询 tag
	describeReq := tag.NewDescribeTagsRequest()
	describeReq.TagKey = &tagKey
	describeReq.TagValue = &tagValue

	describeJSON := describeReq.ToJsonString()
	t.Logf("Step 2 - Describe Tags: %s", describeJSON)

	// 3. 删除 tag
	deleteReq := tag.NewDeleteTagRequest()
	deleteReq.TagKey = &tagKey
	deleteReq.TagValue = &tagValue

	deleteJSON := deleteReq.ToJsonString()
	t.Logf("Step 3 - Delete Tag: %s", deleteJSON)

	t.Log("Tag resource workflow test completed successfully")
}

func TestTagAttachmentWorkflow(t *testing.T) {
	// 模拟完整的 tag_attachment 资源工作流程

	tagKey := "environment"
	tagValue := "staging"
	resource := "qcs::cvm:ap-shanghai:uin/100000000001:instance/ins-abcd1234"

	// 1. 添加资源标签
	addReq := tag.NewAddResourceTagRequest()
	addReq.TagKey = &tagKey
	addReq.TagValue = &tagValue
	addReq.Resource = &resource

	addJSON := addReq.ToJsonString()
	t.Logf("Step 1 - Add Resource Tag: %s", addJSON)

	// 2. 查询资源标签 (使用 GetResources)
	getReq := tag.NewGetResourcesRequest()
	getReq.ResourceList = []*string{&resource}

	getJSON := getReq.ToJsonString()
	t.Logf("Step 2 - Get Resources: %s", getJSON)

	// 3. 删除资源标签
	deleteReq := tag.NewDeleteResourceTagRequest()
	deleteReq.TagKey = &tagKey
	deleteReq.Resource = &resource

	deleteJSON := deleteReq.ToJsonString()
	t.Logf("Step 3 - Delete Resource Tag: %s", deleteJSON)

	t.Log("Tag attachment workflow test completed successfully")
}

// 辅助函数
func strPtr(s string) *string {
	return &s
}
