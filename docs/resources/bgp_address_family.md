---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_bgp_address_family Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the BGP (VRF) address family configuration.
  API Documentation: bgpGr https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:DomAf/
  Parent resources
  nxosbgpvrf https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/bgp_vrf
---

# nxos_bgp_address_family (Resource)

This resource can manage the BGP (VRF) address family configuration.

- API Documentation: [bgpGr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:DomAf/)

### Parent resources

- [nxos_bgp_vrf](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/bgp_vrf)

## Example Usage

```terraform
resource "nxos_bgp_address_family" "example" {
  vrf                          = "default"
  address_family               = "ipv4-ucast"
  critical_nexthop_timeout     = 1800
  non_critical_nexthop_timeout = 1800
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `address_family` (String) Address Family.
  - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`, `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`, `ipv4-mdt`
  - Default value: `ipv4-ucast`
- `vrf` (String) VRF name.

### Optional

- `critical_nexthop_timeout` (Number) The next-hop address tracking delay timer for critical next-hop reachability routes.
  - Range: `1`-`4294967295`
  - Default value: `3000`
- `device` (String) A device name from the provider configuration.
- `non_critical_nexthop_timeout` (Number) The next-hop address tracking delay timer for non-critical next-hop reachability routes.
  - Range: `1`-`4294967295`
  - Default value: `10000`

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_bgp_address_family.example "sys/bgp/inst/dom-[default]/af-[ipv4-ucast]"
```