package dto

type GetServiceCardDTO struct {
	ID           string  `json:"id,omitempty"`
	ServiceName  string  `json:"serviceName,omitempty"`
	CustomerName string  `json:"customerName,omitempty"`
	Description  string  `json:"description,omitempty"`
	StartDate    string  `json:"startDate,omitempty"`
	EndDate      string  `json:"endDate,omitempty"`
	State        string  `json:"state,omitempty"`
	Total        float32 `json:"total,omitempty"`
}
