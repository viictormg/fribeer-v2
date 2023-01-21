package port

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IProductAdapter interface {
	CreateProductAdapter(product entity.ProductEntity, companyID string) (dto.CreationDTO, error)
	CreateServiceAdapter(service entity.ProductEntity, companyID string) (dto.CreationDTO, error)
}
