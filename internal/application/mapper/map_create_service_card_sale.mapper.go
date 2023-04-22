package mapper

import (
	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapperCreateCardService(services []dto.DetailProductToCreateSale, customer, sale, companyID string) []entity.ServiceCardEntity {
	var cards []entity.ServiceCardEntity

	for _, service := range services {

		if service.TypeProduct == constants.TypeProductService && service.IsFrequency {
			card := entity.ServiceCardEntity{
				ID:          uuid.NewString(),
				Customer:    customer,
				ServiceName: service.Name,
				Sale:        sale,
				StartDate:   service.StartDate,
				EndDate:     service.EndDate,
				SaleDetail:  service.ID,
				State:       constants.StateServiceCard[constants.StateIDCurrent],
				Company:     companyID,
			}
			cards = append(cards, card)
		}
	}

	return cards
}
