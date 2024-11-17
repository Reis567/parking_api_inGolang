package request



type UserRequest struct {
    FirstName string `json:"first_name" binding:"required,alpha,min=2"`  // Primeiro nome obrigatório
    LastName  string `json:"last_name" binding:"required,alpha,min=2"`   // Último nome obrigatório
    Email     string `json:"email" binding:"required,email"`             // Email obrigatório
    Password  string `json:"password" binding:"required,password"`       // Senha obrigatória
    Age       int8   `json:"age" binding:"gte=0,lte=120"`                // Idade entre 0 e 120 anos
}
