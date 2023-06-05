package db

import (
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) DB {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Save(b Book) error {
	args := map[string]interface{}{
		"title":  b.Title,
		"isbn":   b.ISBN,
		"author": b.Author,
	}
	_, err := p.db.NamedExec("INSERT INTO books (title, isbn, author) VALUES (:title, :isbn, :author)", args)
	return err
}
