package db

import "errors"

var ErrDuplicate = errors.New("duplicate item")

type Book struct {
	Title  string
	ISBN   string
	Author string
}

type DB interface {
	Save(b Book) error
}
