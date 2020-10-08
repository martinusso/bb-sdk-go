package cobranca

import (
	"encoding/json"
	"fmt"
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

func (c client) BaixarBoleto(numero string) error {
	cred := rest.Credentials{
		Bearer:   c.cc.Token,
		AppKey:   c.cc.AppKey,
		Endpoint: c.endpoint(),
	}
	params := url.Values{}

	path := fmt.Sprintf("/boletos/%s/baixar", numero)

	res, err := rest.NewClient(cred, params).Post(path, nil)
	if err != nil {
		return err
	}
	fmt.Println(string(res.Body))
	return nil
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
	err = json.Unmarshal(res.Body, &boletos)
	return
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
	err = json.Unmarshal(res.Body, &boleto)
	return
}

func (c client) endpoint() string {
	if c.cc.Staging {
		return endpointCobrancasStaging
	}
	return endpointCobrancasProd
}
