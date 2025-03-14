package repository

import (
	"database/sql"

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
		h := domain.History{}
		err := rows.Scan(&h.Name, &h.Duration)
		if err != nil {
			return nil, err
		}
		pomodoros = append(pomodoros, h)
	}
	return pomodoros, nil
}

func (r *Repository) Post(value string) error {
	_, err := r.database.Exec("insert into Pomodoro (name, duration) values ($1, $2)",
		"pomodoro", value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Close() {
	r.database.Close()
}
