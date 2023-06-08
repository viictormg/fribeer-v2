package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (s *ServiceCardHandler) GetServiceCardByIDHandler(c echo.Context) error {
	payload := utils.GetClaimsToken(c)

	id := c.Param("id")

	card, err := s.serviceUsecase.GetServiceCardByIDUsecase(payload.CompanyID, id)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Timestamp: time.Now(),
			Message:   constants.MessageErrorFind,
		}
		return c.JSON(http.StatusConflict, response)
	}
	response := infradto.Response{
		Success:   true,
		Timestamp: time.Now(),
		Data:      card,
	}
	return c.JSON(http.StatusOK, response)
}
