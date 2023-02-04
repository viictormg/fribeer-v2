package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type CustomerUsecase struct {
	peopleService port.IPeopleService
}

func NewCustomerUsecase(peopleService port.IPeopleService) *CustomerUsecase {
	return &CustomerUsecase{
		peopleService: peopleService,
	}
}
