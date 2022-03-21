// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureInterfaceVLAN(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:testAccNxosFeatureInterfaceVLANConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
				),
			},
			{
				Config:testAccNxosFeatureInterfaceVLANConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_interface_vlan.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_interface_vlan.test",
				ImportState:   true,
				ImportStateId: "sys/fm/ifvlan",
			},
		},
	})
}

func testAccNxosFeatureInterfaceVLANConfig_minimum() string {
	return `
	resource "nxos_feature_interface_vlan" "test" {
	}
	`
}

func testAccNxosFeatureInterfaceVLANConfig_all() string {
	return `
	resource "nxos_feature_interface_vlan" "test" {
		admin_state = "enabled"
	}
	`
}