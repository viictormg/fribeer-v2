package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type ServiceModel struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Cost        float64 `json:""`
	IsFrequency bool    `json:"isFrequency"`
	IDUnitTime  string  `json:"idUnitTime"`
	Duration    int     `json:"duration"`
}

// {"name":"Servicio masajes","description":"","price":50000,"cost":0,"isFrequency":true,"duration":2,"idUnitTime":"0e574690-51c8-11ed-860e-005056a61a3a"}
func (s *ServiceModel) Validate() []string {
	err := validation.ValidateStruct(s,
		validation.Field(&s.Name, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&s.Price, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}

	return []string{}
}
