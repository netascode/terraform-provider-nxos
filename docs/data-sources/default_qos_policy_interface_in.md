---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_default_qos_policy_interface_in Data Source - terraform-provider-nxos"
subcategory: ""
description: |-
  This data source can read the default QoS policy interface in configuration.
  API Documentation: ipqosIf https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:If/
---

# nxos_default_qos_policy_interface_in (Data Source)

This data source can read the default QoS policy interface in configuration.

- API Documentation: [ipqosIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:If/)

## Example Usage

```terraform
data "nxos_default_qos_policy_interface_in" "example" {
  interface_id = "eth1/10"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **interface_id** (String) Must match first field in the output of `show intf brief`. Example: `eth1/1`.

### Read-Only

- **id** (String) The distinguished name of the object.

