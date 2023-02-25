package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (saleHandler *SaleHandler) CreateSaleHandler(c echo.Context) error {
	var body model.CreateSaleModel

	err := c.Bind(&body)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	creation, _ := saleHandler.saleUsecase.CreateSaleUsecase(body, "")

	response := infradto.Response{
		Success:   true,
		Data:      creation,
		Message:   constants.MessageCreate,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
