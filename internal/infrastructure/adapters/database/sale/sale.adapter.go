package adapters

import "database/sql"

type SaleAdapter struct {
	db *sql.DB
}

func NewSaleAdapter(db *sql.DB) *SaleAdapter {
	return &SaleAdapter{
		db: db,
	}
}
