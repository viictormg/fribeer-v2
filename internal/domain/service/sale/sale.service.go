package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type SaleService struct {
	saleAdapter    port.ISaleAdapter
	productAdapter port.IProductAdapter
}

func NewSaleService(saleAdapter port.ISaleAdapter, productAdapter port.IProductAdapter) *SaleService {
	return &SaleService{
		saleAdapter:    saleAdapter,
		productAdapter: productAdapter,
	}
}
