/*
Use this data source to query the COS buckets of the current Cloud user.

# Example Usage

```hcl

	data "tencentcloudenterprise_csp_buckets" "csp_buckets" {
	  result_output_file = "mytestpath"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_csp_buckets", CNDescription{
		TerraformTypeCN: "获取存储桶列表（CSP）",
		DescriptionCN:   "提供CSP存储桶数据源，用于查询当前云用户的CSP存储桶。",
		AttributesCN: map[string]string{
			"owner":       "桶的所有者信息，包含 id 和 display_name 字段，分别表示所有者的唯一标识和显示名称",
			"bucket_list": "CSP桶列表",
			"name":        "桶名称",
			//"location":    "桶所在地域",
			"create_date":        "桶创建时间",
			"result_output_file": "用于保存数据源查询结果, 在前端可视化界面使用时，该参数不可用",
		},
	})
}

func dataSourceTencentCloudCspBuckets() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query the CSP buckets of the current Cloud user.",
		Read:        dataSourceTencentCloudCspBucketsRead,

		Schema: map[string]*schema.Schema{
			//"bucket_prefix": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Description: "A prefix string to filter results by bucket name.",
			//},
			//"tags": {
			//	Type:        schema.TypeMap,
			//	Optional:    true,
			//	Description: "Tags to filter bucket.",
			//},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// computed
			"owner": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Information about the owner of the bucket, including the id and display_name fields, which represent the unique identifier and display name of the owner respectively.",
			},
			"bucket_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Csp bucket list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "bucket name.",
						},
						//"location": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "bucket region.",
						//},
						"create_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "bucket create date.",
						},
					},
				},
			},
			//"bucket_list": {
			//	Type:        schema.TypeList,
			//	Computed:    true,
			//	Description: "A list of bucket. Each element contains the following attributes:",
			//	Elem: &schema.Resource{
			//		Schema: map[string]*schema.Schema{
			//			"bucket": {
			//				Type:        schema.TypeString,
			//				Computed:    true,
			//				Description: "Bucket name, the format likes `<bucket>-<appid>`.",
			//			},
			//			"cors_rules": {
			//				Type:        schema.TypeList,
			//				Computed:    true,
			//				Description: "A list of CORS rule configurations.",
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"allowed_origins": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Elem:        &schema.Schema{Type: schema.TypeString},
			//							Description: "Specifies which origins are allowed.",
			//						},
			//						"allowed_methods": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Elem:        &schema.Schema{Type: schema.TypeString},
			//							Description: "Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.",
			//						},
			//						"allowed_headers": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Elem:        &schema.Schema{Type: schema.TypeString},
			//							Description: "Specifies which headers are allowed.",
			//						},
			//						"max_age_seconds": {
			//							Type:        schema.TypeInt,
			//							Computed:    true,
			//							Description: "Specifies time in seconds that browser can cache the response for a preflight request.",
			//						},
			//						"expose_headers": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Elem:        &schema.Schema{Type: schema.TypeString},
			//							Description: "Specifies expose header in the response.",
			//						},
			//					},
			//				},
			//			},
			//			"lifecycle_rules": {
			//				Type:        schema.TypeList,
			//				Computed:    true,
			//				Description: "The lifecycle configuration of a bucket.",
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"filter_prefix": {
			//							Type:        schema.TypeString,
			//							Computed:    true,
			//							Description: "Object key prefix identifying one or more objects to which the rule applies.",
			//						},
			//						"transition": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Description: "Specifies a period in the object's transitions.",
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									"date": {
			//										Type:        schema.TypeString,
			//										Computed:    true,
			//										Description: "Specifies the date after which you want the corresponding action to take effect.",
			//									},
			//									"days": {
			//										Type:        schema.TypeInt,
			//										Computed:    true,
			//										Description: "Specifies the number of days after object creation when the specific rule action takes effect.",
			//									},
			//									"storage_class": {
			//										Type:        schema.TypeString,
			//										Computed:    true,
			//										Description: "Specifies the storage class to which you want the object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.",
			//									},
			//								},
			//							},
			//						},
			//						"expiration": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Description: "Specifies a period in the object's expire.",
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									"date": {
			//										Type:        schema.TypeString,
			//										Computed:    true,
			//										Description: "Specifies the date after which you want the corresponding action to take effect.",
			//									},
			//									"days": {
			//										Type:        schema.TypeInt,
			//										Computed:    true,
			//										Description: "Specifies the number of days after object creation when the specific rule action takes effect.",
			//									},
			//								},
			//							},
			//						},
			//						"non_current_transition": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Description: "Specifies when to transition objects of non current versions and the target storage class.",
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									"non_current_days": {
			//										Type:        schema.TypeInt,
			//										Computed:    true,
			//										Description: "Number of days after non current object creation when the specific rule action takes effect.",
			//									},
			//									"storage_class": {
			//										Type:        schema.TypeString,
			//										Computed:    true,
			//										Description: "Specifies the storage class to which you want the non current object to transition. Available values include STANDARD, STANDARD_IA and ARCHIVE.",
			//									},
			//								},
			//							},
			//						},
			//						"non_current_expiration": {
			//							Type:        schema.TypeList,
			//							Computed:    true,
			//							Description: "Specifies when non current object versions shall expire.",
			//							Elem: &schema.Resource{
			//								Schema: map[string]*schema.Schema{
			//									"non_current_days": {
			//										Type:        schema.TypeInt,
			//										Computed:    true,
			//										Description: "Number of days after non current object creation when the specific rule action takes effect. The maximum value is 3650.",
			//									},
			//								},
			//							},
			//						},
			//					},
			//				},
			//			},
			//			"website": {
			//				Type:        schema.TypeList,
			//				Computed:    true,
			//				Description: "A list of one element containing configuration parameters used when the bucket is used as a website.",
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"index_document": {
			//							Type:        schema.TypeString,
			//							Computed:    true,
			//							Description: "COS returns this index document when requests are made to the root domain or any of the subfolders.",
			//						},
			//						"error_document": {
			//							Type:        schema.TypeString,
			//							Computed:    true,
			//							Description: "An absolute path to the document to return in case of a 4XX error.",
			//						},
			//					},
			//				},
			//			},
			//			"origin_pull_rules": {
			//				Type:        schema.TypeList,
			//				Computed:    true,
			//				Description: "Bucket Origin-Pull rules.",
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"priority": {
			//							Type:        schema.TypeInt,
			//							Required:    true,
			//							Description: "Priority of origin-pull rules, do not set the same value for multiple rules.",
			//						},
			//						"sync_back_to_source": {
			//							Type:        schema.TypeBool,
			//							Optional:    true,
			//							Default:     false,
			//							Description: "If `true`, COS will not return 3XX status code when pulling data from an origin server. Currently available zone: ap-beijing, ap-shanghai, ap-singapore, ap-mumbai.",
			//						},
			//						"prefix": {
			//							Type:        schema.TypeString,
			//							Optional:    true,
			//							Default:     "",
			//							Description: "Triggers the origin-pull rule when the requested file name matches this prefix.",
			//						},
			//						"protocol": {
			//							Type:        schema.TypeString,
			//							Optional:    true,
			//							Default:     "",
			//							Description: "The protocol used for COS to access the specified origin server. The available value include `HTTP`, `HTTPS` and `FOLLOW`.",
			//						},
			//						"host": {
			//							Type:        schema.TypeString,
			//							Required:    true,
			//							Description: "Allows only a domain name or IP address. You can optionally append a port number to the address.",
			//						},
			//						"follow_query_string": {
			//							Type:        schema.TypeBool,
			//							Optional:    true,
			//							Default:     true,
			//							Description: "Specifies whether to pass through COS request query string when accessing the origin server.",
			//						},
			//						"follow_redirection": {
			//							Type:        schema.TypeBool,
			//							Optional:    true,
			//							Default:     true,
			//							Description: "Specifies whether to follow 3XX redirect to another origin server to pull data from.",
			//						},
			//						//"copy_origin_data": {
			//						//	Type:		 schema.TypeBool,
			//						//	Optional: 	 true,
			//						//	Default:	 true,
			//						//	Description: "",
			//						//},
			//						"follow_http_headers": {
			//							Type:        schema.TypeList,
			//							Optional:    true,
			//							Description: "Specifies the pass through headers when accessing the origin server.",
			//							Elem:        &schema.Schema{Type: schema.TypeString},
			//						},
			//						"custom_http_headers": {
			//							Type:        schema.TypeMap,
			//							Optional:    true,
			//							Description: "Specifies the custom headers that you can add for COS to access your origin server.",
			//						},
			//						//"redirect_prefix": {
			//						//	Type:		schema.TypeString,
			//						//	Optional:   true,
			//						//	Description: "Prefix for the file to which a request is redirected when the origin-pull rule is triggered.",
			//						//},
			//						//"redirect_suffix": {
			//						//	Type:		schema.TypeString,
			//						//	Optional:   true,
			//						//	Description: "Suffix for the file to which a request is redirected when the origin-pull rule is triggered.",
			//						//},
			//					},
			//				},
			//			},
			//			"origin_domain_rules": {
			//				Type:        schema.TypeList,
			//				Computed:    true,
			//				Description: "Bucket origin domain rules.",
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"domain": {
			//							Type:        schema.TypeString,
			//							Required:    true,
			//							Description: "Specify domain host.",
			//						},
			//						"type": {
			//							Type:        schema.TypeString,
			//							Optional:    true,
			//							Default:     "REST",
			//							Description: "Specify origin domain type, available values: `REST`, `WEBSITE`, `ACCELERATE`, default: `REST`.",
			//						},
			//						"status": {
			//							Type:         schema.TypeString,
			//							Optional:     true,
			//							Default:      "ENABLED",
			//							Description:  "Domain status, default: `ENABLED`.",
			//							ValidateFunc: validateAllowedStringValue([]string{"ENABLED", "DISABLED"}),
			//						},
			//					},
			//				},
			//			},
			//			"acl": {
			//				Type:        schema.TypeString,
			//				Computed:    true,
			//				Description: "Bucket access control configurations.",
			//			},
			//			"acl_body": {
			//				Type:        schema.TypeString,
			//				Computed:    true,
			//				Description: "Bucket verbose acl configurations.",
			//			},
			//			"tags": {
			//				Type:        schema.TypeMap,
			//				Computed:    true,
			//				Description: "The tags of a bucket.",
			//			},
			//			"cos_bucket_url": {
			//				Type:        schema.TypeString,
			//				Computed:    true,
			//				Description: "The URL of this cos bucket.",
			//			},
			//		},
			//	},
			//},
		},
	}
}

func dataSourceTencentCloudCspBucketsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_csp_buckets.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	cspService := CosService{client: meta.(*TencentCloudClient).apiV3Conn, useCspClient: true}
	bucketResult, err := cspService.ListTencentCloudBuckets(ctx)
	if err != nil {
		return err
	}
	log.Printf("get bucket list result: %+v", bucketResult)
	//prefix := d.Get("bucket_prefix").(string)
	//tags := helper.GetTags(d, "tags")

	bucketList := make([]map[string]interface{}, 0, len(bucketResult.Buckets))

	for _, v := range bucketResult.Buckets {
		bucket := make(map[string]interface{})
		//if prefix != "" && !strings.HasPrefix(*v.Name, prefix) {
		//	continue
		//}
		//
		//respTags, err := cosService.GetBucketTags(ctx, *v.Name)
		//if err != nil {
		//	return err
		//}
		//
		//var matchTags bool
		//
		//for k, v := range tags {
		//	if respTags[k] == v {
		//		matchTags = true
		//		break
		//	}
		//}
		//
		//if len(tags) != 0 && !matchTags {
		//	continue
		//}

		bucket["name"] = v.Name
		//bucket["location"] = v.Region
		bucket["create_date"] = v.CreationDate
		//corsRules, err := cosService.GetBucketCors(ctx, *v.Name)
		//if err != nil {
		//	return err
		//}
		//bucket["cors_rules"] = corsRules
		//
		//lifecycleRules, err := cosService.GetDataSourceBucketLifecycle(ctx, *v.Name)
		//if err != nil {
		//	return err
		//}
		//bucket["lifecycle_rules"] = lifecycleRules
		//
		//website, err := cosService.GetBucketWebsite(ctx, *v.Name)
		//if err != nil {
		//	return err
		//}
		//bucket["website"] = website
		//
		//originRules, err := cosService.GetBucketPullOrigin(ctx, *v.Name)
		//if err != nil {
		//	return err
		//}
		//bucket["origin_pull_rules"] = originRules
		//
		//domainRules, err := cosService.GetBucketOriginDomain(ctx, *v.Name)
		//if err == nil {
		//	bucket["origin_domain_rules"] = domainRules
		//}

		//aclBody, err := cosService.GetBucketACL(ctx, *v.Name)
		//
		//if err != nil {
		//	return err
		//}
		//
		//aclXML, err := xml.Marshal(aclBody)
		//
		//if err != nil {
		//	log.Printf("WARN: acl body marshal failed: %s", err.Error())
		//} else {
		//	bucket["acl"] = GetBucketPublicACL(aclBody)
		//	bucket["acl_body"] = string(aclXML)
		//}
		//
		//bucket["tags"] = respTags
		//bucket["cos_bucket_url"] = fmt.Sprintf("%s.cos.%s.myqcloud.com", *v.Name, meta.(*TencentCloudClient).apiV3Conn.Region)
		bucketList = append(bucketList, bucket)
	}

	ids := make([]string, 2)
	ids[0] = "bucket_list"
	ids[1] = bucketResult.Owner.ID
	d.SetId(helper.DataResourceIdsHash(ids))
	if err := d.Set("bucket_list", bucketList); err != nil {
		return fmt.Errorf("setting bucket list error: %s", err.Error())
	}
	owner := make(map[string]interface{})
	owner["id"] = bucketResult.Owner.ID
	owner["display_name"] = bucketResult.Owner.DisplayName
	if err := d.Set("owner", owner); err != nil {
		return fmt.Errorf("setting owner error: %s", err.Error())
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		result := make(map[string]interface{})
		result["bucket_list"] = bucketList
		result["owner"] = owner
		if err = writeToFile(output.(string), result); err != nil {
			return err
		}
	}

	return nil
}
