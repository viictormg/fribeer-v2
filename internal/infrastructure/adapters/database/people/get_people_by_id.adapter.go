package adapters

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (peopleAdapter *PeopleAdapter) GetPeopleByIDAdapter(companyID, id string) (dto.GetPeopleDTO, error) {
	var people dto.GetPeopleDTO

	fmt.Println(companyID, id)

	query := `SELECT p.id, TRIM( CONCAT_WS(" ", p.firstName, p.secondName, p.surname, p.lastSurname)) AS fullName,
				p.documentNumber, p.phone, p.email,
				IF(p.isActive = 1, true, false) active
				FROM People p
				JOIN TypePeople tp ON p.typePeople = tp.id
				WHERE tp.description = 'Cliente'
				AND p.company = ? AND p.id = ?`

	err := peopleAdapter.db.QueryRow(query, companyID, id).Scan(
		&people.ID,
		&people.FullName,
		&people.Document,
		&people.Phone,
		&people.Email,
		&people.Active,
	)
	if err != nil {
		logrus.Error(err)
		return people, err
	}

	return people, nil
}
