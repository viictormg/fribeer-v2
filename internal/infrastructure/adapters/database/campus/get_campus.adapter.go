package adapters

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (c *CampusAdapter) GetCampusAdapter(companyID string) ([]dto.GetCampusDTO, error) {
	campus := []dto.GetCampusDTO{}

	const query = `SELECT id, name FROM Campus WHERE company = ? `

	rows, err := c.db.Query(query, companyID)

	if err != nil {
		return campus, err
	}

	defer rows.Close()

	for rows.Next() {
		var camp = dto.GetCampusDTO{}

		err = rows.Scan(
			&camp.ID,
			&camp.Name,
		)

		if err != nil {
			defer rows.Close()
			return campus, err
		}

		campus = append(campus, camp)
	}
	return campus, nil
}
