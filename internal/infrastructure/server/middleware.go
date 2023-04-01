package server

// func MiddlewareFribeer(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authHeader := c.Request().Header.Get("Authorization")
// 		if authHeader == "" {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token de autenticación no proporcionado"})
// 		}
// 		tokenString := authHeader[len("Bearer "):]
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("error al verificar la firma del token")
// 			}
// 			return "todoesculpadelabareta", nil
// 		})
// 		if err != nil || !token.Valid {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token de autenticación no válido"})
// 		}
// 		return next(c)
// 	}
// }
