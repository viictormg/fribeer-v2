package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type CustomerHandler struct {
	customerUsecase port.ICustomerUsecase
}

func NewCustomerHandler(customerUsecase port.ICustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}
