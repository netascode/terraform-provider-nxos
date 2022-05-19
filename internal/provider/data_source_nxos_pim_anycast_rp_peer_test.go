// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMAnycastRPPeer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPIMAnycastRPPeerPrerequisitesConfig + testAccDataSourceNxosPIMAnycastRPPeerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_anycast_rp_peer.test", "address", "10.1.1.1/32"),
					resource.TestCheckResourceAttr("data.nxos_pim_anycast_rp_peer.test", "rp_set_address", "20.1.1.1/32"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMAnycastRPPeerPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  content = {
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/pim/inst/dom-[default]/acastrpfunc"
  class_name = "pimAcastRPFuncP"
  content = {
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

const testAccDataSourceNxosPIMAnycastRPPeerConfig = `

resource "nxos_pim_anycast_rp_peer" "test" {
  vrf_name = "default"
  address = "10.1.1.1/32"
  rp_set_address = "20.1.1.1/32"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
}

data "nxos_pim_anycast_rp_peer" "test" {
  vrf_name = "default"
  address = "10.1.1.1/32"
  rp_set_address = "20.1.1.1/32"
  depends_on = [nxos_pim_anycast_rp_peer.test]
}
`