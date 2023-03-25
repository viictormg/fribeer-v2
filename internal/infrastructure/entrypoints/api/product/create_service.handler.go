package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (productHandler *ProductHandler) CreateServiceHandler(c echo.Context) error {
	var service model.ServiceModel

	err := c.Bind(&service)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   constants.ErrorDecodeBody,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	errValidation := service.Validate()
	if len(errValidation) > 0 {
		response := infradto.Response{
			Success:   false,
			Message:   constants.MessageErrorCreate,
			Error:     errValidation,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	creation, err := productHandler.productUsecase.CreateServiceUsecase(service, constants.CompanyIDTest)
	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   constants.MessageErrorCreate,
			Error:     []string{err.Error()},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)

	}

	response := infradto.Response{
		Success:   true,
		Message:   constants.MessageCreate,
		Data:      creation,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
