// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosVRFAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosVRFAddressFamilyPrerequisitesConfig + testAccDataSourceNxosVRFAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_vrf_address_family.test", "address_family", "ipv4-ucast"),
				),
			},
		},
	})
}

const testAccDataSourceNxosVRFAddressFamilyPrerequisitesConfig = `
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

resource "nxos_rest" "PreReq2" {
  dn = "sys/inst-[VRF1]/dom-[VRF1]"
  class_name = "rtctrlDom"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosVRFAddressFamilyConfig = `

resource "nxos_vrf_address_family" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_vrf_address_family" "test" {
  vrf = "VRF1"
  address_family = "ipv4-ucast"
  depends_on = [nxos_vrf_address_family.test]
}
`
