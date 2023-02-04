package mapper

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapCreatePeopleModelPeopleEntity(customer model.CustomerCreateModel) entity.PeopleEntity {
	return entity.PeopleEntity{
		FirstName:  customer.FirstName,
		TypePeople: constants.TypePeopleCustomerID,
		SecondName: customer.SecondName,
		Phone:      customer.Phone,
	}
}
