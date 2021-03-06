// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosFeatureLLDP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosFeatureLLDPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_feature_lldp.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosFeatureLLDPConfig = `

resource "nxos_feature_lldp" "test" {
  admin_state = "enabled"
}

data "nxos_feature_lldp" "test" {
  depends_on = [nxos_feature_lldp.test]
}
`
