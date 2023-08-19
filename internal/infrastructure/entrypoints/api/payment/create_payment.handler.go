package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (p *PaymentHandler) CreatePaymentHandler(c echo.Context) error {
	var payment model.PaymentModel

	err := c.Bind(&payment)

	if err != nil {
		fmt.Println(err)
		response := infradto.Response{
			Success:   false,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	errValidation := payment.Validate()

	if len(errValidation) > 0 {
		response := infradto.Response{
			Success:   false,
			Error:     errValidation,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)

	}

	return nil
}
