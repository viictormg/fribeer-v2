package mapper

import (
	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapProductsToSaleDetail(producs []dto.DetailProductToCreateSale, saleID string) []entity.SaleDetailEntity {
	var details []entity.SaleDetailEntity

	for _, d := range producs {
		detail := entity.SaleDetailEntity{
			ID:           uuid.NewString(),
			Product:      d.ID,
			Sale:         saleID,
			Price:        d.Price,
			Quantity:     d.Quantity,
			Subtotal:     d.Subtotal,
			Discount:     d.Discount,
			Observations: d.Observations,
			IsActive:     true,
		}
		details = append(details, detail)
	}
	return details
}
