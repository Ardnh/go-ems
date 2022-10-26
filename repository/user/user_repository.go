package user

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.User)
	FindByUsername(ctx context.Context, tx *sql.Tx, name string) (domain.User, error)
}
