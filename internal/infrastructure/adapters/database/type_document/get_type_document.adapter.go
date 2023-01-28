package adapters

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (typeDocumentAdapter *TypeDocumentAdapter) GetTypeDocumentAdapter() ([]dto.TypeDocumentDTO, error) {
	var typeDocuments []dto.TypeDocumentDTO

	const query = `SELECT id, name, code FROM TypeDocument;`

	rows, err := typeDocumentAdapter.db.Query(query)

	if err != nil {
		return typeDocuments, err
	}

	defer rows.Close()

	for rows.Next() {
		var typeDocument dto.TypeDocumentDTO
		err = rows.Scan(
			&typeDocument.ID,
			&typeDocument.Name,
			&typeDocument.Code,
		)

		if err != nil {
			defer rows.Close()
			return typeDocuments, err
		}
		typeDocuments = append(typeDocuments, typeDocument)
	}
	return typeDocuments, nil
}
