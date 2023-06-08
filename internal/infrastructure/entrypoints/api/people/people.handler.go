package people

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type PeopleHandler struct {
	peopleUsecase port.IPeopleUsecase
}

func NewPeopleHandler(peopleUsecase port.IPeopleUsecase) *PeopleHandler {
	return &PeopleHandler{
		peopleUsecase: peopleUsecase,
	}
}
