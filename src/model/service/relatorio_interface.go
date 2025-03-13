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
	CalcularLotacaoHistorica(periodo, tipo string) (float64, *rest_err.RestErr)
	CalcularTempoMedioPermanencia(inicio, fim time.Time) (float64, *rest_err.RestErr)
}



type relatoriosService struct {
	registroRepo repository.RegistroEstacionamentoRepository // Atualizado para usar a interface do repositório
	vagaRepo     repository.VagaRepository        
	relatorioRepo     repository.RelatoriosRepository  
}

func NewRelatoriosService(
	registroRepo repository.RegistroEstacionamentoRepository,
	vagaRepo repository.VagaRepository,
	relatorioRepo repository.RelatoriosRepository,
) RelatoriosService {
	return &relatoriosService{
		registroRepo: registroRepo,
		vagaRepo:     vagaRepo,
		relatorioRepo:     relatorioRepo,
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



func (s *relatoriosService) CalcularLotacaoHistorica(periodo, tipo string) (float64, *rest_err.RestErr) {
	now := time.Now()
	var inicio time.Time

	// Determinar a data de início com base no período solicitado
	switch periodo {
	case "diario":
		// Início do dia atual
		inicio = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	case "semanal":
		// 7 dias atrás
		inicio = now.AddDate(0, 0, -7)
	case "mensal":
		// 30 dias atrás
		inicio = now.AddDate(0, 0, -30)
	default:
		return 0, rest_err.NewBadRequestError("Periodo inválido. Use 'diario', 'semanal' ou 'mensal'.")
	}
	fim := now

	// Obter os registros de estacionamento no período
	registros, err := s.registroRepo.FindRegistrosPorPeriodo(inicio, fim)
	if err != nil {
		return 0, err
	}

	// Criar um conjunto de IDs únicos de vagas que foram utilizadas no período
	uniqueVagas := make(map[uint]struct{})
	for _, reg := range registros {
		uniqueVagas[reg.GetVagaID()] = struct{}{}
	}

	var totalVagas int
	var usadas int

	if tipo != "" {
		// Se o filtro por tipo for informado, buscar todas as vagas e filtrar pelo tipo
		vagas, repoErr := s.vagaRepo.FindAllVagas()
		if repoErr != nil {
			return 0, repoErr
		}

		// Filtrar as vagas que possuem o tipo desejado
		filteredVagas := make(map[uint]struct{})
		for _, vaga := range vagas {
			if vaga.GetTipo() == tipo {
				filteredVagas[vaga.GetID()] = struct{}{}
			}
		}
		totalVagas = len(filteredVagas)
		// Contar quantas das vagas utilizadas estão entre as vagas filtradas
		for id := range uniqueVagas {
			if _, ok := filteredVagas[id]; ok {
				usadas++
			}
		}
	} else {
		// Sem filtro de tipo, contar todas as vagas disponíveis
		var countErr *rest_err.RestErr
		totalVagas, countErr = s.vagaRepo.CountTotalVagas()
		if countErr != nil {
			return 0, countErr
		}
		usadas = len(uniqueVagas)
	}

	if totalVagas == 0 {
		return 0, rest_err.NewNotFoundError("Nenhuma vaga encontrada para o filtro informado")
	}

	// Calcular a taxa de ocupação (porcentagem de vagas utilizadas)
	ocupacao := (float64(usadas) / float64(totalVagas)) * 100.0
	return ocupacao, nil
}


func (s *relatoriosService) CalcularTempoMedioPermanencia(inicio, fim time.Time) (float64, *rest_err.RestErr) {
	logger.Info("Iniciando cálculo do tempo médio de permanência", zap.Time("inicio", inicio), zap.Time("fim", fim))

	tempoMedio, err := s.relatorioRepo.CalcularTempoMedioPermanencia(inicio, fim)
	if err != nil {
		logger.Error("Erro ao calcular tempo médio de permanência", zap.Error(err))
		return 0, err
	}

	logger.Info("Tempo médio de permanência calculado com sucesso", zap.Float64("tempoMedio", tempoMedio))
	return tempoMedio, nil
}
