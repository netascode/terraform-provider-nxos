// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPPeerTemplateMaxPrefix(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPeerTemplateMaxPrefixPrerequisitesConfig + testAccNxosBGPPeerTemplateMaxPrefixConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "template_name", "SPINE-PEERS"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "action", "log"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "maximum_prefix", "10000"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "restart_time", "0"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_template_max_prefix.test", "threshold", "30"),
				),
			},
			{
				ResourceName:  "nxos_bgp_peer_template_max_prefix.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]/af-[ipv4-ucast]/maxpfxp",
			},
		},
	})
}

const testAccNxosBGPPeerTemplateMaxPrefixPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]"
  class_name = "bgpPeerCont"
  content = {
      name = "SPINE-PEERS"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]/af-[ipv4-ucast]"
  class_name = "bgpPeerAf"
  content = {
      type = "ipv4-ucast"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

func testAccNxosBGPPeerTemplateMaxPrefixConfig_minimum() string {
	return `
	resource "nxos_bgp_peer_template_max_prefix" "test" {
		vrf = "default"
		template_name = "SPINE-PEERS"
		address_family = "ipv4-ucast"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}

func testAccNxosBGPPeerTemplateMaxPrefixConfig_all() string {
	return `
	resource "nxos_bgp_peer_template_max_prefix" "test" {
		vrf = "default"
		template_name = "SPINE-PEERS"
		address_family = "ipv4-ucast"
		action = "log"
		maximum_prefix = 10000
		restart_time = 0
		threshold = 30
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}