package entity

type UserEntity struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Age       int8   `json:"age"`
}
