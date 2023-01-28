package adapters

import "database/sql"

type MeasureAdaterUnit struct {
	db *sql.DB
}

func NewMeasureAdaterUnit(db *sql.DB) *MeasureAdaterUnit {
	return &MeasureAdaterUnit{
		db: db,
	}
}
