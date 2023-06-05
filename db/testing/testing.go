package testing

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var TestInstance = struct {
	Driver string
	Conn   string
}{
	Driver: "pgx",
	Conn:   "postgres://test:test@localhost:5432/test",
}

type TestDB struct {
	DB     *sqlx.DB
	DBName string
}

func Setup(t *testing.T) TestDB {
	var (
		dbname  = fmt.Sprintf("test_%d", rand.Int())
		connstr = fmt.Sprintf("postgres://test:test@localhost:5432/%s?sslmode=disable", dbname)
	)

	conn, err := sqlx.Open(TestInstance.Driver, TestInstance.Conn)
	if err != nil {
		t.Errorf("error opening conn: %s", err)
	}
	defer conn.Close()

	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s OWNER test;", dbname))
	if err != nil {
		t.Errorf("error creating database: %s", err)
	}

	testconn, err := sqlx.Open("postgres", connstr)
	if err != nil {
		t.Errorf("error opening test conn: %s", err)
	}

	mig, err := migrate.New("file://migrations", connstr)
	if err != nil {
		t.Errorf("error creating migrations: %s", err)
	}
	err = mig.Up()
	if err != nil {
		t.Errorf("error running migrations: %s", err)
	}

	return TestDB{
		DB:     testconn,
		DBName: dbname,
	}
}

func Fixture(tdb TestDB, stmts []string, t *testing.T) {
	conn, err := sqlx.Open(TestInstance.Driver, TestInstance.Conn)
	if err != nil {
		t.Errorf("error opening conn: %s", err)
	}
	defer conn.Close()

	for _, stmt := range stmts {
		_, err = conn.Exec(stmt)
		if err != nil {
			t.Errorf("error executing fixture statement `%s`: %s", stmt, err)
		}
	}
}

func Teardown(tdb TestDB, t *testing.T) {
	conn, err := sqlx.Open(TestInstance.Driver, TestInstance.Conn)
	if err != nil {
		t.Errorf("error opening conn: %s", err)
	}
	defer conn.Close()

	_, err = conn.Exec(fmt.Sprintf("DROP DATABASE %s WITH (FORCE)", tdb.DBName))
	if err != nil {
		t.Errorf("error dropping test database: %s", err)
	}
}
