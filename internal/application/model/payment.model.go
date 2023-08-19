package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type PaymentModel struct {
	TypePayment   string   `json:"typePayment"`
	PaymentMethod string   `json:"idPaymentMethod"`
	Total         float64  `json:"total"`
	Observations  string   `json:"observations"`
	Reason        string   `json:"reason"`
	State         string   `json:"idState"`
	IDsDetails    []string `json:"idsDetails"`
}

func (p *PaymentModel) Validate() []string {
	err := validation.ValidateStruct(p,
		validation.Field(&p.TypePayment, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.PaymentMethod, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.Total, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.Total, validation.Min(0.0).Error(constants.FieldMustBeGreaterThanZero)),
		validation.Field(&p.IDsDetails, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}
	return []string{}

}
