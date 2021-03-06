// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosIPv4AccessListPolicyIngressInterfaceInstace(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosIPv4AccessListPolicyIngressInterfaceInstacePrerequisitesConfig + testAccDataSourceNxosIPv4AccessListPolicyIngressInterfaceInstaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_access_list_policy_ingress_interface_instace.test", "name", "ACL1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4AccessListPolicyIngressInterfaceInstacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/acl/ipv4/policy/ingress/intf-[eth1/10]"
  class_name = "aclIf"
  content = {
      name = "eth1/10"
  }
}

`

const testAccDataSourceNxosIPv4AccessListPolicyIngressInterfaceInstaceConfig = `

resource "nxos_ipv4_access_list_policy_ingress_interface_instace" "test" {
  interface_id = "eth1/10"
  name = "ACL1"
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_ipv4_access_list_policy_ingress_interface_instace" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_ipv4_access_list_policy_ingress_interface_instace.test]
}
`
