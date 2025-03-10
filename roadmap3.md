
---

### 1. Histórico de Estacionamento (Registros)

- **Endpoint:** `GET /registros/historico`  
- **Função:** Listar todos os registros de estacionamento (entrada e saída) com a possibilidade de filtrar por período (diário, semanal, mensal), por placa ou por status.  
- **Exemplo:**  
  - Parâmetros de query: `?dataInicio=2025-01-01&dataFim=2025-01-31&placa=ABC1234`

---

### 2. Relatório de Lotação (Ocupação)

- **Endpoint:** `GET /relatorios/lotacao`  
- **Função:** Gerar um relatório histórico da ocupação do estacionamento, podendo ser filtrado por períodos (diário, semanal, mensal) e/ou por tipo de vaga.  
- **Exemplo:**  
  - Parâmetros de query: `?periodo=diario&tipo=carro`

---

### 3. Histórico de Pagamentos

- **Endpoint:** `GET /pagamento/historico`  
- **Função:** Listar os pagamentos realizados, permitindo filtrar por data, status (por exemplo, Aberto, Concluído, Cancelado) ou método de pagamento.  
- **Exemplo:**  
  - Parâmetros de query: `?dataInicio=2025-01-01&dataFim=2025-01-31&status=Concluido`

---

### 4. Consultar Reservas Ativas/Agendamentos

- **Endpoint:** `GET /agendamentos/reservas`  
- **Função:** Listar todas as reservas (agendamentos) futuras e/ou ativas, possibilitando ao usuário visualizar suas reservas ou para a administração monitorar a demanda.  
- **Exemplo:**  
  - Parâmetros de query: `?status=confirmada`

---

### 5. Cancelamento de Reserva

- **Endpoint:** `POST /agendamentos/cancelar/{id}`  
- **Função:** Permitir o cancelamento de uma reserva/agendamento, atualizando seu status para "cancelado".  
- **Exemplo:**  
  - Passar o ID do agendamento na URL e, opcionalmente, um corpo com justificativa.

---

### 6. Veículos Atualmente Estacionados

- **Endpoint:** `GET /veiculos/ativos`  
- **Função:** Retornar uma lista de veículos que estão atualmente estacionados (por exemplo, registros com status "entrada" sem horário de saída).  
- **Exemplo:**  
  - Esse endpoint pode usar os registros de estacionamento para identificar quais veículos ainda estão ativos.

---

