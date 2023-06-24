package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
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
	err = body.Validate()

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	payload := utils.GetClaimsToken(c)

	creation, err := saleHandler.saleUsecase.CreateSaleUsecase(body, payload.CompanyID, constants.CampusIDTest)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{"pendiente definicion de error"},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	response := infradto.Response{
		Success:   true,
		Data:      creation,
		Message:   constants.MessageCreate,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
