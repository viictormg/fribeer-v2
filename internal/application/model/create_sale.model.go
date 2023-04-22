package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
)

type CreateSaleModel struct {
	CustomerID string                   `json:"customerID"`
	State      string                   `json:"state"`
	SaleDate   string                   `json:"saleDate"`
	Comments   string                   `json:"comments"`
	Products   []ProductCreateSaleModel `json:"products"`
}

type ProductCreateSaleModel struct {
	ProductID string  `json:"productID"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Quantity  float64 `json:"quantity"`
}

func (s CreateSaleModel) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.SaleDate, validation.Date("2006-01-02 15:04:05").Error(constants.DateFormatInvalid)),
		// validation.Field(&s.Products, validation.Each(validation.Date("2006-01-02 15:04:05").Error(constants.DateFormatInvalid))),
	)
}
