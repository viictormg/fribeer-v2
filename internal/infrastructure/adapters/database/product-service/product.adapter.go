package adapters

import "database/sql"

type ProductAdapter struct {
	db *sql.DB
}

func NewProductAdapter(db *sql.DB) *ProductAdapter {
	return &ProductAdapter{
		db: db,
	}
}
