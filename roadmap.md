

### **1. Gerenciamento de Vagas**
- **Endpoints**
  - Criar vaga (POST /vagas)
  - Listar vagas (GET /vagas)
  - Detalhar vaga (GET /vagas/{id})
  - Atualizar vaga (PUT /vagas/{id})
  - Excluir vaga (DELETE /vagas/{id})

- **Dados**
  - Tipo de vaga: carro, moto
  - Status: disponível, ocupada, reservada
  - Localização: setor, número

---

### **2. Registro de Estacionamento**
- **Endpoints**
  - Registrar entrada (POST /estacionamento/entrada)
  - Registrar saída (POST /estacionamento/saida)
  - Listar históricos (GET /estacionamento/historico)
  - Detalhar registro (GET /estacionamento/historico/{id})

- **Dados**
  - Placa do veículo
  - Hora de entrada e saída
  - Cálculo automático de valor com base no tempo estacionado
  - Identificação da vaga usada

---

### **3. Agendamento de Vagas**
- **Endpoints**
  - Criar agendamento (POST /agendamentos)
  - Listar agendamentos (GET /agendamentos)
  - Cancelar agendamento (DELETE /agendamentos/{id})

- **Dados**
  - Data e hora desejadas
  - Tipo de vaga
  - Placa do veículo

---

### **4. Relatórios Financeiros**
- **Endpoints**
  - Relatório por período (GET /relatorios/financeiro)
  - Relatório por tipo de vaga (GET /relatorios/financeiro/vagas)

- **Dados**
  - Total arrecadado no período
  - Número de vagas ocupadas por tipo
  - Taxa média de ocupação

---

### **5. Calendário de Veículos Estacionados**
- **Endpoints**
  - Listar veículos por data (GET /calendario)
  - Detalhar veículos de um dia específico (GET /calendario/{data})

- **Dados**
  - Placa do veículo
  - Horário de permanência
  - Vaga ocupada

---

### **6. Configurações**
- **Endpoints**
  - Atualizar valores por tipo de vaga (PUT /configuracoes/valores)
  - Listar valores configurados (GET /configuracoes/valores)

---

#### **Próximos Passos**
1. **Modelagem do Banco de Dados**:
   - Definir tabelas para vagas, agendamentos, registros, veículos, e valores configuráveis.
2. **Implementação Gradual**:
   - Começar com os CRUDs básicos de vagas e registros.
3. **Autenticação**:
   - Implementar JWT para gerenciar acesso (clientes e administradores).
4. **Testes**:
   - Criar testes unitários e de integração para cada funcionalidade.
5. **Documentação**:
   - Usar Swagger ou Postman para documentar a API.
