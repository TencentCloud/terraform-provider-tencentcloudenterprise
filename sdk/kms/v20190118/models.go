// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190118

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type GetParametersForImportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的唯一标识，用于指定目标导入密钥材料的CMK。

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 导入密钥材料需要的token，用于作为 ImportKeyMaterial 的参数。

		ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`
		// 用于加密密钥材料的RSA公钥，base64编码。使用PublicKey base64解码后的公钥将导入密钥进行加密后作为 ImportKeyMaterial 的参数。

		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
		// 该导出token和公钥的有效期，超过该时间后无法导入，需要重新调用GetParametersForImport获取。

		ParametersValidTo *uint64 `json:"ParametersValidTo,omitempty" name:"ParametersValidTo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetParametersForImportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetParametersForImportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAlgorithmsRequest struct {
	*tchttp.BaseRequest
}

func (r *ListAlgorithmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAlgorithmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyRequest struct {
	*tchttp.BaseRequest

	// CMK全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBaradSubModulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Barad 子模块

		SubModules []*string `json:"SubModules,omitempty" name:"SubModules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetBaradSubModulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBaradSubModulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelKeyArchiveRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *CancelKeyArchiveRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelKeyArchiveRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSMLicensesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回各个版本的License，如果没有配置License，则返回空数组

		Licenses []*TsmLicense `json:"Licenses,omitempty" name:"Licenses"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTSMLicensesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSMLicensesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥属性信息

		KeyMetadata *KeyMetadata `json:"KeyMetadata,omitempty" name:"KeyMetadata"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlgorithmInfo struct {

	// 算法的标识

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// 算法的名称

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

type DescribeWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白盒密钥信息

		KeyInfo *WhiteboxKeyInfo `json:"KeyInfo,omitempty" name:"KeyInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteBoxKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAlgorithmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 本地区支持的对称加密算法

		SymmetricAlgorithms []*AlgorithmInfo `json:"SymmetricAlgorithms,omitempty" name:"SymmetricAlgorithms"`
		// 本地区支持的非对称加密算法

		AsymmetricAlgorithms []*AlgorithmInfo `json:"AsymmetricAlgorithms,omitempty" name:"AsymmetricAlgorithms"`
		// 支持的签名验签算法

		AsymmetricSignVerifyAlgorithms []*AlgorithmInfo `json:"AsymmetricSignVerifyAlgorithms,omitempty" name:"AsymmetricSignVerifyAlgorithms"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAlgorithmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListAlgorithmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2EncryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加密后经过base64编码的密文

		Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
		// 加密使用的CMK的全局唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsymmetricSm2EncryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricSm2EncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 设备指纹列表

		DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitempty" name:"DeviceFingerprints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateAsymmetricKeyPairRequest struct {
	*tchttp.BaseRequest

	// CMK全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 指定生成非对称密钥对算法。支持：SM2

	KeySpec *string `json:"KeySpec,omitempty" name:"KeySpec"`
	// X509格式的pem类型公钥字符串， 支持 RSA2048 和 SM2 公钥， 用于对返回的明文私钥进行加密。 若为空，则不对明文密钥对加密。

	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`
	// 非对称加密算法，配合  EncryptionPublicKey  对返回数据进行加密。 目前支持：SM2（C1C3C2）, SM2_C1C3C2_ASN1(返回SM2C1C3C2 ASN1格式密文)，RSAES_PKCS1_V1_5， RSAES_OAEP_SHA_1， RSAES_OAEP_SHA_256。 若为空，则默认为 SM2。

	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

func (r *GenerateAsymmetricKeyPairRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateAsymmetricKeyPairRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 签名是否有效。true：签名有效，false：签名无效。

		SignatureValid *bool `json:"SignatureValid,omitempty" name:"SignatureValid"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyByAsymmetricKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2DecryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 解密后的明文，base64编码

		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsymmetricSm2DecryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricSm2DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableWhiteBoxKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK列表数组

		Keys []*Key `json:"Keys,omitempty" name:"Keys"`
		// CMK的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableWhiteBoxKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// KMS服务是否开通， true 表示已开通

		ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`
		// 服务不可用类型： 0-未购买，1-正常， 2-欠费停服， 3-资源释放

		InvalidType *int64 `json:"InvalidType,omitempty" name:"InvalidType"`
		// 用户等级

		UserLevel *uint64 `json:"UserLevel,omitempty" name:"UserLevel"`
		// 过期时间

		ProExpireTime *uint64 `json:"ProExpireTime,omitempty" name:"ProExpireTime"`
		// 标志

		ProRenewFlag *uint64 `json:"ProRenewFlag,omitempty" name:"ProRenewFlag"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetServiceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImportKeyMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCloudResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindCloudResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRotationRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableKeyRotationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeyRotationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableWhiteBoxKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableWhiteBoxKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteWhiteBoxKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteBoxKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateAsymmetricKeyPairResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的全局唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 若调用时未提供 EncryptionPublicKey，该字段值为P8格式PEM类型私钥明文。 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥对P8格式PEM类型私钥明文进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取P8格式PEM类型私钥明文。

		PrivateKeyPlaintext *string `json:"PrivateKeyPlaintext,omitempty" name:"PrivateKeyPlaintext"`
		// P1格式PEM类型返回的明文公钥。

		PublicKeyPlaintext *string `json:"PublicKeyPlaintext,omitempty" name:"PublicKeyPlaintext"`
		// 使用 CMK 加密私钥后的密文，用户需要自行保存该密文，KMS不托管生成的密钥对。可以通过Decrypt接口从密文中获取密钥明文。

		PrivateKeyCiphertextBlob *string `json:"PrivateKeyCiphertextBlob,omitempty" name:"PrivateKeyCiphertextBlob"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateAsymmetricKeyPairResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateAsymmetricKeyPairResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReEncryptRequest struct {
	*tchttp.BaseRequest

	// 需要重新加密的密文

	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
	// 重新加密使用的CMK，如果为空，则使用密文原有的CMK重新加密（若密钥没有轮换则密文不会刷新）

	DestinationKeyId *string `json:"DestinationKeyId,omitempty" name:"DestinationKeyId"`
	// CiphertextBlob 密文加密时使用的key/value对的json字符串。如果加密时未使用，则为空

	SourceEncryptionContext *string `json:"SourceEncryptionContext,omitempty" name:"SourceEncryptionContext"`
	// 重新加密使用的key/value对的json字符串，如果使用该字段，则返回的新密文在解密时需要填入相同的字符串

	DestinationEncryptionContext *string `json:"DestinationEncryptionContext,omitempty" name:"DestinationEncryptionContext"`
}

func (r *ReEncryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReEncryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncryptByWhiteBoxResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 初始化向量，加密算法会使用到, base64编码。如果由调用方在入参中传入，则原样返回。如果调用方没有传入，则后端服务随机生成，并返回

		InitializationVector *string `json:"InitializationVector,omitempty" name:"InitializationVector"`
		// 加密后的密文，base64编码

		CipherText *string `json:"CipherText,omitempty" name:"CipherText"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EncryptByWhiteBoxResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EncryptByWhiteBoxResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2EncryptRequest struct {
	*tchttp.BaseRequest

	// 调用CreateKey生成的CMK全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持 1024 字节

	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
}

func (r *AsymmetricSm2EncryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricSm2EncryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelKeyArchiveResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelKeyArchiveResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelKeyArchiveResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncryptByWhiteBoxRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 初始化向量，大小为 16 Bytes，加密算法会使用到, base64编码；如果不传，则由后端服务随机生成。用户需要自行保存该值，作为解密的参数。

	InitializationVector *string `json:"InitializationVector,omitempty" name:"InitializationVector"`
	// 待加密的文本， base64编码，文本的原始长度最大不超过4KB

	PlainText *string `json:"PlainText,omitempty" name:"PlainText"`
}

func (r *EncryptByWhiteBoxRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EncryptByWhiteBoxRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetParametersForImportRequest struct {
	*tchttp.BaseRequest

	// CMK的唯一标识，获取密钥参数的CMK必须是EXTERNAL类型，即在CreateKey时指定Type=2 类型的CMK。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 指定加密密钥材料的算法，目前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256

	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" name:"WrappingAlgorithm"`
	// 指定加密密钥材料的类型，目前只支持RSA_2048

	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" name:"WrappingKeySpec"`
}

func (r *GetParametersForImportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetParametersForImportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest

	// 签名算法，支持的算法：SM2DSA（ASN1 格式），SM2DSA_ASN1（ASN1 格式），SM2DSA_RAW（原始格式），ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，消息摘要长度（Base64编码前的长度）必须等于32字节

	Message *string `json:"Message,omitempty" name:"Message"`
	// 密钥的唯一标识

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。

	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
}

func (r *SignByAsymmetricKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SignByAsymmetricKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的属性信息列表

		KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateKeyDescriptionRequest struct {
	*tchttp.BaseRequest

	// 新的描述信息，最大支持1024字节

	Description *string `json:"Description,omitempty" name:"Description"`
	// 需要修改描述信息的CMK ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *UpdateKeyDescriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateKeyDescriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *EnableKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListKeysRequest struct {
	*tchttp.BaseRequest

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 含义跟 SQL 查询的 Limit 一致，表示本次获最多获取 Limit 个元素。缺省值为10，最大值为200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk

	Role *uint64 `json:"Role,omitempty" name:"Role"`
}

func (r *ListKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Key struct {

	// CMK的全局唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

type CancelKeyDeletionRequest struct {
	*tchttp.BaseRequest

	// 需要被取消删除的CMK的唯一标志

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *CancelKeyDeletionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelKeyDeletionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricRsaDecryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 解密后的明文，base64编码

		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AsymmetricRsaDecryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricRsaDecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用于加密的密钥，base64编码

		EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`
		// 用于解密的密钥，base64编码

		DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`
		// 白盒密钥的全局唯一标识符

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误

		TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`
		// 标签操作的返回信息

		TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWhiteBoxKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DecryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的全局唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 若调用时未提供 EncryptionPublicKey，该字段值为 Base64 编码的明文，需进行 Base64 解码以获取明文。 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取明文。

		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DecryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DecryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxServiceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的白盒密钥服务是否可用

		ServiceEnabled *bool `json:"ServiceEnabled,omitempty" name:"ServiceEnabled"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteBoxServiceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxServiceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateRandomRequest struct {
	*tchttp.BaseRequest

	// 生成的随机数的长度。最小值为1， 最大值为1024。

	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`
}

func (r *GenerateRandomRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateRandomRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateRandomResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生成的随机数的明文，该明文使用base64编码，用户需要使用base64解码得到明文。

		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateRandomResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateRandomResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TsmLicense struct {

	// License版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// TSM AppId

	TsmAppId *string `json:"TsmAppId,omitempty" name:"TsmAppId"`
	// 证书，这里返回的是字符串格式

	Cert *string `json:"Cert,omitempty" name:"Cert"`
}

type CreateWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest

	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。Alias不可重复。

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 密钥的描述，最大1024字节

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建密钥所有的算法类型，支持的取值：AES_256,SM4

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// 标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateWhiteBoxKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWhiteBoxKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DisableWhiteBoxKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableWhiteBoxKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTSMLicensesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeTSMLicensesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTSMLicensesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReEncryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 重新加密后的密文

		CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
		// 重新加密使用的CMK

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 重新加密前密文使用的CMK

		SourceKeyId *string `json:"SourceKeyId,omitempty" name:"SourceKeyId"`
		// true表示密文已经重新加密。同一个CMK进行重加密，在密钥没有发生轮换的情况下不会进行实际重新加密操作，返回原密文

		ReEncrypted *bool `json:"ReEncrypted,omitempty" name:"ReEncrypted"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReEncryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReEncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WhiteboxKeyInfo struct {

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字. 不可重复

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 创建者

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 密钥的描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 密钥创建时间，Unix时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 白盒密钥的状态， 取值为：Enabled | Disabled

	Status *string `json:"Status,omitempty" name:"Status"`
	// 创建者

	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 密钥所用的算法类型

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
	// 白盒加密密钥，base64编码

	EncryptKey *string `json:"EncryptKey,omitempty" name:"EncryptKey"`
	// 白盒解密密钥，base64编码

	DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`
	// 资源ID，格式：creatorUin/$creatorUin/$keyId

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 是否有设备指纹与当前密钥绑定

	DeviceFingerprintBind *bool `json:"DeviceFingerprintBind,omitempty" name:"DeviceFingerprintBind"`
}

type GenerateDataKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的全局唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 若调用时未提供 EncryptionPublicKey，该字段值为生成的数据密钥 DataKey 的 Base64 编码的明文，需进行 Base64 解码以获取 DataKey 明文。 若调用时提供了 EncryptionPublicKey，则该字段值为使用 EncryptionPublicKey 公钥进行非对称加密后的 Base64 编码的密文。需在 Base64 解码后，使用用户上传的公钥对应的私钥进行进一步解密，以获取 DataKey 明文。

		Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
		// 数据密钥DataKey加密后的密文，用户需要自行保存该密文，KMS不托管用户的数据密钥。可以通过Decrypt接口从CiphertextBlob中获取数据密钥DataKey明文

		CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenerateDataKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDataKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *GetRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SignByAsymmetricKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 签名，Base64编码

		Signature *string `json:"Signature,omitempty" name:"Signature"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SignByAsymmetricKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SignByAsymmetricKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyRequest struct {
	*tchttp.BaseRequest

	// CMK 的描述，最大1024字节

	Description *string `json:"Description,omitempty" name:"Description"`
	// 作为密钥更容易辨识，更容易被人看懂的别名， 不可为空，1-60个字母数字 - _ 的组合，首字符必须为字母或者数字。以 kms- 作为前缀的用于云产品使用，Alias 不可重复。

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 指定key的用途，默认为  "ENCRYPT_DECRYPT" 表示创建对称加解密密钥，其它支持用途 “ASYMMETRIC_DECRYPT_RSA_2048” 表示创建用于加解密的RSA2048非对称密钥，“ASYMMETRIC_DECRYPT_SM2” 表示创建用于加解密的SM2非对称密钥。可选值：ENCRYPT_DECRYPT，ASYMMETRIC_DECRYPT_RSA_2048，ASYMMETRIC_DECRYPT_SM2，ASYMMETRIC_SIGN_VERIFY_SM2，ASYMMETRIC_SIGN_VERIFY_RSA_2048，ASYMMETRIC_SIGN_VERIFY_ECC

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// 指定key类型，默认为1，1表示默认类型，由KMS创建CMK密钥，2 表示EXTERNAL 类型，该类型需要用户导入密钥材料，参考 GetParametersForImport 和 ImportKeyMaterial 接口

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 所属项目 Id。

	TCEProjectId *string `json:"TCEProjectId,omitempty" name:"TCEProjectId"`
}

func (r *CreateKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DecryptRequest struct {
	*tchttp.BaseRequest

	// 待解密的密文数据

	CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
	// key/value对的json字符串，如果Encrypt指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024字符

	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。

	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`
	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。

	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

func (r *DecryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DecryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源ID列表

		ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
		// 资源ID总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyRequest struct {
	*tchttp.BaseRequest

	// CMK的唯一标识。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *GetPublicKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPublicKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWhiteBoxKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BuyServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BuyServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BuyServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeysRequest struct {
	*tchttp.BaseRequest

	// 需要批量启用的CMK Id 列表， CMK数量最大支持100

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *EnableKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeviceFingerprint struct {

	// 指纹信息，由设备指纹采集工具采集获得，格式满足正则表达式：^[0-9a-f]{8}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{14}[\-][0-9a-f]{16}$

	Identity *string `json:"Identity,omitempty" name:"Identity"`
	// 描述信息，如：IP，设备名称等，最大1024字节

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeWhiteBoxDecryptKeyRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDecryptKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxDecryptKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest

	// 新的别名，1-60个字符或数字的组合

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// CMK的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAliasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableWhiteBoxKeysRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符列表。注意：要确保所有提供的KeyId是格式有效的，没有重复，个数不超过50个，并且都是有效存在的。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *EnableWhiteBoxKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableWhiteBoxKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *GetServiceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetServiceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArchiveKeyRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ArchiveKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArchiveKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListKeyDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 返回的属性信息列表。

		KeyMetadatas []*KeyMetadata `json:"KeyMetadatas,omitempty" name:"KeyMetadatas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListKeyDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListKeyDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type BuyServiceRequest struct {
	*tchttp.BaseRequest

	// 开通服务类型，默认为0，当前只支持0，表示基础版。

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *BuyServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BuyServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxServiceStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeWhiteBoxServiceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxServiceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImportKeyMaterialRequest struct {
	*tchttp.BaseRequest

	// 使用GetParametersForImport 返回的PublicKey加密后的密钥材料base64编码。对于国密版本region的KMS，导入的密钥材料长度要求为 128 bit，FIPS版本region的KMS， 导入的密钥材料长度要求为 256 bit。

	EncryptedKeyMaterial *string `json:"EncryptedKeyMaterial,omitempty" name:"EncryptedKeyMaterial"`
	// 通过调用GetParametersForImport获得的导入令牌。

	ImportToken *string `json:"ImportToken,omitempty" name:"ImportToken"`
	// 密钥材料过期时间 unix 时间戳，不指定或者 0 表示密钥材料不会过期，若指定过期时间，需要大于当前时间点，最大支持 2147443200。

	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
	// 指定导入密钥材料的CMK，需要和GetParametersForImport 指定的CMK相同。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *ImportKeyMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ImportKeyMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArchiveKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArchiveKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArchiveKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListKeyDetailRequest struct {
	*tchttp.BaseRequest

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为10，最大值为200

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据创建者角色筛选，默认 0 表示用户自己创建的cmk， 1 表示授权其它云产品自动创建的cmk

	Role *uint64 `json:"Role,omitempty" name:"Role"`
	// 根据CMK创建时间排序， 0 表示按照降序排序，1表示按照升序排序

	OrderType *uint64 `json:"OrderType,omitempty" name:"OrderType"`
	// 根据CMK状态筛选， 0表示全部CMK， 1 表示仅查询Enabled CMK， 2 表示仅查询Disabled CMK，3 表示查询PendingDelete 状态的CMK(处于计划删除状态的Key)，4 表示查询 PendingImport 状态的CMK

	KeyState *uint64 `json:"KeyState,omitempty" name:"KeyState"`
	// 根据KeyId或者Alias进行模糊匹配查询

	SearchKeyAlias *string `json:"SearchKeyAlias,omitempty" name:"SearchKeyAlias"`
	// 根据CMK类型筛选， "TENCENT_KMS" 表示筛选密钥材料由KMS创建的CMK， "EXTERNAL" 表示筛选密钥材料需要用户导入的 EXTERNAL类型CMK，"ALL" 或者不设置表示两种类型都查询，大小写敏感。

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 根据CMK的KeyUsage筛选，ALL表示筛选全部，可使用的参数为：ALL 或 ENCRYPT_DECRYPT 或 ASYMMETRIC_DECRYPT_RSA_2048 或 ASYMMETRIC_DECRYPT_SM2，为空则默认筛选ENCRYPT_DECRYPT类型

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// 标签过滤条件

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *ListKeyDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListKeyDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCloudResourceRequest struct {
	*tchttp.BaseRequest

	// cmk的ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 云产品的唯一性标识符

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *BindCloudResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCloudResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetBaradSubModulesRequest struct {
	*tchttp.BaseRequest
}

func (r *GetBaradSubModulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetBaradSubModulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncryptRequest struct {
	*tchttp.BaseRequest

	// 调用CreateKey生成的CMK全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 被加密的明文数据，该字段必须使用base64编码，原文最大长度支持4K

	Plaintext *string `json:"Plaintext,omitempty" name:"Plaintext"`
	// key/value对的json字符串，如果指定了该参数，则在调用Decrypt API时需要提供同样的参数，最大支持1024个字符

	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
}

func (r *EncryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EncryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricRsaDecryptRequest struct {
	*tchttp.BaseRequest

	// CMK的唯一标识

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 使用PublicKey加密的密文，Base64编码

	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
	// 在使用公钥加密时对应的算法：当前支持RSAES_PKCS1_V1_5、RSAES_OAEP_SHA_1、RSAES_OAEP_SHA_256

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

func (r *AsymmetricRsaDecryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricRsaDecryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DescribeWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRotationRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 轮换周期，最小单位：天，取值范围：7~365，默认值为365

	RotateDays *uint64 `json:"RotateDays,omitempty" name:"RotateDays"`
}

func (r *EnableKeyRotationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeyRotationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableWhiteBoxKeyRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥的全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *EnableWhiteBoxKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableWhiteBoxKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverwriteWhiteBoxDeviceFingerprintsRequest struct {
	*tchttp.BaseRequest

	// 白盒密钥ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 设备指纹列表，如果列表为空，则表示删除该密钥对应的所有指纹信息。列表最大长度不超过200。

	DeviceFingerprints []*DeviceFingerprint `json:"DeviceFingerprints,omitempty" name:"DeviceFingerprints"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverwriteWhiteBoxDeviceFingerprintsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRotationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeyRotationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelKeyDeletionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一标志被取消删除的CMK。

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelKeyDeletionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteImportedKeyMaterialResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteImportedKeyMaterialResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImportedKeyMaterialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxDecryptKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白盒解密密钥，base64编码

		DecryptKey *string `json:"DecryptKey,omitempty" name:"DecryptKey"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteBoxDecryptKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxDecryptKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxKeyDetailsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：密钥的状态，0：disabled，1：enabled

	KeyStatus *int64 `json:"KeyStatus,omitempty" name:"KeyStatus"`
	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为0, 表示不分页

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 标签过滤条件

	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

func (r *DescribeWhiteBoxKeyDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxKeyDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleKeyDeletionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 计划删除执行时间

		DeletionDate *uint64 `json:"DeletionDate,omitempty" name:"DeletionDate"`
		// 唯一标志被计划删除的CMK

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScheduleKeyDeletionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScheduleKeyDeletionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcZoneData struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// vpc节点地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribeResourceIdsRequest struct {
	*tchttp.BaseRequest

	// 含义跟 SQL 查询的 Offset 一致，表示本次获取从按一定顺序排列数组的第 Offset 个元素开始，缺省为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 含义跟 SQL 查询的 Limit 一致，表示本次最多获取 Limit 个元素。缺省值为0, 表示不分页

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeResourceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWhiteBoxKeyDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白盒密钥信息列表

		KeyInfos []*WhiteboxKeyInfo `json:"KeyInfos,omitempty" name:"KeyInfos"`
		// key总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWhiteBoxKeyDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWhiteBoxKeyDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCloudResourceRequest struct {
	*tchttp.BaseRequest

	// cmk的ID

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 云产品的唯一性标识符

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 资源/实例ID，由调用方根据自己的云产品特征来定义，以字符串形式做存储。

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *UnbindCloudResourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindCloudResourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableKeyRotationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableKeyRotationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableKeyRotationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可用region列表

		Regions []*string `json:"Regions,omitempty" name:"Regions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverwriteWhiteBoxDeviceFingerprintsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OverwriteWhiteBoxDeviceFingerprintsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateKeyDescriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateKeyDescriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateKeyDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCloudResourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindCloudResourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindCloudResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeysRequest struct {
	*tchttp.BaseRequest

	// 需要批量禁用的CMK Id 列表，CMK数量最大支持100

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DisableKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableWhiteBoxKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableWhiteBoxKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableWhiteBoxKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduleKeyDeletionRequest struct {
	*tchttp.BaseRequest

	// CMK的唯一标志

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 计划删除时间区间[7,30]

	PendingWindowInDays *uint64 `json:"PendingWindowInDays,omitempty" name:"PendingWindowInDays"`
}

func (r *ScheduleKeyDeletionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScheduleKeyDeletionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagFilter struct {

	// 标签键

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签值

	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type DeleteImportedKeyMaterialRequest struct {
	*tchttp.BaseRequest

	// 指定需要删除密钥材料的EXTERNAL CMK。

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DeleteImportedKeyMaterialRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteImportedKeyMaterialRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetPublicKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的唯一标识。

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 经过base64编码的公钥内容。

		PublicKey *string `json:"PublicKey,omitempty" name:"PublicKey"`
		// PEM格式的公钥内容。

		PublicKeyPem *string `json:"PublicKeyPem,omitempty" name:"PublicKeyPem"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetPublicKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetPublicKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EncryptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加密后经过base64编码的密文

		CiphertextBlob *string `json:"CiphertextBlob,omitempty" name:"CiphertextBlob"`
		// 加密使用的CMK的全局唯一标识

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EncryptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EncryptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetKeyRotationStatusRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *GetKeyRotationStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetKeyRotationStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AsymmetricSm2DecryptRequest struct {
	*tchttp.BaseRequest

	// CMK的唯一标识

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 使用PublicKey加密的密文，Base64编码。密文长度不能超过 2048 字节。

	Ciphertext *string `json:"Ciphertext,omitempty" name:"Ciphertext"`
}

func (r *AsymmetricSm2DecryptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AsymmetricSm2DecryptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKeysRequest struct {
	*tchttp.BaseRequest

	// 查询CMK的ID列表，批量查询一次最多支持100个KeyId

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *DescribeKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableKeyRequest struct {
	*tchttp.BaseRequest

	// CMK唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
}

func (r *DisableKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableWhiteBoxKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableWhiteBoxKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableWhiteBoxKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyByAsymmetricKeyRequest struct {
	*tchttp.BaseRequest

	// 密钥的唯一标识

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 签名值，通过调用KMS签名接口生成

	SignatureValue *string `json:"SignatureValue,omitempty" name:"SignatureValue"`
	// 消息原文或消息摘要。如果提供的是消息原文，则消息原文的长度（Base64编码前的长度）不超过4096字节。如果提供的是消息摘要，则消息摘要长度（Base64编码前的长度）必须等于32字节

	Message *string `json:"Message,omitempty" name:"Message"`
	// 消息类型：RAW，DIGEST，如果不传，默认为RAW，表示消息原文。

	MessageType *string `json:"MessageType,omitempty" name:"MessageType"`
	// 签名算法，支持的算法：SM2DSA，ECC_P256_R1，RSA_PSS_SHA_256，RSA_PKCS1_SHA_256

	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

func (r *VerifyByAsymmetricKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyByAsymmetricKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateKeyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// CMK的全局唯一标识符

		KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
		// 作为密钥更容易辨识，更容易被人看懂的别名

		Alias *string `json:"Alias,omitempty" name:"Alias"`
		// 密钥创建时间，unix时间戳

		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
		// CMK的描述

		Description *string `json:"Description,omitempty" name:"Description"`
		// CMK的状态

		KeyState *string `json:"KeyState,omitempty" name:"KeyState"`
		// CMK的用途

		KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
		// 标签操作的返回码. 0: 成功；1: 内部错误；2: 业务处理错误

		TagCode *uint64 `json:"TagCode,omitempty" name:"TagCode"`
		// 标签操作的返回信息

		TagMsg *string `json:"TagMsg,omitempty" name:"TagMsg"`
		// 项目绑定结果：0 成功，1 失败。

		ProjectBindingResultCode *uint64 `json:"ProjectBindingResultCode,omitempty" name:"ProjectBindingResultCode"`
		// 项目绑定结果描述。

		ProjectBindingMessage *string `json:"ProjectBindingMessage,omitempty" name:"ProjectBindingMessage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKeyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenerateDataKeyRequest struct {
	*tchttp.BaseRequest

	// CMK全局唯一标识符

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 指定生成Datakey的加密算法以及Datakey大小，AES_128或者AES_256。KeySpec 和 NumberOfBytes 必须指定一个

	KeySpec *string `json:"KeySpec,omitempty" name:"KeySpec"`
	// 生成的DataKey的长度，同时指定NumberOfBytes和KeySpec时，以NumberOfBytes为准。最小值为1， 最大值为1024。KeySpec 和 NumberOfBytes 必须指定一个

	NumberOfBytes *uint64 `json:"NumberOfBytes,omitempty" name:"NumberOfBytes"`
	// key/value对的json字符串，如果使用该字段，则返回的DataKey在解密时需要填入相同的字符串

	EncryptionContext *string `json:"EncryptionContext,omitempty" name:"EncryptionContext"`
	// PEM 格式公钥字符串，支持 RSA2048 和 SM2 公钥，用于对返回数据中的 Plaintext 值进行加密。若为空，则不对 Plaintext 值加密。

	EncryptionPublicKey *string `json:"EncryptionPublicKey,omitempty" name:"EncryptionPublicKey"`
	// 非对称加密算法，配合 EncryptionPublicKey 对返回数据进行加密。目前支持：SM2（以 C1C3C2 格式返回密文），SM2_C1C3C2_ASN1 （以 C1C3C2 ASN1 格式返回密文），RSAES_PKCS1_V1_5，RSAES_OAEP_SHA_1，RSAES_OAEP_SHA_256。若为空，则默认为 SM2。

	EncryptionAlgorithm *string `json:"EncryptionAlgorithm,omitempty" name:"EncryptionAlgorithm"`
}

func (r *GenerateDataKeyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenerateDataKeyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetKeyRotationStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 密钥轮换是否开启

		KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetKeyRotationStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetKeyRotationStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyMetadata struct {

	// CMK的全局唯一标识

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 作为密钥更容易辨识，更容易被人看懂的别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 密钥创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// CMK的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// CMK的状态， 取值为：Enabled | Disabled | PendingDelete | PendingImport

	KeyState *string `json:"KeyState,omitempty" name:"KeyState"`
	// CMK用途，取值为: ENCRYPT_DECRYPT | ASYMMETRIC_DECRYPT_RSA_2048 | ASYMMETRIC_DECRYPT_SM2

	KeyUsage *string `json:"KeyUsage,omitempty" name:"KeyUsage"`
	// CMK类型，2 表示符合FIPS标准，4表示符合国密标准

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 创建者

	CreatorUin *uint64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 是否开启了密钥轮换功能

	KeyRotationEnabled *bool `json:"KeyRotationEnabled,omitempty" name:"KeyRotationEnabled"`
	// CMK的创建者，用户创建的为 user，授权各云产品自动创建的为对应的产品名

	Owner *string `json:"Owner,omitempty" name:"Owner"`
	// 在密钥轮换开启状态下，下次轮换的时间

	NextRotateTime *uint64 `json:"NextRotateTime,omitempty" name:"NextRotateTime"`
	// 计划删除的时间

	DeletionDate *uint64 `json:"DeletionDate,omitempty" name:"DeletionDate"`
	// CMK 密钥材料类型，由KMS创建的为： TENCENT_KMS， 由用户导入的类型为：EXTERNAL

	Origin *string `json:"Origin,omitempty" name:"Origin"`
	// 在Origin为  EXTERNAL 时有效，表示密钥材料的有效日期， 0 表示不过期

	ValidTo *uint64 `json:"ValidTo,omitempty" name:"ValidTo"`
	// 资源ID，格式：creatorUin/$creatorUin/$keyId

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 轮换天数

	RotateDays *uint64 `json:"RotateDays,omitempty" name:"RotateDays"`
	// 上一次轮换的时间

	LastRotateTime *uint64 `json:"LastRotateTime,omitempty" name:"LastRotateTime"`
}
