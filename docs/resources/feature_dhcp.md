---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_dhcp Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the DHCP feature configuration.
  API Documentation: fmDhcp https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Dhcp/
---

# nxos_feature_dhcp (Resource)

This resource can manage the DHCP feature configuration.

- API Documentation: [fmDhcp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Dhcp/)

## Example Usage

```terraform
resource "nxos_feature_dhcp" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **admin_state** (String) Administrative state.
  - Default value: `disabled`

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_feature_dhcp.example "sys/fm/dhcp"
```