package controllers

import (
	"log"
	"net/http"

	"github.com/AchmadAlli/golang-todo-app/app"
	"github.com/AchmadAlli/golang-todo-app/app/repositories"
	"github.com/AchmadAlli/golang-todo-app/app/services"
	"github.com/AchmadAlli/golang-todo-app/app/utils"
)

var service *services.TodoService

func ListenTodo(app *app.App) {
	log.Println("listening todo controller")
	service = services.CreateService(&repositories.TodoRepo{DB: app.DB})
	route := app.Router.PathPrefix("/todos").Subrouter()

	route.HandleFunc("/", index)
}

func index(writer http.ResponseWriter, req *http.Request) {
	todos, err := service.Index()

	if err != nil {
		utils.ErrorResponse(writer, 400, err.Error())
	}

	utils.Response(writer, http.StatusOK, todos)
}
