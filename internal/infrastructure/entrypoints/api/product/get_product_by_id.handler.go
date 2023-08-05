package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (productHandler *ProductHandler) GetProductByIDHandler(c echo.Context) error {
	id := c.Param("id")
	payload := utils.GetClaimsToken(c)

	product, err := productHandler.productUsecase.GetProductByIDUsecase(id, payload.CompanyID)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   constants.MessageErrorFind,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	response := infradto.Response{
		Success:   true,
		Message:   constants.MessageFound,
		Data:      product,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, response)

}
