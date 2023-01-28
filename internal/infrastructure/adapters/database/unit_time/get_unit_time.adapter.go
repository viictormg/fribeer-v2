package adapters

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (unitTimeAdapter *UnitTimeAdapter) GetUnitTimeAdapter() ([]dto.UnitTimeDTO, error) {
	var unitTimes []dto.UnitTimeDTO

	const query = `SELECT id, frequency, description,  IF(isActive = 1, true, false) FROM UnitTime`

	rows, err := unitTimeAdapter.db.Query(query)

	if err != nil {
		return unitTimes, err
	}

	defer rows.Close()

	for rows.Next() {
		var unitTime dto.UnitTimeDTO
		err = rows.Scan(
			&unitTime.ID,
			&unitTime.Frequency,
			&unitTime.Description,
			&unitTime.IsActive,
		)

		if err != nil {
			defer rows.Close()
			return unitTimes, err
		}
		unitTimes = append(unitTimes, unitTime)
	}
	return unitTimes, nil
}
