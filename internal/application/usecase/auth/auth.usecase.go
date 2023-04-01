package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type AuthUsecase struct {
	accountService port.IAccountService
	campusService  port.ICampusService
}

func NewAuthUsecase(accountService port.IAccountService, campusService port.ICampusService) *AuthUsecase {
	return &AuthUsecase{
		accountService: accountService,
		campusService:  campusService,
	}
}
