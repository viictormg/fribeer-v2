package port

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IServiceCardAdapter interface {
	CreateServiceCardAdapter(card entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error)
}
