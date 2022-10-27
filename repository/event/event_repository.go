package event

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/domain"
)

type EventRepostory interface {
	Save(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	Update(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	Delete(ctx context.Context, tx *sql.Tx, event domain.Event)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Event, error)
	FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) []domain.Event
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Event
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Event
}
