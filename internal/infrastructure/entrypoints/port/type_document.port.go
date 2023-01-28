package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type ITypeDocumentUsecase interface {
	GetTypeDocumentUsecase() ([]dto.TypeDocumentDTO, error)
}
