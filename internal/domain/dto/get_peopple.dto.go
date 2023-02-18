package dto

type GetPeopleDTO struct {
	ID           string `json:"id"`
	NameComplete string `json:"nameComplete"`
	Document     string `json:"document"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
}
