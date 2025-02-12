package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	"github.com/authelia/authelia/v4/internal/middlewares"
)

// OpenIDConnectConfigurationWellKnownGET handles requests to a .well-known endpoint (RFC5785) which returns the
// OpenID Connect Discovery 1.0 metadata.
//
// https://datatracker.ietf.org/doc/html/rfc5785
//
// https://openid.net/specs/openid-connect-discovery-1_0.html
func OpenIDConnectConfigurationWellKnownGET(ctx *middlewares.AutheliaCtx) {
	issuer, err := ctx.ExternalRootURL()
	if err != nil {
		ctx.Logger.Errorf("Error occurred determining OpenID Connect issuer details: %+v", err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	wellKnown := ctx.Providers.OpenIDConnect.GetOpenIDConnectWellKnownConfiguration(issuer)

	ctx.SetContentType("application/json")

	if err = json.NewEncoder(ctx).Encode(wellKnown); err != nil {
		ctx.Logger.Errorf("Error occurred in JSON encode: %+v", err)
		// TODO: Determine if this is the appropriate error code here.
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}
}

// OAuthAuthorizationServerWellKnownGET handles requests to a .well-known endpoint (RFC5785) which returns the
// OAuth 2.0 Authorization Server Metadata (RFC8414).
//
// https://datatracker.ietf.org/doc/html/rfc5785
//
// https://datatracker.ietf.org/doc/html/rfc8414
func OAuthAuthorizationServerWellKnownGET(ctx *middlewares.AutheliaCtx) {
	issuer, err := ctx.ExternalRootURL()
	if err != nil {
		ctx.Logger.Errorf("Error occurred determining OpenID Connect issuer details: %+v", err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)

		return
	}

	wellKnown := ctx.Providers.OpenIDConnect.GetOAuth2WellKnownConfiguration(issuer)

	ctx.SetContentType("application/json")

	if err = json.NewEncoder(ctx).Encode(wellKnown); err != nil {
		ctx.Logger.Errorf("Error occurred in JSON encode: %+v", err)
		// TODO: Determine if this is the appropriate error code here.
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)

		return
	}
}
