package cobranca

import (
	"fmt"
	"strconv"
)

// CodigoAceite identifica se o boleto de cobrança foi aceito (reconhecimento da dívida pelo Pagador), sendo:
// "A" - ACEITE
// "N" - NAO ACEITE
type CodigoAceite string

const (
	Aceite    CodigoAceite = "A"
	NaoAceite CodigoAceite = "N"
)

// CodigoModalidade identifica a característica dos boletos dentro das modalidades de cobrança existentes no BB, sendo:
// 1 - SIMPLES
// 4 - VINCULADA
type CodigoModalidade int

const (
	Simples   CodigoModalidade = 1
	Vinculada CodigoModalidade = 4
)

func (m CodigoModalidade) String() string {
	if m == Simples || m == Vinculada {
		return strconv.Itoa(int(m))
	}
	return ""
}

// BoletoVencido indica se o boleto está vencido "S" ou não "N"
type BoletoVencido string

const (
	Vencido    BoletoVencido = "S"
	NaoVencido BoletoVencido = "N"
)

// IndicadorSituacao indica a situação do boleto, sendo:
// A - boletos em ser
// B - boletos liquidados/baixados/protestados
type IndicadorSituacao string

const (
	BoletosEmSer                         IndicadorSituacao = "A"
	BoletosLiquidadosBaixadosProtestados IndicadorSituacao = "B"
)

// TipoDesconto indica cmo o desconto será concedido, sendo:
// 0 - SEM DESCONTO
// 1 - VLR FIXO ATE A DATA INFORMADA
// 2 - PERCENTUAL ATE A DATA INFORMADA
type TipoDesconto int

const (
	SemDesconto TipoDesconto = iota
	ValorFixo
	Percentual
)

// TipoJurosMora indica o código utilizado pela FEBRABAN para identificar o tipo de taxa de juros, sendo:
// 0 - DISPENSAR
// 1 - VALOR DIA ATRASO
// 2 - TAXA MENSAL
// 3 - ISENTO
type TipoJurosMora int

const (
	Dispensar TipoJurosMora = iota
	ValorDiaAtraso
	TaxaMensal
	Isento
)

// TipoMulta indica o código para identificação do tipo de multa para o Título de Cobrança, sendo:
// 0 - Sem multa
// 1 - Valor da Multa
// 2 - Percentual da Multa
type TipoMulta int

const (
	SemMulta TipoMulta = iota
	ValorDaMulta
	PercentualDaMulta
)

// TipoRegistro indica o tipo do registro do pagador, sendo:
// 1 - Pessoa Física
// 2 - Pessoa Jurídica
type TipoRegistro int

const (
	PessoaFisica TipoRegistro = iota + 1
	PessoaJuridica
)

// SituacaoBoleto indica o código da situação atual do boleto
type SituacaoBoleto int

const (
	SituacaoBoletoNormal                          SituacaoBoleto = iota + 1 // 01 NORMAL
	SituacaoBoletoMovimentoCartorio                                         // 02 MOVIMENTO CARTORIO
	SituacaoBoletoEmCartorio                                                // 03  EM CARTORIO
	SituacaoBoletoTituloComOcorrenciaCartorio                               // 04  TITULO COM OCORRENCIA DE CARTORIO
	SituacaoBoletoProtestadoEletronico                                      // 05  PROTESTADO ELETRONICO
	SituacaoBoletoLiquidado                                                 // 06  LIQUIDADO
	SituacaoBoletoBaixado                                                   // 07  BAIXADO
	SituacaoBoletoTituloComPendenciaCartorio                                // 08  TITULO COM PENDENCIA DE CARTORIO
	SituacaoBoletoTituloProtestadoManual                                    // 09  TITULO PROTESTADO MANUAL
	SituacaoBoletoTituloBaixadoPagoEmCartorio                               // 10  TITULO BAIXADO/PAGO EM CARTORIO
	SituacaoBoletoTituloLiquidadoProtestado                                 // 11  TITULO LIQUIDADO/PROTESTADO
	SituacaoBoletoTituloLiquidadoPagoEmCartorio                             // 12  TITULO LIQUID/PGCRTO
	SituacaoBoletoTituloProtestadoAguardandoBaixa                           // 13  TITULO PROTESTADO AGUARDANDO BAIXA
	SituacaoBoletoTituloEmLiquidacao                                        // 14  TITULO EM LIQUIDACAO
	SituacaoBoletoTituloAgendado                                            // 15  TITULO AGENDADO
	SituacaoBoletoTituloCreditado                                           // 16  TITULO CREDITADO
	SituacaoBoletoPagoEmCheque                                              // 17  PAGO EM CHEQUE - AGUARD.LIQUIDACAO
	SituacaoBoletoPagoParcialmente                                          // 18  PAGO PARCIALMENTE
	SituacaoBoletoPagoParcialmenteCreditado                                 // 19  PAGO PARCIALMENTE CREDITADO
)

func (s SituacaoBoleto) String() string {
	if s > 0 {
		return fmt.Sprintf("%02d", s)
	}
	return ""
}
