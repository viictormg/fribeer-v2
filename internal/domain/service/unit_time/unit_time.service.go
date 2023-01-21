package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type UnitTimeService struct {
	unitTimeAdapter port.IUnitTimeAdapter
}

func NewUnitTimeService(unitTimeAdapter port.IUnitTimeAdapter) *UnitTimeService {
	return &UnitTimeService{
		unitTimeAdapter: unitTimeAdapter,
	}
}
