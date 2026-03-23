package tencentcloud

import (
	"context"
	"encoding/json"
	"log"

	open "terraform-provider-tencentcloudenterprise/sdk/open/v20201202"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

type OpenService struct {
	client *connectivity.TencentCloudClient
}

// DescribeOidcConfigByName retrieves OIDC configuration by IdP name
func (me *OpenService) DescribeOidcConfigByName(ctx context.Context, idpName string) (oidcConfig *OidcConfigDetail, errRet error) {
	logId := getLogId(ctx)

	request := open.NewListIdentityProviderRequest()
	response := open.NewListIdentityProviderResponse()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseOpenClient().ListIdentityProvider(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil || response.Response.Item == nil || response.Response.Item.Data == nil {
		return
	}

	// Iterate through the provider list to find the OIDC config with matching name
	for _, item := range response.Response.Item.Data.List {
		if item == nil {
			continue
		}

		// Check if this is an OIDC provider
		if item.Oidc == nil || *item.Oidc == "" {
			continue
		}

		// Check if the IdpName matches
		if item.Name != nil && *item.Name == idpName {
			// Parse OIDC JSON string
			var oidcData map[string]interface{}
			if err := json.Unmarshal([]byte(*item.Oidc), &oidcData); err != nil {
				log.Printf("[WARN]%s failed to parse OIDC JSON for item %v: %s\n", logId, item.Id, err.Error())
				continue
			}

			oidcConfig = &OidcConfigDetail{
				Id:      item.Id,
				IdpName: item.Name,
			}

			// Extract OIDC-specific fields from the JSON
			// Based on actual API response, these fields are in the Oidc JSON string:
			if val, ok := oidcData["IdentityUrl"].(string); ok {
				oidcConfig.IdentityUrl = &val
			}
			if val, ok := oidcData["IdentityKey"].(string); ok {
				oidcConfig.IdentityKey = &val
			}
			if val, ok := oidcData["ClientId"].(string); ok {
				oidcConfig.ClientId = &val
			}
			if val, ok := oidcData["AuthorizationEndpoint"].(string); ok {
				oidcConfig.AuthorizationEndpoint = &val
			}
			if val, ok := oidcData["ResponseType"].(string); ok {
				oidcConfig.ResponseType = &val
			}
			if val, ok := oidcData["ResponseMode"].(string); ok {
				oidcConfig.ResponseMode = &val
			}
			if val, ok := oidcData["EmailField"].(string); ok {
				oidcConfig.EmailField = &val
			}
			if val, ok := oidcData["NickNameField"].(string); ok {
				oidcConfig.NickNameField = &val
			}
			if val, ok := oidcData["PhoneNumField"].(string); ok {
				oidcConfig.PhoneNumField = &val
			}
			if val, ok := oidcData["LoginAccountField"].(string); ok {
				oidcConfig.LoginAccountField = &val
			}
			if val, ok := oidcData["CountryCodeField"].(string); ok {
				oidcConfig.CountryCodeField = &val
			}

			// Handle Scope array
			if scopeVal, ok := oidcData["Scope"].([]interface{}); ok {
				scope := make([]*string, 0, len(scopeVal))
				for _, s := range scopeVal {
					if str, ok := s.(string); ok {
						scope = append(scope, &str)
					}
				}
				if len(scope) > 0 {
					oidcConfig.Scope = scope
				}
			}

			// Fields that are NOT in the Oidc JSON, but in the item level:
			// Protocol - not in response, need to set from request or default
			protocol := "oidc" // Default protocol
			oidcConfig.Protocol = &protocol

			// Remark is in item.Desc, not in Oidc JSON
			if item.Desc != nil {
				oidcConfig.Remark = item.Desc
			}

			// IsSyncIdpUser is in item level, not in Oidc JSON
			if item.IsSyncIdpUser != nil {
				oidcConfig.IsSyncIdpUser = item.IsSyncIdpUser
			}

			// LogoutUrl - check if it's in Oidc JSON (API docs unclear)
			if val, ok := oidcData["LogoutUrl"].(string); ok {
				oidcConfig.LogoutUrl = &val
			}

			return
		}
	}

	return
}

// DescribeSamlConfigByName retrieves SAML configuration by IdP name
func (me *OpenService) DescribeSamlConfigByName(ctx context.Context, idpName string) (samlConfig *SamlConfigDetail, errRet error) {
	logId := getLogId(ctx)

	request := open.NewListIdentityProviderRequest()
	response := open.NewListIdentityProviderResponse()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		ratelimit.Check(request.GetAction())
		result, e := me.client.UseOpenClient().ListIdentityProvider(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})

	if err != nil {
		errRet = err
		return
	}

	if response == nil || response.Response == nil || response.Response.Item == nil || response.Response.Item.Data == nil {
		return
	}

	// Iterate through the provider list to find the SAML config with matching name
	for _, item := range response.Response.Item.Data.List {
		if item == nil {
			continue
		}

		// Check if this is a SAML provider
		if item.SamlMetaData == nil || *item.SamlMetaData == "" {
			continue
		}

		// Check if the IdpName matches
		if item.Name != nil && *item.Name == idpName {
			samlConfig = &SamlConfigDetail{
				Id:           item.Id,
				IdpName:      item.Name,
				SamlMetaData: item.SamlMetaData,
			}

			// Protocol - not in response, need to set from request or default
			protocol := "saml" // Default protocol
			samlConfig.Protocol = &protocol

			// Remark is in item.Desc
			if item.Desc != nil {
				samlConfig.Remark = item.Desc
			}

			// IsSyncIdpUser is in item level
			if item.IsSyncIdpUser != nil {
				samlConfig.IsSyncIdpUser = item.IsSyncIdpUser
			}

			// AssistDomain is in item level
			if item.AssistDomain != nil {
				samlConfig.AssistDomain = item.AssistDomain
			}

			return
		}
	}

	return
}
