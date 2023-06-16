package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleService *SaleService) GetDetailProductsService(products []model.ProductCreateSaleModel, companyID string) dto.DetailSaleToCreateSale {
	var productsDetail []dto.DetailProductToCreateSale
	var total, discount float64

	for _, product := range products {
		p, err := saleService.productAdapter.GetProductByIDAdapter(product.ProductID, companyID)
		if err != nil {
			return dto.DetailSaleToCreateSale{
				Details:  []dto.DetailProductToCreateSale{},
				Total:    0,
				Discount: 0,
				Error:    errors.New(constants.MessageNotFound),
			}
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
			Discount:    product.Discount,
		}
		total += productDetail.Subtotal
		discount += productDetail.Discount

		productsDetail = append(productsDetail, productDetail)
	}

	return dto.DetailSaleToCreateSale{
		Details:  productsDetail,
		Total:    total - discount,
		Discount: discount,
		Error:    nil,
	}
}
