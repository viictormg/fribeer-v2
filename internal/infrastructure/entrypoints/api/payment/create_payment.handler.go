package payment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (p *PaymentHandler) CreatePaymentHandler(c echo.Context) error {
	var payment model.PaymentModel

	payload := utils.GetClaimsToken(c)

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

	creation, err := p.paymentUsecase.CreatePaymentUsecase(payload.CompanyID, payment)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Error:     []string{err.Error()},
			Message:   constants.MessageErrorCreate,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
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
