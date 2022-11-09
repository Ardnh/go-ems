package advertisement

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
)

type AdvertisementRepositoryImpl struct {
}

func NewAdvertisementRepositoryImpl() AdvertisementRepository {
	return &AdvertisementRepositoryImpl{}
}

func (repository *AdvertisementRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, ads domain.Advertisement) domain.Advertisement {
	SQL := "INSERT INTO advertisement (user_id, event_id, facebook_url, twitter_url, instagram_url, banner_url) VALUES ( ?,?,?,?,?,? );"
	result, err := tx.ExecContext(ctx, SQL, ads.UserId, ads.EventId, ads.FacebookUrl, ads.TwitterUrl, ads.InstagramUrl, ads.BannerUrl)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	ads.Id = int(id)

	return ads
}

func (repository *AdvertisementRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ads domain.Advertisement) domain.Advertisement {
	SQL := "UPDATE advertisement SET event_id = ?, facebook_url = ?, twitter_url = ?, instagram_url = ?, banner_url = ? WHERE id = ? AND user_id = ?;"
	_, err := tx.ExecContext(ctx, SQL, ads.EventId, ads.FacebookUrl, ads.TwitterUrl, ads.InstagramUrl, ads.BannerUrl, ads.Id, ads.UserId)
	helper.PanicIfError(err)

	return ads
}

func (repository *AdvertisementRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, advertiseId int, userId int) {
	SQL := "DELETE FROM advertisement WHERE id = ? AND user_id = ?;"
	_, err := tx.ExecContext(ctx, SQL, advertiseId, userId)
	helper.PanicIfError(err)
}

func (repository *AdvertisementRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, idAds int) (domain.Advertisement, error) {
	SQL := "SELECT id, user_id, facebook_url, twitter_url, instagram_url, banner_url FROM advertisement WHERE id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, idAds)
	helper.PanicIfError(err)
	defer rows.Close()

	var advertise domain.Advertisement
	if rows.Next() {
		err := rows.Scan(&advertise.Id, &advertise.UserId, &advertise.EventId, &advertise.FacebookUrl, &advertise.TwitterUrl, &advertise.InstagramUrl, &advertise.BannerUrl)
		helper.PanicIfError(err)

		return advertise, nil
	} else {

		return advertise, errors.New("advertise not found")
	}
}

func (repository *AdvertisementRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Advertisement {
	SQL := "SELECT * FROM advertisement WHERE user_id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)

	var advertisement []domain.Advertisement
	for rows.Next() {
		ads := domain.Advertisement{}
		err := rows.Scan(&ads.Id, &ads.UserId, &ads.EventId, &ads.FacebookUrl, &ads.TwitterUrl, &ads.InstagramUrl, &ads.BannerUrl)
		helper.PanicIfError(err)

		advertisement = append(advertisement, ads)
	}

	return advertisement
}

func (repository *AdvertisementRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Advertisement {

	SQL := "SELECT * FROM advertisement;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var advertisement []domain.Advertisement
	for rows.Next() {
		ads := domain.Advertisement{}
		err := rows.Scan(&ads.Id, &ads.UserId, &ads.EventId, &ads.FacebookUrl, &ads.TwitterUrl, &ads.InstagramUrl, &ads.BannerUrl)
		helper.PanicIfError(err)

		advertisement = append(advertisement, ads)
	}

	return advertisement
}
