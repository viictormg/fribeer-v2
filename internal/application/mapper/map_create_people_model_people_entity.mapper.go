package mapper

import (
	"github.com/viictormg/fribeer-v2/internal/application/model"
	"github.com/viictormg/fribeer-v2/internal/domain/constants"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

func MapCreatePeopleModelPeopleEntity(customer model.CustomerCreateModel) entity.PeopleEntity {
	return entity.PeopleEntity{
		FirstName:      customer.FirstName,
		SecondName:     customer.SecondName,
		Surname:        customer.Surname,
		LastSurname:    customer.LastSurname,
		TypePeople:     constants.TypePeopleCustomerID,
		DocumentNumber: customer.DocumentNumber,
		Birthdate:      customer.Birthdate,
		Address:        customer.Address,
		Phone:          customer.Phone,
		Email:          customer.Email,
		TypeDocument:   customer.TypeDocument,
	}
}
