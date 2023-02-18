package dto

type GetPeopleDTO struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Document string `json:"document"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
