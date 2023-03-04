package service

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleService *SaleService) CreateDetailSaleService(saleDetails []entity.SaleDetailEntity, trx *sql.Tx) (*sql.Tx, error) {
	for _, detail := range saleDetails {

		trx, err := saleService.saleAdapter.CreateDetailSaleAdapter(detail, trx)
		if err != nil {
			return trx, err
		}
	}

	return trx, nil
}
