package mapper

import (
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapperCreateCardService(services []dto.DetailProductToCreateSale) []entity.ServiceCardEntity {
	var cards []entity.ServiceCardEntity

	for _, service := range services {

		if service.TypeProduct == constants.TypeProductService && service.IsFrequency {
			card := entity.ServiceCardEntity{
				ServiceName: service.Name,
			}
			cards = append(cards, card)
		}
	}

	return cards
}
