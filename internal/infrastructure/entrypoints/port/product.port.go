package port

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type IProductUsecase interface {
	CreateProductUsecase(product model.ProductModel, companyID string) (dto.CreationDTO, error)
	CreateServiceUsecase(product model.ServiceModel, companyID string) (dto.CreationDTO, error)
	GetProductsUsecase(typeProduct, companyID string) ([]dto.ProductResponseGet, error)
}
