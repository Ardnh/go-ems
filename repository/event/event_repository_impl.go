package event

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Ardnh/go-ems/helper"
	"github.com/Ardnh/go-ems/model/domain"
)

type EventRepositoryImpl struct {
}

func NewRepositoryImpl() EventRepostory {
	return &EventRepositoryImpl{}
}

func (repository *EventRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	SQL := "INSERT INTO events(user_id, category_id, name, tagline, description, organizer, start_date, end_date, registration_start_date, registration_end_date, location, capacity, banner_url, status) VALUES ( ?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	result, err := tx.ExecContext(ctx, SQL, event.UserId, event.CategoryId, event.Name, event.Tagline, event.Description, event.Organizer, event.StartDate, event.EndDate, event.RegistrationStartDate, event.RegistrationEndDate, event.Location, event.Capacity, event.BannerUrl, event.Status)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	event.Id = int(id)

	return event
}

func (repository *EventRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event {
	SQL := "UPDATE events SET category_id = ?, name = ? , tagline = ?, description = ?, organizer = ?, start_date = ?, end_date = ?, registration_start_date = ?, registration_end_date = ?, location = ?, capacity = ?, banner_url = ?, status = ? WHERE id = ? AND user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, event.CategoryId, event.Name, event.Tagline, event.Description, event.Organizer, event.StartDate, event.EndDate, event.RegistrationStartDate, event.RegistrationEndDate, event.Location, event.Capacity, event.BannerUrl, event.Status, event.Id, event.UserId)
	helper.PanicIfError(err)

	return event
}

func (repository *EventRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, event domain.Event) {
	SQL := "DELETE FROM events WHERE id = ? AND user_id = ?;"
	_, err := tx.ExecContext(ctx, SQL, event.Id, event.UserId)
	helper.PanicIfError(err)
}

func (repository *EventRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Event, error) {
	SQL := "SELECT * FROM events WHERE id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var event domain.Event
	if rows.Next() {
		err := rows.Scan(&event.Id, &event.UserId, &event.CategoryId, &event.Name, &event.Tagline, &event.Description, &event.Organizer, &event.StartDate, &event.EndDate, &event.RegistrationStartDate, &event.RegistrationEndDate, &event.Location, &event.Capacity, &event.BannerUrl, &event.Status)
		helper.PanicIfError(err)

		return event, nil
	} else {
		return event, errors.New("event not found")
	}
}

func (repository *EventRepositoryImpl) FindByCategoryId(ctx context.Context, tx *sql.Tx, categoryId int) []domain.Event {
	SQL := "SELECT * FROM events WHERE category_id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		event := domain.Event{}
		err := rows.Scan(&event.Id, &event.UserId, &event.CategoryId, &event.Name, &event.Tagline, &event.Description, &event.Organizer, &event.StartDate, &event.EndDate, &event.RegistrationStartDate, &event.RegistrationEndDate, &event.Location, &event.Capacity, &event.BannerUrl, &event.Status)
		helper.PanicIfError(err)

		events = append(events, event)
	}

	return events
}

func (repository *EventRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) []domain.Event {
	SQL := "SELECT * FROM events WHERE user_id = ?;"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		event := domain.Event{}
		err := rows.Scan(&event.Id, &event.UserId, &event.CategoryId, &event.Name, &event.Tagline, &event.Description, &event.Organizer, &event.StartDate, &event.EndDate, &event.RegistrationStartDate, &event.RegistrationEndDate, &event.Location, &event.Capacity, &event.BannerUrl, &event.Status)
		helper.PanicIfError(err)

		events = append(events, event)
	}

	return events
}

func (repository *EventRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Event {
	SQL := "SELECT * FROM events;"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		event := domain.Event{}
		err := rows.Scan(&event.Id, &event.UserId, &event.CategoryId, &event.Name, &event.Tagline, &event.Description, &event.Organizer, &event.StartDate, &event.EndDate, &event.RegistrationStartDate, &event.RegistrationEndDate, &event.Location, &event.Capacity, &event.BannerUrl, &event.Status)
		helper.PanicIfError(err)

		events = append(events, event)
	}

	return events
}
