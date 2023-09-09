package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (saleService *SaleService) GetSalesService(companyID string, paramsFind dto.ParamsFindSales) ([]dto.SaleDTO, error) {
	return saleService.saleAdapter.GetSalesAdapter(companyID, paramsFind)
}
