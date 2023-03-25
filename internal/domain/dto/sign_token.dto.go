package dto

import "github.com/golang-jwt/jwt/v4"

type SignTokenDTO struct {
	AccountID string `json:"accountID"`
	CompanyID string `json:"companyID"`
	PeopleID  string `json:"PeopleID"`
}

type CustomClaims struct {
	SignTokenDTO SignTokenDTO `json:"data,omitempty"`
	jwt.RegisteredClaims
}
