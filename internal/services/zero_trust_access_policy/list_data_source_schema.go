// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessPoliciesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPoliciesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The UUID of the policy",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"decision": schema.StringAttribute{
							Description: "The action Access will take if a user matches this policy. Infrastructure application policies can only use the Allow action.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"allow",
									"deny",
									"non_identity",
									"bypass",
								),
							},
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match the policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPoliciesExcludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPoliciesIncludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Access policy.",
							Computed:    true,
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match the policy, a user must meet all of the Require rules.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustAccessPoliciesRequireDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"email": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireEmailDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the user.",
												Computed:    true,
											},
										},
									},
									"email_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireEmailListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created email list.",
												Computed:    true,
											},
										},
									},
									"email_domain": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"domain": schema.StringAttribute{
												Description: "The email domain to match.",
												Computed:    true,
											},
										},
									},
									"everyone": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireEveryoneDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"ip": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireIPDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"ip": schema.StringAttribute{
												Description: "An IPv4 or IPv6 CIDR block.",
												Computed:    true,
											},
										},
									},
									"ip_list": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireIPListDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created IP list.",
												Computed:    true,
											},
										},
									},
									"certificate": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireCertificateDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{},
									},
									"group": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireGroupDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of a previously created Access group.",
												Computed:    true,
											},
										},
									},
									"azure_ad": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireAzureADDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "The ID of an Azure group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Azure identity provider.",
												Computed:    true,
											},
										},
									},
									"github_organization": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Github identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the organization.",
												Computed:    true,
											},
										},
									},
									"gsuite": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"email": schema.StringAttribute{
												Description: "The email of the Google Workspace group.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Google Workspace identity provider.",
												Computed:    true,
											},
										},
									},
									"okta": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireOktaDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your Okta identity provider.",
												Computed:    true,
											},
											"name": schema.StringAttribute{
												Description: "The name of the Okta group.",
												Computed:    true,
											},
										},
									},
									"saml": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireSAMLDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description: "The name of the SAML attribute.",
												Computed:    true,
											},
											"attribute_value": schema.StringAttribute{
												Description: "The SAML attribute value to look for.",
												Computed:    true,
											},
											"identity_provider_id": schema.StringAttribute{
												Description: "The ID of your SAML identity provider.",
												Computed:    true,
											},
										},
									},
									"service_token": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"token_id": schema.StringAttribute{
												Description: "The ID of a Service Token.",
												Computed:    true,
											},
										},
									},
									"any_valid_service_token": schema.SingleNestedAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireAnyValidServiceTokenDataSourceModel](ctx),
										Attributes:  map[string]schema.Attribute{},
									},
									"external_evaluation": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"evaluate_url": schema.StringAttribute{
												Description: "The API endpoint containing your business logic.",
												Computed:    true,
											},
											"keys_url": schema.StringAttribute{
												Description: "The API endpoint containing the key that Access uses to verify that the response came from your API.",
												Computed:    true,
											},
										},
									},
									"geo": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireGeoDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"country_code": schema.StringAttribute{
												Description: "The country code that should be matched.",
												Computed:    true,
											},
										},
									},
									"auth_method": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"auth_method": schema.StringAttribute{
												Description: "The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176#section-2.",
												Computed:    true,
											},
										},
									},
									"device_posture": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"integration_uid": schema.StringAttribute{
												Description: "The ID of a device posture integration.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustAccessPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustAccessPoliciesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
