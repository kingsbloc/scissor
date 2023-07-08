package dto

import "net/http"

type AddUserDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

func (a *AddUserDto) Bind(r *http.Request) error {
	validationErr := validate.Struct(a)
	return validationErr
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (a *LoginDto) Bind(r *http.Request) error {
	validationErr := validate.Struct(a)
	return validationErr
}
