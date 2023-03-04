package adapters

import (
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (productAdapter *ProductAdapter) GetProductByIDAdapter(id, companyID string) (dto.ProductResponseGet, error) {
	var product dto.ProductResponseGet

	fmt.Println(companyID, id)

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
					WHERE p.company  = ?
					AND p.id = ?`

	err := productAdapter.db.QueryRow(query, companyID, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.TypeProduct,
		&product.MeasureUnit,
		&product.QuantityStock,
		&product.MinStock,
		&product.Price,
		&product.Cost,
		&product.IsFrequency,
		&product.UnitTime,
		&product.Duration,
		&product.IsActive)
	if err != nil {
		fmt.Println(err)
		return product, err
	}

	return product, nil
}
