package adapters

import (
	"errors"

	"github.com/google/uuid"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
	"github.com/viictormg/fribeer-v2/internal/domain/entity"
)

const errCreatePeopleDatabase = "error create people database"

func (peopleAdapter *PeopleAdapter) CreatePeopleAdapter(people entity.PeopleEntity, company string) (dto.CreationDTO, error) {
	id := uuid.NewString()

	query := `INSERT INTO People(id, typePeople, firstName, secondName, 
								surname, lastSurname, typeDocument, documentNumber, 
								birthdate, phone, address, email, company, isActive) 
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := peopleAdapter.db.Query(query,
		id,
		people.TypePeople,
		people.FirstName,
		people.SecondName,
		people.Surname,
		people.LastSurname,
		people.TypeDocument,
		people.DocumentNumber,
		people.Birthdate,
		people.Phone,
		people.Address,
		people.Email,
		company,
		true,
	)

	if err != nil {
		return dto.CreationDTO{}, errors.New(errCreatePeopleDatabase)
	}

	return dto.CreationDTO{ID: id}, nil
}
