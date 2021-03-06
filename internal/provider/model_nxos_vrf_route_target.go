// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFRouteTarget struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Vrf     types.String `tfsdk:"vrf"`
	Af_type types.String `tfsdk:"address_family"`
	Rt_type types.String `tfsdk:"route_target_address_family"`
	Type    types.String `tfsdk:"direction"`
	Rtt     types.String `tfsdk:"route_target"`
}

func (data VRFRouteTarget) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]/rttp-[%s]/ent-[%s]", data.Vrf.Value, data.Af_type.Value, data.Rt_type.Value, data.Type.Value, data.Rtt.Value)
}

func (data VRFRouteTarget) getClassName() string {
	return "rtctrlRttEntry"
}

func (data VRFRouteTarget) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("rtt", data.Rtt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFRouteTarget) fromBody(res gjson.Result) {
	data.Rtt.Value = res.Get("*.attributes.rtt").String()
}

func (data *VRFRouteTarget) fromPlan(plan VRFRouteTarget) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Vrf.Value = plan.Vrf.Value
	data.Af_type.Value = plan.Af_type.Value
	data.Rt_type.Value = plan.Rt_type.Value
	data.Type.Value = plan.Type.Value
}
