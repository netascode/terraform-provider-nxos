// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosQueuingQOSPolicyMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosQueuingQOSPolicyMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map.test", "name", "PM1"),
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map.test", "match_type", "match-any"),
				),
			},
			{
				ResourceName:  "nxos_queuing_qos_policy_map.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/queuing/p/name-[PM1]",
			},
		},
	})
}

func testAccNxosQueuingQOSPolicyMapConfig_minimum() string {
	return `
	resource "nxos_queuing_qos_policy_map" "test" {
		name = "PM1"
	}
	`
}

func testAccNxosQueuingQOSPolicyMapConfig_all() string {
	return `
	resource "nxos_queuing_qos_policy_map" "test" {
		name = "PM1"
		match_type = "match-any"
	}
	`
}
