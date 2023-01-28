package adapters

import "database/sql"

type UnitTimeAdapter struct {
	db *sql.DB
}

func NewUnitTimeAdapter(db *sql.DB) *UnitTimeAdapter {
	return &UnitTimeAdapter{
		db: db,
	}
}
