package CRUD

type UserUpdate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}