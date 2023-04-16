package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (authUsecase *AuthUsecase) GetUserUsecase(payload dto.SignTokenDTO) (dto.GetUserDTO, error) {
	user, err := authUsecase.accountService.GetUserService(payload.AccountID)

	if err != nil {
		return dto.GetUserDTO{}, err
	}

	c, err := authUsecase.campusService.GetCampusService(payload.CompanyID)

	if err != nil {
		return dto.GetUserDTO{}, err
	}
	user.Campus = c
	return user, nil
}
