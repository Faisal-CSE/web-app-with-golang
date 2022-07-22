package database

import "log"

func RemoveRecordById(id int) error {
	if id < 0 {
		id = 0
	}
	_, err := db.Exec(
		`
		DELETE FROM clientes WHERE id = $1`,
		id,
	)
	if err == nil {
		log.Println("Record successfully deleted!")
	}
	return err
}
