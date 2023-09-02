package usecase

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/application/mapper"
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (p *PaymentUsecase) CreatePaymentUsecase(companyID string, payment model.PaymentModel) (dto.CreationDTO, error) {
	var totalToPAy float64

	for _, idSale := range payment.IDsDetails {
		sale, err := p.saleService.GetSaleByIDService(companyID, idSale)
		if err != nil {
			logrus.Error(err)
			return dto.CreationDTO{}, err
		}

		totalToPAy += sale.Total
	}

	fmt.Println("Total to pay ", totalToPAy)
	//Crear encabezado

	paymentEntity := mapper.MapPaymentModelToPaymentEntity(payment, totalToPAy)
	creation, _, err := p.paymentService.CreatePaymentService(paymentEntity)

	fmt.Println(creation)
	fmt.Println(err)

	//guardar detalle

	//Cambiar el estado de las compras
	return dto.CreationDTO{}, nil

}
