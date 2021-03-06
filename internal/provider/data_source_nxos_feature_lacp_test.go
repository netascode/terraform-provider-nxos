// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosFeatureLACP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosFeatureLACPConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_feature_lacp.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosFeatureLACPConfig = `

resource "nxos_feature_lacp" "test" {
  admin_state = "enabled"
}

data "nxos_feature_lacp" "test" {
  depends_on = [nxos_feature_lacp.test]
}
`
