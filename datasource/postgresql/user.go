package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const ( //todo move to config
	hostname     = "localhost"
	hostPort     = 5432
	username     = "postgres"
	password     = "password"
	databaseName = "user"
)

type UserDB struct {
	*sql.DB
}

func NewUserDBConnection() *UserDB {

	//pgConString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	//	username, password, hostname, hostPort, databaseName)

	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable",
		hostPort, hostname, username, password, databaseName)

	result, err := sql.Open("postgres", pgConString)
	if err != nil {
		log.Fatalf("couldn't connect to DB : %s", err.Error())
	}

	if err = result.Ping(); err != nil {
		log.Fatalf("couldn't ping the DB : %s", err.Error())
	}

	log.Printf("success connecting to db : %s", databaseName)
	return &UserDB{result}
}
