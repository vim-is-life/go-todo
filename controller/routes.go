package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	model "github.com/vim-is-life/go-todo/model"
)

func serveTodosToDisplay(w http.ResponseWriter) {
	todoList := model.GetAllTodos()
	tmpl := template.Must(template.ParseFiles("views/todos.gohtml"))

	err := tmpl.ExecuteTemplate(w, "Todos", todoList)
	model.LogErr(err)
}

// getIdxPage serves a parsed html template containing all the todos
func getIdxPage(w http.ResponseWriter, r *http.Request) {
	log.Println("getIdxPage: not implemented yet!")
}

// markTodo inverts the completion state of a todoItem.
// The id of the item we want to change will be in r. After the state has been
// changed this function will re-serve the index page.
func markTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("markTodo: not implemented yet!")
	serveTodosToDisplay(w)
}

// createTodo will parse information from an html form to add a todo to the DB.
// The information for the new todo will be in r. After the state has been
// changed this function will re-serve the index page.
func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("createTodo: not implemented yet!")
	serveTodosToDisplay(w)
}

// deleteTodo removes a todo from the todo list.
// Information about the todo item to delete will be in r.
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("deleteTodo: not implemented yet!")
	serveTodosToDisplay(w)
}

// SetupAndRun starts the server and runs it.
// The port to run it on is specified by the APP_PORT environment variable.
func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", getIdxPage)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/createTodo", createTodo).Methods("POST")

	// port must be in form ':abdc' where abcd are numbers
	appPort := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(appPort, mux))
}
