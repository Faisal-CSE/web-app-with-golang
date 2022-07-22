package database

import (
	"fmt"
	"log"
)

func GetRecordCount() {
	var count int

	err := db.QueryRow("SELECT COUNT(*) FROM clientes").Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Number of rows are %v\n", count)
	}
}

func GetRecordByID(Id int) {
	if Id < 0 {
		Id = 0
	}
	SQL := `SELECT "id","name" FROM "clientes" WHERE id = $1`

	rows, err := db.Queryx(SQL, Id)
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		p := struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}{}
		err = rows.StructScan(&p)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ID: %v\tClient Name: %v\n", p.ID, p.Name)
	}
}

func ListRecords(page int) {

	limit := 5
	offset := limit * (page - 1)

	SQL := `SELECT "id","name" FROM "clientes" ORDER BY "id" LIMIT $2 OFFSET $1`

	rows, err := db.Queryx(SQL, offset, limit)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next() {
		p := struct {
			ID   int    `db:"id"`
			Name string `db:"name"`
		}{}
		err = rows.StructScan(&p)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ID: %v\tClient Name: %v\n", p.ID, p.Name)
	}
}
