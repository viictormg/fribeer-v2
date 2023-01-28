package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IUnitTimeUsecase interface {
	GetUnitTimeUsecase() ([]dto.UnitTimeDTO, error)
}
