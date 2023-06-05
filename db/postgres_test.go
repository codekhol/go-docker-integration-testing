package db_test

import (
	"testing"

	"github.com/codekhol/test/db"
	dbtesting "github.com/codekhol/test/db/testing"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
)

func TestSave(t *testing.T) {
	tdb := dbtesting.Setup(t)
	defer dbtesting.Teardown(tdb, t)

	theDB := db.New(tdb.DB)
	err := theDB.Save(db.Book{
		Title:  "Go",
		ISBN:   "123",
		Author: "Go team",
	})
	if err != nil {
		t.Errorf("error saving book: %s", err)
	}
}
