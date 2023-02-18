package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (customerHandler *CustomerHandler) GetCustomerHandler(c echo.Context) error {

	customers, err := customerHandler.customerUsecase.GetCustomerUsecase(constants.CompanyIDTest)

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
		Data:      customers,
		Message:   constants.MessageFound,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)

}
