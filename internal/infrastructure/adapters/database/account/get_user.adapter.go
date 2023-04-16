package adapters

import (
	"github.com/sirupsen/logrus"
	"github.com/viictormg/fribeer-v2/internal/domain/dto"
)

func (a *AccountAdapter) GetUserAdapter(AccountID string) (dto.GetUserDTO, error) {
	var user dto.GetUserDTO

	query := `SELECT username, ac.email, IFNULL(ac.avatar, '') AS avatar, CONCAT_WS(p.firstName, p.secondName, p.surname, p.lastSurname) AS fullName
				FROM Account  ac
				JOIN People p ON ac.people = p.id
				WHERE ac.id = ?`

	err := a.db.QueryRow(query, AccountID).Scan(
		&user.UserName,
		&user.Email,
		&user.Avatar,
		&user.FullName,
	)
	if err != nil {
		logrus.Error(err)
		return user, err
	}

	return user, nil
}
