package entity

type SaleCreationEntity struct {
	ID           string
	Customer     string
	SaleDate     string
	Total        float64
	Observations string
	State        string
	Campus       string
	Company      string
	CreateBy     string
	IsActive     bool
}
