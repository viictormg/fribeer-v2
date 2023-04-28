package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (s *ServiceCardHandler) GetServiceCardHandler(c echo.Context) error {
	var cards []dto.GetServiceCardDTO
	card := dto.GetServiceCardDTO{
		ID:           "test",
		ServiceName:  "prueba",
		CustomerName: "juan perez",
		Description:  "cualquier cosa",
		StartDate:    "2006-01-02 15:04:05",
		EndDate:      "2006-01-02 15:04:05",
		State:        "Terminado",
		Total:        1000,
	}

	cards = append(cards, card, card)

	response := infradto.Response{
		Success:   true,
		Timestamp: time.Now(),
		Data:      cards,
	}
	return c.JSON(http.StatusOK, response)

}
