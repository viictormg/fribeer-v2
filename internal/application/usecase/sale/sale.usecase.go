package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type SaleUsecase struct {
	saleService port.ISaleService
	serviceCard port.IServiceCardService
}

func NewSaleUsecase(saleService port.ISaleService, serviceCard port.IServiceCardService) *SaleUsecase {
	return &SaleUsecase{
		saleService: saleService,
		serviceCard: serviceCard,
	}
}
