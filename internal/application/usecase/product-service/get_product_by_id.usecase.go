package usecase

import (
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (ProductUsecase *ProductUsecase) GetProductByIDUsecase(id, companyID string) (dto.ProductResponseGet, error) {
	fmt.Println(id)
	return ProductUsecase.productService.GetProductByIDService(id, companyID)
}
