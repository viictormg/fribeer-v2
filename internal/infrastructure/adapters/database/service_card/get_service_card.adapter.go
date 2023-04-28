package adapters

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (s *ServiceCardAdapter) GetServiceCardAdapter(companyID string) ([]dto.GetServiceCardDTO, error) {
	cards := []dto.GetServiceCardDTO{}

	query := `SELECT id, serviceName, description 
				FROM ServiceCard WHERE company = ?;`

	rows, err := s.db.Query(query, companyID)

	if err != nil {
		return cards, err
	}

	defer rows.Close()

	for rows.Next() {
		var card = dto.GetServiceCardDTO{}

		err = rows.Scan(
			&card.ID,
			&card.ServiceName,
			&card.Description,
		)

		if err != nil {
			defer rows.Close()
			return cards, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
