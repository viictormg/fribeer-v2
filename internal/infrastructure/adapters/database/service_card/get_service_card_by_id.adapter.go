package adapters

import (
	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (s *ServiceCardAdapter) GetServiceCardByIDAdapter(companyID, id string) (dto.GetServiceCardDTO, error) {
	var card dto.GetServiceCardDTO
	query := `SELECT s.id, serviceName, s.description, s.startDate, s.endDate, ABS(DATEDIFF(NOW(), s.endDate)) AS expiration,
				CONCAT_WS(" ",p.firstName, p.secondName, p.surname, p.lastSurname)AS customer,
				st.name, st.code
				FROM ServiceCard s
				JOIN People p ON s.customer = p.id
				JOIN State st ON  s.state = st.id
				WHERE s.company = ? AND s.id = ? `

	err := s.db.QueryRow(query, companyID, id).Scan(
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
		return card, err
	}

	return card, nil
}
