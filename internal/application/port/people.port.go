package port

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IPeopleService interface {
	CreatePeopleService(people entity.PeopleEntity, companyID string) (dto.CreationDTO, error)
	GetCustomerService(companyID string) ([]dto.GetPeopleDTO, error)
	GetPersonByIDService(companyID, id string) (dto.GetPeopleDTO, error)
}
