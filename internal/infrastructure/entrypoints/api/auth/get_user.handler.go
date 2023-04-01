package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (authHandler *AuthHandler) GetUserHandler(c echo.Context) error {
	// fmt.Println(claims.SignTokenDTO.CompanyID)

	token := strings.Split(c.Request().Header.Get("Authorization"), ".")
	rawDecodedText, _ := base64.StdEncoding.DecodeString(token[1])

	var claims dto.CustomClaims

	json.Unmarshal(rawDecodedText, &claims)

	fmt.Println(claims.SignTokenDTO.CompanyID)

	return nil
}
