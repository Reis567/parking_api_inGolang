package request

// UserRequest representa um usuário do sistema
type UserRequest struct {
    ID        string `json:"id" binding:"required,uuid4"`                 // O ID deve ser um UUID e é obrigatório
    FirstName string `json:"first_name" binding:"required,alpha,min=2"`  // Primeiro nome é obrigatório, só letras, mínimo 2 caracteres
    LastName  string `json:"last_name" binding:"required,alpha,min=2"`   // Último nome é obrigatório, só letras, mínimo 2 caracteres
    Email     string `json:"email" binding:"required,email"`             // Email é obrigatório e deve ser válido
    Password  string `json:"password" binding:"required,min=6,regexp=^(?=.*[A-Z])(?=.*[!@#$&*]).*$"` // Senha é obrigatória, min. 6 caracteres, uma maiúscula, um símbolo
    Age       int8   `json:"age" binding:"gte=0,lte=120"`                // Idade deve ser entre 0 e 120 anos
    CreatedAt string `json:"created_at" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"` // Data no formato ISO8601, opcional
    UpdatedAt string `json:"updated_at" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"` // Data no formato ISO8601, opcional
}
