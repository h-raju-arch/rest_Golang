package main

import (
	"fmt"
	"harish/TodoApp/db"
	"harish/TodoApp/types"

	"github.com/gin-gonic/gin"
)

// type Todo struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	Done        bool   `json:"done"`
// }

func main() {

	db.InitDB()

	defer db.CloseDb()

	router := gin.Default()

	//test
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello world"})
	})

	// addTodo
	router.POST("/todo", func(ctx *gin.Context) {
		var input struct {
			Title       string `json:"title"`
			Description string `json:"Description"`
		}
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(401, gin.H{
				"error": "Invalid Todo item",
			})
			return
		}

		todo := types.Todo{input.Title, input.Description, false}
		fmt.Printf("%+v", todo)

		id := db.AddTodo(todo)

		ctx.JSON(200, gin.H{
			"message": "Todo created successfully",
			"id":      id,
		})
	})

	//getAll Todos

	router.GET("/todos", func(ctx *gin.Context) {
		data := db.GetTodos()
		ctx.JSON(200, data)
	})
	router.Run(":8080")
}
