package service

import (
	"context"
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (saleService *SaleService) CreateSaleService(sale entity.SaleCreationEntity, ctx context.Context) (dto.CreationDTO, *sql.Tx, error) {
	return saleService.saleAdapter.CreateSaleAdapter(sale, ctx)
}
