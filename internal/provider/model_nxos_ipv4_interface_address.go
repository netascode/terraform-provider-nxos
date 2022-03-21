// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type IPv4InterfaceAddress struct {
	Dn types.String `tfsdk:"id"`
	Dom types.String `tfsdk:"vrf"`
	Id types.String `tfsdk:"interface_id"`
	Addr types.String `tfsdk:"address"`
}

func (data IPv4InterfaceAddress) getDn() string {
	return fmt.Sprintf("sys/ipv4/inst/dom-[%s]/if-[%s]/addr-[%s]", data.Dom.Value, data.Id.Value, data.Addr.Value)
}

func (data IPv4InterfaceAddress) getClassName() string {
	return "ipv4Addr"
}

func (data IPv4InterfaceAddress) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("addr", data.Addr.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *IPv4InterfaceAddress) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Addr.Value = res.Get("*.attributes.addr").String()
}

func (data *IPv4InterfaceAddress) fromPlan(plan IPv4InterfaceAddress) {
	data.Dom.Value = plan.Dom.Value
	data.Id.Value = plan.Id.Value
}