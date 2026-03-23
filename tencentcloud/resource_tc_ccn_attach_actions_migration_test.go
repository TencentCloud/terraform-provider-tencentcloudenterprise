package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	ccn "terraform-provider-tencentcloudenterprise/sdk/ccn/v20170312"
	common2 "terraform-provider-tencentcloudenterprise/sdk/common"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
)

type testAccCcnCrossAccountAttachmentContext struct {
	ownerCcnID     string
	ownerUIN       string
	instanceID     string
	instanceRegion string
	instanceType   string
}

func TestAccTencentCloudCcnInstancesAcceptAttach_basic(t *testing.T) {
	ccnName := fmt.Sprintf("ci-temp-test-ccn-accept-%d", time.Now().UnixNano())
	vpcName := fmt.Sprintf("ci-temp-test-vpc-accept-%d", time.Now().UnixNano())
	keyName := "tencentcloudenterprise_ccn_instances_accept_attach.example"
	attachCtx := &testAccCcnCrossAccountAttachmentContext{}

	defer func() {
		if err := testAccCleanupCcnCrossAccountAttachment(attachCtx); err != nil {
			t.Logf("cleanup accepted ccn attachment failed: %v", err)
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCcnCrossAccount(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnCrossAccountBaseConfig(ccnName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccPreparePendingCcnAttachment("tencentcloudenterprise_ccn.owner", "tencentcloudenterprise_vpc.sub", "data.tencentcloudenterprise_user_info.owner", attachCtx),
					testAccCheckRemoteCcnAttachmentState(attachCtx, "PENDING"),
				),
			},
			{
				Config: testAccCcnAcceptAttachConfig(ccnName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(keyName, "id"),
					testAccCheckRemoteCcnAttachmentState(attachCtx, "ACTIVE"),
				),
			},
			{
				PreConfig: func() {
					if err := testAccCleanupCcnCrossAccountAttachment(attachCtx); err != nil {
						t.Fatalf("cleanup accepted ccn attachment failed: %v", err)
					}
				},
				Config: testAccCcnCrossAccountBaseConfig(ccnName, vpcName),
			},
		},
	})
}

func TestAccTencentCloudCcnInstancesRejectAttach_basic(t *testing.T) {
	ccnName := fmt.Sprintf("ci-temp-test-ccn-reject-%d", time.Now().UnixNano())
	vpcName := fmt.Sprintf("ci-temp-test-vpc-reject-%d", time.Now().UnixNano())
	keyName := "tencentcloudenterprise_ccn_instances_reject_attach.example"
	attachCtx := &testAccCcnCrossAccountAttachmentContext{}

	defer func() {
		if err := testAccCleanupCcnCrossAccountAttachment(attachCtx); err != nil {
			t.Logf("cleanup rejected ccn attachment failed: %v", err)
		}
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCcnCrossAccount(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnCrossAccountBaseConfig(ccnName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccPreparePendingCcnAttachment("tencentcloudenterprise_ccn.owner", "tencentcloudenterprise_vpc.sub", "data.tencentcloudenterprise_user_info.owner", attachCtx),
					testAccCheckRemoteCcnAttachmentState(attachCtx, "PENDING"),
				),
			},
			{
				Config: testAccCcnRejectAttachConfig(ccnName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(keyName, "id"),
					testAccCheckRemoteCcnAttachmentState(attachCtx, "REJECTED"),
				),
			},
			{
				PreConfig: func() {
					if err := testAccCleanupCcnCrossAccountAttachment(attachCtx); err != nil {
						t.Fatalf("cleanup rejected ccn attachment failed: %v", err)
					}
				},
				Config: testAccCcnCrossAccountBaseConfig(ccnName, vpcName),
			},
		},
	})
}

