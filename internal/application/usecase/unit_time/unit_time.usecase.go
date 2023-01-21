package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type UnitTimeUseCase struct {
	unitTimeService port.IUnitTimeService
}

func NewUnitTimeUseCase(unitTimeService port.IUnitTimeService) *UnitTimeUseCase {
	return &UnitTimeUseCase{
		unitTimeService: unitTimeService,
	}
}
