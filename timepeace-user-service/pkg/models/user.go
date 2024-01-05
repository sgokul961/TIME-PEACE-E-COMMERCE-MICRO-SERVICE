package models

type UserLoginCheck struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	PhNo     string `json:"phno"`
	Password string `json:"password"`
}
