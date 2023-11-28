// api.go is the API implementation for interactions with the backend outside of the browser.
// The format for communication will be semicolon-delimited (;) CSV. Since we
// have tabular data, CSV works well for us. Also, compared to JSON, CSV is
// smaller and seems to be faster to serialize and deserialize.
//
// Specs for the CSV format used to communicate:
// - it will have no header
// - it will have 5 columns
//   - column 1 will be unsigned int
//   - column 2 will be string
//   - column 3 will be string
//   - column 4 will be int matching the underlying values of the TodoKind type
//     (see model/defs)
//   - column 5 will be int matching the underlying values of the TodoState type
//     (see model/defs)
//
// An example response to get all the todos from the api, for instance, could
// look like the following:
// 3;finish todo app;;-1;-1
// 1;learn rust;learn rust to ascend to a higher plane;0;0
package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/vim-is-life/go-todo/model"
)

// getTodosApi will write a CSV of the todos to w
func getTodosApi(w http.ResponseWriter, r *http.Request) {
	todoList := model.GetAllTodos()
	cw := csv.NewWriter(w)
	cw.Comma = ';'
	const numCols = 5

	for _, todoItem := range todoList {
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

// createTodoApi will add a todo item to the list from the API.
func createTodoApi(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	model.LogErr(err)

	log.Println(r.FormValue("kind"))
	newKind, err := strconv.Atoi(r.FormValue("kind"))
	if err != nil {
		log.Println(err)
	}

	// note that
	// - we don't need to worry about id because db handles this for us
	// - we don't need to worry about state because it will default to StateTodo
	newTodoItem := model.TodoItem{
		Name: r.FormValue("name"),
		Kind: model.TodoKind(newKind),
		Desc: r.FormValue("desc"),
	}

	// fmt.Printf("%+v\n", newTodoItem)
	model.AddTodo(newTodoItem)
}

// markTodoApi will change the state of a todo item from Todo->InProgress->Done.
func markTodoApi(w http.ResponseWriter, r *http.Request) {
	todo_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("couldn't parse id from url", err)
	}

	model.MarkDone(uint(todo_id))
}

// markTodoApi will change the state of a todo item from Todo->InProgress->Done.
func updateTodoApi(w http.ResponseWriter, r *http.Request) {
	log.Println("updateTodoApi not implemented yet")
	// todo_id, err := strconv.Atoi(mux.Vars(r)["id"])
	// if err != nil {
	// 	log.Println("couldn't parse id from url", err)
	// }

	// model.MarkDone(uint(todo_id))
}

// deleteTodoApi will remove a todo item from the list from the API.
func deleteTodoApi(w http.ResponseWriter, r *http.Request) {
	todo_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("couldn't parse id from url", err)
	}

	model.DeleteTodo(uint(todo_id))
}
