package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleService *SaleService) GetDetailProductsService(products []model.ProductCreateSaleModel, companyID string) ([]dto.DetailProductToCreateSale, float64, error) {
	var productsDetail []dto.DetailProductToCreateSale
	var total float64

	for _, product := range products {
		p, err := saleService.productAdapter.GetProductByIDAdapter(product.ProductID, companyID)
		if err != nil {
			return []dto.DetailProductToCreateSale{}, 0, errors.New(constants.MessageNotFound)
		}

		productDetail := dto.DetailProductToCreateSale{
			ID:          uuid.NewString(),
			ProductID:   p.ID,
			Price:       p.Price,
			Name:        p.Description,
			Quantity:    product.Quantity,
			Subtotal:    p.Price * product.Quantity,
			TypeProduct: p.TypeProduct,
			IsFrequency: p.IsFrequency,
			StartDate:   product.StartDate,
			EndDate:     product.EndDate,
			// Discount:    product.,
		}
		total += productDetail.Subtotal
		productsDetail = append(productsDetail, productDetail)
	}

	return productsDetail, total, nil
}
