package adapters

import "database/sql"

type AccountAdapter struct {
	db *sql.DB
}

func NewAccountAdapter(db *sql.DB) *AccountAdapter {
	return &AccountAdapter{
		db: db,
	}
}
