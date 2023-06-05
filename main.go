package main

import (
	"fmt"

	"github.com/codekhol/test/db"
	dbtesting "github.com/codekhol/test/db/testing"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	conn, err := sqlx.Open(dbtesting.TestInstance.Driver, dbtesting.TestInstance.Conn)
	if err != nil {
		panic(fmt.Sprintf("error opening conn: %s\n", err))
	}
	theDB := db.New(conn)
	err = theDB.Save(db.Book{
		Title:  "Go",
		ISBN:   "123",
		Author: "Go team",
	})
	if err != nil {
		panic(fmt.Sprintf("error saving book: %s\n", err))
	}

	fmt.Println("finished successfully")
}
