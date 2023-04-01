package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type ICampusService interface {
	GetCampusService(companyID string) ([]dto.GetCampusDTO, error)
}
