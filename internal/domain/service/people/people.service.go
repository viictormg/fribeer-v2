package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type PeopleService struct {
	peopleAdapter port.IPeopleAdapter
}

func NewPeopleService(peopleAdapter port.IPeopleAdapter) *PeopleService {
	return &PeopleService{
		peopleAdapter: peopleAdapter,
	}
}
