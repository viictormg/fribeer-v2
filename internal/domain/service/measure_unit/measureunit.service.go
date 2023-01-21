package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type MeasureUnitService struct {
	measureUnitAdater port.IUnitMeasureAdapter
}

func NewMeasureUnitService(measureUnitAdater port.IUnitMeasureAdapter) *MeasureUnitService {
	return &MeasureUnitService{
		measureUnitAdater: measureUnitAdater,
	}
}
