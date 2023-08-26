package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (saleService *SaleService) GetSaleByIDService(companyID, id string) (dto.SaleDTO, error) {
	return saleService.saleAdapter.GetSaleByIDAdapter(companyID, id)
}
