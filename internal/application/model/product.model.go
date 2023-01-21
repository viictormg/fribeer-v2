package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	MeasureUnit string  `json:"measureUnit"`
	Price       float64 `json:"price"`
	Cost        float64 `json:"cost"`
	MinStock    int     `json:"minStock"`
}

func (p *Product) Validate() []string {
	err := validation.ValidateStruct(p,
		validation.Field(&p.Name, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.MeasureUnit, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.Price, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}

	return []string{}
}
