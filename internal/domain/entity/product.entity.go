package entity

import "time"

type ProductEntity struct {
	Id            string
	Name          string
	Description   string
	TypeProduct   string
	MeasureUnit   string
	QuantityStock int
	MinStock      int
	Price         float64
	Cost          float64
	IsFrequency   bool
	UnitTime      string
	Duration      int
	Company       string
	CreateAt      time.Time
	CreateBy      string
	UpdateAt      time.Time
	UpdateBy      string
	IsActive      bool
}
