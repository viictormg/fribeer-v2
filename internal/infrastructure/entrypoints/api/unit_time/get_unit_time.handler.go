package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"

	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (UnitTimeHandler *UnitTimeHandler) GetUnitTimeHandler(c echo.Context) error {
	unitTimes, err := UnitTimeHandler.unitTimeUseCase.GetUnitTimeUsecase()

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := infradto.Response{
		Success:   false,
		Error:     []string{constants.ErrorDecodeBody},
		Timestamp: time.Now(),
		Data:      unitTimes,
	}
	return c.JSON(http.StatusOK, response)

}
