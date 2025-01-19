package service

import (
	"meu-novo-projeto/src/configuration/rest_err"
	"time"

	"meu-novo-projeto/src/configuration/logger"


	"go.uber.org/zap"
	"meu-novo-projeto/src/model/repository"
)

// RelatoriosService define os métodos para relatórios financeiros e operacionais
type RelatoriosService interface {
	CalcularReceita(inicio, fim time.Time) (float64, *rest_err.RestErr)
	CalcularOcupacaoAtual() (float64, *rest_err.RestErr)
	VeiculosMaisFrequentes(inicio, fim time.Time) ([]string, *rest_err.RestErr)
}



type relatoriosService struct {
	registroRepo repository.RegistroEstacionamentoRepository // Atualizado para usar a interface do repositório
	vagaRepo     repository.VagaRepository                   // Atualizado para usar a interface do repositório
}

func NewRelatoriosService(
	registroRepo repository.RegistroEstacionamentoRepository,
	vagaRepo repository.VagaRepository,
) RelatoriosService {
	return &relatoriosService{
		registroRepo: registroRepo,
		vagaRepo:     vagaRepo,
	}
}


// CalcularReceita calcula a receita total em um intervalo de datas
func (s *relatoriosService) CalcularReceita(inicio, fim time.Time) (float64, *rest_err.RestErr) {
	logger.Info("Init CalcularReceita", zap.Time("inicio", inicio), zap.Time("fim", fim))

	// Buscar registros por período usando o repositório
	registros, err := s.registroRepo.FindRegistrosPorPeriodo(inicio, fim)
	if err != nil {
		logger.Error("Erro ao buscar registros por período", zap.Error(err))
		return 0, err
	}

	// Calcular a receita total
	var receitaTotal float64
	for _, registro := range registros {
		receitaTotal += registro.GetValorCobrado()
	}

	logger.Info("Receita calculada com sucesso", zap.Float64("receitaTotal", receitaTotal))
	return receitaTotal, nil
}


func (s *relatoriosService) CalcularOcupacaoAtual() (float64, *rest_err.RestErr) {
	logger.Info("Init CalcularOcupacaoAtual")

	// Contar vagas ocupadas
	vagasOcupadas, err := s.vagaRepo.CountVagasPorStatus("ocupada")
	if err != nil {
		logger.Error("Erro ao buscar vagas ocupadas", zap.Error(err))
		return 0, err
	}

	// Contar total de vagas
	totalVagas, err := s.vagaRepo.CountTotalVagas()
	if err != nil || totalVagas == 0 {
		logger.Error("Erro ao calcular total de vagas", zap.Error(err))
		return 0, rest_err.NewInternalServerError("Erro ao calcular ocupação", err)
	}

	// Calcular porcentagem
	porcentagem := (float64(vagasOcupadas) / float64(totalVagas)) * 100
	logger.Info("Ocupação calculada com sucesso", zap.Float64("ocupacao", porcentagem))
	return porcentagem, nil
}


// VeiculosMaisFrequentes retorna uma lista de placas de veículos mais frequentes em um intervalo de datas
func (s *relatoriosService) VeiculosMaisFrequentes(inicio, fim time.Time) ([]string, *rest_err.RestErr) {
	logger.Info("Init VeiculosMaisFrequentes", zap.Time("inicio", inicio), zap.Time("fim", fim))

	registros, err := s.registroRepo.FindRegistrosPorPeriodo(inicio, fim)
	if err != nil {
		logger.Error("Erro ao buscar registros por período", zap.Error(err))
		return nil, err
	}

	frequencia := make(map[string]int)
	for _, registro := range registros {
		frequencia[registro.GetPlaca()]++
	}

	// Extrair os veículos mais frequentes
	maisFrequentes := extrairMaisFrequentes(frequencia)
	logger.Info("Veículos mais frequentes calculados com sucesso", zap.Strings("veiculos", maisFrequentes))
	return maisFrequentes, nil
}

// Função auxiliar para extrair os veículos mais frequentes
func extrairMaisFrequentes(frequencia map[string]int) []string {
	// Implementar lógica de ordenação e extração, se necessário
	var veiculos []string
	for placa := range frequencia {
		veiculos = append(veiculos, placa)
	}
	return veiculos
}
