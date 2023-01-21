package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type ProductService struct {
	productAdapter port.IProductAdapter
}

func NewProductServie(productAdapter port.IProductAdapter) *ProductService {
	return &ProductService{
		productAdapter: productAdapter,
	}
}
