package domain

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

type History struct {
	Name     string
	Duration string
	Date     time.Time
}

const (
	PomodoroRecord = "pomodoro"
	BreakRecord    = "break"
)

type Repository interface {
	Get() ([]History, error)
	Post(value, recordName string) error
	Close()
}

type Api interface {
	Send(word string) (string, error)
}

var (
	ErrNotFound = leveldb.ErrNotFound
)
