/*
Provides an available image for the user.

The Images data source fetch proper image, which could be one of the private images of the user and images of system resources provided by TencentCloud, as well as other public images and those available on the image market.

~> **NOTE:** This data source will be deprecated, please use `cloud_cvm_images` instead.

# Example Usage

```hcl

	data "tencentcloudenterprise_cvm_image" "my_favorate_image" {
	  os_name = "centos"

	  filter {
	    name   = "image-type"
	    values = ["PUBLIC_IMAGE"]
	  }
	}

```
*/
package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	// 注册资源中文描述提供者映射
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_image", CNDescription{
		TerraformTypeCN: "CVM 镜像",
		DescriptionCN:   "提供CVM镜像数据源，用于为用户查询可用的镜像信息。",
		AttributesCN: map[string]string{
			"filter":             "一个或多个名称/值对过滤器",
			"image_name_regex":   "应用于返回的镜像列表的正则表达式注意：这不是通配符",
			"os_name":            "将应用于返回的镜像列表中的 os_name 属性的模糊匹配字符串",
			"result_output_file": "用于保存结果的文件",
			"image_id":           "表示某种镜像唯一性的镜像 ID，可用于实例创建或重置",
			"image_name":         "此镜像的名称",
			"name":               "过滤器的键，有效键：`image-id`, `image-type`, `image-name`",
			"values":             "过滤器的值",
		},
	})
}

func dataSourceTencentCloudImage() *schema.Resource {
	return &schema.Resource{
		Description: "Provides an available image for the user.",
		Read:        dataSourceTencentCloudImageRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "One or more name/value pairs to filter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Key of the filter, valid keys: `image-id`, `image-type`, `image-name`.",
						},
						"values": {
							Type:        schema.TypeList,
							Required:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Values of the filter.",
						},
					},
				},
			},
			"image_name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNameRegex,
				Description:  "A regex string to apply to the image list returned by TencentCloud. **NOTE**: it is not wildcard, should look like `image_name_regex = \"^CentOS\\s+6\\.8\\s+64\\w*\"`.",
			},
			"os_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNotEmpty,
				Description:  "A string to apply with fuzzy match to the os_name attribute on the image list returned by TencentCloud. **NOTE**: when os_name is provided, highest priority is applied in this field instead of `image_name_regex`.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An image id indicate the uniqueness of a certain image,  which can be used for instance creation or resetting.",
			},
			"image_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of this image.",
			},
		},
	}
}

func dataSourceTencentCloudImageRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_image.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	filter := make(map[string][]string)
	filters, ok := d.GetOk("filter")
	if ok {
		for _, v := range filters.(*schema.Set).List() {
			vv := v.(map[string]interface{})
			name := vv["name"].(string)
			filter[name] = []string{}
			for _, vvv := range vv["values"].([]interface{}) {
				filter[name] = append(filter[name], vvv.(string))
			}
		}
	}

	var images []*cvm.Image
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		images, errRet = cvmService.DescribeImagesByFilter(ctx, filter, "")
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if len(images) == 0 {
		return errors.New("No image found")
	}

	var osName string
	if v, ok := d.GetOk("os_name"); ok {
		osName = v.(string)
	}

	var regImageName string
	var imageNameRegex *regexp.Regexp
	if v, ok := d.GetOk("image_name_regex"); ok {
		regImageName = v.(string)
		imageNameRegex, err = regexp.Compile(regImageName)
		if err != nil {
			return fmt.Errorf("image_name_regex format error,%s", err.Error())
		}
	}

	var resultImageId string
	images = sortImages(images)
	for _, image := range images {
		if osName != "" {
			if strings.Contains(strings.ToLower(*image.OsName), strings.ToLower(osName)) {
				resultImageId = *image.ImageId
				_ = d.Set("image_name", *image.ImageName)
				break
			}
			continue
		}

		if imageNameRegex != nil {
			if imageNameRegex.MatchString(*image.ImageName) {
				resultImageId = *image.ImageId
				_ = d.Set("image_name", *image.ImageName)
				break
			}
			continue
		}

		resultImageId = *image.ImageId
		_ = d.Set("image_name", *image.ImageName)
		break
	}

	if resultImageId == "" {
		return errors.New("No image found")
	}

	d.SetId(helper.DataResourceIdHash(resultImageId))
	if err := d.Set("image_id", resultImageId); err != nil {
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), resultImageId); err != nil {
			return err
		}
	}

	return nil
}
