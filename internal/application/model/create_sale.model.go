package model

type CreateSaleModel struct {
	CustomerID string                   `json:"customerID"`
	State      string                   `json:"state"`
	SaleDate   string                   `json:"saleDate"`
	Comments   string                   `json:"comments"`
	Products   []ProductCreateSaleModel `json:"products"`
}

type ProductCreateSaleModel struct {
	ProductID   string  `json:"productID"`
	InitialDate string  `json:"initialDate"`
	EndDate     string  `json:"endDate"`
	Quantity    float64 `json:"quantity"`
}
