package main

import (
	"github.com/AchmadAlli/golang-todo-app/app"
	"github.com/AchmadAlli/golang-todo-app/app/controllers"
	"github.com/AchmadAlli/golang-todo-app/app/database"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	app := app.App{}
	app.AddDatabase(db)
	app.InitRoutes()
	listenService(&app, db)

	app.Run()
}

func listenService(app *app.App, db *gorm.DB) {
	todoRouteGroup := app.Router.PathPrefix("/todos").Subrouter()
	controllers.ListenTodo(todoRouteGroup)
}
