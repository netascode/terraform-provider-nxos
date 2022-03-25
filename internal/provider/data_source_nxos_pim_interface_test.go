// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all() + testAccNxosPIMVRFConfig_all(),
			},
			{
				Config: testAccNxosPIMConfig_all() + testAccNxosPIMInstanceConfig_all() + testAccNxosPIMVRFConfig_all() + testAccNxosPIMInterfaceConfig_all(),
			},
			{
				Config: testAccDataSourceNxosPIMInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "name", "default"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "bfd", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "dr_priority", "10"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "passive", "false"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "sparse_mode", "true"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMInterfaceConfig = `
data "nxos_pim_interface" "test" {
  name = "default"
  interface_id = "eth1/10"
}
`
