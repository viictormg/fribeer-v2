package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (saleHandler *SaleHandler) GetSalesHandler(c echo.Context) error {
	sales, err := saleHandler.saleUsecase.GetSalesUsecase(constants.CompanyIDTest, constants.CampusIDTest)
	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Message:   constants.MessageNotFound,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}
	response := infradto.Response{
		Success:   true,
		Data:      sales,
		Message:   constants.MessageFound,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, response)
}
