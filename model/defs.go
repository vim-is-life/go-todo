package model

import "fmt"

// TODO look at having this later be restricted on how long name and desc can be?
// - maybe name is no more than 200 chars
// TodoItem represents an item in our todo list
type TodoItem struct {
	TodoId uint
	Name   string
	Desc   string
	Kind   TodoKind
	State  TodoState
}

// IsDone returns true if the todo item is done
func (ti TodoItem) IsDone() bool {
	return ti.State == StateDone
}

// TodoKind represents the todo item categories we'll allow
type TodoKind int

const (
	KindUncategorized TodoKind = iota - 1
	KindProject
	KindHomework
	KindReading
	KindStudy
)

// Method to get a string representation of a kind of todo
func (tk TodoKind) String() string {
	switch tk {
	case KindUncategorized:
		return "Uncategorized"
	case KindProject:
		return "Project"
	case KindHomework:
		return "Homework"
	case KindReading:
		return "Reading"
	case KindStudy:
		return "Study"
	}
	// if we're here we somehow didn't fit in the switch
	return fmt.Sprint(tk, " is not in our list of kinds of todos")
}

//end TodoKind

// TodoState represents the state of a todo item.
// (ie whether it's done, being worked on, or not started)
type TodoState int

const (
	StateTodo TodoState = iota
	StateInProgress
	StateDone
)

// Method to get a string representation of a TodoState
func (ts TodoState) String() string {
	switch ts {
	case StateTodo:
		return "Todo"
	case StateInProgress:
		return "In Progress"
	case StateDone:
		return "Done"
	}
	return fmt.Sprint(ts, " is not in our list of states of todos")
}

//end TodoState
