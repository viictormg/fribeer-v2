package adapters

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const errorCreateProductDataBase = "error create product database"

func (productAdapter *ProductAdapter) CreateProductAdapter(product entity.ProductEntity, companyID string) (dto.CreationDTO, error) {
	id := uuid.NewString()

	query := `INSERT INTO 
				Product (id, name, description, measureUnit, price, cost, minStock, typeProduct, company, isActive) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := productAdapter.db.Query(
		query,
		id,
		product.Name,
		product.Description,
		product.MeasureUnit,
		product.Price,
		product.Cost,
		product.MinStock,
		product.TypeProduct,
		companyID,
		product.IsActive,
	)

	if err != nil {
		fmt.Println(errorCreateProductDataBase, err)
		return dto.CreationDTO{}, errors.New(errorCreateProductDataBase)
	}

	return dto.CreationDTO{
		ID: id,
	}, nil
}
