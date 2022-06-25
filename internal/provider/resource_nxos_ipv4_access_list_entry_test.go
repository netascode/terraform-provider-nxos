// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosIPv4AccessListEntry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv4AccessListEntryPrerequisitesConfig + testAccNxosIPv4AccessListEntryConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "name", "ACL1"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "sequence_number", "10"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "ack", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "dscp", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "destination_port_1", "443"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "destination_port_2", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "destination_port_operator", "eq"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "destination_prefix", "10.1.1.0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "destination_prefix_length", "24"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "est", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "fin", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "fragment", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "http_option_type", "invalid"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "icmp_code", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "icmp_type", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "logging", "true"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "psh", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "rev", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "rst", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "source_port_1", "443"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "source_port_2", "0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "source_port_operator", "eq"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "source_prefix", "20.1.0.0"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "source_prefix_length", "16"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "syn", "false"),
					resource.TestCheckResourceAttr("nxos_ipv4_access_list_entry.test", "urg", "false"),
				),
			},
			{
				ResourceName:  "nxos_ipv4_access_list_entry.test",
				ImportState:   true,
				ImportStateId: "sys/acl/ipv4/name-[ACL1]//seq-[10]",
			},
		},
	})
}

const testAccNxosIPv4AccessListEntryPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/acl/ipv4/name-[ACL1]"
  class_name = "ipv4aclACL"
  content = {
      name = "ACL1"
  }
}

`

func testAccNxosIPv4AccessListEntryConfig_minimum() string {
	return `
	resource "nxos_ipv4_access_list_entry" "test" {
		name = "ACL1"
		sequence_number = 10
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosIPv4AccessListEntryConfig_all() string {
	return `
	resource "nxos_ipv4_access_list_entry" "test" {
		name = "ACL1"
		sequence_number = 10
		ack = false
		action = "permit"
		dscp = 0
		destination_port_1 = 443
		destination_port_2 = 0
		destination_port_operator = "eq"
		destination_prefix = "10.1.1.0"
		destination_prefix_length = "24"
		est = false
		fin = false
		fragment = false
		http_option_type = "invalid"
		icmp_code = 0
		icmp_type = 0
		logging = true
		protocol = "tcp"
		psh = false
		rev = false
		rst = false
		source_port_1 = 443
		source_port_2 = 0
		source_port_operator = "eq"
		source_prefix = "20.1.0.0"
		source_prefix_length = "16"
		syn = false
		urg = false
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}