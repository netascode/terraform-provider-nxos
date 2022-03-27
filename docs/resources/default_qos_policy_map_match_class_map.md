---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_default_qos_policy_map_match_class_map Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the default QoS policy map match class map configuration.
  API Documentation: ipqosMatchCMap https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:MatchCMap/
  Parent resources
  nxosdefaultqospolicymap https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map
  Child resources
  nxosdefaultqospolicymapmatchclassmapsetqosgroup https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map_set_qos_groupnxosdefaultqospolicymapmatchclassmappolice https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map_police
  Referenced resources
  nxosdefaultqosclassmap https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_class_map
---

# nxos_default_qos_policy_map_match_class_map (Resource)

This resource can manage the default QoS policy map match class map configuration.

- API Documentation: [ipqosMatchCMap](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:MatchCMap/)

### Parent resources

- [nxos_default_qos_policy_map](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map)

### Child resources

- [nxos_default_qos_policy_map_match_class_map_set_qos_group](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map_set_qos_group)
- [nxos_default_qos_policy_map_match_class_map_police](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_policy_map_match_class_map_police)

### Referenced resources

- [nxos_default_qos_class_map](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/default_qos_class_map)

## Example Usage

```terraform
resource "nxos_default_qos_policy_map_match_class_map" "example" {
  policy_map_name = "PM1"
  name            = "Voice"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) Class map name.
- **policy_map_name** (String) Policy map name.

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_default_qos_policy_map_match_class_map.example "sys/ipqos/dflt/p/name-[PM1]/cmap-[Voice]"
```