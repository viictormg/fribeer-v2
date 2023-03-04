package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (productService *ProductService) GetProductByIDService(companyID string, id string) (dto.ProductResponseGet, error) {
	return productService.productAdapter.GetProductByIDAdapter(companyID, id)
}
