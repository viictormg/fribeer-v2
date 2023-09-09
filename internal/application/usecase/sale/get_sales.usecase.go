package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (saleUsecase *SaleUsecase) GetSalesUsecase(companyID string, paramsFind dto.ParamsFindSales) ([]dto.SaleDTO, error) {
	return saleUsecase.saleService.GetSalesService(companyID, paramsFind)
}
