package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (campusService *CampusService) GetCampusService(companyID string) ([]dto.GetCampusDTO, error) {
	return campusService.campusAdapter.GetCampusAdapter(companyID)
}
