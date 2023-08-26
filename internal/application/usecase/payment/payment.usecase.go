package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type PaymentUsecase struct {
	paymentService port.IPaymentService
	saleService    port.ISaleService
}

func NewPaymentUsecase(paymentService port.IPaymentService, saleService port.ISaleService) *PaymentUsecase {
	return &PaymentUsecase{
		paymentService: paymentService,
		saleService:    saleService,
	}
}
