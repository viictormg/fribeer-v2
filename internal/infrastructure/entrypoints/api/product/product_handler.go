package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type ProductHandler struct {
	productUsecase port.IProductUsecase
}

func NewProductHandler(productUsecase port.IProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}
