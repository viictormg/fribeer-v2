package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (s *ServiceCardHandler) GetServiceCardHandler(c echo.Context) error {

	payload := utils.GetClaimsToken(c)

	cards, err := s.serviceUsecase.GetServiceCardUsecase(payload.CompanyID)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Timestamp: time.Now(),
			Message:   constants.MessageErrorFind,
			Data:      cards,
		}
		return c.JSON(http.StatusConflict, response)

	}

	response := infradto.Response{
		Success:   true,
		Timestamp: time.Now(),
		Data:      cards,
	}
	return c.JSON(http.StatusOK, response)

}
