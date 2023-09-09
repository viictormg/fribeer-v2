package adapters

import (
	"fmt"
	"regexp"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleAdapter *SaleAdapter) GetSalesAdapter(companyID string, paramsFind dto.ParamsFindSales) ([]dto.SaleDTO, error) {
	sales := []dto.SaleDTO{}
	query := fmt.Sprintf(`SELECT sa.id, pe.firstName, sa.saleDate, sa.total, st.name
				FROM Sale sa
				JOIN People pe ON sa.customer = pe.id
				JOIN State st ON sa.state = st.id
				WHERE sa.company = ? 
				%s`, GetConditionSales(paramsFind))

	fmt.Println(query)
	rows, err := saleAdapter.db.Query(query, companyID)

	if err != nil {
		logrus.Error(err)
		logrus.Info(regexp.QuoteMeta(query))
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

func GetConditionSales(params dto.ParamsFindSales) string {
	var condition string
	if params.CampusID != "" {
		condition += fmt.Sprintf("AND sa.campus = '%s' ", params.CampusID)
	}
	if params.CustomerID != "" {
		condition += fmt.Sprintf("AND sa.customer = '%s' ", params.CustomerID)
	}

	return condition
}
