package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (unitTimeService *UnitTimeService) GetUnitTimeService() ([]dto.UnitTimeDTO, error) {
	return unitTimeService.unitTimeAdapter.GetUnitTimeAdapter()
}
