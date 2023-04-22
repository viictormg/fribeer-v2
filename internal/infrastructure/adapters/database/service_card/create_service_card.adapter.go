package adapters

import (
	"context"
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (s *ServiceCardAdapter) CreateServiceCardAdapter(card entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error) {

	query := `INSERT INTO SaleDetail (id, serviceName, description, customer, sale, saleDetail, startDate, endDate, state, company, isActive)
		VALUES (?,?,?,?,?,?,?,?,?,?,?)`

	_, err := trx.ExecContext(context.Background(), query,
		card.ID,
		card.ServiceName,
		card.Description,
		card.Customer,
		card.Sale,
		card.SaleDetail,
		card.StartDate,
		card.EndDate,
		card.State,
		card.Company,
		true,
	)

	if err != nil {
		return trx, err
	}

	return trx, nil
}
