package service

import (
	"strings"

	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (productService *ProductService) GetProductsService(typeProduct, companyID string) ([]dto.ProductResponseGet, error) {

	if typeProduct != "" {
		typeProduct = constants.TypeProducts[strings.ToUpper(typeProduct)]

		return productService.productAdapter.GetProductsByTypeAdapter(typeProduct, companyID)
	}
	return productService.productAdapter.GetProductsAllAdapter(companyID)
}
