package service

import (
	"database/sql"

	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func (s *ServiceCardSerivice) CreateServiceCardService(cards []entity.ServiceCardEntity, trx *sql.Tx) (*sql.Tx, error) {

	for _, card := range cards {
		trx, err := s.serviceCardAdapter.CreateServiceCardAdapter(card, trx)
		if err != nil {
			if err != nil {
				return trx, err
			}
		}
	}

	return nil, nil
}
