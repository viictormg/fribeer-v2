package dto

type GetServiceCardDTO struct {
	ID           string  `json:"id"`
	ServiceName  string  `json:"serviceName"`
	CustomerName string  `json:"customerName"`
	Description  string  `json:"description"`
	StartDate    string  `json:"startDate"`
	EndDate      string  `json:"endDate"`
	State        string  `json:"state"`
	Total        float32 `json:"total"`
}
