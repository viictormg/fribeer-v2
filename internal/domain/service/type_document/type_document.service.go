package service

import "github.com/viictormg/fribeer-v2/internal/domain/port"

type TypeDocumentService struct {
	typeDocumentAdapter port.ITypeDocumentAdapter
}

func NewTypeDocumentService(typeDocumentAdapter port.ITypeDocumentAdapter) *TypeDocumentService {
	return &TypeDocumentService{
		typeDocumentAdapter: typeDocumentAdapter,
	}
}
