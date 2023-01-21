package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type ProductUsecase struct {
	productService port.IProductService
}

func NewProductUsecase(productService port.IProductService) *ProductUsecase {
	return &ProductUsecase{
		productService: productService,
	}
}
