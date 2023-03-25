package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type AuthUsecase struct {
	accountService port.IAccountService
}

func NewAuthUsecase(accountService port.IAccountService) *AuthUsecase {
	return &AuthUsecase{
		accountService: accountService,
	}
}
