// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r ObservatoryScheduledTestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"url": schema.StringAttribute{
				Description:   "A URL.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"count": schema.Float64Attribute{
				Description: "Number of items affected.",
				Computed:    true,
			},
			"frequency": schema.StringAttribute{
				Description: "The frequency of the test.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
				},
			},
			"region": schema.StringAttribute{
				Description: "A test region.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asia-east1", "asia-northeast1", "asia-northeast2", "asia-south1", "asia-southeast1", "australia-southeast1", "europe-north1", "europe-southwest1", "europe-west1", "europe-west2", "europe-west3", "europe-west4", "europe-west8", "europe-west9", "me-west1", "southamerica-east1", "us-central1", "us-east1", "us-east4", "us-south1", "us-west1"),
				},
			},
		},
	}
}
