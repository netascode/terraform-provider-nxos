// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

type PIMVRF struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	AdminSt types.String `tfsdk:"admin_state"`
	Bfd     types.Bool   `tfsdk:"bfd"`
}

func (data PIMVRF) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]", data.Name.Value)
}

func (data PIMVRF) getClassName() string {
	return "pimDom"
}

func (data PIMVRF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.Value).
		Set("adminSt", data.AdminSt.Value).
		Set("bfd", strconv.FormatBool(data.Bfd.Value))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PIMVRF) fromBody(res gjson.Result) {
	data.Name.Value = res.Get("*.attributes.name").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
	data.Bfd.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.bfd").String())
}

func (data *PIMVRF) fromPlan(plan PIMVRF) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
}
