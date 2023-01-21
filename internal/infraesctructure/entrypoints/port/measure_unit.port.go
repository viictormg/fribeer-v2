package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IMeasureUnitUsecase interface {
	GetMeasureUnitUsecase() ([]dto.MeasureUnitDTO, error)
}
