package model

// User representa um usuário do sistema
type User struct {
    ID        string `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"-"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
