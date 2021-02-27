package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mkruczek/user-store/config"
	"log"
)

type UserDB struct {
	*sql.DB
}

func NewUserDBConnection(cfg *config.Config) *UserDB {
	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Port, cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)

	result, err := sql.Open("postgres", pgConString)
	if err != nil {
		log.Fatalf("couldn't connect to DB : %s", err.Error())
	}

	if err = result.Ping(); err != nil {
		log.Fatalf("couldn't ping the DB : %s", err.Error())
	}

	log.Printf("success connecting to db : %s", cfg.DB.Name)
	return &UserDB{result}
}
