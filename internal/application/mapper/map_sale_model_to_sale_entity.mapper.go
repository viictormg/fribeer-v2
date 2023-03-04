package mapper

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MappSaleModelToSaleEntity(sale model.CreateSaleModel, total float64, companyID, campus string) entity.SaleCreationEntity {
	return entity.SaleCreationEntity{
		Customer:     sale.CustomerID,
		SaleDate:     sale.SaleDate,
		Total:        total,
		Observations: sale.Comments,
		State:        constants.StateSaleConstant[sale.State],
		Company:      companyID,
		Campus:       campus,
	}
}
