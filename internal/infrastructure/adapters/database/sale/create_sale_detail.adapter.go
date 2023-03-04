package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleAdapter *SaleAdapter) CreateDetailSaleAdapter(sale entity.SaleDetailEntity, trx *sql.Tx) (*sql.Tx, error) {
	id := uuid.NewString()

	query := `INSERT INTO SaleDetail (id, product, sale, price, quantity, subtotal, discount, company, isActive)
					VALUES (?,?,?,?,?,?,?,?,?)`

	_, err := trx.ExecContext(context.Background(), query,
		id,
		sale.Product,
		sale.Sale,
		sale.Price,
		sale.Quantity,
		sale.Subtotal,
		sale.Discount,
		sale.Company,
		true,
	)

	if err != nil {
		return trx, err
	}

	return trx, nil
}
