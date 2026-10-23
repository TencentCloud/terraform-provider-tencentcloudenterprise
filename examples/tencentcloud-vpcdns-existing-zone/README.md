# Attach a new VPC to an existing private DNS zone

This example creates a VPC and associates it with a private DNS zone that already exists. The zone does not need to be recreated, declared as a resource, or imported into this Terraform state.

Use a provider version that includes `tencentcloudenterprise_vpcdns_zone_vpc_attachment`. Configure credentials and the API endpoint through the provider environment variables, and supply `existing_zone_id`, `region`, `vpc_name`, and `vpc_cidr` for your account.

```shell
terraform init
terraform plan -out=tfplan
terraform apply tfplan
```

The attachment preserves other VPC associations. Removing the attachment removes only its association; the existing zone remains. Destroying this entire example also deletes the VPC created by it.

If another Terraform configuration manages the zone's `vpc_set`, ensure it does not overwrite independently managed attachments. Keep one owner for each association; when delegating associations to attachment resources, exclude `vpc_set` from updates in the zone configuration (for example, with `lifecycle { ignore_changes = [vpc_set] }`).

To adopt an association that already exists, import it instead of creating it again:

```shell
terraform import tencentcloudenterprise_vpcdns_zone_vpc_attachment.existing_zone 'zone-id#vpc-id'
```

Replace both IDs with the real IDs and ensure the configured VPC matches the imported association. If that VPC also already exists, import it into `tencentcloudenterprise_vpc.new` before planning this example.
