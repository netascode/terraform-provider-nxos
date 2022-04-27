// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosHMMInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosHMMInterfacePrerequisitesConfig + testAccDataSourceNxosHMMInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_hmm_interface.test", "interface_id", "vlan10"),
					resource.TestCheckResourceAttr("data.nxos_hmm_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_hmm_interface.test", "mode", "anycastGW"),
				),
			},
		},
	})
}

const testAccDataSourceNxosHMMInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/hmm"
  class_name = "fmHmm"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/hmm"
  class_name = "hmmEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/hmm/fwdinst"
  class_name = "hmmFwdInst"
  content = {
      adminSt = "enabled"
      amac = "20:20:00:00:10:10"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosHMMInterfaceConfig = `

resource "nxos_hmm_interface" "test" {
  interface_id = "vlan10"
  admin_state = "enabled"
  mode = "anycastGW"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_hmm_interface" "test" {
  interface_id = "vlan10"
  depends_on = [nxos_hmm_interface.test]
}
`
