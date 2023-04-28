package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (s *ServiceCardUsecase) GetServiceCardUsecase(companyID string) ([]dto.GetServiceCardDTO, error) {
	return s.serviceCerdSerice.GetServiceCardService(companyID)
}
