package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type ICampusAdapter interface {
	GetCampusAdapter(companyID string) ([]dto.GetCampusDTO, error)
}
