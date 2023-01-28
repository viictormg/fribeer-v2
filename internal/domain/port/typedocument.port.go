package port

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

type ITypeDocumentAdapter interface {
	GetTypeDocumentAdapter() ([]dto.TypeDocumentDTO, error)
}
