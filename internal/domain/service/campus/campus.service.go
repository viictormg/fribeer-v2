package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type CampusService struct {
	campusAdapter port.ICampusAdapter
}

func NewCampusService(campusAdapter port.ICampusAdapter) *CampusService {
	return &CampusService{
		campusAdapter: campusAdapter,
	}
}
