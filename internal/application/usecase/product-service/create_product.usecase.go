package usecase

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const ()

func (productUsecase *ProductUsecase) CreateProductUsecase(product model.ProductModel, companyID string) (dto.CreationDTO, error) {
	productEntity := entity.ProductEntity{
		Name:        product.Name,
		Description: product.Description,
		MeasureUnit: product.MeasureUnit,
		Price:       product.Price,
		Cost:        product.Cost,
		MinStock:    product.MinStock,
		TypeProduct: constants.IDProductTypeProduct,
		IsActive:    true,
	}

	return productUsecase.productService.CreateProductService(productEntity, companyID)
}
