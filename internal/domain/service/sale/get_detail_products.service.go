package service

import (
	"errors"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleService *SaleService) GetDetailProductsService(products []model.ProductCreateSaleModel, companyID string) ([]entity.SaleDetailEntity, float64, error) {
	var productsDetail []entity.SaleDetailEntity
	var total float64

	for _, product := range products {
		p, err := saleService.productAdapter.GetProductByIDAdapter(product.ProductID, companyID)
		if err != nil {
			fmt.Println(err)
			return []entity.SaleDetailEntity{}, 0, errors.New(constants.MessageNotFound)
		}

		productDetail := entity.SaleDetailEntity{
			Product:  p.ID,
			Price:    p.Price,
			Quantity: product.Quantity,
			Subtotal: p.Price * product.Quantity,
			Company:  companyID,
		}
		total += productDetail.Subtotal
		productsDetail = append(productsDetail, productDetail)
	}

	return productsDetail, total, nil
}
