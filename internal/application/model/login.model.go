package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type LoginModel struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (l *LoginModel) Validate() []string {
	err := validation.ValidateStruct(l,
		validation.Field(&l.User, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&l.Password, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}

	return []string{}
}
