package domain

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type History struct {
	Name     string
	Duration string
}

type Repository interface {
	Get() ([]History, error)
	Post(value string) error
	Close()
}

type Api interface {
	Send(word string) (string, error)
}

var (
	ErrNotFound = leveldb.ErrNotFound
)
