package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (measureUnitUsecase *MeasureUnitUsecase) GetMeasureUnitUsecase() ([]dto.MeasureUnitDTO, error) {
	return measureUnitUsecase.measureUnitService.GetMeasureUnitService()
}
