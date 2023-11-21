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
	db, err := sql.Open("sqlite3", "./todolist.sqlite")
	LogErr(err)
	err = db.Ping()
	LogErr(err)

	const migrationStr = `CREATE TABLE IF NOT EXISTS TodoList (
		todo_id     INTEGER  NOT NULL UNIQUE PRIMARY KEY AUTOINCREMENT,
		name        TEXT     NOT NULL UNIQUE,
		desc        TEXT     NOT NULL UNIQUE,
		kind        INTEGER  NOT NULL DEFAULT -1,
		state       INTEGER  NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(migrationStr)
	LogErr(err)

	// queryStr := `INSERT INTO TodoList (name, desc, kind, state)
	// 	VALUES ('%s', '%s', %d, %d);`

	// db.Exec(fmt.Sprintf(queryStr, "learn rust", "learn rust to ascend to a higher plane",
	// 	KindProject, StateTodo))
	// db.Exec(fmt.Sprintf(queryStr, "learn haskell", "learn hs to ascend to an even higher plane",
	// 	KindUncategorized, StateTodo))
}

// AddTodo saves a todo item into the database.
func AddTodo(todoToAdd TodoItem) {
	const queryStr = `INSERT INTO TodoList (name, desc, kind, state)
		VALUES ('%s', '%s', %d, %d);`
	_, err := db.Exec(fmt.Sprintf(queryStr, todoToAdd.name,
		todoToAdd.desc, todoToAdd.kind, todoToAdd.state))
	LogErr(err)
}

// GetAllTodos returns all todos in the DB as a slice of TodoItems
func GetAllTodos() []TodoItem {
	todos := []TodoItem{}
	rows, err := db.Query(`SELECT * FROM TodoList`)
	LogErr(err)
	defer rows.Close()

	for rows.Next() {
		var newTodo TodoItem
		err = rows.Scan(&newTodo.todo_id, &newTodo.name,
			&newTodo.desc, &newTodo.kind, &newTodo.state)
		LogErr(err)
		todos = append(todos, newTodo)
	}

	return todos
}

// UpdateTodo will modify a todo of a given ID to match the passed TodoItem
func UpdateTodo(newTodoInfo TodoItem) {
	const queryStr = `UPDATE TodoList SET
			name=%s,
			desc=%s,
			kind=%d,
			state=%d
	WHERE id=%d;`

	_, err := db.Exec(fmt.Sprintf(queryStr, newTodoInfo.name,
		newTodoInfo.desc, newTodoInfo.kind, newTodoInfo.state,
		newTodoInfo.todo_id))
	LogErr(err)
}

// MarkDone will toggle the state of the todo item with the given ID.
func MarkDone(todo_id uint) {
	queryStr := `SELECT state FROM TodoList WHERE id=%d;`
	var currentState int
	err := db.QueryRow(fmt.Sprintf(queryStr, todo_id)).Scan(&currentState)
	LogErr(err)

	var newState TodoState
	switch currentState {
	case int(StateTodo), int(StateInProgress):
		newState = StateDone
	case int(StateDone):
		newState = StateTodo
	}

	queryStr = `UPDATE TodoList SET state=%d WHERE id=%d;`
	_, err = db.Exec(fmt.Sprintf(queryStr, newState, todo_id))
	LogErr(err)
}

// DeleteTodo will delete the todo item with the given ID.
func DeleteTodo(todo_id uint) {
	const queryStr = `DELETE FROM TodoList WHERE id=%d;`
	_, err := db.Exec(fmt.Sprintf(queryStr, todo_id))
	LogErr(err)
}
