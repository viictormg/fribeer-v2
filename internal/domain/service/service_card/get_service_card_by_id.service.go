package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (s *ServiceCardSerivice) GetServiceCardByIDService(companyID, id string) (dto.GetServiceCardDTO, error) {
	return s.serviceCardAdapter.GetServiceCardByIDAdapter(companyID, id)
}
