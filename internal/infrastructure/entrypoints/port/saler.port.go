package port

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type ISaleUsecase interface {
	CreateSaleUsecase(sale model.CreateSaleModel, companyID, campus string) (dto.CreationDTO, error)
	GetSalesUsecase(companyID string, paramsFind dto.ParamsFindSales) ([]dto.SaleDTO, error)
}
