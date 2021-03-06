---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_bgp_peer Data Source - terraform-provider-nxos"
subcategory: "BGP"
description: |-
  This data source can read the BGP peer configuration.
  API Documentation: bgpPeer https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/
---

# nxos_bgp_peer (Data Source)

This data source can read the BGP peer configuration.

- API Documentation: [bgpPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/)

## Example Usage

```terraform
data "nxos_bgp_peer" "example" {
  asn     = "65001"
  vrf     = "default"
  address = "192.168.0.1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address` (String) Peer address.
- `asn` (String) Autonomous system number.
- `vrf` (String) VRF name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `description` (String) Peer description.
- `id` (String) The distinguished name of the object.
- `peer_template` (String) Peer template name.
- `peer_type` (String) Neighbor Fabric Type.
- `remote_asn` (String) Peer autonomous system number.
- `source_interface` (String) Source Interface. Must match first field in the output of `show intf brief`.


