### **📌 Roadmap do Projeto de Sistema de Estacionamento**  

---

### **1. Domínios e Estruturas de Dados (✅ COMPLETO)**
- **Domínios Implementados:**
  - Agendamento
  - Configuração
  - Pagamento
  - Plano de Cobrança
  - Registro de Estacionamento
  - Usuário
  - Vaga
  - Veículo  

- **Status:** ✅ Todos os domínios necessários para o projeto foram implementados e definidos.

---

### **2. Repositórios (✅ COMPLETO)**
- CRUD completo para todos os domínios:
  - Agendamento
  - Configuração
  - Pagamento
  - Plano de Cobrança
  - Registro de Estacionamento
  - Usuário
  - Vaga
  - Veículo  

- **Status:** ✅ Repositórios prontos com métodos de criação, leitura, atualização e exclusão.

---

### **3. Serviços (✅ COMPLETO)**
- Lógica de negócios implementada:
  - Agendamento de vagas
  - Planos de cobrança
  - Registro de estacionamento
  - Configurações de tarifas
  - Processamento de pagamentos
  - Cálculo de ocupação e receitas  

- **Status:** ✅ Serviços prontos.

---

### **4. Controllers (✅ COMPLETO - CRUD)**
- **Controllers com operações básicas:**  
  - AgendamentoControllerInterface  
  - CalendarioControllerInterface  
  - RegistroControllerInterface  
  - RelatoriosController  
  - UserControllerInterface  
  - VagaControllerInterface  
  - VeiculoControllerInterface  
  - BillingPlanControllerInterface  

- **Status:** ✅ Controladores de CRUD implementados.

---

### **5. 📍 Novos Endpoints Necessários - Funções Essenciais**
  
#### **🚗 5.1. Entrada de Veículo (Reserva ou Sem Reserva)**
- **Endpoint:** `POST /estacionamento/entrada`
- **Função:** Registrar a entrada de um veículo, associando-o a uma vaga disponível.
- **Campos Requeridos:**  
  - `placa`, `modelo`, `cor`  
  - Tipo de vaga (carro, moto)  
  - ID do plano de cobrança (opcional, se for por período pré-definido)  

- **Cenários:**  
  - Se o cliente tiver reserva: associar a vaga reservada.  
  - Se não houver reserva: buscar a primeira vaga disponível.  
  - Registrar hora de entrada e status "ocupada".  

#### **🚗 5.2. Finalizar Uso da Vaga**
- **Endpoint:** `POST /estacionamento/saida/{registroID}`
- **Função:** Finalizar o uso de uma vaga e calcular o valor total a ser pago.
- **Campos Requeridos:**  
  - `registroID` (identificador do registro de estacionamento)  
  - Hora de saída (atual)  

- **Cálculo:**  
  - Se for "por hora usada": calcular diferença entre hora de entrada e saída e aplicar tarifa.  
  - Se for plano (6h, 12h, mensal): verificar tempo e aplicar valor pré-definido.

---

#### **🚗 5.3. Pagamento**
- **Endpoint:** `POST /pagamento`
- **Função:** Registrar o pagamento após o uso da vaga ou no início, dependendo do plano.  
- **Cenários:**  
  - Para uso por hora: pagamento ao final.  
  - Para planos de período: pagamento antecipado.

---

#### **🚗 5.4. Verificar Disponibilidade de Vagas**
- **Endpoint:** `GET /vagas/disponiveis`
- **Função:** Listar todas as vagas disponíveis por tipo (carro ou moto).  

---

#### **🚗 5.5. Relatórios de Ocupação e Receitas**
- **Endpoint:**  
  - `GET /relatorios/ocupacao`  
  - `GET /relatorios/financeiro`  

- **Cenários:**  
  - Por período (diário, semanal, mensal)  
  - Por tipo de vaga  

---

### **6. Integração e Requisitos Adicionais**

#### **6.1. Middleware de Autenticação e Permissões (✅ EM PROGRESSO)**
- Implementar JWT para proteger endpoints.
- Usuário do estacionamento deve ter acesso a todas as funções.

#### **6.2. Integração de Pagamento (🔄 PENDENTE)**
- Implementar integração com gateway de pagamento online (opcional).

#### **6.3. Testes Unitários e de Integração (🔄 PENDENTE)**
- Testar todos os endpoints usando Postman ou Swagger.
  
#### **6.4. Documentação da API (🔄 PENDENTE)**
- Criar documentação utilizando Swagger.

#### **6.5. Deploy e Configuração do Ambiente de Produção (🔄 PENDENTE)**
- Configurar ambiente seguro com monitoramento.

---

### **🚀 Próximos Passos Imediatos**
1. **Implementar os endpoints críticos:**  
   - Entrada de veículo  
   - Finalização do uso e pagamento  

2. **Configurar os cálculos de tarifas nos serviços existentes.**

3. **Implementar e testar os fluxos de entrada, saída e pagamento.**  

💡 **Objetivo:** Concluir o fluxo básico de um cliente estacionando, utilizando a vaga e realizando o pagamento.