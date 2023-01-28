package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type ITypeDocumentService interface {
	GetTypeDocumentService() ([]dto.TypeDocumentDTO, error)
}
