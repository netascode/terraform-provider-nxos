// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type QueuingQOSPolicyMapMatchClassMap struct {
	Dn              types.String `tfsdk:"id"`
	Policy_map_name types.String `tfsdk:"policy_map_name"`
	Name            types.String `tfsdk:"name"`
}

func (data QueuingQOSPolicyMapMatchClassMap) getDn() string {
	return fmt.Sprintf("sys/ipqos/queuing/p/name-[%s]/cmap-[%s]", data.Policy_map_name.Value, data.Name.Value)
}

func (data QueuingQOSPolicyMapMatchClassMap) getClassName() string {
	return "ipqosMatchCMap"
}

func (data QueuingQOSPolicyMapMatchClassMap) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *QueuingQOSPolicyMapMatchClassMap) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Name.Value = res.Get("*.attributes.name").String()
}

func (data *QueuingQOSPolicyMapMatchClassMap) fromPlan(plan QueuingQOSPolicyMapMatchClassMap) {
	data.Policy_map_name.Value = plan.Policy_map_name.Value
}