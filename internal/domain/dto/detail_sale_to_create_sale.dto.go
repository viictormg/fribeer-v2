package dto

type DetailSaleToCreateSale struct {
	Details  []DetailProductToCreateSale
	Total    float64
	Error    error
	Discount float64
}
