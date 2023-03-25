package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type AccountService struct {
	accountAdapter port.IAccountAdapter
}

func NewAccountService(accountAdapter port.IAccountAdapter) *AccountService {
	return &AccountService{
		accountAdapter: accountAdapter,
	}
}
