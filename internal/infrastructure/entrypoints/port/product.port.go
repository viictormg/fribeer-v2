package port

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type IProductUsecase interface {
	CreateProductUsecase(product model.Product, companyID string) (dto.CreationDTO, error)
	CreateServiceUsecase(product model.Service, companyID string) (dto.CreationDTO, error)
}
