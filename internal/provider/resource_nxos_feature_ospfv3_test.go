// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureOSPFv3(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureOSPFv3Config_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_ospfv3.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_ospfv3.test",
				ImportState:   true,
				ImportStateId: "sys/fm/ospfv3",
			},
		},
	})
}

func testAccNxosFeatureOSPFv3Config_minimum() string {
	return `
	resource "nxos_feature_ospfv3" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureOSPFv3Config_all() string {
	return `
	resource "nxos_feature_ospfv3" "test" {
		admin_state = "enabled"
	}
	`
}