package adapters

import "database/sql"

type ServiceCardAdapter struct {
	db *sql.DB
}

func NewServiceCardAdapter(db *sql.DB) *ServiceCardAdapter {
	return &ServiceCardAdapter{
		db: db,
	}
}
