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

// IsInProgress returns true if the todo item is currently in progress
func (ti TodoItem) IsInProgress() bool {
	return ti.State == StateInProgress
}

// GetValidKinds returns a slice of all valid kinds of todos
func (ti TodoItem) GetValidKinds() map[int]string {
	todoKinds := make(map[int]string)
	todoKinds[int(KindUncategorized)] = KindUncategorized.String()
	todoKinds[int(KindProject)] = KindProject.String()
	todoKinds[int(KindHomework)] = KindHomework.String()
	todoKinds[int(KindReading)] = KindReading.String()
	todoKinds[int(KindStudy)] = KindStudy.String()
	return todoKinds
}

//end TodoItem

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
	return fmt.Sprintf("%d is not in our list of kinds of todos", tk)
}

// GetAllTodoKinds returns a map of TodoKinds and their string representations.
func GetAllTodoKinds() map[int]string {
	todoKinds := make(map[int]string)
	todoKinds[int(KindUncategorized)] = KindUncategorized.String()
	todoKinds[int(KindProject)] = KindProject.String()
	todoKinds[int(KindHomework)] = KindHomework.String()
	todoKinds[int(KindReading)] = KindReading.String()
	todoKinds[int(KindStudy)] = KindStudy.String()
	return todoKinds
}

//end TodoKind

// TodoState represents the state of a todo item.
// (ie whether it's done, being worked on, or not started)
type TodoState int

const (
	StateInProgress TodoState = iota - 1
	StateTodo
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
	return fmt.Sprintf("%d is not in our list of states of todos", ts)
}

//end TodoState
