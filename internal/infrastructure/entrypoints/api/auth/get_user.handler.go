package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (authHandler *AuthHandler) GetUserHandler(c echo.Context) error {
	// fmt.Println(claims.SignTokenDTO.CompanyID)
	payload := utils.GetClaimsToken(c)

	user, err := authHandler.loginUsecase.GetUserUsecase(payload)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   "error obteniendo usuario",
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}
	response := infradto.Response{
		Success: true,
		Message: constants.MessageFoundSingular,
		Data:    user,
	}

	return c.JSON(http.StatusOK, response)
}
