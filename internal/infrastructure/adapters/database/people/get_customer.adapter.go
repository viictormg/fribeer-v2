package adapters

import (
	"errors"
	"fmt"

	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

const errorDBGetCustomer = "error consulting db"

func (peopleAdapter *PeopleAdapter) GetCustomerAdapter(companyID string) ([]dto.GetPeopleDTO, error) {
	var customers []dto.GetPeopleDTO

	const query = `SELECT p.id, CONCAT_WS(" ", p.firstName, p.secondName, p.surname, p.lastSurname) AS nameComplete,
						p.documentNumber, p.phone, p.email
						FROM People p
						JOIN TypePeople tp ON p.typePeople = tp.id
						WHERE tp.description = 'Cliente'
						AND p.company = ?`

	rows, err := peopleAdapter.db.Query(query, companyID)

	if err != nil {
		fmt.Println(err.Error())
		return customers, errors.New(errorDBGetCustomer)
	}

	defer rows.Close()

	for rows.Next() {
		var customer dto.GetPeopleDTO

		err = rows.Scan(
			&customer.ID,
			&customer.NameComplete,
			&customer.Document,
			&customer.Phone,
			&customer.Email,
		)

		if err != nil {
			return []dto.GetPeopleDTO{}, err
		}

		customers = append(customers, customer)
	}

	return customers, nil

}
