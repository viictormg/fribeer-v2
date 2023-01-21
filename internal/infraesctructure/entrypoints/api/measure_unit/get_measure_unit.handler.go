package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"

	infradto "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api"
)

func (measuerUnitHandler *MeasuerUnitHandler) GetMeasureUnitHandler(c echo.Context) error {
	measureUnits, err := measuerUnitHandler.measureUnitUsecase.GetMeasureUnitUsecase()

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
		Data:      measureUnits,
	}
	return c.JSON(http.StatusOK, response)

}