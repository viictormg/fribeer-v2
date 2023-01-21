package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type UnitMeasureService struct {
	unitmeasureAdater port.IUnitMeasureAdapter
}

func NewMeasureUnitService(unitmeasureAdater port.IUnitMeasureAdapter) *UnitMeasureService {
	return &UnitMeasureService{
		unitmeasureAdater: unitmeasureAdater,
	}
}
