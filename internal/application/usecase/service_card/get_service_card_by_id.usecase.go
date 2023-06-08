package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (s *ServiceCardUsecase) GetServiceCardByIDUsecase(CompanyID, id string) (dto.GetServiceCardDTO, error) {
	return s.serviceCerdService.GetServiceCardByIDService(CompanyID, id)
}
