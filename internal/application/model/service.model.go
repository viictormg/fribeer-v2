package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type Service struct {
	Name        string
	Description string
	Price       float64
	Cost        float64
	IsFrequency bool
	UnitTime    string
	Duration    int
}

func (s *Service) Validate() []string {
	err := validation.ValidateStruct(s,
		validation.Field(&s.Name, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&s.Price, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}

	return []string{}
}
