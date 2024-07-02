// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RegionalHostnameResultEnvelope struct {
	Result RegionalHostnameModel `json:"result,computed"`
}

type RegionalHostnameResultDataSourceEnvelope struct {
	Result RegionalHostnameDataSourceModel `json:"result,computed"`
}

type RegionalHostnamesResultDataSourceEnvelope struct {
	Result RegionalHostnamesDataSourceModel `json:"result,computed"`
}

type RegionalHostnameModel struct {
	ID        types.String                      `tfsdk:"id" json:"-,computed"`
	ZoneID    types.String                      `tfsdk:"zone_id" path:"zone_id"`
	RegionKey types.String                      `tfsdk:"region_key" json:"region_key"`
	Hostname  types.String                      `tfsdk:"hostname" json:"hostname"`
	CreatedOn types.String                      `tfsdk:"created_on" json:"created_on,computed"`
	Errors    *[]*RegionalHostnameErrorsModel   `tfsdk:"errors" json:"errors,computed"`
	Messages  *[]*RegionalHostnameMessagesModel `tfsdk:"messages" json:"messages,computed"`
	Success   types.Bool                        `tfsdk:"success" json:"success,computed"`
}

type RegionalHostnameErrorsModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type RegionalHostnameMessagesModel struct {
	Code    types.Int64  `tfsdk:"code" json:"code"`
	Message types.String `tfsdk:"message" json:"message"`
}

type RegionalHostnameDataSourceModel struct {
}

type RegionalHostnamesDataSourceModel struct {
}
