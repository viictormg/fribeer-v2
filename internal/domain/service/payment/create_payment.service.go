package service

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (p *PaymentService) CreatePaymentService(payment entity.PaymentEntity) (dto.CreationDTO, *sql.Tx, error) {
	return dto.CreationDTO{}, nil, nil
}
