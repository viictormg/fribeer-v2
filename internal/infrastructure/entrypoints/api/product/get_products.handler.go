package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (productHandler *ProductHandler) GetProductsHandler(c echo.Context) error {
	typeProduct := c.QueryParam("type")
	payload := utils.GetClaimsToken(c)

	products, err := productHandler.productUsecase.GetProductsUsecase(typeProduct, payload.CompanyID)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   constants.MessageErrorFind,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	if len(products) == 0 {
		response := infradto.Response{
			Success:   true,
			Message:   constants.MessageNotFound,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusOK, response)
	}

	response := infradto.Response{
		Success:   true,
		Message:   constants.MessageFound,
		Data:      products,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, response)

}
