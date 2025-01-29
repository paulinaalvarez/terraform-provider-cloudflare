// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared_config_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/services/zero_trust_tunnel_cloudflared_config"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/test_helpers"
)

func TestZeroTrustTunnelCloudflaredConfigModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*zero_trust_tunnel_cloudflared_config.ZeroTrustTunnelCloudflaredConfigModel)(nil)
	schema := zero_trust_tunnel_cloudflared_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
