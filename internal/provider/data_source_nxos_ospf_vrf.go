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

type dataSourceOSPFVRFType struct{}

func (t dataSourceOSPFVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the OSPF VRF configuration.", "ospfDom", "Routing%20and%20Forwarding/ospf:Dom/").String,

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
			"instance_name": {
				MarkdownDescription: "OSPF instance name.",
				Type:                types.StringType,
				Required:            true,
			},
			"name": {
				MarkdownDescription: "VRF name.",
				Type:                types.StringType,
				Required:            true,
			},
			"admin_state": {
				MarkdownDescription: "Administrative state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"bandwidth_reference": {
				MarkdownDescription: "Bandwidth reference value.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"banwidth_reference_unit": {
				MarkdownDescription: "Bandwidth reference unit.",
				Type:                types.StringType,
				Computed:            true,
			},
			"distance": {
				MarkdownDescription: "Administrative distance preference.",
				Type:                types.Int64Type,
				Computed:            true,
			},
			"router_id": {
				MarkdownDescription: "Router ID.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceOSPFVRFType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceOSPFVRF{
		provider: provider,
	}, diags
}

type dataSourceOSPFVRF struct {
	provider provider
}

func (d dataSourceOSPFVRF) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state OSPFVRF

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
