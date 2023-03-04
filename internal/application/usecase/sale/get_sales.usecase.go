package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (saleUsecase *SaleUsecase) GetSalesUsecase(companyID, campus string) ([]dto.SaleDTO, error) {
	return saleUsecase.saleService.GetSalesService(companyID, campus)
}
