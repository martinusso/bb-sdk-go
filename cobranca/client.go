package cobranca

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/martinusso/bb-sdk-go"
	"github.com/martinusso/bb-sdk-go/internal/rest"
)

const (
	endpointCobrancasStaging = "https://api.hm.bb.com.br/cobrancas/v1"
	endpointCobrancasProd    = "https://api.bb.com.br/cobrancas/v1"
)

type client struct {
	cc bb.ClientCredentials
}

func NewClient(cc bb.ClientCredentials) client {
	return client{
		cc: cc}
}

func (c client) ListarBoletos(p ListaBoletosParams) (boletos BoletosListagem, err error) {
	cred := rest.Credentials{
		Bearer:   c.cc.Token,
		AppKey:   c.cc.AppKey,
		Endpoint: c.endpoint(),
	}

	res, err := rest.NewClient(cred, p.Values()).Get("/boletos")
	if err != nil {
		return
	}

	if res.Code == http.StatusOK {
		err = json.Unmarshal(res.Body, &boletos)
		return
	}

	var cobErr *ErrorListaBoletos
	err = json.Unmarshal(res.Body, &cobErr)
	if err != nil {
		return
	}
	return boletos, cobErr
}

func (c client) RegistrarBoleto(b Boleto) (boleto RegistroBoleto, err error) {
	if err = b.Validate(); err != nil {
		return
	}
	if b.IndicadorPermissaoRecebimentoParcial == "" {
		b.IndicadorPermissaoRecebimentoParcial = "N"
	}

	cred := rest.Credentials{
		Bearer:   c.cc.Token,
		AppKey:   c.cc.AppKey,
		Endpoint: c.endpoint(),
	}
	params := url.Values{}

	res, err := rest.NewClient(cred, params).Post("/boletos", b)
	if err != nil {
		return
	}

	if res.Code == http.StatusCreated {
		err = json.Unmarshal(res.Body, &boleto)
		return
	}

	var cobErr *bb.ErrorBB
	err = json.Unmarshal(res.Body, &cobErr)
	if err != nil {
		return
	}
	return boleto, cobErr
}

func (c client) endpoint() string {
	if c.cc.Staging {
		return endpointCobrancasStaging
	}
	return endpointCobrancasProd
}
