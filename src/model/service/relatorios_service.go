package service

import (
	"meu-novo-projeto/src/configuration/logger"
	"meu-novo-projeto/src/configuration/rest_err"
	"meu-novo-projeto/src/model"
	"go.uber.org/zap"
	"time"
)


func (s *relatoriosService) CalcularReceita(inicio, fim time.Time) (float64, *rest_err.RestErr) {
    registros, err := s.registroRepo.FindRegistrosPorPeriodo(inicio, fim)
    if err != nil {
        return 0, err
    }

    var receitaTotal float64
    for _, registro := range registros {
        receitaTotal += registro.GetValorCobrado()
    }

    return receitaTotal, nil
}

