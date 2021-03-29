package cobranca

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/martinusso/bb-sdk-go"
	"github.com/martinusso/bb-sdk-go/internal/rest"
)

const (
	endpointCobrancasStaging = "https://api.hm.bb.com.br/cobrancas/v2"
	endpointCobrancasProd    = "https://api.bb.com.br/cobrancas/v2"
)

type client struct {
	cc bb.ClientCredentials
}

func NewClient(cc bb.ClientCredentials) client {
	return client{
		cc: cc}
}

func (c client) BaixarBoleto(numeroConvenio, nossoNumero string) (baixa RetornoBaixa, err error) {
	cred := rest.Credentials{
		Bearer:   c.cc.Token,
		AppKey:   c.cc.AppKey,
		Endpoint: c.endpoint(),
	}
	params := url.Values{}

	path := fmt.Sprintf("/boletos/%s/baixar", nossoNumero)

	payload := struct {
		NumeroConvenio string `json:"numeroConvenio"`
	}{
		NumeroConvenio: numeroConvenio,
	}

	res, err := rest.NewClient(cred, params).Post(path, payload)
	if err != nil {
		return
	}

	if res.Code == http.StatusOK {
		err = json.Unmarshal(res.Body, &baixa)
		return
	}

	var cobErr *ErrorBaixaBoleto
	err = json.Unmarshal(res.Body, &cobErr)
	if err != nil {
		return
	}

	if cobErr.Error() != "" {
		err = cobErr
	}
	return
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
	if cobErr.Error() != "" {
		return boletos, cobErr
	}
	return boletos, nil

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

	var cobErros *bb.ErrosV1
	err = json.Unmarshal(res.Body, &cobErros)
	if err == nil && cobErros != nil && cobErros.Error() != "" {
		return boleto, cobErros
	}

	var cobErr *bb.ErrorBB
	err = json.Unmarshal(res.Body, &cobErr)
	if err != nil {
		return
	}

	if cobErr.Error() != "" {
		return boleto, cobErr
	}

	return boleto, errors.New(string(res.Body))
}

func (c client) endpoint() string {
	if c.cc.Staging {
		return endpointCobrancasStaging
	}
	return endpointCobrancasProd
}
