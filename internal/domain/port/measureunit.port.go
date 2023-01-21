package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IUnitMeasureAdapter interface {
	GetMeasureUnitAdapter() ([]dto.MeasureUnitDTO, error)
}
