# Terraform Provider For TencentCloudEnterprise

<div>
  <p>
    <br>
      Tencent Cloud Enterprise(TCE) Infrastructure Automation for Terraform.
    <br>
  </p>
</div>

## Website

- [Tencent Cloud Enterprise](https://www.tencentcloud.com/solutions/tce)
- [腾讯专有云](https://cloud.tencent.com/solution/tce)

## Requirements

* [Terraform](https://www.terraform.io/downloads.html) 1.5.x
* [Go](https://golang.org/doc/install) 1.13.x (to build the provider plugin)

## Build

Clone repository
```
cd $GOPATH/src
git clone github.com/tencentcloudstack/terraform-provider-tencentcloudenterprise.git
```

Download Dependencies
```sh
go mod tidy
```

Build executable binary
```sh
cd $GOPATH/src/github.com/tencentcloudstack/terraform-provider-tencentcloudenterprise

go build -o terraform-provider-cloud
```


## Configuration

### Environment Variables
```sh
export TENCENTCLOUD_DOMAIN=api3.your-domain.com
export TENCENTCLOUD_PROTOCOL=HTTP
export TENCENTCLOUD_REGION=your-region

# debug mode on
export TF_LOG=DEBUG
export TF_LOG_PATH=./terraform.log
# local provider mode on
export TF_CLI_CONFIG_FILE=/Users/developer/workspace/dev.tfrc
```

Configure your `dev.tfrc` file as follows:
```text
provider_installation {
 # Use /home/developer/tmp/terraform-null as an overridden package directory
 # for the hashicorp/null provider. This disables the version and checksum
 # verifications for this provider and forces Terraform to look for the
 # null provider plugin in the given directory.
 
 # Specify the provider project directory
 dev_overrides {
     "myorg/cloud" = "$GOPATH/src/github.com/tencentcloudstack/terraform-provider-tencentcloudenterprise"
 }
}
```

### Tenant Credential Configuration

```
export TENCENTCLOUD_SECRET_ID  = ID-EXAMPLE
export TENCENTCLOUD_SECRET_KEY = KEY-EXAMPLE
```

## Developer Guide

### DEBUG

You will need to set an environment variable named ``TF_LOG``, for more info please refer to [Terraform official doc](https://www.terraform.io/docs/internals/debugging.html):

```
export TF_LOG=DEBUG
```

### Avoid ``terraform init``

```
export TF_SKIP_PROVIDER_VERIFY=1
```

This will disable the verify steps, so after you update this provider, you won't need to create new resources, but use previously saved state.

### License

Terraform-Provider-TencentCloudEnterprise is under the Mozilla Public License 2.0. See the [LICENSE](LICENSE) file for details.