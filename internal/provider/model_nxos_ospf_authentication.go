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

type OSPFAuthentication struct {
	Device           types.String `tfsdk:"device"`
	Dn               types.String `tfsdk:"id"`
	Inst             types.String `tfsdk:"instance_name"`
	Name             types.String `tfsdk:"vrf_name"`
	Id               types.String `tfsdk:"interface_id"`
	Key              types.String `tfsdk:"key"`
	KeyId            types.Int64  `tfsdk:"key_id"`
	KeySecureMode    types.Bool   `tfsdk:"key_secure_mode"`
	Keychain         types.String `tfsdk:"keychain"`
	Md5key           types.String `tfsdk:"md5_key"`
	Md5keySecureMode types.Bool   `tfsdk:"md5_key_secure_mode"`
	Type             types.String `tfsdk:"type"`
}

func (data OSPFAuthentication) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/if-[%s]/authnew", data.Inst.Value, data.Name.Value, data.Id.Value)
}

func (data OSPFAuthentication) getClassName() string {
	return "ospfAuthNewP"
}

func (data OSPFAuthentication) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("key", data.Key.Value).
		Set("keyId", strconv.FormatInt(data.KeyId.Value, 10)).
		Set("keySecureMode", strconv.FormatBool(data.KeySecureMode.Value)).
		Set("keychain", data.Keychain.Value).
		Set("md5key", data.Md5key.Value).
		Set("md5keySecureMode", strconv.FormatBool(data.Md5keySecureMode.Value)).
		Set("type", data.Type.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *OSPFAuthentication) fromBody(res gjson.Result) {
	data.KeyId.Value = res.Get("*.attributes.keyId").Int()
	data.KeySecureMode.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.keySecureMode").String())
	data.Keychain.Value = res.Get("*.attributes.keychain").String()
	data.Md5keySecureMode.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.md5keySecureMode").String())
	data.Type.Value = res.Get("*.attributes.type").String()
}

func (data *OSPFAuthentication) fromPlan(plan OSPFAuthentication) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Inst.Value = plan.Inst.Value
	data.Name.Value = plan.Name.Value
	data.Id.Value = plan.Id.Value
	data.Key.Value = plan.Key.Value
	data.Md5key.Value = plan.Md5key.Value
}
