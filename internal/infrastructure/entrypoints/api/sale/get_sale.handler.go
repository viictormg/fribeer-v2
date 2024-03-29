package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (saleHandler *SaleHandler) GetSalesHandler(c echo.Context) error {
	payload := utils.GetClaimsToken(c)

	params := dto.ParamsFindSales{
		CustomerID: c.QueryParam("customerID"),
	}

	sales, err := saleHandler.saleUsecase.GetSalesUsecase(payload.CompanyID, params)
	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Message:   constants.MessageNotFound,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	if len(sales) == 0 {
		response := infradto.Response{
			Success:   false,
			Message:   constants.MessageNotFound,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusOK, response)
	}

	response := infradto.Response{
		Success:   true,
		Data:      sales,
		Message:   constants.MessageFound,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, response)
}
