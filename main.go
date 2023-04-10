package main

import (
	"todo-list-rest-go/db"
	"todo-list-rest-go/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3030"))
}
