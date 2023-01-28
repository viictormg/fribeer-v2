package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (ProductUsecase *ProductUsecase) GetProductsUsecase(typeProduct, companyID string) ([]dto.ProductResponseGet, error) {
	return ProductUsecase.productService.GetProductsService(typeProduct, companyID)
}
