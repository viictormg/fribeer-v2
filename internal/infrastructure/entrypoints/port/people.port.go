package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IPeopleUsecase interface {
	GetPersonByIDUsecase(companyID, id string) (dto.GetPeopleDTO, error)
}
