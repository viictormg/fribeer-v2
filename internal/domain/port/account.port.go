package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IAccountAdapter interface {
	GetSignTokenLoginAdapter(user, password string) (dto.SignTokenDTO, error)
}
