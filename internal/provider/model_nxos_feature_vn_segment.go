// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureVNSegment struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureVNSegment) getDn() string {
	return "sys/fm/vnsegment"
}

func (data FeatureVNSegment) getClassName() string {
	return "fmVnSegment"
}

func (data FeatureVNSegment) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureVNSegment) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *FeatureVNSegment) fromPlan(plan FeatureVNSegment) {
	data.Device = plan.Device
}
