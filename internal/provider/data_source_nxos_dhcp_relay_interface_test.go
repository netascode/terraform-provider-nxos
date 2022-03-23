// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDHCPRelayInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureDHCPConfig_all(),
			},
			{
				Config: testAccNxosFeatureDHCPConfig_all() + testAccNxosDHCPRelayInterfaceConfig_all(),
			},
			{
				Config: testAccDataSourceNxosDHCPRelayInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_interface.test", "interface_id", "eth1/10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDHCPRelayInterfaceConfig = `
data "nxos_dhcp_relay_interface" "test" {
  interface_id = "eth1/10"
}
`
