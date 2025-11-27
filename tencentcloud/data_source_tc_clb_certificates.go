/*
Use this data source to query SSL certificate.

# Example Usage

```hcl

	data "tencentcloudenterprise_clb_certificates" "foo" {
	  result_output_file = "foo2.json"
	  with_cert = true
	  cert_ids = ["76cXqklW"]
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_clb_certificates", CNDescription{
		TerraformTypeCN: "负载均衡证书",
		DescriptionCN:   "提供负载均衡证书数据源，用于查询SSL证书的详细信息。",
		AttributesCN: map[string]string{
			"cert_type":          "要查询的SSL证书的类型可用值包括：“CA”和“SVR”",
			"cert_ids":           "要查询的SSL证书的ID",
			"with_cert":          "是否返回SSL证书的内容",
			"result_output_file": "用于保存结果",
			"certificates":       "证书的信息列表每个元素都包含以下属性：",
			"id":                 "SSL证书的ID",
			"name":               "SSL证书的名称",
			"type":               "SSL证书的类型",
			"project_id":         "SSL证书的项目ID",
			"cert":               "SSL证书的内容",
			"key":                "SSL证书的密钥",
			"product_zh_name":    "证书颁发机构",
			"domain":             "SSL证书的主域",
			"status":             "SSL证书的状态",
			"begin_time":         "SSL证书的开始时间",
			"end_time":           "SSL证书的结束时间",
			"create_time":        "SSL证书的创建时间",
			"subject_names":      "SSL证书中包含的所有域包括主域名",
		},
	})
}

func dataSourceTencentCloudClbCertificates() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the SSL certificate of the CLB instance.",
		Read:        dataSourceTencentCloudClbCertificatesRead,
		Schema: map[string]*schema.Schema{
			"cert_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the SSL certificate to be queried. Available values includes: `CA` and `SVR`.",
			},
			"cert_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "ID of the SSL certificate to be queried.",
			},
			"with_cert": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to return the content of the SSL certificate.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// computed
			"certificates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of certificate. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the SSL certificate.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the SSL certificate.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the SSL certificate.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Project ID of the SSL certificate.",
						},
						"cert": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Content of the SSL certificate.",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key of the SSL certificate.",
						},
						"product_zh_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate authority.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Primary domain of the SSL certificate.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Status of the SSL certificate.",
						},
						"begin_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Beginning time of the SSL certificate.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Ending time of the SSL certificate.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the SSL certificate.",
						},
						"subject_names": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "ALL domains included in the SSL certificate. Including the primary domain name.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudClbCertificatesRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("data_source.cloud_ssl_certificates.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		certType *string
		request  = clb.NewDescribeCertsRequest()
	)

	if raw, ok := d.GetOk("cert_type"); ok {
		certType = helper.String(raw.(string))
		request.CertType = certType
	}

	if v, ok := d.GetOk("cert_ids"); ok {
		instanceIDsSet := v.(*schema.Set).List()
		ids := make([]*string, 0, len(instanceIDsSet))
		for _, vv := range instanceIDsSet {
			ids = append(ids, helper.String(vv.(string)))
		}
		request.CertIdList = ids
	}

	if raw, ok := d.GetOk("with_cert"); ok {
		withCert := raw.(bool)
		request.WithCert = &withCert
	}

	clbService := ClbService{client: m.(*TencentCloudClient).apiV3Conn}
	certs, err := clbService.DescribeCert(ctx, request)
	if err != nil {
		return err
	}

	certificates := make([]map[string]interface{}, 0, len(certs))
	ids := make([]string, 0, len(certs))
	for _, certificate := range certs {
		if nilNames := CheckNil(certificate, map[string]string{
			"CertificateId":   "id",
			"Alias":           "name",
			"CertificateType": "type",
			"ProjectId":       "project id",
			"ProductZhName":   "product zh name",
			"Domain":          "domain",
			"Status":          "status",
			"CertBeginTime":   "begin time",
			"CertEndTime":     "end time",
			"InsertTime":      "create time",
		}); len(nilNames) > 0 {
			return fmt.Errorf("certificate %v are nil", nilNames)
		}

		ids = append(ids, *certificate.CertID)

		projectId, err := strconv.Atoi(*certificate.ProjectId)
		if err != nil {
			return err
		}

		m := map[string]interface{}{
			"id":              *certificate.CertID,
			"name":            *certificate.Alias,
			"type":            *certificate.CertType,
			"project_id":      projectId,
			"product_zh_name": *certificate.ProductZhName,
			"domain":          *certificate.Domain,
			"status":          *certificate.Status,
			"begin_time":      *certificate.CertBeginTime,
			"end_time":        *certificate.CertEndTime,
			"create_time":     *certificate.InsertTime,
		}

		if len(certificate.SubjectAltName) > 0 {
			subjectAltNames := make([]string, 0, len(certificate.SubjectAltName))
			for _, name := range certificate.SubjectAltName {
				subjectAltNames = append(subjectAltNames, *name)
			}
			m["subject_names"] = subjectAltNames
		}

		certificates = append(certificates, m)
	}

	_ = d.Set("certificates", certificates)
	d.SetId(helper.DataResourceIdsHash(ids))

	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), certificates); err != nil {
			log.Printf("[CRITAL]%s output file[%s] fail, reason[%s]",
				logId, output.(string), err.Error())
			return err
		}
	}

	return nil
}
