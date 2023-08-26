package model

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type PaymentModel struct {
	TypePayment   string   `json:"idTypePayment"`
	PaymentMethod string   `json:"idPaymentMethod"`
	Observations  string   `json:"observations"`
	Reason        string   `json:"reason"`
	State         string   `json:"idState"`
	IDsDetails    []string `json:"idsDetails"`
}

func (p *PaymentModel) Validate() []string {
	err := validation.ValidateStruct(p,
		validation.Field(&p.TypePayment, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.PaymentMethod, validation.Required.Error(constants.FieldRequiredMessage)),
		validation.Field(&p.IDsDetails, validation.Required.Error(constants.FieldRequiredMessage)),
	)

	if err != nil {
		return strings.Split(err.Error(), "; ")
	}
	return []string{}

}
