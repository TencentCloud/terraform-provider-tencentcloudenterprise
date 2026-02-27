# Terraform Enterprise Provider 文档生成工具

## 简介

这个工具用于自动生成 TencentCloud Enterprise Provider 的 Terraform 文档。它能确保文档格式统一、内容准确，并且与代码保持同步。

## 为什么需要这个工具？

手动编写 Terraform 文档存在以下问题：

* **格式不统一**：章节顺序、空行数量、命名风格在不同产品之间可能存在差异
* **细节错误**：空格数量、缩进、符号使用可能不一致
* **内容不准确**：参数的必填/选填状态可能与代码不符，参数列表可能有遗漏
* **维护成本高**：文档更新不及时，需要大量时间和精力

自动生成工具可以：
* 保证文档格式完全一致
* 确保文档与代码同步
* 自动提取所有参数和属性
* 节省大量时间和精力

## 如何使用

### 1. 确保代码符合文档规范

#### 资源/数据源文件注释格式

在 `resource_tc_*.go` 或 `data_source_tc_*.go` 文件开头添加注释：

\`\`\`go
/*
Provides a CVM instance resource.

~> **NOTE:** You can launch an CVM instance for a VPC network via specifying parameter `vpc_id`. One instance can only belong to one VPC.

Example Usage

\`\`\`hcl
resource "tencentcloudenterprise_cvm_instance" "example" {
  instance_name     = "example-instance"
  availability_zone = "ap-guangzhou-3"
  image_id          = "img-xxx"
  instance_type     = "S5.MEDIUM2"
  
  system_disk {
    disk_type = "CLOUD_PREMIUM"
    disk_size = 50
  }
}
\`\`\`

Import

tencentcloudenterprise_cvm_instance can be imported using the id, e.g.

\`\`\`
$ terraform import tencentcloudenterprise_cvm_instance.example ins-xxxxxxxx
\`\`\`
*/
package tencentcloud
\`\`\`

#### Schema 定义规范

每个 Schema 字段必须包含 Description：

\`\`\`go
"instance_name": {
    Type:        schema.TypeString,
    Required:    true,
    Description: "The name of the instance.",
},

"vpc_id": {
    Type:        schema.TypeString,
    Optional:    true,
    Computed:    true,
    ForceNew:    true,
    Description: "The ID of a VPC network. If you want to create instances in a VPC network, this parameter must be set.",
},

"private_ip": {
    Type:        schema.TypeString,
    Computed:    true,
    Description: "The private IP of the instance.",
},
\`\`\`

#### Provider 文件注释

在 `provider.go` 中维护资源列表：

\`\`\`go
/*
...

Resources List

Provider Data Sources
  tencentcloudenterprise_availability_zones
  tencentcloudenterprise_availability_regions

CVM
  Data Source
    tencentcloudenterprise_cvm_instances
    tencentcloudenterprise_cvm_images
    
  Resource
    tencentcloudenterprise_cvm_instance
    tencentcloudenterprise_cvm_key_pair

VPC
  Data Source
    tencentcloudenterprise_vpc_instances
    tencentcloudenterprise_vpc_subnets
    
  Resource
    tencentcloudenterprise_vpc
    tencentcloudenterprise_vpc_subnet
*/
package tencentcloud
\`\`\`

### 2. 运行文档生成工具

\`\`\`bash
cd terraform-provider-tencentcloudenterprise/gendoc
go run *.go
\`\`\`

### 3. 检查生成的文档

文档将生成在 `../docs/` 目录下：

* `docs/index.md` - Provider 首页文档
* `docs/resources/*.md` - 各个资源的文档
* `docs/data-sources/*.md` - 各个数据源的文档

## 文档结构

生成的文档包含以下部分：

1. **名称和描述**：从代码注释中提取
2. **Example Usage**：从代码注释中提取示例代码
3. **Argument Reference**：从 Schema 定义中自动生成参数列表
4. **Attributes Reference**：从 Schema 定义中自动生成导出属性列表
5. **Import**：从代码注释中提取导入说明

## 注意事项

* 所有 Description 必须以英文句号 `.` 或冒号 `:` 结尾
* Description 开头和结尾不能有空格
* Description 中不能包含中文标点符号
* 数据源必须包含 `result_output_file` 参数
* 不要在数据源的 Schema 中设置 `ForceNew`
* 不要同时设置 `Required` 和 `Optional`

## 开发指南

如果需要修改文档模板，请编辑 `template.go` 文件中的 `docTPL` 和 `idxTPL` 常量。

模板使用 Go 的 `text/template` 语法，支持以下变量：

* `{{.name}}` - 资源/数据源名称
* `{{.description}}` - 详细描述
* `{{.example}}` - 示例代码
* `{{.arguments}}` - 参数列表
* `{{.attributes}}` - 属性列表
* `{{.import}}` - 导入说明
