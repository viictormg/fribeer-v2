package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (customerHandler *CustomerHandler) CreateCustomerHandler(c echo.Context) error {
	var customer model.CustomerCreateModel

	err := c.Bind(&customer)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	err = customer.Validate()

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     strings.Split(err.Error(), "; "),
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	creation, err := customerHandler.customerUsecase.CreateCustomerUsecase(customer)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)

	}

	response := infradto.Response{
		Success:   true,
		Data:      creation,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)

}
