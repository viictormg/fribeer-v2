package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleUsecase *SaleUsecase) CreateSaleUsecase(sale model.CreateSaleModel, companyID, campus string) (dto.CreationDTO, error) {
	//Pendiente logia de descuentos en el detalle

	responseDetail := saleUsecase.saleService.GetDetailProductsService(sale.Products, companyID)

	if responseDetail.Error != nil {
		logrus.Error(responseDetail.Error)
		return dto.CreationDTO{}, responseDetail.Error
	}

	ctx := context.Background()

	saleEntity := mapper.MappSaleModelToSaleEntity(sale, responseDetail.Total, companyID, campus)
	saleCreated, trx, err := saleUsecase.saleService.CreateSaleService(saleEntity, ctx)

	if err != nil {

		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err
	}

	saleDetailsToCreate := mapper.MapProductsToSaleDetail(responseDetail.Details, saleCreated.ID)
	trx, err = saleUsecase.saleService.CreateDetailSaleService(saleDetailsToCreate, trx, companyID)

	if err != nil {
		logrus.Error(err)

		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err
	}

	cardsSerivces := mapper.MapperCreateCardService(responseDetail.Details, saleEntity.Customer, saleCreated.ID, companyID)

	if len(cardsSerivces) > 0 {
		trx, err = saleUsecase.serviceCard.CreateServiceCardService(cardsSerivces, trx)
		if err != nil {
			logrus.Error("CARDDDDDD", err)

			trx.Rollback()
			trx.Commit()
			return dto.CreationDTO{}, err
		}
	}

	//Creacion de tarjetas de servicio

	trx.Commit()

	return saleCreated, nil
}
