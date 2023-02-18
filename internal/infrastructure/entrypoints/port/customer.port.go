package port

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type ICustomerUsecase interface {
	CreateCustomerUsecase(customer model.CustomerCreateModel, companyID string) (dto.CreationDTO, error)
}
