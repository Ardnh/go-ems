package superuser

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
)

type SuperUserRepositoryImpl struct {
}

func NewSuperUserRepository() SuperUserRepository {
	return &SuperUserRepositoryImpl{}
}

func (repository *SuperUserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.SuperUser) {

	SQL := "INSERT INTO super_user ( firstname, lastname, username, organization, password ) VALUES ( ?,?,?,?,? )"
	_, err := tx.ExecContext(ctx, SQL, user.FirstName, user.LastName, user.UserName, user.Password)
	helper.PanicIfError(err)

}

func (repository *SuperUserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, name string) (domain.SuperUser, error) {

	SQL := "SELECT id, firstname, lastname, username, password FROM super_user WHERE username = ?;"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.SuperUser
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}
