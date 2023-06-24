package adapters

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

const errorDBGetCustomer = "error consulting db"

func (peopleAdapter *PeopleAdapter) GetCustomerAdapter(companyID string) ([]dto.GetPeopleDTO, error) {
	customers := []dto.GetPeopleDTO{}

	const query = `SELECT p.id, CONCAT_WS(" ", p.firstName, p.secondName, p.surname, p.lastSurname) AS fullName,
						p.documentNumber, p.phone, p.email,
						IF(p.isActive = 1, true, false) active
						FROM People p
						JOIN TypePeople tp ON p.typePeople = tp.id
						WHERE tp.description = 'Cliente'
						AND p.company = ?`

	rows, err := peopleAdapter.db.Query(query, companyID)

	if err != nil {
		logrus.Error(err)
		return customers, errors.New(errorDBGetCustomer)
	}

	defer rows.Close()

	for rows.Next() {
		var customer dto.GetPeopleDTO
		err = rows.Scan(
			&customer.ID,
			&customer.FullName,
			&customer.Document,
			&customer.Phone,
			&customer.Email,
			&customer.Active,
		)

		if err != nil {
			logrus.Error(err)
			return []dto.GetPeopleDTO{}, err
		}

		customers = append(customers, customer)
	}
	return customers, nil

}
