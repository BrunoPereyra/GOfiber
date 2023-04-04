package validator

import "github.com/go-playground/validator"

type UserFind struct {
	NameUser string `json:"nameUser" validate:"required,max=12"`
}

func (u *UserFind) ValidateUserFind() error {
	validate := validator.New()
	return validate.Struct(u)
}
