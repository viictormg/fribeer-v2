package adapters

import "database/sql"

type PeopleAdapter struct {
	db *sql.DB
}

func NewPeopleAdapter(db *sql.DB) *PeopleAdapter {
	return &PeopleAdapter{
		db: db,
	}
}
