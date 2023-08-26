package adapters

import (
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleAdapter *SaleAdapter) GetSaleByIDAdapter(companyID, id string) (dto.SaleDTO, error) {
	sale := dto.SaleDTO{}
	query := `SELECT sa.id, pe.firstName, sa.saleDate, sa.total, st.name
				FROM Sale sa
				JOIN People pe ON sa.customer = pe.id
				JOIN State st ON sa.state = st.id
				WHERE sa.company = ?
				AND sa.id = ?`

	rows, err := saleAdapter.db.Query(query, companyID, id)

	if err != nil {
		return dto.SaleDTO{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&sale.ID,
			&sale.Customer,
			&sale.Date,
			&sale.Total,
			&sale.State,
		)

		if err != nil {
			defer rows.Close()
			return dto.SaleDTO{}, err
		}

	}

	return sale, nil

}
