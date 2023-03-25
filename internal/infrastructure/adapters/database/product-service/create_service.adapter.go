package adapters

import (
	"errors"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const errorCreateServiceDataBase = "error create service database"

func (productAdapter *ProductAdapter) CreateServiceAdapter(product entity.ProductEntity, companyID string) (dto.CreationDTO, error) {
	id := uuid.NewString()

	var unitTime interface{}

	if product.UnitTime == "" {
		unitTime = nil
	} else {
		unitTime = product.UnitTime
	}

	query := `INSERT INTO 
				Product (id, name, description, price, cost, isFrequency, unitTime, duration, typeProduct, company) 
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := productAdapter.db.Query(
		query,
		id,
		product.Name,
		product.Description,
		product.Price,
		product.Cost,
		product.IsFrequency,
		unitTime,
		product.Duration,
		product.TypeProduct,
		companyID,
	)

	if err != nil {
		return dto.CreationDTO{}, errors.New(errorCreateServiceDataBase)
	}

	return dto.CreationDTO{
		ID: id,
	}, nil
}

// la unidad de tiempo va a ser requerida en caso de que el isFrequency este en true
