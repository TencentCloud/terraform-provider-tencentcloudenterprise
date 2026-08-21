# TKE Kubernetes User Permissions Local Validation

This example validates `tencentcloudenterprise_tke_kubernetes_user_permissions` with a provider binary built from the current source checkout. Terraform loads the binary through a development override, so no provider release is required.

## Safety

Use an existing test cluster and a dedicated test user. The resource treats `permissions` as the complete desired permission set for `target_uin`.

- Applying an update can remove permissions omitted from the configuration.
- Destroying the resource removes the managed cluster and namespace bindings.
- Do not use a production account or an account with permissions that must be preserved.
- Keep credentials, Terraform state, plans, and logs outside the repository.

## Prerequisites

- Go 1.23.x
- Terraform CLI 1.5 or later
- Credentials that can manage RBAC bindings in the target TKE cluster
- The API domain and region of the target TCE environment

## 1. Build the provider

Run these commands from the repository root:

```bash
PROVIDER_BIN_DIR="${TMPDIR:-/tmp}/tencentcloudenterprise-provider-bin"
mkdir -p "${PROVIDER_BIN_DIR}"
go build -o "${PROVIDER_BIN_DIR}/terraform-provider-tencentcloudenterprise" .
```

Keep `PROVIDER_BIN_DIR` exported in the shell used for the remaining steps:

```bash
export PROVIDER_BIN_DIR
```

## 2. Create an isolated validation directory

Run these commands from the repository root:

```bash
VALIDATION_DIR="${TMPDIR:-/tmp}/tke-kubernetes-user-permissions-validation"
mkdir -p "${VALIDATION_DIR}"
cp examples/tencentcloud-tke-kubernetes-user-permissions/* "${VALIDATION_DIR}/"
cd "${VALIDATION_DIR}"
```

Create the active Terraform CLI configuration from the template:

```bash
sed "s|<PROVIDER_BIN_DIR>|${PROVIDER_BIN_DIR}|g" dev.tfrc.example > dev.tfrc
export TF_CLI_CONFIG_FILE="${VALIDATION_DIR}/dev.tfrc"
```

The override value must be the directory containing the binary, not the binary path.

## 3. Configure access

Set credentials and endpoint information in environment variables:

```bash
export TENCENTCLOUD_SECRET_ID="<SECRET_ID>"
export TENCENTCLOUD_SECRET_KEY="<SECRET_KEY>"
export TENCENTCLOUD_REGION="<REGION>"
export TENCENTCLOUD_DOMAIN="<API_DOMAIN>"
export TENCENTCLOUD_PROTOCOL="HTTP"
```

Use `HTTPS` when required by the environment. `TENCENTCLOUD_DOMAIN` must contain only the API domain, without a URL path.

Do not place credentials in `terraform.tfvars`, `dev.tfrc`, or tracked files.

## 4. Configure the test resource

Copy the variable template:

```bash
cp terraform.tfvars.example terraform.tfvars
```

Replace the placeholder values with an existing test cluster ID and a dedicated test user UIN:

```hcl
cluster_id = "cls-xxxxxxxx"
target_uin = "100000000000"

permissions = [
  {
    role_name = "tke:ro"
    role_type = "cluster"
  }
]
```

## 5. Initialize and confirm the local override

```bash
terraform init
terraform validate
```

Terraform must print a warning that provider development overrides are in effect for:

```text
TencentCloud/tencentcloudenterprise
```

Stop if the warning is absent. Check `TF_CLI_CONFIG_FILE`, the absolute provider directory in `dev.tfrc`, and the binary name.

## 6. Validate create and read

```bash
terraform plan -out=create.tfplan
terraform apply create.tfplan
terraform state show tencentcloudenterprise_tke_kubernetes_user_permissions.validation
terraform output
```

Expected results:

- The resource ID is `cluster_id:target_uin`.
- The cluster-level `tke:ro` permission is present.
- The target TKE cluster contains the corresponding ClusterRoleBinding.

## 7. Validate update

Add a namespace permission to `terraform.tfvars`:

```hcl
permissions = [
  {
    role_name = "tke:ro"
    role_type = "cluster"
  },
  {
    role_name = "tke:dev"
    role_type = "namespace"
    namespace = "default"
  }
]
```

Apply the update:

```bash
terraform plan -out=update.tfplan
terraform apply update.tfplan
terraform state show tencentcloudenterprise_tke_kubernetes_user_permissions.validation
```

Expected results:

- Both permissions appear in state.
- The cluster-level binding remains present.
- A RoleBinding is present in the `default` namespace.

## 8. Validate idempotency

Without changing the configuration, run:

```bash
terraform plan
```

The expected result is `No changes`.

## 9. Validate import

Set the import values:

```bash
export TKE_CLUSTER_ID="<CLUSTER_ID>"
export TKE_TARGET_UIN="<TARGET_UIN>"
```

Back up the local state, remove only the local state entry, and import the existing bindings:

```bash
terraform state pull > pre-import.tfstate
terraform state rm tencentcloudenterprise_tke_kubernetes_user_permissions.validation
terraform import \
  tencentcloudenterprise_tke_kubernetes_user_permissions.validation \
  "${TKE_CLUSTER_ID}:${TKE_TARGET_UIN}"
terraform plan
```

`terraform state rm` does not delete remote RBAC bindings. The final plan should report no changes.

## 10. Clean up

Review the target user before cleanup. Destroy removes the permissions managed by this resource:

```bash
terraform plan -destroy -out=destroy.tfplan
terraform apply destroy.tfplan
```

Confirm in the TKE cluster that the test user's ClusterRoleBinding and RoleBinding were removed.

## Troubleshooting

Enable Terraform logs only in the isolated validation directory:

```bash
export TF_LOG=DEBUG
export TF_LOG_PATH="${VALIDATION_DIR}/terraform-debug.log"
```

Common checks:

- Rebuild the binary after every source change.
- Confirm `TF_CLI_CONFIG_FILE` points to the generated `dev.tfrc`.
- Confirm the override directory contains `terraform-provider-tencentcloudenterprise`.
- Confirm the endpoint protocol, domain, region, credentials, and TKE RBAC privileges.
- Remove or securely archive state, plans, and debug logs when validation is complete.
