package database

import (
	"log"
)

func InsertNewRecord(clientName string) error {
	if clientName == "" {
		clientName = "DUMMY CLIENT"
	}

	_, err := db.Exec(
		`
		INSERT INTO clientes (name)
		VALUES ($1)
		`,
		clientName,
	)
	if err == nil {
		log.Println("Record successfully inserted!")
	}
	return err

}
