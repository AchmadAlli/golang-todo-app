package main

import "github.com/AchmadAlli/golang-todo-app/app"

func main() {
	app := app.App{}
	app.InitRoutes()
	app.Run()
}
