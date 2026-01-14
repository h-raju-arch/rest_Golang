package db

import (
	"harish/TodoApp/types"
	"log"
)

func AddTodo(todo types.Todo) int {
	query := `INSERT  INTO todo (title,description) VALUES($1,$2) RETURNING id`

	var id int
	err := Db.QueryRow(query, todo.Title, todo.Description).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}
