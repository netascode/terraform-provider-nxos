---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim_ssm_policy Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the PIM SSM policy configuration.
  API Documentation: pimSSMPatP https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMPatP/
---

# nxos_pim_ssm_policy (Resource)

This resource can manage the PIM SSM policy configuration.

- API Documentation: [pimSSMPatP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMPatP/)

## Example Usage

```terraform
resource "nxos_pim_ssm_policy" "example" {
  vrf_name = "default"
  name     = "SSM"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **vrf_name** (String) VRF name.

### Optional

- **name** (String) Policy name.

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_pim_ssm_policy.example "sys/pim/inst/dom-[default]/ssm"
```