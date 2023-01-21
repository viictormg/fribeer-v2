package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IUnitTimeAdapter interface {
	GetUnitTimeAdapter() ([]dto.UnitTimeDTO, error)
}
