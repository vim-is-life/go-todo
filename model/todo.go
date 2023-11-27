// Package model manages DB interactions as well as validates input
package model

// TODO add code to validate input at some point

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

// LogErr fatally logs any non-nil error passed to it.
func LogErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// InitDB will open the database and create the table if it doesn't exist.
func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "./todolist.sqlite")
	LogErr(err)
	err = db.Ping()
	LogErr(err)

	const migrationStr = `CREATE TABLE IF NOT EXISTS TodoList (
		Todo_id     INTEGER  NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
		Name        TEXT     NOT NULL UNIQUE,
		Desc        TEXT     NOT NULL DEFAULT '',
		Kind        INTEGER  NOT NULL DEFAULT -1,
		State       INTEGER  NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(migrationStr)
	LogErr(err)

	rows, err := db.Query(`SELECT * FROM TodoList;`)
	LogErr(err)
	isFreshDB := true
	for rows.Next() {
		isFreshDB = false
		break
	}
	defer rows.Close()

	if isFreshDB {
		queryStr := `INSERT INTO TodoList (Name, Desc, Kind, State) VALUES ('%s', '%s', %d, %d);`
		db.Exec(fmt.Sprintf(queryStr, "learn rust", "learn rust to ascend to a higher plane",
			KindProject, StateTodo))
	}
	// db.Exec(fmt.Sprintf(queryStr, "learn haskell", "learn hs to ascend to an even higher plane",
	// 	KindUncategorized, StateTodo))
}

// AddTodo saves a todo item into the database.
func AddTodo(todoToAdd TodoItem) {
	const queryStr = `INSERT INTO TodoList(Name, Desc, Kind, State)
		VALUES('%s', '%s', %d, %d);`
	_, err := db.Exec(fmt.Sprintf(queryStr, todoToAdd.Name,
		todoToAdd.Desc, todoToAdd.Kind, todoToAdd.State))

	// const queryStr = `INSERT INTO TodoList(Name, Desc, Kind, State)
	// 	VALUES(?, ?, ?, ?);`
	// stmt, err := db.Prepare(queryStr)
	if err != nil {
		log.Println(err)
	}
	// // stmt.Exec(todoToAdd.Name, todoToAdd.Desc, todoToAdd.Kind, todoToAdd.State)
	// defer stmt.Close()
}

// GetAllTodos returns all todos in the DB as a slice of TodoItems
func GetAllTodos() []TodoItem {
	err := db.Ping()
	LogErr(err)
	todos := []TodoItem{}
	rows, err := db.Query(`SELECT * FROM TodoList ORDER BY State`)
	LogErr(err)

	for rows.Next() {
		var newTodo TodoItem
		err = rows.Scan(&newTodo.TodoId, &newTodo.Name,
			&newTodo.Desc, &newTodo.Kind, &newTodo.State)
		LogErr(err)
		todos = append(todos, newTodo)
	}

	defer rows.Close()
	return todos
}

// UpdateTodo will modify a todo of a given ID to match the passed TodoItem
func UpdateTodo(newTodoInfo TodoItem) {
	const queryStr = `UPDATE TodoList SET
			Name=%s,
			Desc=%s,
			Kind=%d,
			State=%d
	WHERE Todo_id=%d;`

	_, err := db.Exec(fmt.Sprintf(queryStr, newTodoInfo.Name,
		newTodoInfo.Desc, newTodoInfo.Kind, newTodoInfo.State,
		newTodoInfo.TodoId))
	LogErr(err)
}

// MarkDone will toggle the state of the todo item with the given ID.
func MarkDone(todo_id uint) {
	queryStr := `SELECT State FROM TodoList WHERE Todo_id=%d;`
	var currentState int
	err := db.QueryRow(fmt.Sprintf(queryStr, todo_id)).Scan(&currentState)
	LogErr(err)

	var newState TodoState
	switch currentState {
	case int(StateTodo):
		newState = StateInProgress
	case int(StateInProgress):
		newState = StateDone
	case int(StateDone):
		newState = StateTodo
	}

	queryStr = `UPDATE TodoList SET State=%d WHERE Todo_id=%d;`
	_, err = db.Exec(fmt.Sprintf(queryStr, newState, todo_id))
	LogErr(err)
}

// DeleteTodo will delete the todo item with the given ID.
func DeleteTodo(todo_id uint) {
	const queryStr = `DELETE FROM TodoList WHERE Todo_id=%d;`
	_, err := db.Exec(fmt.Sprintf(queryStr, todo_id))
	LogErr(err)
}
