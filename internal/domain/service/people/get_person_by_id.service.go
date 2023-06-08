package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (peopleService *PeopleService) GetPersonByIDService(companyID, id string) (dto.GetPeopleDTO, error) {
	return peopleService.peopleAdapter.GetPeopleByIDAdapter(companyID, id)
}
