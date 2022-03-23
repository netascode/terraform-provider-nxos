// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type resourcePhysicalInterfaceType struct{}

func (t resourcePhysicalInterfaceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This resource can manage a physical interface.\n\n- API Documentation: [l1PhysIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/l1:PhysIf/)",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"interface_id": {
				MarkdownDescription: helpers.NewDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"fec_mode": {
				MarkdownDescription: helpers.NewDescription("FEC mode.").AddStringEnumDescription("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"access_vlan": {
				MarkdownDescription: helpers.NewDescription("Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"admin_state": {
				MarkdownDescription: helpers.NewDescription("Administrative port state.").AddStringEnumDescription("up", "down").AddDefaultValueDescription("up").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("up", "down"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("up"),
				},
			},
			"auto_negotiation": {
				MarkdownDescription: helpers.NewDescription("Administrative port auto-negotiation.").AddStringEnumDescription("on", "off", "25G").AddDefaultValueDescription("on").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("on", "off", "25G"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("on"),
				},
			},
			"bandwidth": {
				MarkdownDescription: helpers.NewDescription("The bandwidth parameter for a routed interface, port channel, or subinterface.").AddIntegerRangeDescription(0, 3200000000).AddDefaultValueDescription("0").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 3200000000),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"delay": {
				MarkdownDescription: helpers.NewDescription("The administrative port delay time.").AddIntegerRangeDescription(1, 16777215).AddDefaultValueDescription("1").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 16777215),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(1),
				},
			},
			"description": {
				MarkdownDescription: helpers.NewDescription("Interface description.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"duplex": {
				MarkdownDescription: helpers.NewDescription("Duplex.").AddStringEnumDescription("auto", "full", "half").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("auto", "full", "half"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"layer": {
				MarkdownDescription: helpers.NewDescription("Administrative port layer.").AddStringEnumDescription("Layer2", "Layer3").AddDefaultValueDescription("Layer2").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("Layer2", "Layer3"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("Layer2"),
				},
			},
			"link_logging": {
				MarkdownDescription: helpers.NewDescription("Administrative link logging.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("default", "enable", "disable"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("default"),
				},
			},
			"medium": {
				MarkdownDescription: helpers.NewDescription("The administrative port medium type.").AddStringEnumDescription("broadcast", "p2p").AddDefaultValueDescription("broadcast").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("broadcast", "p2p"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("broadcast"),
				},
			},
			"mode": {
				MarkdownDescription: helpers.NewDescription("Administrative port mode.").AddStringEnumDescription("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag").AddDefaultValueDescription("access").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("access"),
				},
			},
			"mtu": {
				MarkdownDescription: helpers.NewDescription("Administrative port MTU.").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(576, 9216),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(1500),
				},
			},
			"native_vlan": {
				MarkdownDescription: helpers.NewDescription("Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"speed": {
				MarkdownDescription: helpers.NewDescription("Administrative port speed.").AddStringEnumDescription("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"speed_group": {
				MarkdownDescription: helpers.NewDescription("Speed group.").AddStringEnumDescription("unknown", "1000", "10000", "40000", "auto", "25000").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("unknown", "1000", "10000", "40000", "auto", "25000"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"trunk_vlans": {
				MarkdownDescription: helpers.NewDescription("List of trunk VLANs.").AddDefaultValueDescription("1-4094").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("1-4094"),
				},
			},
			"uni_directional_ethernet": {
				MarkdownDescription: helpers.NewDescription("UDE (Uni-Directional Ethernet).").AddStringEnumDescription("disable", "send-only", "receive-only").AddDefaultValueDescription("disable").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("disable", "send-only", "receive-only"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("disable"),
				},
			},
		},
	}, nil
}

func (t resourcePhysicalInterfaceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourcePhysicalInterface{
		provider: provider,
	}, diags
}

type resourcePhysicalInterface struct {
	provider provider
}

func (r resourcePhysicalInterface) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePhysicalInterface) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.Value))

	res, err := r.provider.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePhysicalInterface) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourcePhysicalInterface) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.Value))

	res, err := r.provider.client.DeleteDn(state.Dn.Value)
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourcePhysicalInterface) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
