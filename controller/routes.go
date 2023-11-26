package controller

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	model "github.com/vim-is-life/go-todo/model"
)

func serveTodosToDisplay(w http.ResponseWriter) {
	todoList := model.GetAllTodos()
	tmpl := template.Must(template.ParseFiles("views/todos.gohtml"))

	// NOTE: since we call ExecuteTemplate, this will only work on the "Todos"
	// named template in template html file
	err := tmpl.ExecuteTemplate(w, "Todos", todoList)
	model.LogErr(err)
}

// getIdxPage serves a parsed html template containing all the todos
func getIdxPage(w http.ResponseWriter, r *http.Request) {
	todoList := model.GetAllTodos()
	tmpl := template.Must(template.ParseFiles("views/todos.gohtml"))

	// NOTE: since we call execute, this will work on the WHOLE template html
	// file
	err := tmpl.Execute(w, todoList)
	model.LogErr(err)
}

// markTodo inverts the completion state of a todoItem.
// The id of the item we want to change will be in r. After the state has been
// changed this function will re-serve the index page.
func markTodo(w http.ResponseWriter, r *http.Request) {
	todo_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("couldn't parse id from url", err)
	}

	model.MarkDone(uint(todo_id))
	serveTodosToDisplay(w)
}

// createTodo will parse information from an html form to add a todo to the DB.
// The information for the new todo will be in r. After the state has been
// changed this function will re-serve the index page.
func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	model.LogErr(err)

	newKind, err := strconv.Atoi(r.FormValue("newTodoKind"))
	if err != nil {
		log.Println(err)
	}
	// model.LogErr(err)

	// note that
	// - we don't need to worry about id because db handles this for us
	// - we don't need to worry about state because it will default to StateTodo
	newTodoItem := model.TodoItem{
		Name: r.FormValue("newTodoName"),
		Kind: model.TodoKind(newKind),
		Desc: r.FormValue("newTodoDesc"),
	}

	// fmt.Printf("%+v\n", newTodoItem)
	model.AddTodo(newTodoItem)
	serveTodosToDisplay(w)
}

// deleteTodo removes a todo from the todo list.
// Information about the todo item to delete will be in r.
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todo_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("couldn't parse id from url", err)
	}

	model.DeleteTodo(uint(todo_id))
	serveTodosToDisplay(w)
}

// SetupAndRun starts the server and runs it.
// The port to run it on is specified by the APP_PORT environment variable.
func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", getIdxPage)
	mux.HandleFunc("/markTodo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/delete/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/createTodo", createTodo).Methods("POST")

	// port must be in form ':abdc' where abcd are numbers
	appPort := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(appPort, mux))
}
