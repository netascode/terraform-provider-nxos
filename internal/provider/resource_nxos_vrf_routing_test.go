// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosVRFRouting(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosVRFRoutingPrerequisitesConfig + testAccNxosVRFRoutingConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_vrf_routing.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("nxos_vrf_routing.test", "route_distinguisher", "rd:unknown:0:0"),
				),
			},
			{
				ResourceName:  "nxos_vrf_routing.test",
				ImportState:   true,
				ImportStateId: "sys/inst-[VRF1]/dom-[VRF1]",
			},
		},
	})
}

const testAccNxosVRFRoutingPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/inst-[VRF1]"
  class_name = "l3Inst"
  content = {
      name = "VRF1"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ipv4/inst/dom-[VRF1]"
  class_name = "ipv4Dom"
  content = {
      name = "VRF1"
  }
}

`

func testAccNxosVRFRoutingConfig_minimum() string {
	return `
	resource "nxos_vrf_routing" "test" {
		vrf = "VRF1"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosVRFRoutingConfig_all() string {
	return `
	resource "nxos_vrf_routing" "test" {
		vrf = "VRF1"
		route_distinguisher = "rd:unknown:0:0"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
