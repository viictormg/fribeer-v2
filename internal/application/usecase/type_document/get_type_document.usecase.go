package usecase

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (typeDocumentUsecase *TypeDocumentUseCase) GetTypeDocumentUsecase() ([]dto.TypeDocumentDTO, error) {
	return typeDocumentUsecase.typeDocumentService.GetTypeDocumentService()
}
