package port

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IServiceCardService interface {
	CreateServiceCardService(cards []entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error)
	GetServiceCardService(companyID string) ([]dto.GetServiceCardDTO, error)
	GetServiceCardByIDService(companyID, id string) (dto.GetServiceCardDTO, error)
	GetAllServiceCardService() ([]dto.GetServiceCardDTO, error)
}
