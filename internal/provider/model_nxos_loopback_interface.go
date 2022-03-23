// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type LoopbackInterface struct {
	Dn      types.String `tfsdk:"id"`
	Id      types.String `tfsdk:"interface_id"`
	AdminSt types.String `tfsdk:"admin_state"`
	Descr   types.String `tfsdk:"description"`
}

func (data LoopbackInterface) getDn() string {
	return fmt.Sprintf("sys/intf/lb-[%s]", data.Id.Value)
}

func (data LoopbackInterface) getClassName() string {
	return "l3LbRtdIf"
}

func (data LoopbackInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("adminSt", data.AdminSt.Value).
		Set("descr", data.Descr.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *LoopbackInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
	data.Descr.Value = res.Get("*.attributes.descr").String()
}

func (data *LoopbackInterface) fromPlan(plan LoopbackInterface) {
}
