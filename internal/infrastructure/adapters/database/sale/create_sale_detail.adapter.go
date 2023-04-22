package adapters

import (
	"context"
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleAdapter *SaleAdapter) CreateDetailSaleAdapter(sale entity.SaleDetailEntity, trx *sql.Tx, companyID string) (*sql.Tx, error) {

	query := `INSERT INTO SaleDetail (id, product, sale, price, quantity, subtotal, discount, company, isActive)
					VALUES (?,?,?,?,?,?,?,?,?)`

	_, err := trx.ExecContext(context.Background(), query,
		sale.ID,
		sale.Product,
		sale.Sale,
		sale.Price,
		sale.Quantity,
		sale.Subtotal,
		sale.Discount,
		companyID,
		true,
	)

	if err != nil {
		return trx, err
	}

	return trx, nil
}
