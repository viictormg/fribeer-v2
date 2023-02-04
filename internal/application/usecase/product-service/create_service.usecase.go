package usecase

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const (
	idProductService = "1aba8311-67e5-11ed-867b-005056a61a3a"
)

func (ProductUsecase *ProductUsecase) CreateServiceUsecase(service model.ServiceModel, companyID string) (dto.CreationDTO, error) {
	serviceEntity := entity.ProductEntity{
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
		Cost:        service.Cost,
		IsFrequency: service.IsFrequency,
		UnitTime:    service.IDUnitTime,
		Duration:    service.Duration,
		TypeProduct: idProductService,
	}
	return ProductUsecase.productService.CreateServiceService(serviceEntity, companyID)
}
