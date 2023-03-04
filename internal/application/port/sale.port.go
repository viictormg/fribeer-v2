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
	GetDetailProductsService(product []model.ProductCreateSaleModel, companyID string) ([]entity.SaleDetailEntity, float64, error)
}
