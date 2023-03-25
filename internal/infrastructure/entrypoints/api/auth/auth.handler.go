package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type AuthHandler struct {
	loginUsecase port.IAuthUsecase
}

func NewAuthHandler(loginUsecase port.IAuthUsecase) *AuthHandler {
	return &AuthHandler{
		loginUsecase: loginUsecase,
	}
}
