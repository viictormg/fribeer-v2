package port

import (
	"context"
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type ISaleAdapter interface {
	CreateSaleAdapter(sale entity.SaleCreationEntity, ctx context.Context) (dto.CreationDTO, *sql.Tx, error)
	CreateDetailSaleAdapter(sale entity.SaleDetailEntity, trx *sql.Tx, companyID string) (*sql.Tx, error)
	GetSalesAdapter(companyID, campus string) ([]dto.SaleDTO, error)
	GetSaleByIDAdapter(companyID, id string) (dto.SaleDTO, error)
}
