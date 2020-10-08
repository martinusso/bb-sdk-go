# bb-sdk-go

## Endpoints

## Autenticação

### Bearer

```
// import "github.com/martinusso/bb-sdk-go"

const tokenBasic = "ZXlKcFpDST...="
tokenBearer, err := bb.OAuth{Staging: true}.Bearer(tokenBasic)
if err != nil {
	log.Fatal(err)
}
```

### Credenciais

```
// import bb "github.com/martinusso/bb-sdk-go"

cc := bb.ClientCredentials{
	Token:   tokenBearer.AccessToken, // obtido em bb.OAuth
	AppKey:  developer_application_key, // 
	Staging: true,
}
```

## Serviços

### Cria um boleto bancário

```
// import "github.com/martinusso/bb-sdk-go/cobranca"

client := cobranca.NewClient(cc)

b := cobranca.Boleto{
	// ...
}

boletoRegistrado, err := client.Registar(b)
```

### Lista boletos

```
params := cobranca.ListaBoletosParams{
	IndicadorSituacao: cobranca.BoletosEmSer,
	// Pesquisar por boletos liquidados/baixados/protestados:
	// IndicadorSituacao: cobranca.BoletosLiquidadosBaixadosProtestados,
	AgenciaBeneficiario: "452",
	ContaBeneficiario:   "123873",
	// Pesquisar por boletos de acordo com a situação
	// CodigoEstadoTituloCobranca: cobranca.SituacaoBoletoLiquidado,
}

lista, err := cobranca.NewClient(cc).ListarBoletos(params)
if err != nil {
	log.Fatal(err)
}
fmt.Print("Quantidade boletos listados: ")
fmt.Println(len(lista.Boletos))
```
