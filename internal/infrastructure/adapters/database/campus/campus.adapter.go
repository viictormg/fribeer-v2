package adapters

import "database/sql"

type CampusAdapter struct {
	db *sql.DB
}

func NewCampusAdapter(db *sql.DB) *CampusAdapter {
	return &CampusAdapter{
		db: db,
	}
}
