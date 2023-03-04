package service

import (
	"database/sql"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleService *SaleService) CreateDetailSaleService(saleDetails []entity.SaleDetailEntity, idSale string, trx *sql.Tx) (*sql.Tx, error) {
	for i, detail := range saleDetails {
		fmt.Println("DDDD", i)

		detail.Sale = idSale
		trx, err := saleService.saleAdapter.CreateDetailSaleAdapter(detail, trx)
		if err != nil {
			return trx, err
		}
	}

	return trx, nil
}
