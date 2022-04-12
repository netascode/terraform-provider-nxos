// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRF struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Descr  types.String `tfsdk:"description"`
	Encap  types.String `tfsdk:"encap"`
}

func (data VRF) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]", data.Name.Value)
}

func (data VRF) getClassName() string {
	return "l3Inst"
}

func (data VRF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.Value).
		Set("descr", data.Descr.Value).
		Set("encap", data.Encap.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRF) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Name.Value = res.Get("*.attributes.name").String()
	data.Descr.Value = res.Get("*.attributes.descr").String()
	data.Encap.Value = res.Get("*.attributes.encap").String()
}

func (data *VRF) fromPlan(plan VRF) {
	data.Device = plan.Device
}
