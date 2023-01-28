package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type MeasuerUnitHandler struct {
	measureUnitUsecase port.IMeasureUnitUsecase
}

func NewMeasuerUnitHandler(measureUnitUsecase port.IMeasureUnitUsecase) *MeasuerUnitHandler {
	return &MeasuerUnitHandler{
		measureUnitUsecase: measureUnitUsecase,
	}
}
