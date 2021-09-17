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

	// migrate(db)

	app := app.App{}
	app.AddDatabase(db)
	app.InitRoutes()
	listenServices(&app)

	app.Run()
}

func listenServices(app *app.App) {
	controllers.ListenTodo(app)
}

func migrate(db *gorm.DB) {
	database.MigrateTodo(db)
}
