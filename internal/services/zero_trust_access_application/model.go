// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationResultEnvelope struct {
	Result ZeroTrustAccessApplicationModel `json:"result"`
}

type ZeroTrustAccessApplicationModel struct {
	ID                       types.String                                      `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                                      `tfsdk:"account_id" path:"account_id"`
	ZoneID                   types.String                                      `tfsdk:"zone_id" path:"zone_id"`
	AllowAuthenticateViaWARP types.Bool                                        `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp"`
	AppLauncherLogoURL       types.String                                      `tfsdk:"app_launcher_logo_url" json:"app_launcher_logo_url"`
	BgColor                  types.String                                      `tfsdk:"bg_color" json:"bg_color"`
	CustomDenyMessage        types.String                                      `tfsdk:"custom_deny_message" json:"custom_deny_message"`
	CustomDenyURL            types.String                                      `tfsdk:"custom_deny_url" json:"custom_deny_url"`
	CustomNonIdentityDenyURL types.String                                      `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url"`
	Domain                   types.String                                      `tfsdk:"domain" json:"domain"`
	HeaderBgColor            types.String                                      `tfsdk:"header_bg_color" json:"header_bg_color"`
	LogoURL                  types.String                                      `tfsdk:"logo_url" json:"logo_url"`
	Name                     types.String                                      `tfsdk:"name" json:"name"`
	OptionsPreflightBypass   types.Bool                                        `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass"`
	SameSiteCookieAttribute  types.String                                      `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute"`
	ServiceAuth401Redirect   types.Bool                                        `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect"`
	SkipInterstitial         types.Bool                                        `tfsdk:"skip_interstitial" json:"skip_interstitial"`
	Type                     types.String                                      `tfsdk:"type" json:"type"`
	AllowedIdPs              *[]types.String                                   `tfsdk:"allowed_idps" json:"allowed_idps"`
	CustomPages              *[]types.String                                   `tfsdk:"custom_pages" json:"custom_pages"`
	SelfHostedDomains        *[]types.String                                   `tfsdk:"self_hosted_domains" json:"self_hosted_domains"`
	Tags                     *[]types.String                                   `tfsdk:"tags" json:"tags"`
	CORSHeaders              *ZeroTrustAccessApplicationCORSHeadersModel       `tfsdk:"cors_headers" json:"cors_headers"`
	FooterLinks              *[]*ZeroTrustAccessApplicationFooterLinksModel    `tfsdk:"footer_links" json:"footer_links"`
	LandingPageDesign        *ZeroTrustAccessApplicationLandingPageDesignModel `tfsdk:"landing_page_design" json:"landing_page_design"`
	Policies                 *[]*ZeroTrustAccessApplicationPoliciesModel       `tfsdk:"policies" json:"policies"`
	SaaSApp                  *ZeroTrustAccessApplicationSaaSAppModel           `tfsdk:"saas_app" json:"saas_app"`
	SCIMConfig               *ZeroTrustAccessApplicationSCIMConfigModel        `tfsdk:"scim_config" json:"scim_config"`
	AppLauncherVisible       types.Bool                                        `tfsdk:"app_launcher_visible" json:"app_launcher_visible,computed_optional"`
	AutoRedirectToIdentity   types.Bool                                        `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed_optional"`
	EnableBindingCookie      types.Bool                                        `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed_optional"`
	HTTPOnlyCookieAttribute  types.Bool                                        `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute,computed_optional"`
	PathCookieAttribute      types.Bool                                        `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute,computed_optional"`
	SessionDuration          types.String                                      `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	SkipAppLauncherLoginPage types.Bool                                        `tfsdk:"skip_app_launcher_login_page" json:"skip_app_launcher_login_page,computed_optional"`
}

type ZeroTrustAccessApplicationCORSHeadersModel struct {
	AllowAllHeaders  types.Bool                     `tfsdk:"allow_all_headers" json:"allow_all_headers,computed_optional"`
	AllowAllMethods  types.Bool                     `tfsdk:"allow_all_methods" json:"allow_all_methods,computed_optional"`
	AllowAllOrigins  types.Bool                     `tfsdk:"allow_all_origins" json:"allow_all_origins,computed_optional"`
	AllowCredentials types.Bool                     `tfsdk:"allow_credentials" json:"allow_credentials,computed_optional"`
	AllowedHeaders   customfield.List[types.String] `tfsdk:"allowed_headers" json:"allowed_headers,computed_optional"`
	AllowedMethods   customfield.List[types.String] `tfsdk:"allowed_methods" json:"allowed_methods,computed_optional"`
	AllowedOrigins   customfield.List[types.String] `tfsdk:"allowed_origins" json:"allowed_origins,computed_optional"`
	MaxAge           types.Float64                  `tfsdk:"max_age" json:"max_age,computed_optional"`
}

type ZeroTrustAccessApplicationFooterLinksModel struct {
	Name types.String `tfsdk:"name" json:"name"`
	URL  types.String `tfsdk:"url" json:"url"`
}

type ZeroTrustAccessApplicationLandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color" json:"button_color,computed_optional"`
	ButtonTextColor types.String `tfsdk:"button_text_color" json:"button_text_color,computed_optional"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,computed_optional"`
	Message         types.String `tfsdk:"message" json:"message,computed_optional"`
	Title           types.String `tfsdk:"title" json:"title,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesModel struct {
	ID         types.String `tfsdk:"id" json:"id,computed_optional"`
	Precedence types.Int64  `tfsdk:"precedence" json:"precedence,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppModel struct {
	AuthType                      types.String                                                    `tfsdk:"auth_type" json:"auth_type,computed_optional"`
	ConsumerServiceURL            types.String                                                    `tfsdk:"consumer_service_url" json:"consumer_service_url"`
	CreatedAt                     timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomAttributes              *ZeroTrustAccessApplicationSaaSAppCustomAttributesModel         `tfsdk:"custom_attributes" json:"custom_attributes"`
	DefaultRelayState             types.String                                                    `tfsdk:"default_relay_state" json:"default_relay_state"`
	IdPEntityID                   types.String                                                    `tfsdk:"idp_entity_id" json:"idp_entity_id"`
	NameIDFormat                  types.String                                                    `tfsdk:"name_id_format" json:"name_id_format"`
	NameIDTransformJsonata        types.String                                                    `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata"`
	PublicKey                     types.String                                                    `tfsdk:"public_key" json:"public_key,computed_optional"`
	SAMLAttributeTransformJsonata types.String                                                    `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata"`
	SPEntityID                    types.String                                                    `tfsdk:"sp_entity_id" json:"sp_entity_id"`
	SSOEndpoint                   types.String                                                    `tfsdk:"sso_endpoint" json:"sso_endpoint"`
	UpdatedAt                     timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AccessTokenLifetime           types.String                                                    `tfsdk:"access_token_lifetime" json:"access_token_lifetime"`
	AllowPKCEWithoutClientSecret  types.Bool                                                      `tfsdk:"allow_pkce_without_client_secret" json:"allow_pkce_without_client_secret"`
	AppLauncherURL                types.String                                                    `tfsdk:"app_launcher_url" json:"app_launcher_url"`
	ClientID                      types.String                                                    `tfsdk:"client_id" json:"client_id"`
	ClientSecret                  types.String                                                    `tfsdk:"client_secret" json:"client_secret"`
	CustomClaims                  *ZeroTrustAccessApplicationSaaSAppCustomClaimsModel             `tfsdk:"custom_claims" json:"custom_claims"`
	GrantTypes                    *[]types.String                                                 `tfsdk:"grant_types" json:"grant_types"`
	GroupFilterRegex              types.String                                                    `tfsdk:"group_filter_regex" json:"group_filter_regex"`
	HybridAndImplicitOptions      *ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel `tfsdk:"hybrid_and_implicit_options" json:"hybrid_and_implicit_options"`
	RedirectURIs                  *[]types.String                                                 `tfsdk:"redirect_uris" json:"redirect_uris"`
	RefreshTokenOptions           *ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel      `tfsdk:"refresh_token_options" json:"refresh_token_options"`
	Scopes                        *[]types.String                                                 `tfsdk:"scopes" json:"scopes"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesModel struct {
	FriendlyName types.String                                                                           `tfsdk:"friendly_name" json:"friendly_name,computed_optional"`
	Name         types.String                                                                           `tfsdk:"name" json:"name,computed_optional"`
	NameFormat   types.String                                                                           `tfsdk:"name_format" json:"name_format,computed_optional"`
	Required     types.Bool                                                                             `tfsdk:"required" json:"required,computed_optional"`
	Source       customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel] `tfsdk:"source" json:"source,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed_optional"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsModel struct {
	Name     types.String                                                                       `tfsdk:"name" json:"name,computed_optional"`
	Required types.Bool                                                                         `tfsdk:"required" json:"required,computed_optional"`
	Scope    types.String                                                                       `tfsdk:"scope" json:"scope,computed_optional"`
	Source   customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel] `tfsdk:"source" json:"source,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed_optional"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint" json:"return_access_token_from_authorization_endpoint,computed_optional"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint" json:"return_id_token_from_authorization_endpoint,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel struct {
	Lifetime types.String `tfsdk:"lifetime" json:"lifetime,computed_optional"`
}

type ZeroTrustAccessApplicationSCIMConfigModel struct {
	IdPUID             types.String                                                                      `tfsdk:"idp_uid" json:"idp_uid"`
	RemoteURI          types.String                                                                      `tfsdk:"remote_uri" json:"remote_uri"`
	Authentication     customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigAuthenticationModel] `tfsdk:"authentication" json:"authentication,computed_optional"`
	DeactivateOnDelete types.Bool                                                                        `tfsdk:"deactivate_on_delete" json:"deactivate_on_delete,computed_optional"`
	Enabled            types.Bool                                                                        `tfsdk:"enabled" json:"enabled,computed_optional"`
	Mappings           customfield.NestedObjectList[ZeroTrustAccessApplicationSCIMConfigMappingsModel]   `tfsdk:"mappings" json:"mappings,computed_optional"`
}

type ZeroTrustAccessApplicationSCIMConfigAuthenticationModel struct {
	Password         types.String    `tfsdk:"password" json:"password"`
	Scheme           types.String    `tfsdk:"scheme" json:"scheme"`
	User             types.String    `tfsdk:"user" json:"user"`
	Token            types.String    `tfsdk:"token" json:"token"`
	AuthorizationURL types.String    `tfsdk:"authorization_url" json:"authorization_url"`
	ClientID         types.String    `tfsdk:"client_id" json:"client_id"`
	ClientSecret     types.String    `tfsdk:"client_secret" json:"client_secret"`
	TokenURL         types.String    `tfsdk:"token_url" json:"token_url"`
	Scopes           *[]types.String `tfsdk:"scopes" json:"scopes"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsModel struct {
	Schema           types.String                                                                          `tfsdk:"schema" json:"schema"`
	Enabled          types.Bool                                                                            `tfsdk:"enabled" json:"enabled,computed_optional"`
	Filter           types.String                                                                          `tfsdk:"filter" json:"filter,computed_optional"`
	Operations       customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel] `tfsdk:"operations" json:"operations,computed_optional"`
	TransformJsonata types.String                                                                          `tfsdk:"transform_jsonata" json:"transform_jsonata,computed_optional"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel struct {
	Create types.Bool `tfsdk:"create" json:"create,computed_optional"`
	Delete types.Bool `tfsdk:"delete" json:"delete,computed_optional"`
	Update types.Bool `tfsdk:"update" json:"update,computed_optional"`
}
