package dto

type MeasureUnitDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
