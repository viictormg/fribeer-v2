package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type ServiceCardHandler struct {
	serviceUsecase port.IServiceCardUsecase
}

func NewServiceCardHandler(serviceUsecase port.IServiceCardUsecase) *ServiceCardHandler {
	return &ServiceCardHandler{
		serviceUsecase: serviceUsecase,
	}
}
