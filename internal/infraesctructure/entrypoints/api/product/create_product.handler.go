package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infraesctructure/entrypoints/api"
)

func (productHandler *ProductHandler) CreateProductHandler(c echo.Context) error {
	var product model.Product

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

	companyID := "4f76fe5e-527f-11ed-867b-005056a61a3a"

	creation, err := productHandler.productUsecase.CreateProductUsecase(product, companyID)
	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     errValidation,
			Data:      creation,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	response := infradto.Response{
		Success:   true,
		Error:     errValidation,
		Data:      creation,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
