package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (accountService *AccountService) GetSignTokenLoginService(user, password string) (dto.SignTokenDTO, error) {
	return accountService.accountAdapter.GetSignTokenLoginAdapter(user, password)
}
