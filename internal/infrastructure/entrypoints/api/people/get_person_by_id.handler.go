package people

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	infradto "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/api"
	"github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/utils"
)

func (people *PeopleHandler) GetPersonByIDHandler(c echo.Context) error {
	id := c.Param("id")

	payload := utils.GetClaimsToken(c)

	person, err := people.peopleUsecase.GetPersonByIDUsecase(payload.CompanyID, id)

	if err != nil {
		response := infradto.Response{
			Success:   false,
			Message:   constants.ErrorDecodeBody,
			Error:     []string{constants.ErrorDecodeBody},
			Timestamp: time.Now(),
		}
		return c.JSON(http.StatusConflict, response)
	}

	response := infradto.Response{
		Success:   true,
		Message:   constants.MessageFound,
		Data:      person,
		Timestamp: time.Now(),
	}
	return c.JSON(http.StatusOK, response)

}
