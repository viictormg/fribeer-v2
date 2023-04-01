package dto

type GetUserDTO struct {
	UserName string         `json:"userName"`
	Email    string         `json:"email"`
	FullName string         `json:"fullName"`
	Avatar   string         `json:"avatar"`
	Campus   []GetCampusDTO `json:"offices"`
}
