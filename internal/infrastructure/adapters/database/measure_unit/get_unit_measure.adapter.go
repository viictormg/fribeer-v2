package adapters

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (measureAdaterUnit *MeasureAdaterUnit) GetMeasureUnitAdapter() ([]dto.MeasureUnitDTO, error) {
	var measureUnits []dto.MeasureUnitDTO

	const query = `SELECT id, description,  IF(isActive = 1, true, false) FROM MeasureUnit`

	rows, err := measureAdaterUnit.db.Query(query)

	if err != nil {
		return measureUnits, err
	}

	defer rows.Close()

	for rows.Next() {
		var measureUnit dto.MeasureUnitDTO
		err = rows.Scan(
			&measureUnit.ID,
			&measureUnit.Description,
			&measureUnit.IsActive,
		)

		if err != nil {
			defer rows.Close()
			return measureUnits, err
		}
		measureUnits = append(measureUnits, measureUnit)
	}
	return measureUnits, nil
}
