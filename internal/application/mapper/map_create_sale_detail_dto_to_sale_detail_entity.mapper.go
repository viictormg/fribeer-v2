package mapper

import (
	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapProductsToSaleDetail(producs []entity.SaleDetailEntity, saleID string) []entity.SaleDetailEntity {
	var details []entity.SaleDetailEntity
	id := uuid.NewString()

	for _, d := range producs {
		detail := entity.SaleDetailEntity{
			ID:           id,
			Product:      d.Product,
			Sale:         saleID,
			Price:        d.Price,
			Quantity:     d.Quantity,
			Subtotal:     d.Subtotal,
			Discount:     d.Discount,
			Observations: d.Observations,
			Company:      d.Company,
			IsActive:     true,
		}
		details = append(details, detail)
	}
	return details
}
