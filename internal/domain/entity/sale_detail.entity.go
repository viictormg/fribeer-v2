package entity

type SaleDetailEntity struct {
	ID           string
	Product      string
	Sale         string
	Price        float64
	Quantity     float64
	Subtotal     float64
	Discount     float64
	Observations string
	Company      string
	IsActive     bool
	TypeProduct  string
}
