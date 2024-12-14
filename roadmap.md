### **Roadmap Detalhado para API de Estacionamento**

---

### **1. Gerenciamento de Vagas**
- **Objetivo**: Gerenciar as vagas disponíveis no estacionamento, incluindo criação, atualização, exclusão e status em tempo real.
- **Endpoints**:
  - **Criar Vaga** (POST /vagas)
  - **Listar Todas as Vagas** (GET /vagas)
  - **Detalhar Vaga por ID** (GET /vagas/{id})
  - **Atualizar Dados da Vaga** (PUT /vagas/{id})
  - **Excluir Vaga** (DELETE /vagas/{id})
  - **Listar Vagas Disponíveis** (GET /vagas/disponiveis)
  - **Listar Vagas Ocupadas** (GET /vagas/ocupadas)

- **Dados**:
  - **Tipo de vaga**: carro, moto.
  - **Status da vaga**: disponível, ocupada, reservada, manutenção.
  - **Localização**: setor, número.
  - **Serial da Vaga**: identificador único baseado em regras específicas.

- **Métricas**:
  - Total de vagas no estacionamento.
  - Porcentagem de ocupação atual.
  - Histórico de ocupação nos últimos dias/semanas.

---

### **2. Registro de Estacionamento**
- **Objetivo**: Registrar a entrada e saída de veículos, calcular valores com base no tempo estacionado e gerenciar o histórico de uso.
- **Endpoints**:
  - **Registrar Entrada** (POST /estacionamento/entrada)
  - **Registrar Saída** (POST /estacionamento/saida)
  - **Listar Históricos de Estacionamento** (GET /estacionamento/historico)
  - **Detalhar Registro por ID** (GET /estacionamento/historico/{id})
  - **Estatísticas por Veículo** (GET /estacionamento/veiculos/{placa}/historico)

- **Dados**:
  - Placa do veículo.
  - Hora de entrada e saída.
  - Tipo de vaga usada.
  - Identificador da vaga (serial).
  - Valor total cobrado (calculado automaticamente).
  - Tempo total estacionado.

- **Funcionalidades Adicionais**:
  - Cálculo de valores configuráveis por tipo de vaga (carro/moto).
  - Histórico de uso detalhado por veículo (placa).

---

### **3. Agendamento de Vagas**
- **Objetivo**: Permitir que os usuários reservem vagas com antecedência para datas específicas.
- **Endpoints**:
  - **Criar Agendamento** (POST /agendamentos)
  - **Listar Agendamentos Ativos** (GET /agendamentos)
  - **Cancelar Agendamento** (DELETE /agendamentos/{id})
  - **Verificar Disponibilidade** (GET /agendamentos/disponibilidade)

- **Dados**:
  - Data e hora do agendamento.
  - Tipo de vaga desejado.
  - Placa do veículo.
  - Confirmação automática do status da vaga para "reservada".

---

### **4. Relatórios Financeiros e Operacionais**
- **Objetivo**: Gerar insights financeiros e operacionais para melhorar a gestão do estacionamento.
- **Endpoints**:
  - **Relatório Financeiro por Período** (GET /relatorios/financeiro?inicio={data}&fim={data})
  - **Relatório de Ocupação por Tipo de Vaga** (GET /relatorios/ocupacao)
  - **Relatório Detalhado de Veículos Estacionados** (GET /relatorios/veiculos)

- **Dados**:
  - Total arrecadado por período.
  - Porcentagem de ocupação média por tipo de vaga.
  - Lista de veículos que mais estacionaram.
  - Total de reservas e quantas foram efetivamente utilizadas.

---

### **5. Calendário e Visualização de Ocupação**
- **Objetivo**: Exibir a ocupação do estacionamento em formato de calendário e permitir visualizações rápidas por dia.
- **Endpoints**:
  - **Listar Veículos Estacionados por Data** (GET /calendario)
  - **Detalhar Ocupação em uma Data Específica** (GET /calendario/{data})

- **Dados**:
  - Lista de veículos (placa).
  - Tempo de permanência.
  - Status das vagas (disponíveis/ocupadas/reservadas).

- **Funcionalidades Adicionais**:
  - Visualização gráfica com porcentagem de ocupação em tempo real.
  - Histórico visual de ocupação por semanas/meses.

---

### **6. Configurações**
- **Objetivo**: Permitir ajustes dinâmicos de valores e outros parâmetros da API.
- **Endpoints**:
  - **Atualizar Valores por Tipo de Vaga** (PUT /configuracoes/valores)
  - **Listar Configurações de Valores** (GET /configuracoes/valores)

- **Dados**:
  - Valores por hora para carros e motos.
  - Tempo mínimo cobrado e valores adicionais.

---

### **7. Estatísticas Avançadas**
- **Objetivo**: Fornecer informações para análises detalhadas sobre o desempenho do estacionamento.
- **Endpoints**:
  - **Porcentagem de Ocupação Atual** (GET /estatisticas/ocupacao)
  - **Relatório de Ocupação nos Últimos Dias** (GET /estatisticas/ocupacao?dias=7)
  - **Estatísticas por Tipo de Vaga** (GET /estatisticas/vagas)

- **Dados**:
  - Total de vagas disponíveis atualmente.
  - Comparação de ocupação por tipo de vaga (carro/moto).
  - Dias mais movimentados.

---

### **8. Autenticação e Permissões**
- **Objetivo**: Proteger endpoints críticos e permitir controle de acesso baseado em perfis.
- **Endpoints**:
  - **Registrar Usuário** (POST /usuarios)
  - **Login** (POST /login)
  - **Gerenciamento de Permissões** (Admin Only)
  - JWT obrigatório para todos os endpoints protegidos.

- **Funcionalidades Adicionais**:
  - Perfis: Administrador, Operador, Cliente.
  - Controle de acesso granular.

---

### **Próximos Passos**
1. **Implementação do CRUD de Vagas e Registro de Estacionamento**.
2. **Desenvolvimento de funcionalidades de Agendamento**.
3. **Criação de Relatórios Financeiros e Operacionais**.
4. **Integração com APIs externas (opcional)**:
   - Consulta de placas de veículos.
   - Pagamento online integrado.
5. **Testes Unitários e de Integração**.
6. **Documentação da API**:
   - Utilizar Swagger ou Postman para mapear todos os endpoints.
7. **Deploy**:
   - Configurar ambiente seguro para produção com monitoramento.

