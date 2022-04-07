// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type dataSourceBGPInstanceDomainPeerContainerAfMaxPrefixType struct{}

func (t dataSourceBGPInstanceDomainPeerContainerAfMaxPrefixType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the BGP peer template Maximum Prefix Policy configuration.", "bgpMaxPfxP", "Routing%20and%20Forwarding/bgp:MaxPfxP/").String,

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
			},
			"vrf": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"template_name": {
				MarkdownDescription: "Peer template name.",
				Type:                types.StringType,
				Required:            true,
			},
			"af": {
				MarkdownDescription: "Address Family.",
				Type:                types.StringType,
				Required:            true,
			},
			"action": {
				MarkdownDescription: "Action to do when limit is exceeded.",
				Type:                types.StringType,
				Computed:            true,
			},
			"maximum_prefix": {
				MarkdownDescription: "Maximum number of prefixes allowed from the peer.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"restart_time": {
				MarkdownDescription: "The period of time in minutes before restarting the peer when the prefix limit is reached.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"threshold": {
				MarkdownDescription: "The period of time in minutes before restarting the peer when the prefix limit is reached.",
				Type:                types.Int64Type,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceBGPInstanceDomainPeerContainerAfMaxPrefixType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceBGPInstanceDomainPeerContainerAfMaxPrefix{
		provider: provider,
	}, diags
}

type dataSourceBGPInstanceDomainPeerContainerAfMaxPrefix struct {
	provider provider
}

func (d dataSourceBGPInstanceDomainPeerContainerAfMaxPrefix) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state BGPInstanceDomainPeerContainerAfMaxPrefix

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn(), nxos.OverrideUrl(d.provider.devices[config.Device.Value]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
