package dto

type UnitTimeDTO struct {
	ID          string `json:"id"`
	Frequency   string `json:"frequency"`
	Singular    string `json:"singular"`
	Plural      string `json:"plural"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
	Code        string `json:"code"`
}
