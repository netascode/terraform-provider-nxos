// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosBGPVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPVRFPrerequisitesConfig + testAccDataSourceNxosBGPVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("data.nxos_bgp_vrf.test", "router_id", "1.1.1.1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPVRFPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosBGPVRFConfig = `

resource "nxos_bgp_vrf" "test" {
  asn = "65001"
  name = "default"
  router_id = "1.1.1.1"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_bgp_vrf" "test" {
  asn = "65001"
  name = "default"
  depends_on = [nxos_bgp_vrf.test]
}
`
