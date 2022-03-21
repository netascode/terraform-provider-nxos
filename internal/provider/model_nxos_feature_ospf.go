// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureOSPF struct {
	Dn types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureOSPF) getDn() string {
	return "sys/fm/ospf"
}

func (data FeatureOSPF) getClassName() string {
	return "fmOspf"
}

func (data FeatureOSPF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureOSPF) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *FeatureOSPF) fromPlan(plan FeatureOSPF) {
}