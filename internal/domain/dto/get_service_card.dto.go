package dto

type GetServiceCardDTO struct {
	ID           string
	ServiceName  string
	CustomerName string
	Description  string
	StartDate    string
	EndDate      string
	State        string
	Total        float32
}
