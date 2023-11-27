package controller

import (
	"log"
	"net/http"
)

// getTodosApi will write a CSV of the todos to w
func getTodosApi(w http.ResponseWriter, r *http.Request) {
	log.Println("ERROR: getTodosApi not implemented yet")
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
