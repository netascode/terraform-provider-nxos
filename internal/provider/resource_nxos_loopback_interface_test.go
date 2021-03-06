// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosLoopbackInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosLoopbackInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_loopback_interface.test", "interface_id", "lo123"),
					resource.TestCheckResourceAttr("nxos_loopback_interface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("nxos_loopback_interface.test", "description", "My Description"),
				),
			},
			{
				ResourceName:  "nxos_loopback_interface.test",
				ImportState:   true,
				ImportStateId: "sys/intf/lb-[lo123]",
			},
		},
	})
}

func testAccNxosLoopbackInterfaceConfig_minimum() string {
	return `
	resource "nxos_loopback_interface" "test" {
		interface_id = "lo123"
	}
	`
}

func testAccNxosLoopbackInterfaceConfig_all() string {
	return `
	resource "nxos_loopback_interface" "test" {
		interface_id = "lo123"
		admin_state = "down"
		description = "My Description"
	}
	`
}
