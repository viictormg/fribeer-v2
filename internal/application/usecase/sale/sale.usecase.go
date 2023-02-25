package usecase

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type SaleUsecase struct {
}

func NewSaleUsecase() *SaleUsecase {
	return &SaleUsecase{}
}

func (saleUsecase *SaleUsecase) CreateSaleUsecase(sale model.CreateSaleModel, companyID string) (dto.CreationDTO, error) {
	return dto.CreationDTO{}, nil
}
