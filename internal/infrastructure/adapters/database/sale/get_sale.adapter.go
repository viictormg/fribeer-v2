package adapters

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleAdapter *SaleAdapter) GetSalesAdapter(companyID, campus string) ([]dto.SaleDTO, error) {
	sales := []dto.SaleDTO{}
	query := `SELECT sa.id, pe.firstName, sa.saleDate, sa.total, st.name
				FROM Sale sa
				JOIN People pe ON sa.customer = pe.id
				JOIN State st ON sa.state = st.id
				WHERE sa.company = ?
				AND sa.campus = ?`

	rows, err := saleAdapter.db.Query(query, companyID, campus)

	if err != nil {
		return []dto.SaleDTO{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var sale dto.SaleDTO
		err = rows.Scan(
			&sale.ID,
			&sale.Customer,
			&sale.Date,
			&sale.Total,
			&sale.State,
		)

		if err != nil {
			defer rows.Close()
			return []dto.SaleDTO{}, err
		}

		sales = append(sales, sale)
	}

	return sales, nil

}
