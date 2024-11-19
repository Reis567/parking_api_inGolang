package request



type UserLogin struct {

    Email     string `json:"email" binding:"required,email"`             // Email obrigatório
    Password  string `json:"password" binding:"required,password"`       // Senha obrigatória

}
