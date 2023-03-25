package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (authHandler *AuthHandler) LoginHandler(c echo.Context) error {
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

	token, err := authHandler.loginUsecase.LoginUsecase(login.User, login.Password)

	fmt.Println(token)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   "usuario no encontrado",
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	response := infradto.Response{
		Success:   true,
		Message:   "login exitoso",
		Timestamp: time.Now(),
		Data:      token,
	}
	return c.JSON(http.StatusOK, response)

}
