package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (authHandler *AuthHandler) Login(c echo.Context) error {
	var login model.LoginModel

	err := c.Bind(&login)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	errValidation := login.Validate()

	if len(errValidation) > 0 {
		response := infradto.Response{
			Success:   false,
			Message:   errValidation,
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return nil
}
