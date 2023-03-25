package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IAccountService interface {
	GetSignTokenLoginService(user, password string) (dto.SignTokenDTO, error)
}
