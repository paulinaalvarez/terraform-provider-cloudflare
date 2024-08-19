// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AuthenticatedOriginPullsCertificateDataSource{}

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate from the authority expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Computed:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "The zone's leaf certificate.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: "Status of the certificate activation.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"initializing",
						"pending_deployment",
						"pending_deletion",
						"active",
						"deleted",
						"deployment_timed_out",
						"deletion_timed_out",
					),
				},
			},
			"uploaded_on": schema.StringAttribute{
				Description: "This is the time the certificate was uploaded.",
				Computed:    true,
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *AuthenticatedOriginPullsCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AuthenticatedOriginPullsCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("certificate_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("certificate_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}
