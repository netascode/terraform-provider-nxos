// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosEVPNVNIRouteTargetDirection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosEVPNVNIRouteTargetDirectionPrerequisitesConfig + testAccDataSourceNxosEVPNVNIRouteTargetDirectionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_evpn_vni_route_target_direction.test", "direction", "import"),
				),
			},
		},
	})
}

const testAccDataSourceNxosEVPNVNIRouteTargetDirectionPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/nvo"
  class_name = "fmNvo"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  delete = false
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/evpn"
  class_name = "rtctrlL2Evpn"
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/evpn/bdevi-[vxlan-123456]"
  class_name = "rtctrlBDEvi"
  content = {
      encap = "vxlan-123456"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosEVPNVNIRouteTargetDirectionConfig = `

resource "nxos_evpn_vni_route_target_direction" "test" {
  encap = "vxlan-123456"
  direction = "import"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_evpn_vni_route_target_direction" "test" {
  encap = "vxlan-123456"
  direction = "import"
  depends_on = [nxos_evpn_vni_route_target_direction.test]
}
`
