package payment

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type PaymentHandler struct {
	paymentUsecase port.IPaymentUsecase
}

func NewPaymentHandler(paymentUsecase port.IPaymentUsecase) *PaymentHandler {
	return &PaymentHandler{
		paymentUsecase: paymentUsecase,
	}
}
