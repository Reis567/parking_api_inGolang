package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"time"
)

// RelatoriosService define os métodos para relatórios financeiros e operacionais
type RelatoriosService interface {
	CalcularReceita(inicio, fim time.Time) (float64, *rest_err.RestErr)
	CalcularOcupacaoAtual() (float64, *rest_err.RestErr)
	VeiculosMaisFrequentes(inicio, fim time.Time) ([]string, *rest_err.RestErr)
}
