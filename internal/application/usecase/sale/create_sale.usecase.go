package usecase

import (
	"context"

	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleUsecase *SaleUsecase) CreateSaleUsecase(sale model.CreateSaleModel, companyID, campus string) (dto.CreationDTO, error) {
	saleDetails, total, err := saleUsecase.saleService.GetDetailProductsService(sale.Products, companyID)
	if err != nil {
		return dto.CreationDTO{}, err
	}

	ctx := context.Background()

	saleEntity := mapper.MappSaleModelToSaleEntity(sale, total, companyID, campus)
	saleCreated, trx, err := saleUsecase.saleService.CreateSaleService(saleEntity, ctx)

	if err != nil {
		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err
	}

	saleDetails = mapper.MapProductsToSaleDetail(saleDetails, saleCreated.ID)

	trx, err = saleUsecase.saleService.CreateDetailSaleService(saleDetails, trx)

	if err != nil {
		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err

	}

	trx.Commit()

	return saleCreated, nil
}
