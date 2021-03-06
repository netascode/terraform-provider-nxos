// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBridgeDomain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBridgeDomainConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bridge_domain.test", "fabric_encap", "vlan-10"),
					resource.TestCheckResourceAttr("nxos_bridge_domain.test", "access_encap", "unknown"),
					resource.TestCheckResourceAttr("nxos_bridge_domain.test", "name", "VLAN10"),
				),
			},
			{
				ResourceName:  "nxos_bridge_domain.test",
				ImportState:   true,
				ImportStateId: "sys/bd/bd-[vlan-10]",
			},
		},
	})
}

func testAccNxosBridgeDomainConfig_minimum() string {
	return `
	resource "nxos_bridge_domain" "test" {
		fabric_encap = "vlan-10"
	}
	`
}

func testAccNxosBridgeDomainConfig_all() string {
	return `
	resource "nxos_bridge_domain" "test" {
		fabric_encap = "vlan-10"
		access_encap = "unknown"
		name = "VLAN10"
	}
	`
}
