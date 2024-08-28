// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_rule

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingRuleResultDataSourceEnvelope struct {
	Result EmailRoutingRuleDataSourceModel `json:"result,computed"`
}

type EmailRoutingRuleResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[EmailRoutingRuleDataSourceModel] `json:"result,computed"`
}

type EmailRoutingRuleDataSourceModel struct {
	RuleIdentifier types.String                                                          `tfsdk:"rule_identifier" path:"rule_identifier"`
	ZoneIdentifier types.String                                                          `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool                                                            `tfsdk:"enabled" json:"enabled,computed"`
	ID             types.String                                                          `tfsdk:"id" json:"id,computed"`
	Name           types.String                                                          `tfsdk:"name" json:"name,computed"`
	Priority       types.Float64                                                         `tfsdk:"priority" json:"priority,computed"`
	Tag            types.String                                                          `tfsdk:"tag" json:"tag,computed"`
	Actions        customfield.NestedObjectList[EmailRoutingRuleActionsDataSourceModel]  `tfsdk:"actions" json:"actions,computed"`
	Matchers       customfield.NestedObjectList[EmailRoutingRuleMatchersDataSourceModel] `tfsdk:"matchers" json:"matchers,computed"`
	Filter         *EmailRoutingRuleFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *EmailRoutingRuleDataSourceModel) toListParams() (params email_routing.RuleListParams, diags diag.Diagnostics) {
	params = email_routing.RuleListParams{}

	if !m.Filter.Enabled.IsNull() {
		params.Enabled = cloudflare.F(email_routing.RuleListParamsEnabled(m.Filter.Enabled.ValueBool()))
	}

	return
}

type EmailRoutingRuleActionsDataSourceModel struct {
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.List   `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleMatchersDataSourceModel struct {
	Field types.String `tfsdk:"field" json:"field,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type EmailRoutingRuleFindOneByDataSourceModel struct {
	ZoneIdentifier types.String `tfsdk:"zone_identifier" path:"zone_identifier"`
	Enabled        types.Bool   `tfsdk:"enabled" query:"enabled"`
}
