// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type SVIInterfaceVRF struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Id     types.String `tfsdk:"interface_id"`
	TDn    types.String `tfsdk:"vrf_dn"`
}

func (data SVIInterfaceVRF) getDn() string {
	return fmt.Sprintf("sys/intf/svi-[%s]/rtvrfMbr", data.Id.Value)
}

func (data SVIInterfaceVRF) getClassName() string {
	return "nwRtVrfMbr"
}

func (data SVIInterfaceVRF) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("tDn", data.TDn.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *SVIInterfaceVRF) fromBody(res gjson.Result) {
	data.TDn.Value = res.Get("*.attributes.tDn").String()
}

func (data *SVIInterfaceVRF) fromPlan(plan SVIInterfaceVRF) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Id.Value = plan.Id.Value
}
