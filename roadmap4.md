 Rotas adicionais recomendadas:
1. Administração do estacionamento
GET /estacionamento/status → Obtém o status geral do estacionamento (total de vagas, ocupação atual, receita do dia, etc.).
POST /estacionamento/config → Atualiza configurações gerais (horários de funcionamento, preços, regras, etc.).
GET /estacionamento/config → Obtém as configurações atuais.
2. Gestão de clientes
GET /users/me → Retorna os detalhes do usuário logado (útil para o app do cliente).
GET /users/:id/historico → Retorna o histórico de estacionamento do usuário.
GET /users/:id/veiculos → Lista os veículos cadastrados pelo usuário.
3. Controle de estacionamento
GET /vagas/status → Retorna a disponibilidade em tempo real das vagas (exemplo: por setor ou andar).
GET /vagas/setor/:setorID → Retorna a ocupação de um setor específico.
GET /vagas/reservadas → Lista vagas que estão reservadas.
4. Gestão de pagamentos e faturamento
GET /pagamento/:id → Consulta um pagamento específico.
GET /pagamento/pending → Lista pagamentos pendentes.
GET /pagamento/recibos/:userId → Retorna os recibos do usuário.
5. Notificações e suporte
POST /suporte/abrir-chamado → Usuário pode abrir um chamado para suporte (problema com vaga, cobrança, etc.).
GET /suporte/chamados → Lista chamados abertos e resolvidos.
POST /notificacoes/enviar → Admin pode enviar notificações para usuários.
6. Segurança e monitoramento
GET /seguranca/cameras → Lista câmeras ativas no estacionamento.
POST /seguranca/alerta → Registra um alerta de segurança (exemplo: veículo suspeito, incidente, etc.).