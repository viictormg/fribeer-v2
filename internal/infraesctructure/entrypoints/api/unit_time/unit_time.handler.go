package api

import "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/port"

type UnitTimeHandler struct {
	unitTimeUseCase port.IUnitTimeUsecase
}

func NewUnitTimeHandler(GetUnitTimeUsecase port.IUnitTimeUsecase) *UnitTimeHandler {
	return &UnitTimeHandler{
		unitTimeUseCase: GetUnitTimeUsecase,
	}
}
