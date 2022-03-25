// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type DefaultQOSClassMapDSCP struct {
	Dn             types.String `tfsdk:"id"`
	Class_map_name types.String `tfsdk:"class_map_name"`
	Val            types.String `tfsdk:"value"`
}

func (data DefaultQOSClassMapDSCP) getDn() string {
	return fmt.Sprintf("sys/ipqos/dflt/c/name-[%s]/dscp-[%v]", data.Class_map_name.Value, data.Val.Value)
}

func (data DefaultQOSClassMapDSCP) getClassName() string {
	return "ipqosDscp"
}

func (data DefaultQOSClassMapDSCP) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("val", data.Val.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *DefaultQOSClassMapDSCP) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Val.Value = res.Get("*.attributes.val").String()
}

func (data *DefaultQOSClassMapDSCP) fromPlan(plan DefaultQOSClassMapDSCP) {
	data.Class_map_name.Value = plan.Class_map_name.Value
}