package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"

	model "github.com/vim-is-life/go-todo/model"
)

// getTodosApi will write a CSV of the todos to w
func getTodosApi(w http.ResponseWriter, r *http.Request) {
	todoList := model.GetAllTodos()
	cw := csv.NewWriter(w)
	const numCols = 5

	for _, todoItem := range todoList {
		// record := make([]string, numCols)
		record := []string{
			fmt.Sprintf("%d", todoItem.TodoId),
			todoItem.Name,
			todoItem.Desc,
			fmt.Sprintf("%d", todoItem.Kind),
			fmt.Sprintf("%d", todoItem.State),
		}

		if err := cw.Write(record); err != nil {
			log.Println("Error writing csv record: ", err)
			break
		}
	}

	cw.Flush()
}

// markTodoApi will change the state of a todo item from Todo->InProgress->Done.
func markTodoApi(w http.ResponseWriter, r *http.Request) {
	log.Println("ERROR: markTodoApi not implemented yet")
	getTodosApi(w, r)
}

// deleteTodoApi will remove a todo item from the list from the API.
func deleteTodoApi(w http.ResponseWriter, r *http.Request) {
	log.Println("ERROR: deleteTodoApi not implemented yet")
	getTodosApi(w, r)
}

// createTodoApi will add a todo item to the list from the API.
func createTodoApi(w http.ResponseWriter, r *http.Request) {
	log.Println("ERROR: createTodoApi not implemented yet")
	getTodosApi(w, r)
}
