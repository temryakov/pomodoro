package repository

import (
	"database/sql"
	"time"

	"github.com/temryakov/pomodoro/domain"
)

type Repository struct {
	database *sql.DB
}

func NewRepository(db *sql.DB) domain.Repository {
	return &Repository{
		database: db,
	}
}

func (r *Repository) Get() ([]domain.History, error) {
	rows, err := r.database.Query("select * from Pomodoro")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	pomodoros := []domain.History{}

	for rows.Next() {
		var seconds int64
		h := domain.History{}
		err := rows.Scan(&h.Name, seconds, &h.Date)
		if err != nil {
			return nil, err
		}
		h.Duration = time.Duration(seconds) * time.Second
		pomodoros = append(pomodoros, h)
	}
	return pomodoros, nil
}

func (r *Repository) Post(duration time.Duration, recordName string) error {
	_, err := r.database.Exec("insert into Pomodoro (name, duration, end_time) values ($1, $2, $3)",
		recordName, int(duration.Seconds()), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Close() {
	r.database.Close()
}

func (r *Repository) Clear() error {
	_, err := r.database.Exec("delete from Pomodoro")
	if err != nil {
		return err
	}
	return nil
}
