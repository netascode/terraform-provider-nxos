// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosFeatureInterfaceVLAN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosFeatureInterfaceVLANConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_feature_interface_vlan.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosFeatureInterfaceVLANConfig = `

resource "nxos_feature_interface_vlan" "test" {
  admin_state = "enabled"
}

data "nxos_feature_interface_vlan" "test" {
  depends_on = [nxos_feature_interface_vlan.test]
}
`
