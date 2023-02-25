package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type SaleHandler struct {
	saleUsecase port.ISaleUsecase
}

func NewSaleHandler(saleUsecase port.ISaleUsecase) *SaleHandler {
	return &SaleHandler{
		saleUsecase: saleUsecase,
	}
}
