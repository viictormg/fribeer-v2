package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IMeasureUnitService interface {
	GetMeasureUnitService() ([]dto.MeasureUnitDTO, error)
}
