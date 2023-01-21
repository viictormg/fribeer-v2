package dto

type UnitTimeDTO struct {
	ID          string `json:"id"`
	Frequency   string `json:"frequency"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
