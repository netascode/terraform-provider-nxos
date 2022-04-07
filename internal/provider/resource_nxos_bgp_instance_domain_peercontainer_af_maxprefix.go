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

type resourceBGPInstanceDomainPeerContainerAfMaxPrefixType struct{}

func (t resourceBGPInstanceDomainPeerContainerAfMaxPrefixType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the BGP peer template Maximum Prefix Policy configuration.", "bgpMaxPfxP", "Routing%20and%20Forwarding/bgp:MaxPfxP/").AddParents("bgp_instance_domain_peercontainer_af").String,

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"vrf": {
				MarkdownDescription: helpers.NewAttributeDescription("VRF name.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"template_name": {
				MarkdownDescription: helpers.NewAttributeDescription("Peer template name.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"af": {
				MarkdownDescription: helpers.NewAttributeDescription("Address Family.").AddStringEnumDescription("ipv4-ucast", "ipv4-mcast", "vpnv4-ucast", "ipv6-ucast", "ipv6-mcast", "vpnv6-ucast", "vpnv6-mcast", "l2vpn-evpn", "ipv4-lucast", "ipv6-lucast", "lnkstate", "ipv4-mvpn", "ipv6-mvpn", "l2vpn-vpls", "ipv4-mdt").AddDefaultValueDescription("ipv4-ucast").String,
				Type:                types.StringType,
				Required:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("ipv4-ucast", "ipv4-mcast", "vpnv4-ucast", "ipv6-ucast", "ipv6-mcast", "vpnv6-ucast", "vpnv6-mcast", "l2vpn-evpn", "ipv4-lucast", "ipv6-lucast", "lnkstate", "ipv4-mvpn", "ipv6-mvpn", "l2vpn-vpls", "ipv4-mdt"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"action": {
				MarkdownDescription: helpers.NewAttributeDescription("Action to do when limit is exceeded.").AddDefaultValueDescription("shut").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("shut"),
				},
			},
			"maximum_prefix": {
				MarkdownDescription: helpers.NewAttributeDescription("Maximum number of prefixes allowed from the peer.").AddIntegerRangeDescription(0, 4294967295).String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
			},
			"restart_time": {
				MarkdownDescription: helpers.NewAttributeDescription("The period of time in minutes before restarting the peer when the prefix limit is reached.").AddDefaultValueDescription("0").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"threshold": {
				MarkdownDescription: helpers.NewAttributeDescription("The period of time in minutes before restarting the peer when the prefix limit is reached.").AddIntegerRangeDescription(0, 100).AddDefaultValueDescription("0").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 100),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(0),
				},
			},
		},
	}, nil
}

func (t resourceBGPInstanceDomainPeerContainerAfMaxPrefixType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceBGPInstanceDomainPeerContainerAfMaxPrefix{
		provider: provider,
	}, diags
}

type resourceBGPInstanceDomainPeerContainerAfMaxPrefix struct {
	provider provider
}

func (r resourceBGPInstanceDomainPeerContainerAfMaxPrefix) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state BGPInstanceDomainPeerContainerAfMaxPrefix

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
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

func (r resourceBGPInstanceDomainPeerContainerAfMaxPrefix) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state BGPInstanceDomainPeerContainerAfMaxPrefix

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.Value))

	res, err := r.provider.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[state.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceBGPInstanceDomainPeerContainerAfMaxPrefix) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state BGPInstanceDomainPeerContainerAfMaxPrefix

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
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

func (r resourceBGPInstanceDomainPeerContainerAfMaxPrefix) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state BGPInstanceDomainPeerContainerAfMaxPrefix

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.Value))

	res, err := r.provider.client.DeleteDn(state.Dn.Value, nxos.OverrideUrl(r.provider.devices[state.Device.Value]))
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

func (r resourceBGPInstanceDomainPeerContainerAfMaxPrefix) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
