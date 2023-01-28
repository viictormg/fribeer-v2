package port

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IProductService interface {
	CreateProductService(product entity.ProductEntity, companyID string) (dto.CreationDTO, error)
	CreateServiceService(service entity.ProductEntity, companyID string) (dto.CreationDTO, error)
	GetProductsService(typeProduct, companyID string) ([]dto.ProductResponseGet, error)
}
