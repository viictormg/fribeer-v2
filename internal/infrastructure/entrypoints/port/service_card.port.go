package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type IServiceCardUsecase interface {
	GetServiceCardUsecase(CompanyID string) ([]dto.GetServiceCardDTO, error)
	GetServiceCardByIDUsecase(CompanyID, id string) (dto.GetServiceCardDTO, error)
	JobServiceCardUsecase() error
}
