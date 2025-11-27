/*
Provides a resource to create a dcdb account

# Example Usage

```hcl

	resource "tencentcloudenterprise_dcdb_account" "account" {
		instance_id = "tdsqlshard-973xatu3"
		user_name = "mysql"
		host = "127.0.0.1"
		password = "===password==="
		read_only = 0
		description = "test modify ---this is a test account"
	}

```
Import

dcdb account can be imported using the id, e.g.
```
$ terraform import cloud_dcdb_account.account account_id
```
*/
package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dcdb_account", CNDescription{
		TerraformTypeCN: "数据库账号",
		DescriptionCN:   "提供DCDB数据库账号资源，用于创建和管理DCDB数据库账号。",
		AttributesCN: map[string]string{
			"instance_id":  "实例id",
			"user_name":    "帐户名称",
			"host":         "db主机",
			"password":     "暗语",
			"read_only":    "帐户是否为只读0表示不是只读帐户",
			"description":  "帐户的描述",
			"delay_thresh": "如果备份机器的延迟超过此参数中设置的值，系统将认为备份机器发生故障建议将参数值设置为大于10当ReadOnly设置为1或2时，此参数将生效",
		},
	})
}

func resourceTencentCloudDcdbAccount() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a dcdb account",
		Read:        resourceTencentCloudDcdbAccountRead,
		Create:      resourceTencentCloudDcdbAccountCreate,
		Update:      resourceTencentCloudDcdbAccountUpdate,
		Delete:      resourceTencentCloudDcdbAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},

			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account name.",
			},

			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Db host.",
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password.",
				Sensitive:   true,
			},

			"read_only": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether the account is readonly. 0 means not a readonly account.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description for account.",
			},
			"delay_thresh": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "If the delay of the backup machine exceeds the value set in this parameter, " +
					"the system will consider the backup machine to have a failure. " +
					"It is recommended to set the parameter value greater than 10. " +
					"This parameter takes effect when ReadOnly is set to 1 or 2.",
			},
		},
	}
}

func resourceTencentCloudDcdbAccountCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_account.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = dcdb.NewCreateAccountRequest()
		response   *dcdb.CreateAccountResponse
		instanceId string
		userName   string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_name"); ok {
		userName = v.(string)
		request.UserName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("host"); ok {

		request.Host = helper.String(v.(string))
	}

	if v, ok := d.GetOk("password"); ok {

		request.Password = helper.String(v.(string))
	}

	if v, ok := d.GetOk("read_only"); ok {
		request.ReadOnly = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("description"); ok {

		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("delay_thresh"); ok {
		request.DelayThresh = helper.Int64(int64(v.(int)))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().CreateAccount(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		if result.Response == nil {
			return retryError(errors.New("response is nil"))
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dcdb account failed, reason:%+v", logId, err)
		return err
	}

	instanceId = *response.Response.InstanceId

	d.SetId(instanceId + FILED_SP + userName)
	return resourceTencentCloudDcdbAccountRead(d, meta)
}

func resourceTencentCloudDcdbAccountRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_account.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	userName := idSplit[1]

	accounts, err := service.DescribeDcdbAccount(ctx, instanceId, userName)

	if err != nil {
		return err
	}

	if len(accounts) == 0 {
		d.SetId("")
		return fmt.Errorf("resource `account` %s does not exist", instanceId)
	}

	log.Printf("[DEBUG]cloud_dcdb_account.read Users:%v", accounts[0])
	if accounts[0].UserName != nil {
		_ = d.Set("user_name", accounts[0].UserName)
	}

	if accounts[0].Host != nil {
		_ = d.Set("host", accounts[0].Host)
	}

	if accounts[0].ReadOnly != nil {
		_ = d.Set("read_only", accounts[0].ReadOnly)
	}

	if accounts[0].Description != nil {
		_ = d.Set("description", accounts[0].Description)
	}
	if accounts[0].DelayThresh != nil {
		_ = d.Set("delay_thresh", accounts[0].DelayThresh)
	}

	return nil
}

func resourceTencentCloudDcdbAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_account.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.TODO(), logIdKey, logId)

	request := dcdb.NewModifyAccountDescriptionRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	instanceId := idSplit[0]
	userName := idSplit[1]

	request.InstanceId = &instanceId
	request.UserName = &userName

	if d.HasChange("instance_id") {
		return fmt.Errorf("`instance_id` do not support change now.")
	}

	if d.HasChange("user_name") {
		return fmt.Errorf("`user_name` do not support change now.")
	}

	if d.HasChange("password") {
		// return fmt.Errorf("`password` do not support change now.")
		if v, ok := d.GetOk("password"); ok {
			request := dcdb.NewResetAccountPasswordRequest()
			request.InstanceId = &instanceId
			request.UserName = &userName
			if v, ok := d.GetOk("host"); ok {
				request.Host = helper.String(v.(string))
			}
			request.Password = helper.String(v.(string))

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().ResetAccountPassword(request)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId,
						request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s operate dcdb resetAccountPasswordOperation failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	if d.HasChange("read_only") {
		return fmt.Errorf("`read_only` do not support change now.")
	}

	if d.HasChange("description") {
		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}
		if v, ok := d.GetOk("host"); ok {
			request.Host = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseDcdbClient().ModifyAccountDescription(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dcdb account failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDcdbAccountRead(d, meta)
}

func resourceTencentCloudDcdbAccountDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_dcdb_account.delete")()
	defer inconsistentCheck(d, meta)()
	var host string
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken, %s", d.Id())
	}

	if v, ok := d.GetOk("host"); ok {
		host = v.(string)
	} else {
		return fmt.Errorf("host is broken, %s", host)
	}

	instanceId := idSplit[0]
	userName := idSplit[1]

	if err := service.DeleteDcdbAccountById(ctx, instanceId, userName, host); err != nil {
		return err
	}

	conf := BuildStateChangeConf([]string{}, []string{"deleted"}, readRetryTimeout, time.Second,
		service.DcdbAccountRefreshFunc(instanceId, userName, []string{}))
	if _, e := conf.WaitForStateContext(ctx); e != nil {
		return e
	}

	return nil
}
