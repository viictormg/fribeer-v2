package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (measureUnitService *MeasureUnitService) GetMeasureUnitService() ([]dto.MeasureUnitDTO, error) {
	return measureUnitService.measureUnitAdater.GetMeasureUnitAdapter()
}
