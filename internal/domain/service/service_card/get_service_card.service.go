package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (s *ServiceCardSerivice) GetServiceCardService(companyID string) ([]dto.GetServiceCardDTO, error) {
	return s.serviceCardAdapter.GetServiceCardAdapter(companyID)
}
