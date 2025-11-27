/*
Use this data source to query detailed information of dcdb accounts.

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_accounts" "foo" {
	  instance_id = cloud_dcdb_account.foo.instance_id
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_accounts", CNDescription{
		TerraformTypeCN: "云数据库账号信息",
		DescriptionCN:   "提供DCDB账户数据源，用于查询DCDB账户的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "实例id",
			"list":               "云数据库帐户信息",
			"user_name":          "用户名",
			"host":               "用户可以从哪个主机登录（对应MySQL用户的主机字段，UserName+host唯一标识用户，以IP形式，IP段以%结尾；支持填写%；如果为空，默认为%）",
			"description":        "用户备注信息",
			"create_time":        "创建时间",
			"update_time":        "上次更新时间",
			"read_only":          "只读标志，0:否，1:该账号的SQL请求优先在备机上执行，在备机不可用时选择主机执行2：优先选择备用机执行，当备用机不可用时操作失败",
			"delay_thresh":       "如果备用机延迟超过该参数的设置值，系统会认为备用机有故障，建议参数值大于10当ReadOnly选择1和2时，此参数生效",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbAccounts() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud database account information.",
		Read:        dataSourceTencentCloudDcdbAccountsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cloud database account information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User Name.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "From which host the user can log in (corresponding to the host field of MySQL users, UserName + Host uniquely identifies a user, in the form of IP, the IP segment ends with %; supports filling in %; if it is empty, it defaults to %).",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User remarks info.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last update time.",
						},
						"read_only": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Read-only flag, 0: No, 1: The SQL request of this account is preferentially executed on the standby machine, and the host is selected for execution when the standby machine is unavailable. 2: The standby machine is preferentially selected for execution, and the operation fails when the standby machine is unavailable.",
						},
						"delay_thresh": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "If the standby machine delay exceeds the setting value of this parameter, the system will consider that the standby machine is faulty and recommend that the parameter value be greater than 10. This parameter takes effect when ReadOnly selects 1 and 2.",
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

func dataSourceTencentCloudDcdbAccountsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_accounts.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	dcdbService := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var dbAccountList []*dcdb.DBAccount
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := dcdbService.DescribeDcdbAccountsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		dbAccountList = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb list failed, reason:%+v", logId, err)
		return err
	}

	retList := []interface{}{}
	if dbAccountList != nil {
		ids := make([]string, 0, len(dbAccountList))
		for _, dbA := range dbAccountList {
			listMap := map[string]interface{}{}
			if dbA.UserName != nil {
				listMap["user_name"] = dbA.UserName
			}
			if dbA.Host != nil {
				listMap["host"] = dbA.Host
			}
			if dbA.Description != nil {
				listMap["description"] = dbA.Description
			}
			if dbA.CreateTime != nil {
				listMap["create_time"] = dbA.CreateTime
			}
			if dbA.UpdateTime != nil {
				listMap["update_time"] = dbA.UpdateTime
			}
			if dbA.ReadOnly != nil {
				listMap["read_only"] = dbA.ReadOnly
			}
			if dbA.DelayThresh != nil {
				listMap["delay_thresh"] = dbA.DelayThresh
			}
			//if dbA.SlaveConst != nil {
			//	listMap["slave_const"] = dbA.SlaveConst
			//}
			ids = append(ids, *dbA.UserName+FILED_SP+*dbA.Host)
			retList = append(retList, listMap)
		}

		d.SetId(helper.DataResourceIdsHash(ids))
		_ = d.Set("list", retList)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), retList); e != nil {
			return e
		}
	}

	return nil
}
