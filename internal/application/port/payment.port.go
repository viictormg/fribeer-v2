package port

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

type IPaymentService interface {
	CreatePaymentService(payment entity.PaymentEntity) (dto.CreationDTO, *sql.Tx, error)
}
