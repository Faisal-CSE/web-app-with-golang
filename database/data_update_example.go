package database

import "log"

func UpdateRecordById(id int, clientName string) error {
	if id < 0 {
		id = 0
	}
	_, err := db.Exec(
		`
		UPDATE clientes set name = $1 WHERE id = $2`,
		clientName,
		id,
	)
	if err == nil {
		log.Println("Record successfully updated!")
	}
	return err
}
