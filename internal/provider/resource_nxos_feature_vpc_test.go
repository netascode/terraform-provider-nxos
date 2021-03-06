// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureVPC(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosFeatureVPCConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_vpc.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_vpc.test",
				ImportState:   true,
				ImportStateId: "sys/fm/vpc",
			},
		},
	})
}

func testAccNxosFeatureVPCConfig_minimum() string {
	return `
	resource "nxos_feature_vpc" "test" {
		admin_state = "enabled"
	}
	`
}

func testAccNxosFeatureVPCConfig_all() string {
	return `
	resource "nxos_feature_vpc" "test" {
		admin_state = "enabled"
	}
	`
}
