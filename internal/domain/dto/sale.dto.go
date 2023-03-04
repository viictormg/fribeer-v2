package dto

type SaleDTO struct {
	ID       string  `json:"id"`
	Customer string  `json:"customer"`
	Date     string  `json:"date"`
	Total    float64 `json:"total"`
	State    string  `json:"state"`
}
