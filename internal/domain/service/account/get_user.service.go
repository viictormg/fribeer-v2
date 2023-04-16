package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (accountService *AccountService) GetUserService(AccountID string) (dto.GetUserDTO, error) {
	return accountService.accountAdapter.GetUserAdapter(AccountID)
}
