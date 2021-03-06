// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPrerequisitesConfig + testAccNxosBGPConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_bgp.test",
				ImportState:   true,
				ImportStateId: "sys/bgp",
			},
		},
	})
}

const testAccNxosBGPPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bgp"
  class_name = "fmBgp"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

`

func testAccNxosBGPConfig_minimum() string {
	return `
	resource "nxos_bgp" "test" {
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}

func testAccNxosBGPConfig_all() string {
	return `
	resource "nxos_bgp" "test" {
		admin_state = "enabled"
  		depends_on = [nxos_rest.PreReq0, ]
	}
	`
}
