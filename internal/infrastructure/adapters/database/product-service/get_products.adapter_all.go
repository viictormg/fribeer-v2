package adapters

import (
	"errors"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

const errDBGetAllProducts = "error consulting db"

func (productAdapter *ProductAdapter) GetProductsAllAdapter(companyID string) ([]dto.ProductResponseGet, error) {
	var products []dto.ProductResponseGet

	const query = `SELECT p.id, name, p.description, tp.description AS typeProduct,
						IFNULL(mu.description, "") AS measureUnit, 
						IFNULL(p.quantityStock, 0) AS quantityStock, 
						IFNULL(MinStock, 0) AS minStock, 
						Price, Cost, IF(IsFrequency = 1, true, false) AS isFrequency,
						IFNULL(ut.frequency, "") AS unitTime, IFNULL(p.Duration, 0) AS duration,
						IF(p.isActive = 1,true, false) AS isActive,
						IFNULL(ut.code, "") AS unitTimeCode
					FROM Product p
					INNER JOIN TypeProduct tp ON p.TypeProduct = tp.id
					LEFT JOIN MeasureUnit mu ON p.measureUnit = mu.id
					LEFT JOIN UnitTime ut ON ut.id = p.unitTime
					WHERE p.company = ?`

	rows, err := productAdapter.db.Query(query, companyID)

	if err != nil {
		return products, errors.New(errDBGetAllProducts)
	}

	defer rows.Close()

	for rows.Next() {
		var product dto.ProductResponseGet

		err = rows.Scan(
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
			&product.IsActive,
			&product.UnitTimeCode,
		)

		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

// RENAME TABLE account TO Account;
// RENAME TABLE accountingrecords TO Accountingrecords;
// RENAME TABLE campus TO Aampus;
// RENAME TABLE company TO Company;
// RENAME TABLE config TO Config;
// RENAME TABLE historystock TO Historystock;
// RENAME TABLE jobs TO Jobs;
// RENAME TABLE measureunit TO Measureunit;
// RENAME TABLE module TO Module;
// RENAME TABLE modulesforcompany TO ModulesForCompany;
// RENAME TABLE notifications TO Notifications;
// RENAME TABLE payment TO Payment;
// RENAME TABLE paymentdetail TO PaymentDetail;
// RENAME TABLE paymentmethod TO PaymentMethod;
// RENAME TABLE people TO People;
// RENAME TABLE permission TO Permission;
// RENAME TABLE rol TO Rol;
// RENAME TABLE sale TO Sale;
// RENAME TABLE saledetail TO Saledetail;
// RENAME TABLE servicecard TO ServiceCard;
// RENAME TABLE state TO State;
// RENAME TABLE typecompany TO TypeCompany;
// RENAME TABLE typedocument TO TypeDocument;
// RENAME TABLE typepayment TO TypePayment;
// RENAME TABLE typepeople TO TypePeople;
// RENAME TABLE typeproduct TO Typeproduct;
// RENAME TABLE unittime TO UnitTime;
