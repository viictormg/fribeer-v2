package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (peopleService *PeopleService) GetCustomerService(companyID string) ([]dto.GetPeopleDTO, error) {
	return peopleService.peopleAdapter.GetCustomerAdapter(companyID)
}
