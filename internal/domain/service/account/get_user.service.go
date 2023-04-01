package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (accountService *AccountService) GetUserService(AccountID, PeopleID string) (dto.GetUserDTO, error) {

	return dto.GetUserDTO{UserName: "service"}, nil
}
