package database

import (
	"fmt"
	"log"
)

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
		fmt.Printf("ID: %v\tNome: %v\n", p.ID, p.Name)
	}
}
