package service

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (productService *ProductService) CreateServiceService(service entity.ProductEntity, companyID string) (dto.CreationDTO, error) {
	return productService.productAdapter.CreateServiceAdapter(service, companyID)
}
