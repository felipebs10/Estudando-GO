package request

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"` // Uma variável que ocupa menos espaço
}
