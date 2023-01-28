package api

import "github.com/viictormg/fribeer-v2/internal/infrastructure/entrypoints/port"

type TypeDocumentHandler struct {
	typeDocumentUseCase port.ITypeDocumentUsecase
}

func NewTypeDocumentHandler(GetTypeDocumentUsecase port.ITypeDocumentUsecase) *TypeDocumentHandler {
	return &TypeDocumentHandler{
		typeDocumentUseCase: GetTypeDocumentUsecase,
	}
}
