package bb

import (
	"encoding/json"
	"net/url"

	"github.com/martinusso/bb-sdk-go/internal/rest"
)

const (
	endpointAuthTokenStaging = "https://oauth.hm.bb.com.br"
	endpointAuthToken        = "https://oauth.desenv.bb.com.br"
	oAuthToken               = "/oauth/token"
)

type OAuth struct {
	Staging bool
}

func (o OAuth) Bearer(basic string) (response OAuthResponse, err error) {
	c := rest.Credentials{
		Basic:    basic,
		Endpoint: o.Endpoint(),
	}
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("scope", "cobrancas.boletos-requisicao cobrancas.boletos-info")
	res, err := rest.NewClient(c, params).FormData(oAuthToken)
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Body, &response)
	return
}

func (o OAuth) Endpoint() string {
	if o.Staging {
		return endpointAuthTokenStaging
	}
	return endpointAuthToken
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
