package sale

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
)

func (saleHandler *SaleHandler) CreateSaleHandler(c echo.Context) error {

	response := infradto.Response{
		Success:   true,
		Data:      dto.CreationDTO{ID: "ddd"},
		Message:   constants.MessageCreate,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusCreated, response)
}
