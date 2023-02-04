package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (productHandler *ProductHandler) CreateProductHandler(c echo.Context) error {
	var product model.ProductModel

	err := c.Bind(&product)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	errValidation := product.Validate()

	if len(errValidation) > 0 {
		response := infradto.Response{
			Success:   false,
			Error:     errValidation,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	creation, err := productHandler.productUsecase.CreateProductUsecase(product, constants.CompanyIDTest)
	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Message:   constants.MessageErrorCreate,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	response := infradto.Response{
		Success:   true,
		Error:     errValidation,
		Data:      creation,
		Message:   constants.MessageCreate,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
