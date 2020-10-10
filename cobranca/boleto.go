package cobranca

import "fmt"

// NossoNumero é o número de identificação do boleto, com 20 dígitos, que deverá ser formatado da seguinte forma:
// "000" + (número do convênio com 7 dígitos) + (10 algarismos - se necessário, completar com zeros à esquerda)
func NossoNumero(contrato, numero string) string {
	return fmt.Sprintf("000%s%010s", contrato, numero)
}
