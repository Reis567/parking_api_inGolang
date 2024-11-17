package response

type UserResponse struct {
    ID        uint   `json:"id"`           // Alterado de string para uint
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Age       int8   `json:"age"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
