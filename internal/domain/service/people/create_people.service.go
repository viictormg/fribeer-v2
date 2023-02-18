package service

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (peopleService *PeopleService) CreatePeopleService(people entity.PeopleEntity, companyID string) (dto.CreationDTO, error) {
	return peopleService.peopleAdapter.CreatePeopleAdapter(people, companyID)
}
