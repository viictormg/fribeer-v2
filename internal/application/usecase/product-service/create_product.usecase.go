package usecase

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const (
	idProductTypeProduct = "1ac6e11c-67e5-11ed-867b-005056a61a3a"
)

func (productUsecase *ProductUsecase) CreateProductUsecase(product model.Product, companyID string) (dto.CreationDTO, error) {
	productEntity := entity.ProductEntity{
		Name:        product.Name,
		Description: product.Description,
		MeasureUnit: product.MeasureUnit,
		Price:       product.Price,
		Cost:        product.Cost,
		MinStock:    product.MinStock,
		TypeProduct: idProductTypeProduct,
		IsActive:    true,
	}

	return productUsecase.productService.CreateProductService(productEntity, companyID)
}
