package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (unitTimeUsecase *UnitTimeUseCase) GetUnitTimeUsecase() ([]dto.UnitTimeDTO, error) {
	return unitTimeUsecase.unitTimeService.GetUnitTimeService()
}
