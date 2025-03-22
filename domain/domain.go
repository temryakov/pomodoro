package domain

import (
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

type History struct {
	Name     string
	Duration time.Duration
	Date     time.Time
}

type Repository interface {
	Get() ([]History, error)
	Post(duration time.Duration, recordName string) error
	Close()
	Clear() error
}

type Api interface {
	Send(word string) (string, error)
}

var (
	ErrNotFound = leveldb.ErrNotFound
)
