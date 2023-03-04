package usecase

import (
	"context"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleUsecase *SaleUsecase) CreateSaleUsecase(sale model.CreateSaleModel, companyID, campus string) (dto.CreationDTO, error) {
	_, total, err := saleUsecase.saleService.GetDetailProductsService(sale.Products, companyID)
	if err != nil {
		return dto.CreationDTO{}, err
	}

	ctx := context.Background()

	saleEntity := mapper.MappSaleModelToSaleEntity(sale, total, companyID, campus)
	saleCreated, trx, err := saleUsecase.saleService.CreateSaleService(saleEntity, ctx)

	if err != nil {
		fmt.Println(err)
		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err

	}
	trx.Commit()

	return saleCreated, nil
}