func testAccPreCheckCcnCrossAccount(t *testing.T) {
	t.Helper()

	if v := os.Getenv(PROVIDER_DOMAIN); v == "" {
		os.Setenv(PROVIDER_DOMAIN, "api3.yfm4.fsphere.cn")
	}
	if v := os.Getenv(PROVIDER_PROTOCOL); v == "" {
		os.Setenv(PROVIDER_PROTOCOL, "HTTP")
	}
	if v := os.Getenv(PROVIDER_CSP_DOMAIN); v == "" {
		os.Setenv(PROVIDER_CSP_DOMAIN, "csp.yfm4.fsphere.cn")
	}
	if v := os.Getenv(PROVIDER_COS_DOMAIN); v == "" {
		os.Setenv(PROVIDER_COS_DOMAIN, "cospub.yfm4.fsphere.cn")
	}
	if v := os.Getenv(PROVIDER_REGION); v == "" {
		os.Setenv(PROVIDER_REGION, defaultRegion)
	}

	ownerSecretID := os.Getenv(PROVIDER_SECRET_ID)
	ownerSecretKey := os.Getenv(PROVIDER_SECRET_KEY)
	if ownerSecretID == "" || ownerSecretKey == "" {
		ownerSecretID = os.Getenv(COMMON_PROVIDER_SECRET_ID)
		ownerSecretKey = os.Getenv(COMMON_PROVIDER_SECRET_KEY)
	}
	if ownerSecretID == "" || ownerSecretKey == "" {
		t.Fatalf("%s/%s or %s/%s must be set for CCN cross-account acceptance tests\n",
			PROVIDER_SECRET_ID, PROVIDER_SECRET_KEY, COMMON_PROVIDER_SECRET_ID, COMMON_PROVIDER_SECRET_KEY)
	}

	subSecretID := os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_ID)
	subSecretKey := os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_KEY)
	if subSecretID == "" || subSecretKey == "" {
		t.Fatalf("%s and %s must be set for CCN cross-account acceptance tests\n",
			SUB_ACCOUNT_PROVIDER_SECRET_ID, SUB_ACCOUNT_PROVIDER_SECRET_KEY)
	}

	os.Setenv(PROVIDER_SECRET_ID, ownerSecretID)
	os.Setenv(PROVIDER_SECRET_KEY, ownerSecretKey)
}

func testAccPreparePendingCcnAttachment(ccnResourceName, vpcResourceName, ownerUserInfoName string, attachCtx *testAccCcnCrossAccountAttachmentContext) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ccnRS, ok := s.RootModule().Resources[ccnResourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", ccnResourceName)
		}
		vpcRS, ok := s.RootModule().Resources[vpcResourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", vpcResourceName)
		}
		userInfoRS, ok := s.RootModule().Resources[ownerUserInfoName]
		if !ok {
			return fmt.Errorf("resource %s not found", ownerUserInfoName)
		}

		attachCtx.ownerCcnID = ccnRS.Primary.ID
		attachCtx.instanceID = vpcRS.Primary.ID
		attachCtx.instanceRegion = testAccCcnCrossAccountRegion()
		attachCtx.instanceType = CNN_INSTANCE_TYPE_VPC
		attachCtx.ownerUIN = userInfoRS.Primary.Attributes["owner_uin"]
		if attachCtx.ownerUIN == "" {
			attachCtx.ownerUIN = userInfoRS.Primary.Attributes["uin"]
		}
		if attachCtx.ownerUIN == "" {
			return fmt.Errorf("owner uin not found in %s", ownerUserInfoName)
		}

		info, err := testAccDescribeCrossAccountAttachment(attachCtx)
		if err != nil {
			return err
		}
		if info != nil && info.State != nil {
			if strings.EqualFold(*info.State, "PENDING") {
				return nil
			}
			return fmt.Errorf("unexpected attachment state before action: %s", *info.State)
		}

		logID := getLogId(contextNil)
		request := ccn.NewAttachCcnInstancesRequest()
		request.CcnId = helper.String(attachCtx.ownerCcnID)
		request.CcnUin = helper.String(attachCtx.ownerUIN)
		request.Instances = []*ccn.CcnInstance{
			{
				InstanceId:     helper.String(attachCtx.instanceID),
				InstanceRegion: helper.String(attachCtx.instanceRegion),
				InstanceType:   helper.String(attachCtx.instanceType),
			},
		}

		subClient := testAccCcnSubAccountClient()
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := subClient.apiV3Conn.UseCcnClient().AttachCcnInstances(request)
			if e != nil {
				return retryError(e)
			}

			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logID, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			return nil
		})
		if err != nil {
			return err
		}

		return testAccWaitForCrossAccountAttachmentState(attachCtx, "PENDING")
	}
}

func testAccCheckRemoteCcnAttachmentState(attachCtx *testAccCcnCrossAccountAttachmentContext, expectedState string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return testAccWaitForCrossAccountAttachmentState(attachCtx, expectedState)
	}
}

func testAccWaitForCrossAccountAttachmentState(attachCtx *testAccCcnCrossAccountAttachmentContext, expectedState string) error {
	return resource.Retry(readRetryTimeout, func() *resource.RetryError {
		info, err := testAccDescribeCrossAccountAttachment(attachCtx)
		if err != nil {
			return retryError(err)
		}
		if info == nil || info.State == nil {
			return resource.RetryableError(fmt.Errorf("cross-account ccn attachment not ready"))
		}
		if !strings.EqualFold(*info.State, expectedState) {
			return resource.RetryableError(fmt.Errorf("cross-account ccn attachment state is %s, expect %s", *info.State, expectedState))
		}
		return nil
	})
}

