package models

// RegistrationForm is designed for
// initial registration data store.
type RegistrationForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64,containsany=!@#$%^&*()?=+-_"`
}
