// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosPIMSSMPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMSSMPolicyPrerequisitesConfig + testAccNxosPIMSSMPolicyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_ssm_policy.test", "vrf_name", "default"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_policy.test", "name", "SSM"),
				),
			},
			{
				ResourceName:  "nxos_pim_ssm_policy.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst/dom-[default]/ssm",
			},
		},
	})
}

const testAccNxosPIMSSMPolicyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  content = {
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosPIMSSMPolicyConfig_minimum() string {
	return `
	resource "nxos_pim_ssm_policy" "test" {
		vrf_name = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosPIMSSMPolicyConfig_all() string {
	return `
	resource "nxos_pim_ssm_policy" "test" {
		vrf_name = "default"
		name = "SSM"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
