package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	env "github.com/web-app-with-golang/project_env"
)

var (
	db  *sqlx.DB
	err error
)

func PostgresDbConnection() *sqlx.DB {
	db, err = sqlx.Connect("postgres", env.GoDotEnvVariable("DSN_PG"))
	if err != nil {
		fmt.Printf("error open db: %v\n", err)
		return nil
	}
	return db
}
