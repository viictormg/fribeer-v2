package usecase

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (saleUsecase *SaleUsecase) CreateSaleUsecase(sale model.CreateSaleModel, companyID, campus string) (dto.CreationDTO, error) {
	//Pendiente logia de descuentos en el detalle

	saleDetails, total, err := saleUsecase.saleService.GetDetailProductsService(sale.Products, companyID)

	if err != nil {
		logrus.Error(err)
		return dto.CreationDTO{}, err
	}

	ctx := context.Background()

	saleEntity := mapper.MappSaleModelToSaleEntity(sale, total, companyID, campus)
	saleCreated, trx, err := saleUsecase.saleService.CreateSaleService(saleEntity, ctx)

	if err != nil {
		logrus.Error(err)

		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err
	}

	saleDetailsToCreate := mapper.MapProductsToSaleDetail(saleDetails, saleCreated.ID)
	trx, err = saleUsecase.saleService.CreateDetailSaleService(saleDetailsToCreate, trx, companyID)

	if err != nil {
		logrus.Error(err)

		trx.Rollback()
		trx.Commit()
		return dto.CreationDTO{}, err
	}

	cardsSerivces := mapper.MapperCreateCardService(saleDetails)

	if len(cardsSerivces) > 0 {
		trx, err = saleUsecase.serviceCard.CreateServiceCardService(cardsSerivces, trx)
		if err != nil {
			logrus.Error(err)

			trx.Rollback()
			trx.Commit()
			return dto.CreationDTO{}, err
		}
	}

	fmt.Println(cardsSerivces)

	//Creacion de tarjetas de servicio

	trx.Commit()

	return saleCreated, nil
}
