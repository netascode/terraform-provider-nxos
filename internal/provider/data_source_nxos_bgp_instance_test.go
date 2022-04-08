// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosBGPInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPInstancePrerequisitesConfig + testAccDataSourceNxosBGPInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_instance.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_bgp_instance.test", "asn", "65001"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
}

`

const testAccDataSourceNxosBGPInstanceConfig = `

resource "nxos_bgp_instance" "test" {
  admin_state = "enabled"
  asn = "65001"
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_bgp_instance" "test" {
  depends_on = [nxos_bgp_instance.test]
}
`