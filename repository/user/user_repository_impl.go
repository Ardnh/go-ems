package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) {

	SQL := "INSERT INTO user ( firstname, lastname, username, organization, password ) VALUES ( ?,?,?,?,? )"
	_, err := tx.ExecContext(ctx, SQL, user.FirstName, user.LastName, user.UserName, user.Organization, user.Password)
	helper.PanicIfError(err)

}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, name string) (domain.User, error) {

	SQL := "SELECT id, firstname, lastname, username, organization, password FROM user WHERE username = ?;"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.UserName, &user.Organization, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}

}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "UPDATE user SET firstname = ?, lastname = ?, organization = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.FirstName, user.LastName, user.Organization, user.Id)
	helper.PanicIfError(err)

	return err
}
