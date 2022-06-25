---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_ipv4_access_list_entry Data Source - terraform-provider-nxos"
subcategory: "IPv4"
description: |-
  This data source can read IPv4 Access List Entries.
  API Documentation: ipv4aclACE https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/ipv4acl:ACE/
---

# nxos_ipv4_access_list_entry (Data Source)

This data source can read IPv4 Access List Entries.

- API Documentation: [ipv4aclACE](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/ipv4acl:ACE/)

## Example Usage

```terraform
data "nxos_ipv4_access_list_entry" "example" {
  name            = "ACL1"
  sequence_number = 10
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Access list name.
- `sequence_number` (Number) Sequence number.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `ack` (Boolean) Match TCP ACK flag.
- `action` (String) Action.
- `destination_address_group` (String) Destination address group.
- `destination_port_1` (Number) First destination port number.
- `destination_port_2` (Number) Second destination port number.
- `destination_port_group` (String) Destination port group.
- `destination_port_mask` (Number) Destination port mask number.
- `destination_port_operator` (String) Destination port operator.
- `destination_prefix` (String) Destination prefix.
- `destination_prefix_length` (String) Destination prefix length.
- `destination_prefix_mask` (String) Destination prefix mask.
- `dscp` (Number) Match DSCP.
- `est` (Boolean) Match TCP EST flag.
- `fin` (Boolean) Match TCP FIN flag.
- `fragment` (Boolean) Match non-initial fragment.
- `http_option_type` (String) HTTP option method.
- `icmp_code` (Number) ICMP code.
- `icmp_type` (Number) ICMP type.
- `id` (String) The distinguished name of the object.
- `logging` (Boolean) Log matches against ACL entry.
- `packet_length_1` (String) First packet length. Either `invalid` or a number between 19 and 9210.
- `packet_length_2` (String) Second packet length. Either `invalid` or a number between 19 and 9210.
- `packet_length_operator` (String) Packet length operator.
- `precedence` (Number) Precedence.
- `protocol` (String) Protocol name or number.
- `protocol_mask` (String) Protocol mask name or number.
- `psh` (Boolean) Match TCP PSH flag.
- `redirect` (String) Redirect action.
- `remark` (String) ACL comment.
- `rev` (Boolean) Match TCP REV flag.
- `rst` (Boolean) Match TCP RST flag.
- `source_address_group` (String) Source address group.
- `source_port_1` (Number) First source port number.
- `source_port_2` (Number) Second source port number.
- `source_port_group` (String) Source port group.
- `source_port_mask` (Number) Source port mask number.
- `source_port_operator` (String) Source port operator.
- `source_prefix` (String) Source prefix.
- `source_prefix_length` (String) Source prefix length.
- `source_prefix_mask` (String) Source prefix mask.
- `syn` (Boolean) Match TCP SYN flag.
- `time_range` (String) Time range name.
- `ttl` (Number) TTL.
- `urg` (Boolean) Match TCP URG flag.
- `vlan` (Number) VLAN ID.
- `vni` (String) NVE VNI ID. Either `invalid` or a number between 0 and 16777216.

