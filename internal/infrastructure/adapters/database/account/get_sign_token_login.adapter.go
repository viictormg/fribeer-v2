package adapters

import "github.com/viictormg/fribeer-v2/internal/domain/dto"

func (accountAdapter *AccountAdapter) GetSignTokenLoginAdapter(user, password string) (dto.SignTokenDTO, error) {
	var signToken dto.SignTokenDTO

	query := `SELECT ac.id AS accountID, pe.id AS peopleID, pe.company AS companyID
				FROM Account ac
				INNER JOIN People pe ON ac.people = pe.id
				WHERE (ac.username = ? AND password = ?) OR (ac.email = ? AND password = ?);`

	err := accountAdapter.db.QueryRow(query, user, password, user, password).Scan(
		&signToken.AccountID,
		&signToken.PeopleID,
		&signToken.CompanyID,
	)
	if err != nil {
		return signToken, err
	}

	return signToken, nil
}
