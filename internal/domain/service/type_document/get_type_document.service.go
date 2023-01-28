package service

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (typeDocumentService *TypeDocumentService) GetTypeDocumentService() ([]dto.TypeDocumentDTO, error) {
	return typeDocumentService.typeDocumentAdapter.GetTypeDocumentAdapter()
}
