package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type MeasureUnitUsecase struct {
	measureUnitService port.IMeasureUnitService
}

func NewMeasureUnitUsecase(measureUnitService port.IMeasureUnitService) *MeasureUnitUsecase {
	return &MeasureUnitUsecase{
		measureUnitService: measureUnitService,
	}
}
