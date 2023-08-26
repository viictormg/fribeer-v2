package mapper

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapPaymentModelToPaymentEntity(payment model.PaymentModel, total float64) entity.PaymentEntity {
	return entity.PaymentEntity{
		TypePayment:   payment.TypePayment,
		Total:         total,
		PaymentMethod: payment.PaymentMethod,
	}
}
