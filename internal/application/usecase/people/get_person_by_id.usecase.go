package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (p *PeopleUsecase) GetPersonByIDUsecase(companyID, id string) (dto.GetPeopleDTO, error) {
	return p.peopleService.GetPersonByIDService(companyID, id)
}
