package adapters

import "database/sql"

type TypeDocumentAdapter struct {
	db *sql.DB
}

func NewTypeDocumentAdapter(db *sql.DB) *TypeDocumentAdapter {
	return &TypeDocumentAdapter{
		db: db,
	}
}
