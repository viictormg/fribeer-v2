package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleAdapter *SaleAdapter) CreateSaleAdapter(sale entity.SaleCreationEntity, ctx context.Context) (dto.CreationDTO, *sql.Tx, error) {
	id := uuid.NewString()

	trx, err := saleAdapter.db.BeginTx(ctx, nil)
	if err != nil {
		return dto.CreationDTO{}, trx, err
	}

	query := `INSERT INTO Sale (id, customer, saleDate, total, observations, state, campus, company, isActive) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = trx.ExecContext(ctx, query,
		id,
		sale.Customer,
		sale.SaleDate,
		sale.Total,
		sale.Observations,
		sale.State,
		sale.Campus,
		sale.Company,
		true,
	)

	if err != nil {
		return dto.CreationDTO{}, trx, err
	}

	// trx.Rollback()
	// defer trx.Commit()

	return dto.CreationDTO{ID: id}, trx, err

}
