// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureUDLD(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureUDLDConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_udld.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_udld.test",
				ImportState:   true,
				ImportStateId: "sys/fm/udld",
			},
		},
	})
}

func testAccNxosFeatureUDLDConfig_minimum() string {
	return `
	resource "nxos_feature_udld" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureUDLDConfig_all() string {
	return `
	resource "nxos_feature_udld" "test" {
		admin_state = "enabled"
	}
	`
}
