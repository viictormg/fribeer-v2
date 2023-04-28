package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type ServiceCardUsecase struct {
	serviceCerdSerice port.IServiceCardService
}

func NewServiceUsecase(serviceCerdSerice port.IServiceCardService) *ServiceCardUsecase {
	return &ServiceCardUsecase{
		serviceCerdSerice: serviceCerdSerice,
	}
}
