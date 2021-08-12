package main

import (
	"github.com/AchmadAlli/golang-todo-app/app"
	"github.com/AchmadAlli/golang-todo-app/app/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	app := app.App{}
	app.AddDatabase(db)
	app.InitRoutes()
	app.Run()
}
