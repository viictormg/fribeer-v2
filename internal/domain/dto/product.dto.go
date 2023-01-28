package dto

type ProductResponseGet struct {
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	TypeProduct   string  `json:"typeProduct"`
	MeasureUnit   string  `json:"measureUnit"`
	QuantityStock int     `json:"quantityStock"`
	MinStock      int     `json:"minStock"`
	Price         float64 `json:"price"`
	Cost          float64 `json:"cost"`
	IsFrequency   bool    `json:"isFrequency"`
	UnitTime      string  `json:"unitTime"`
	Duration      int     `json:"duration"`
	IsActive      bool    `json:"isActive"`
}
