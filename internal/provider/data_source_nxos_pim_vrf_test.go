// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMVRF(t *testing.T) {
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
				Config: testAccDataSourceNxosPIMVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("data.nxos_pim_vrf.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_pim_vrf.test", "bfd", "true"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMVRFConfig = `
data "nxos_pim_vrf" "test" {
  name = "default"
}
`
