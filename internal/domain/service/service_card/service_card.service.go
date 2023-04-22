package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type ServiceCardSerivice struct {
	serviceCardAdapter port.IServiceCardAdapter
}

func NewServiceCardSerivice(serviceCardAdapter port.IServiceCardAdapter) *ServiceCardSerivice {
	return &ServiceCardSerivice{
		serviceCardAdapter: serviceCardAdapter,
	}
}
