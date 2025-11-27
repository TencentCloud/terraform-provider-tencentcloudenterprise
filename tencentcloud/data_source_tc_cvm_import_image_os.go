/*
Use this data source to query detailed information of cvm import_image_os

# Example Usage

```hcl
data "tencentcloudenterprise_cvm_import_image_os" "import_image_os" {
}
```
*/
package tencentcloud

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	// 注册资源中文描述提供者映射
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_import_image_os", CNDescription{
		TerraformTypeCN: "CVM 导入镜像支持的操作系统",
		AttributesCN: map[string]string{
			"import_image_os_list_supported": "导入镜像支持的操作系统类型",
			"import_image_os_version_set":    "导入镜像支持的操作系统版本",
			"result_output_file":             "用于保存结果的文件",
			"windows":                        "支持导入的 Windows 操作系统注意：此字段可能返回空，表示取不到有效值",
			"linux":                          "支持导入的 Linux 操作系统注意：此字段可能返回空，表示取不到有效值",
			"os_name":                        "操作系统类型",
			"os_versions":                    "支持的操作系统版本",
			"architecture":                   "支持的操作系统架构",
		},
	})
}

func dataSourceTencentCloudCvmImportImageOs() *schema.Resource {
	return &schema.Resource{
		Description: "Query detailed information of cvm import_image_os",
		Read:        dataSourceTencentCloudCvmImportImageOsRead,
		Schema: map[string]*schema.Schema{
			"import_image_os_list_supported": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Supported operating system types of imported images.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"windows": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Supported Windows OS Note: This field may return null, indicating that no valid values can be obtained.",
						},
						"linux": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Supported Linux OS Note: This field may return null, indicating that no valid values can be obtained.",
						},
					},
				},
			},

			"import_image_os_version_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Supported operating system versions of imported images.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"os_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operating system type.",
						},
						"os_versions": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Supported operating system versions.",
						},
						"architecture": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Supported operating system architecture.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCvmImportImageOsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_import_image_os.read")()
	defer inconsistentCheck(d, meta)()

	var (
		request  = cvm.NewDescribeImportImageOsRequest()
		response = cvm.NewDescribeImportImageOsResponse()
	)
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		resule, e := meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DescribeImportImageOs(request)
		if e != nil {
			return retryError(e)
		}
		response = resule
		return nil
	})
	if err != nil {
		return err
	}

	if response == nil || response.Response == nil {
		d.SetId("")
		return fmt.Errorf("Response is null")
	}
	imageOsList := response.Response.ImportImageOsListSupported
	importImageOsVersionSet := response.Response.ImportImageOsInfoSupported
	result := make(map[string]interface{})
	if imageOsList != nil {
		imageOsListMap := make(map[string]interface{})

		if len(imageOsList.Windows) != 0 {
			windowsList := make([]string, 0, len(imageOsList.Windows))
			for _, v := range imageOsList.Windows {
				windowsList = append(windowsList, *v)
			}
			imageOsListMap["windows"] = windowsList
		}

		if len(imageOsList.Linux) != 0 {
			linuxList := make([]string, 0, len(imageOsList.Linux))
			for _, v := range imageOsList.Linux {
				linuxList = append(linuxList, *v)
			}
			imageOsListMap["linux"] = linuxList
		}

		result["import_image_os_list_supported"] = imageOsListMap
		_ = d.Set("import_image_os_list_supported", []map[string]interface{}{imageOsListMap})
	}
	if importImageOsVersionSet != nil {
		tmpList := make([]map[string]interface{}, 0)
		for _, osVersion := range importImageOsVersionSet {
			osVersionMap := map[string]interface{}{}

			if osVersion.OsName != nil {
				osVersionMap["os_name"] = osVersion.OsName
			}

			if osVersion.OsVersions != nil {
				osVersionMap["os_versions"] = osVersion.OsVersions
			}

			if osVersion.Architecture != nil {
				osVersionMap["architecture"] = osVersion.Architecture
			}

			tmpList = append(tmpList, osVersionMap)
		}
		result["import_image_os_version_set"] = tmpList
		_ = d.Set("import_image_os_version_set", tmpList)
	}

	d.SetId(helper.BuildToken())
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), result); e != nil {
			return e
		}
	}
	return nil
}
