package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IUnitTimeService interface {
	GetUnitTimeService() ([]dto.UnitTimeDTO, error)
}
