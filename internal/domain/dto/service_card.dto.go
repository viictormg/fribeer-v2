package dto

type ServiceCardDTO struct {
	ID               string `json:"id"`
	FullNameCustomer string `json:"fullNameCustomer"`
	ServiceName      string `json:"serviceName"`
	Description      string `json:"description"`
	Customer         string `json:"customer"`
	Sale             string `json:"sale"`
	SaleDetail       string `json:"saleDetail"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
	State            string `json:"state"`
	Company          string `json:"company"`
	CreateAt         string `json:"createAt"`
	CreateBy         string `json:"createBy"`
	CpdateAt         string `json:"cpdateAt"`
	IsActive         bool   `json:"isActive"`
}
