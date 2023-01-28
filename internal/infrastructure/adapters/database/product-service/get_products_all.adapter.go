package adapters

import (
	"errors"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

const errDBGetAllProducts = "error consulting db"

func (productAdapter *ProductAdapter) GetProducts(companyID, typeProduct string) ([]dto.ProductResponseGet, error) {
	var products []dto.ProductResponseGet

	const query = `SELECT p.id, name, p.description, tp.description AS typeProduct,
						IFNULL(mu.description, "") AS measureUnit, 
						IFNULL(p.quantityStock, 0) AS quantityStock, 
						IFNULL(MinStock, 0) AS minStock, 
						Price, Cost, IF(IsFrequency = 1, true, false) AS isFrequency,
						IFNULL(ut.frequency, "") AS unitTime, IFNULL(p.Duration, 0) AS duration,
						IF(p.isActive = 1,true, false) AS isActive
					FROM Product p
					INNER JOIN TypeProduct tp ON p.TypeProduct = tp.id
					LEFT JOIN MeasureUnit mu ON p.measureUnit = mu.id
					LEFT JOIN UnitTime ut ON ut.id = p.unitTime
					WHERE p.company  = ?`

	rows, err := productAdapter.db.Query(query)

	if err != nil {
		fmt.Println("error db products", err.Error())
		return products, errors.New(errDBGetAllProducts)
	}

	defer rows.Close()

	// for rows.Next() {
	// 	var product dto.ProductResponseGet

	// 	err = rows.Scan(
	// 	// &product
	// 	)

	// }

	return products, nil
}
