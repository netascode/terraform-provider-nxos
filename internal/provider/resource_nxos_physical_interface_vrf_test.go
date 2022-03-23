// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosPhysicalInterfaceVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPhysicalInterfaceConfig_all(),
			},
			{
				Config: testAccNxosPhysicalInterfaceConfig_all() + testAccNxosPhysicalInterfaceVRFConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_physical_interface_vrf.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_physical_interface_vrf.test", "vrf_dn", "sys/inst-default"),
				),
			},
			{
				Config: testAccNxosPhysicalInterfaceConfig_all() + testAccNxosPhysicalInterfaceVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_physical_interface_vrf.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_physical_interface_vrf.test", "vrf_dn", "sys/inst-default"),
				),
			},
			{
				ResourceName:  "nxos_physical_interface_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/intf/phys-[eth1/10]/rtvrfMbr",
			},
		},
	})
}

func testAccNxosPhysicalInterfaceVRFConfig_minimum() string {
	return `
	resource "nxos_physical_interface_vrf" "test" {
		interface_id = "eth1/10"
		vrf_dn = "sys/inst-default"
	}
	`
}

func testAccNxosPhysicalInterfaceVRFConfig_all() string {
	return `
	resource "nxos_physical_interface_vrf" "test" {
		interface_id = "eth1/10"
		vrf_dn = "sys/inst-default"
	}
	`
}
