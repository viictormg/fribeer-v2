package usecase

import "github.com/viictormg/fribeer-v2/internal/application/port"

type TypeDocumentUseCase struct {
	typeDocumentService port.ITypeDocumentService
}

func NewTypeDocumentUseCase(typeDocumentService port.ITypeDocumentService) *TypeDocumentUseCase {
	return &TypeDocumentUseCase{
		typeDocumentService: typeDocumentService,
	}
}
