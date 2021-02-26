package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/mkruczek/user-store/config"
)

func DoMagicWithDB(cfg *config.Config) error {

	if !cfg.DB.Migration.Run {
		return nil
	}

	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Port, cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		return err
	}

	driver, err := pg.WithInstance(db, &pg.Config{"", cfg.DB.Name, cfg.DB.Schema, cfg.DB.Timeout})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", cfg.DB.Migration.Files),
		cfg.DB.Name,
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Steps(1) //number of files to call
	if err != nil {
		return err
	}

	return nil
}
