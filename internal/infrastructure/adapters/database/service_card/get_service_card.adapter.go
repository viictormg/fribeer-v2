package adapters

import (
	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (s *ServiceCardAdapter) GetServiceCardAdapter(companyID string) ([]dto.GetServiceCardDTO, error) {
	cards := []dto.GetServiceCardDTO{}

	query := `SELECT s.id, serviceName, s.description, s.startDate, s.endDate, ABS(DATEDIFF(NOW(), s.endDate)) AS expiration,
				CONCAT_WS(" ",p.firstName, p.secondName, p.surname, p.lastSurname)AS customer,
				st.name, st.code
				FROM ServiceCard s
				JOIN People p ON s.customer = p.id
				JOIN State st ON  s.state = st.id
				WHERE s.company = ?;`

	rows, err := s.db.Query(query, companyID)

	if err != nil {
		logrus.Error(err)
		return cards, err
	}

	defer rows.Close()

	for rows.Next() {
		var card = dto.GetServiceCardDTO{}

		err = rows.Scan(
			&card.ID,
			&card.ServiceName,
			&card.Description,
			&card.StartDate,
			&card.EndDate,
			&card.Expiration,
			&card.CustomerName,
			&card.State,
			&card.StateCode,
		)

		if err != nil {
			logrus.Error(err)
			defer rows.Close()
			return cards, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (s *ServiceCardAdapter) GeAllServicesCardAdapter() ([]dto.GetServiceCardDTO, error) {
	cards := []dto.GetServiceCardDTO{}

	query := `SELECT s.id, serviceName, s.description, s.startDate, s.endDate, DATEDIFF(NOW(), s.endDate) AS expiration,
				CONCAT_WS(" ",p.firstName, p.secondName, p.surname, p.secondName)AS customer,
				st.name, st.code
				FROM ServiceCard s
				JOIN People p ON s.customer = p.id
				JOIN State st ON  s.state = st.id`

	rows, err := s.db.Query(query)

	if err != nil {
		logrus.Error(err)
		return cards, err
	}

	defer rows.Close()

	for rows.Next() {
		var card = dto.GetServiceCardDTO{}

		err = rows.Scan(
			&card.ID,
			&card.ServiceName,
			&card.Description,
			&card.StartDate,
			&card.EndDate,
			&card.Expiration,
			&card.CustomerName,
			&card.State,
			&card.StateCode,
		)

		if err != nil {
			logrus.Error(err)
			defer rows.Close()
			return cards, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
