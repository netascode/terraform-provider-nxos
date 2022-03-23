// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosIPv4InterfaceAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosIPv4InterfaceConfig_all(),
			},
			{
				Config: testAccNxosIPv4InterfaceConfig_all() + testAccNxosIPv4InterfaceAddressConfig_all(),
			},
			{
				Config: testAccDataSourceNxosIPv4InterfaceAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_interface_address.test", "address", "24.63.46.49/30"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4InterfaceAddressConfig = `
data "nxos_ipv4_interface_address" "test" {
  vrf = "default"
  interface_id = "eth1/59"
  address = "24.63.46.49/30"
}
`
