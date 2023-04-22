package dto

type DetailProductToCreateSale struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	TypeProduct   string  `json:"typeProduct"`
	Quantity      float64 `json:"quantity"`
	MeasureUnit   string  `json:"measureUnit"`
	QuantityStock int     `json:"quantityStock"`
	MinStock      int     `json:"minStock"`
	Price         float64 `json:"price"`
	IsFrequency   bool    `json:"isFrequency"`
	IsActive      bool    `json:"isActive"`
	Subtotal      float64 `json:"subtotal"`
	Discount      float64 `json:"discount"`
	Observations  string  `json:"observations"`
}