func testAccDescribeCrossAccountAttachment(attachCtx *testAccCcnCrossAccountAttachmentContext) (*ccn.CcnAttachedInstance, error) {
	logID := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logID)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	return service.DescribeCcnAttachedInstanceByFilter(
		ctx,
		attachCtx.ownerCcnID,
		attachCtx.instanceType,
		attachCtx.instanceRegion,
		attachCtx.instanceID,
	)
}

func testAccCleanupCcnCrossAccountAttachment(attachCtx *testAccCcnCrossAccountAttachmentContext) error {
	if attachCtx.ownerCcnID == "" || attachCtx.instanceID == "" || attachCtx.instanceRegion == "" || attachCtx.instanceType == "" {
		return nil
	}

	info, err := testAccDescribeCrossAccountAttachment(attachCtx)
	if err != nil {
		return err
	}
	if info == nil || info.State == nil {
		return nil
	}

	switch strings.ToUpper(*info.State) {
	case "REJECTED", "DELETED":
		return nil
	}

	logID := getLogId(contextNil)
	request := ccn.NewDetachCcnInstancesRequest()
	request.CcnId = helper.String(attachCtx.ownerCcnID)
	request.Instances = []*ccn.CcnInstance{
		{
			InstanceId:     helper.String(attachCtx.instanceID),
			InstanceRegion: helper.String(attachCtx.instanceRegion),
			InstanceType:   helper.String(attachCtx.instanceType),
		},
	}

	subClient := testAccCcnSubAccountClient()
	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := subClient.apiV3Conn.UseCcnClient().DetachCcnInstances(request)
		if e != nil {
			latest, describeErr := testAccDescribeCrossAccountAttachment(attachCtx)
			if describeErr == nil && latest == nil {
				return nil
			}
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logID, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		return err
	}

	return resource.Retry(readRetryTimeout, func() *resource.RetryError {
		latest, e := testAccDescribeCrossAccountAttachment(attachCtx)
		if e != nil {
			return retryError(e)
		}
		if latest != nil {
			return resource.RetryableError(fmt.Errorf("cross-account ccn attachment still exists during cleanup"))
		}
		return nil
	})
}

func testAccCcnSubAccountClient() *TencentCloudClient {
	subSecretID := os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_ID)
	subSecretKey := os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_KEY)
	baseClient := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn

	apiConn := &connectivity.TencentCloudClient{
		Credential: common.NewTokenCredential(subSecretID, subSecretKey, ""),
		CredentialTce: common2.NewTokenCredential(
			subSecretID,
			subSecretKey,
			"",
		),
		Region:    baseClient.Region,
		Protocol:  baseClient.Protocol,
		Domain:    baseClient.Domain,
		CspDomain: baseClient.CspDomain,
		CosDomain: baseClient.CosDomain,
	}

	return &TencentCloudClient{apiV3Conn: apiConn}
}

func testAccCcnCrossAccountRegion() string {
	if region := os.Getenv(PROVIDER_REGION); region != "" {
		return region
	}
	return defaultRegion
}

func testAccCcnCrossAccountBaseConfig(ccnName, vpcName string) string {
	return fmt.Sprintf(`
variable "region" {
  default = %q
}

provider "cloud" {
  alias      = "sub"
  secret_id  = %q
  secret_key = %q
}

data "tencentcloudenterprise_user_info" "owner" {}

resource "tencentcloudenterprise_ccn" "owner" {
  name        = %q
  description = "ci-temp-test-ccn-attach-actions"
  qos         = "AG"
}

resource "tencentcloudenterprise_vpc" "sub" {
  provider     = cloud.sub
  name         = %q
  cidr_block   = "10.0.0.0/16"
  dns_servers  = ["119.29.29.29", "8.8.8.8"]
  is_multicast = false
}
`, testAccCcnCrossAccountRegion(), os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_ID), os.Getenv(SUB_ACCOUNT_PROVIDER_SECRET_KEY), ccnName, vpcName)
}

func testAccCcnAcceptAttachConfig(ccnName, vpcName string) string {
	return testAccCcnCrossAccountBaseConfig(ccnName, vpcName) + `
resource "tencentcloudenterprise_ccn_instances_accept_attach" "example" {
  ccn_id = tencentcloudenterprise_ccn.owner.id

  instances {
    instance_id     = tencentcloudenterprise_vpc.sub.id
    instance_region = var.region
    instance_type   = "VPC"
  }
}
`
}

func testAccCcnRejectAttachConfig(ccnName, vpcName string) string {
	return testAccCcnCrossAccountBaseConfig(ccnName, vpcName) + `
resource "tencentcloudenterprise_ccn_instances_reject_attach" "example" {
  ccn_id = tencentcloudenterprise_ccn.owner.id

  instances {
    instance_id     = tencentcloudenterprise_vpc.sub.id
    instance_region = var.region
    instance_type   = "VPC"
  }
}
`
}
