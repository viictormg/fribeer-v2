package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (saleHandler *SaleHandler) GetSalesHandler(c echo.Context) error {
	payload := utils.GetClaimsToken(c)

	sales, err := saleHandler.saleUsecase.GetSalesUsecase(payload.CompanyID, constants.CampusIDTest)
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
