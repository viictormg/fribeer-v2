package adapters

import (
	"errors"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

const errDBGetAllProducts = "error consulting db"

func (productAdapter *ProductAdapter) GetProducts() ([]dto.ProductResponseGetAll, error) {
	var products []dto.ProductResponseGetAll

	const query = "SELECT FROM products"

	rows, err := productAdapter.db.Query(query)

	if err != nil {
		fmt.Println("error db products all", err.Error())
		return products, errors.New(errDBGetAllProducts)
	}

	defer rows.Close()

	for rows.Next() {
		// var product dto.ProductResponseGetAll

		err = rows.Scan(
		// &product
		)

	}

	return products, nil
}
