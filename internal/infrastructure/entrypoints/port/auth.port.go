package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IAuthUsecase interface {
	LoginUsecase(user, password string) (string, error)
	GetUserUsecase(payload dto.SignTokenDTO) (dto.GetUserDTO, error)
}
