package main

import (
	controller "github.com/vim-is-life/go-todo/controller"
	model "github.com/vim-is-life/go-todo/model"
)

func main() {
	model.InitDB()
	controller.SetupAndRun()
}
