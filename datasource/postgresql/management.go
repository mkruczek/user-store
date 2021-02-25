package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"log"
	"time"
)

func DoMagicWithDB() {

	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		5432, "localhost", "postgres", "password", "users")

	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		log.Fatalf("couldn't connect to DB : %s", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{"", "users", "public", 10 * time.Second})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../user-store/datasource/postgresql/migrations",
		"users",
		driver,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	m.Steps(2)
}
