package usecase

import (
	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (customerUsecase *CustomerUsecase) CreateCustomerUsecase(customer model.CustomerCreateModel, companyID string) (dto.CreationDTO, error) {
	customerEntity := mapper.MapCreatePeopleModelPeopleEntity(customer)

	return customerUsecase.peopleService.CreatePeopleService(customerEntity, companyID)
}
