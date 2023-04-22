package port

import (
	"context"
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type ISaleService interface {
	CreateSaleService(sale entity.SaleCreationEntity, ctx context.Context) (dto.CreationDTO, *sql.Tx, error)
	GetDetailProductsService(product []model.ProductCreateSaleModel, companyID string) ([]dto.DetailProductToCreateSale, float64, error)
	CreateDetailSaleService(saleDetails []entity.SaleDetailEntity, trx *sql.Tx, companyID string) (*sql.Tx, error)
	GetSalesService(companyID, campus string) ([]dto.SaleDTO, error)
}
