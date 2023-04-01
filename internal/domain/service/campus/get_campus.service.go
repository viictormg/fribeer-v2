package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (campusService *CampusService) GetCampusService(companyID string) ([]dto.GetCampusDTO, error) {
	a := dto.GetCampusDTO{ID: "ssddd"}
	var b []dto.GetCampusDTO

	b = append(b, a)

	return b, nil
}
