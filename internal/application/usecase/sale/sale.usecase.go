package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type SaleUsecase struct {
	saleService port.ISaleService
}

func NewSaleUsecase(saleService port.ISaleService) *SaleUsecase {
	return &SaleUsecase{
		saleService: saleService,
	}
}
