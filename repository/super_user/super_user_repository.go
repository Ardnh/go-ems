package superuser

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/domain"
)

type SuperUserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, user domain.SuperUser)
	FindByUsername(ctx context.Context, tx *sql.Tx, name string) (domain.SuperUser, error)
}
