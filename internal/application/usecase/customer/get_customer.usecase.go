package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (customerUsecase *CustomerUsecase) GetCustomerUsecase(companyID string) ([]dto.GetPeopleDTO, error) {
	return customerUsecase.peopleService.GetCustomerService(companyID)
}
