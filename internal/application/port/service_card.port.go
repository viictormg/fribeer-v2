package port

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IServiceCardService interface {
	CreateServiceCardService(cards []entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error)
}
