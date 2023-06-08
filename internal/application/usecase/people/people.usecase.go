package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type PeopleUsecase struct {
	peopleService port.IPeopleService
}

func NewPeopleUsecase(peopleService port.IPeopleService) *PeopleUsecase {
	return &PeopleUsecase{
		peopleService: peopleService,
	}
}
