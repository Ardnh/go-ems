package advertisement

import (
	"context"
	"database/sql"

	"github.com/Ardnh/go-ems/model/domain"
)

type AdvertisementRepositoryImpl struct {
}

func NewAdvertisementRepositoryImpl() AdvertisementRepository {
	return &AdvertisementRepositoryImpl{}
}

func (repository *AdvertisementRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, ads domain.Advertisement) domain.Advertisement {

	return ads
}

func (repository *AdvertisementRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ads domain.Advertisement) domain.Advertisement {

	return ads
}

func (repository *AdvertisementRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ads domain.Advertisement) {

}

func (repository *AdvertisementRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) domain.Advertisement {

	var advertise domain.Advertisement

	return advertise
}

func (repository *AdvertisementRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Advertisement {

	var advertisement []domain.Advertisement

	return advertisement
}
