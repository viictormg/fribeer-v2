package port

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

type IPaymentUsecase interface {
	CreatePaymentUsecase(companyID string, payment model.PaymentModel) (dto.CreationDTO, error)
}
