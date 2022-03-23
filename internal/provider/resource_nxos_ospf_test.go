// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosOSPF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccNxosOSPFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ospf.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_ospf.test",
				ImportState:   true,
				ImportStateId: "sys/ospf",
			},
		},
	})
}

func testAccNxosOSPFConfig_minimum() string {
	return `
	resource "nxos_ospf" "test" {
	}
	`
}

func testAccNxosOSPFConfig_all() string {
	return `
	resource "nxos_ospf" "test" {
		admin_state = "enabled"
	}
	`
}
