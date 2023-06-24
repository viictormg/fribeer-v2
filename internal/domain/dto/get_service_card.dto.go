package dto

type GetServiceCardDTO struct {
	ID           string `json:"id"`
	ServiceName  string `json:"serviceName"`
	SaleID       string `json:"saleID"`
	CustomerName string `json:"customerName"`
	Description  string `json:"description"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	State        string `json:"state"`
	StateCode    string `json:"stateCode"`
	Expiration   int64  `json:"expiration"`
	CustomerID   string `json:"customerID"`
}
