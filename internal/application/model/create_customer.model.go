package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type CustomerCreateModel struct {
	FirstName      string `json:"firstName"`
	SecondName     string `json:"secondName"`
	Surname        string `json:"surname"`
	LastSurname    string `json:"lastSurname"`
	TypeDocument   string `json:"typeDocument"`
	DocumentNumber string `json:"documentNumber"`
	Birthdate      string `json:"birthdate"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Email          string `json:"email"`
}

func (c *CustomerCreateModel) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.FirstName, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&c.Surname, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&c.DocumentNumber, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&c.Phone, validation.Required.Error(constants.FieldRequiredMessage)),
	)
}
