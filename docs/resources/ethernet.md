---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_ethernet Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage global ethernet settings.
  API Documentation: ethpmInst https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/ethpm:Inst/
---

# nxos_ethernet (Resource)

This resource can manage global ethernet settings.

- API Documentation: [ethpmInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/ethpm:Inst/)

## Example Usage

```terraform
resource "nxos_ethernet" "example" {
  mtu = 9216
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **mtu** (Number) System jumbo MTU.
  - Range: `576`-`9216`
  - Default value: `9216`

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_ethernet.example "sys/ethpm/inst"
```