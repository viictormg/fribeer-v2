package port

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IServiceCardAdapter interface {
	CreateServiceCardAdapter(card entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error)
	GetServiceCardAdapter(companyID string) ([]dto.GetServiceCardDTO, error)
	GeAllServicesCardAdapter() ([]dto.GetServiceCardDTO, error)
	GetServiceCardByIDAdapter(companyID, id string) (dto.GetServiceCardDTO, error)
}
