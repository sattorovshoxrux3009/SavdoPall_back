package models

type CreateAdmin struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=255"` // Adminning ismi
	LastName  string `json:"last_name" validate:"required,min=3,max=255"`  // Adminning familiyasi
	Username  string `json:"username" validate:"required,min=3,max=255"`   // Adminning foydalanuvchi nomi
	Password  string `json:"password" validate:"required,min=6"`           // Parol
}
type Login struct {
	Username string `json:"username" validate:"required,min=3,max=255"` // Adminning foydalanuvchi nomi
	Password string `json:"password" validate:"required,min=6"`         // Parol
}
