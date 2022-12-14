package advertisement

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/domain"
)

type AdvertisementRepository interface {
	Save(ctx context.Context, tx *sql.Tx, advertise domain.Advertisement) domain.Advertisement
	Update(ctx context.Context, tx *sql.Tx, advertise domain.Advertisement) domain.Advertisement
	Delete(ctx context.Context, tx *sql.Tx, advertiseId int, userId int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Advertisement, error)
	FindByUserId(ctx context.Context, tx *sql.Tx, UserId int) []domain.Advertisement
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Advertisement
}
