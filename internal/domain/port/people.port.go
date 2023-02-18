package port

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IPeopleAdapter interface {
	CreatePeopleAdapter(people entity.PeopleEntity, companyID string) (dto.CreationDTO, error)
	GetCustomerAdapter(companyID string) ([]dto.GetPeopleDTO, error)
}
