package utils

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func GetClaimsToken(c echo.Context) dto.SignTokenDTO {
	var claims dto.CustomClaims

	token := strings.Split(c.Request().Header.Get("Authorization"), ".")
	rawDecodedText, _ := base64.StdEncoding.DecodeString(token[1])

	json.Unmarshal(rawDecodedText, &claims)

	return claims.SignTokenDTO
}
