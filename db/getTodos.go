package db

import (
	"harish/TodoApp/types"
	"log"
)

func GetTodos() []types.TodoOutput {
	query := `SELECT id,title,description,done FROM todo`

	rows, err := Db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	data := []types.TodoOutput{}
	var title string
	var description string
	var id int
	var done bool

	for rows.Next() {
		err1 := rows.Scan(&id, &title, &description, &done)
		if err1 != nil {
			log.Fatal(err1)
		}
		data = append(data, types.TodoOutput{id, title, description, done})
	}

	return data
}
