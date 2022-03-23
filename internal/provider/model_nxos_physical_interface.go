// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type PhysicalInterface struct {
	Dn                     types.String `tfsdk:"id"`
	Id                     types.String `tfsdk:"interface_id"`
	FECMode                types.String `tfsdk:"fec_mode"`
	AccessVlan             types.String `tfsdk:"access_vlan"`
	AdminSt                types.String `tfsdk:"admin_state"`
	AutoNeg                types.String `tfsdk:"auto_negotiation"`
	Bw                     types.Int64  `tfsdk:"bandwidth"`
	Delay                  types.Int64  `tfsdk:"delay"`
	Descr                  types.String `tfsdk:"description"`
	Duplex                 types.String `tfsdk:"duplex"`
	Layer                  types.String `tfsdk:"layer"`
	LinkLog                types.String `tfsdk:"link_logging"`
	Medium                 types.String `tfsdk:"medium"`
	Mode                   types.String `tfsdk:"mode"`
	Mtu                    types.Int64  `tfsdk:"mtu"`
	NativeVlan             types.String `tfsdk:"native_vlan"`
	Speed                  types.String `tfsdk:"speed"`
	SpeedGroup             types.String `tfsdk:"speed_group"`
	TrunkVlans             types.String `tfsdk:"trunk_vlans"`
	UniDirectionalEthernet types.String `tfsdk:"uni_directional_ethernet"`
}

func (data PhysicalInterface) getDn() string {
	return fmt.Sprintf("sys/intf/phys-[%s]", data.Id.Value)
}

func (data PhysicalInterface) getClassName() string {
	return "l1PhysIf"
}

func (data PhysicalInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("FECMode", data.FECMode.Value).
		Set("accessVlan", data.AccessVlan.Value).
		Set("adminSt", data.AdminSt.Value).
		Set("autoNeg", data.AutoNeg.Value).
		Set("bw", strconv.FormatInt(data.Bw.Value, 10)).
		Set("delay", strconv.FormatInt(data.Delay.Value, 10)).
		Set("descr", data.Descr.Value).
		Set("duplex", data.Duplex.Value).
		Set("layer", data.Layer.Value).
		Set("linkLog", data.LinkLog.Value).
		Set("medium", data.Medium.Value).
		Set("mode", data.Mode.Value).
		Set("mtu", strconv.FormatInt(data.Mtu.Value, 10)).
		Set("nativeVlan", data.NativeVlan.Value).
		Set("speed", data.Speed.Value).
		Set("speedGroup", data.SpeedGroup.Value).
		Set("trunkVlans", data.TrunkVlans.Value).
		Set("uniDirectionalEthernet", data.UniDirectionalEthernet.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PhysicalInterface) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.FECMode.Value = res.Get("*.attributes.FECMode").String()
	data.AccessVlan.Value = res.Get("*.attributes.accessVlan").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
	data.AutoNeg.Value = res.Get("*.attributes.autoNeg").String()
	data.Bw.Value = res.Get("*.attributes.bw").Int()
	data.Delay.Value = res.Get("*.attributes.delay").Int()
	data.Descr.Value = res.Get("*.attributes.descr").String()
	data.Duplex.Value = res.Get("*.attributes.duplex").String()
	data.Layer.Value = res.Get("*.attributes.layer").String()
	data.LinkLog.Value = res.Get("*.attributes.linkLog").String()
	data.Medium.Value = res.Get("*.attributes.medium").String()
	data.Mode.Value = res.Get("*.attributes.mode").String()
	data.Mtu.Value = res.Get("*.attributes.mtu").Int()
	data.NativeVlan.Value = res.Get("*.attributes.nativeVlan").String()
	data.Speed.Value = res.Get("*.attributes.speed").String()
	data.SpeedGroup.Value = res.Get("*.attributes.speedGroup").String()
	data.TrunkVlans.Value = res.Get("*.attributes.trunkVlans").String()
	data.UniDirectionalEthernet.Value = res.Get("*.attributes.uniDirectionalEthernet").String()
}

func (data *PhysicalInterface) fromPlan(plan PhysicalInterface) {
}
