// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosIPv4Interface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv4InterfacePrerequisitesConfig + testAccNxosIPv4InterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "drop_glean", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "forward", "disabled"),
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "unnumbered", "unspecified"),
					resource.TestCheckResourceAttr("nxos_ipv4_interface.test", "urpf", "disabled"),
				),
			},
			{
				ResourceName:  "nxos_ipv4_interface.test",
				ImportState:   true,
				ImportStateId: "sys/ipv4/inst/dom-[default]/if-[eth1/10]",
			},
		},
	})
}

const testAccNxosIPv4InterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ipv4/inst/dom-[default]"
  class_name = "ipv4Dom"
  content = {
  }
}

`

func testAccNxosIPv4InterfaceConfig_minimum() string {
	return `
	resource "nxos_ipv4_interface" "test" {
		vrf = "default"
		interface_id = "eth1/10"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosIPv4InterfaceConfig_all() string {
	return `
	resource "nxos_ipv4_interface" "test" {
		vrf = "default"
		interface_id = "eth1/10"
		drop_glean = "disabled"
		forward = "disabled"
		unnumbered = "unspecified"
		urpf = "disabled"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}
