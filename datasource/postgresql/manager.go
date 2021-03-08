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
	"github.com/mkruczek/user-store/utils/logger"
)

func DoMagicWithDB(cfg *config.Config, LOG logger.Logger) error {
	if !cfg.DB.Migration.Run {
		LOG.Infof("skip DB migration")
		return nil
	}

	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Port, cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		return err
	}
	defer db.Close()

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

	err = m.Steps(cfg.DB.Migration.Steps)
	if err != nil {
		return err
	}

	return nil
}
