---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_tacacs Resource - terraform-provider-nxos"
subcategory: "Feature"
description: |-
  This resource can manage the TACACS+ feature configuration.
  API Documentation: fmTacacsplus https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Tacacsplus/
---

# nxos_feature_tacacs (Resource)

This resource can manage the TACACS+ feature configuration.

- API Documentation: [fmTacacsplus](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Tacacsplus/)

## Example Usage

```terraform
resource "nxos_feature_tacacs" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `admin_state` (String) Administrative state.
  - Choices: `enabled`, `disabled`

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_feature_tacacs.example "sys/fm/tacacsplus"
```
